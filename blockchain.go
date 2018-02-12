package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	// fmt.Println("******TODO: IMPLEMENT InitBlockchain!******")
	spew.Dump(Blockchain)
	// Fill me in, noble warrior.
	block := Block{"Genesis Block", time.Now().Unix(), []byte{}, []byte{}}
	block.Hash = block.calculateHash()
	Blockchain = append(Blockchain, block)
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	// fmt.Println("******TODO: IMPLEMENT NewBlock!******")
	block := Block{Data: data, Timestamp: time.Now().Unix(), PrevHash: oldBlock.Hash, Hash: []byte{}}
	block.Hash = block.calculateHash()
	// fmt.Println("data: " + block.Data)
	// fmt.Printf("timestamp: %d", block.Timestamp)
	// fmt.Println()
	// fmt.Printf("preHash: %x", block.PrevHash)
	// fmt.Println()
	// fmt.Printf("currentHash: %x", block.Hash)
	// fmt.Println()
	// fmt.Println("******TODO: END NewBlock!******")
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	return block
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	// fmt.Println("******TODO: IMPLEMENT AddBlock!******")
	spew.Dump(Blockchain)
	// Fill me in, brave wizard.
	prevBlock := Blockchain[len(Blockchain)-1]
	// fmt.Println(prevBlock)
	// fmt.Println(b)
	if bytes.Compare(b.PrevHash, prevBlock.Hash) != 0 {
		// return errors.New("Error block")
		return fmt.Errorf("New Block should have Hash: %x, but has Hash: %x",
			prevBlock.Hash, b.PrevHash)
	}
	Blockchain = append(Blockchain, b)
	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
