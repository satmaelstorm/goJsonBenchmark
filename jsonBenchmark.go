//go:generate go run github.com/mailru/easyjson/easyjson -pkg
package jsonBenchmark

import (
	_ "embed"
	"encoding/json"
	"github.com/valyala/fastjson"
)

//easyjson:json
type ApiAnswer struct {
	Response Response `json:"response"`
}

//easyjson:json
type Response struct {
	Rows []RowElement `json:"rows"`
}

//easyjson:json
type RowElement struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	AddInfo     AddInfo `json:"add_info"`
	RecordCount int     `json:"record_count"`
}

//easyjson:json
type AddInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

//go:embed test.json
var sampleJson []byte

//go:embed object.json
var sampleObject []byte

var sampleResponse ApiAnswer

var skipFields = []string{"id", "short_description", "inner_object_two", "int_field2", "pi"}
var fields = []string{"title", "inner_object", "description", "time_field", "bool_field", "int_field", "e"}

var fastJsonParserPool fastjson.ParserPool

func init() {
	_ = json.Unmarshal(sampleJson, &sampleResponse)
}
