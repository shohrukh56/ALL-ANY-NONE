package main

import (
	"fmt"
)

var items = []account{
	{id: 1, name: "Sasha", balance: 5000},
	{id: 2, name: "Masha", balance: 9000},
	{id: 3, name: "Petya", balance: 15_000},
}
func ExampleIndex_NoResult() {
	result := Index(items, func(item account) bool {
		return item.balance < 5_000
	})
	fmt.Println(result)
	// Output: -1
}

func ExampleIndex_HasResult() {
	result := Index(items, func(item account) bool {
		return item.balance > 5_000
	})
	fmt.Println(result)
	// Output: 1
}

func ExampleAll_True() {
	result := All(items, func(item account) bool {
		return item.balance > 1_000
	})
	fmt.Println(result)
	// Output: true
}
func ExampleAll_False() {
	result := All(items, func(item account) bool {
		return item.balance > 5000_000
	})
	fmt.Println(result)
	// Output: false
}

func ExampleAny_True() {
	result := Any(items, func(item account) bool {
		return item.balance > 9_000
	})
	fmt.Println(result)
	// Output: true
}

func ExampleAny_False() {
	result := Any(items, func(item account) bool {
		return item.balance < 1_000
	})
	fmt.Println(result)
	// Output: false
}

func ExampleNone_True() {
	result := None(items, func(item account) bool {
		return item.balance > 9_000
	})
	fmt.Println(result)
	// Output: false
}
func ExampleNone_False() {
	result := None(items, func(item account) bool {
		return item.balance < 1_000
	})
	fmt.Println(result)
	// Output: true
}

func ExampleFind_NoResult_HasError() {
	result, ok := Find(items, func(item account) bool {
		return item.balance < 1_000
	})

	fmt.Println(result, ok)
	// Output: {0  0} Not found
}

func ExampleFind_HasResult_NoError() {
	result, ok := Find(items, func(item account) bool {
		return item.balance == 15_000
	})
	fmt.Println(result, ok)
	// Output: {3 Petya 15000} <nil>
}
