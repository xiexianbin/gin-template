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

package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/xiexianbin/gin-template/pkg/log"
)

// Load loads .env file from specified paths
func Load(paths ...string) error {
	var path string
	if len(paths) > 0 {
		path = paths[0]
	} else {
		// Default search for .env file in current and parent directories
		path = findEnvFile()
		// Not finding a .env file is not considered an error
		if path == "" {
			return nil
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open .env file failed: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip empty lines and comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // invalid file format
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// If the environment variable already exists, do not overwrite
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read .env file failed: %v", err)
	}

	return nil
}

// findEnvFile searches for .env file
func findEnvFile() string {
	// First check current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	envPath := filepath.Join(currentDir, ".env")
	if _, err := os.Stat(envPath); err == nil {
		return envPath
	}

	// Check parent directory
	parentDir := filepath.Dir(currentDir)
	envPath = filepath.Join(parentDir, ".env")
	if _, err := os.Stat(envPath); err == nil {
		return envPath
	}

	// Check executable directory
	if ex, err := os.Executable(); err == nil {
		envPath = filepath.Join(filepath.Dir(ex), ".env")
		if _, err := os.Stat(envPath); err == nil {
			return filepath.Dir(ex)
		}
	}

	return ""
}

// GetStr retrieves environment variable, returns default value if not exists
func GetStr(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// MustGet retrieves environment variable, panics if not exists
func MustGet(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("env %s not found", key))
}

// GetInt retrieves integer environment variable
func GetInt(key string, defaultValue int) int {
	valueStr := GetStr(key, "")
	if valueStr == "" {
		return defaultValue
	}

	var value int
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		return defaultValue
	}
	// d, err := strconv.Atoi(valueStr)
	return value
}

// GetFloat retrieves float64 environment variable
func GetFloat64(key string, o float64) float64 {
	v, found := os.LookupEnv(key)
	if found && v != "" {
		d, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.CoreLogger.Sugar().Panicw("failed to convert to float", zap.Any(key, v), zap.Error(err))
		} else {
			return d
		}
	}
	return o
}

// GetBool retrieves boolean environment variable
func GetBool(key string, defaultValue bool) bool {
	valueStr := GetStr(key, "")
	if valueStr == "" {
		return defaultValue
	}

	b, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return b
}

// GetDuration retrieves duration environment variable
// Supported formats: "30s", "5m", "1h", "2h45m"
func GetDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := GetStr(key, "")
	if valueStr == "" {
		return defaultValue
	}

	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	return duration
}
