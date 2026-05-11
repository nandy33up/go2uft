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

uintptr_t hs_create_md_api_static(const char* flowPath);
void hs_release_md_api_static(uintptr_t api);
const char* hs_md_GetApiVersion_static(uintptr_t api);
int hs_md_Init_static(uintptr_t api, void* config);
int hs_md_Join_static(uintptr_t api);
int hs_md_RegisterFront_static(uintptr_t api, const char* frontAddress);
int hs_md_RegisterFensServer_static(uintptr_t api, const char* fensAddress, const char* accountID);
void hs_md_RegisterSpi_static(uintptr_t api, uintptr_t spi);
int hs_md_ReqDepthMarketDataSubscribe_static(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqDepthMarketDataCancel_static(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqForQuoteSubscribe_static(uintptr_t api, void* pReq, int count, int requestID);
int hs_md_ReqForQuoteCancel_static(uintptr_t api, void* pReq, int count, int requestID);
const char* hs_md_GetApiErrorMsg_static(uintptr_t api, int errorCode);
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

func CreateMdApi(options ...MdOption) thost.CHSMdApi {
	api := &CMdApi{
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

	api.apiPtr = uintptr(C.hs_create_md_api_static(cflowPath))
	return api
}

func (api *CMdApi) GetApiVersion() string {
	cs := C.hs_md_GetApiVersion_static(C.uintptr_t(api.apiPtr))
	return C.GoString(cs)
}

func (api *CMdApi) ReleaseApi() {
	C.hs_md_RegisterSpi_static(C.uintptr_t(api.apiPtr), C.uintptr_t(0))
	C.hs_release_md_api_static(C.uintptr_t(api.apiPtr))
	if api.handle != 0 {
		api.handle.Delete()
		api.handle = 0
	}
}

func (api *CMdApi) Init(pInitCfgField *thost.CHSInitConfigField, pExtraParam unsafe.Pointer) int {
	return int(C.hs_md_Init_static(C.uintptr_t(api.apiPtr), unsafe.Pointer(pInitCfgField)))
}

func (api *CMdApi) Join() int {
	return int(C.hs_md_Join_static(C.uintptr_t(api.apiPtr)))
}

func (api *CMdApi) RegisterFront(pszFrontAddress string) int {
	addr := C.CString(pszFrontAddress)
	defer C.free(unsafe.Pointer(addr))
	return int(C.hs_md_RegisterFront_static(C.uintptr_t(api.apiPtr), addr))
}

func (api *CMdApi) RegisterFensServer(pszFensAddress string, pszAccountID string) int {
	fensAddr := C.CString(pszFensAddress)
	defer C.free(unsafe.Pointer(fensAddr))
	accountID := C.CString(pszAccountID)
	defer C.free(unsafe.Pointer(accountID))
	return int(C.hs_md_RegisterFensServer_static(C.uintptr_t(api.apiPtr), fensAddr, accountID))
}

func (api *CMdApi) RegisterSpi(pSpi thost.CHSMdSpi) {
	api.spi = pSpi
	C.hs_md_RegisterSpi_static(C.uintptr_t(api.apiPtr), C.uintptr_t(api.handle))
}

func (api *CMdApi) ReqDepthMarketDataSubscribe(pReqDepthMarketDataSubscribe []thost.CHSReqDepthMarketDataField, nCount int, nRequestID int) int {
	if len(pReqDepthMarketDataSubscribe) == 0 {
		return -1
	}
	return int(C.hs_md_ReqDepthMarketDataSubscribe_static(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqDepthMarketDataSubscribe[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *CMdApi) ReqDepthMarketDataCancel(pReqDepthMarketDataCancel []thost.CHSReqDepthMarketDataField, nCount int, nRequestID int) int {
	if len(pReqDepthMarketDataCancel) == 0 {
		return -1
	}
	return int(C.hs_md_ReqDepthMarketDataCancel_static(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqDepthMarketDataCancel[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *CMdApi) ReqForQuoteSubscribe(pReqForQuoteSubscribe []thost.CHSReqForQuoteField, nCount int, nRequestID int) int {
	if len(pReqForQuoteSubscribe) == 0 {
		return -1
	}
	return int(C.hs_md_ReqForQuoteSubscribe_static(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqForQuoteSubscribe[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *CMdApi) ReqForQuoteCancel(pReqForQuoteCancel []thost.CHSReqForQuoteField, nCount int, nRequestID int) int {
	if len(pReqForQuoteCancel) == 0 {
		return -1
	}
	return int(C.hs_md_ReqForQuoteCancel_static(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqForQuoteCancel[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *CMdApi) GetApiErrorMsg(nErrorCode int) string {
	if nErrorCode < 0 {
		nErrorCode = -nErrorCode
	}
	cs := C.hs_md_GetApiErrorMsg_static(C.uintptr_t(api.apiPtr), C.int(nErrorCode))
	if cs == nil {
		return ""
	}
	return thost.BytesToGBK([]byte(C.GoString(cs)))
}

//export uft_quote_OnFrontConnected
func uft_quote_OnFrontConnected(v uintptr) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnFrontConnected()
}

//export uft_quote_OnFrontDisconnected
func uft_quote_OnFrontDisconnected(v uintptr, nReason C.int) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnFrontDisconnected(int(nReason))
}

//export uft_quote_OnRspDepthMarketDataSubscribe
func uft_quote_OnRspDepthMarketDataSubscribe(v uintptr, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRspDepthMarketDataSubscribe((*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export uft_quote_OnRspDepthMarketDataCancel
func uft_quote_OnRspDepthMarketDataCancel(v uintptr, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRspDepthMarketDataCancel((*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export uft_quote_OnRtnDepthMarketData
func uft_quote_OnRtnDepthMarketData(v uintptr, pDepthMarketData *C.struct_CHSDepthMarketDataField) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRtnDepthMarketData((*thost.CHSDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)))
}

//export uft_quote_OnRspForQuoteSubscribe
func uft_quote_OnRspForQuoteSubscribe(v uintptr, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRspForQuoteSubscribe((*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export uft_quote_OnRspForQuoteCancel
func uft_quote_OnRspForQuoteCancel(v uintptr, pRspInfo *C.struct_CHSRspInfoField, nRequestID C.int, bIsLast C.bool) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRspForQuoteCancel((*thost.CHSRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
}

//export uft_quote_OnRtnForQuote
func uft_quote_OnRtnForQuote(v uintptr, pForQuote *C.struct_CHSForQuoteField) {
	api := cgo.Handle(v).Value().(*CMdApi)
	api.spi.OnRtnForQuote((*thost.CHSForQuoteField)(unsafe.Pointer(pForQuote)))
}
