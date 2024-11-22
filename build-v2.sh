#!/usr/bin/env bash

go mod download

go generate .
go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o sfcrypt-v2 github.com/kmou424/sfcrypt/cmd/v2/sfcrypt