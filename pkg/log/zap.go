// Copyright 2025 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package log implements the app logging package by uber zap.
package log

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once

	Logger     *zap.Logger
	CoreLogger *zap.Logger // with core flag
)

func Init() {
	once.Do(func() {
		var err error
		// ref zap.NewProduction()
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		config := zap.Config{
			Level:       zap.NewAtomicLevelAt(GetLogLevel()),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    encoderConfig,
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}
		Logger, err = config.Build()
		if err != nil {
			panic(fmt.Errorf("init zap log err: %w", err))
		}
		defer Logger.Sync()

		// init Logger for compontent
		CoreLogger = Logger.With(zap.String("component", "core"))

		CoreLogger.Info("zap logger init done")
	})
}

// GetLogLevel return zap logger level
func GetLogLevel() (zapLogLevel zapcore.Level) {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch logLevel {
	case "debug":
		zapLogLevel = zap.DebugLevel
	case "info":
		zapLogLevel = zap.InfoLevel
	case "warn":
		zapLogLevel = zap.WarnLevel
	case "error":
		zapLogLevel = zap.ErrorLevel
	case "dpanic":
		zapLogLevel = zap.DPanicLevel
	case "panic":
		zapLogLevel = zap.PanicLevel
	case "fatal":
		zapLogLevel = zap.FatalLevel
	default:
		zapLogLevel = zap.DebugLevel
	}
	return
}
