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
pkg: github.com/satmaelstorm/goJsonBenchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkStdJsonEncode-8                  420640              2524 ns/op            1112 B/op          6 allocs/op
BenchmarkEasyJsonEncode-8                1000000              1003 ns/op             968 B/op          6 allocs/op
BenchmarkFFJsonEncode-8                   732212              1596 ns/op             528 B/op         16 allocs/op
BenchmarkJsoniterJsonEncode-8            1314940               981.2 ns/op          1112 B/op          6 allocs/op
BenchmarkStdJsonDecode-8                  180108              6610 ns/op             304 B/op         13 allocs/op
BenchmarkEasyJsonDecode-8                 633112              2105 ns/op              42 B/op          8 allocs/op
BenchmarkFFJsonDecode-8                   323800              3143 ns/op             642 B/op         15 allocs/op
BenchmarkJsoniterDecode-8                 292714              4150 ns/op             664 B/op         29 allocs/op
BenchmarkFastJsonDecode-8                1234142               951.5 ns/op             0 B/op          0 allocs/op
BenchmarkGJsonDecode-8                   2475152               465.9 ns/op             3 B/op          1 allocs/op
BenchmarkFastJsonDecodeParallel-8        3799425               338.0 ns/op             0 B/op          0 allocs/op
BenchmarkGJsonDecodeParallel-8           7916830               154.1 ns/op             3 B/op          1 allocs/op
BenchmarkFastJsonDelete-8                 275527              3672 ns/op            3914 B/op          9 allocs/op
BenchmarkSJsonDelete-8                     38605             29378 ns/op           38096 B/op         20 allocs/op
BenchmarkSJsonDeleteBySet-8                15432             78291 ns/op           37584 B/op         40 allocs/op
BenchmarkGJsonDeleteByMap-8                41868             34799 ns/op           20013 B/op         48 allocs/op
```

*IMHO*:
1. If you need to encode or decode into a structure, and you want to avoid code generation, take **jsoniter**, otherwise, an **easyjson** one is a great way
2. If you need a small number of fields from a clear location, you can take **gjson**
3. If you need to remove fields from the json without decoding it - **fastjson** will do the best

