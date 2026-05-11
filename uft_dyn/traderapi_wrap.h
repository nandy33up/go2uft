#ifndef _HS_TRADERAPI_WRAP_H_
#define _HS_TRADERAPI_WRAP_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void go_td_OnFrontConnected(uintptr_t handle, int _dummy);
void go_td_OnFrontDisconnected(uintptr_t handle, int reason);
void go_td_OnRspAuthenticate(uintptr_t handle, void* pRspAuthenticateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspSubmitUserSystemInfo(uintptr_t handle, void* pRspUserSystemInfoField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspUserLogin(uintptr_t handle, void* pRspUserLoginField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspUserPasswordUpdate(uintptr_t handle, void* pRspUserPasswordUpdateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorOrderInsert(uintptr_t handle, void* pRspOrderInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOrderAction(uintptr_t handle, void* pRspOrderActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorExerciseOrderInsert(uintptr_t handle, void* pRspExerciseOrderInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspExerciseOrderAction(uintptr_t handle, void* pRspExerciseOrderActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorLockInsert(uintptr_t handle, void* pRspLockInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspForQuoteInsert(uintptr_t handle, void* pRspForQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorQuoteInsert(uintptr_t handle, void* pRspQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQuoteAction(uintptr_t handle, void* pRspQuoteActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorCombActionInsert(uintptr_t handle, void* pRspCombActionInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryMaxOrderVolume(uintptr_t handle, void* pRspQueryMaxOrderVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryLockVolume(uintptr_t handle, void* pRspQryLockVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryExerciseVolume(uintptr_t handle, void* pRspQueryExerciseVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombVolume(uintptr_t handle, void* pRspQryCombVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPosition(uintptr_t handle, void* pRspQryPositionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTradingAccount(uintptr_t handle, void* pRspQryTradingAccountField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOrder(uintptr_t handle, void* pRspQryOrderField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTrade(uintptr_t handle, void* pRspQryTradeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExercise(uintptr_t handle, void* pRspQryExerciseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryLock(uintptr_t handle, void* pRspQryLockField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombAction(uintptr_t handle, void* pRspQryCombActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryForQuote(uintptr_t handle, void* pRspQryForQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryQuote(uintptr_t handle, void* pRspQryQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPositionCombineDetail(uintptr_t handle, void* pRspQryPositionCombineDetailField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryInstrument(uintptr_t handle, void* pRspQryInstrumentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCoveredShort(uintptr_t handle, void* pRspQryCoveredShortField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExerciseAssign(uintptr_t handle, void* pRspQryExerciseAssignField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspTransfer(uintptr_t handle, void* pRspTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTransfer(uintptr_t handle, void* pRspQryTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBankBalance(uintptr_t handle, void* pRspQueryBankBalanceField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBankAccount(uintptr_t handle, void* pRspQueryBankAccountField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspMultiCentreFundTrans(uintptr_t handle, void* pRspMultiCentreFundTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBillContent(uintptr_t handle, void* pRspQueryBillContentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspBillConfirm(uintptr_t handle, void* pRspBillConfirmField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryMargin(uintptr_t handle, void* pRspQryMarginField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCommission(uintptr_t handle, void* pRspQryCommissionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPositionDetail(uintptr_t handle, void* pRspQryPositionDetailField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExchangeRate(uintptr_t handle, void* pRspQryExchangeRateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySysConfig(uintptr_t handle, void* pRspQrySysConfigField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryDepthMarketData(uintptr_t handle, void* pRspQryDepthMarketDataField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspFundTrans(uintptr_t handle, void* pRspFundTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryFundTrans(uintptr_t handle, void* pRspQryFundTransField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryClientNotice(uintptr_t handle, void* pRspQryClientNoticeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptUnderly(uintptr_t handle, void* pRspQryOptUnderlyField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySecuDepthMarket(uintptr_t handle, void* pRspQrySecuDepthMarketField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryHistOrder(uintptr_t handle, void* pRspQryHistOrderField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryHistTrade(uintptr_t handle, void* pRspQryHistTradeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryWithdrawFund(uintptr_t handle, void* pRspQryWithdrawFundField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombInstrument(uintptr_t handle, void* pRspQryCombInstrumentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySeatID(uintptr_t handle, void* pRspQrySeatIDField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptionSelfClose(uintptr_t handle, void* pRspOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptionSelfCloseAction(uintptr_t handle, void* pRspOptionSelfCloseActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptionSelfCloseResult(uintptr_t handle, void* pRspQryOptionSelfCloseResultField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptionSelfClose(uintptr_t handle, void* pRspQryOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorOptQuoteInsert(uintptr_t handle, void* pRspOptQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptQuoteAction(uintptr_t handle, void* pRspOptQuoteActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptQuote(uintptr_t handle, void* pRspQryOptQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptCombStrategy(uintptr_t handle, void* pRspQryOptCombStrategyField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRtnTrade(uintptr_t handle, void* pRtnTradeField);
void go_td_OnRtnOrder(uintptr_t handle, void* pRtnOrderField);
void go_td_OnRtnExercise(uintptr_t handle, void* pRtnExerciseField);
void go_td_OnRtnCombAction(uintptr_t handle, void* pRtnCombActionField);
void go_td_OnRtnLock(uintptr_t handle, void* pRtnLockField);
void go_td_OnErrRtnOrderAction(uintptr_t handle, void* pRtnOrderActionField);
void go_td_OnRtnClientNotice(uintptr_t handle, void* pRtnClientNoticeField);
void go_td_OnRtnForQuote(uintptr_t handle, void* pRtnForQuoteField);
void go_td_OnRtnQuote(uintptr_t handle, void* pRtnQuoteField);
void go_td_OnRtnExchangeStatus(uintptr_t handle, void* pRtnExchangeStatusField);
void go_td_OnRtnProductStatus(uintptr_t handle, void* pRtnProductStatusField);
void go_td_OnRtnOptionSelfClose(uintptr_t handle, void* pRtnOptionSelfCloseField);
void go_td_OnRtnOptQuote(uintptr_t handle, void* pRtnOptQuoteField);
void go_td_OnRtnTransfer(uintptr_t handle, void* pRtnTransferField);
void go_td_OnErrRtnOptQuoteAction(uintptr_t handle, void* pRtnOptQuoteActionField);

uintptr_t hs_create_td_api(const char* dllPath);
void hs_release_td_api(uintptr_t api);
const char* hs_td_GetApiVersion(uintptr_t api);
int hs_td_Init(uintptr_t api, void* config);
int hs_td_Join(uintptr_t api);
int hs_td_RegisterSubModel(uintptr_t api, int eSubType);
int hs_td_RegisterFront(uintptr_t api, const char* frontAddress);
int hs_td_RegisterFensServer(uintptr_t api, const char* fensAddress, const char* accountID);
void hs_td_RegisterSpi(uintptr_t api, uintptr_t spi);
const char* hs_td_GetApiErrorMsg(int errorCode);
int hs_td_GetTradingDate(uintptr_t api);
int hs_td_BindSessionID(uintptr_t api, uint8_t nSessionID);
int hs_td_ReqAuthenticate(uintptr_t api, void* pReqAuthenticateField, int nRequestID);
int hs_td_ReqSubmitUserSystemInfo(uintptr_t api, void* pReqUserSystemInfoField, int nRequestID);
int hs_td_ReqUserLogin(uintptr_t api, void* pReqUserLoginField, int nRequestID);
int hs_td_ReqUserPasswordUpdate(uintptr_t api, void* pReqUserPasswordUpdateField, int nRequestID);
int hs_td_ReqOrderInsert(uintptr_t api, void* pReqOrderInsertField, int nRequestID);
int hs_td_ReqOrderAction(uintptr_t api, void* pReqOrderActionField, int nRequestID);
int hs_td_ReqExerciseOrderInsert(uintptr_t api, void* pReqExerciseOrderInsertField, int nRequestID);
int hs_td_ReqExerciseOrderAction(uintptr_t api, void* pReqExerciseOrderActionField, int nRequestID);
int hs_td_ReqLockInsert(uintptr_t api, void* pReqLockInsertField, int nRequestID);
int hs_td_ReqForQuoteInsert(uintptr_t api, void* pReqForQuoteInsertField, int nRequestID);
int hs_td_ReqQuoteInsert(uintptr_t api, void* pReqQuoteInsertField, int nRequestID);
int hs_td_ReqQuoteAction(uintptr_t api, void* pReqQuoteActionField, int nRequestID);
int hs_td_ReqCombActionInsert(uintptr_t api, void* pReqCombActionInsertField, int nRequestID);
int hs_td_ReqQueryMaxOrderVolume(uintptr_t api, void* pReqQueryMaxOrderVolumeField, int nRequestID);
int hs_td_ReqQryLockVolume(uintptr_t api, void* pReqQryLockVolumeField, int nRequestID);
int hs_td_ReqQueryExerciseVolume(uintptr_t api, void* pReqQueryExerciseVolumeField, int nRequestID);
int hs_td_ReqQryCombVolume(uintptr_t api, void* pReqQryCombVolumeField, int nRequestID);
int hs_td_ReqQryPosition(uintptr_t api, void* pReqQryPositionField, int nRequestID);
int hs_td_ReqQryTradingAccount(uintptr_t api, void* pReqQryTradingAccountField, int nRequestID);
int hs_td_ReqQryOrder(uintptr_t api, void* pReqQryOrderField, int nRequestID);
int hs_td_ReqQryTrade(uintptr_t api, void* pReqQryTradeField, int nRequestID);
int hs_td_ReqQryExercise(uintptr_t api, void* pReqQryExerciseField, int nRequestID);
int hs_td_ReqQryLock(uintptr_t api, void* pReqQryLockField, int nRequestID);
int hs_td_ReqQryCombAction(uintptr_t api, void* pReqQryCombActionField, int nRequestID);
int hs_td_ReqQryForQuote(uintptr_t api, void* pReqQryForQuoteField, int nRequestID);
int hs_td_ReqQryQuote(uintptr_t api, void* pReqQryQuoteField, int nRequestID);
int hs_td_ReqQryPositionCombineDetail(uintptr_t api, void* pReqQryPositionCombineDetailField, int nRequestID);
int hs_td_ReqQryInstrument(uintptr_t api, void* pReqQryInstrumentField, int nRequestID);
int hs_td_ReqQryCoveredShort(uintptr_t api, void* pReqQryCoveredShortField, int nRequestID);
int hs_td_ReqQryExerciseAssign(uintptr_t api, void* pReqQryExerciseAssignField, int nRequestID);
int hs_td_ReqTransfer(uintptr_t api, void* pReqTransferField, int nRequestID);
int hs_td_ReqQryTransfer(uintptr_t api, void* pReqQryTransferField, int nRequestID);
int hs_td_ReqQueryBankBalance(uintptr_t api, void* pReqQueryBankBalanceField, int nRequestID);
int hs_td_ReqQueryBankAccount(uintptr_t api, void* pReqQueryBankAccountField, int nRequestID);
int hs_td_ReqMultiCentreFundTrans(uintptr_t api, void* pReqMultiCentreFundTransField, int nRequestID);
int hs_td_ReqQueryBillContent(uintptr_t api, void* pReqQueryBillContentField, int nRequestID);
int hs_td_ReqBillConfirm(uintptr_t api, void* pReqBillConfirmField, int nRequestID);
int hs_td_ReqQryMargin(uintptr_t api, void* pReqQryMarginField, int nRequestID);
int hs_td_ReqQryCommission(uintptr_t api, void* pReqQryCommissionField, int nRequestID);
int hs_td_ReqQryExchangeRate(uintptr_t api, void* pReqQryExchangeRateField, int nRequestID);
int hs_td_ReqQryPositionDetail(uintptr_t api, void* pReqQryPositionDetailField, int nRequestID);
int hs_td_ReqQrySysConfig(uintptr_t api, void* pReqQrySysConfigField, int nRequestID);
int hs_td_ReqQryDepthMarketData(uintptr_t api, void* pReqQryDepthMarketDataField, int nRequestID);
int hs_td_ReqFundTrans(uintptr_t api, void* pReqFundTransferField, int nRequestID);
int hs_td_ReqQryFundTrans(uintptr_t api, void* pReqQryFundTransField, int nRequestID);
int hs_td_ReqQryClientNotice(uintptr_t api, void* pReqQryClientNoticeField, int nRequestID);
int hs_td_ReqQryOptUnderly(uintptr_t api, void* pReqQryOptUnderlyField, int nRequestID);
int hs_td_ReqQrySecuDepthMarket(uintptr_t api, void* pReqQrySecuDepthMarketField, int nRequestID);
int hs_td_ReqQryHistOrder(uintptr_t api, void* pReqQryHistOrderField, int nRequestID);
int hs_td_ReqQryHistTrade(uintptr_t api, void* pReqQryHistTradeField, int nRequestID);
int hs_td_ReqQryWithdrawFund(uintptr_t api, void* pReqQryWithdrawFundField, int nRequestID);
int hs_td_ReqQryCombInstrument(uintptr_t api, void* pReqQryCombInstrumentField, int nRequestID);
int hs_td_ReqQrySeatID(uintptr_t api, void* pReqQrySeatIDField, int nRequestID);
int hs_td_ReqOptionSelfClose(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID);
int hs_td_ReqOptionSelfCloseAction(uintptr_t api, void* pReqOptionSelfCloseActionField, int nRequestID);
int hs_td_ReqQryOptionSelfCloseResult(uintptr_t api, void* pReqQryOptionSelfCloseResultField, int nRequestID);
int hs_td_ReqQryOptionSelfClose(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID);
int hs_td_ReqOptQuoteInsert(uintptr_t api, void* pReqOptQuoteInsertField, int nRequestID);
int hs_td_ReqOptQuoteAction(uintptr_t api, void* pReqOptQuoteActionField, int nRequestID);
int hs_td_ReqQryOptQuote(uintptr_t api, void* pReqQryOptQuoteField, int nRequestID);
int hs_td_ReqQryOptCombStrategy(uintptr_t api, void* pReqQryOptCombStrategyField, int nRequestID);

#ifdef __cplusplus
}
#endif

#endif
