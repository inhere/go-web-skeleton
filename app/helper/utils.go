package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

// LocTime get local time
func LocTime() time.Time {
	// loc, _ := time.LoadLocation(Timezone)
	// return time.Now().In(loc)
	return time.Now().Local()
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

// GetRootPath Get abs path of the project
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}
