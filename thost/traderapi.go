package thost

import "unsafe"

type CHSTradeApi interface {
	GetApiVersion() string
	ReleaseApi()
	Init(pInitCfgField *CHSInitConfigField, pExtraParam unsafe.Pointer) int
	Join() int
	RegisterSubModel(eSubType SUB_TERT_TYPE) int
	RegisterFront(pszFrontAddress string) int
	RegisterFensServer(pszFensAddress string, pszAccountID string) int
	RegisterSpi(pSpi CHSTradeSpi)
	GetApiErrorMsg(nErrorCode int) string
	GetTradingDate() int
	BindSessionID(nSessionID uint8) int
	ReqAuthenticate(pReqAuthenticateField *CHSReqAuthenticateField, nRequestID int) int
	ReqSubmitUserSystemInfo(pReqUserSystemInfo *CHSReqUserSystemInfoField, nRequestID int) int
	ReqUserLogin(pReqUserLoginField *CHSReqUserLoginField, nRequestID int) int
	ReqUserPasswordUpdate(pReqUserPasswordUpdateField *CHSReqUserPasswordUpdateField, nRequestID int) int
	ReqOrderInsert(pReqOrderInsert *CHSReqOrderInsertField, nRequestID int) int
	ReqOrderAction(pReqOrderAction *CHSReqOrderActionField, nRequestID int) int
	ReqExerciseOrderInsert(pReqExerciseOrderInsert *CHSReqExerciseOrderInsertField, nRequestID int) int
	ReqExerciseOrderAction(pReqExerciseOrderActionField *CHSReqExerciseOrderActionField, nRequestID int) int
	ReqLockInsert(pReqLockInsert *CHSReqLockInsertField, nRequestID int) int
	ReqForQuoteInsert(pReqForQuoteInsertField *CHSReqForQuoteInsertField, nRequestID int) int
	ReqQuoteInsert(pReqQuoteInsert *CHSReqQuoteInsertField, nRequestID int) int
	ReqQuoteAction(pReqQuoteAction *CHSReqQuoteActionField, nRequestID int) int
	ReqCombActionInsert(pReqCombActionInsert *CHSReqCombActionInsertField, nRequestID int) int
	ReqQueryMaxOrderVolume(pReqQueryMaxOrderVolume *CHSReqQueryMaxOrderVolumeField, nRequestID int) int
	ReqQryLockVolume(pReqQryLockVolume *CHSReqQryLockVolumeField, nRequestID int) int
	ReqQueryExerciseVolume(pReqQueryExerciseVolume *CHSReqQueryExerciseVolumeField, nRequestID int) int
	ReqQryCombVolume(pReqQryCombVolume *CHSReqQryCombVolumeField, nRequestID int) int
	ReqQryPosition(pReqQryPosition *CHSReqQryPositionField, nRequestID int) int
	ReqQryTradingAccount(pReqQryTradingAccount *CHSReqQryTradingAccountField, nRequestID int) int
	ReqQryOrder(pReqQryOrder *CHSReqQryOrderField, nRequestID int) int
	ReqQryTrade(pReqQryTrade *CHSReqQryTradeField, nRequestID int) int
	ReqQryExercise(pReqQryExercise *CHSReqQryExerciseField, nRequestID int) int
	ReqQryLock(pReqQryLock *CHSReqQryLockField, nRequestID int) int
	ReqQryCombAction(pReqQryCombAction *CHSReqQryCombActionField, nRequestID int) int
	ReqQryForQuote(pReqQryForQuote *CHSReqQryForQuoteField, nRequestID int) int
	ReqQryQuote(pReqQryQuote *CHSReqQryQuoteField, nRequestID int) int
	ReqQryPositionCombineDetail(pReqQryPositionCombineDetail *CHSReqQryPositionCombineDetailField, nRequestID int) int
	ReqQryInstrument(pReqQryInstrument *CHSReqQryInstrumentField, nRequestID int) int
	ReqQryCoveredShort(pReqQryCoveredShort *CHSReqQryCoveredShortField, nRequestID int) int
	ReqQryExerciseAssign(pReqQryExerciseAssign *CHSReqQryExerciseAssignField, nRequestID int) int
	ReqTransfer(pReqTransfer *CHSReqTransferField, nRequestID int) int
	ReqQryTransfer(pReqQryTransfer *CHSReqQryTransferField, nRequestID int) int
	ReqQueryBankBalance(pReqQueryBankBalance *CHSReqQueryBankBalanceField, nRequestID int) int
	ReqQueryBankAccount(pReqQueryBankAccount *CHSReqQueryBankAccountField, nRequestID int) int
	ReqMultiCentreFundTrans(pReqMultiCentreFundTransfer *CHSReqMultiCentreFundTransField, nRequestID int) int
	ReqQueryBillContent(pReqQueryBillContent *CHSReqQueryBillContentField, nRequestID int) int
	ReqBillConfirm(pReqBillConfirm *CHSReqBillConfirmField, nRequestID int) int
	ReqQryMargin(pReqQryMargin *CHSReqQryMarginField, nRequestID int) int
	ReqQryCommission(pReqQryCommission *CHSReqQryCommissionField, nRequestID int) int
	ReqQryExchangeRate(pReqQryExchangeRate *CHSReqQryExchangeRateField, nRequestID int) int
	ReqQryPositionDetail(pReqQryPositionDetail *CHSReqQryPositionDetailField, nRequestID int) int
	ReqQrySysConfig(pReqQrySysConfig *CHSReqQrySysConfigField, nRequestID int) int
	ReqQryDepthMarketData(pReqQryDepthMarketData *CHSReqQryDepthMarketDataField, nRequestID int) int
	ReqFundTrans(pReqFundTransfer *CHSReqFundTransField, nRequestID int) int
	ReqQryFundTrans(pReqQryFundTrans *CHSReqQryFundTransField, nRequestID int) int
	ReqQryClientNotice(pReqQryClientNotice *CHSReqQryClientNoticeField, nRequestID int) int
	ReqQryOptUnderly(pReqQryOptUnderly *CHSReqQryOptUnderlyField, nRequestID int) int
	ReqQrySecuDepthMarket(pReqQrySecuDepthMarket *CHSReqQrySecuDepthMarketField, nRequestID int) int
	ReqQryHistOrder(pReqQryHistOrder *CHSReqQryHistOrderField, nRequestID int) int
	ReqQryHistTrade(pReqQryHistTrade *CHSReqQryHistTradeField, nRequestID int) int
	ReqQryWithdrawFund(pReqQryWithdrawFund *CHSReqQryWithdrawFundField, nRequestID int) int
	ReqQryCombInstrument(pReqQryCombInstrument *CHSReqQryCombInstrumentField, nRequestID int) int
	ReqQrySeatID(pReqQrySeatID *CHSReqQrySeatIDField, nRequestID int) int
	ReqOptionSelfClose(pReqOptionSelfClose *CHSReqOptionSelfCloseField, nRequestID int) int
	ReqOptionSelfCloseAction(pReqOptionSelfCloseAction *CHSReqOptionSelfCloseActionField, nRequestID int) int
	ReqQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *CHSReqQryOptionSelfCloseResultField, nRequestID int) int
	ReqQryOptionSelfClose(pReqOptionSelfClose *CHSReqOptionSelfCloseField, nRequestID int) int
	ReqOptQuoteInsert(pReqOptQuoteInsert *CHSReqOptQuoteInsertField, nRequestID int) int
	ReqOptQuoteAction(pReqOptQuoteAction *CHSReqOptQuoteActionField, nRequestID int) int
	ReqQryOptQuote(pReqQryOptQuote *CHSReqQryOptQuoteField, nRequestID int) int
	ReqQryOptCombStrategy(pReqQryOptCombStrategy *CHSReqQryOptCombStrategyField, nRequestID int) int
}

type CHSTradeSpi interface {
	OnFrontConnected()
	OnFrontDisconnected(nResult int)
	OnRspAuthenticate(pRspAuthenticate *CHSRspAuthenticateField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspSubmitUserSystemInfo(pRspUserSystemInfo *CHSRspUserSystemInfoField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspUserLogin(pRspUserLogin *CHSRspUserLoginField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspUserPasswordUpdate(pRspUserPasswordUpdate *CHSRspUserPasswordUpdateField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorOrderInsert(pRspOrderInsert *CHSRspOrderInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspOrderAction(pRspOrderAction *CHSRspOrderActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorExerciseOrderInsert(pRspExerciseOrderInsert *CHSRspExerciseOrderInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspExerciseOrderAction(pRspExerciseOrderAction *CHSRspExerciseOrderActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorLockInsert(pRspLockInsert *CHSRspLockInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspForQuoteInsert(pRspForQuoteInsert *CHSRspForQuoteInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorQuoteInsert(pRspQuoteInsert *CHSRspQuoteInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQuoteAction(pRspQuoteAction *CHSRspQuoteActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorCombActionInsert(pRspCombActionInsert *CHSRspCombActionInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQueryMaxOrderVolume(pRspQueryMaxOrderVolume *CHSRspQueryMaxOrderVolumeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryLockVolume(pRspQryLockVolume *CHSRspQryLockVolumeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQueryExerciseVolume(pRspQueryExerciseVolume *CHSRspQueryExerciseVolumeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryCombVolume(pRspQryCombVolume *CHSRspQryCombVolumeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryPosition(pRspQryPosition *CHSRspQryPositionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryTradingAccount(pRspQryTradingAccount *CHSRspQryTradingAccountField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOrder(pRspQryOrder *CHSOrderField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryTrade(pRspQryTrade *CHSTradeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryExercise(pRspQryExercise *CHSExerciseField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryLock(pRspQryLock *CHSLockField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryCombAction(pRspQryCombAction *CHSCombActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryForQuote(pRspQryForQuote *CHSForQuoteField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryQuote(pRspQryQuote *CHSQuoteField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryPositionCombineDetail(pRspQryPositionCombineDetail *CHSRspQryPositionCombineDetailField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryInstrument(pRspQryInstrument *CHSRspQryInstrumentField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryCoveredShort(pRspQryCoveredShort *CHSRspQryCoveredShortField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryExerciseAssign(pRspQryExerciseAssign *CHSRspQryExerciseAssignField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspTransfer(pRspTransfer *CHSRspTransferField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryTransfer(pRspQryTransfer *CHSRspQryTransferField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQueryBankBalance(pRspQueryBankBalance *CHSRspQueryBankBalanceField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQueryBankAccount(pRspQueryBankAccount *CHSRspQueryBankAccountField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspMultiCentreFundTrans(pRspMultiCentreFundTransfer *CHSRspMultiCentreFundTransField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQueryBillContent(pRspQueryBillContent *CHSRspQueryBillContentField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspBillConfirm(pRspBillConfirm *CHSRspBillConfirmField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryMargin(pRspQryMargin *CHSRspQryMarginField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryCommission(pRspQryCommission *CHSRspQryCommissionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryPositionDetail(pRspQryPositionDetail *CHSRspQryPositionDetailField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryExchangeRate(pRspQryExchangeRate *CHSRspQryExchangeRateField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQrySysConfig(pRspQrySysConfig *CHSRspQrySysConfigField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryDepthMarketData(pRspQryDepthMarketData *CHSDepthMarketDataField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspFundTrans(pRspFundTransfer *CHSRspFundTransField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryFundTrans(pRspQryFundTrans *CHSRspQryFundTransField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryClientNotice(pRspQryClientNotice *CHSClientNoticeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOptUnderly(pRspQryOptUnderly *CHSRspQryOptUnderlyField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQrySecuDepthMarket(pRspQrySecuDepthMarket *CHSRspQrySecuDepthMarketField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryHistOrder(pRspQryHistOrder *CHSOrderField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryHistTrade(pRspQryHistTrade *CHSTradeField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryWithdrawFund(pRspQryWithdrawFund *CHSRspQryWithdrawFundField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryCombInstrument(pRspQryCombInstrument *CHSRspQryCombInstrumentField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQrySeatID(pRspQrySeatID *CHSRspQrySeatIDField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspOptionSelfClose(pReqOptionSelfClose *CHSRspOptionSelfCloseField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspOptionSelfCloseAction(pReqOptionSelfCloseAction *CHSRspOptionSelfCloseActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *CHSRspQryOptionSelfCloseResultField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOptionSelfClose(pRspQryOptionSelfClose *CHSOptionSelfCloseField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspErrorOptQuoteInsert(pRspOptQuoteInsert *CHSRspOptQuoteInsertField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspOptQuoteAction(pRspOptQuoteAction *CHSRspOptQuoteActionField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOptQuote(pRspQryOptQuote *CHSOptQuoteField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryOptCombStrategy(pRspQryOptCombStrategy *CHSRspQryOptCombStrategyField, pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRtnTrade(pRtnTrade *CHSTradeField)
	OnRtnOrder(pRtnOrder *CHSOrderField)
	OnRtnExercise(pRtnExercise *CHSExerciseField)
	OnRtnCombAction(pRtnCombAction *CHSCombActionField)
	OnRtnLock(pRtnLock *CHSLockField)
	OnErrRtnOrderAction(pRtnOrder *CHSOrderActionField)
	OnRtnClientNotice(pRtnClientNotice *CHSClientNoticeField)
	OnRtnForQuote(pRtnForQuote *CHSForQuoteField)
	OnRtnQuote(pRtnQuote *CHSQuoteField)
	OnRtnExchangeStatus(pRtnExchangeStatus *CHSExchangeStatusField)
	OnRtnProductStatus(pRtnProductStatus *CHSProductStatusField)
	OnRtnOptionSelfClose(pRtnOptionSelfClose *CHSOptionSelfCloseField)
	OnRtnOptQuote(pRtnOptQuote *CHSOptQuoteField)
	OnRtnTransfer(pRtnTransfer *CHSTransferField)
	OnErrRtnOptQuoteAction(pRtnOptQuoteAction *CHSOptQuoteActionField)
}
