package entity_test

import (
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/Yushchuk6/total-transactions/entity"
)

func TestFromJSON(t *testing.T) {
	tests := []struct {
		json string
		want entity.EthBlock
	}{
		{
			json: `{
				"result": {
				"number": "0x2a",
				"transactions": [
					{"value": "0x2a"},
					{"value": "0x2a"}
					]
				}
			}`,
			want: entity.EthBlock{
				entity.Result{
					Number: "0x2a",
					Transactions: []entity.Transaction{
						{Value: "0x2a"},
						{Value: "0x2a"},
					},
				},
			},
		},
	}
	for _, test := range tests {
		var eb entity.EthBlock
		reader := strings.NewReader(test.json)
		err := eb.FromJSON(reader)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(eb, test.want) {
			t.Errorf("Error converting jSON to ETHBlock, want %v got %v", test.want, eb)
		}
	}
}

func TestToBlock(t *testing.T) {
	tests := []struct {
		eb   entity.EthBlock
		want entity.Block
	}{
		{
			eb: entity.EthBlock{
				entity.Result{
					Number: "0x2a",
					Transactions: []entity.Transaction{
						{Value: "0x2a"},
						{Value: "0x2a"},
					},
				},
			},
			want: entity.Block{
				ID:           *big.NewInt(42),
				Transactions: 2,
				Amount:       84e-18,
			},
		},
	}
	for _, test := range tests {
		got, err := test.eb.ToBlock()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(*got, test.want) {
			t.Errorf("Error converting jSON to ETHBlock, want %v got %v", test.want, *got)
		}
	}
}
