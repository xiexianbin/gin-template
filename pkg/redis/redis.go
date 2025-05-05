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

// Package redis implements a Redis client by github.com/redis/go-redis.
// - github.com/go-redis/redis/v8 for redis 6, https://github.com/redis/go-redis/blob/v8.11.5/Makefile#L19
// - github.com/redis/go-redis/v9 for redis 7, https://github.com/redis/go-redis/blob/v9.5.2/Makefile#L34
// ref https://github.com/redis/go-redis?tab=readme-ov-file#quickstart
package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/xiexianbin/gin-template/pkg/util"
)

var (
	rdb  *redis.Client
	rdbc *redis.ClusterClient
	once sync.Once
)

func InitRedis(redisDSN string) {
	once.Do(func() {
		// Redis Single
		if redisDSN != "" {
			opt, err := redis.ParseURL(redisDSN)
			util.Mustf(err, "parse Redis by REDIS_DSN failed")

			opt.DialTimeout = 3 * time.Second // no time unit = seconds
			opt.ReadTimeout = 6 * time.Second

			// connection pool
			opt.PoolFIFO = true
			opt.PoolSize = 10

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewClient
			rdb = redis.NewClient(opt)
			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis err")
			return
		}
	})
}

func InitRedisSentinel(redisSentinelDSN string) {
	once.Do(func() {
		// Redis Sentinel Cluster
		if redisSentinelDSN != "" {
			opt, err := redis.ParseClusterURL(redisSentinelDSN)
			util.Mustf(err, "init Redis by REDIS_SENTINEL_DSN failed")

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewFailoverClient
			rdb = redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    "master",
				ClientName:    opt.ClientName,
				SentinelAddrs: opt.Addrs,
				Password:      opt.Password,

				DialTimeout: 3 * time.Second, // no time unit = seconds
				ReadTimeout: 6 * time.Second,

				PoolFIFO: true,
				PoolSize: 10,
			})

			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis sentinel err")
			return
		}
	})
}

func InitRedisCluster(redisClusterNSN string) {
	// Redis Cluster
	once.Do(func() {
		if redisClusterNSN != "" {
			opt, err := redis.ParseClusterURL(redisClusterNSN)
			util.Mustf(err, "parse Redis by REDIS_CLUTER_DSN failed")

			opt.DialTimeout = 3 * time.Second // no time unit = seconds
			opt.ReadTimeout = 6 * time.Second

			// connection pool
			opt.PoolFIFO = true
			opt.PoolSize = 10

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewClusterClient
			rdbc = redis.NewClusterClient(opt)
			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis cluster err")
			return
		}
	})
}

// Pool return redis.ClusterClient or redis.Client
// Usage demo:
//
//	Pool().Get(...)
func Pool() IRedis {
	if rdbc != nil {
		return rdbc
	}
	if rdb != nil {
		return rdb
	}
	panic("no redis client and cluster client init")
}

// ClientPool return redis client
func ClientPool() *redis.Client {
	if rdb != nil {
		return rdb
	}
	panic("no redis client init")
}

// ClusterClientPool return redis cluster client
func ClusterClientPool() *redis.Client {
	if rdb != nil {
		return rdb
	}
	panic("no redis cluster client init")
}
