#!/bin/sh

../bin/qvmctl deploy -language golang -privkey ../data/a/privkey -rpc http://localhost:8545 ./contracts/sample.go
