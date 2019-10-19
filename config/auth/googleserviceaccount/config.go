// Copyright (c) 2019, Noel Cower.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package googleserviceaccount contains a spin CLI config structure to use Google service accounts
// for oauth2.
package googleserviceaccount

import "golang.org/x/oauth2"

type GoogleServiceAccountConfig struct {
	File string `yaml:"file"`

	CachedToken *oauth2.Token `yaml:"cachedToken,omitempty"`
}

func (g *GoogleServiceAccountConfig) IsEnabled() bool {
	return g != nil
}
