package utils

import (
	"net"
	"os"
	"strings"
	// "encoding/json"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"time"
)

// Replaces replace multi strings
// pairs - [old => new]
// can also use:
// strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	for old, newVal := range pairs {
		str = strings.Replace(str, old, newVal, -1)
	}

	return str
}

// GenMd5 生成32位md5字串
func GenMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encode
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// InternalIP
func InternalIP() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error())
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				// os.Stdout.WriteString(ipNet.IP.String() + "\n")
				ip = ipNet.IP.String()
				return
			}
		}
	}

	// os.Exit(0)
	return
}

// WriteJsonFile
func WriteJsonFile(filePath string, data interface{}) error {
	jsonBytes, err := JsonEncode(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, jsonBytes, 0664)
}

// ReadJsonFile
func ReadJsonFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return JsonDecode(content, v)
}

// JsonEncode encode data to json bytes. use it instead of json.Marshal
func JsonEncode(v interface{}) ([]byte, error) {
	var parser = jsoniter.ConfigCompatibleWithStandardLibrary

	return parser.Marshal(v)
}

// JsonEncode decode json bytes to data. use it instead of json.Unmarshal
func JsonDecode(json []byte, v interface{}) error {
	var parser = jsoniter.ConfigCompatibleWithStandardLibrary

	return parser.Unmarshal(json, v)
}

// Filling filling a model from submitted data
// data 提交过来的数据结构体
// model 定义表模型的数据结构体
// 相当于是在合并两个结构体(data 必须是 model 的子集)
func Filling(data interface{}, model interface{}) error {
	jsonBytes, _ := JsonEncode(data)

	return JsonDecode(jsonBytes, model)
}

// FormatDate
// str eg "2018-01-14T21:45:54+08:00"
func FormatDate(str string) string {
	// 先将时间转换为字符串
	tt, _ := time.Parse("2006-01-02T15:04:05Z07:00", str)

	// 格式化时间
	return tt.Format("2006-01-02 15:04:05")
}

// TransDateToTime
func TransDateToTime(date string) (t time.Time, ok bool) {
	var layout string

	switch len(date) {
	case 10: // 2006-01-02
		layout = "2006-01-02"
	case 19: // 2006-01-02 12:24:36
		layout = "2006-01-02 15:04:05"
	default:
		return
	}

	t, err := time.ParseInLocation(layout, date, time.Local)
	ok = err == nil

	return
}

// CalcElapsedTime 计算运行时间消耗 单位 ms(毫秒)
func CalcElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}

// GetRootPath Get abs path of the project
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}

// Substr
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length

	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// TransStruct2Map translate structure to map
func TransStruct2Map(obj interface{}) (mp map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		fmt.Sprintf("%d %+v\n", i, t.Field(i))
		mp[t.Field(i).Name] = v.Field(i).Interface()
	}

	return
}
