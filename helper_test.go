package helper

import (
	"testing"
)

func TestLog(t *testing.T) {
	t.Skip()
}

func BenchmarkLog(b *testing.B) {
	b.Skip()
}

func TestUniqid(t *testing.T) {

}

func BenchmarkUniqid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uniqid()
	}
}

func TestParseSearchString(t *testing.T) {

}

func TestJoinInt(t *testing.T) {
	ints := []int{1, 2, 3, 4}

	glue := ","
	if JoinInt(ints, glue) != "1,2,3,4" {
		t.Error(JoinInt(ints, glue), " != ", "1,2,3,4")
	}
	glue = " ,"
	if JoinInt(ints, glue) != "1 ,2 ,3 ,4" {
		t.Error(JoinInt(ints, glue), " != ", "1 ,2 ,3 ,4")
	}
	glue = ""
	if JoinInt(ints, glue) != "1234" {
		t.Error(JoinInt(ints, glue), " != ", "1234")
	}
	glue = "/"
	if JoinInt(ints, glue) != "1/2/3/4" {
		t.Error(JoinInt(ints, glue), " != ", "1/2/3/4")
	}
	glue = ","
	ints = []int{}
	if JoinInt(ints, glue) != "" {
		t.Error()
	}
	ints = []int{1}
	if JoinInt(ints, glue) != "1" {
		t.Error(JoinInt(ints, glue), " != ", 1)
	}
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
