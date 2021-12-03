# Json encoder and decoder benchmark

Test Encode for:
1. Build-in encoder
2. https://github.com/mailru/easyjson
3. https://github.com/json-iterator/go

Test Decode with access to one field for:
1. Build-in decoder
2. https://github.com/mailru/easyjson
3. https://github.com/json-iterator/go
4. https://github.com/valyala/fastjson (only one field extracted)
5. https://github.com/tidwall/gjson (only one field extracted)
6. Parallel https://github.com/valyala/fastjson (only one field extracted)
7. Parallel https://github.com/tidwall/gjson (only one field extracted)

Test Delete unnecessary field from json document without parsing it into structure:
1. https://github.com/valyala/fastjson
2. https://github.com/tidwall/sjson
3. https://github.com/tidwall/sjson by setting necessary fields in new document (bad idea)
4. https://github.com/tidwall/gjson by parse to map[interface{}]interface{}, delete fields and Marshal with jsoniter 

```
goos: linux
goarch: amd64
pkg: github.com/satmaelstorm/jsonBenchmark
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkStdJsonEncode-12                 455626              2907 ns/op            1112 B/op          6 allocs/op
BenchmarkEasyJsonEncode-12               1000000              1325 ns/op             968 B/op          6 allocs/op
BenchmarkJsoniterJsonEncode-12           1000000              1960 ns/op            1112 B/op          6 allocs/op
BenchmarkStdJsonDecode-12                 228279              5276 ns/op             336 B/op         13 allocs/op
BenchmarkEasyJsonDecode-12                663830              1794 ns/op              42 B/op          8 allocs/op
BenchmarkJsoniterDecode-12                366958              3283 ns/op             664 B/op         29 allocs/op
BenchmarkFastJsonDecode-12               1634101               730.6 ns/op             0 B/op          0 allocs/op
BenchmarkGJsonDecode-12                  3851779               318.0 ns/op             3 B/op          1 allocs/op
BenchmarkFastJsonDecodeParallel-12       7178468               170.7 ns/op             0 B/op          0 allocs/op
BenchmarkGJsonDecodeParallel-12         17168725                69.63 ns/op            3 B/op          1 allocs/op
BenchmarkFastJsonDelete-12                426006              5862 ns/op            3916 B/op          9 allocs/op
BenchmarkSJsonDelete-12                    52360             43182 ns/op           38096 B/op         20 allocs/op
BenchmarkSJsonDeleteBySet-12               12549             97480 ns/op           37584 B/op         40 allocs/op
BenchmarkGJsonDeleteByMap-12               47029             35157 ns/op           19958 B/op         46 allocs/op
```

*IMHO*:
1. If you need to encode or decode into a structure, and you want to avoid code generation, take **jsoniter**, otherwise, an **easyjson** one is a great way
2. If you need a small number of fields from a clear location, you can take **gjson**
3. If you need to remove fields from the json without decoding it - **fastjson** will do the best

