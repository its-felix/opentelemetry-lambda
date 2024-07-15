//go:build lambdacomponents.custom && (lambdacomponents.all || lambdacomponents.receiver.all || lambdacomponents.receiver.telemetryapi)

// Copyright The OpenTelemetry Authors
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

package receiver

import (
	"github.com/open-telemetry/opentelemetry-lambda/collector/receiver/telemetryapireceiver"
	"go.opentelemetry.io/collector/receiver"
)

func init() {
	Factories = append(Factories, func(extensionId string) receiver.Factory {
		return telemetryapireceiver.NewFactory(extensionId)
	})
}
