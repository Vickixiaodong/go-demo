package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const blocksBucket = "blocksBucket"

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open("chain-data.db", 0600, nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Fatal(err.Error())
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Fatal(err.Error())
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Fatal(err.Error())
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Fatal(err.Error())
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Fatal(err.Error())
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}
