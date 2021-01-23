## Simple Encoding Benchmark

A benchmark comparison of the following:
- standard Go library JSON Marhsal/Unmarshal
- [easyjson](https://github.com/mailru/easyjson) custom Marshal/Unmarshal
- standard Go library gob Encode/Decode

Benchmarks executed on the following setup:  
MacBook Pro (Retina, 13-inch, Mid 2014)  
2,6 GHz Dual-Core Intel Core i5  
8 GB 1600 MHz DDR3  

[Optional] Update easyjson custom encoder/decoder for project structs with  
```bash
cd encoding && easyjson -all encoding.go 
```
Run benchmarks with the following command  
```bash
cd encoding && go test -bench=.
```

Results look like the following
```bash
goos: darwin
goarch: amd64
pkg: github.com/fpapadopou/go-encoding-benchmark/encoding
Benchmark_StdJSONEncode-4                 888596              1321 ns/op
Benchmark_EasyJsonEncode-4               2997316               400 ns/op
Benchmark_StdGobEncode-4                  172783              5985 ns/op
Benchmark_StdJSONEncode_MidSize-4           2502            431536 ns/op
Benchmark_EasyJsonEncode_MidSize-4         13312             91171 ns/op
Benchmark_StdGobEncode_MidSize-4            9175            142470 ns/op
Benchmark_StdJSONEncode_LargeSize-4           25          44295806 ns/op
Benchmark_EasyJsonEncode_LargeSize-4         122           9148548 ns/op
Benchmark_StdGobEncode_LargeSize-4            78          14089659 ns/op
Benchmark_StdJSONDecode-4                 581384              1977 ns/op
Benchmark_EasyJsonDecode-4               1703674               680 ns/op
Benchmark_StdGobDecode-4                   38575             29062 ns/op
Benchmark_StdJSONDecode_MidSize-4           1959            609825 ns/op
Benchmark_EasyJsonDecode_MidSize-4          4534            286501 ns/op
Benchmark_StdGobDecode_MidSize-4            7022            201647 ns/op
Benchmark_StdJSONDecode_LargeSize-4           16          67875257 ns/op
Benchmark_EasyJsonDecode_LargeSize-4          37          28440457 ns/op
Benchmark_StdGobDecode_LargeSize-4            73          13713334 ns/op
PASS
ok      github.com/fpapadopou/go-encoding-benchmark/encoding    29.169s

```

Encoding results visualization  
![Encoding duration graph](images/Encoding%20operation%20duration.png)

Decoding results visualization  
![Decoding duration graph](images/Decoding%20operation%20duration.png)
