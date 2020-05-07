package main

import "errors"

var (
	ErrGlobalQueueNotFound      = errors.New("Global queue coudn't be found")
	ErrGlobalQueueCantBeUpdated = errors.New("Can't update global queue")
	ErrGlobalQueueIsEmpty       = errors.New("Global queue is empty")
	ErrElementNotFound          = errors.New("Element in global key was not found")
)
