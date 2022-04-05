package jsonBenchmark

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/tidwall/gjson"
	"testing"
)

func BenchmarkStdJsonEncode(b *testing.B) {
	var res []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ = json.Marshal(sampleResponse)
	}
	_ = res
}

func BenchmarkEasyJsonEncode(b *testing.B) {
	var res []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ = easyjson.Marshal(sampleResponse)
	}
	_ = res
}

func BenchmarkFFJsonEncode(b *testing.B) {
	var res []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ = ffjson.MarshalFast(&ffSampleResponse)
	}
	_ = res
}

func BenchmarkJsoniterJsonEncode(b *testing.B) {
	var res []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ = jsoniter.Marshal(sampleResponse)
	}
	_ = res
}

func BenchmarkStdJsonDecode(b *testing.B) {
	var result ApiAnswer
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(sampleJson, &result)
		needId = result.Response.Rows[0].AddInfo.Id
	}
	_ = needId
}

func BenchmarkEasyJsonDecode(b *testing.B) {
	var result ApiAnswer
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(sampleJson, &result)
		needId = result.Response.Rows[0].AddInfo.Id
	}
	_ = needId
}

func BenchmarkFFJsonDecode(b *testing.B) {
	var result FFApiAnswer
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ffjson.UnmarshalFast(sampleJson, &result)
		needId = result.Response.Rows[0].AddInfo.Id
	}
	_ = needId
}

func BenchmarkJsoniterDecode(b *testing.B) {
	var result ApiAnswer
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = jsoniter.Unmarshal(sampleJson, &result)
		needId = result.Response.Rows[0].AddInfo.Id
	}
	_ = needId
}

func BenchmarkFastJsonDecode(b *testing.B) {
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := fastJsonParserPool.Get()
		obj, _ := parser.ParseBytes(sampleJson)
		needId = obj.GetInt("response", "rows", "0", "add_info", "id")
		fastJsonParserPool.Put(parser)
	}
	_ = needId
}

func BenchmarkGJsonDecode(b *testing.B) {
	var needId int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needId = int(gjson.GetBytes(sampleJson, "response.rows.0.add_info.id").Int())
	}
	_ = needId
}

func BenchmarkFastJsonDecodeParallel(b *testing.B) {
	var needId int
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			parser := fastJsonParserPool.Get()
			obj, _ := parser.ParseBytes(sampleJson)
			needId = obj.GetInt("response", "rows", "0", "add_info", "id")
			fastJsonParserPool.Put(parser)
		}
	})
	_ = needId
}

func BenchmarkGJsonDecodeParallel(b *testing.B) {
	var needId int
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			needId = int(gjson.GetBytes(sampleJson, "response.rows.0.add_info.id").Int())
		}
	})
	_ = needId
}

func BenchmarkFastJsonDelete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteUnnecessaryFieldsFastJson(map[int][]byte{44866196: sampleObject}, skipFields)
	}
}

func BenchmarkSJsonDelete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteUnnecessaryFieldsSJson(map[int][]byte{44866196: sampleObject}, skipFields)
	}
}

func BenchmarkSJsonDeleteBySet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteUnnecessaryFieldsSJsonBySet(map[int][]byte{44866196: sampleObject}, fields)
	}
}

func BenchmarkGJsonDeleteByMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteUnnecessaryFieldsGJsonMap(map[int][]byte{44866196: sampleObject}, skipFields)
	}
}

const needId = 135

func TestJsonDecode(t *testing.T) {
	var resultStd ApiAnswer
	var resultEasy ApiAnswer

	_ = json.Unmarshal(sampleJson, &resultStd)
	needIdStd := resultStd.Response.Rows[0].AddInfo.Id

	_ = easyjson.Unmarshal(sampleJson, &resultEasy)
	needIdEasyJson := resultEasy.Response.Rows[0].AddInfo.Id

	needIdGJson := int(gjson.GetBytes(sampleJson, "response.rows.0.add_info.id").Int())
	parser := fastJsonParserPool.Get()
	obj, _ := parser.ParseBytes(sampleJson)
	needIdFastJson := obj.GetInt("response", "rows", "0", "add_info", "id")
	fastJsonParserPool.Put(parser)

	if needIdStd != needId {
		t.Errorf("needIdStd invalid!")
	}

	if needIdEasyJson != needId {
		t.Errorf("needIdEasyJson invalid!")
	}

	if needIdGJson != needId {
		t.Errorf("needIdGJson invalid!")
	}

	if needIdFastJson != needId {
		t.Errorf("needIdFastJson invalid!")
	}
}
