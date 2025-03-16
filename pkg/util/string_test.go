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
	"reflect"
	"testing"
)

func TestStringAny(t *testing.T) {
	type person struct {
		Name    string
		Age     int
		Hobbies []string
	}
	type args struct {
		x any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "struct-test",
			args: args{
				x: &person{
					Name:    "Alice",
					Age:     30,
					Hobbies: []string{"Reading", "Hiking"},
				},
			},
			want: "{\"Name\":\"Alice\",\"Age\":30,\"Hobbies\":[\"Reading\",\"Hiking\"]}",
		},
		{
			name: "array-test",
			args: args{
				x: []int{1, 2, 3},
			},
			want: "[1, 2, 3]",
		},
		{
			name: "map-test",
			args: args{
				x: map[string]int{"a": 1},
			},
			want: "{\"a\": 1}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringAny(tt.args.x); got != tt.want {
				t.Errorf("StringAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "StringToBytes",
			args: args{s: "abc"},
			want: []byte{'a', 'b', 'c'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "StringToBytes",
			args: args{b: []byte{'a', 'b', 'c'}},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
