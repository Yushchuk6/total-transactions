package handler

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/Yushchuk6/total-transactions/entity"
	"github.com/julienschmidt/httprouter"
)

func GetTotalByID(client entity.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id, ok := new(big.Int).SetString(params.ByName("id"), 10)
		if !ok {
			log.Println(entity.ErrIncorrectValue)
			return
		}
		tag := entity.BigIntToHex(id)

		m := map[string]string{
			"module":  "proxy",
			"action":  "eth_getBlockByNumber",
			"tag":     tag,
			"boolean": "true",
		}
		client.AddParams(m)

		eb, err := ethGetBlockByNumber(client.URL.String())
		if err == entity.ErrResponveServerNOTOK {
			fmt.Fprintf(w, "APi server responded with NOTOK (Max rate limit reached)")
			return
		} else if err != nil {
			log.Println(err)
			return
		}
		b, err := eb.ToBlock()
		if err != nil {
			log.Println(err)
			return
		}
		b.ToJSON(w)
	})
}

func ethGetBlockByNumber(url string) (*entity.EthBlock, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	eb := new(entity.EthBlock)
	err = eb.FromJSON(resp.Body)
	if err != nil {
		return nil, entity.ErrResponveServerNOTOK
	}
	return eb, nil
}
