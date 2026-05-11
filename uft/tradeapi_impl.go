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

/*
#cgo linux LDFLAGS: -fPIC -L. -L${SRCDIR}/../v3.7.4/sdk/linux.x64 -Wl,-rpath=${SRCDIR}/../v3.7.4/sdk/linux.x64 -lHSMdApi -lHSTradeApi -lldptcpsdk -lt2sdk -lHsFutuSystemInfo -lsmf_api -linfocertsdk -ldl -lpthread -lstdc++
#cgo linux CXXFLAGS: -std=c++17
#cgo linux CPPFLAGS: -fPIC -I. -I${SRCDIR}/../v3.7.4/sdk/include/

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

#include "HSDataType.h"
#include "HSStruct.h"

uintptr_t hs_create_td_api_static(const char* flowPath);
void hs_release_td_api_static(uintptr_t api);
const char* hs_td_GetApiVersion_static(uintptr_t api);
int hs_td_Init_static(uintptr_t api, void* config);
int hs_td_Join_static(uintptr_t api);
int hs_td_RegisterSubModel_static(uintptr_t api, int eSubType);
int hs_td_RegisterFront_static(uintptr_t api, const char* frontAddress);
int hs_td_RegisterFensServer_static(uintptr_t api, const char* fensAddress, const char* accountID);
void hs_td_RegisterSpi_static(uintptr_t api, uintptr_t spi);
const char* hs_td_GetApiErrorMsg_static(int errorCode);
int hs_td_GetTradingDate_static(uintptr_t api);
int hs_td_BindSessionID_static(uintptr_t api, uint8_t nSessionID);
int hs_td_ReqAuthenticate_static(uintptr_t api, void* pReqAuthenticateField, int nRequestID);
int hs_td_ReqSubmitUserSystemInfo_static(uintptr_t api, void* pReqUserSystemInfoField, int nRequestID);
int hs_td_ReqUserLogin_static(uintptr_t api, void* pReqUserLoginField, int nRequestID);
int hs_td_ReqUserPasswordUpdate_static(uintptr_t api, void* pReqUserPasswordUpdateField, int nRequestID);
int hs_td_ReqOrderInsert_static(uintptr_t api, void* pReqOrderInsertField, int nRequestID);
int hs_td_ReqOrderAction_static(uintptr_t api, void* pReqOrderActionField, int nRequestID);
int hs_td_ReqExerciseOrderInsert_static(uintptr_t api, void* pReqExerciseOrderInsertField, int nRequestID);
int hs_td_ReqExerciseOrderAction_static(uintptr_t api, void* pReqExerciseOrderActionField, int nRequestID);
int hs_td_ReqLockInsert_static(uintptr_t api, void* pReqLockInsertField, int nRequestID);
int hs_td_ReqForQuoteInsert_static(uintptr_t api, void* pReqForQuoteInsertField, int nRequestID);
int hs_td_ReqQuoteInsert_static(uintptr_t api, void* pReqQuoteInsertField, int nRequestID);
int hs_td_ReqQuoteAction_static(uintptr_t api, void* pReqQuoteActionField, int nRequestID);
int hs_td_ReqCombActionInsert_static(uintptr_t api, void* pReqCombActionInsertField, int nRequestID);
int hs_td_ReqQueryMaxOrderVolume_static(uintptr_t api, void* pReqQueryMaxOrderVolumeField, int nRequestID);
int hs_td_ReqQryLockVolume_static(uintptr_t api, void* pReqQryLockVolumeField, int nRequestID);
int hs_td_ReqQueryExerciseVolume_static(uintptr_t api, void* pReqQueryExerciseVolumeField, int nRequestID);
int hs_td_ReqQryCombVolume_static(uintptr_t api, void* pReqQryCombVolumeField, int nRequestID);
int hs_td_ReqQryPosition_static(uintptr_t api, void* pReqQryPositionField, int nRequestID);
int hs_td_ReqQryTradingAccount_static(uintptr_t api, void* pReqQryTradingAccountField, int nRequestID);
int hs_td_ReqQryOrder_static(uintptr_t api, void* pReqQryOrderField, int nRequestID);
int hs_td_ReqQryTrade_static(uintptr_t api, void* pReqQryTradeField, int nRequestID);
int hs_td_ReqQryExercise_static(uintptr_t api, void* pReqQryExerciseField, int nRequestID);
int hs_td_ReqQryLock_static(uintptr_t api, void* pReqQryLockField, int nRequestID);
int hs_td_ReqQryCombAction_static(uintptr_t api, void* pReqQryCombActionField, int nRequestID);
int hs_td_ReqQryForQuote_static(uintptr_t api, void* pReqQryForQuoteField, int nRequestID);
int hs_td_ReqQryQuote_static(uintptr_t api, void* pReqQryQuoteField, int nRequestID);
int hs_td_ReqQryPositionCombineDetail_static(uintptr_t api, void* pReqQryPositionCombineDetailField, int nRequestID);
int hs_td_ReqQryInstrument_static(uintptr_t api, void* pReqQryInstrumentField, int nRequestID);
int hs_td_ReqQryCoveredShort_static(uintptr_t api, void* pReqQryCoveredShortField, int nRequestID);
int hs_td_ReqQryExerciseAssign_static(uintptr_t api, void* pReqQryExerciseAssignField, int nRequestID);
int hs_td_ReqTransfer_static(uintptr_t api, void* pReqTransferField, int nRequestID);
int hs_td_ReqQryTransfer_static(uintptr_t api, void* pReqQryTransferField, int nRequestID);
int hs_td_ReqQueryBankBalance_static(uintptr_t api, void* pReqQueryBankBalanceField, int nRequestID);
int hs_td_ReqQueryBankAccount_static(uintptr_t api, void* pReqQueryBankAccountField, int nRequestID);
int hs_td_ReqMultiCentreFundTrans_static(uintptr_t api, void* pReqMultiCentreFundTransferField, int nRequestID);
int hs_td_ReqQueryBillContent_static(uintptr_t api, void* pReqQueryBillContentField, int nRequestID);
int hs_td_ReqBillConfirm_static(uintptr_t api, void* pReqBillConfirmField, int nRequestID);
int hs_td_ReqQryMargin_static(uintptr_t api, void* pReqQryMarginField, int nRequestID);
int hs_td_ReqQryCommission_static(uintptr_t api, void* pReqQryCommissionField, int nRequestID);
int hs_td_ReqQryExchangeRate_static(uintptr_t api, void* pReqQryExchangeRateField, int nRequestID);
int hs_td_ReqQryPositionDetail_static(uintptr_t api, void* pReqQryPositionDetailField, int nRequestID);
int hs_td_ReqQrySysConfig_static(uintptr_t api, void* pReqQrySysConfigField, int nRequestID);
int hs_td_ReqQryDepthMarketData_static(uintptr_t api, void* pReqQryDepthMarketDataField, int nRequestID);
int hs_td_ReqFundTrans_static(uintptr_t api, void* pReqFundTransferField, int nRequestID);
int hs_td_ReqQryFundTrans_static(uintptr_t api, void* pReqQryFundTransField, int nRequestID);
int hs_td_ReqQryClientNotice_static(uintptr_t api, void* pReqQryClientNoticeField, int nRequestID);
int hs_td_ReqQryOptUnderly_static(uintptr_t api, void* pReqQryOptUnderlyField, int nRequestID);
int hs_td_ReqQrySecuDepthMarket_static(uintptr_t api, void* pReqQrySecuDepthMarketField, int nRequestID);
int hs_td_ReqQryHistOrder_static(uintptr_t api, void* pReqQryHistOrderField, int nRequestID);
int hs_td_ReqQryHistTrade_static(uintptr_t api, void* pReqQryHistTradeField, int nRequestID);
int hs_td_ReqQryWithdrawFund_static(uintptr_t api, void* pReqQryWithdrawFundField, int nRequestID);
int hs_td_ReqQryCombInstrument_static(uintptr_t api, void* pReqQryCombInstrumentField, int nRequestID);
int hs_td_ReqQrySeatID_static(uintptr_t api, void* pReqQrySeatIDField, int nRequestID);
int hs_td_ReqOptionSelfClose_static(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID);
int hs_td_ReqOptionSelfCloseAction_static(uintptr_t api, void* pReqOptionSelfCloseActionField, int nRequestID);
int hs_td_ReqQryOptionSelfCloseResult_static(uintptr_t api, void* pReqQryOptionSelfCloseResultField, int nRequestID);
int hs_td_ReqQryOptionSelfClose_static(uintptr_t api, void* pReqQryOptionSelfCloseField, int nRequestID);
int hs_td_ReqOptQuoteInsert_static(uintptr_t api, void* pReqOptQuoteInsertField, int nRequestID);
int hs_td_ReqOptQuoteAction_static(uintptr_t api, void* pReqOptQuoteActionField, int nRequestID);
int hs_td_ReqQryOptQuote_static(uintptr_t api, void* pReqQryOptQuoteField, int nRequestID);
int hs_td_ReqQryOptCombStrategy_static(uintptr_t api, void* pReqQryOptCombStrategyField, int nRequestID);
*/
import "C"
import (
	"os"
	"path/filepath"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/nandy33up/go2uft/thost"
)

func CreateTradeApi(options ...TradeOption) thost.CHSTradeApi {
	api := &CTradeApi{
		flowPath:   defaultFlowPath,
		production: defaultIsProductionMode,
	}
	api.handle = cgo.NewHandle(api)
	for _, opt := range options {
		opt(api)
	}
	if api.flowPath != "" {
		var err error
		if strings.HasSuffix(api.flowPath, "/") {
			err = os.MkdirAll(api.flowPath, os.ModePerm)
		} else {
			parentDir := filepath.Dir(api.flowPath)
			err = os.MkdirAll(parentDir, os.ModePerm)
		}
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
	}
	cflowPath := C.CString(api.flowPath)
	defer C.free(unsafe.Pointer(cflowPath))

	api.apiPtr = uintptr(C.hs_create_td_api_static(cflowPath))
	return api
}

func (api *CTradeApi) GetApiVersion() string {
	cs := C.hs_td_GetApiVersion_static(C.uintptr_t(api.apiPtr))
	return C.GoString(cs)
}

func (api *CTradeApi) ReleaseApi() {
	C.hs_td_RegisterSpi_static(C.uintptr_t(api.apiPtr), C.uintptr_t(0))
	C.hs_release_td_api_static(C.uintptr_t(api.apiPtr))
	if api.handle != 0 {
		api.handle.Delete()
		api.handle = 0
	}
}

func (api *CTradeApi) Init(pInitCfgField *thost.CHSInitConfigField, pExtraParam unsafe.Pointer) int {
	return int(C.hs_td_Init_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pInitCfgField)))
}

func (api *CTradeApi) Join() int {
	return int(C.hs_td_Join_static(C.uintptr_t(api.apiPtr)))
}

func (api *CTradeApi) RegisterSubModel(eSubType thost.SUB_TERT_TYPE) int {
	return int(C.hs_td_RegisterSubModel_static(C.uintptr_t(api.apiPtr), C.int(eSubType)))
}

func (api *CTradeApi) RegisterFront(pszFrontAddress string) int {
	addr := C.CString(pszFrontAddress)
	defer C.free(unsafe.Pointer(addr))
	return int(C.hs_td_RegisterFront_static(C.uintptr_t(api.apiPtr), addr))
}

func (api *CTradeApi) RegisterFensServer(pszFensAddress string, pszAccountID string) int {
	fensAddr := C.CString(pszFensAddress)
	defer C.free(unsafe.Pointer(fensAddr))
	accountID := C.CString(pszAccountID)
	defer C.free(unsafe.Pointer(accountID))
	return int(C.hs_td_RegisterFensServer_static(C.uintptr_t(api.apiPtr), fensAddr, accountID))
}

func (api *CTradeApi) RegisterSpi(pSpi thost.CHSTradeSpi) {
	api.spi = pSpi
}

func (api *CTradeApi) GetApiErrorMsg(nErrorCode int) string {
	cs := C.hs_td_GetApiErrorMsg_static(C.int(nErrorCode))
	if cs == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func (api *CTradeApi) GetTradingDate() int {
	return int(C.hs_td_GetTradingDate_static(C.uintptr_t(api.apiPtr)))
}

func (api *CTradeApi) BindSessionID(nSessionID uint8) int {
	return int(C.hs_td_BindSessionID_static(C.uintptr_t(api.apiPtr), C.uint8_t(nSessionID)))
}

func (api *CTradeApi) ReqAuthenticate(pReqAuthenticateField *thost.CHSReqAuthenticateField, nRequestID int) int {
	return int(C.hs_td_ReqAuthenticate_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqAuthenticateField), C.int(nRequestID)))
}

func (api *CTradeApi) ReqSubmitUserSystemInfo(pReqUserSystemInfo *thost.CHSReqUserSystemInfoField, nRequestID int) int {
	return int(C.hs_td_ReqSubmitUserSystemInfo_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserSystemInfo), C.int(nRequestID)))
}

func (api *CTradeApi) ReqUserLogin(pReqUserLoginField *thost.CHSReqUserLoginField, nRequestID int) int {
	return int(C.hs_td_ReqUserLogin_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserLoginField), C.int(nRequestID)))
}

func (api *CTradeApi) ReqUserPasswordUpdate(pReqUserPasswordUpdateField *thost.CHSReqUserPasswordUpdateField, nRequestID int) int {
	return int(C.hs_td_ReqUserPasswordUpdate_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqUserPasswordUpdateField), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOrderInsert(pReqOrderInsert *thost.CHSReqOrderInsertField, nRequestID int) int {
	return int(C.hs_td_ReqOrderInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOrderInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOrderAction(pReqOrderAction *thost.CHSReqOrderActionField, nRequestID int) int {
	return int(C.hs_td_ReqOrderAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOrderAction), C.int(nRequestID)))
}

func (api *CTradeApi) ReqExerciseOrderInsert(pReqExerciseOrderInsert *thost.CHSReqExerciseOrderInsertField, nRequestID int) int {
	return int(C.hs_td_ReqExerciseOrderInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqExerciseOrderInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqExerciseOrderAction(pReqExerciseOrderActionField *thost.CHSReqExerciseOrderActionField, nRequestID int) int {
	return int(C.hs_td_ReqExerciseOrderAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqExerciseOrderActionField), C.int(nRequestID)))
}

func (api *CTradeApi) ReqLockInsert(pReqLockInsert *thost.CHSReqLockInsertField, nRequestID int) int {
	return int(C.hs_td_ReqLockInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqLockInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqForQuoteInsert(pReqForQuoteInsertField *thost.CHSReqForQuoteInsertField, nRequestID int) int {
	return int(C.hs_td_ReqForQuoteInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqForQuoteInsertField), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQuoteInsert(pReqQuoteInsert *thost.CHSReqQuoteInsertField, nRequestID int) int {
	return int(C.hs_td_ReqQuoteInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQuoteInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQuoteAction(pReqQuoteAction *thost.CHSReqQuoteActionField, nRequestID int) int {
	return int(C.hs_td_ReqQuoteAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQuoteAction), C.int(nRequestID)))
}

func (api *CTradeApi) ReqCombActionInsert(pReqCombActionInsert *thost.CHSReqCombActionInsertField, nRequestID int) int {
	return int(C.hs_td_ReqCombActionInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqCombActionInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQueryMaxOrderVolume(pReqQueryMaxOrderVolume *thost.CHSReqQueryMaxOrderVolumeField, nRequestID int) int {
	return int(C.hs_td_ReqQueryMaxOrderVolume_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryMaxOrderVolume), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryLockVolume(pReqQryLockVolume *thost.CHSReqQryLockVolumeField, nRequestID int) int {
	return int(C.hs_td_ReqQryLockVolume_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryLockVolume), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQueryExerciseVolume(pReqQueryExerciseVolume *thost.CHSReqQueryExerciseVolumeField, nRequestID int) int {
	return int(C.hs_td_ReqQueryExerciseVolume_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryExerciseVolume), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryCombVolume(pReqQryCombVolume *thost.CHSReqQryCombVolumeField, nRequestID int) int {
	return int(C.hs_td_ReqQryCombVolume_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombVolume), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryPosition(pReqQryPosition *thost.CHSReqQryPositionField, nRequestID int) int {
	return int(C.hs_td_ReqQryPosition_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPosition), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryTradingAccount(pReqQryTradingAccount *thost.CHSReqQryTradingAccountField, nRequestID int) int {
	return int(C.hs_td_ReqQryTradingAccount_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTradingAccount), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOrder(pReqQryOrder *thost.CHSReqQryOrderField, nRequestID int) int {
	return int(C.hs_td_ReqQryOrder_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOrder), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryTrade(pReqQryTrade *thost.CHSReqQryTradeField, nRequestID int) int {
	return int(C.hs_td_ReqQryTrade_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTrade), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryExercise(pReqQryExercise *thost.CHSReqQryExerciseField, nRequestID int) int {
	return int(C.hs_td_ReqQryExercise_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExercise), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryLock(pReqQryLock *thost.CHSReqQryLockField, nRequestID int) int {
	return int(C.hs_td_ReqQryLock_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryLock), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryCombAction(pReqQryCombAction *thost.CHSReqQryCombActionField, nRequestID int) int {
	return int(C.hs_td_ReqQryCombAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombAction), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryForQuote(pReqQryForQuote *thost.CHSReqQryForQuoteField, nRequestID int) int {
	return int(C.hs_td_ReqQryForQuote_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryForQuote), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryQuote(pReqQryQuote *thost.CHSReqQryQuoteField, nRequestID int) int {
	return int(C.hs_td_ReqQryQuote_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryQuote), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryPositionCombineDetail(pReqQryPositionCombineDetail *thost.CHSReqQryPositionCombineDetailField, nRequestID int) int {
	return int(C.hs_td_ReqQryPositionCombineDetail_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPositionCombineDetail), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryInstrument(pReqQryInstrument *thost.CHSReqQryInstrumentField, nRequestID int) int {
	return int(C.hs_td_ReqQryInstrument_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryInstrument), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryCoveredShort(pReqQryCoveredShort *thost.CHSReqQryCoveredShortField, nRequestID int) int {
	return int(C.hs_td_ReqQryCoveredShort_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCoveredShort), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryExerciseAssign(pReqQryExerciseAssign *thost.CHSReqQryExerciseAssignField, nRequestID int) int {
	return int(C.hs_td_ReqQryExerciseAssign_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExerciseAssign), C.int(nRequestID)))
}

func (api *CTradeApi) ReqTransfer(pReqTransfer *thost.CHSReqTransferField, nRequestID int) int {
	return int(C.hs_td_ReqTransfer_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqTransfer), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryTransfer(pReqQryTransfer *thost.CHSReqQryTransferField, nRequestID int) int {
	return int(C.hs_td_ReqQryTransfer_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryTransfer), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQueryBankBalance(pReqQueryBankBalance *thost.CHSReqQueryBankBalanceField, nRequestID int) int {
	return int(C.hs_td_ReqQueryBankBalance_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBankBalance), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQueryBankAccount(pReqQueryBankAccount *thost.CHSReqQueryBankAccountField, nRequestID int) int {
	return int(C.hs_td_ReqQueryBankAccount_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBankAccount), C.int(nRequestID)))
}

func (api *CTradeApi) ReqMultiCentreFundTrans(pReqMultiCentreFundTransfer *thost.CHSReqMultiCentreFundTransField, nRequestID int) int {
	return int(C.hs_td_ReqMultiCentreFundTrans_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqMultiCentreFundTransfer), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQueryBillContent(pReqQueryBillContent *thost.CHSReqQueryBillContentField, nRequestID int) int {
	return int(C.hs_td_ReqQueryBillContent_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQueryBillContent), C.int(nRequestID)))
}

func (api *CTradeApi) ReqBillConfirm(pReqBillConfirm *thost.CHSReqBillConfirmField, nRequestID int) int {
	return int(C.hs_td_ReqBillConfirm_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqBillConfirm), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryMargin(pReqQryMargin *thost.CHSReqQryMarginField, nRequestID int) int {
	return int(C.hs_td_ReqQryMargin_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryMargin), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryCommission(pReqQryCommission *thost.CHSReqQryCommissionField, nRequestID int) int {
	return int(C.hs_td_ReqQryCommission_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCommission), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryExchangeRate(pReqQryExchangeRate *thost.CHSReqQryExchangeRateField, nRequestID int) int {
	return int(C.hs_td_ReqQryExchangeRate_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryExchangeRate), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryPositionDetail(pReqQryPositionDetail *thost.CHSReqQryPositionDetailField, nRequestID int) int {
	return int(C.hs_td_ReqQryPositionDetail_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryPositionDetail), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQrySysConfig(pReqQrySysConfig *thost.CHSReqQrySysConfigField, nRequestID int) int {
	return int(C.hs_td_ReqQrySysConfig_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySysConfig), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryDepthMarketData(pReqQryDepthMarketData *thost.CHSReqQryDepthMarketDataField, nRequestID int) int {
	return int(C.hs_td_ReqQryDepthMarketData_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryDepthMarketData), C.int(nRequestID)))
}

func (api *CTradeApi) ReqFundTrans(pReqFundTransfer *thost.CHSReqFundTransField, nRequestID int) int {
	return int(C.hs_td_ReqFundTrans_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqFundTransfer), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryFundTrans(pReqQryFundTrans *thost.CHSReqQryFundTransField, nRequestID int) int {
	return int(C.hs_td_ReqQryFundTrans_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryFundTrans), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryClientNotice(pReqQryClientNotice *thost.CHSReqQryClientNoticeField, nRequestID int) int {
	return int(C.hs_td_ReqQryClientNotice_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryClientNotice), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOptUnderly(pReqQryOptUnderly *thost.CHSReqQryOptUnderlyField, nRequestID int) int {
	return int(C.hs_td_ReqQryOptUnderly_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptUnderly), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQrySecuDepthMarket(pReqQrySecuDepthMarket *thost.CHSReqQrySecuDepthMarketField, nRequestID int) int {
	return int(C.hs_td_ReqQrySecuDepthMarket_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySecuDepthMarket), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryHistOrder(pReqQryHistOrder *thost.CHSReqQryHistOrderField, nRequestID int) int {
	return int(C.hs_td_ReqQryHistOrder_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryHistOrder), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryHistTrade(pReqQryHistTrade *thost.CHSReqQryHistTradeField, nRequestID int) int {
	return int(C.hs_td_ReqQryHistTrade_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryHistTrade), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryWithdrawFund(pReqQryWithdrawFund *thost.CHSReqQryWithdrawFundField, nRequestID int) int {
	return int(C.hs_td_ReqQryWithdrawFund_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryWithdrawFund), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryCombInstrument(pReqQryCombInstrument *thost.CHSReqQryCombInstrumentField, nRequestID int) int {
	return int(C.hs_td_ReqQryCombInstrument_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryCombInstrument), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQrySeatID(pReqQrySeatID *thost.CHSReqQrySeatIDField, nRequestID int) int {
	return int(C.hs_td_ReqQrySeatID_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQrySeatID), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOptionSelfClose(pReqOptionSelfClose *thost.CHSReqOptionSelfCloseField, nRequestID int) int {
	return int(C.hs_td_ReqOptionSelfClose_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfClose), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOptionSelfCloseAction(pReqOptionSelfCloseAction *thost.CHSReqOptionSelfCloseActionField, nRequestID int) int {
	return int(C.hs_td_ReqOptionSelfCloseAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfCloseAction), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult *thost.CHSReqQryOptionSelfCloseResultField, nRequestID int) int {
	return int(C.hs_td_ReqQryOptionSelfCloseResult_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptionSelfCloseResult), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOptionSelfClose(pReqOptionSelfClose *thost.CHSReqOptionSelfCloseField, nRequestID int) int {
	return int(C.hs_td_ReqQryOptionSelfClose_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptionSelfClose), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOptQuoteInsert(pReqOptQuoteInsert *thost.CHSReqOptQuoteInsertField, nRequestID int) int {
	return int(C.hs_td_ReqOptQuoteInsert_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptQuoteInsert), C.int(nRequestID)))
}

func (api *CTradeApi) ReqOptQuoteAction(pReqOptQuoteAction *thost.CHSReqOptQuoteActionField, nRequestID int) int {
	return int(C.hs_td_ReqOptQuoteAction_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqOptQuoteAction), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOptQuote(pReqQryOptQuote *thost.CHSReqQryOptQuoteField, nRequestID int) int {
	return int(C.hs_td_ReqQryOptQuote_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptQuote), C.int(nRequestID)))
}

func (api *CTradeApi) ReqQryOptCombStrategy(pReqQryOptCombStrategy *thost.CHSReqQryOptCombStrategyField, nRequestID int) int {
	return int(C.hs_td_ReqQryOptCombStrategy_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pReqQryOptCombStrategy), C.int(nRequestID)))
}

//export hs_trade_callback_OnFrontConnected
func hs_trade_callback_OnFrontConnected(v uintptr) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnFrontConnected()
}

//export hs_trade_callback_OnFrontDisconnected
func hs_trade_callback_OnFrontDisconnected(v uintptr, nReason C.int) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnFrontDisconnected(int(nReason))
}

//export hs_trade_callback_OnRspAuthenticate
func hs_trade_callback_OnRspAuthenticate(v uintptr, pRspAuthenticateField *C.struct_CHSRspAuthenticateField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspAuthenticate((*thost.CHSRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspSubmitUserSystemInfo
func hs_trade_callback_OnRspSubmitUserSystemInfo(v uintptr, pRspUserSystemInfoField *C.struct_CHSRspUserSystemInfoField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspSubmitUserSystemInfo((*thost.CHSRspUserSystemInfoField)(unsafe.Pointer(pRspUserSystemInfoField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspUserLogin
func hs_trade_callback_OnRspUserLogin(v uintptr, pRspUserLoginField *C.struct_CHSRspUserLoginField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspUserLogin((*thost.CHSRspUserLoginField)(unsafe.Pointer(pRspUserLoginField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspUserPasswordUpdate
func hs_trade_callback_OnRspUserPasswordUpdate(v uintptr, pRspUserPasswordUpdateField *C.struct_CHSRspUserPasswordUpdateField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspUserPasswordUpdate((*thost.CHSRspUserPasswordUpdateField)(unsafe.Pointer(pRspUserPasswordUpdateField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorOrderInsert
func hs_trade_callback_OnRspErrorOrderInsert(v uintptr, pRspOrderInsertField *C.struct_CHSRspOrderInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorOrderInsert((*thost.CHSRspOrderInsertField)(unsafe.Pointer(pRspOrderInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspOrderAction
func hs_trade_callback_OnRspOrderAction(v uintptr, pRspOrderActionField *C.struct_CHSRspOrderActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspOrderAction((*thost.CHSRspOrderActionField)(unsafe.Pointer(pRspOrderActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorExerciseOrderInsert
func hs_trade_callback_OnRspErrorExerciseOrderInsert(v uintptr, pRspExerciseOrderInsertField *C.struct_CHSRspExerciseOrderInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorExerciseOrderInsert((*thost.CHSRspExerciseOrderInsertField)(unsafe.Pointer(pRspExerciseOrderInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspExerciseOrderAction
func hs_trade_callback_OnRspExerciseOrderAction(v uintptr, pRspExerciseOrderActionField *C.struct_CHSRspExerciseOrderActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspExerciseOrderAction((*thost.CHSRspExerciseOrderActionField)(unsafe.Pointer(pRspExerciseOrderActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorLockInsert
func hs_trade_callback_OnRspErrorLockInsert(v uintptr, pRspLockInsertField *C.struct_CHSRspLockInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorLockInsert((*thost.CHSRspLockInsertField)(unsafe.Pointer(pRspLockInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspForQuoteInsert
func hs_trade_callback_OnRspForQuoteInsert(v uintptr, pRspForQuoteInsertField *C.struct_CHSRspForQuoteInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspForQuoteInsert((*thost.CHSRspForQuoteInsertField)(unsafe.Pointer(pRspForQuoteInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorQuoteInsert
func hs_trade_callback_OnRspErrorQuoteInsert(v uintptr, pRspQuoteInsertField *C.struct_CHSRspQuoteInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorQuoteInsert((*thost.CHSRspQuoteInsertField)(unsafe.Pointer(pRspQuoteInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQuoteAction
func hs_trade_callback_OnRspQuoteAction(v uintptr, pRspQuoteActionField *C.struct_CHSRspQuoteActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQuoteAction((*thost.CHSRspQuoteActionField)(unsafe.Pointer(pRspQuoteActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorCombActionInsert
func hs_trade_callback_OnRspErrorCombActionInsert(v uintptr, pRspCombActionInsertField *C.struct_CHSRspCombActionInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorCombActionInsert((*thost.CHSRspCombActionInsertField)(unsafe.Pointer(pRspCombActionInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQueryMaxOrderVolume
func hs_trade_callback_OnRspQueryMaxOrderVolume(v uintptr, pRspQueryMaxOrderVolumeField *C.struct_CHSRspQueryMaxOrderVolumeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQueryMaxOrderVolume((*thost.CHSRspQueryMaxOrderVolumeField)(unsafe.Pointer(pRspQueryMaxOrderVolumeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryLockVolume
func hs_trade_callback_OnRspQryLockVolume(v uintptr, pRspQryLockVolumeField *C.struct_CHSRspQryLockVolumeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryLockVolume((*thost.CHSRspQryLockVolumeField)(unsafe.Pointer(pRspQryLockVolumeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQueryExerciseVolume
func hs_trade_callback_OnRspQueryExerciseVolume(v uintptr, pRspQueryExerciseVolumeField *C.struct_CHSRspQueryExerciseVolumeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQueryExerciseVolume((*thost.CHSRspQueryExerciseVolumeField)(unsafe.Pointer(pRspQueryExerciseVolumeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryCombVolume
func hs_trade_callback_OnRspQryCombVolume(v uintptr, pRspQryCombVolumeField *C.struct_CHSRspQryCombVolumeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryCombVolume((*thost.CHSRspQryCombVolumeField)(unsafe.Pointer(pRspQryCombVolumeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryPosition
func hs_trade_callback_OnRspQryPosition(v uintptr, pRspQryPositionField *C.struct_CHSRspQryPositionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryPosition((*thost.CHSRspQryPositionField)(unsafe.Pointer(pRspQryPositionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryTradingAccount
func hs_trade_callback_OnRspQryTradingAccount(v uintptr, pRspQryTradingAccountField *C.struct_CHSRspQryTradingAccountField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryTradingAccount((*thost.CHSRspQryTradingAccountField)(unsafe.Pointer(pRspQryTradingAccountField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOrder
func hs_trade_callback_OnRspQryOrder(v uintptr, pRspQryOrderField *C.struct_CHSOrderField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOrder((*thost.CHSOrderField)(unsafe.Pointer(pRspQryOrderField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryTrade
func hs_trade_callback_OnRspQryTrade(v uintptr, pRspQryTradeField *C.struct_CHSTradeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryTrade((*thost.CHSTradeField)(unsafe.Pointer(pRspQryTradeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryExercise
func hs_trade_callback_OnRspQryExercise(v uintptr, pRspQryExerciseField *C.struct_CHSExerciseField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryExercise((*thost.CHSExerciseField)(unsafe.Pointer(pRspQryExerciseField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryLock
func hs_trade_callback_OnRspQryLock(v uintptr, pRspQryLockField *C.struct_CHSLockField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryLock((*thost.CHSLockField)(unsafe.Pointer(pRspQryLockField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryCombAction
func hs_trade_callback_OnRspQryCombAction(v uintptr, pRspQryCombActionField *C.struct_CHSCombActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryCombAction((*thost.CHSCombActionField)(unsafe.Pointer(pRspQryCombActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryForQuote
func hs_trade_callback_OnRspQryForQuote(v uintptr, pRspQryForQuoteField *C.struct_CHSForQuoteField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryForQuote((*thost.CHSForQuoteField)(unsafe.Pointer(pRspQryForQuoteField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryQuote
func hs_trade_callback_OnRspQryQuote(v uintptr, pRspQryQuoteField *C.struct_CHSQuoteField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryQuote((*thost.CHSQuoteField)(unsafe.Pointer(pRspQryQuoteField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryPositionCombineDetail
func hs_trade_callback_OnRspQryPositionCombineDetail(v uintptr, pRspQryPositionCombineDetailField *C.struct_CHSRspQryPositionCombineDetailField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryPositionCombineDetail((*thost.CHSRspQryPositionCombineDetailField)(unsafe.Pointer(pRspQryPositionCombineDetailField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryInstrument
func hs_trade_callback_OnRspQryInstrument(v uintptr, pRspQryInstrumentField *C.struct_CHSRspQryInstrumentField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryInstrument((*thost.CHSRspQryInstrumentField)(unsafe.Pointer(pRspQryInstrumentField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryCoveredShort
func hs_trade_callback_OnRspQryCoveredShort(v uintptr, pRspQryCoveredShortField *C.struct_CHSRspQryCoveredShortField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryCoveredShort((*thost.CHSRspQryCoveredShortField)(unsafe.Pointer(pRspQryCoveredShortField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryExerciseAssign
func hs_trade_callback_OnRspQryExerciseAssign(v uintptr, pRspQryExerciseAssignField *C.struct_CHSRspQryExerciseAssignField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryExerciseAssign((*thost.CHSRspQryExerciseAssignField)(unsafe.Pointer(pRspQryExerciseAssignField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspTransfer
func hs_trade_callback_OnRspTransfer(v uintptr, pRspTransferField *C.struct_CHSRspTransferField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspTransfer((*thost.CHSRspTransferField)(unsafe.Pointer(pRspTransferField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryTransfer
func hs_trade_callback_OnRspQryTransfer(v uintptr, pRspQryTransferField *C.struct_CHSRspQryTransferField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryTransfer((*thost.CHSRspQryTransferField)(unsafe.Pointer(pRspQryTransferField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQueryBankBalance
func hs_trade_callback_OnRspQueryBankBalance(v uintptr, pRspQueryBankBalanceField *C.struct_CHSRspQueryBankBalanceField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQueryBankBalance((*thost.CHSRspQueryBankBalanceField)(unsafe.Pointer(pRspQueryBankBalanceField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQueryBankAccount
func hs_trade_callback_OnRspQueryBankAccount(v uintptr, pRspQueryBankAccountField *C.struct_CHSRspQueryBankAccountField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQueryBankAccount((*thost.CHSRspQueryBankAccountField)(unsafe.Pointer(pRspQueryBankAccountField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspMultiCentreFundTrans
func hs_trade_callback_OnRspMultiCentreFundTrans(v uintptr, pRspMultiCentreFundTransferField *C.struct_CHSRspMultiCentreFundTransField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspMultiCentreFundTrans((*thost.CHSRspMultiCentreFundTransField)(unsafe.Pointer(pRspMultiCentreFundTransferField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQueryBillContent
func hs_trade_callback_OnRspQueryBillContent(v uintptr, pRspQueryBillContentField *C.struct_CHSRspQueryBillContentField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQueryBillContent((*thost.CHSRspQueryBillContentField)(unsafe.Pointer(pRspQueryBillContentField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspBillConfirm
func hs_trade_callback_OnRspBillConfirm(v uintptr, pRspBillConfirmField *C.struct_CHSRspBillConfirmField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspBillConfirm((*thost.CHSRspBillConfirmField)(unsafe.Pointer(pRspBillConfirmField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryMargin
func hs_trade_callback_OnRspQryMargin(v uintptr, pRspQryMarginField *C.struct_CHSRspQryMarginField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryMargin((*thost.CHSRspQryMarginField)(unsafe.Pointer(pRspQryMarginField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryCommission
func hs_trade_callback_OnRspQryCommission(v uintptr, pRspQryCommissionField *C.struct_CHSRspQryCommissionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryCommission((*thost.CHSRspQryCommissionField)(unsafe.Pointer(pRspQryCommissionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryPositionDetail
func hs_trade_callback_OnRspQryPositionDetail(v uintptr, pRspQryPositionDetailField *C.struct_CHSRspQryPositionDetailField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryPositionDetail((*thost.CHSRspQryPositionDetailField)(unsafe.Pointer(pRspQryPositionDetailField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryExchangeRate
func hs_trade_callback_OnRspQryExchangeRate(v uintptr, pRspQryExchangeRateField *C.struct_CHSRspQryExchangeRateField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryExchangeRate((*thost.CHSRspQryExchangeRateField)(unsafe.Pointer(pRspQryExchangeRateField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQrySysConfig
func hs_trade_callback_OnRspQrySysConfig(v uintptr, pRspQrySysConfigField *C.struct_CHSRspQrySysConfigField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQrySysConfig((*thost.CHSRspQrySysConfigField)(unsafe.Pointer(pRspQrySysConfigField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryDepthMarketData
func hs_trade_callback_OnRspQryDepthMarketData(v uintptr, pRspQryDepthMarketDataField *C.struct_CHSDepthMarketDataField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryDepthMarketData((*thost.CHSDepthMarketDataField)(unsafe.Pointer(pRspQryDepthMarketDataField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspFundTrans
func hs_trade_callback_OnRspFundTrans(v uintptr, pRspFundTransferField *C.struct_CHSRspFundTransField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspFundTrans((*thost.CHSRspFundTransField)(unsafe.Pointer(pRspFundTransferField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryFundTrans
func hs_trade_callback_OnRspQryFundTrans(v uintptr, pRspQryFundTransField *C.struct_CHSRspQryFundTransField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryFundTrans((*thost.CHSRspQryFundTransField)(unsafe.Pointer(pRspQryFundTransField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryClientNotice
func hs_trade_callback_OnRspQryClientNotice(v uintptr, pRspQryClientNoticeField *C.struct_CHSClientNoticeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryClientNotice((*thost.CHSClientNoticeField)(unsafe.Pointer(pRspQryClientNoticeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOptUnderly
func hs_trade_callback_OnRspQryOptUnderly(v uintptr, pRspQryOptUnderlyField *C.struct_CHSRspQryOptUnderlyField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOptUnderly((*thost.CHSRspQryOptUnderlyField)(unsafe.Pointer(pRspQryOptUnderlyField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQrySecuDepthMarket
func hs_trade_callback_OnRspQrySecuDepthMarket(v uintptr, pRspQrySecuDepthMarketField *C.struct_CHSRspQrySecuDepthMarketField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQrySecuDepthMarket((*thost.CHSRspQrySecuDepthMarketField)(unsafe.Pointer(pRspQrySecuDepthMarketField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryHistOrder
func hs_trade_callback_OnRspQryHistOrder(v uintptr, pRspQryHistOrderField *C.struct_CHSOrderField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryHistOrder((*thost.CHSOrderField)(unsafe.Pointer(pRspQryHistOrderField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryHistTrade
func hs_trade_callback_OnRspQryHistTrade(v uintptr, pRspQryHistTradeField *C.struct_CHSTradeField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryHistTrade((*thost.CHSTradeField)(unsafe.Pointer(pRspQryHistTradeField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryWithdrawFund
func hs_trade_callback_OnRspQryWithdrawFund(v uintptr, pRspQryWithdrawFundField *C.struct_CHSRspQryWithdrawFundField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryWithdrawFund((*thost.CHSRspQryWithdrawFundField)(unsafe.Pointer(pRspQryWithdrawFundField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryCombInstrument
func hs_trade_callback_OnRspQryCombInstrument(v uintptr, pRspQryCombInstrumentField *C.struct_CHSRspQryCombInstrumentField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryCombInstrument((*thost.CHSRspQryCombInstrumentField)(unsafe.Pointer(pRspQryCombInstrumentField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQrySeatID
func hs_trade_callback_OnRspQrySeatID(v uintptr, pRspQrySeatIDField *C.struct_CHSRspQrySeatIDField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQrySeatID((*thost.CHSRspQrySeatIDField)(unsafe.Pointer(pRspQrySeatIDField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspOptionSelfClose
func hs_trade_callback_OnRspOptionSelfClose(v uintptr, pRspOptionSelfCloseField *C.struct_CHSRspOptionSelfCloseField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspOptionSelfClose((*thost.CHSRspOptionSelfCloseField)(unsafe.Pointer(pRspOptionSelfCloseField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspOptionSelfCloseAction
func hs_trade_callback_OnRspOptionSelfCloseAction(v uintptr, pRspOptionSelfCloseActionField *C.struct_CHSRspOptionSelfCloseActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspOptionSelfCloseAction((*thost.CHSRspOptionSelfCloseActionField)(unsafe.Pointer(pRspOptionSelfCloseActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOptionSelfCloseResult
func hs_trade_callback_OnRspQryOptionSelfCloseResult(v uintptr, pRspQryOptionSelfCloseResultField *C.struct_CHSRspQryOptionSelfCloseResultField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOptionSelfCloseResult((*thost.CHSRspQryOptionSelfCloseResultField)(unsafe.Pointer(pRspQryOptionSelfCloseResultField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOptionSelfClose
func hs_trade_callback_OnRspQryOptionSelfClose(v uintptr, pRspQryOptionSelfCloseField *C.struct_CHSOptionSelfCloseField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOptionSelfClose((*thost.CHSOptionSelfCloseField)(unsafe.Pointer(pRspQryOptionSelfCloseField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspErrorOptQuoteInsert
func hs_trade_callback_OnRspErrorOptQuoteInsert(v uintptr, pRspOptQuoteInsertField *C.struct_CHSRspOptQuoteInsertField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspErrorOptQuoteInsert((*thost.CHSRspOptQuoteInsertField)(unsafe.Pointer(pRspOptQuoteInsertField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspOptQuoteAction
func hs_trade_callback_OnRspOptQuoteAction(v uintptr, pRspOptQuoteActionField *C.struct_CHSRspOptQuoteActionField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspOptQuoteAction((*thost.CHSRspOptQuoteActionField)(unsafe.Pointer(pRspOptQuoteActionField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOptQuote
func hs_trade_callback_OnRspQryOptQuote(v uintptr, pRspQryOptQuoteField *C.struct_CHSOptQuoteField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOptQuote((*thost.CHSOptQuoteField)(unsafe.Pointer(pRspQryOptQuoteField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRspQryOptCombStrategy
func hs_trade_callback_OnRspQryOptCombStrategy(v uintptr, pRspQryOptCombStrategyField *C.struct_CHSRspQryOptCombStrategyField, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRspQryOptCombStrategy((*thost.CHSRspQryOptCombStrategyField)(unsafe.Pointer(pRspQryOptCombStrategyField)), (*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export hs_trade_callback_OnRtnTrade
func hs_trade_callback_OnRtnTrade(v uintptr, pRtnTradeField *C.struct_CHSTradeField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnTrade((*thost.CHSTradeField)(unsafe.Pointer(pRtnTradeField)))
}

//export hs_trade_callback_OnRtnOrder
func hs_trade_callback_OnRtnOrder(v uintptr, pRtnOrderField *C.struct_CHSOrderField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnOrder((*thost.CHSOrderField)(unsafe.Pointer(pRtnOrderField)))
}

//export hs_trade_callback_OnRtnExercise
func hs_trade_callback_OnRtnExercise(v uintptr, pRtnExerciseField *C.struct_CHSExerciseField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnExercise((*thost.CHSExerciseField)(unsafe.Pointer(pRtnExerciseField)))
}

//export hs_trade_callback_OnRtnCombAction
func hs_trade_callback_OnRtnCombAction(v uintptr, pRtnCombActionField *C.struct_CHSCombActionField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnCombAction((*thost.CHSCombActionField)(unsafe.Pointer(pRtnCombActionField)))
}

//export hs_trade_callback_OnRtnLock
func hs_trade_callback_OnRtnLock(v uintptr, pRtnLockField *C.struct_CHSLockField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnLock((*thost.CHSLockField)(unsafe.Pointer(pRtnLockField)))
}

//export hs_trade_callback_OnErrRtnOrderAction
func hs_trade_callback_OnErrRtnOrderAction(v uintptr, pRtnOrderActionField *C.struct_CHSOrderActionField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnErrRtnOrderAction((*thost.CHSOrderActionField)(unsafe.Pointer(pRtnOrderActionField)))
}

//export hs_trade_callback_OnRtnClientNotice
func hs_trade_callback_OnRtnClientNotice(v uintptr, pRtnClientNoticeField *C.struct_CHSClientNoticeField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnClientNotice((*thost.CHSClientNoticeField)(unsafe.Pointer(pRtnClientNoticeField)))
}

//export hs_trade_callback_OnRtnForQuote
func hs_trade_callback_OnRtnForQuote(v uintptr, pRtnForQuoteField *C.struct_CHSForQuoteField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnForQuote((*thost.CHSForQuoteField)(unsafe.Pointer(pRtnForQuoteField)))
}

//export hs_trade_callback_OnRtnQuote
func hs_trade_callback_OnRtnQuote(v uintptr, pRtnQuoteField *C.struct_CHSQuoteField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnQuote((*thost.CHSQuoteField)(unsafe.Pointer(pRtnQuoteField)))
}

//export hs_trade_callback_OnRtnExchangeStatus
func hs_trade_callback_OnRtnExchangeStatus(v uintptr, pRtnExchangeStatusField *C.struct_CHSExchangeStatusField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnExchangeStatus((*thost.CHSExchangeStatusField)(unsafe.Pointer(pRtnExchangeStatusField)))
}

//export hs_trade_callback_OnRtnProductStatus
func hs_trade_callback_OnRtnProductStatus(v uintptr, pRtnProductStatusField *C.struct_CHSProductStatusField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnProductStatus((*thost.CHSProductStatusField)(unsafe.Pointer(pRtnProductStatusField)))
}

//export hs_trade_callback_OnRtnOptionSelfClose
func hs_trade_callback_OnRtnOptionSelfClose(v uintptr, pRtnOptionSelfCloseField *C.struct_CHSOptionSelfCloseField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnOptionSelfClose((*thost.CHSOptionSelfCloseField)(unsafe.Pointer(pRtnOptionSelfCloseField)))
}

//export hs_trade_callback_OnRtnOptQuote
func hs_trade_callback_OnRtnOptQuote(v uintptr, pRtnOptQuoteField *C.struct_CHSOptQuoteField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnOptQuote((*thost.CHSOptQuoteField)(unsafe.Pointer(pRtnOptQuoteField)))
}

//export hs_trade_callback_OnRtnTransfer
func hs_trade_callback_OnRtnTransfer(v uintptr, pRtnTransferField *C.struct_CHSTransferField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnRtnTransfer((*thost.CHSTransferField)(unsafe.Pointer(pRtnTransferField)))
}

//export hs_trade_callback_OnErrRtnOptQuoteAction
func hs_trade_callback_OnErrRtnOptQuoteAction(v uintptr, pRtnOptQuoteActionField *C.struct_CHSOptQuoteActionField) {
	api := cgo.Handle(v).Value().(*CTradeApi)
	api.spi.OnErrRtnOptQuoteAction((*thost.CHSOptQuoteActionField)(unsafe.Pointer(pRtnOptQuoteActionField)))
}
