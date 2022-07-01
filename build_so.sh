#!/bin/bash

cd "$(dirname $(readlink -f $0))/go"

go build -o ../cpp/lib/libmsgagent.so -buildmode=c-shared cgo/send_message_c.go
mv ../cpp/lib/libmsgagent.h ../cpp/include