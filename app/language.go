package app

import (
	"fmt"
	"github.com/gookit/i18n"
	"strings"
)

func initLanguage() {
	// conf := map[string]string{
	// 	"langDir": "res/lang",
	// 	"allowed": "en:English|zh-CN:简体中文",
	// 	"default": "en",
	// }
	conf, _ := Cfg.StringMap("lang")
	fmt.Printf("language - %v\n", conf)

	// en:English|zh-CN:简体中文
	langList := strings.Split(conf["allowed"], "|")
	languages := make(map[string]string, len(langList))

	for _, str := range langList {
		item := strings.Split(str, ":")
		languages[item[0]] = item[1]
	}

	// init and load data
	i18n.Init(conf["langDir"], conf["default"], languages)
}

// Tr
func Tr(lang string, key string, args ...interface{}) string {
	return i18n.Tr(lang, key, args...)
}

// Dtr translate from default lang
func Dtr(key string, args ...interface{}) string {
	return i18n.DefTr(key, args...)
}
