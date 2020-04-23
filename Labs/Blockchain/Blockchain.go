package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Transaction struct {
	user   string
	amount int
}

type Block struct {
	prevBlock  *Block
	prevHash   string
	nonce      int
	body       []Transaction
	difficulty int
}

var blockChain = []Block{}

func (theBlock *Block) Add(previousBlock *Block) {
	if previousBlock == nil {
		theBlock.prevBlock = nil
		theBlock.prevHash = ""
		return
	}

	theBlock.prevBlock = previousBlock
	theBlock.prevHash = CreateHash(*previousBlock)
}

func CreateHash(b Block) string {
	prefix := strings.Repeat("0", b.difficulty)
	hash := ""
	for {
		plaintext := string(b.nonce)
		for _, value := range b.body {
			plaintext = plaintext + value.user +
				string(value.amount)
		}

		data := sha256.Sum256([]byte(plaintext))
		hash = hex.EncodeToString(data[:])
		if hash[:b.difficulty] == prefix {
			break
		}
		b.nonce++
	}

	return hash
}

// func VerifyBlockchain() bool {
// 	for _, block := range blockChain {
// 		if block.prevHash != "" {
// 			if block.prevHash != CreateHash(*block.prevBlock) {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func main() {

	var block1 Block
	var newT = Transaction{user: "Bob", amount: 5}
	block1.body = append(block1.body, newT)
	block1.difficulty = 3
	block1.Add(nil)
	blockChain = append(blockChain, block1)
	var block2 Block
	newT = Transaction{user: "Bob", amount: 5}
	block2.body = append(block2.body, newT)
	newT = Transaction{user: "Carol", amount: 15}
	block2.body = append(block2.body, newT)
	block2.difficulty = 3
	block2.Add(&block1)
	blockChain = append(blockChain, block2)

	fmt.Println(blockChain)

}
