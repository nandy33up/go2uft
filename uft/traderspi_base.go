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

var _ thost.CHSTradeSpi = &BaseTradeSpi{}

type BaseTradeSpi struct {
	/// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontConnectedCallback func()

	/// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nResult 错误原因
	OnFrontDisconnectedCallback func(int)

	/// 客户端认证响应
	OnRspAuthenticateCallback func(*thost.CHSRspAuthenticateField, *thost.CHSRspInfoField, int, bool)

	/// 用户系统信息上报响应
	OnRspSubmitUserSystemInfoCallback func(*thost.CHSRspUserSystemInfoField, *thost.CHSRspInfoField, int, bool)

	/// 登录请求响应
	OnRspUserLoginCallback func(*thost.CHSRspUserLoginField, *thost.CHSRspInfoField, int, bool)

	/// 用户口令更新请求响应
	OnRspUserPasswordUpdateCallback func(*thost.CHSRspUserPasswordUpdateField, *thost.CHSRspInfoField, int, bool)

	/// 报单录入错误响应
	OnRspErrorOrderInsertCallback func(*thost.CHSRspOrderInsertField, *thost.CHSRspInfoField, int, bool)

	/// 报单操作请求响应
	OnRspOrderActionCallback func(*thost.CHSRspOrderActionField, *thost.CHSRspInfoField, int, bool)

	/// 行权录入错误响应
	OnRspErrorExerciseOrderInsertCallback func(*thost.CHSRspExerciseOrderInsertField, *thost.CHSRspInfoField, int, bool)

	/// 行权操作请求响应
	OnRspExerciseOrderActionCallback func(*thost.CHSRspExerciseOrderActionField, *thost.CHSRspInfoField, int, bool)

	/// 锁仓录入错误响应
	OnRspErrorLockInsertCallback func(*thost.CHSRspLockInsertField, *thost.CHSRspInfoField, int, bool)

	/// 询价录入响应
	OnRspForQuoteInsertCallback func(*thost.CHSRspForQuoteInsertField, *thost.CHSRspInfoField, int, bool)

	/// 报价录入错误响应
	OnRspErrorQuoteInsertCallback func(*thost.CHSRspQuoteInsertField, *thost.CHSRspInfoField, int, bool)

	/// 报价操作请求响应
	OnRspQuoteActionCallback func(*thost.CHSRspQuoteActionField, *thost.CHSRspInfoField, int, bool)

	/// 组合操作录入错误响应
	OnRspErrorCombActionInsertCallback func(*thost.CHSRspCombActionInsertField, *thost.CHSRspInfoField, int, bool)

	/// 查询最大报单数量响应
	OnRspQueryMaxOrderVolumeCallback func(*thost.CHSRspQueryMaxOrderVolumeField, *thost.CHSRspInfoField, int, bool)

	/// 查询锁仓数量响应
	OnRspQryLockVolumeCallback func(*thost.CHSRspQryLockVolumeField, *thost.CHSRspInfoField, int, bool)

	/// 查询行权数量响应
	OnRspQueryExerciseVolumeCallback func(*thost.CHSRspQueryExerciseVolumeField, *thost.CHSRspInfoField, int, bool)

	/// 查询组合数量响应
	OnRspQryCombVolumeCallback func(*thost.CHSRspQryCombVolumeField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询持仓响应
	OnRspQryPositionCallback func(*thost.CHSRspQryPositionField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询资金账户响应
	OnRspQryTradingAccountCallback func(*thost.CHSRspQryTradingAccountField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询报单响应
	OnRspQryOrderCallback func(*thost.CHSOrderField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询成交响应
	OnRspQryTradeCallback func(*thost.CHSTradeField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询行权响应
	OnRspQryExerciseCallback func(*thost.CHSExerciseField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询锁仓响应
	OnRspQryLockCallback func(*thost.CHSLockField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询组合响应
	OnRspQryCombActionCallback func(*thost.CHSCombActionField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询询价响应
	OnRspQryForQuoteCallback func(*thost.CHSForQuoteField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询报价响应
	OnRspQryQuoteCallback func(*thost.CHSQuoteField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询持仓组合明细响应
	OnRspQryPositionCombineDetailCallback func(*thost.CHSRspQryPositionCombineDetailField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询合约响应
	OnRspQryInstrumentCallback func(*thost.CHSRspQryInstrumentField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询融券响应
	OnRspQryCoveredShortCallback func(*thost.CHSRspQryCoveredShortField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询行权指派响应
	OnRspQryExerciseAssignCallback func(*thost.CHSRspQryExerciseAssignField, *thost.CHSRspInfoField, int, bool)

	/// 请求转账响应
	OnRspTransferCallback func(*thost.CHSRspTransferField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询转账响应
	OnRspQryTransferCallback func(*thost.CHSRspQryTransferField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询银行余额响应
	OnRspQueryBankBalanceCallback func(*thost.CHSRspQueryBankBalanceField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询银行账户响应
	OnRspQueryBankAccountCallback func(*thost.CHSRspQueryBankAccountField, *thost.CHSRspInfoField, int, bool)

	/// 请求多中心资金调拨响应
	OnRspMultiCentreFundTransCallback func(*thost.CHSRspMultiCentreFundTransField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询账单内容响应
	OnRspQueryBillContentCallback func(*thost.CHSRspQueryBillContentField, *thost.CHSRspInfoField, int, bool)

	/// 请求账单确认响应
	OnRspBillConfirmCallback func(*thost.CHSRspBillConfirmField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询保证金响应
	OnRspQryMarginCallback func(*thost.CHSRspQryMarginField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询手续费响应
	OnRspQryCommissionCallback func(*thost.CHSRspQryCommissionField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询持仓明细响应
	OnRspQryPositionDetailCallback func(*thost.CHSRspQryPositionDetailField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询汇率响应
	OnRspQryExchangeRateCallback func(*thost.CHSRspQryExchangeRateField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询系统配置响应
	OnRspQrySysConfigCallback func(*thost.CHSRspQrySysConfigField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询行情响应
	OnRspQryDepthMarketDataCallback func(*thost.CHSDepthMarketDataField, *thost.CHSRspInfoField, int, bool)

	/// 请求资金转账响应
	OnRspFundTransCallback func(*thost.CHSRspFundTransField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询资金转账响应
	OnRspQryFundTransCallback func(*thost.CHSRspQryFundTransField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询客户通知响应
	OnRspQryClientNoticeCallback func(*thost.CHSClientNoticeField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询期权标的响应
	OnRspQryOptUnderlyCallback func(*thost.CHSRspQryOptUnderlyField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询证券行情响应
	OnRspQrySecuDepthMarketCallback func(*thost.CHSRspQrySecuDepthMarketField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询历史报单响应
	OnRspQryHistOrderCallback func(*thost.CHSOrderField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询历史成交响应
	OnRspQryHistTradeCallback func(*thost.CHSTradeField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询提款响应
	OnRspQryWithdrawFundCallback func(*thost.CHSRspQryWithdrawFundField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询组合合约响应
	OnRspQryCombInstrumentCallback func(*thost.CHSRspQryCombInstrumentField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询席位响应
	OnRspQrySeatIDCallback func(*thost.CHSRspQrySeatIDField, *thost.CHSRspInfoField, int, bool)

	/// 请求期权自对冲响应
	OnRspOptionSelfCloseCallback func(*thost.CHSRspOptionSelfCloseField, *thost.CHSRspInfoField, int, bool)

	/// 请求期权自对冲操作响应
	OnRspOptionSelfCloseActionCallback func(*thost.CHSRspOptionSelfCloseActionField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询期权自对冲结果响应
	OnRspQryOptionSelfCloseResultCallback func(*thost.CHSRspQryOptionSelfCloseResultField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询期权自对冲响应
	OnRspQryOptionSelfCloseCallback func(*thost.CHSOptionSelfCloseField, *thost.CHSRspInfoField, int, bool)

	/// 请求期权报价录入错误响应
	OnRspErrorOptQuoteInsertCallback func(*thost.CHSRspOptQuoteInsertField, *thost.CHSRspInfoField, int, bool)

	/// 请求期权报价操作响应
	OnRspOptQuoteActionCallback func(*thost.CHSRspOptQuoteActionField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询期权报价响应
	OnRspQryOptQuoteCallback func(*thost.CHSOptQuoteField, *thost.CHSRspInfoField, int, bool)

	/// 请求查询期权组合策略响应
	OnRspQryOptCombStrategyCallback func(*thost.CHSRspQryOptCombStrategyField, *thost.CHSRspInfoField, int, bool)

	/// 成交通知
	OnRtnTradeCallback func(*thost.CHSTradeField)

	/// 报单通知
	OnRtnOrderCallback func(*thost.CHSOrderField)

	/// 行权通知
	OnRtnExerciseCallback func(*thost.CHSExerciseField)

	/// 组合操作通知
	OnRtnCombActionCallback func(*thost.CHSCombActionField)

	/// 锁仓通知
	OnRtnLockCallback func(*thost.CHSLockField)

	/// 报单操作错误回报
	OnErrRtnOrderActionCallback func(*thost.CHSOrderActionField)

	/// 客户通知通知
	OnRtnClientNoticeCallback func(*thost.CHSClientNoticeField)

	/// 询价通知
	OnRtnForQuoteCallback func(*thost.CHSForQuoteField)

	/// 报价通知
	OnRtnQuoteCallback func(*thost.CHSQuoteField)

	/// 交易所状态通知
	OnRtnExchangeStatusCallback func(*thost.CHSExchangeStatusField)

	/// 产品状态通知
	OnRtnProductStatusCallback func(*thost.CHSProductStatusField)

	/// 期权自对冲通知
	OnRtnOptionSelfCloseCallback func(*thost.CHSOptionSelfCloseField)

	/// 期权报价通知
	OnRtnOptQuoteCallback func(*thost.CHSOptQuoteField)

	/// 转账通知
	OnRtnTransferCallback func(*thost.CHSTransferField)

	/// 期权报价操作错误回报
	OnErrRtnOptQuoteActionCallback func(*thost.CHSOptQuoteActionField)
}

// / 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *BaseTradeSpi) OnFrontConnected() {
	if s.OnFrontConnectedCallback != nil {
		s.OnFrontConnectedCallback()
	}
}

// / 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
// /@param nResult 错误原因
func (s *BaseTradeSpi) OnFrontDisconnected(nResult int) {
	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(nResult)
	}
}

// / 客户端认证响应
func (s *BaseTradeSpi) OnRspAuthenticate(pRspAuthenticate *thost.CHSRspAuthenticateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspAuthenticateCallback != nil {
		s.OnRspAuthenticateCallback(pRspAuthenticate, pRspInfo, nRequestID, bIsLast)
	}
}

// / 用户系统信息上报响应
func (s *BaseTradeSpi) OnRspSubmitUserSystemInfo(pRspUserSystemInfo *thost.CHSRspUserSystemInfoField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspSubmitUserSystemInfoCallback != nil {
		s.OnRspSubmitUserSystemInfoCallback(pRspUserSystemInfo, pRspInfo, nRequestID, bIsLast)
	}
}

// / 登录请求响应
func (s *BaseTradeSpi) OnRspUserLogin(pRspUserLogin *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserLoginCallback != nil {
		s.OnRspUserLoginCallback(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
	}
}

// / 用户口令更新请求响应
func (s *BaseTradeSpi) OnRspUserPasswordUpdate(pRspUserPasswordUpdate *thost.CHSRspUserPasswordUpdateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserPasswordUpdateCallback != nil {
		s.OnRspUserPasswordUpdateCallback(pRspUserPasswordUpdate, pRspInfo, nRequestID, bIsLast)
	}
}

// / 报单录入错误响应
func (s *BaseTradeSpi) OnRspErrorOrderInsert(pRspOrderInsert *thost.CHSRspOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorOrderInsertCallback != nil {
		s.OnRspErrorOrderInsertCallback(pRspOrderInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 报单操作请求响应
func (s *BaseTradeSpi) OnRspOrderAction(pRspOrderAction *thost.CHSRspOrderActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOrderActionCallback != nil {
		s.OnRspOrderActionCallback(pRspOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 行权录入错误响应
func (s *BaseTradeSpi) OnRspErrorExerciseOrderInsert(pRspExerciseOrderInsert *thost.CHSRspExerciseOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorExerciseOrderInsertCallback != nil {
		s.OnRspErrorExerciseOrderInsertCallback(pRspExerciseOrderInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 行权操作请求响应
func (s *BaseTradeSpi) OnRspExerciseOrderAction(pRspExerciseOrderAction *thost.CHSRspExerciseOrderActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspExerciseOrderActionCallback != nil {
		s.OnRspExerciseOrderActionCallback(pRspExerciseOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 锁仓录入错误响应
func (s *BaseTradeSpi) OnRspErrorLockInsert(pRspLockInsert *thost.CHSRspLockInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorLockInsertCallback != nil {
		s.OnRspErrorLockInsertCallback(pRspLockInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 询价录入响应
func (s *BaseTradeSpi) OnRspForQuoteInsert(pRspForQuoteInsert *thost.CHSRspForQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteInsertCallback != nil {
		s.OnRspForQuoteInsertCallback(pRspForQuoteInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 报价录入错误响应
func (s *BaseTradeSpi) OnRspErrorQuoteInsert(pRspQuoteInsert *thost.CHSRspQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorQuoteInsertCallback != nil {
		s.OnRspErrorQuoteInsertCallback(pRspQuoteInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 报价操作请求响应
func (s *BaseTradeSpi) OnRspQuoteAction(pRspQuoteAction *thost.CHSRspQuoteActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQuoteActionCallback != nil {
		s.OnRspQuoteActionCallback(pRspQuoteAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 组合操作录入错误响应
func (s *BaseTradeSpi) OnRspErrorCombActionInsert(pRspCombActionInsert *thost.CHSRspCombActionInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorCombActionInsertCallback != nil {
		s.OnRspErrorCombActionInsertCallback(pRspCombActionInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 查询最大报单数量响应
func (s *BaseTradeSpi) OnRspQueryMaxOrderVolume(pRspQueryMaxOrderVolume *thost.CHSRspQueryMaxOrderVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryMaxOrderVolumeCallback != nil {
		s.OnRspQueryMaxOrderVolumeCallback(pRspQueryMaxOrderVolume, pRspInfo, nRequestID, bIsLast)
	}
}

// / 查询锁仓数量响应
func (s *BaseTradeSpi) OnRspQryLockVolume(pRspQryLockVolume *thost.CHSRspQryLockVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryLockVolumeCallback != nil {
		s.OnRspQryLockVolumeCallback(pRspQryLockVolume, pRspInfo, nRequestID, bIsLast)
	}
}

// / 查询行权数量响应
func (s *BaseTradeSpi) OnRspQueryExerciseVolume(pRspQueryExerciseVolume *thost.CHSRspQueryExerciseVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryExerciseVolumeCallback != nil {
		s.OnRspQueryExerciseVolumeCallback(pRspQueryExerciseVolume, pRspInfo, nRequestID, bIsLast)
	}
}

// / 查询组合数量响应
func (s *BaseTradeSpi) OnRspQryCombVolume(pRspQryCombVolume *thost.CHSRspQryCombVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombVolumeCallback != nil {
		s.OnRspQryCombVolumeCallback(pRspQryCombVolume, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询持仓响应
func (s *BaseTradeSpi) OnRspQryPosition(pRspQryPosition *thost.CHSRspQryPositionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryPositionCallback != nil {
		s.OnRspQryPositionCallback(pRspQryPosition, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询资金账户响应
func (s *BaseTradeSpi) OnRspQryTradingAccount(pRspQryTradingAccount *thost.CHSRspQryTradingAccountField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradingAccountCallback != nil {
		s.OnRspQryTradingAccountCallback(pRspQryTradingAccount, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询报单响应
func (s *BaseTradeSpi) OnRspQryOrder(pRspQryOrder *thost.CHSOrderField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOrderCallback != nil {
		s.OnRspQryOrderCallback(pRspQryOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询成交响应
func (s *BaseTradeSpi) OnRspQryTrade(pRspQryTrade *thost.CHSTradeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradeCallback != nil {
		s.OnRspQryTradeCallback(pRspQryTrade, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询行权响应
func (s *BaseTradeSpi) OnRspQryExercise(pRspQryExercise *thost.CHSExerciseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExerciseCallback != nil {
		s.OnRspQryExerciseCallback(pRspQryExercise, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询锁仓响应
func (s *BaseTradeSpi) OnRspQryLock(pRspQryLock *thost.CHSLockField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryLockCallback != nil {
		s.OnRspQryLockCallback(pRspQryLock, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询组合响应
func (s *BaseTradeSpi) OnRspQryCombAction(pRspQryCombAction *thost.CHSCombActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombActionCallback != nil {
		s.OnRspQryCombActionCallback(pRspQryCombAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询询价响应
func (s *BaseTradeSpi) OnRspQryForQuote(pRspQryForQuote *thost.CHSForQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryForQuoteCallback != nil {
		s.OnRspQryForQuoteCallback(pRspQryForQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询报价响应
func (s *BaseTradeSpi) OnRspQryQuote(pRspQryQuote *thost.CHSQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryQuoteCallback != nil {
		s.OnRspQryQuoteCallback(pRspQryQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询持仓组合明细响应
func (s *BaseTradeSpi) OnRspQryPositionCombineDetail(pRspQryPositionCombineDetail *thost.CHSRspQryPositionCombineDetailField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryPositionCombineDetailCallback != nil {
		s.OnRspQryPositionCombineDetailCallback(pRspQryPositionCombineDetail, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询合约响应
func (s *BaseTradeSpi) OnRspQryInstrument(pRspQryInstrument *thost.CHSRspQryInstrumentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInstrumentCallback != nil {
		s.OnRspQryInstrumentCallback(pRspQryInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询融券响应
func (s *BaseTradeSpi) OnRspQryCoveredShort(pRspQryCoveredShort *thost.CHSRspQryCoveredShortField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCoveredShortCallback != nil {
		s.OnRspQryCoveredShortCallback(pRspQryCoveredShort, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询行权指派响应
func (s *BaseTradeSpi) OnRspQryExerciseAssign(pRspQryExerciseAssign *thost.CHSRspQryExerciseAssignField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExerciseAssignCallback != nil {
		s.OnRspQryExerciseAssignCallback(pRspQryExerciseAssign, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求转账响应
func (s *BaseTradeSpi) OnRspTransfer(pRspTransfer *thost.CHSRspTransferField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspTransferCallback != nil {
		s.OnRspTransferCallback(pRspTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询转账响应
func (s *BaseTradeSpi) OnRspQryTransfer(pRspQryTransfer *thost.CHSRspQryTransferField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTransferCallback != nil {
		s.OnRspQryTransferCallback(pRspQryTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询银行余额响应
func (s *BaseTradeSpi) OnRspQueryBankBalance(pRspQueryBankBalance *thost.CHSRspQueryBankBalanceField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryBankBalanceCallback != nil {
		s.OnRspQueryBankBalanceCallback(pRspQueryBankBalance, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询银行账户响应
func (s *BaseTradeSpi) OnRspQueryBankAccount(pRspQueryBankAccount *thost.CHSRspQueryBankAccountField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryBankAccountCallback != nil {
		s.OnRspQueryBankAccountCallback(pRspQueryBankAccount, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求多中心资金调拨响应
func (s *BaseTradeSpi) OnRspMultiCentreFundTrans(pRspMultiCentreFundTransfer *thost.CHSRspMultiCentreFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspMultiCentreFundTransCallback != nil {
		s.OnRspMultiCentreFundTransCallback(pRspMultiCentreFundTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询账单内容响应
func (s *BaseTradeSpi) OnRspQueryBillContent(pRspQueryBillContent *thost.CHSRspQueryBillContentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryBillContentCallback != nil {
		s.OnRspQueryBillContentCallback(pRspQueryBillContent, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求账单确认响应
func (s *BaseTradeSpi) OnRspBillConfirm(pRspBillConfirm *thost.CHSRspBillConfirmField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspBillConfirmCallback != nil {
		s.OnRspBillConfirmCallback(pRspBillConfirm, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询保证金响应
func (s *BaseTradeSpi) OnRspQryMargin(pRspQryMargin *thost.CHSRspQryMarginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryMarginCallback != nil {
		s.OnRspQryMarginCallback(pRspQryMargin, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询手续费响应
func (s *BaseTradeSpi) OnRspQryCommission(pRspQryCommission *thost.CHSRspQryCommissionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCommissionCallback != nil {
		s.OnRspQryCommissionCallback(pRspQryCommission, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询持仓明细响应
func (s *BaseTradeSpi) OnRspQryPositionDetail(pRspQryPositionDetail *thost.CHSRspQryPositionDetailField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryPositionDetailCallback != nil {
		s.OnRspQryPositionDetailCallback(pRspQryPositionDetail, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询汇率响应
func (s *BaseTradeSpi) OnRspQryExchangeRate(pRspQryExchangeRate *thost.CHSRspQryExchangeRateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExchangeRateCallback != nil {
		s.OnRspQryExchangeRateCallback(pRspQryExchangeRate, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询系统配置响应
func (s *BaseTradeSpi) OnRspQrySysConfig(pRspQrySysConfig *thost.CHSRspQrySysConfigField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySysConfigCallback != nil {
		s.OnRspQrySysConfigCallback(pRspQrySysConfig, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询行情响应
func (s *BaseTradeSpi) OnRspQryDepthMarketData(pRspQryDepthMarketData *thost.CHSDepthMarketDataField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryDepthMarketDataCallback != nil {
		s.OnRspQryDepthMarketDataCallback(pRspQryDepthMarketData, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求资金转账响应
func (s *BaseTradeSpi) OnRspFundTrans(pRspFundTransfer *thost.CHSRspFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspFundTransCallback != nil {
		s.OnRspFundTransCallback(pRspFundTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询资金转账响应
func (s *BaseTradeSpi) OnRspQryFundTrans(pRspQryFundTrans *thost.CHSRspQryFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryFundTransCallback != nil {
		s.OnRspQryFundTransCallback(pRspQryFundTrans, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询客户通知响应
func (s *BaseTradeSpi) OnRspQryClientNotice(pRspQryClientNotice *thost.CHSClientNoticeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryClientNoticeCallback != nil {
		s.OnRspQryClientNoticeCallback(pRspQryClientNotice, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询期权标的响应
func (s *BaseTradeSpi) OnRspQryOptUnderly(pRspQryOptUnderly *thost.CHSRspQryOptUnderlyField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptUnderlyCallback != nil {
		s.OnRspQryOptUnderlyCallback(pRspQryOptUnderly, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询证券行情响应
func (s *BaseTradeSpi) OnRspQrySecuDepthMarket(pRspQrySecuDepthMarket *thost.CHSRspQrySecuDepthMarketField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySecuDepthMarketCallback != nil {
		s.OnRspQrySecuDepthMarketCallback(pRspQrySecuDepthMarket, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询历史报单响应
func (s *BaseTradeSpi) OnRspQryHistOrder(pRspQryHistOrder *thost.CHSOrderField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryHistOrderCallback != nil {
		s.OnRspQryHistOrderCallback(pRspQryHistOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询历史成交响应
func (s *BaseTradeSpi) OnRspQryHistTrade(pRspQryHistTrade *thost.CHSTradeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryHistTradeCallback != nil {
		s.OnRspQryHistTradeCallback(pRspQryHistTrade, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询提款响应
func (s *BaseTradeSpi) OnRspQryWithdrawFund(pRspQryWithdrawFund *thost.CHSRspQryWithdrawFundField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryWithdrawFundCallback != nil {
		s.OnRspQryWithdrawFundCallback(pRspQryWithdrawFund, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询组合合约响应
func (s *BaseTradeSpi) OnRspQryCombInstrument(pRspQryCombInstrument *thost.CHSRspQryCombInstrumentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombInstrumentCallback != nil {
		s.OnRspQryCombInstrumentCallback(pRspQryCombInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询席位响应
func (s *BaseTradeSpi) OnRspQrySeatID(pRspQrySeatID *thost.CHSRspQrySeatIDField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySeatIDCallback != nil {
		s.OnRspQrySeatIDCallback(pRspQrySeatID, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求期权自对冲响应
func (s *BaseTradeSpi) OnRspOptionSelfClose(pReqOptionSelfClose *thost.CHSRspOptionSelfCloseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOptionSelfCloseCallback != nil {
		s.OnRspOptionSelfCloseCallback(pReqOptionSelfClose, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求期权自对冲操作响应
func (s *BaseTradeSpi) OnRspOptionSelfCloseAction(pReqOptionSelfCloseAction *thost.CHSRspOptionSelfCloseActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOptionSelfCloseActionCallback != nil {
		s.OnRspOptionSelfCloseActionCallback(pReqOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询期权自对冲结果响应
func (s *BaseTradeSpi) OnRspQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *thost.CHSRspQryOptionSelfCloseResultField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptionSelfCloseResultCallback != nil {
		s.OnRspQryOptionSelfCloseResultCallback(pReqQryOptionSelfCloseResult, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询期权自对冲响应
func (s *BaseTradeSpi) OnRspQryOptionSelfClose(pRspQryOptionSelfClose *thost.CHSOptionSelfCloseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptionSelfCloseCallback != nil {
		s.OnRspQryOptionSelfCloseCallback(pRspQryOptionSelfClose, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求期权报价录入错误响应
func (s *BaseTradeSpi) OnRspErrorOptQuoteInsert(pRspOptQuoteInsert *thost.CHSRspOptQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorOptQuoteInsertCallback != nil {
		s.OnRspErrorOptQuoteInsertCallback(pRspOptQuoteInsert, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求期权报价操作响应
func (s *BaseTradeSpi) OnRspOptQuoteAction(pRspOptQuoteAction *thost.CHSRspOptQuoteActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOptQuoteActionCallback != nil {
		s.OnRspOptQuoteActionCallback(pRspOptQuoteAction, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询期权报价响应
func (s *BaseTradeSpi) OnRspQryOptQuote(pRspQryOptQuote *thost.CHSOptQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptQuoteCallback != nil {
		s.OnRspQryOptQuoteCallback(pRspQryOptQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// / 请求查询期权组合策略响应
func (s *BaseTradeSpi) OnRspQryOptCombStrategy(pRspQryOptCombStrategy *thost.CHSRspQryOptCombStrategyField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptCombStrategyCallback != nil {
		s.OnRspQryOptCombStrategyCallback(pRspQryOptCombStrategy, pRspInfo, nRequestID, bIsLast)
	}
}

// / 成交通知
func (s *BaseTradeSpi) OnRtnTrade(pRtnTrade *thost.CHSTradeField) {
	if s.OnRtnTradeCallback != nil {
		s.OnRtnTradeCallback(pRtnTrade)
	}
}

// / 报单通知
func (s *BaseTradeSpi) OnRtnOrder(pRtnOrder *thost.CHSOrderField) {
	if s.OnRtnOrderCallback != nil {
		s.OnRtnOrderCallback(pRtnOrder)
	}
}

// / 行权通知
func (s *BaseTradeSpi) OnRtnExercise(pRtnExercise *thost.CHSExerciseField) {
	if s.OnRtnExerciseCallback != nil {
		s.OnRtnExerciseCallback(pRtnExercise)
	}
}

// / 组合操作通知
func (s *BaseTradeSpi) OnRtnCombAction(pRtnCombAction *thost.CHSCombActionField) {
	if s.OnRtnCombActionCallback != nil {
		s.OnRtnCombActionCallback(pRtnCombAction)
	}
}

// / 锁仓通知
func (s *BaseTradeSpi) OnRtnLock(pRtnLock *thost.CHSLockField) {
	if s.OnRtnLockCallback != nil {
		s.OnRtnLockCallback(pRtnLock)
	}
}

// / 报单操作错误回报
func (s *BaseTradeSpi) OnErrRtnOrderAction(pRtnOrder *thost.CHSOrderActionField) {
	if s.OnErrRtnOrderActionCallback != nil {
		s.OnErrRtnOrderActionCallback(pRtnOrder)
	}
}

// / 客户通知通知
func (s *BaseTradeSpi) OnRtnClientNotice(pRtnClientNotice *thost.CHSClientNoticeField) {
	if s.OnRtnClientNoticeCallback != nil {
		s.OnRtnClientNoticeCallback(pRtnClientNotice)
	}
}

// / 询价通知
func (s *BaseTradeSpi) OnRtnForQuote(pRtnForQuote *thost.CHSForQuoteField) {
	if s.OnRtnForQuoteCallback != nil {
		s.OnRtnForQuoteCallback(pRtnForQuote)
	}
}

// / 报价通知
func (s *BaseTradeSpi) OnRtnQuote(pRtnQuote *thost.CHSQuoteField) {
	if s.OnRtnQuoteCallback != nil {
		s.OnRtnQuoteCallback(pRtnQuote)
	}
}

// / 交易所状态通知
func (s *BaseTradeSpi) OnRtnExchangeStatus(pRtnExchangeStatus *thost.CHSExchangeStatusField) {
	if s.OnRtnExchangeStatusCallback != nil {
		s.OnRtnExchangeStatusCallback(pRtnExchangeStatus)
	}
}

// / 产品状态通知
func (s *BaseTradeSpi) OnRtnProductStatus(pRtnProductStatus *thost.CHSProductStatusField) {
	if s.OnRtnProductStatusCallback != nil {
		s.OnRtnProductStatusCallback(pRtnProductStatus)
	}
}

// / 期权自对冲通知
func (s *BaseTradeSpi) OnRtnOptionSelfClose(pRtnOptionSelfClose *thost.CHSOptionSelfCloseField) {
	if s.OnRtnOptionSelfCloseCallback != nil {
		s.OnRtnOptionSelfCloseCallback(pRtnOptionSelfClose)
	}
}

// / 期权报价通知
func (s *BaseTradeSpi) OnRtnOptQuote(pRtnOptQuote *thost.CHSOptQuoteField) {
	if s.OnRtnOptQuoteCallback != nil {
		s.OnRtnOptQuoteCallback(pRtnOptQuote)
	}
}

// / 转账通知
func (s *BaseTradeSpi) OnRtnTransfer(pRtnTransfer *thost.CHSTransferField) {
	if s.OnRtnTransferCallback != nil {
		s.OnRtnTransferCallback(pRtnTransfer)
	}
}

// / 期权报价操作错误回报
func (s *BaseTradeSpi) OnErrRtnOptQuoteAction(pRtnOptQuoteAction *thost.CHSOptQuoteActionField) {
	if s.OnErrRtnOptQuoteActionCallback != nil {
		s.OnErrRtnOptQuoteActionCallback(pRtnOptQuoteAction)
	}
}
