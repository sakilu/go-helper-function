package helper

import (
	"fmt"
	"time"
)

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
