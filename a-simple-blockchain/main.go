package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targeBits = 24

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targeBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targeBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func IntToHex(i int64) []byte {
	return []byte(fmt.Sprintf("%x", i))
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containint \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\n\n")

	return nonce, hash[:]
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to sasaxie")
	bc.AddBlock("Send 2  more BTC to sasaxie")

	for _, block := range bc.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %x\n", block.Nonce)
		fmt.Println()
	}
}

/*
Mining the block containint "Genesis Block"
000000f882b9c702e19abfd62109cdeb15d05bc255698a2d78891d0b6ab840cf

Mining the block containint "Send 1 BTC to sasaxie"
00000007dc10cdc3285c3d3c3649cfb5ca6d8e886d1f69689beaee3fd1656800

Mining the block containint "Send 2  more BTC to sasaxie"
000000e34077705ab88f91905a23558bfce4edb9fd4839783291b893c726490e

Prev Hash:
Data: Genesis Block
Hash: 000000f882b9c702e19abfd62109cdeb15d05bc255698a2d78891d0b6ab840cf
Nonce: f133bb

Prev Hash: 000000f882b9c702e19abfd62109cdeb15d05bc255698a2d78891d0b6ab840cf
Data: Send 1 BTC to sasaxie
Hash: 00000007dc10cdc3285c3d3c3649cfb5ca6d8e886d1f69689beaee3fd1656800
Nonce: 1ba237

Prev Hash: 00000007dc10cdc3285c3d3c3649cfb5ca6d8e886d1f69689beaee3fd1656800
Data: Send 2  more BTC to sasaxie
Hash: 000000e34077705ab88f91905a23558bfce4edb9fd4839783291b893c726490e
Nonce: f09505
*/
