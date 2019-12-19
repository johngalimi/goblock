package main

import (
	"fmt"
	"hash/fnv"
	"encoding/json"
)

type M map[string]int

type BlockContents struct {
	BlockNumber int `json:"blockNumber"`
	ParentHash uint32 `json:"parentHash"`
	TransactionCount int `json:"transactionCount"`
	Transactions []M `json:"transactions"`
}

func convertBlock(block BlockContents) string {
	bytes, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func hashBlock(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	var testTransactions []M

	testTransaction1 := M{"john":5, "james":-5}
	testTransaction2 := M{"jill":-7, "joe":7}

	testTransactions = append(testTransactions, testTransaction1, testTransaction2)

	testBlock :=  BlockContents{
		BlockNumber: 1,
		ParentHash: 12345,
		TransactionCount: 7,
		Transactions: testTransactions,
	}

	convertedblock := convertBlock(testBlock)
	hashedblock := hashBlock(convertedblock)

	fmt.Println(convertedblock)
	fmt.Println(hashedblock)
}
