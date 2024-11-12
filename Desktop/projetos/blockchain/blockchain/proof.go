package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// TAKE THE DATA FROM THE BLOCK
// CREATE A COUNTER NAMED NONCE WHICH STARTS IN 0
// CREATE A HASH OF THE DATA + THE COUNTER
// CHECK IF THE HASH MEETS A SET OF REQUIREMENTS
// REQUIREMENTS: THE FIRST FEW BYTES MUST CONTAIN 0s

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
