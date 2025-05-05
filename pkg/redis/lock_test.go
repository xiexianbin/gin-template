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
	"sync"
	"testing"
	"time"
)

func TestAcquireDistributedLock(t *testing.T) {
	lockKey := "lockKey"
	lockExpire := 3 // time.Second
	wg := sync.WaitGroup{}

	// Create multiple coroutines to simulate concurrent scenarios
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// try acquire lock
			ok, err := AcquireDistributedLock(context.TODO(), lockKey, fmt.Sprintf("v-%d", i), lockExpire)
			if err != nil {
				t.Error(err)
			}
			if ok {
				defer ReleaseDistributedLock(context.TODO(), lockKey)

				// 进入临界区
				t.Logf("id %d successfully acquired the lock", i)
				time.Sleep(time.Second)
				t.Logf("id %d successfully release the lock", i)
			} else {
				t.Logf("id %d successfully acquired the lock failed", i)
			}
		}(i)
	}
	wg.Wait()
}
