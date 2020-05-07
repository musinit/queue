## Queue smart-contract for Fabric network
Test task, that will be removed in a while from here

# Getting started
1. Place queue folder inside fabric-samples/chaincode
2. Launch network from chaincode-docker-devmode as showed in https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html
Then:
    go mod vendor
    go build
    CORE_CHAINCODE_ID_NAME=mycc:0 CORE_PEER_TLS_ENABLED=false ./queue -peer.address peer:7052

4. Request examples:
    peer chaincode invoke -n mycc -c '{"Args":["InitLedger"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["Enqueue", "value 2"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["Dequeue"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["Length"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["Swap", "ID_1", "ID_2"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["FilterByValue", "beginning of the string"]}' -C myc