#!/bin/bash
rm -rf agent/go/sendpb
protoc -I . \
  --go_out .  \
  --go-grpc_out .  \
  ./proto/*
mv github.com/ICBench/icb-message-agent agent/go/sendpb
rm -rf github.com
