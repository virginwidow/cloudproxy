//  File: linux_process_factory_unittests.cc
//  Author: Tom Roeder <tmroeder@google.com>
//
//  Description: Unit tests for LinuxProcessFactory.
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
#include "tao/linux_process_factory.h"

#include <unistd.h>
#include <signal.h>

#include <gflags/gflags.h>
#include <glog/logging.h>
#include <gtest/gtest.h>

#include "tao/fd_message_channel.h"
#include "tao/pipe_factory.h"
#include "tao/util.h"

using namespace tao;

DEFINE_string(short_program, "/bin/true",
              "A short program to execute, "
              "preferably one that will stop on its own");

DEFINE_string(long_program, "/bin/sleep",
              "A long program to execute, "
              "preferably one that will run for a while");

DEFINE_string(long_program_arg, "5", "An argument for the long program");

class LinuxProcessFactoryTest : public ::testing::Test {
 protected:
  virtual void SetUp() {
    path_ = FLAGS_short_program;
    long_path_ = FLAGS_long_program;
    long_arg_ = FLAGS_long_program_arg;
  }
  LinuxProcessFactory factory_;
  string path_, long_path_, long_arg_;
};

TEST_F(LinuxProcessFactoryTest, SubprinTest) {
  string subprin0, subprin1;
  ASSERT_TRUE(factory_.MakeHostedProgramSubprin(0, path_, &subprin0));
  ASSERT_TRUE(factory_.MakeHostedProgramSubprin(1, path_, &subprin1));
  // subprin1 should include the id
  EXPECT_TRUE(subprin0.size() < subprin1.size());

  int id0, id1;
  string hash0, hash1;
  string ext0, ext1;

  ASSERT_TRUE(
      factory_.ParseHostedProgramSubprin(subprin0, &id0, &hash0, &ext0));
  ASSERT_TRUE(
      factory_.ParseHostedProgramSubprin(subprin1, &id1, &hash1, &ext1));
  EXPECT_EQ(id0, 0);
  EXPECT_EQ(id1, 1);
  EXPECT_EQ(hash0, hash1);
  EXPECT_EQ("", ext0);
  EXPECT_EQ("", ext1);

  ASSERT_TRUE(factory_.ParseHostedProgramSubprin(subprin0 + "::Test1::Test2",
                                                  &id0, &hash0, &ext0));
  EXPECT_EQ(id0, 0);
  EXPECT_EQ(hash0, hash1);
  EXPECT_EQ("Test1::Test2", ext0);
}

TEST_F(LinuxProcessFactoryTest, StartTest) {
  PipeFactory pipe_factory;
  string subprin;
  scoped_ptr<HostedLinuxProcess> child;
  ASSERT_TRUE(factory_.MakeHostedProgramSubprin(0, path_, &subprin));
  ASSERT_TRUE(factory_.StartHostedProgram(pipe_factory, path_, list<string>{},
                                           subprin, &child));
  EXPECT_TRUE(child->pid > 0);
  EXPECT_EQ(subprin, child->subprin);
  sleep(1);
  // it should have already stopped
  ASSERT_TRUE(factory_.StopHostedProgram(child.get(), SIGTERM));
  EXPECT_EQ(child->pid, 0);
}

TEST_F(LinuxProcessFactoryTest, StartStopTest) {
  PipeFactory pipe_factory;
  string subprin;
  scoped_ptr<HostedLinuxProcess> child;
  ASSERT_TRUE(factory_.MakeHostedProgramSubprin(0, long_path_, &subprin));
  ASSERT_TRUE(factory_.StartHostedProgram(
      pipe_factory, long_path_, list<string>{long_arg_}, subprin, &child));
  EXPECT_TRUE(child->pid > 0);
  EXPECT_EQ(subprin, child->subprin);
  sleep(1);
  // it should still be running
  ASSERT_TRUE(factory_.StopHostedProgram(child.get(), SIGTERM));
  EXPECT_EQ(child->pid, 0);
}
