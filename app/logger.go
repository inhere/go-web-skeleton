package app

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"os"
	"strings"
)

// logger instance
var Logger *zap.Logger

// initLog init log setting
func initLogger() {
	// newGenericLogger()
	newRotatedLogger()

	Logger.Info("logger construction succeeded")
}

func newGenericLogger() {
	var err error
	var cfg zap.Config

	conf, _ := Cfg.StringMap("log")
	logFile := conf["logFile"]
	errFile := conf["errFile"]

	// replace
	logFile = strings.NewReplacer(
		"{date}", LocTime().Format("20060102"),
		"{hostname}", Hostname,
	).Replace(logFile)

	errFile = strings.NewReplacer(
		"{date}", LocTime().Format("20060102"),
		"{hostname}", Hostname,
	).Replace(errFile)

	// create config
	if Debug {
		// cfg = zap.NewDevelopmentConfig()
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Development = true
		cfg.OutputPaths = []string{"stdout", logFile}
		cfg.ErrorOutputPaths = []string{"stderr", errFile}
	} else {
		cfg = zap.NewProductionConfig()
		cfg.OutputPaths = []string{logFile}
		cfg.ErrorOutputPaths = []string{errFile}
	}

	// init some defined fields to log
	cfg.InitialFields = map[string]interface{}{
		"hostname": Hostname,
		// "context": map[string]interface{}{},
	}

	// create logger
	Logger, err = cfg.Build()

	if err != nil {
		panic(err)
	}
}

// see https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
func newRotatedLogger() {
	var cfg zap.Config

	conf, _ := Cfg.StringMap("log")
	logFile := conf["logFile"]
	// errFile := conf["errFile"]

	// replace
	logFile = strings.NewReplacer(
		"{date}", LocTime().Format("20060102"),
		"{hostname}", Hostname,
	).Replace(logFile)

	fmt.Printf("logger - file=%s\n", logFile)

	// create config
	if Debug {
		// cfg = zap.NewDevelopmentConfig()
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Development = true
		cfg.OutputPaths = []string{"stdout", logFile}
		// cfg.ErrorOutputPaths = []string{"stderr", errFile}
	} else {
		cfg = zap.NewProductionConfig()
		cfg.OutputPaths = []string{logFile}
		// cfg.ErrorOutputPaths = []string{errFile}
	}

	// lumberjack.Logger is already safe for concurrent use, so we don't need to lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: logFile,
		MaxSize:  10, // megabytes
		// MaxBackups: 3,
		// MaxAge: 28, // days
	})

	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		w,
		cfg.Level,
	)

	// init some defined fields to log
	Logger = zap.New(core).With(zap.String("hostname", Hostname))
}

// see https://godoc.org/go.uber.org/zap#hdr-Frequently_Asked_Questions
func advLogger() {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	topicDebugging := zapcore.AddSync(ioutil.Discard)
	topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		// high
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		// low
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}

// LogToContext
func LogToContext() zap.Field {
	return zap.Namespace("context")
}
