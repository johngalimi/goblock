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
	hashed := fnv.New32a()
	hashed.Write([]byte(s))

	return hashed.Sum32()
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

func updateAccount(txn Transaction, state Transaction) Transaction {

	stateCopy := state

	for key := range txn {
		_, exists := stateCopy[key]
		if exists {
			stateCopy[key] += txn[key]
		} else {
			stateCopy[key] = txn[key]
		}
	}

	return stateCopy
}

func validateTransaction(txn Transaction, state Transaction) bool {

	sum := 0

	for _, value := range txn {
		sum += value
	}
	if sum != 0 {
		return false
	}

	for key := range txn {

		_, exists := state[key]
		accountBalance := 0

		if exists {
			accountBalance = state[key]
		}

		if (accountBalance - txn[key]) < 0 {
			return false
		}
	}
	return true
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

	state := Transaction{"party_a": 170, "party_b": 95}

	fmt.Println(state)
	fmt.Println(txnList[0])

	if validateTransaction(txnList[0], state) {
		newAcct := updateAccount(txnList[0], state)
		fmt.Println(newAcct)
	} else {
		fmt.Println("failed")
	}
}
