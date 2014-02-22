//  File: heserver.cc
//  Author: Kevin Walsh <kwalsh@holycross.edu>
//
//  Description: An example http echo server application using HttpEchoServer
//
//  Copyright (c) 2014, Kevin Walsh.  All rights reserved.
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

#include <gflags/gflags.h>
#include <glog/logging.h>

#include "cloudproxy/http_echo_server.h"
#include "tao/util.h"

DEFINE_string(address, "localhost", "The address to listen on");
DEFINE_string(port, "8080", "The port to listen on");

int main(int argc, char **argv) {
  tao::InitializeApp(&argc, &argv, true);

  cloudproxy::HttpEchoServer hes(FLAGS_address, FLAGS_port);

  LOG(INFO) << "HttpEchoServer listening on " << FLAGS_address << ":"
            << FLAGS_port;
  CHECK(hes.Listen(false /* not single channel */))
      << "Could not listen for http connections";
  return 0;
}
