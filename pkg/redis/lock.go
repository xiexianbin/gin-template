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
	"fmt"
)

// AcquireDistributedLock acquire distributed lock by key
func AcquireDistributedLock(ctx context.Context, key, value string, expire int) (bool, error) {
	lua_script := `
			local lockValue = redis.call('get', KEYS[1])
			if lockValue == nil then
					redis.call('set', KEYS[1], ARGV[1], 'EX', ARGV[2])
					return redis.call('get', KEYS[1])
			elseif lockValue == ARGV[1] then
					redis.call('expire', KEYS[1], ARGV[2])
					return redis.call('get', KEYS[1])
			else
					return nil
			end
	`

	// Use the Lua script to try to get the lock
	eval := rdb.Eval(ctx, lua_script,
		[]string{key},
		[]string{value, fmt.Sprintf("%d", expire)})
	result, err := eval.Result()
	if err != nil {
		return false, err
	}

	if eval.Val() == nil {
		return false, fmt.Errorf("failed to acquire lock")
	}

	return result.(string) == value, nil
}

// ReleaseDistributedLock release distributed lock by key
func ReleaseDistributedLock(ctx context.Context, key string) error {
	return rdb.Del(ctx, key).Err()
}
