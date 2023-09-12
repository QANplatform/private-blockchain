#!/bin/sh

../bin/qvmctl call -r QVM_INIT_MAXUSER -m 536870912 -privkey ../data/a/privkey -rpc http://localhost:8545 0x3E004D3de2B9Bbda912CF2C63db4312683449fa9 register qvml2
