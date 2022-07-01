#!/bin/bash

rm -rf go/proto
protoc -I . \
  --go_out .  \
  --go-grpc_out .  \
  ./proto/*
mv github.com/ICBench/icb-message-agent go/proto
rm -rf github.com
