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

	amount := 1 + rand.Intn(maxValue - 1)
	aValue := sign * amount

	return Transaction{"party_a": aValue, "party_b": -1 * aValue}
}

func createTransactions(maxValue int, numTransactions int) []Transaction {

	var transactionList []Transaction

	for i := 1; i < numTransactions; i++ {
		transactionList = append(transactionList, generateTransaction(maxValue))
	}

	return transactionList
}

func main() {

	txnList := createTransactions(100, 20)

	testBlock :=  BlockContents{
		BlockNumber: 1,
		ParentHash: 12345,
		TransactionCount: 7,
		Transactions: txnList[:5],
	}

	convertedblock := convertBlock(testBlock)
	hashedblock := hashBlock(convertedblock)

	fmt.Println(convertedblock)
	fmt.Println(hashedblock)
}
