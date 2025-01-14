// Copyright The HTNN Authors.
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

package key_auth

import (
	"net/url"

	"mosn.io/htnn/api/pkg/filtermanager/api"
	"mosn.io/htnn/types/plugins/key_auth"
)

func factory(c interface{}, callbacks api.FilterCallbackHandler) api.Filter {
	return &filter{
		callbacks: callbacks,
		config:    c.(*config),
	}
}

type filter struct {
	api.PassThroughFilter

	callbacks api.FilterCallbackHandler
	config    *config
}

func (f *filter) verify(value string) api.ResultAction {
	c, ok := f.callbacks.LookupConsumer(key_auth.Name, value)
	if !ok {
		return &api.LocalResponse{Code: 401, Msg: "invalid key"}
	}

	f.callbacks.SetConsumer(c)
	return api.Continue
}

func (f *filter) DecodeHeaders(headers api.RequestHeaderMap, endStream bool) api.ResultAction {
	config := f.config
	var query url.Values
	for _, key := range config.Keys {
		var vals []string
		if key.Source == key_auth.Source_QUERY {
			if query == nil {
				query = headers.Url().Query()
			}
			vals = query[key.Name]
		} else {
			vals = headers.Values(key.Name)
		}

		n := len(vals)
		if n == 1 {
			return f.verify(vals[0])
		}
		if n > 1 {
			return &api.LocalResponse{Code: 401, Msg: "duplicate key found"}
		}
	}
	return api.Continue
}
