#include "mdapi_wrap.h"

extern "C" {
    extern void uft_quote_OnFrontConnected(uintptr_t handle);
    extern void uft_quote_OnFrontDisconnected(uintptr_t handle, int reason);
    extern void uft_quote_OnRspDepthMarketDataSubscribe(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
    extern void uft_quote_OnRspDepthMarketDataCancel(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
    extern void uft_quote_OnRtnDepthMarketData(uintptr_t handle, void* pDepthMarketData);
    extern void uft_quote_OnRspForQuoteSubscribe(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
    extern void uft_quote_OnRspForQuoteCancel(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
    extern void uft_quote_OnRtnForQuote(uintptr_t handle, void* pForQuote);
}

void HSMdSpiWrap::OnFrontConnected() {
    uft_quote_OnFrontConnected(m_handle);
}

void HSMdSpiWrap::OnFrontDisconnected(int nResult) {
    uft_quote_OnFrontDisconnected(m_handle, nResult);
}

void HSMdSpiWrap::OnRspDepthMarketDataSubscribe(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    uft_quote_OnRspDepthMarketDataSubscribe(m_handle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRspDepthMarketDataCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    uft_quote_OnRspDepthMarketDataCancel(m_handle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRtnDepthMarketData(CHSDepthMarketDataField *pDepthMarketData) {
    uft_quote_OnRtnDepthMarketData(m_handle, pDepthMarketData);
}

void HSMdSpiWrap::OnRspForQuoteSubscribe(CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    uft_quote_OnRspForQuoteSubscribe(m_handle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRspForQuoteCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    uft_quote_OnRspForQuoteCancel(m_handle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRtnForQuote(CHSForQuoteField *pForQuote) {
    uft_quote_OnRtnForQuote(m_handle, pForQuote);
}

extern "C" {

uintptr_t hs_create_md_api_static(const char* flowPath) {
    CHSMdApi* api = NewMdApi(flowPath);
    if (!api) return 0;
    HSMdApiWrap* wrap = new HSMdApiWrap(api);
    return reinterpret_cast<uintptr_t>(wrap);
}

void hs_release_md_api_static(uintptr_t api) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (wrap) {
        wrap->GetApi()->ReleaseApi();
        delete wrap;
    }
}

const char* hs_md_GetApiVersion_static(uintptr_t api) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return "";
    return GetMdApiVersion();
}

int hs_md_Init_static(uintptr_t api, void* config) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->Init(reinterpret_cast<CHSInitConfigField*>(config));
}

int hs_md_Join_static(uintptr_t api) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->Join();
}

int hs_md_RegisterFront_static(uintptr_t api, const char* frontAddress) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->RegisterFront(frontAddress);
}

int hs_md_RegisterFensServer_static(uintptr_t api, const char* fensAddress, const char* accountID) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->RegisterFensServer(fensAddress, accountID);
}

void hs_md_RegisterSpi_static(uintptr_t api, uintptr_t handle) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return;
    HSMdSpiWrap* spi = new HSMdSpiWrap(handle);
    wrap->SetSpi(spi);
    wrap->GetApi()->RegisterSpi(spi);
}

int hs_md_ReqDepthMarketDataSubscribe_static(uintptr_t api, void* pReq, int count, int requestID) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqDepthMarketDataSubscribe(
        reinterpret_cast<CHSReqDepthMarketDataField*>(pReq), count, requestID);
}

int hs_md_ReqDepthMarketDataCancel_static(uintptr_t api, void* pReq, int count, int requestID) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqDepthMarketDataCancel(
        reinterpret_cast<CHSReqDepthMarketDataField*>(pReq), count, requestID);
}

int hs_md_ReqForQuoteSubscribe_static(uintptr_t api, void* pReq, int count, int requestID) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqForQuoteSubscribe(
        reinterpret_cast<CHSReqForQuoteField*>(pReq), count, requestID);
}

int hs_md_ReqForQuoteCancel_static(uintptr_t api, void* pReq, int count, int requestID) {
    HSMdApiWrap* wrap = reinterpret_cast<HSMdApiWrap*>(api);
    if (!wrap) return -1;
    return wrap->GetApi()->ReqForQuoteCancel(
        reinterpret_cast<CHSReqForQuoteField*>(pReq), count, requestID);
}

const char* hs_md_GetApiErrorMsg_static(int errorCode) {
    return nullptr;
}

}