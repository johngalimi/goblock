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

func hashBlock(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func convertBlock(block BlockContents) string {
	bytes, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func main() {
	testBlock :=  BlockContents{
		BlockNumber: 1,
		ParentHash: 12345,
		TransactionCount: 7,
	}
	hashedblock := hashBlock(convertBlock(testBlock))
	fmt.Println(hashedblock)
}
