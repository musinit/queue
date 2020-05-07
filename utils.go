package main

import (
	"errors"
	"math/rand"
	"time"
)

// GetWorldState takes the first transaction arg as the key and sets
// what is found in the world state for that key in the transaction context
func GetWorldState(ctx CustomTransactionContextInterface) error {
	existing, err := ctx.GetStub().GetState(queueGlobalKey)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	ctx.SetData(existing)

	return nil
}

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000)
}
