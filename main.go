package main

import (
	"bytes"
	"crypto/sha256"
	_"fmt"
	"strconv"
)

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
