//go:generate go run github.com/pquerna/ffjson $GOFILE
package jsonBenchmark

import "encoding/json"

type FFApiAnswer struct {
	Response FFResponse `json:"response"`
}

type FFResponse struct {
	Rows []FFRowElement `json:"rows"`
}

type FFRowElement struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	AddInfo     FFAddInfo `json:"add_info"`
	RecordCount int       `json:"record_count"`
}

type FFAddInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

var ffSampleResponse FFApiAnswer

func init() {
	_ = json.Unmarshal(sampleJson, &ffSampleResponse)
}
