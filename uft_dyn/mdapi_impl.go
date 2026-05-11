package uft_dyn

/*
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
*/
import "C"

import (
	"sync"
	"unsafe"

	"github.com/nandy33up/go2uft/thost"
)

type MdOption func(api *MdApi)

type MdApi struct {
	apiPtr  uintptr
	dllPath string
	spi     CHSMdSpi
	mu      sync.RWMutex
}

type CHSMdSpi interface {
	thost.CHSMdSpi
}

var (
	mdSpiMap     = make(map[uintptr]CHSMdSpi)
	mdSpiMapLock sync.RWMutex
)

func DynamicLibPath(path string) MdOption {
	return func(api *MdApi) {
		api.dllPath = path
	}
}

func CreateMdApi(options ...MdOption) *MdApi {
	api := &MdApi{}
	for _, opt := range options {
		opt(api)
	}
	api.apiPtr = uintptr(C.hs_create_md_api(cString(api.dllPath)))
	return api
}

func (api *MdApi) ReleaseApi() {
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.apiPtr != 0 {
		C.hs_release_md_api(C.uintptr_t(api.apiPtr))
		api.apiPtr = 0
	}
}

func (api *MdApi) GetApiVersion() string {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return ""
	}
	cs := C.hs_md_GetApiVersion(C.uintptr_t(api.apiPtr))
	if cs == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func (api *MdApi) Init(pInitCfgField *thost.CHSInitConfigField, pExtraParam unsafe.Pointer) int {
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_md_Init(C.uintptr_t(api.apiPtr), unsafe.Pointer(pInitCfgField)))
}

func (api *MdApi) Join() int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_md_Join(C.uintptr_t(api.apiPtr)))
}

func (api *MdApi) RegisterFront(pszFrontAddress string) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_md_RegisterFront(C.uintptr_t(api.apiPtr), cString(pszFrontAddress)))
}

func (api *MdApi) RegisterFensServer(pszFensAddress string, pszAccountID string) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	return int(C.hs_md_RegisterFensServer(C.uintptr_t(api.apiPtr), cString(pszFensAddress), cString(pszAccountID)))
}

func (api *MdApi) RegisterSpi(pSpi CHSMdSpi) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.spi = pSpi
	if api.apiPtr != 0 && pSpi != nil {
		mdSpiMapLock.Lock()
		mdSpiMap[api.apiPtr] = pSpi
		mdSpiMapLock.Unlock()
		C.hs_md_RegisterSpi(C.uintptr_t(api.apiPtr), C.uintptr_t(api.apiPtr))
	}
}

func (api *MdApi) ReqDepthMarketDataSubscribe(pReqDepthMarketDataSubscribe []thost.CHSReqDepthMarketDataField, nCount int, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	if len(pReqDepthMarketDataSubscribe) == 0 {
		return -1
	}
	return int(C.hs_md_ReqDepthMarketDataSubscribe(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqDepthMarketDataSubscribe[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *MdApi) ReqDepthMarketDataCancel(pReqDepthMarketDataCancel []thost.CHSReqDepthMarketDataField, nCount int, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	if len(pReqDepthMarketDataCancel) == 0 {
		return -1
	}
	return int(C.hs_md_ReqDepthMarketDataCancel(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqDepthMarketDataCancel[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *MdApi) ReqForQuoteSubscribe(pReqForQuoteSubscribe []thost.CHSReqForQuoteField, nCount int, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	if len(pReqForQuoteSubscribe) == 0 {
		return -1
	}
	return int(C.hs_md_ReqForQuoteSubscribe(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqForQuoteSubscribe[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *MdApi) ReqForQuoteCancel(pReqForQuoteCancel []thost.CHSReqForQuoteField, nCount int, nRequestID int) int {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.apiPtr == 0 {
		return -1
	}
	if len(pReqForQuoteCancel) == 0 {
		return -1
	}
	return int(C.hs_md_ReqForQuoteCancel(
		C.uintptr_t(api.apiPtr),
		unsafe.Pointer(&pReqForQuoteCancel[0]),
		C.int(nCount),
		C.int(nRequestID)))
}

func (api *MdApi) GetApiErrorMsg(nErrorCode int) string {
	cs := C.hs_md_GetApiErrorMsg(C.int(nErrorCode))
	if cs == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func getMdSpi(handle uintptr) CHSMdSpi {
	mdSpiMapLock.RLock()
	defer mdSpiMapLock.RUnlock()
	return mdSpiMap[handle]
}

func cString(s string) *C.char {
	if s == "" {
		return nil
	}
	return C.CString(s)
}

//export go_md_OnFrontConnected
func go_md_OnFrontConnected(v C.uintptr_t, _dummy C.int) {
	spi := getMdSpi(uintptr(v))
	if spi == nil {
		return
	}
	spi.OnFrontConnected()
}

//export go_md_OnFrontDisconnected
func go_md_OnFrontDisconnected(v uintptr, p0 C.int) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnFrontDisconnected(int(p0))
}

//export go_md_OnRspDepthMarketDataSubscribe
func go_md_OnRspDepthMarketDataSubscribe(v uintptr, p0 unsafe.Pointer, p1 C.int, p2 C.bool) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspDepthMarketDataSubscribe((*thost.CHSRspInfoField)(p0), int(p1), bool(p2))
}

//export go_md_OnRspDepthMarketDataCancel
func go_md_OnRspDepthMarketDataCancel(v uintptr, p0 unsafe.Pointer, p1 C.int, p2 C.bool) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspDepthMarketDataCancel((*thost.CHSRspInfoField)(p0), int(p1), bool(p2))
}

//export go_md_OnRtnDepthMarketData
func go_md_OnRtnDepthMarketData(v uintptr, p0 unsafe.Pointer) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnDepthMarketData((*thost.CHSDepthMarketDataField)(p0))
}

//export go_md_OnRspForQuoteSubscribe
func go_md_OnRspForQuoteSubscribe(v uintptr, p0 unsafe.Pointer, p1 C.int, p2 C.bool) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspForQuoteSubscribe((*thost.CHSRspInfoField)(p0), int(p1), bool(p2))
}

//export go_md_OnRspForQuoteCancel
func go_md_OnRspForQuoteCancel(v uintptr, p0 unsafe.Pointer, p1 C.int, p2 C.bool) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRspForQuoteCancel((*thost.CHSRspInfoField)(p0), int(p1), bool(p2))
}

//export go_md_OnRtnForQuote
func go_md_OnRtnForQuote(v uintptr, p0 unsafe.Pointer) {
	spi := getMdSpi(v)
	if spi == nil {
		return
	}
	spi.OnRtnForQuote((*thost.CHSForQuoteField)(p0))
}
