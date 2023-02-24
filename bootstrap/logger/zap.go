package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"skeleton/config"
	"time"
)

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func Init() {
	date := time.Now().Format("2006-01-02")
	debug := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zap.DebugLevel
	})
	other := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l > zap.DebugLevel
	})
	all := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= config.YAML.Logger.Level
	})

	//Cores
	cores := [...]zapcore.Core{
		//主要是设置writer和日志分割
		zapcore.NewCore(getEncoder(), getWrite(fmt.Sprintf("storage/logs/%s/debug.log", date)), debug),
		zapcore.NewCore(getEncoder(), getWrite(fmt.Sprintf("storage/logs/%s/other.log", date)), other),
		zapcore.NewCore(getEncoder(), os.Stdout, all),
	}

	Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	SugarLogger = Logger.Sugar()
}

// 定义配置
func getZapConfig() zapcore.EncoderConfig {
	zapConfig := zap.NewProductionEncoderConfig()

	//日志时间
	zapConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("[2006-01-02 15:04:05]"))
	}

	//日志等级
	zapConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString("[" + l.CapitalString() + "]")
	}

	//触发点
	zapConfig.EncodeCaller = func(ec zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString("[" + ec.TrimmedPath() + "]")
	}

	return zapConfig
}

// 定义Encoder
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(getZapConfig())
}

// 获取文件
func getWrite(filename string) zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   true,
	}
	return zapcore.AddSync(l)
}
