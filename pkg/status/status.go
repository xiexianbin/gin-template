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

// Package status implements errors returned by the app.
package status

import (
	"fmt"

	"github.com/xiexianbin/gin-template/pkg/code"
)

type Status struct {
	// The status code, which should be an enum value of
	// [pkg/code.Code].
	Code code.Code `json:"code,omitempty"`
	// A developer-facing error message, which should be in English.
	Message string `json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details []any `json:"details,omitempty"`
}

// Err returns an immutable error representing s; returns nil if s.Code() is OK.
func (s *Status) Err() error {
	if s.Code == code.CodeOK || s.Code == code.CodeCreated || s.Code == code.CodeAccepted {
		return nil
	}
	return &StatusError{s: s}
}

// String return Status 's Message
func (x *Status) String() string {
	return x.Message
}

// New returns a Status representing c and msg.
func New(c code.Code, msg string) *Status {
	return &Status{
		Code:    c,
		Message: msg,
	}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c code.Code, format string, a ...any) *Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c code.Code, msg string) error {
	return New(c, msg).Err()
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c code.Code, format string, a ...any) error {
	return Error(c, fmt.Sprintf(format, a...))
}

// Code returns the Code of the error if it is a Status error or if it wraps a
// Status error. If that is not the case, it returns code.CodeOK if err is nil, or
// code.Unknown otherwise.
// func Code(err error) code.Code {
// 	// Don't use FromError to avoid allocation of OK status.
// 	if err == nil {
// 		return code.CodeOK
// 	}

// 	return
// }

// StatusError wraps a pointer of a Status. It implements error.
type StatusError struct {
	s *Status
}

func (e *StatusError) Error() string {
	return e.s.String()
}

// Status returns the Status represented by se.
func (e *StatusError) Status() *Status {
	return e.s
}

// Is implements future error.Is functionality.
// A Error is equivalent if the code and message are identical.
func (e *StatusError) Is(target error) bool {
	tse, ok := target.(*StatusError)
	if !ok {
		return false
	}
	return e.s.Code == tse.s.Code
}
