package main

import (
	"C"
	"github.com/dgraph-io/dgraph/gql"
	"encoding/json"
)

//export Parse
func Parse(query *C.char, variables_json *C.char) *C.char {
	var variables map[string]string
	json.Unmarshal([]byte(C.GoString(variables_json)), &variables)
	parsed, err := gql.Parse(gql.Request{ Str: C.GoString(query), Variables: variables })
	if err != nil {
		json_bytes, err := json.Marshal(map[string]string{ "error": err.Error() })
		if err != nil { panic(err) }
		return C.CString(string(json_bytes))
	}
	json_bytes, err := json.Marshal(parsed)
	if err != nil {
		json_bytes, err := json.Marshal(map[string]string{ "error": err.Error() })
		if err != nil { panic(err) }
		return C.CString(string(json_bytes))
	}
	return C.CString(string(json_bytes))
}

func main() {}
