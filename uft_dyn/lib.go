package uft_dyn

/*
#cgo CXXFLAGS: -std=c++17 -I${SRCDIR}/../v3.7.4/sdk/include
#cgo LDFLAGS: -ldl -lpthread
#cgo windows LDFLAGS: -lws2_32

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

uintptr_t hs_create_md_api(const char* dllPath);
void hs_release_md_api(uintptr_t api);
const char* hs_md_GetApiVersion(uintptr_t api);
int hs_md_Init(uintptr_t api, void* config);
int hs_md_Join(uintptr_t api);
int hs_md_RegisterFront(uintptr_t api, const char* frontAddress);
int hs_md_RegisterFensServer(uintptr_t api, const char* fensAddress, const char* accountID);
void hs_md_RegisterSpi(uintptr_t api, uintptr_t spi);
int hs_md_ReqDepthMarketDataSubscribe(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqDepthMarketDataCancel(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqForQuoteSubscribe(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqForQuoteCancel(uintptr_t api, void* pReq, int count, int requestID);
const char* hs_md_GetApiErrorMsg(int errorCode);

// C function declarations for TradeApi
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
*/
import "C"
import "unsafe"

type HSApiHandle uintptr

func freeCString(cs *C.char) {
	if cs != nil {
		C.free(unsafe.Pointer(cs))
	}
}
