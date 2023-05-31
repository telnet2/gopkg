package bench

import "testing"

const HeaderKey = "X-Header-Key"

var (
	headers = map[string]string{HeaderKey: "value"}

	headerKeyValue string
)

var Headers = struct {
	Key string
}{
	Key: HeaderKey,
}

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
