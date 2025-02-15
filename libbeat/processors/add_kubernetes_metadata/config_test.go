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

package add_kubernetes_metadata

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/common"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		cfg   map[string]interface{}
		error bool
	}{
		{
			cfg: map[string]interface{}{
				"scope": "foo",
			},
			error: true,
		},
		{
			cfg: map[string]interface{}{
				"scope": "cluster",
			},
			error: false,
		},
		{
			cfg:   map[string]interface{}{},
			error: false,
		},
	}

	for _, test := range tests {
		cfg := common.MustNewConfigFrom(test.cfg)
		c := defaultKubernetesAnnotatorConfig()

		err := cfg.Unpack(&c)
		if test.error {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
	}
}
