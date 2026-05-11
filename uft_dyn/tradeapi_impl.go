package uft_dyn

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

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

import (
	"sync"
	"unsafe"

	"github.com/nandy33up/go2uft/thost"
)

type TradeOption func(api *TradeApi)

type TradeApi struct {
	apiPtr  uintptr
	dllPath string
	spi     CHSTradeSpi
	mu      sync.RWMutex
}

type CHSTradeSpi interface {
	thost.CHSTradeSpi
}

var (
	tdSpiMap     = make(map[uintptr]CHSTradeSpi)
	tdSpiMapLock sync.RWMutex
)

func TradeDynamicLibPath(path string) TradeOption {
	return func(api *TradeApi) {
		api.dllPath = path
	}
}

func CreateTradeApi(options ...TradeOption) *TradeApi {
	api := &TradeApi{}
	for _, opt := range options {
		opt(api)
	}
	api.apiPtr = uintptr(C.hs_create_td_api(cString(api.dllPath)))
	return api
}

func (api *TradeApi) ReleaseApi() {
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.apiPtr != 0 {
		C.hs_release_td_api(C.uintptr_t(api.apiPtr))
		api.apiPtr = 0
	}
}

func (api *TradeApi) GetApiVersion() string {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return ""
	}
	cs := C.hs_td_GetApiVersion(C.uintptr_t(api.apiPtr))
	if cs == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func (api *TradeApi) Init(pInitCfgField *thost.CHSInitConfigField, pExtraParam unsafe.Pointer) int {
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_Init(C.uintptr_t(api.apiPtr), unsafe.Pointer(pInitCfgField)))
}

func (api *TradeApi) Join() int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_Join(C.uintptr_t(api.apiPtr)))
}

func (api *TradeApi) RegisterSubModel(eSubType thost.SUB_TERT_TYPE) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_RegisterSubModel(C.uintptr_t(api.apiPtr), C.int(eSubType)))
}

func (api *TradeApi) RegisterFront(pszFrontAddress string) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_RegisterFront(C.uintptr_t(api.apiPtr), cString(pszFrontAddress)))
}

func (api *TradeApi) RegisterFensServer(pszFensAddress string, pszAccountID string) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_RegisterFensServer(C.uintptr_t(api.apiPtr), cString(pszFensAddress), cString(pszAccountID)))
}

func (api *TradeApi) RegisterSpi(pSpi CHSTradeSpi) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.spi = pSpi
	if api.apiPtr != 0 && pSpi != nil {
		tdSpiMapLock.Lock()
		tdSpiMap[api.apiPtr] = pSpi
		tdSpiMapLock.Unlock()
		C.hs_td_RegisterSpi(C.uintptr_t(api.apiPtr), C.uintptr_t(api.apiPtr))
	}
}

func (api *TradeApi) GetApiErrorMsg(nErrorCode int) string {
	cs := C.hs_td_GetApiErrorMsg(C.int(nErrorCode))
	if cs == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func (api *TradeApi) GetTradingDate() int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return 0
	}
	return int(C.hs_td_GetTradingDate(C.uintptr_t(api.apiPtr)))
}

func (api *TradeApi) BindSessionID(nSessionID uint8) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_BindSessionID(C.uintptr_t(api.apiPtr), C.uint8_t(nSessionID)))
}

func (api *TradeApi) ReqAuthenticate(pReqAuthenticateField *thost.CHSReqAuthenticateField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqAuthenticate(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqAuthenticateField), C.int(nRequestID)))
}

func (api *TradeApi) ReqSubmitUserSystemInfo(pReqUserSystemInfoField *thost.CHSReqUserSystemInfoField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqSubmitUserSystemInfo(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserSystemInfoField), C.int(nRequestID)))
}

func (api *TradeApi) ReqUserLogin(pReqUserLoginField *thost.CHSReqUserLoginField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqUserLogin(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserLoginField), C.int(nRequestID)))
}

func (api *TradeApi) ReqUserPasswordUpdate(pReqUserPasswordUpdateField *thost.CHSReqUserPasswordUpdateField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqUserPasswordUpdate(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserPasswordUpdateField), C.int(nRequestID)))
}

func (api *TradeApi) ReqOrderInsert(pReqOrderInsert *thost.CHSReqOrderInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOrderInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOrderInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqOrderAction(pReqOrderAction *thost.CHSReqOrderActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOrderAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOrderAction), C.int(nRequestID)))
}

func (api *TradeApi) ReqExerciseOrderInsert(pReqExerciseOrderInsert *thost.CHSReqExerciseOrderInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqExerciseOrderInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqExerciseOrderInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqExerciseOrderAction(pReqExerciseOrderActionField *thost.CHSReqExerciseOrderActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqExerciseOrderAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqExerciseOrderActionField), C.int(nRequestID)))
}

func (api *TradeApi) ReqLockInsert(pReqLockInsert *thost.CHSReqLockInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqLockInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqLockInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqForQuoteInsert(pReqForQuoteInsertField *thost.CHSReqForQuoteInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqForQuoteInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqForQuoteInsertField), C.int(nRequestID)))
}

func (api *TradeApi) ReqQuoteInsert(pReqQuoteInsert *thost.CHSReqQuoteInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQuoteInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQuoteInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqQuoteAction(pReqQuoteAction *thost.CHSReqQuoteActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQuoteAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQuoteAction), C.int(nRequestID)))
}

func (api *TradeApi) ReqCombActionInsert(pReqCombActionInsert *thost.CHSReqCombActionInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqCombActionInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqCombActionInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqQueryMaxOrderVolume(pReqQueryMaxOrderVolume *thost.CHSReqQueryMaxOrderVolumeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQueryMaxOrderVolume(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryMaxOrderVolume), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryLockVolume(pReqQryLockVolume *thost.CHSReqQryLockVolumeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryLockVolume(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryLockVolume), C.int(nRequestID)))
}

func (api *TradeApi) ReqQueryExerciseVolume(pReqQueryExerciseVolume *thost.CHSReqQueryExerciseVolumeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQueryExerciseVolume(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryExerciseVolume), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryCombVolume(pReqQryCombVolume *thost.CHSReqQryCombVolumeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryCombVolume(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombVolume), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryPosition(pReqQryPosition *thost.CHSReqQryPositionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryPosition(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPosition), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryTradingAccount(pReqQryTradingAccount *thost.CHSReqQryTradingAccountField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryTradingAccount(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTradingAccount), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOrder(pReqQryOrder *thost.CHSReqQryOrderField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOrder(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOrder), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryTrade(pReqQryTrade *thost.CHSReqQryTradeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryTrade(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTrade), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryExercise(pReqQryExercise *thost.CHSReqQryExerciseField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryExercise(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExercise), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryLock(pReqQryLock *thost.CHSReqQryLockField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryLock(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryLock), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryCombAction(pReqQryCombAction *thost.CHSReqQryCombActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryCombAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombAction), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryForQuote(pReqQryForQuote *thost.CHSReqQryForQuoteField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryForQuote(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryForQuote), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryQuote(pReqQryQuote *thost.CHSReqQryQuoteField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryQuote(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryQuote), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryPositionCombineDetail(pReqQryPositionCombineDetail *thost.CHSReqQryPositionCombineDetailField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryPositionCombineDetail(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPositionCombineDetail), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryInstrument(pReqQryInstrument *thost.CHSReqQryInstrumentField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryInstrument(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryInstrument), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryCoveredShort(pReqQryCoveredShort *thost.CHSReqQryCoveredShortField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryCoveredShort(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCoveredShort), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryExerciseAssign(pReqQryExerciseAssign *thost.CHSReqQryExerciseAssignField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryExerciseAssign(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExerciseAssign), C.int(nRequestID)))
}

func (api *TradeApi) ReqTransfer(pReqTransfer *thost.CHSReqTransferField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqTransfer(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqTransfer), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryTransfer(pReqQryTransfer *thost.CHSReqQryTransferField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryTransfer(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTransfer), C.int(nRequestID)))
}

func (api *TradeApi) ReqQueryBankBalance(pReqQueryBankBalance *thost.CHSReqQueryBankBalanceField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQueryBankBalance(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBankBalance), C.int(nRequestID)))
}

func (api *TradeApi) ReqQueryBankAccount(pReqQueryBankAccount *thost.CHSReqQueryBankAccountField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQueryBankAccount(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBankAccount), C.int(nRequestID)))
}

func (api *TradeApi) ReqMultiCentreFundTrans(pReqMultiCentreFundTransfer *thost.CHSReqMultiCentreFundTransField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqMultiCentreFundTrans(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqMultiCentreFundTransfer), C.int(nRequestID)))
}

func (api *TradeApi) ReqQueryBillContent(pReqQueryBillContent *thost.CHSReqQueryBillContentField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQueryBillContent(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBillContent), C.int(nRequestID)))
}

func (api *TradeApi) ReqBillConfirm(pReqBillConfirm *thost.CHSReqBillConfirmField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqBillConfirm(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqBillConfirm), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryMargin(pReqQryMargin *thost.CHSReqQryMarginField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryMargin(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryMargin), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryCommission(pReqQryCommission *thost.CHSReqQryCommissionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryCommission(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCommission), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryPositionDetail(pReqQryPositionDetail *thost.CHSReqQryPositionDetailField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryPositionDetail(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPositionDetail), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryExchangeRate(pReqQryExchangeRate *thost.CHSReqQryExchangeRateField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryExchangeRate(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExchangeRate), C.int(nRequestID)))
}

func (api *TradeApi) ReqQrySysConfig(pReqQrySysConfig *thost.CHSReqQrySysConfigField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQrySysConfig(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySysConfig), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryDepthMarketData(pReqQryDepthMarketData *thost.CHSReqQryDepthMarketDataField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryDepthMarketData(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryDepthMarketData), C.int(nRequestID)))
}

func (api *TradeApi) ReqFundTrans(pReqFundTransfer *thost.CHSReqFundTransField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqFundTrans(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqFundTransfer), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryFundTrans(pReqQryFundTrans *thost.CHSReqQryFundTransField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryFundTrans(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryFundTrans), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryClientNotice(pReqQryClientNotice *thost.CHSReqQryClientNoticeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryClientNotice(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryClientNotice), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOptUnderly(pReqQryOptUnderly *thost.CHSReqQryOptUnderlyField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOptUnderly(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptUnderly), C.int(nRequestID)))
}

func (api *TradeApi) ReqQrySecuDepthMarket(pReqQrySecuDepthMarket *thost.CHSReqQrySecuDepthMarketField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQrySecuDepthMarket(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySecuDepthMarket), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryHistOrder(pReqQryHistOrder *thost.CHSReqQryHistOrderField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryHistOrder(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryHistOrder), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryHistTrade(pReqQryHistTrade *thost.CHSReqQryHistTradeField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryHistTrade(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryHistTrade), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryWithdrawFund(pReqQryWithdrawFund *thost.CHSReqQryWithdrawFundField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryWithdrawFund(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryWithdrawFund), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryCombInstrument(pReqQryCombInstrument *thost.CHSReqQryCombInstrumentField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryCombInstrument(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombInstrument), C.int(nRequestID)))
}

func (api *TradeApi) ReqQrySeatID(pReqQrySeatID *thost.CHSReqQrySeatIDField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQrySeatID(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySeatID), C.int(nRequestID)))
}

func (api *TradeApi) ReqOptionSelfClose(pReqOptionSelfClose *thost.CHSReqOptionSelfCloseField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOptionSelfClose(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfClose), C.int(nRequestID)))
}

func (api *TradeApi) ReqOptionSelfCloseAction(pReqOptionSelfCloseAction *thost.CHSReqOptionSelfCloseActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOptionSelfCloseAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfCloseAction), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *thost.CHSReqQryOptionSelfCloseResultField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOptionSelfCloseResult(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptionSelfCloseResult), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOptionSelfClose(pReqOptionSelfClose *thost.CHSReqOptionSelfCloseField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOptionSelfClose(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfClose), C.int(nRequestID)))
}

func (api *TradeApi) ReqOptQuoteInsert(pReqOptQuoteInsert *thost.CHSReqOptQuoteInsertField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOptQuoteInsert(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptQuoteInsert), C.int(nRequestID)))
}

func (api *TradeApi) ReqOptQuoteAction(pReqOptQuoteAction *thost.CHSReqOptQuoteActionField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqOptQuoteAction(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptQuoteAction), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOptQuote(pReqQryOptQuote *thost.CHSReqQryOptQuoteField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOptQuote(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptQuote), C.int(nRequestID)))
}

func (api *TradeApi) ReqQryOptCombStrategy(pReqQryOptCombStrategy *thost.CHSReqQryOptCombStrategyField, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_td_ReqQryOptCombStrategy(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptCombStrategy), C.int(nRequestID)))
}

func getTdSpi(handle uintptr) CHSTradeSpi {
	tdSpiMapLock.RLock()
	defer tdSpiMapLock.RUnlock()
	return tdSpiMap[handle]
}

//export go_td_OnFrontConnected
func go_td_OnFrontConnected(v C.uintptr_t, _dummy C.int) {
	spi := getTdSpi(uintptr(v))
	if spi == nil {
		return
	}
	spi.OnFrontConnected()
}

//export go_td_OnFrontDisconnected
func go_td_OnFrontDisconnected(v uintptr, p0 C.int) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnFrontDisconnected(int(p0))
}

//export go_td_OnRspAuthenticate
func go_td_OnRspAuthenticate(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspAuthenticate((*thost.CHSRspAuthenticateField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspSubmitUserSystemInfo
func go_td_OnRspSubmitUserSystemInfo(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspSubmitUserSystemInfo((*thost.CHSRspUserSystemInfoField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspUserLogin
func go_td_OnRspUserLogin(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspUserLogin((*thost.CHSRspUserLoginField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspUserPasswordUpdate
func go_td_OnRspUserPasswordUpdate(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspUserPasswordUpdate((*thost.CHSRspUserPasswordUpdateField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorOrderInsert
func go_td_OnRspErrorOrderInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorOrderInsert((*thost.CHSRspOrderInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspOrderAction
func go_td_OnRspOrderAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspOrderAction((*thost.CHSRspOrderActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorExerciseOrderInsert
func go_td_OnRspErrorExerciseOrderInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorExerciseOrderInsert((*thost.CHSRspExerciseOrderInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspExerciseOrderAction
func go_td_OnRspExerciseOrderAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspExerciseOrderAction((*thost.CHSRspExerciseOrderActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorLockInsert
func go_td_OnRspErrorLockInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorLockInsert((*thost.CHSRspLockInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspForQuoteInsert
func go_td_OnRspForQuoteInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspForQuoteInsert((*thost.CHSRspForQuoteInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorQuoteInsert
func go_td_OnRspErrorQuoteInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorQuoteInsert((*thost.CHSRspQuoteInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQuoteAction
func go_td_OnRspQuoteAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQuoteAction((*thost.CHSRspQuoteActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorCombActionInsert
func go_td_OnRspErrorCombActionInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorCombActionInsert((*thost.CHSRspCombActionInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQueryMaxOrderVolume
func go_td_OnRspQueryMaxOrderVolume(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQueryMaxOrderVolume((*thost.CHSRspQueryMaxOrderVolumeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQueryExerciseVolume
func go_td_OnRspQueryExerciseVolume(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQueryExerciseVolume((*thost.CHSRspQueryExerciseVolumeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryLockVolume
func go_td_OnRspQryLockVolume(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryLockVolume((*thost.CHSRspQryLockVolumeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryCombVolume
func go_td_OnRspQryCombVolume(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryCombVolume((*thost.CHSRspQryCombVolumeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryPosition
func go_td_OnRspQryPosition(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryPosition((*thost.CHSRspQryPositionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryTradingAccount
func go_td_OnRspQryTradingAccount(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryTradingAccount((*thost.CHSRspQryTradingAccountField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOrder
func go_td_OnRspQryOrder(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOrder((*thost.CHSOrderField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryTrade
func go_td_OnRspQryTrade(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryTrade((*thost.CHSTradeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryExercise
func go_td_OnRspQryExercise(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryExercise((*thost.CHSExerciseField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryLock
func go_td_OnRspQryLock(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryLock((*thost.CHSLockField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryCombAction
func go_td_OnRspQryCombAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryCombAction((*thost.CHSCombActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryForQuote
func go_td_OnRspQryForQuote(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryForQuote((*thost.CHSForQuoteField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryQuote
func go_td_OnRspQryQuote(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryQuote((*thost.CHSQuoteField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryPositionCombineDetail
func go_td_OnRspQryPositionCombineDetail(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryPositionCombineDetail((*thost.CHSRspQryPositionCombineDetailField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryInstrument
func go_td_OnRspQryInstrument(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryInstrument((*thost.CHSRspQryInstrumentField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryCoveredShort
func go_td_OnRspQryCoveredShort(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryCoveredShort((*thost.CHSRspQryCoveredShortField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryExerciseAssign
func go_td_OnRspQryExerciseAssign(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryExerciseAssign((*thost.CHSRspQryExerciseAssignField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspTransfer
func go_td_OnRspTransfer(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspTransfer((*thost.CHSRspTransferField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryTransfer
func go_td_OnRspQryTransfer(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryTransfer((*thost.CHSRspQryTransferField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQueryBankBalance
func go_td_OnRspQueryBankBalance(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQueryBankBalance((*thost.CHSRspQueryBankBalanceField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQueryBankAccount
func go_td_OnRspQueryBankAccount(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQueryBankAccount((*thost.CHSRspQueryBankAccountField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspMultiCentreFundTrans
func go_td_OnRspMultiCentreFundTrans(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspMultiCentreFundTrans((*thost.CHSRspMultiCentreFundTransField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQueryBillContent
func go_td_OnRspQueryBillContent(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQueryBillContent((*thost.CHSRspQueryBillContentField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspBillConfirm
func go_td_OnRspBillConfirm(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspBillConfirm((*thost.CHSRspBillConfirmField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryMargin
func go_td_OnRspQryMargin(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryMargin((*thost.CHSRspQryMarginField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryCommission
func go_td_OnRspQryCommission(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryCommission((*thost.CHSRspQryCommissionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryExchangeRate
func go_td_OnRspQryExchangeRate(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryExchangeRate((*thost.CHSRspQryExchangeRateField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryPositionDetail
func go_td_OnRspQryPositionDetail(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryPositionDetail((*thost.CHSRspQryPositionDetailField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQrySysConfig
func go_td_OnRspQrySysConfig(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQrySysConfig((*thost.CHSRspQrySysConfigField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryDepthMarketData
func go_td_OnRspQryDepthMarketData(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryDepthMarketData((*thost.CHSDepthMarketDataField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspFundTrans
func go_td_OnRspFundTrans(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspFundTrans((*thost.CHSRspFundTransField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryFundTrans
func go_td_OnRspQryFundTrans(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryFundTrans((*thost.CHSRspQryFundTransField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryClientNotice
func go_td_OnRspQryClientNotice(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryClientNotice((*thost.CHSClientNoticeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOptUnderly
func go_td_OnRspQryOptUnderly(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOptUnderly((*thost.CHSRspQryOptUnderlyField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQrySecuDepthMarket
func go_td_OnRspQrySecuDepthMarket(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQrySecuDepthMarket((*thost.CHSRspQrySecuDepthMarketField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryHistOrder
func go_td_OnRspQryHistOrder(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryHistOrder((*thost.CHSOrderField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryHistTrade
func go_td_OnRspQryHistTrade(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryHistTrade((*thost.CHSTradeField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryWithdrawFund
func go_td_OnRspQryWithdrawFund(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryWithdrawFund((*thost.CHSRspQryWithdrawFundField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryCombInstrument
func go_td_OnRspQryCombInstrument(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryCombInstrument((*thost.CHSRspQryCombInstrumentField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQrySeatID
func go_td_OnRspQrySeatID(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQrySeatID((*thost.CHSRspQrySeatIDField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspOptionSelfClose
func go_td_OnRspOptionSelfClose(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspOptionSelfClose((*thost.CHSRspOptionSelfCloseField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspOptionSelfCloseAction
func go_td_OnRspOptionSelfCloseAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspOptionSelfCloseAction((*thost.CHSRspOptionSelfCloseActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOptionSelfClose
func go_td_OnRspQryOptionSelfClose(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOptionSelfClose((*thost.CHSOptionSelfCloseField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOptionSelfCloseResult
func go_td_OnRspQryOptionSelfCloseResult(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOptionSelfCloseResult((*thost.CHSRspQryOptionSelfCloseResultField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspOptQuoteAction
func go_td_OnRspOptQuoteAction(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspOptQuoteAction((*thost.CHSRspOptQuoteActionField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspErrorOptQuoteInsert
func go_td_OnRspErrorOptQuoteInsert(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspErrorOptQuoteInsert((*thost.CHSRspOptQuoteInsertField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOptQuote
func go_td_OnRspQryOptQuote(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOptQuote((*thost.CHSOptQuoteField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRspQryOptCombStrategy
func go_td_OnRspQryOptCombStrategy(v uintptr, p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.int, p3 C.bool) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspQryOptCombStrategy((*thost.CHSRspQryOptCombStrategyField)(p0), (*thost.CHSRspInfoField)(p1), int(p2), bool(p3))
}

//export go_td_OnRtnTrade
func go_td_OnRtnTrade(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnTrade((*thost.CHSTradeField)(p0))
}

//export go_td_OnRtnOrder
func go_td_OnRtnOrder(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnOrder((*thost.CHSOrderField)(p0))
}

//export go_td_OnRtnExercise
func go_td_OnRtnExercise(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnExercise((*thost.CHSExerciseField)(p0))
}

//export go_td_OnRtnCombAction
func go_td_OnRtnCombAction(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnCombAction((*thost.CHSCombActionField)(p0))
}

//export go_td_OnRtnLock
func go_td_OnRtnLock(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnLock((*thost.CHSLockField)(p0))
}

//export go_td_OnErrRtnOrderAction
func go_td_OnErrRtnOrderAction(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnErrRtnOrderAction((*thost.CHSOrderActionField)(p0))
}

//export go_td_OnRtnClientNotice
func go_td_OnRtnClientNotice(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnClientNotice((*thost.CHSClientNoticeField)(p0))
}

//export go_td_OnRtnQuote
func go_td_OnRtnQuote(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnQuote((*thost.CHSQuoteField)(p0))
}

//export go_td_OnRtnForQuote
func go_td_OnRtnForQuote(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnForQuote((*thost.CHSForQuoteField)(p0))
}

//export go_td_OnRtnExchangeStatus
func go_td_OnRtnExchangeStatus(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnExchangeStatus((*thost.CHSExchangeStatusField)(p0))
}

//export go_td_OnRtnProductStatus
func go_td_OnRtnProductStatus(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnProductStatus((*thost.CHSProductStatusField)(p0))
}

//export go_td_OnRtnOptionSelfClose
func go_td_OnRtnOptionSelfClose(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnOptionSelfClose((*thost.CHSOptionSelfCloseField)(p0))
}

//export go_td_OnRtnOptQuote
func go_td_OnRtnOptQuote(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnOptQuote((*thost.CHSOptQuoteField)(p0))
}

//export go_td_OnErrRtnOptQuoteAction
func go_td_OnErrRtnOptQuoteAction(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnErrRtnOptQuoteAction((*thost.CHSOptQuoteActionField)(p0))
}

//export go_td_OnRtnTransfer
func go_td_OnRtnTransfer(v uintptr, p0 unsafe.Pointer) {
	spi := getTdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnTransfer((*thost.CHSTransferField)(p0))
}
