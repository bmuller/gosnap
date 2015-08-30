package main

/*
f := `{ "facebook_entry": {"properties": {
	"publisher_name": {"index": "not_analyzed", "type": "string"},
	"source_url": {"index": "not_analyzed", "type": "string" },
	"facebook_performance": {"properties": {"page_insights": {"type": "nested" } } }
}}}`
*/

import (
	"fmt"
	"strings"
	"errors"
)

type Array []interface{}
type Obj map[string]interface{}

func ToJson(obj interface{}) (string, error) {
	switch obj.(type) {
	case Array:
		r := make([]string, len(obj.(Array)))
		for item, i := range obj.(Array) {
			if j, err := ToJson(item); err == nil {
				//r[i] = j
				fmt.Println(i)
				fmt.Println(j)
			} else {
				return "", err
			}
		}
		return "[" + strings.Join(r, ",") + "]", nil
	}
	return "", errors.New("broke")
}

func main() {
	blah := Array{Obj{"one": Obj{"two": 3}}, Obj{"three": false}}
	fmt.Println(ToJson(blah))
}
