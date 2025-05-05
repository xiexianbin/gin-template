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

package redis

import (
	"context"

	"go.uber.org/zap"
)

// RedisLogger implemente github.com/redis/go-redis/v9/internal/#Logging logging interface
type RedisLogger struct {
	ZapLogger *zap.Logger
}

func (l *RedisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	reqId := ctx.Value("req-id").(string)
	if reqId != "" {
		l.ZapLogger.Sugar().With(zap.String("req_id", reqId)).Debugf(format, v...)
	} else {
		l.ZapLogger.Sugar().Debugf(format, v...)
	}
}
