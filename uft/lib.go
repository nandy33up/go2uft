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

package uft

import (
	"runtime/cgo"

	"github.com/nandy33up/go2uft/thost"
)

const (
	defaultFlowPath         = ""
	defaultIsProductionMode = true
)

type MdOption func(api *CMdApi)

func MdFlowPath(path string) MdOption {
	return func(api *CMdApi) {
		api.flowPath = path
	}
}

func MdProductionMode(production bool) MdOption {
	return func(api *CMdApi) {
		api.production = production
	}
}

type CMdApi struct {
	apiPtr     uintptr
	spi        thost.CHSMdSpi
	handle     cgo.Handle
	flowPath   string
	production bool
}

type TradeOption func(api *CTradeApi)

func TradeFlowPath(path string) TradeOption {
	return func(api *CTradeApi) {
		api.flowPath = path
	}
}

func TradeProductionMode(production bool) TradeOption {
	return func(api *CTradeApi) {
		api.production = production
	}
}

type CTradeApi struct {
	apiPtr     uintptr
	spi        thost.CHSTradeSpi
	handle     cgo.Handle
	flowPath   string
	production bool
}
