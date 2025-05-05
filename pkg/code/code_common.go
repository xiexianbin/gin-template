// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// https://github.com/golang/go/blob/go1.22.3/src/net/http/Code.go
package code

const (
	// common err code biz id is 100000
	CodeOK Code = 100200 + iota
	CodeCreated
	CodeAccepted
	CodeNonAuthoritativeInfo
	CodeNoContent

	CodeBadRequest Code = 100400 + iota
	CodeUnauthorized
	CodePaymentRequired
	CodeForbidden
	CodeNotFound // 100404
	CodeMethodNotAllowed
	CodeNotAcceptable
	CodeProxyAuthRequired
	CodeRequestTimeout
	CodeConflict
	CodeUnKnown
)

var codeCommonMessages = map[Code]string{
	CodeOK: "OK",

	CodeCreated:              "Created",
	CodeAccepted:             "Accepted",
	CodeNonAuthoritativeInfo: "Non-Authoritative Information",
	CodeNoContent:            "No Content",

	CodeBadRequest:        "Bad Request",
	CodeUnauthorized:      "Unauthorized",
	CodePaymentRequired:   "Payment Required",
	CodeForbidden:         "Forbidden",
	CodeNotFound:          "Not Found",
	CodeMethodNotAllowed:  "Method Not Allowed",
	CodeNotAcceptable:     "Not Acceptable",
	CodeProxyAuthRequired: "Proxy Authentication Required",
	CodeRequestTimeout:    "Request Timeout",
	CodeConflict:          "Conflict",

	CodeUnKnown: "UnKnown",
}
