package main

import (
	"crypto/sha256"
	"fmt"
)

type Transaction struct {
	user   string
	amount int
}

var blockChain []Block

type Block struct {
	previous     *Block
	previousHash string
	nonce        int
	body         []Transaction
}

func CreateHash(b Block) string {
	plaintext := string(b.nonce)
	for _, k := range b.body {
		plaintext = plaintext + k.user + string(k.amount)
	}
	temp := sha256.Sum256([]byte(plaintext))
	return string(temp[:])
}

func VerifyBlockChain() bool {
	verified := true
	for _, b := range blockChain {
		if b.previousHash == "" {
			continue
		}

		if b.previousHash != CreateHash(*b.previous) {
			verified = false
			continue
		}
	}
	return verified
}

func (theBlock *Block) Add(prevBlock *Block) {
	if prevBlock == nil {
		theBlock.previous = nil
		theBlock.previousHash = ""
		return
	}
	theBlock.previous = prevBlock
	theBlock.previousHash = CreateHash(*prevBlock)
}

func (theBlock *Block) Mine() {
	blockChain = append(blockChain, *theBlock)
}

func main() {

	var block1 Block
	var newT = Transaction{user: "Bob", amount: 5}
	block1.body = append(block1.body, newT)
	block1.Add(nil)
	block1.Mine()

	var block2 Block
	newT = Transaction{user: "Bob", amount: 5}
	block2.body = append(block2.body, newT)
	newT = Transaction{user: "Carol", amount: 15}
	block2.body = append(block2.body, newT)
	block2.Add(&block1)
	block2.Mine()

	//newT = Transaction{user: "Carol", amount: 15}
	//block1.body = append(block1.body, newT)
	//block1.Add(nil)

	if VerifyBlockChain() {
		fmt.Println("verified")
	} else {
		fmt.Println("failed verification")
	}
}
