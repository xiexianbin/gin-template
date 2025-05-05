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

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// RemoveNonASCII Removes all non-ASCII characters from a string
func RemoveNonASCII(s string) string {
	var result strings.Builder
	for _, r := range s {
		if r <= 127 {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// RemoveNonUTF8 Removes all no-UTF-8 characters from a string
func RemoveNonUTF8(s string) string {
	return strings.Map(func(r rune) rune {
		if r == utf8.RuneError {
			return -1
		}
		return r
	}, s)
}

// RemoveControlChars remove (unicode)control chars
func RemoveControlChars(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsControl(r) {
			result = append(result, r)
		}
	}
	return string(result)

	// remove ascii control chars(0-31 and 127)
	// re := regexp.MustCompile("[\x00-\x1F\x7F]")
	// return re.ReplaceAllString(s, "")

	// result := make([]byte, 0, len(s))
	// for i := 0; i < len(s); i++ {
	// 		b := s[i]
	// 		if b > 31 && b != 127 {
	// 				result = append(result, b)
	// 		}
	// }
	// return string(result)

}

// StringAny convert any to string
func StringAny(x any) string {
	// 处理所有类型的 nil 值
	if x == nil {
		return "nil"
	}

	v := reflect.ValueOf(x)

	// 判断是否为结构体或结构体指针
	isStruct := false
	switch v.Kind() {
	case reflect.Struct:
		isStruct = true
	case reflect.Ptr:
		if v.Elem().Kind() == reflect.Struct {
			isStruct = true
		}
	}

	// 优先尝试 JSON 序列化结构体
	if isStruct {
		if data, err := json.Marshal(x); err == nil {
			return string(data)
		}
	}

	// 其他类型，递归处理
	return stringAny(reflect.ValueOf(x), true)
}

func stringAny(v reflect.Value, topLevel bool) string {
	if !v.IsValid() {
		return "<invalid>"
	}

	// 处理nil值
	if (v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Func) && v.IsNil() {
		return "nil"
	}

	// 获取实际值
	v = reflect.Indirect(v)

	// 处理实现了 Stringer 接口的类型
	if stringer, ok := v.Interface().(fmt.Stringer); ok {
		if v.Kind() == reflect.Ptr && v.IsNil() {
			return "nil"
		}
		return stringer.String()
	}

	// 处理 time.Time 特殊类型
	if t, ok := v.Interface().(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	// 根据类型处理
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.6f", v.Float())

	case reflect.String:
		return strconv.Quote(v.String())

	case reflect.Slice, reflect.Array:
		return formatSlice(v)

	case reflect.Map:
		return formatMap(v)

	case reflect.Struct:
		return formatStruct(v)

	case reflect.Ptr:
		return "&" + stringAny(v.Elem(), false)

	case reflect.Interface:
		return stringAny(v.Elem(), false)

	case reflect.Func:
		return "func(" + v.Type().String() + ")"

	case reflect.Chan:
		return "chan " + v.Type().Elem().String()

	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

func formatSlice(v reflect.Value) string {
	buf := bytes.NewBufferString("[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(stringAny(v.Index(i), false))
	}
	buf.WriteString("]")
	return buf.String()
}

func formatMap(v reflect.Value) string {
	buf := bytes.NewBufferString("{")
	keys := v.MapKeys()
	for i, key := range keys {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(stringAny(key, false))
		buf.WriteString(": ")
		buf.WriteString(stringAny(v.MapIndex(key), false))
	}
	buf.WriteString("}")
	return buf.String()
}

func formatStruct(v reflect.Value) string {
	t := v.Type()
	buf := bytes.NewBufferString(t.Name() + "{")
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		field := t.Field(i)
		// 不导出字段跳过
		if field.PkgPath != "" {
			continue
		}
		buf.WriteString(field.Name)
		buf.WriteString(": ")
		buf.WriteString(stringAny(v.Field(i), false))
	}
	buf.WriteString("}")
	return buf.String()
}

// StringToBytes convert string to []byte
// https://github.com/golang/go/issues/53003#issuecomment-1140276077
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString convert []byte to string
// https://github.com/golang/go/issues/53003#issuecomment-1140276077
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}
