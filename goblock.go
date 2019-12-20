package main

import (
	"fmt"
	"hash/fnv"
	"encoding/json"
	"math/rand"
	"time"
)

type Transaction map[string]int

type BlockContents struct {
	BlockNumber int `json:"blockNumber"`
	ParentHash uint32 `json:"parentHash"`
	TransactionCount int `json:"transactionCount"`
	Transactions []Transaction `json:"transactions"`
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

func generateTransaction(maxValue int) Transaction {
	rand.Seed(time.Now().UnixNano())

	sign := 1
	if rand.Float32() < 0.5 {
		sign = -sign
	}

	return Transaction{"max":maxValue, "sign":sign}
}

func main() {

	var testTransactions []Transaction

	testTransaction1 := Transaction{"john":5, "james":-5}
	testTransaction2 := Transaction{"jill":-7, "joe":7}

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

	fmt.Println(generateTransaction(100))
}
