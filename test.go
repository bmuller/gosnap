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
	"encoding/json"
	"strings"
	"strconv"
)

type Pathable interface {
	Set(key string, value interface{}) Pathable
	ToJson() (string, error)
}

type Array []interface{}

type Obj map[string]interface{}

func toJson(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

func (a Array) ToJson() (string, error) {
	return toJson(a)
}

func (o Obj) ToJson() (string, error) {
	return toJson(o)
}

func (o Obj) Set(key string, value interface{}) Pathable {
	parts := strings.Split(key, ".")
	nkey := strings.Join(parts[1:], ".")

	switch v := anything.(type) {
	case string:
	
	if len(parts) == 1 {
		o[parts[0]] = value
	} else if curr, found := o[parts[0]]; found {

		switch v := curr.(type) {
		case Pathable:
			p := curr.(Pathable)
			o[parts[0]] = p.Set(nkey, value)
		default:
			
	} else if i, err := strconv.Atoi(parts[1]); err == nil && i >= 0 {
		nvalue := make(Array, i+1)
		o[parts[0]] = nvalue.Set(nkey, value)
	} else {
		nvalue := Obj{}
		o[parts[0]] = nvalue.Set(nkey, value)
	}
	return o
}

func (a Array) Set(key string, value interface{}) Pathable {
	parts := strings.Split(key, ".")
	nkey := strings.Join(parts[1:], ".")

	index, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("Cannot index array with value: " + parts[0])
	}

	if index >= len(a) {
		n := make(Array, index+1)
		for i:=0; i<index; i++ {
			if i < len(a) {
				n[i] = a[i]
			} else {
				n[i] = nil
			}
		}
		a = n
	}

	if len(parts) == 1 {
		a[index] = value
	} else if a[index] != nil {
		p := a[index].(Pathable)
		a[index] = p.Set(nkey, value)
	} else if i, err := strconv.Atoi(parts[1]); err == nil && i >= 0 {
		nvalue := make(Array, i+1)
		a[index] = nvalue.Set(nkey, value)
	} else {
		nvalue := Obj{}
		a[index] = nvalue.Set(nkey, value)
	}
	return a
}

func main() {
	//blah := Array{Obj{"one": Obj{"two": 3}}, Obj{"three": false}}	
	//blah.Set("facebook_entry.properties.publisher_name", Obj{"index": "not_analyzed", "type": "string"})
	//blah.Set("facebook_entry.properties.source_url", Obj{"index": "not_analyzed", "type": "string"})
	//blah.Set("facebook_entry.properties.facebook_performance.properties.page_insights.type", "nested")
	blah := Obj{}
	blah.Set("one.two", 1)
	blah.Set("one.two.three", "one hundred")	
	//blah.Set("three.four.five.six", "five")
	blah.Set("one.two", Obj{"one": 1})	
	if r, err := blah.ToJson(); err == nil {
		fmt.Println(r)
	} else {
		fmt.Println("Error: ", err)
	}
}
