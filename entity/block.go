package entity

import (
	"encoding/json"
	"io"
	"math/big"
)

//Block required data
type Block struct {
	ID           big.Int `json:"-"`
	Transactions uint    `json:"transactions"`
	Amount       float64 `json:"amount"`
}

//NewBlock create ETHBlock
func NewBlock(id big.Int, transactions uint, amount float64) (*Block, error) {
	if amount < 0 {
		return nil, ErrVariableOutOfRange
	}
	b := &Block{
		ID:           id,
		Transactions: transactions,
		Amount:       amount,
	}
	return b, nil
}

//ToJSON encode to json from io.Writer
func (b *Block) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}
