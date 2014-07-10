//  File: util.h
//  Author: Tom Roeder <tmroeder@google.com>
//
// Description: Useful functions for CloudProxy
//
//  Copyright (c) 2013, Google Inc.  All rights reserved.
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
#ifndef CLOUDPROXY_UTIL_H_
#define CLOUDPROXY_UTIL_H_

// #include <stdio.h>

/// These basic utilities from the standard library are used extensively
/// throughout the CloudProxy implementation, so we include them here.
#include <list>
#include <memory>
#include <set>
#include <sstream>
#include <string>

/// These basic utilities from Keyczar and OpenSSL are used extensively
/// throughout the CloudProxy implementation, so we include them here.
#include <keyczar/base/base64w.h>
#include <keyczar/base/basictypes.h>  // DISALLOW_COPY_AND_ASSIGN
#include <keyczar/base/file_util.h>
#include <keyczar/base/scoped_ptr.h>
#include <keyczar/base/values.h>  // for ScopedSafeString
// #include <keyczar/base/stl_util-inl.h>
#include <keyczar/openssl/util.h>
#include <openssl/crypto.h>
#include <openssl/evp.h>
#include <openssl/ssl.h>

#include "tao/util.h"

namespace keyczar {
class Signer;
class Verifier;
}

namespace tao {
class Signer;
}

namespace cloudproxy {
/// These basic utilities from Keyczar and the standard library are used
/// extensively throughout the CloudProxy implementation, so we import them into
/// the cloudproxy namespace here.
/// @{

using std::list;
using std::set;
using std::string;
using std::stringstream;
using std::unique_ptr;  // TODO(kwalsh) Discuss unique_ptr vs. scoped_ptr.
// using std::make_unique;  // TODO(kwalsh) Discuss unique_ptr vs. scoped_ptr.

// using keyczar::base::FilePath;  // Why isn't this in keyczar::base ?
// using keyczar::base::ReadFileToString; // Define our own version below.
using keyczar::base::Base64WDecode;      // NOLINT
using keyczar::base::Base64WEncode;      // NOLINT
using keyczar::base::CreateDirectory;    // NOLINT
using keyczar::base::Delete;             // NOLINT
using keyczar::base::DirectoryExists;    // NOLINT
using keyczar::base::PathExists;         // NOLINT
using keyczar::base::ScopedSafeString;   // NOLINT
using keyczar::base::WriteStringToFile;  // NOLINT

using tao::make_unique;
using tao::CallUnlessNull;

/// @}


/// Clean up an OpenSSL SSL connection.
/// @param ssl The connection to clean up.
void ssl_cleanup(SSL *ssl);

/// A smart pointer to an OpenSSL SSL_CTX.
typedef scoped_ptr_malloc<SSL_CTX, CallUnlessNull<SSL_CTX, SSL_CTX_free>>
    ScopedSSLCtx;

/// A smart pointer to an SSL object.
typedef scoped_ptr_malloc<SSL, CallUnlessNull<SSL, ssl_cleanup>> ScopedSSL;

/// Prepare an SSL_CTX for a server to accepts connections from clients.
/// Peer certificates will be required.
/// @param key The private signing key and x509 certificate to use.
/// @param cert A serialized PEM-format x509 certificate for the key.
/// @param ctx The OpenSSL context to prepare.
bool SetUpSSLServerCtx(const tao::Signer &key, const string &cert,
                       ScopedSSLCtx *ctx);

/// Prepare an SSL_CTX for a client to connect to a server.
/// @param key The private signing key and x509 certificate to use.
/// @param ctx The OpenSSL context to prepare.
bool SetUpSSLClientCtx(const tao::Signer &key, const string &cert,
                       ScopedSSLCtx *ctx);

}  // namespace cloudproxy
#endif  // CLOUDPROXY_UTIL_H_
