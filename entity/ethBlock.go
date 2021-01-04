package entity

import (
	"encoding/json"
	"io"
	"math/big"
)

//EthBlock struct of json responce form Ethereum API
type EthBlock struct {
	Result Result `json:"result"`
}

type Result struct {
	Number       string        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Value string `json:"value"`
}

//FromJSON decode json from io.Reader
func (eb *EthBlock) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(eb)
}

//ToBlock convert ETHBlock to Block
func (eb EthBlock) ToBlock() (*Block, error) {
	trnArr := eb.Result.Transactions
	var id, amountWei *big.Int = new(big.Int), new(big.Int)
	var amountETH float64
	var transactions uint

	for _, transaction := range trnArr {
		v, err := HexToBigInt(transaction.Value)
		if err != nil {
			return nil, err
		}
		amountWei.Add(amountWei, v)
	}

	amountETH = WeiToETH(amountWei)
	id, err := HexToBigInt(eb.Result.Number)
	if err != nil {
		return nil, err
	}
	transactions = uint(len(trnArr))

	return NewBlock(*id, transactions, amountETH)
}
