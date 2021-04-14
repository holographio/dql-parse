package main

import (
	"C"
	"github.com/dgraph-io/dgraph/gql"
	"encoding/json"
	"fmt"
)

//export Parse
func Parse(query *C.char, variables_json *C.char, result_pointer **C.char) C.int {
	var variables map[string]string
	json.Unmarshal([]byte(C.GoString(variables_json)), &variables)
	parsed, err := gql.Parse(gql.Request{ Str: C.GoString(query), Variables: variables })
	// if err != nil { panic(err) }
	// if err != nil { return C.CString("error") }
	// if err != nil { return "error" }
	if err != nil { fmt.Println("#error", err) } // TODO:
	json_bytes, err := json.Marshal(parsed)
	if err != nil { fmt.Println("#error", err) } // TODO:
	s := string(json_bytes)
	length := len(s)
	*result_pointer = C.CString(s)
	return C.int(length)
}

func main() {}
