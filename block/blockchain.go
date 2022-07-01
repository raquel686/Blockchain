package main

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Owner struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	CCI      string `json:"CCI"`
	Age      string `json:"Age"`
}

type Block struct {
	timestamp    time.Time
	transactions []Owner
	prevhash     []byte
	Hash         []byte
}

func main() {
	bytes, _ := os.Open("sample_owners.csv")

	r := csv.NewReader(bytes)

	var owners []Owner

	for {
		attribute, err := r.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		owners = append(owners, Owner{
			Name:     attribute[0],
			LastName: attribute[1],
			CCI:      attribute[2],
			Age:      attribute[3],
		})

	}

	contBlock := 1
	var auxPHash []byte

	for _, o := range owners {
		abc := []Owner{o}
		xyz := Blocks(abc, auxPHash)
		fmt.Println("This is out Block", contBlock)
		Print(xyz)
		auxPHash = xyz.Hash
		contBlock++
	}
}

func Blocks(transactions []Owner, prevhash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevhash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions []Owner, prevhash []byte) []byte {
	input := append(prevhash, time.String()...)
	for transactions := range transactions {
		input = append(input, string(rune(transactions))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block) {
	fmt.Printf("\t time: %s \n", block.timestamp.String())
	fmt.Printf("\t prevhash: %x \n", block.prevhash)
	fmt.Printf("\t hash: %x \n", block.Hash)
	Transaction(block)
}

func Transaction(block *Block) {
	fmt.Println("\t Transactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t %v = Name: %s - LastName: %s - CCI: %s - Age. %s\n", i, transaction.Name, transaction.LastName, transaction.CCI, transaction.Age)
	}
}
