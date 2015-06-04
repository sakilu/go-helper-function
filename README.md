# go-helper-function
## functions
    func ParseSearchString(t *testing.T)
    拆空白字串加星號,供mysql全文檢索用
    
    func JoinInt(ints []int, glue string) string
    將int陣列join為字串 JoinInt([]int{1, 2, 3, 4}, ',') => 1,2,3,4
    
    Base64Encode(message string) string 
    Base64Encode
    
    func ToHmacSha1(message string, secret string) string
    string to HmacSha1
    
    func SendMail(user string, password string, host string, port string, msg string, subject string, from string, to []string) error
    寄送email by smtp
    
    func NowTime() string
    返回現在時間 格式:2014-12-21 21:20:32
    
    func NowTime8601() string
    返回現在時間 Time8601格式:2013-02-25T18:30:00.000Z
    
    ToInt(i interface{}) int
    轉型為int
    
    ToString(s interface{}) string
    轉型為string
    
    
