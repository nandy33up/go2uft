#ifndef _HS_MD_API_WRAP_H_
#define _HS_MD_API_WRAP_H_

#include "HSMdApi.h"

class HSMdSpiWrap : public CHSMdSpi {
public:
    HSMdSpiWrap(uintptr_t handle) : m_handle(handle) {}
    
    virtual void OnFrontConnected();
    virtual void OnFrontDisconnected(int nResult);
    virtual void OnRspDepthMarketDataSubscribe(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspDepthMarketDataCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRtnDepthMarketData(CHSDepthMarketDataField *pDepthMarketData);
    virtual void OnRspForQuoteSubscribe(CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspForQuoteCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRtnForQuote(CHSForQuoteField *pForQuote);

private:
    uintptr_t m_handle;
};

class HSMdApiWrap {
public:
    HSMdApiWrap(CHSMdApi* api) : m_api(api) {}
    ~HSMdApiWrap() { if (m_spi) delete m_spi; }
    
    CHSMdApi* GetApi() { return m_api; }
    void SetSpi(HSMdSpiWrap* spi) { m_spi = spi; }
    
private:
    CHSMdApi* m_api;
    HSMdSpiWrap* m_spi;
};

extern "C" {
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
}

#endif