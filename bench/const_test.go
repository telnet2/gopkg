package bench

import "testing"

var (
	headers = map[string]string{HeaderKey: "value"}

	headerKeyValue string
)

func GetHeader(key string) string {
	return headers[key]
}

func BenchmarkRefConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		headerKeyValue = GetHeader(HeaderKey)
	}
}

func BenchmarkRefStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		headerKeyValue = GetHeader(Headers.Key)
	}
}

func TestRef(t *testing.T) {
	headerKeyValue = GetHeader(HeaderKey)
	headerKeyValue = GetHeader(Headers.Key)
}
