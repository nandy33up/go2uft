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

package uft

import (
	"github.com/nandy33up/go2uft/thost"
)

var _ thost.CHSMdSpi = &BaseMdSpi{}

type BaseMdSpi struct {
	/// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontConnectedCallback func()

	/// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nResult 错误原因
	OnFrontDisconnectedCallback func(int)

	/// 订阅行情应答
	OnRspDepthMarketDataSubscribeCallback func(*thost.CHSRspInfoField, int, bool)

	/// 取消订阅行情应答
	OnRspDepthMarketDataCancelCallback func(*thost.CHSRspInfoField, int, bool)

	/// 深度行情通知
	OnRtnDepthMarketDataCallback func(*thost.CHSDepthMarketDataField)

	/// 订阅询价应答
	OnRspForQuoteSubscribeCallback func(*thost.CHSRspInfoField, int, bool)

	/// 取消订阅询价应答
	OnRspForQuoteCancelCallback func(*thost.CHSRspInfoField, int, bool)

	/// 询价通知
	OnRtnForQuoteCallback func(*thost.CHSForQuoteField)
}

// / 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *BaseMdSpi) OnFrontConnected() {
	if s.OnFrontConnectedCallback != nil {
		s.OnFrontConnectedCallback()
	}
}

// / 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
// /@param nResult 错误原因
func (s *BaseMdSpi) OnFrontDisconnected(nResult int) {
	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(nResult)
	}
}

// / 订阅行情应答
func (s *BaseMdSpi) OnRspDepthMarketDataSubscribe(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspDepthMarketDataSubscribeCallback != nil {
		s.OnRspDepthMarketDataSubscribeCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// / 取消订阅行情应答
func (s *BaseMdSpi) OnRspDepthMarketDataCancel(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspDepthMarketDataCancelCallback != nil {
		s.OnRspDepthMarketDataCancelCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// / 深度行情通知
func (s *BaseMdSpi) OnRtnDepthMarketData(pDepthMarketData *thost.CHSDepthMarketDataField) {
	if s.OnRtnDepthMarketDataCallback != nil {
		s.OnRtnDepthMarketDataCallback(pDepthMarketData)
	}
}

// / 订阅询价应答
func (s *BaseMdSpi) OnRspForQuoteSubscribe(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteSubscribeCallback != nil {
		s.OnRspForQuoteSubscribeCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// / 取消订阅询价应答
func (s *BaseMdSpi) OnRspForQuoteCancel(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteCancelCallback != nil {
		s.OnRspForQuoteCancelCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// / 询价通知
func (s *BaseMdSpi) OnRtnForQuote(pForQuote *thost.CHSForQuoteField) {
	if s.OnRtnForQuoteCallback != nil {
		s.OnRtnForQuoteCallback(pForQuote)
	}
}
