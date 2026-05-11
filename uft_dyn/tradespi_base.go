package uft_dyn

import (
	"github.com/nandy33up/go2uft/thost"
)

var _ thost.CHSTradeSpi = &BaseTradeSpi{}

type BaseTradeSpi struct {
	OnFrontConnectedCallback         func()
	OnFrontDisconnectedCallback      func(int)
	OnHeartBeatWarningCallback       func(int)
	OnRspErrorCallback               func(*thost.CHSRspInfoField, int, bool)
	OnRspUserLoginCallback           func(*thost.CHSRspUserLoginField, *thost.CHSRspInfoField, int, bool)
	OnRspUserSystemInfoCallback      func(*thost.CHSRspUserSystemInfoField, *thost.CHSRspInfoField, int, bool)
	OnRspOrderInsertCallback         func(*thost.CHSRspOrderInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspOrderActionCallback         func(*thost.CHSRspOrderActionField, *thost.CHSRspInfoField, int, bool)
	OnRtnOrderCallback               func(*thost.CHSOrderField)
	OnRtnTradeCallback               func(*thost.CHSTradeField)
	OnRspQueryMaxOrderVolumeCallback func(*thost.CHSRspQueryMaxOrderVolumeField, *thost.CHSRspInfoField, int, bool)
	OnRspQueryExerciseVolumeCallback func(*thost.CHSRspQueryExerciseVolumeField, *thost.CHSRspInfoField, int, bool)
	OnRspQryLockVolumeCallback       func(*thost.CHSRspQryLockVolumeField, *thost.CHSRspInfoField, int, bool)
	OnRspQryCombVolumeCallback       func(*thost.CHSRspQryCombVolumeField, *thost.CHSRspInfoField, int, bool)
	OnRspQryPositionCallback         func(*thost.CHSRspQryPositionField, *thost.CHSRspInfoField, int, bool)
	OnRspQryTradingAccountCallback   func(*thost.CHSRspQryTradingAccountField, *thost.CHSRspInfoField, int, bool)
	OnRspExerciseOrderInsertCallback func(*thost.CHSRspExerciseOrderInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspExerciseOrderActionCallback func(*thost.CHSRspExerciseOrderActionField, *thost.CHSRspInfoField, int, bool)
	OnRspLockInsertCallback          func(*thost.CHSRspLockInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspForQuoteInsertCallback      func(*thost.CHSRspForQuoteInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspQuoteInsertCallback         func(*thost.CHSRspQuoteInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspQuoteActionCallback         func(*thost.CHSRspQuoteActionField, *thost.CHSRspInfoField, int, bool)
	OnRspCombActionInsertCallback    func(*thost.CHSRspCombActionInsertField, *thost.CHSRspInfoField, int, bool)
	OnRspUserPasswordUpdateCallback  func(*thost.CHSRspUserPasswordUpdateField, *thost.CHSRspInfoField, int, bool)
	OnRspAuthenticateCallback        func(*thost.CHSRspAuthenticateField, *thost.CHSRspInfoField, int, bool)
}

func (s *BaseTradeSpi) OnFrontConnected() {
	if s.OnFrontConnectedCallback != nil {
		s.OnFrontConnectedCallback()
	}
}

func (s *BaseTradeSpi) OnFrontDisconnected(nReason int) {
	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(nReason)
	}
}

func (s *BaseTradeSpi) OnHeartBeatWarning(nTimeLapse int) {
	if s.OnHeartBeatWarningCallback != nil {
		s.OnHeartBeatWarningCallback(nTimeLapse)
	}
}

func (s *BaseTradeSpi) OnRspError(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspUserLogin(pRspUserLogin *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserLoginCallback != nil {
		s.OnRspUserLoginCallback(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspUserLogout(pUserLogout *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorMsg(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspUserSystemInfo(pRspUserSystemInfo *thost.CHSRspUserSystemInfoField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserSystemInfoCallback != nil {
		s.OnRspUserSystemInfoCallback(pRspUserSystemInfo, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspOrderInsert(pRspOrderInsert *thost.CHSRspOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOrderInsertCallback != nil {
		s.OnRspOrderInsertCallback(pRspOrderInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspOrderAction(pRspOrderAction *thost.CHSRspOrderActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOrderActionCallback != nil {
		s.OnRspOrderActionCallback(pRspOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRtnOrder(pOrder *thost.CHSOrderField) {
	if s.OnRtnOrderCallback != nil {
		s.OnRtnOrderCallback(pOrder)
	}
}

func (s *BaseTradeSpi) OnRtnTrade(pTrade *thost.CHSTradeField) {
	if s.OnRtnTradeCallback != nil {
		s.OnRtnTradeCallback(pTrade)
	}
}

func (s *BaseTradeSpi) OnRspQueryMaxOrderVolume(pRspQueryMaxOrderVolume *thost.CHSRspQueryMaxOrderVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryMaxOrderVolumeCallback != nil {
		s.OnRspQueryMaxOrderVolumeCallback(pRspQueryMaxOrderVolume, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQueryExerciseVolume(pRspQueryExerciseVolume *thost.CHSRspQueryExerciseVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryExerciseVolumeCallback != nil {
		s.OnRspQueryExerciseVolumeCallback(pRspQueryExerciseVolume, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQryLockVolume(pRspQryLockVolume *thost.CHSRspQryLockVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryLockVolumeCallback != nil {
		s.OnRspQryLockVolumeCallback(pRspQryLockVolume, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQryCombVolume(pRspQryCombVolume *thost.CHSRspQryCombVolumeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombVolumeCallback != nil {
		s.OnRspQryCombVolumeCallback(pRspQryCombVolume, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQryPosition(pRspQryPosition *thost.CHSRspQryPositionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryPositionCallback != nil {
		s.OnRspQryPositionCallback(pRspQryPosition, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQryTradingAccount(pRspQryTradingAccount *thost.CHSRspQryTradingAccountField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradingAccountCallback != nil {
		s.OnRspQryTradingAccountCallback(pRspQryTradingAccount, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQryOrder(pRspQryOrder *thost.CHSOrderField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryTrade(pRspQryTrade *thost.CHSTradeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryExercise(pRspQryExercise *thost.CHSExerciseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryLock(pRspQryLock *thost.CHSLockField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryCombAction(pRspQryCombAction *thost.CHSCombActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryForQuote(pRspQryForQuote *thost.CHSForQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryQuote(pRspQryQuote *thost.CHSQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryPositionCombineDetail(pRspQryPositionCombineDetail *thost.CHSRspQryPositionCombineDetailField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryInstrument(pRspQryInstrument *thost.CHSRspQryInstrumentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryCoveredShort(pRspQryCoveredShort *thost.CHSRspQryCoveredShortField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryExerciseAssign(pRspQryExerciseAssign *thost.CHSRspQryExerciseAssignField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspTransfer(pRspTransfer *thost.CHSRspTransferField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryTransfer(pRspQryTransfer *thost.CHSRspQryTransferField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQueryBankBalance(pRspQueryBankBalance *thost.CHSRspQueryBankBalanceField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQueryBankAccount(pRspQueryBankAccount *thost.CHSRspQueryBankAccountField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspMultiCentreFundTrans(pRspMultiCentreFundTransfer *thost.CHSRspMultiCentreFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQueryBillContent(pRspQueryBillContent *thost.CHSRspQueryBillContentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspBillConfirm(pRspBillConfirm *thost.CHSRspBillConfirmField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryMargin(pRspQryMargin *thost.CHSRspQryMarginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryCommission(pRspQryCommission *thost.CHSRspQryCommissionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryExchangeRate(pRspQryExchangeRate *thost.CHSRspQryExchangeRateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryPositionDetail(pRspQryPositionDetail *thost.CHSRspQryPositionDetailField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQrySysConfig(pRspQrySysConfig *thost.CHSRspQrySysConfigField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryDepthMarketData(pRspQryDepthMarketData *thost.CHSDepthMarketDataField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspFundTrans(pRspFundTransfer *thost.CHSRspFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryFundTrans(pRspQryFundTrans *thost.CHSRspQryFundTransField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryClientNotice(pRspQryClientNotice *thost.CHSClientNoticeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryOptUnderly(pRspQryOptUnderly *thost.CHSRspQryOptUnderlyField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQrySecuDepthMarket(pRspQrySecuDepthMarket *thost.CHSRspQrySecuDepthMarketField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryHistOrder(pRspQryHistOrder *thost.CHSOrderField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryHistTrade(pRspQryHistTrade *thost.CHSTradeField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryWithdrawFund(pRspQryWithdrawFund *thost.CHSRspQryWithdrawFundField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryCombInstrument(pRspQryCombInstrument *thost.CHSRspQryCombInstrumentField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQrySeatID(pRspQrySeatID *thost.CHSRspQrySeatIDField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspOptionSelfClose(pReqOptionSelfClose *thost.CHSRspOptionSelfCloseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspOptionSelfCloseAction(pReqOptionSelfCloseAction *thost.CHSRspOptionSelfCloseActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *thost.CHSRspQryOptionSelfCloseResultField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryOptionSelfClose(pRspQryOptionSelfClose *thost.CHSOptionSelfCloseField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorOptQuoteInsert(pRspOptQuoteInsert *thost.CHSRspOptQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspOptQuoteAction(pRspOptQuoteAction *thost.CHSRspOptQuoteActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryOptQuote(pRspQryOptQuote *thost.CHSOptQuoteField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspQryOptCombStrategy(pRspQryOptCombStrategy *thost.CHSRspQryOptCombStrategyField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspExerciseOrderInsert(pRspExerciseOrderInsert *thost.CHSRspExerciseOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspExerciseOrderInsertCallback != nil {
		s.OnRspExerciseOrderInsertCallback(pRspExerciseOrderInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspExerciseOrderAction(pRspExerciseOrderAction *thost.CHSRspExerciseOrderActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspExerciseOrderActionCallback != nil {
		s.OnRspExerciseOrderActionCallback(pRspExerciseOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspLockInsert(pRspLockInsert *thost.CHSRspLockInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspLockInsertCallback != nil {
		s.OnRspLockInsertCallback(pRspLockInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspForQuoteInsert(pRspForQuoteInsert *thost.CHSRspForQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteInsertCallback != nil {
		s.OnRspForQuoteInsertCallback(pRspForQuoteInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQuoteInsert(pRspQuoteInsert *thost.CHSRspQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQuoteInsertCallback != nil {
		s.OnRspQuoteInsertCallback(pRspQuoteInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspQuoteAction(pRspQuoteAction *thost.CHSRspQuoteActionField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQuoteActionCallback != nil {
		s.OnRspQuoteActionCallback(pRspQuoteAction, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspCombActionInsert(pRspCombActionInsert *thost.CHSRspCombActionInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspCombActionInsertCallback != nil {
		s.OnRspCombActionInsertCallback(pRspCombActionInsert, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspUserPasswordUpdate(pRspUserPasswordUpdate *thost.CHSRspUserPasswordUpdateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserPasswordUpdateCallback != nil {
		s.OnRspUserPasswordUpdateCallback(pRspUserPasswordUpdate, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspAuthenticate(pRspAuthenticate *thost.CHSRspAuthenticateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspAuthenticateCallback != nil {
		s.OnRspAuthenticateCallback(pRspAuthenticate, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseTradeSpi) OnRspSubmitUserSystemInfo(pRspUserSystemInfo *thost.CHSRspUserSystemInfoField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorOrderInsert(pRspOrderInsert *thost.CHSRspOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorExerciseOrderInsert(pRspExerciseOrderInsert *thost.CHSRspExerciseOrderInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorLockInsert(pRspLockInsert *thost.CHSRspLockInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorQuoteInsert(pRspQuoteInsert *thost.CHSRspQuoteInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRspErrorCombActionInsert(pRspCombActionInsert *thost.CHSRspCombActionInsertField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseTradeSpi) OnRtnExercise(pRtnExercise *thost.CHSExerciseField) {
}

func (s *BaseTradeSpi) OnRtnCombAction(pRtnCombAction *thost.CHSCombActionField) {
}

func (s *BaseTradeSpi) OnRtnLock(pRtnLock *thost.CHSLockField) {
}

func (s *BaseTradeSpi) OnErrRtnOrderAction(pRtnOrder *thost.CHSOrderActionField) {
}

func (s *BaseTradeSpi) OnRtnClientNotice(pRtnClientNotice *thost.CHSClientNoticeField) {
}

func (s *BaseTradeSpi) OnRtnForQuote(pRtnForQuote *thost.CHSForQuoteField) {
}

func (s *BaseTradeSpi) OnRtnQuote(pRtnQuote *thost.CHSQuoteField) {
}

func (s *BaseTradeSpi) OnRtnExchangeStatus(pRtnExchangeStatus *thost.CHSExchangeStatusField) {
}

func (s *BaseTradeSpi) OnRtnProductStatus(pRtnProductStatus *thost.CHSProductStatusField) {
}

func (s *BaseTradeSpi) OnRtnOptionSelfClose(pRtnOptionSelfClose *thost.CHSOptionSelfCloseField) {
}

func (s *BaseTradeSpi) OnRtnOptQuote(pRtnOptQuote *thost.CHSOptQuoteField) {
}

func (s *BaseTradeSpi) OnRtnTransfer(pRtnTransfer *thost.CHSTransferField) {
}

func (s *BaseTradeSpi) OnErrRtnOptQuoteAction(pRtnOptQuoteAction *thost.CHSOptQuoteActionField) {
}
