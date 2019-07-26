package app

import (
	"os"

	"github.com/inhere/go-web-skeleton/model"
)

// allowed app env name
const (
	PROD = "prod"
	PRE  = "pre"
	TEST = "test"
	DEV  = "dev"
)

// for application
const (
	Timezone = "PRC"
	BaseDate = "2006-01-02 15:04:05"

	Timeout     = 10
	PageSize    = 20
	PageSizeStr = "20"
	MaxPageSize = 100
)

// application info
var (
	Env  = "dev"
	Name = "go-web-skeleton"

	debug bool

	Hostname string
	RootPath string
	GitInfo  model.GitInfoData
	HttpPort = 9440
)

// the app work dir path
var WorkDir, _ = os.Getwd()

// IsEnv current env name check
func IsEnv(env string) bool {
	return env == Env
}

// IsDebug is debug mode
func IsDebug() bool {
	return debug
}
