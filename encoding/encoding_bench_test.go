package encoding

import (
	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	ENCODING
*/
func createObjectWithListSize(size int) object {
	var list []embedded
	for i := 0; i < size; i++ {
		list = append(list, embedded{
			StringProp: "foo",
			IntProp:    100,
		})
	}
	o := object{
		StringProp1: "foo",
		StringProp2: "bar",
		List:        list,
	}

	return o
}

func Benchmark_StdJSONEncode(b *testing.B) {
	o := createObjectWithListSize(1)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdJSONEncode()
	}
}

func Benchmark_EasyJsonEncode(b *testing.B) {
	o := createObjectWithListSize(1)
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(o)
	}
}

func Benchmark_StdGobEncode(b *testing.B) {
	o := createObjectWithListSize(1)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdGobEncode()
	}
}

func Benchmark_StdJSONEncode_MidSize(b *testing.B) {
	o := createObjectWithListSize(1000)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdJSONEncode()
	}
}

func Benchmark_EasyJsonEncode_MidSize(b *testing.B) {
	o := createObjectWithListSize(1000)
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(o)
	}
}

func Benchmark_StdGobEncode_MidSize(b *testing.B) {
	o := createObjectWithListSize(1000)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdGobEncode()
	}
}

func Benchmark_StdJSONEncode_LargeSize(b *testing.B) {
	o := createObjectWithListSize(100000)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdJSONEncode()
	}
}

func Benchmark_EasyJsonEncode_LargeSize(b *testing.B) {
	o := createObjectWithListSize(100000)
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(o)
	}
}

func Benchmark_StdGobEncode_LargeSize(b *testing.B) {
	o := createObjectWithListSize(100000)
	for i := 0; i < b.N; i++ {
		_, _ = o.stdGobEncode()
	}
}

/*
	DECODING
*/
func Benchmark_StdJSONDecode(b *testing.B) {
	i := createObjectWithListSize(1)
	encoded, err := i.stdJSONEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdJSONDecode(encoded)
	}
}

func Benchmark_EasyJsonDecode(b *testing.B) {
	i := createObjectWithListSize(1)
	encoded, err := easyjson.Marshal(i)
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(encoded, &o)
	}
}

func Benchmark_StdGobDecode(b *testing.B) {
	i := createObjectWithListSize(1)
	encoded, err := i.stdGobEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdGobDecode(encoded)
	}
}

func Benchmark_StdJSONDecode_MidSize(b *testing.B) {
	i := createObjectWithListSize(1000)
	encoded, err := i.stdJSONEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdJSONDecode(encoded)
	}
}

func Benchmark_EasyJsonDecode_MidSize(b *testing.B) {
	i := createObjectWithListSize(1000)
	encoded, err := easyjson.Marshal(i)
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(encoded, &o)
	}
}

func Benchmark_StdGobDecode_MidSize(b *testing.B) {
	i := createObjectWithListSize(1000)
	encoded, err := i.stdGobEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdGobDecode(encoded)
	}
}

func Benchmark_StdJSONDecode_LargeSize(b *testing.B) {
	i := createObjectWithListSize(100000)
	encoded, err := i.stdJSONEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdJSONDecode(encoded)
	}
}

func Benchmark_EasyJsonDecode_LargeSize(b *testing.B) {
	i := createObjectWithListSize(100000)
	encoded, err := easyjson.Marshal(i)
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(encoded, &o)
	}
}

func Benchmark_StdGobDecode_LargeSize(b *testing.B) {
	i := createObjectWithListSize(100000)
	encoded, err := i.stdGobEncode()
	require.NoError(b, err)

	var o object
	for i := 0; i < b.N; i++ {
		_ = o.stdGobDecode(encoded)
	}
}
