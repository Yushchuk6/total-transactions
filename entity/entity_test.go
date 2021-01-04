package entity_test

import (
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/Yushchuk6/total-transactions/entity"
)

func TestWeiToETH(t *testing.T) {
	tests := []struct {
		wei  int64
		want float64
	}{
		{
			wei:  0,
			want: 0,
		},
		{
			wei:  42,
			want: 0.000000000000000042,
		},
		{
			wei:  math.MaxInt64,
			want: 9.223372036854776,
		},
	}
	for _, test := range tests {
		got := entity.WeiToETH(new(big.Int).SetInt64(test.wei))
		if got != test.want {
			t.Errorf("Error converting Wei to ETH, want %v got %v", test.want, got)
		}
	}
}

type UintHexTest struct {
	num *big.Int
	hex string
}

var uintHexTests = []UintHexTest{
	{big.NewInt(0), "0x0"},
	{big.NewInt(42), "0x2a"},
	{big.NewInt(math.MaxInt64), "0x7fffffffffffffff"},
}

func TestBigIntToHex(t *testing.T) {
	for _, test := range uintHexTests {
		got := entity.BigIntToHex(test.num)
		if got != test.hex {
			t.Errorf("Error converting Uint to HEX, want %v got %v", test.hex, got)
		}
	}
}

func TestHextoBigInt(t *testing.T) {
	for _, test := range uintHexTests {
		got, err := entity.HexToBigInt(test.hex)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(got, test.num) {
			t.Errorf("Error converting HEX to Uint, want %v got %v", test.num, got)
		}
	}
}
