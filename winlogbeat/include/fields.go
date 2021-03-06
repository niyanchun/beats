// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("winlogbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
