#!/bin/bash
rm -rf go/proto
protoc -I . \
  --go_out .  \
  --go-grpc_out .  \
  ./proto/*
mv github.com/ICBench/icb-message-agent go/proto
rm -rf github.com

cd proto
protoc -I . \
 --grpc_out=../cpp/proto \
  --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` \
  send_message.proto

protoc -I . \
  --cpp_out=../cpp/proto \
  send_message.proto
