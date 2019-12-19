package main

import (
	"fmt"
	"hash/fnv"
	"encoding/json"
)

type BlockContents struct {
	BlockNumber int `json:"blockNumber"`
	ParentHash uint32 `json:"parentHash"`
	TransactionCount int `json:"transactionCount"`
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func convert() string {
	bytes, err := json.Marshal(BlockContents{
		BlockNumber: 1,
		ParentHash: 123456789,
		TransactionCount: 7,
	})
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func main() {
	testconv := convert()
	fmt.Println(testconv)
	fmt.Printf("%T\n", testconv)
	fmt.Println(hash(testconv))
}
