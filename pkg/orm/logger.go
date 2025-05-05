// Copyright 2024 xiexianbin<me@xiexianbin.cn>
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

// ref https://gorm.io/docs/logger.html#Logger
package orm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gorml "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GormLogger implemente gorm.io/gorm/logger/#Interface logging interface
type GormLogger struct {
	ZapLogger *zap.Logger
	Config    gorml.Config
}

var _ gorml.Interface = (*GormLogger)(nil)

func (l *GormLogger) getLogLevel() gorml.LogLevel {
	if l.ZapLogger.Level() == zapcore.ErrorLevel {
		return gorml.Error
	}
	if l.ZapLogger.Level() == zapcore.WarnLevel {
		return gorml.Warn
	}
	if l.ZapLogger.Level() == zapcore.InfoLevel {
		return gorml.Info
	}
	return gorml.Silent
}

// LogMode log mode
func (l *GormLogger) LogMode(level gorml.LogLevel) gorml.Interface {
	// Note(xiexianbin): set log level, default config.LOG_LEVEL
	return l
}

// Info print info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.getLogLevel() >= gorml.Info {
		reqId := ctx.Value("req-id").(string)
		l.ZapLogger.Sugar().With(zap.String("req_id", reqId)).Infof(msg, data...)
	}
}

// Warn print warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.getLogLevel() >= gorml.Warn {
		reqId := ctx.Value("req-id").(string)
		l.ZapLogger.Sugar().With(zap.String("req_id", reqId)).Warnf(msg, data...)
	}
}

// Error print error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.getLogLevel() >= gorml.Error {
		reqId := ctx.Value("req-id").(string)
		l.ZapLogger.Sugar().With(zap.String("req_id", reqId)).Errorf(msg, data...)
	}
}

// Trace print sql message
//
//nolint:cyclop
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.getLogLevel() <= gorml.Silent {
		return
	}

	reqId := ctx.Value("req-id").(string)

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.getLogLevel() >= gorml.Error &&
		(!errors.Is(err, gorm.ErrRecordNotFound) || !l.Config.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.ZapLogger.Sugar().With().Errorw(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.Error(err),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
			)
		} else {
			l.ZapLogger.Sugar().With().Errorw(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.Error(err),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("rows", rows),
				zap.String("sql", sql),
			)
		}
	case elapsed > l.Config.SlowThreshold && l.Config.SlowThreshold != 0 && l.getLogLevel() >= gorml.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.Config.SlowThreshold)
		if rows == -1 {
			l.ZapLogger.Sugar().With().Warnw(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.String("slowLog", slowLog),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
			)
		} else {
			l.ZapLogger.Sugar().With().Warnw(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.String("slowLog", slowLog),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("rows", rows),
				zap.String("sql", sql),
			)
		}
	case l.getLogLevel() == gorml.Info:
		sql, rows := fc()
		if rows == -1 {
			l.ZapLogger.Sugar().With().Infow(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("rows", -1),
				zap.String("sql", sql),
			)
		} else {
			l.ZapLogger.Sugar().With().Infow(
				"",
				zap.String("req_id", reqId),
				zap.String("linenum", utils.FileWithLineNum()),
				zap.Float64("elapse", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("rows", rows),
				zap.String("sql", sql),
			)
		}
	}
}
