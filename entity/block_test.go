// +build entity_block
package entity_test

import (
	"bytes"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/Yushchuk6/total-transactions/entity"
)

type BlockJSONTest struct {
	b    entity.Block
	json string
}

var blockJSONTests = []BlockJSONTest{
	{entity.Block{}, `{"transactions":0,"amount":0}`},
	{entity.Block{*big.NewInt(42), 42, 42}, `{"transactions":42,"amount":42}`},
	{entity.Block{*big.NewInt(math.MaxInt64), math.MaxUint32, math.MaxFloat64}, `{"transactions":4294967295,"amount":1.7976931348623157e+308}`},
}

func TestToJSON(t *testing.T) {
	for _, blockTest := range blockJSONTests {
		var bb bytes.Buffer
		blockTest.b.ToJSON(&bb)

		b := bb.Bytes()
		got := string(b[:len(b)-1])
		if got != blockTest.json {
			t.Errorf("Error converting Block %v to JSON %s, got %s", blockTest.b, blockTest.json, got)
		}
	}
}

func TestNewBlock(t *testing.T) {
	b, err := entity.NewBlock(*big.NewInt(1), 2, 3.5)
	if err != nil {
		t.Error(err)
	}
	want := entity.Block{*big.NewInt(1), 2, 3.5}
	if !reflect.DeepEqual(want, *b) {
		t.Errorf("Error creating Block, want %v got %v", want, *b)
	}
}

func TestNewBlockError(t *testing.T) {
	b, err := entity.NewBlock(*big.NewInt(0), 0, -1)
	if err == nil || b != nil {
		t.Errorf("Error hadling negative numbers to Block result:%v error:%v", b, err)
	}
}
