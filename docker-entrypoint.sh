#!/bin/sh

set -e

# define private and shared data dirs
DATA_PRIV=/data/private
DATA_SHARED=/data/shared

# auto-install dependencies
qanctl install

# write deterministic node key based on account privkey
qanctl util priv2node $DATA_PRIV/privkey > $DATA_PRIV/nodekey

# write self enode to shared dir so other nodes can connect
qanctl util node2enode $DATA_PRIV/nodekey > $DATA_SHARED/$(hostname).enode

# add other participating nodes as peers based on their enode
for NODE in node_a node_b node_c; do
    if test "$NODE" == "$(hostname)"; then
        continue
    fi
    while ! test -f "$DATA_SHARED/$NODE.enode"; do
        echo "Awaiting ENODE of node ${NODE}..."
        sleep 1
    done
    export QAN_STATIC_PEER_${NODE}=$(cat "$DATA_SHARED/$NODE.enode")
done

# launch qand in background
qand &

# print latest confirmed block data in loop (block height + hash + tx count)
qanctl util printblocks http://localhost:8545
