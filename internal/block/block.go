package block

import (
	"time"
	"encoding/json"
)

type Block struct{
	Timestamp time.Time
	LastHash string
	Hash string
	Data interface{}
}

func (b * Block) ToString() string {
	if b == nil {
		return "nil object"
	}

	var marshalled []byte
	var err error
	if marshalled, err = json.Marshal(b); err != nil {
		return err.Error()
	}

	return string(marshalled)
}

func New(t time.Time, lastHash, hash string, data interface{}) Block {
	return Block{
		Timestamp: t,
		LastHash: lastHash,
		Hash: hash,
		Data: data,
	}
}
