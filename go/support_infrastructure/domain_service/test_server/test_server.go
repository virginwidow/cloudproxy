// Copyright (c) 2016, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/x509"
	"flag"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/jlmucb/cloudproxy/go/tao"
	"github.com/jlmucb/cloudproxy/go/tao/auth"

	"github.com/jlmucb/cloudproxy/go/support_infrastructure/domain_service"
)

var hostName = &auth.Prin{
	Type:    "program",
	KeyHash: auth.Str("hostHash")}

var programName = &auth.Prin{
	Type:    "program",
	KeyHash: auth.Str("programHash")}

var network = flag.String("network", "tcp", "The network to use for connections")
var addr = flag.String("addr", "localhost:8124", "The address to listen on")

var domainPass = flag.String("password", "xxx", "The domain password")
var configPath = flag.String("config", "../server/tao.config", "The Tao domain config")

func main() {
	domain, err := tao.LoadDomain(*configPath, []byte(*domainPass))
	if domain == nil {
		log.Printf("domainserver: no domain path - %s, pass - %s, err - %s\n",
			*configPath, *domainPass, err)
	} else if err != nil {
		log.Printf("domainserver: Couldn't load the config path %s: %s\n",
			*configPath, err)
	}
	policyKey, policyCert := domain.Keys, domain.Keys.Cert
	if policyCert == nil {
		log.Fatalln("Policy cert not found")
	}
	hostKey, hostAtt := generateAttestation(policyKey, hostName)
	programKey, programAtt := generateAttestation(hostKey, programName)
	rawEnd1, err := proto.Marshal(hostAtt)
	if err != nil {
		log.Fatalln("Error serializing attestation.")
	}
	programAtt.SerializedEndorsements = [][]byte{rawEnd1}
	cert, err := domain_service.RequestProgramCert(programAtt, programKey.VerifyingKey,
		*network, *addr)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	rootCerts := x509.NewCertPool()
	rootCerts.AddCert(domain.Keys.Cert)
	options := x509.VerifyOptions{Roots: rootCerts}
	_, err = cert.Verify(options)
	if err != nil {
		log.Fatalln("Program cert fails verification check.", err)
	}
	ver, err := tao.VerifierFromX509(cert)
	if err != nil {
		log.Fatalln("Error getting verifier from Program cert", err)
	}
	if v := programKey.VerifyingKey; !v.KeyEqual(cert) {
		log.Fatalf("Key in Program cert %v differs from expected value %v.", ver, v)
	}

	// Test Certificate Revocation.
	serialNumber := cert.SerialNumber
	says := auth.Says{
		Speaker: domain.Keys.SigningKey.ToPrincipal(),
		Message: auth.Pred{
			Name: "revoke",
			Arg:  []auth.Term{auth.Bytes(serialNumber.Bytes())}}}

	att, err := tao.GenerateAttestation(domain.Keys.SigningKey, nil, says)
	if err != nil {
		log.Fatalln("Error generating attestation for certificate revocation.")
	}
	err = domain_service.RequestRevokeCertificate(att, *network, *addr)
	if err != nil {
		log.Fatalln("Error revoking certificate: ", err)
	}
	crl, err := domain_service.RequestCrl(*network, *addr)

	if err != nil {
		log.Fatalln("Error getting CRL: ", err)
	}
	revokedCerts := crl.TBSCertList.RevokedCertificates
	if len(revokedCerts) != 1 {
		log.Fatalf("Revoked 1 cert and got back CRL with %v revoked certs", len(revokedCerts))
	}
	if num := revokedCerts[0].SerialNumber.Int64(); num != serialNumber.Int64() {
		log.Fatalf("Serial number %v doesnt match expected value %v", num, serialNumber)
	}
	log.Println("YAY!")
}

func generateEndorsementCertficate(policyKey *tao.Keys, policyCert *x509.Certificate) (*tao.Keys,
	*x509.Certificate) {
	k, err := tao.NewTemporaryKeys(tao.Signing)
	if k == nil || err != nil {
		log.Fatalln("Can't generate signing key")
	}
	us := "US"
	google := "Google"
	machineName := "Encoded Machine Information"
	details := tao.X509Details{
		Country:      &us,
		Organization: &google,
		CommonName:   &machineName}
	subject := tao.NewX509Name(&details)
	pkInt := tao.PublicKeyAlgFromSignerAlg(*policyKey.SigningKey.Header.KeyType)
	sigInt := tao.SignatureAlgFromSignerAlg(*policyKey.SigningKey.Header.KeyType)
	cert, err := policyKey.SigningKey.CreateSignedX509(
		policyCert, 0, k.SigningKey.GetVerifierFromSigner(),
		pkInt, sigInt, subject)
	if err != nil {
		log.Fatalln(err)
	}
	return k, cert
}

func generateAttestation(signingKey *tao.Keys, delegator *auth.Prin) (*tao.Keys, *tao.Attestation) {
	k, err := tao.NewTemporaryKeys(tao.Signing)
	if k == nil || err != nil {
		log.Fatalln("Can't generate signing key")
	}
	speaksFor := &auth.Speaksfor{
		Delegate:  k.SigningKey.ToPrincipal(),
		Delegator: delegator,
	}
	says := &auth.Says{
		Speaker:    signingKey.SigningKey.ToPrincipal(),
		Time:       nil,
		Expiration: nil,
		Message:    speaksFor,
	}
	att, err := tao.GenerateAttestation(signingKey.SigningKey, nil, *says)
	if err != nil {
		log.Fatalln(err)
	}
	return k, att
}
