package main

import (
	"fmt"
)

type account struct {
	id      int64
	name    string
	balance int64
}

func main() {
}
func Index(items []account, predicate func(item account) bool) int {
	for index, item := range items {
		if predicate(item) {
			return index
		}
	}
	return -1
}

func All(items []account, predicate func(item account) bool) bool {
	return Index(items, func(item account) bool {
		return !predicate(item)
	}) == -1
}

func Any(items []account, predicate func(item account) bool) bool {
	return Index(items, predicate) != -1
}

func None(items []account, predicate func(item account) bool) bool {
	return Index(items, func(item account) bool {
		return predicate(item)
	}) == -1
}

func Find(items []account, predicate func(item account) bool) (account, error) {
	if result := Index(items, predicate); result != -1 {
		return items[ result] , nil
	}

	return account{}, fmt.Errorf("Not found")
}
