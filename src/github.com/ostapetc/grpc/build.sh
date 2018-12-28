#!/usr/bin/env bash

/opt/protoc/bin/protoc -I ./ ./app.proto --go_out=plugins=grpc:./