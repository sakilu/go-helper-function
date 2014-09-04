package helper

import (
	"testing"
)

func TestUniqid(t *testing.T) {

}

func BenchmarkUniqid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uniqid()
	}
}

func TestLog(t *testing.T) {

}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Log(i)
	}
}

func TestParseSearchString(t *testing.T) {

}

func TestJoinInt(t *testing.T) {

}

func TestBase64Encode(t *testing.T) {

}

func TestToHmacSha1(t *testing.T) {

}

func TestSendMail(t *testing.T) {

}

func TestNowTime(t *testing.T) {

}

func TestNowTime8601(t *testing.T) {

}

func TestToInt(t *testing.T) {

}

func TestToString(t *testing.T) {

}
