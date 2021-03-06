// Copyright (c) 2015, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mixnet

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/jlmucb/cloudproxy/go/tao"
)

type DirectoryContext struct {
	keys     *tao.Keys     // Signing keys of this hosted program.
	domain   *tao.Domain   // Policy guard and public key.
	listener net.Listener  // Socket where server listens for proxies/routers
	network  string        // Network protocol, e.g. "tcp"
	timeout  time.Duration // Timeout on read/write/dial.

	dirLock    *sync.Mutex
	directory  []string // List of online servers
	serverKeys [][]byte // NaCL keys
}

func NewDirectoryContext(path, network, addr string, timeout time.Duration,
	x509Identity *pkix.Name, t tao.Tao) (*DirectoryContext, error) {
	dc := new(DirectoryContext)
	var err error
	// Generate keys and get attestation from parent.
	if dc.keys, err = tao.NewTemporaryTaoDelegatedKeys(tao.Signing|tao.Crypting, t); err != nil {
		return nil, err
	}

	// Create a certificate.
	pkInt := tao.PublicKeyAlgFromSignerAlg(*dc.keys.SigningKey.Header.KeyType)
	sigInt := tao.SignatureAlgFromSignerAlg(*dc.keys.SigningKey.Header.KeyType)
	dc.keys.Cert, err = dc.keys.SigningKey.CreateSelfSignedX509(pkInt, sigInt, int64(1), x509Identity)
	if err != nil {
		return nil, err
	}

	// Load domain from local configuration.
	if dc.domain, err = tao.LoadDomain(path, nil); err != nil {
		return nil, err
	}

	// Encode TLS certificate.
	cert, err := tao.EncodeTLSCert(dc.keys)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		RootCAs:            x509.NewCertPool(),
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequestClientCert,
	}

	if dc.listener, err = Listen(network, addr, tlsConfig,
		dc.domain.Guard, dc.domain.Keys.VerifyingKey, dc.keys.Delegation); err != nil {
		return nil, err
	}

	dc.network = network
	dc.timeout = timeout

	dc.dirLock = new(sync.Mutex)
	dc.directory = nil

	return dc, nil
}

func (dc *DirectoryContext) Accept() (net.Conn, error) {
	c, err := dc.listener.Accept()
	if err != nil {
		return nil, err
	}
	go dc.handleConn(c, len(c.(*tls.Conn).ConnectionState().PeerCertificates) > 0)
	return c, nil
}

func (dc *DirectoryContext) handleConn(c net.Conn, fromRouter bool) {
	msg := make([]byte, MaxMsgBytes+1)
	n, err := c.Read(msg)
	if err == io.EOF {
		return
	} else if err != nil {
		errMsg := &DirectoryMessage{
			Type:  DirectoryMessageType_DIRERROR.Enum(),
			Error: proto.String(err.Error()),
		}
		ret, err := proto.Marshal(errMsg)
		if err != nil {
			glog.Error(err)
		}
		c.Write(ret)
	} else if n > MaxMsgBytes {
		errMsg := &DirectoryMessage{
			Type:  DirectoryMessageType_DIRERROR.Enum(),
			Error: proto.String("Too many bytes in this message"),
		}
		ret, err := proto.Marshal(errMsg)
		if err != nil {
			glog.Error(err)
		}
		c.Write(ret)
	}

	var dm DirectoryMessage
	if err = proto.Unmarshal(msg[:n], &dm); err != nil {
		errMsg := &DirectoryMessage{
			Type:  DirectoryMessageType_DIRERROR.Enum(),
			Error: proto.String(err.Error()),
		}
		ret, err := proto.Marshal(errMsg)
		if err != nil {
			glog.Error(err)
		}
		c.Write(ret)
	}

	dc.dirLock.Lock()
	defer dc.dirLock.Unlock()
	if *dm.Type == DirectoryMessageType_REGISTER {
		log.Println("Registering", dm.Addrs)
		if fromRouter {
			dc.directory = append(dc.directory, dm.Addrs...)
			dc.serverKeys = append(dc.serverKeys, dm.Keys...)
		}
		_, err = c.Write([]byte{0}) // Indicate it was successfully written
		if err != nil {
			glog.Error(err)
		}
	} else if *dm.Type == DirectoryMessageType_DELETE {
		log.Println("Deleting", dm.Addrs)
		if fromRouter {
			for _, addr := range dm.Addrs {
				for i := range dc.directory {
					if addr == dc.directory[i] {
						dc.directory[i] = dc.directory[len(dc.directory)-1]
						dc.directory = dc.directory[:len(dc.directory)-1]
						dc.serverKeys[i] = dc.serverKeys[len(dc.serverKeys)-1]
						dc.serverKeys = dc.serverKeys[:len(dc.serverKeys)-1]
						break
					}
				}
			}
		}
		_, err = c.Write([]byte{0}) // Indicate it was successfully written
		if err != nil {
			glog.Error(err)
		}
	} else if *dm.Type == DirectoryMessageType_LIST {
		result := &DirectoryMessage{
			Type:  DirectoryMessageType_DIRECTORY.Enum(),
			Addrs: dc.directory,
			Keys:  dc.serverKeys,
		}
		ret, err := proto.Marshal(result)
		if err != nil {
			glog.Error(err)
		}

		n, err := c.Write(ret)
		if err != nil {
			glog.Error(err)
		} else if n != len(ret) {
			glog.Error("Could not send back all of the directory")
		}
	}
}

func (dc *DirectoryContext) Close() {
	if dc.listener != nil {
		dc.listener.Close()
	}
}

func RegisterRouter(c net.Conn, addrs []string, keys [][]byte) error {
	dm := &DirectoryMessage{
		Type:  DirectoryMessageType_REGISTER.Enum(),
		Addrs: addrs,
		Keys:  keys,
	}
	b, err := proto.Marshal(dm)
	if err != nil {
		return err
	}
	n, err := c.Write(b)
	if err != nil {
		return err
	} else if n != len(b) {
		return errors.New("Couldn't write the whole request")
	}
	c.Read([]byte{0})
	return nil
}

func GetDirectory(c net.Conn) ([]string, [][]byte, error) {
	dm := &DirectoryMessage{
		Type: DirectoryMessageType_LIST.Enum(),
	}
	b, err := proto.Marshal(dm)
	if err != nil {
		return nil, nil, err
	}
	n, err := c.Write(b)
	if err != nil {
		return nil, nil, err
	} else if n != len(b) {
		return nil, nil, errors.New("Couldn't write the whole request")
	}

	msg := make([]byte, MaxMsgBytes+1)
	n, err = c.Read(msg)
	if err != nil {
		return nil, nil, err
	} else if n > MaxMsgBytes {
		return nil, nil, errors.New("Couldn't read the whole response")
	}

	var dir DirectoryMessage
	if err = proto.Unmarshal(msg[:n], &dir); err != nil {
		return nil, nil, err
	}

	return dir.Addrs, dir.Keys, err
}
