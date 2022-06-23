#!/bin/bash
rm -rf agent/go/service
protoc -I . \
  --go_out .  \
  --go-grpc_out .  \
  ./proto/*
mv github.com/ICBench/icb-message-agent agent/go/service
rm -rf github.com
