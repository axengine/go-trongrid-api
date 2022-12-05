package go_trongrid_api

import (
	"encoding/json"
	"fmt"
)

func JsonPretty(v interface{}) {
	bz, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bz))
}
