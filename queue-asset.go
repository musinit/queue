package main

import (
	"context"
	"strings"
)

type Element struct {
	ID           int    `json:"id"`
	Value        string `json:"value"`
	ExtraContext interface{}
}

// Queue just to show interface capabilities
type Queue interface {
	Enqueue(value string)
	Dequeue() (*Element, error)
	AttachExtraContext(id int, contentID ExtraContextKey, value interface{}) error
	FilterByValue(beginsWith string) ([]Element, error)
	Swap(leftID, rightID int) error
	Length() int
}

type QueueAsset struct {
	Elements []Element `json:"elements"`
}

func (qa *QueueAsset) Enqueue(value string) *Element {
	element := Element{
		ID:    GenerateID(),
		Value: value,
	}
	qa.Elements = append(qa.Elements, element)
	return &element
}

func (qa *QueueAsset) Dequeue() (*Element, error) {
	if len(qa.Elements) == 0 {
		return nil, ErrGlobalQueueIsEmpty
	}

	element := qa.Elements[0]
	qa.Elements = qa.Elements[1:]

	return &element, nil
}

func (qa *QueueAsset) FindElement(id int) (int, *Element) {
	for i, element := range qa.Elements {
		if element.ID == id {
			return i, &element
		}
	}
	return 0, nil
}

func (qa *QueueAsset) AttachExtraContext(id int, contentID ExtraContextKey, value string) error {
	_, element := qa.FindElement(id)
	if element == nil {
		return ErrElementNotFound
	}

	extraContext := context.WithValue(context.Background(), contentID, value)
	element.ExtraContext = extraContext
	return nil
}

func (qa *QueueAsset) FilterByValue(beginsWith string) ([]Element, error) {
	var result []Element
	elements := qa.Elements
	if len(elements) == 0 {
		return result, nil
	}

	for _, element := range elements {
		if strings.HasPrefix(element.Value, beginsWith) {
			result = append(result, element)
		}
	}

	return result, nil
}

func (qa *QueueAsset) Swap(leftID, rightID int) error {
	leftIndex, leftElement := qa.FindElement(leftID)
	rightIndex, rightElement := qa.FindElement(rightID)
	if leftElement == nil || rightElement == nil {
		return ErrElementNotFound
	}

	qa.Elements[leftIndex], qa.Elements[rightIndex] = qa.Elements[rightIndex], qa.Elements[leftIndex]
	return nil
}

func (qa *QueueAsset) Length() int {
	return len(qa.Elements)
}

// Custom type, because you shouldn't store base type as Key in context value
// https://golang.org/src/context/context.go?s=16100:16160#L503
type ExtraContextKey int
