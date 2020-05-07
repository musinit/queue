package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// QueueContract base structure to store global queue
type QueueContract struct {
	contractapi.Contract
}

const (
	// Key for storing Global queue in ledger
	queueGlobalKey = "queueKey"
)

// InitLedger init the queue
func (q *QueueContract) InitLedger(ctx CustomTransactionContextInterface) error {
	queue := QueueAsset{}
	queueAsBytes, _ := json.Marshal(queue)
	err := ctx.GetStub().PutState(queueGlobalKey, queueAsBytes)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	return nil
}

// Enqueue to add element in queue (FIFO)
func (q *QueueContract) Enqueue(ctx CustomTransactionContextInterface, value string) (*QueueAsset, error) {
	var err error
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return nil, err
	}

	queue.Enqueue(value)

	err = updateGlobalQueue(ctx, queue)
	if err != nil {
		return nil, ErrGlobalQueueCantBeUpdated
	}

	return queue, nil
}

// Dequeue to get element from queue (FIFO)
func (q *QueueContract) Dequeue(ctx CustomTransactionContextInterface) (*Element, error) {
	var err error
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return nil, err
	}

	element, err := queue.Dequeue()
	if err != nil {
		return nil, err
	}

	err = updateGlobalQueue(ctx, queue)
	if err != nil {
		return nil, ErrGlobalQueueCantBeUpdated
	}

	return element, nil
}

// AttachExtraContext to attach string content in context inside element (can be anything, but string for simplicity)
func (q *QueueContract) AttachExtraContext(ctx CustomTransactionContextInterface, id int, contentID ExtraContextKey, value string) error {
	var err error
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return err
	}

	err = queue.AttachExtraContext(id, contentID, value)
	if err != nil {
		return err
	}

	err = updateGlobalQueue(ctx, queue)
	if err != nil {
		return ErrGlobalQueueCantBeUpdated
	}
	return nil
}

// Swap to replace two elements in queue
func (q *QueueContract) Swap(ctx CustomTransactionContextInterface, leftID, rightID int) error {
	var err error
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return err
	}

	err = queue.Swap(leftID, rightID)
	if err != nil {
		return nil
	}

	err = updateGlobalQueue(ctx, queue)
	if err != nil {
		return ErrGlobalQueueCantBeUpdated
	}

	return nil
}

// FilterByValue to get N elements, which Values begings with 'beginsWith' string
func (q *QueueContract) FilterByValue(ctx CustomTransactionContextInterface, begingsWith string) ([]Element, error) {
	var err error
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return nil, err
	}

	result, err := queue.FilterByValue(begingsWith)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Length to get length
func (q *QueueContract) Length(ctx CustomTransactionContextInterface) (int, error) {
	queue, err := getGlobalQueue(ctx)
	if err != nil {
		return 0, err
	}

	return queue.Length(), nil
}

// UnknownTransactionHandler returns a shim error
// with details of a bad transaction request
func UnknownTransactionHandler(ctx CustomTransactionContextInterface) error {
	fcn, args := ctx.GetStub().GetFunctionAndParameters()
	return fmt.Errorf("Invalid function %s passed with args %v", fcn, args)
}

// getGlobalQueue to get single global queue from state
func getGlobalQueue(ctx CustomTransactionContextInterface) (*QueueAsset, error) {
	queueAsBytes := ctx.GetData()

	if queueAsBytes == nil {
		return nil, ErrGlobalQueueNotFound
	}

	queue := new(QueueAsset)
	_ = json.Unmarshal(queueAsBytes, queue)

	return queue, nil
}

// updateGlobalQueue to update the state of global queue
func updateGlobalQueue(ctx CustomTransactionContextInterface, queue *QueueAsset) error {
	updatedQueueAsBytes, _ := json.Marshal(queue)

	err := ctx.GetStub().PutState(queueGlobalKey, updatedQueueAsBytes)
	if err != nil {
		return ErrGlobalQueueCantBeUpdated
	}

	return nil
}
