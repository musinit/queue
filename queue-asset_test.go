package main

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	queue := QueueAsset{}

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	if size := queue.Length(); size != 3 {
		t.Errorf("wrong count, excepted 3, got %d", queue.Length())
	}
}

func TestDequeue(t *testing.T) {
	queue := QueueAsset{}

	queue.Enqueue("1")
	queue.Enqueue("2")
	element, err := queue.Dequeue()
	if err != nil {
		t.Errorf("error while dequeue: %v", err)
	}

	if size := queue.Length(); size != 1 {
		t.Errorf("wrong count, excepted 1, got %d", queue.Length())
	}

	if val := element.Value; val != "1" {
		t.Errorf("wrong element in dequeue, should be 1, got %s", val)
	}
}

func TestAttachExtraContext(t *testing.T) {
	queue := QueueAsset{}
	element := queue.Enqueue("1")
	contextKey := ExtraContextKey(1)
	contextValue := "testValue"

	err := queue.AttachExtraContext(element.ID, contextKey, contextValue)

	if err != nil {
		t.Errorf("could not attach context to element")
	}
}

func TestFilterByValue(t *testing.T) {
	queue := QueueAsset{}
	queue.Enqueue("h")
	queue.Enqueue("he")
	queue.Enqueue("hello")
	queue.Enqueue("world")

	result, err := queue.FilterByValue("h")

	if err != nil {
		t.Errorf("error while filtering queue")
	}
	if len(result) != 3 {
		t.Errorf("expected 3 number of elements, got %d", len(result))
	}
}

func TestSwap(t *testing.T) {
	queue := QueueAsset{}
	element1 := queue.Enqueue("1")
	element2 := queue.Enqueue("2")

	err := queue.Swap(element1.ID, element2.ID)
	if err != nil {
		t.Errorf("error while swapping queue")
	}

	result, err := queue.Dequeue()
	if err != nil {
		t.Errorf("error while dequeue")
	}
	if result.Value != element2.Value {
		t.Errorf("element 1 and element 2 didn't swapped")
	}
}

func TestFindElement(t *testing.T) {
	queue := QueueAsset{}
	queue.Enqueue("1")
	element2 := queue.Enqueue("2")

	_, result := queue.FindElement(element2.ID)

	if result.Value != element2.Value {
		t.Errorf("got wrong element by it's ID")
	}
}
