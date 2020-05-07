package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	queueContract := new(QueueContract)
	queueContract.BeforeTransaction = GetWorldState
	queueContract.TransactionContextHandler = new(CustomTransactionContext)
	queueContract.UnknownTransaction = UnknownTransactionHandler

	cc, err := contractapi.NewChaincode(queueContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
