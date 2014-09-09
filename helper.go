package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/streadway/simpleuuid"
	"net/smtp"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Uniqid() string {
	uuid, _ := simpleuuid.NewTime(time.Now())
	return uuid.String()
}

func Log(a ...interface{}) {
	fmt.Println(a)
}

func ParseSearchString(source string) string {
	s := strings.Trim(source, " ")
	if s != "" {
		return ""
	}
	re := regexp.MustCompile("\\s+")
	return strings.Trim(re.ReplaceAllString(strings.Trim(source, " ")+" ", "* "), " ")
}

func JoinInt(ints []int, glue string) string {
	return fmt.Sprintf(strings.Trim(strings.Replace(fmt.Sprintf("%d", ints), " ", glue, -1), "[]"))
}

func Base64Encode(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func ToHmacSha1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 寄送 mail
func SendMail(user string, password string, host string, port string, msg string, subject string, from string, to []string) error {
	auth := smtp.PlainAuth("", user, password, host)
	sub := "subject: " + subject + "\r\n\r\n"
	return smtp.SendMail(host+":"+port, auth, from, to, []byte(sub+msg))
}

//返回現在時間 mysql datetime
//格式:2014-12-21 21:20:32
func NowTime() string {
	t := time.Now()
	y, m, d := t.Date()
	H := t.Hour()
	i := t.Minute()
	se := t.Second()
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", y, m, d, H, i, se)
}

//返回現在時間 mysql datetime
//格式:2014-12-21 21:20:32
func NowTime8601() string {
	t := time.Now()
	y, m, d := t.Date()
	H := t.Hour()
	i := t.Minute()
	se := t.Second()
	return fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.000Z", y, m, d, H, i, se)
}

// 轉換成int
func ToInt(i interface{}) int {
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.String {
		v, err := strconv.Atoi(i.(string))
		if err == nil {
			return int(v)
		}
	}
	if kind == reflect.Int {
		return int(i.(int))
	}
	if kind == reflect.Int8 {
		return int(i.(int8))
	}
	if kind == reflect.Int16 {
		return int(i.(int16))
	}
	if kind == reflect.Int32 {
		return int(i.(int32))
	}
	if kind == reflect.Int64 {
		return int(i.(int64))
	}
	if kind == reflect.Uint {
		return int(i.(uint))
	}
	if kind == reflect.Uint8 {
		return int(i.(uint8))
	}
	if kind == reflect.Uint16 {
		return int(i.(uint16))
	}
	if kind == reflect.Uint32 {
		return int(i.(uint32))
	}
	if kind == reflect.Uint64 {
		return int(i.(uint64))
	}
	if kind == reflect.Float32 {
		return int(i.(float32))
	}
	if kind == reflect.Float64 {
		return int(i.(float64))
	}
	if kind == reflect.Bool {
		if i.(bool) == true {
			return int(1)
		} else {
			return int(0)
		}
	}
	return int(0)
}

func ToString(s interface{}) string {
	kind := reflect.TypeOf(s).Kind()
	if kind == reflect.String {
		return s.(string)
	}
	if kind == reflect.Int {
		return strconv.Itoa(s.(int))
	}
	if kind == reflect.Int8 {
		return strconv.Itoa(int(s.(int8)))
	}
	if kind == reflect.Int16 {
		return strconv.Itoa(int(s.(int16)))
	}
	if kind == reflect.Int32 {
		return strconv.Itoa(int(s.(int32)))
	}
	if kind == reflect.Int64 {
		return strconv.Itoa(int(s.(int64)))
	}
	if kind == reflect.Uint {
		return strconv.Itoa(int(s.(uint)))
	}
	if kind == reflect.Uint8 {
		return strconv.Itoa(int(s.(uint8)))
	}
	if kind == reflect.Uint16 {
		return strconv.Itoa(int(s.(uint16)))
	}
	if kind == reflect.Uint32 {
		return strconv.Itoa(int(s.(uint32)))
	}
	if kind == reflect.Uint64 {
		return strconv.Itoa(int(s.(uint64)))
	}
	return ""
}
