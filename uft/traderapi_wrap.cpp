#include "traderapi_wrap.h"

extern "C" {
    extern void go_td_OnFrontConnected(uintptr_t handle);
    extern void go_td_OnFrontDisconnected(uintptr_t handle, int reason);
    extern void go_td_OnRspAuthenticate(uintptr_t handle, void* pRspAuthenticateField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspSubmitUserSystemInfo(uintptr_t handle, void* pRspUserSystemInfoField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspUserLogin(uintptr_t handle, void* pRspUserLoginField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspUserPasswordUpdate(uintptr_t handle, void* pRspUserPasswordUpdateField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorOrderInsert(uintptr_t handle, void* pRspOrderInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspOrderAction(uintptr_t handle, void* pRspOrderActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorExerciseOrderInsert(uintptr_t handle, void* pRspExerciseOrderInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspExerciseOrderAction(uintptr_t handle, void* pRspExerciseOrderActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorLockInsert(uintptr_t handle, void* pRspLockInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspForQuoteInsert(uintptr_t handle, void* pRspForQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorQuoteInsert(uintptr_t handle, void* pRspQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQuoteAction(uintptr_t handle, void* pRspQuoteActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorCombActionInsert(uintptr_t handle, void* pRspCombActionInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQueryMaxOrderVolume(uintptr_t handle, void* pRspQueryMaxOrderVolumeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryLockVolume(uintptr_t handle, void* pRspQryLockVolumeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQueryExerciseVolume(uintptr_t handle, void* pRspQueryExerciseVolumeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryCombVolume(uintptr_t handle, void* pRspQryCombVolumeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryPosition(uintptr_t handle, void* pRspQryPositionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryTradingAccount(uintptr_t handle, void* pRspQryTradingAccountField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOrder(uintptr_t handle, void* pRspQryOrderField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryTrade(uintptr_t handle, void* pRspQryTradeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryExercise(uintptr_t handle, void* pRspQryExerciseField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryLock(uintptr_t handle, void* pRspQryLockField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryCombAction(uintptr_t handle, void* pRspQryCombActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryForQuote(uintptr_t handle, void* pRspQryForQuoteField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryQuote(uintptr_t handle, void* pRspQryQuoteField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryPositionCombineDetail(uintptr_t handle, void* pRspQryPositionCombineDetailField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryInstrument(uintptr_t handle, void* pRspQryInstrumentField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryCoveredShort(uintptr_t handle, void* pRspQryCoveredShortField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryExerciseAssign(uintptr_t handle, void* pRspQryExerciseAssignField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspTransfer(uintptr_t handle, void* pRspTransferField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryTransfer(uintptr_t handle, void* pRspQryTransferField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQueryBankBalance(uintptr_t handle, void* pRspQueryBankBalanceField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQueryBankAccount(uintptr_t handle, void* pRspQueryBankAccountField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspMultiCentreFundTrans(uintptr_t handle, void* pRspMultiCentreFundTransferField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQueryBillContent(uintptr_t handle, void* pRspQueryBillContentField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspBillConfirm(uintptr_t handle, void* pRspBillConfirmField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryMargin(uintptr_t handle, void* pRspQryMarginField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryCommission(uintptr_t handle, void* pRspQryCommissionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryPositionDetail(uintptr_t handle, void* pRspQryPositionDetailField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryExchangeRate(uintptr_t handle, void* pRspQryExchangeRateField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQrySysConfig(uintptr_t handle, void* pRspQrySysConfigField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryDepthMarketData(uintptr_t handle, void* pRspQryDepthMarketDataField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspFundTrans(uintptr_t handle, void* pRspFundTransferField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryFundTrans(uintptr_t handle, void* pRspQryFundTransField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryClientNotice(uintptr_t handle, void* pRspQryClientNoticeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOptUnderly(uintptr_t handle, void* pRspQryOptUnderlyField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQrySecuDepthMarket(uintptr_t handle, void* pRspQrySecuDepthMarketField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryHistOrder(uintptr_t handle, void* pRspQryHistOrderField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryHistTrade(uintptr_t handle, void* pRspQryHistTradeField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryWithdrawFund(uintptr_t handle, void* pRspQryWithdrawFundField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryCombInstrument(uintptr_t handle, void* pRspQryCombInstrumentField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQrySeatID(uintptr_t handle, void* pRspQrySeatIDField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspOptionSelfClose(uintptr_t handle, void* pRspOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspOptionSelfCloseAction(uintptr_t handle, void* pRspOptionSelfCloseActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOptionSelfCloseResult(uintptr_t handle, void* pRspQryOptionSelfCloseResultField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOptionSelfClose(uintptr_t handle, void* pRspQryOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspErrorOptQuoteInsert(uintptr_t handle, void* pRspOptQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspOptQuoteAction(uintptr_t handle, void* pRspOptQuoteActionField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOptQuote(uintptr_t handle, void* pRspQryOptQuoteField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRspQryOptCombStrategy(uintptr_t handle, void* pRspQryOptCombStrategyField, void* pRspInfo, int requestID, bool isLast);
    extern void go_td_OnRtnTrade(uintptr_t handle, void* pRtnTradeField);
    extern void go_td_OnRtnOrder(uintptr_t handle, void* pRtnOrderField);
    extern void go_td_OnRtnExercise(uintptr_t handle, void* pRtnExerciseField);
    extern void go_td_OnRtnCombAction(uintptr_t handle, void* pRtnCombActionField);
    extern void go_td_OnRtnLock(uintptr_t handle, void* pRtnLockField);
    extern void go_td_OnErrRtnOrderAction(uintptr_t handle, void* pRtnOrderActionField);
    extern void go_td_OnRtnClientNotice(uintptr_t handle, void* pRtnClientNoticeField);
    extern void go_td_OnRtnForQuote(uintptr_t handle, void* pRtnForQuoteField);
    extern void go_td_OnRtnQuote(uintptr_t handle, void* pRtnQuoteField);
    extern void go_td_OnRtnExchangeStatus(uintptr_t handle, void* pRtnExchangeStatusField);
    extern void go_td_OnRtnProductStatus(uintptr_t handle, void* pRtnProductStatusField);
    extern void go_td_OnRtnOptionSelfClose(uintptr_t handle, void* pRtnOptionSelfCloseField);
    extern void go_td_OnRtnOptQuote(uintptr_t handle, void* pRtnOptQuoteField);
    extern void go_td_OnRtnTransfer(uintptr_t handle, void* pRtnTransferField);
    extern void go_td_OnErrRtnOptQuoteAction(uintptr_t handle, void* pRtnOptQuoteActionField);
}

void HSTradeSpiWrap::OnFrontConnected() {
    go_td_OnFrontConnected(m_handle);
}

void HSTradeSpiWrap::OnFrontDisconnected(int nResult) {
    go_td_OnFrontDisconnected(m_handle, nResult);
}

void HSTradeSpiWrap::OnRspAuthenticate(CHSRspAuthenticateField *pRspAuthenticate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspAuthenticate(m_handle, pRspAuthenticate, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspSubmitUserSystemInfo(CHSRspUserSystemInfoField *pRspUserSystemInfo, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspSubmitUserSystemInfo(m_handle, pRspUserSystemInfo, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspUserLogin(CHSRspUserLoginField *pRspUserLogin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspUserLogin(m_handle, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspUserPasswordUpdate(CHSRspUserPasswordUpdateField *pRspUserPasswordUpdate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspUserPasswordUpdate(m_handle, pRspUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorOrderInsert(CHSRspOrderInsertField *pRspOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorOrderInsert(m_handle, pRspOrderInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspOrderAction(CHSRspOrderActionField *pRspOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspOrderAction(m_handle, pRspOrderAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorExerciseOrderInsert(CHSRspExerciseOrderInsertField *pRspExerciseOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorExerciseOrderInsert(m_handle, pRspExerciseOrderInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspExerciseOrderAction(CHSRspExerciseOrderActionField *pRspExerciseOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspExerciseOrderAction(m_handle, pRspExerciseOrderAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorLockInsert(CHSRspLockInsertField *pRspLockInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorLockInsert(m_handle, pRspLockInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspForQuoteInsert(CHSRspForQuoteInsertField *pRspForQuoteInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspForQuoteInsert(m_handle, pRspForQuoteInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorQuoteInsert(CHSRspQuoteInsertField* pRspQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorQuoteInsert(m_handle, pRspQuoteInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQuoteAction(CHSRspQuoteActionField* pRspQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQuoteAction(m_handle, pRspQuoteAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorCombActionInsert(CHSRspCombActionInsertField *pRspCombActionInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorCombActionInsert(m_handle, pRspCombActionInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQueryMaxOrderVolume(CHSRspQueryMaxOrderVolumeField *pRspQueryMaxOrderVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQueryMaxOrderVolume(m_handle, pRspQueryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryLockVolume(CHSRspQryLockVolumeField *pRspQryLockVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryLockVolume(m_handle, pRspQryLockVolume, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQueryExerciseVolume(CHSRspQueryExerciseVolumeField *pRspQueryExerciseVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQueryExerciseVolume(m_handle, pRspQueryExerciseVolume, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryCombVolume(CHSRspQryCombVolumeField *pRspQryCombVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryCombVolume(m_handle, pRspQryCombVolume, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryPosition(CHSRspQryPositionField *pRspQryPosition, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryPosition(m_handle, pRspQryPosition, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryTradingAccount(CHSRspQryTradingAccountField *pRspQryTradingAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryTradingAccount(m_handle, pRspQryTradingAccount, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOrder(CHSOrderField *pRspQryOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOrder(m_handle, pRspQryOrder, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryTrade(CHSTradeField *pRspQryTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryTrade(m_handle, pRspQryTrade, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryExercise(CHSExerciseField *pRspQryExercise, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryExercise(m_handle, pRspQryExercise, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryLock(CHSLockField *pRspQryLock, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryLock(m_handle, pRspQryLock, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryCombAction(CHSCombActionField *pRspQryCombAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryCombAction(m_handle, pRspQryCombAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryForQuote(CHSForQuoteField* pRspQryForQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryForQuote(m_handle, pRspQryForQuote, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryQuote(CHSQuoteField* pRspQryQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryQuote(m_handle, pRspQryQuote, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryPositionCombineDetail(CHSRspQryPositionCombineDetailField *pRspQryPositionCombineDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryPositionCombineDetail(m_handle, pRspQryPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryInstrument(CHSRspQryInstrumentField *pRspQryInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryInstrument(m_handle, pRspQryInstrument, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryCoveredShort(CHSRspQryCoveredShortField *pRspQryCoveredShort, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryCoveredShort(m_handle, pRspQryCoveredShort, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryExerciseAssign(CHSRspQryExerciseAssignField *pRspQryExerciseAssign, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryExerciseAssign(m_handle, pRspQryExerciseAssign, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspTransfer(CHSRspTransferField *pRspTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspTransfer(m_handle, pRspTransfer, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryTransfer(CHSRspQryTransferField *pRspQryTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryTransfer(m_handle, pRspQryTransfer, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQueryBankBalance(CHSRspQueryBankBalanceField *pRspQueryBankBalance, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQueryBankBalance(m_handle, pRspQueryBankBalance, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQueryBankAccount(CHSRspQueryBankAccountField *pRspQueryBankAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQueryBankAccount(m_handle, pRspQueryBankAccount, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspMultiCentreFundTrans(CHSRspMultiCentreFundTransField *pRspMultiCentreFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspMultiCentreFundTrans(m_handle, pRspMultiCentreFundTransfer, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQueryBillContent(CHSRspQueryBillContentField *pRspQueryBillContent, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQueryBillContent(m_handle, pRspQueryBillContent, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspBillConfirm(CHSRspBillConfirmField *pRspBillConfirm, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspBillConfirm(m_handle, pRspBillConfirm, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryMargin(CHSRspQryMarginField *pRspQryMargin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryMargin(m_handle, pRspQryMargin, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryCommission(CHSRspQryCommissionField *pRspQryCommission, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryCommission(m_handle, pRspQryCommission, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryPositionDetail(CHSRspQryPositionDetailField *pRspQryPositionDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryPositionDetail(m_handle, pRspQryPositionDetail, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryExchangeRate(CHSRspQryExchangeRateField *pRspQryExchangeRate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryExchangeRate(m_handle, pRspQryExchangeRate, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQrySysConfig(CHSRspQrySysConfigField *pRspQrySysConfig, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQrySysConfig(m_handle, pRspQrySysConfig, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryDepthMarketData(CHSDepthMarketDataField *pRspQryDepthMarketData, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryDepthMarketData(m_handle, pRspQryDepthMarketData, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspFundTrans(CHSRspFundTransField *pRspFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspFundTrans(m_handle, pRspFundTransfer, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryFundTrans(CHSRspQryFundTransField *pRspQryFundTrans, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryFundTrans(m_handle, pRspQryFundTrans, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryClientNotice(CHSClientNoticeField *pRspQryClientNotice, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryClientNotice(m_handle, pRspQryClientNotice, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOptUnderly(CHSRspQryOptUnderlyField *pRspQryOptUnderly, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOptUnderly(m_handle, pRspQryOptUnderly, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQrySecuDepthMarket(CHSRspQrySecuDepthMarketField *pRspQrySecuDepthMarket, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQrySecuDepthMarket(m_handle, pRspQrySecuDepthMarket, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryHistOrder(CHSOrderField *pRspQryHistOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryHistOrder(m_handle, pRspQryHistOrder, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryHistTrade(CHSTradeField *pRspQryHistTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryHistTrade(m_handle, pRspQryHistTrade, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryWithdrawFund(CHSRspQryWithdrawFundField *pRspQryWithdrawFund, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryWithdrawFund(m_handle, pRspQryWithdrawFund, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryCombInstrument(CHSRspQryCombInstrumentField *pRspQryCombInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryCombInstrument(m_handle, pRspQryCombInstrument, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQrySeatID(CHSRspQrySeatIDField *pRspQrySeatID, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQrySeatID(m_handle, pRspQrySeatID, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspOptionSelfClose(CHSRspOptionSelfCloseField *pReqOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspOptionSelfClose(m_handle, pReqOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspOptionSelfCloseAction(CHSRspOptionSelfCloseActionField *pReqOptionSelfCloseAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspOptionSelfCloseAction(m_handle, pReqOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOptionSelfCloseResult(CHSRspQryOptionSelfCloseResultField *pReqQryOptionSelfCloseResult, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOptionSelfCloseResult(m_handle, pReqQryOptionSelfCloseResult, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOptionSelfClose(CHSOptionSelfCloseField *pRspQryOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOptionSelfClose(m_handle, pRspQryOptionSelfClose, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspErrorOptQuoteInsert(CHSRspOptQuoteInsertField *pRspOptQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspErrorOptQuoteInsert(m_handle, pRspOptQuoteInsert, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspOptQuoteAction(CHSRspOptQuoteActionField *pRspOptQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspOptQuoteAction(m_handle, pRspOptQuoteAction, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOptQuote(CHSOptQuoteField *pRspQryOptQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOptQuote(m_handle, pRspQryOptQuote, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRspQryOptCombStrategy(CHSRspQryOptCombStrategyField *pRspQryOptCombStrategy, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_td_OnRspQryOptCombStrategy(m_handle, pRspQryOptCombStrategy, pRspInfo, nRequestID, bIsLast);
}

void HSTradeSpiWrap::OnRtnTrade(CHSTradeField *pRtnTrade) {
    go_td_OnRtnTrade(m_handle, pRtnTrade);
}

void HSTradeSpiWrap::OnRtnOrder(CHSOrderField *pRtnOrder) {
    go_td_OnRtnOrder(m_handle, pRtnOrder);
}

void HSTradeSpiWrap::OnRtnExercise(CHSExerciseField *pRtnExercise) {
    go_td_OnRtnExercise(m_handle, pRtnExercise);
}

void HSTradeSpiWrap::OnRtnCombAction(CHSCombActionField *pRtnCombAction) {
    go_td_OnRtnCombAction(m_handle, pRtnCombAction);
}

void HSTradeSpiWrap::OnRtnLock(CHSLockField *pRtnLock) {
    go_td_OnRtnLock(m_handle, pRtnLock);
}

void HSTradeSpiWrap::OnErrRtnOrderAction(CHSOrderActionField *pRtnOrderAction) {
    go_td_OnErrRtnOrderAction(m_handle, pRtnOrderAction);
}

void HSTradeSpiWrap::OnRtnClientNotice(CHSClientNoticeField *pRtnClientNotice) {
    go_td_OnRtnClientNotice(m_handle, pRtnClientNotice);
}

void HSTradeSpiWrap::OnRtnForQuote(CHSForQuoteField* pRtnForQuote) {
    go_td_OnRtnForQuote(m_handle, pRtnForQuote);
}

void HSTradeSpiWrap::OnRtnQuote(CHSQuoteField* pRtnQuote) {
    go_td_OnRtnQuote(m_handle, pRtnQuote);
}

void HSTradeSpiWrap::OnRtnExchangeStatus(CHSExchangeStatusField* pRtnExchangeStatus) {
    go_td_OnRtnExchangeStatus(m_handle, pRtnExchangeStatus);
}

void HSTradeSpiWrap::OnRtnProductStatus(CHSProductStatusField* pRtnProductStatus) {
    go_td_OnRtnProductStatus(m_handle, pRtnProductStatus);
}

void HSTradeSpiWrap::OnRtnOptionSelfClose(CHSOptionSelfCloseField *pRtnOptionSelfClose) {
    go_td_OnRtnOptionSelfClose(m_handle, pRtnOptionSelfClose);
}

void HSTradeSpiWrap::OnRtnOptQuote(CHSOptQuoteField* pRtnOptQuote) {
    go_td_OnRtnOptQuote(m_handle, pRtnOptQuote);
}

void HSTradeSpiWrap::OnRtnTransfer(CHSTransferField* pRtnTransfer) {
    go_td_OnRtnTransfer(m_handle, pRtnTransfer);
}

void HSTradeSpiWrap::OnErrRtnOptQuoteAction(CHSOptQuoteActionField* pRtnOptQuoteAction) {
    go_td_OnErrRtnOptQuoteAction(m_handle, pRtnOptQuoteAction);
}

extern "C" {

uintptr_t hs_create_td_api_static(const char* flowPath) {
    CHSTradeApi* api = NewTradeApi(flowPath);
    if (!api) return 0;
    HSTradeApiWrap* wrap = new HSTradeApiWrap(api);
    return reinterpret_cast<uintptr_t>(wrap);
}

void hs_release_td_api_static(uintptr_t api) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (wrap) {
        wrap->GetApi()->ReleaseApi();
        delete wrap;
    }
}

const char* hs_td_GetApiVersion_static(uintptr_t api) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return "";
    return GetTradeApiVersion();
}

int hs_td_Init_static(uintptr_t api, void* config) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->Init(reinterpret_cast<CHSInitConfigField*>(config));
}

int hs_td_Join_static(uintptr_t api) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->Join();
}

int hs_td_RegisterSubModel_static(uintptr_t api, int eSubType) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->RegisterSubModel(static_cast<SUB_TERT_TYPE>(eSubType));
}

int hs_td_RegisterFront_static(uintptr_t api, const char* frontAddress) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->RegisterFront(frontAddress);
}

int hs_td_RegisterFensServer_static(uintptr_t api, const char* fensAddress, const char* accountID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->RegisterFensServer(fensAddress, accountID);
}

void hs_td_RegisterSpi_static(uintptr_t api, uintptr_t handle) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return;
    HSTradeSpiWrap* spi = new HSTradeSpiWrap(handle);
    wrap->SetSpi(spi);
    wrap->GetApi()->RegisterSpi(spi);
}

const char* hs_td_GetApiErrorMsg_static(int errorCode) {
    return nullptr;
}

int hs_td_GetTradingDate_static(uintptr_t api) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->GetTradingDate();
}

int hs_td_BindSessionID_static(uintptr_t api, uint8_t nSessionID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->BindSessionID(nSessionID);
}

int hs_td_ReqAuthenticate_static(uintptr_t api, void* pReqAuthenticateField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqAuthenticate(reinterpret_cast<CHSReqAuthenticateField*>(pReqAuthenticateField), nRequestID);
}

int hs_td_ReqSubmitUserSystemInfo_static(uintptr_t api, void* pReqUserSystemInfoField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqSubmitUserSystemInfo(reinterpret_cast<CHSReqUserSystemInfoField*>(pReqUserSystemInfoField), nRequestID);
}

int hs_td_ReqUserLogin_static(uintptr_t api, void* pReqUserLoginField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqUserLogin(reinterpret_cast<CHSReqUserLoginField*>(pReqUserLoginField), nRequestID);
}

int hs_td_ReqUserPasswordUpdate_static(uintptr_t api, void* pReqUserPasswordUpdateField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqUserPasswordUpdate(reinterpret_cast<CHSReqUserPasswordUpdateField*>(pReqUserPasswordUpdateField), nRequestID);
}

int hs_td_ReqOrderInsert_static(uintptr_t api, void* pReqOrderInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOrderInsert(reinterpret_cast<CHSReqOrderInsertField*>(pReqOrderInsertField), nRequestID);
}

int hs_td_ReqOrderAction_static(uintptr_t api, void* pReqOrderActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOrderAction(reinterpret_cast<CHSReqOrderActionField*>(pReqOrderActionField), nRequestID);
}

int hs_td_ReqExerciseOrderInsert_static(uintptr_t api, void* pReqExerciseOrderInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqExerciseOrderInsert(reinterpret_cast<CHSReqExerciseOrderInsertField*>(pReqExerciseOrderInsertField), nRequestID);
}

int hs_td_ReqExerciseOrderAction_static(uintptr_t api, void* pReqExerciseOrderActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqExerciseOrderAction(reinterpret_cast<CHSReqExerciseOrderActionField*>(pReqExerciseOrderActionField), nRequestID);
}

int hs_td_ReqLockInsert_static(uintptr_t api, void* pReqLockInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqLockInsert(reinterpret_cast<CHSReqLockInsertField*>(pReqLockInsertField), nRequestID);
}

int hs_td_ReqForQuoteInsert_static(uintptr_t api, void* pReqForQuoteInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqForQuoteInsert(reinterpret_cast<CHSReqForQuoteInsertField*>(pReqForQuoteInsertField), nRequestID);
}

int hs_td_ReqQuoteInsert_static(uintptr_t api, void* pReqQuoteInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQuoteInsert(reinterpret_cast<CHSReqQuoteInsertField*>(pReqQuoteInsertField), nRequestID);
}

int hs_td_ReqQuoteAction_static(uintptr_t api, void* pReqQuoteActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQuoteAction(reinterpret_cast<CHSReqQuoteActionField*>(pReqQuoteActionField), nRequestID);
}

int hs_td_ReqCombActionInsert_static(uintptr_t api, void* pReqCombActionInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqCombActionInsert(reinterpret_cast<CHSReqCombActionInsertField*>(pReqCombActionInsertField), nRequestID);
}

int hs_td_ReqQueryMaxOrderVolume_static(uintptr_t api, void* pReqQueryMaxOrderVolumeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQueryMaxOrderVolume(reinterpret_cast<CHSReqQueryMaxOrderVolumeField*>(pReqQueryMaxOrderVolumeField), nRequestID);
}

int hs_td_ReqQryLockVolume_static(uintptr_t api, void* pReqQryLockVolumeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryLockVolume(reinterpret_cast<CHSReqQryLockVolumeField*>(pReqQryLockVolumeField), nRequestID);
}

int hs_td_ReqQueryExerciseVolume_static(uintptr_t api, void* pReqQueryExerciseVolumeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQueryExerciseVolume(reinterpret_cast<CHSReqQueryExerciseVolumeField*>(pReqQueryExerciseVolumeField), nRequestID);
}

int hs_td_ReqQryCombVolume_static(uintptr_t api, void* pReqQryCombVolumeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryCombVolume(reinterpret_cast<CHSReqQryCombVolumeField*>(pReqQryCombVolumeField), nRequestID);
}

int hs_td_ReqQryPosition_static(uintptr_t api, void* pReqQryPositionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryPosition(reinterpret_cast<CHSReqQryPositionField*>(pReqQryPositionField), nRequestID);
}

int hs_td_ReqQryTradingAccount_static(uintptr_t api, void* pReqQryTradingAccountField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryTradingAccount(reinterpret_cast<CHSReqQryTradingAccountField*>(pReqQryTradingAccountField), nRequestID);
}

int hs_td_ReqQryOrder_static(uintptr_t api, void* pReqQryOrderField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOrder(reinterpret_cast<CHSReqQryOrderField*>(pReqQryOrderField), nRequestID);
}

int hs_td_ReqQryTrade_static(uintptr_t api, void* pReqQryTradeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryTrade(reinterpret_cast<CHSReqQryTradeField*>(pReqQryTradeField), nRequestID);
}

int hs_td_ReqQryExercise_static(uintptr_t api, void* pReqQryExerciseField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryExercise(reinterpret_cast<CHSReqQryExerciseField*>(pReqQryExerciseField), nRequestID);
}

int hs_td_ReqQryLock_static(uintptr_t api, void* pReqQryLockField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryLock(reinterpret_cast<CHSReqQryLockField*>(pReqQryLockField), nRequestID);
}

int hs_td_ReqQryCombAction_static(uintptr_t api, void* pReqQryCombActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryCombAction(reinterpret_cast<CHSReqQryCombActionField*>(pReqQryCombActionField), nRequestID);
}

int hs_td_ReqQryForQuote_static(uintptr_t api, void* pReqQryForQuoteField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryForQuote(reinterpret_cast<CHSReqQryForQuoteField*>(pReqQryForQuoteField), nRequestID);
}

int hs_td_ReqQryQuote_static(uintptr_t api, void* pReqQryQuoteField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryQuote(reinterpret_cast<CHSReqQryQuoteField*>(pReqQryQuoteField), nRequestID);
}

int hs_td_ReqQryPositionCombineDetail_static(uintptr_t api, void* pReqQryPositionCombineDetailField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryPositionCombineDetail(reinterpret_cast<CHSReqQryPositionCombineDetailField*>(pReqQryPositionCombineDetailField), nRequestID);
}

int hs_td_ReqQryInstrument_static(uintptr_t api, void* pReqQryInstrumentField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryInstrument(reinterpret_cast<CHSReqQryInstrumentField*>(pReqQryInstrumentField), nRequestID);
}

int hs_td_ReqQryCoveredShort_static(uintptr_t api, void* pReqQryCoveredShortField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryCoveredShort(reinterpret_cast<CHSReqQryCoveredShortField*>(pReqQryCoveredShortField), nRequestID);
}

int hs_td_ReqQryExerciseAssign_static(uintptr_t api, void* pReqQryExerciseAssignField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryExerciseAssign(reinterpret_cast<CHSReqQryExerciseAssignField*>(pReqQryExerciseAssignField), nRequestID);
}

int hs_td_ReqTransfer_static(uintptr_t api, void* pReqTransferField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqTransfer(reinterpret_cast<CHSReqTransferField*>(pReqTransferField), nRequestID);
}

int hs_td_ReqQryTransfer_static(uintptr_t api, void* pReqQryTransferField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryTransfer(reinterpret_cast<CHSReqQryTransferField*>(pReqQryTransferField), nRequestID);
}

int hs_td_ReqQueryBankBalance_static(uintptr_t api, void* pReqQueryBankBalanceField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQueryBankBalance(reinterpret_cast<CHSReqQueryBankBalanceField*>(pReqQueryBankBalanceField), nRequestID);
}

int hs_td_ReqQueryBankAccount_static(uintptr_t api, void* pReqQueryBankAccountField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQueryBankAccount(reinterpret_cast<CHSReqQueryBankAccountField*>(pReqQueryBankAccountField), nRequestID);
}

int hs_td_ReqMultiCentreFundTrans_static(uintptr_t api, void* pReqMultiCentreFundTransferField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqMultiCentreFundTrans(reinterpret_cast<CHSReqMultiCentreFundTransField*>(pReqMultiCentreFundTransferField), nRequestID);
}

int hs_td_ReqQueryBillContent_static(uintptr_t api, void* pReqQueryBillContentField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQueryBillContent(reinterpret_cast<CHSReqQueryBillContentField*>(pReqQueryBillContentField), nRequestID);
}

int hs_td_ReqBillConfirm_static(uintptr_t api, void* pReqBillConfirmField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqBillConfirm(reinterpret_cast<CHSReqBillConfirmField*>(pReqBillConfirmField), nRequestID);
}

int hs_td_ReqQryMargin_static(uintptr_t api, void* pReqQryMarginField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryMargin(reinterpret_cast<CHSReqQryMarginField*>(pReqQryMarginField), nRequestID);
}

int hs_td_ReqQryCommission_static(uintptr_t api, void* pReqQryCommissionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryCommission(reinterpret_cast<CHSReqQryCommissionField*>(pReqQryCommissionField), nRequestID);
}

int hs_td_ReqQryExchangeRate_static(uintptr_t api, void* pReqQryExchangeRateField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryExchangeRate(reinterpret_cast<CHSReqQryExchangeRateField*>(pReqQryExchangeRateField), nRequestID);
}

int hs_td_ReqQryPositionDetail_static(uintptr_t api, void* pReqQryPositionDetailField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryPositionDetail(reinterpret_cast<CHSReqQryPositionDetailField*>(pReqQryPositionDetailField), nRequestID);
}

int hs_td_ReqQrySysConfig_static(uintptr_t api, void* pReqQrySysConfigField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQrySysConfig(reinterpret_cast<CHSReqQrySysConfigField*>(pReqQrySysConfigField), nRequestID);
}

int hs_td_ReqQryDepthMarketData_static(uintptr_t api, void* pReqQryDepthMarketDataField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryDepthMarketData(reinterpret_cast<CHSReqQryDepthMarketDataField*>(pReqQryDepthMarketDataField), nRequestID);
}

int hs_td_ReqFundTrans_static(uintptr_t api, void* pReqFundTransferField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqFundTrans(reinterpret_cast<CHSReqFundTransField*>(pReqFundTransferField), nRequestID);
}

int hs_td_ReqQryFundTrans_static(uintptr_t api, void* pReqQryFundTransField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryFundTrans(reinterpret_cast<CHSReqQryFundTransField*>(pReqQryFundTransField), nRequestID);
}

int hs_td_ReqQryClientNotice_static(uintptr_t api, void* pReqQryClientNoticeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryClientNotice(reinterpret_cast<CHSReqQryClientNoticeField*>(pReqQryClientNoticeField), nRequestID);
}

int hs_td_ReqQryOptUnderly_static(uintptr_t api, void* pReqQryOptUnderlyField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOptUnderly(reinterpret_cast<CHSReqQryOptUnderlyField*>(pReqQryOptUnderlyField), nRequestID);
}

int hs_td_ReqQrySecuDepthMarket_static(uintptr_t api, void* pReqQrySecuDepthMarketField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQrySecuDepthMarket(reinterpret_cast<CHSReqQrySecuDepthMarketField*>(pReqQrySecuDepthMarketField), nRequestID);
}

int hs_td_ReqQryHistOrder_static(uintptr_t api, void* pReqQryHistOrderField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryHistOrder(reinterpret_cast<CHSReqQryHistOrderField*>(pReqQryHistOrderField), nRequestID);
}

int hs_td_ReqQryHistTrade_static(uintptr_t api, void* pReqQryHistTradeField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryHistTrade(reinterpret_cast<CHSReqQryHistTradeField*>(pReqQryHistTradeField), nRequestID);
}

int hs_td_ReqQryWithdrawFund_static(uintptr_t api, void* pReqQryWithdrawFundField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryWithdrawFund(reinterpret_cast<CHSReqQryWithdrawFundField*>(pReqQryWithdrawFundField), nRequestID);
}

int hs_td_ReqQryCombInstrument_static(uintptr_t api, void* pReqQryCombInstrumentField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryCombInstrument(reinterpret_cast<CHSReqQryCombInstrumentField*>(pReqQryCombInstrumentField), nRequestID);
}

int hs_td_ReqQrySeatID_static(uintptr_t api, void* pReqQrySeatIDField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQrySeatID(reinterpret_cast<CHSReqQrySeatIDField*>(pReqQrySeatIDField), nRequestID);
}

int hs_td_ReqOptionSelfClose_static(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOptionSelfClose(reinterpret_cast<CHSReqOptionSelfCloseField*>(pReqOptionSelfCloseField), nRequestID);
}

int hs_td_ReqOptionSelfCloseAction_static(uintptr_t api, void* pReqOptionSelfCloseActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOptionSelfCloseAction(reinterpret_cast<CHSReqOptionSelfCloseActionField*>(pReqOptionSelfCloseActionField), nRequestID);
}

int hs_td_ReqQryOptionSelfCloseResult_static(uintptr_t api, void* pReqQryOptionSelfCloseResultField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOptionSelfCloseResult(reinterpret_cast<CHSReqQryOptionSelfCloseResultField*>(pReqQryOptionSelfCloseResultField), nRequestID);
}

int hs_td_ReqQryOptionSelfClose_static(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOptionSelfClose(reinterpret_cast<CHSReqQryOptionSelfCloseField*>(pReqOptionSelfCloseField), nRequestID);
}

int hs_td_ReqOptQuoteInsert_static(uintptr_t api, void* pReqOptQuoteInsertField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOptQuoteInsert(reinterpret_cast<CHSReqOptQuoteInsertField*>(pReqOptQuoteInsertField), nRequestID);
}

int hs_td_ReqOptQuoteAction_static(uintptr_t api, void* pReqOptQuoteActionField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqOptQuoteAction(reinterpret_cast<CHSReqOptQuoteActionField*>(pReqOptQuoteActionField), nRequestID);
}

int hs_td_ReqQryOptQuote_static(uintptr_t api, void* pReqQryOptQuoteField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOptQuote(reinterpret_cast<CHSReqQryOptQuoteField*>(pReqQryOptQuoteField), nRequestID);
}

int hs_td_ReqQryOptCombStrategy_static(uintptr_t api, void* pReqQryOptCombStrategyField, int nRequestID) {
    HSTradeApiWrap* wrap = reinterpret_cast<HSTradeApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqQryOptCombStrategy(reinterpret_cast<CHSReqQryOptCombStrategyField*>(pReqQryOptCombStrategyField), nRequestID);
}

}
