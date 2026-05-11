#ifndef _HS_MDAPI_WRAP_H_
#define _HS_MDAPI_WRAP_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void go_md_OnFrontConnected(uintptr_t handle, int _dummy);
void go_md_OnFrontDisconnected(uintptr_t handle, int reason);
void go_md_OnRspDepthMarketDataSubscribe(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
void go_md_OnRspDepthMarketDataCancel(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
void go_md_OnRtnDepthMarketData(uintptr_t handle, void* pDepthMarketData);
void go_md_OnRspForQuoteSubscribe(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
void go_md_OnRspForQuoteCancel(uintptr_t handle, void* pRspInfo, int requestID, bool isLast);
void go_md_OnRtnForQuote(uintptr_t handle, void* pForQuote);

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

#ifdef __cplusplus
}
#endif

#ifdef __cplusplus

#include <dlfcn.h>
#include <stdio.h>
#include <string.h>

#include "HSMdApi.h"
#include "HSStruct.h"
#include "HSDataType.h"

class HSMdSpiWrap : public CHSMdSpi {
public:
    HSMdSpiWrap(uintptr_t goHandle, const char* dllPath);
    virtual ~HSMdSpiWrap();

    virtual void OnFrontConnected() override;
    virtual void OnFrontDisconnected(int nResult) override;
    virtual void OnRspDepthMarketDataSubscribe(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) override;
    virtual void OnRspDepthMarketDataCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) override;
    virtual void OnRtnDepthMarketData(CHSDepthMarketDataField *pDepthMarketData) override;
    virtual void OnRspForQuoteSubscribe(CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override;
    virtual void OnRspForQuoteCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) override;
    virtual void OnRtnForQuote(CHSForQuoteField *pForQuote) override;

    CHSMdApi* GetApi() { return pUserApi; }

private:
    CHSMdApi* pUserApi;
    uintptr_t goHandle;
    void* dllHandle;
};

class HSMdApiWrap {
public:
    HSMdApiWrap(uintptr_t goHandle, const char* dllPath);
    ~HSMdApiWrap();

    const char* GetApiVersion();
    void ReleaseApi();
    int Init(CHSInitConfigField *pInitConfigField, void *pExtraParam);
    int Join();
    int RegisterFront(const char *pszFrontAddress);
    int RegisterFensServer(const char *pszFensAddress, const char *pszAccountID);
    void RegisterSpi(HSMdSpiWrap* pSpi);
    int ReqDepthMarketDataSubscribe(CHSReqDepthMarketDataField pReq[], int nCount, int nRequestID);
    int ReqDepthMarketDataCancel(CHSReqDepthMarketDataField pReq[], int nCount, int nRequestID);
    int ReqForQuoteSubscribe(CHSReqForQuoteField pReq[], int nCount, int nRequestID);
    int ReqForQuoteCancel(CHSReqForQuoteField pReq[], int nCount, int nRequestID);
    const char* GetApiErrorMsg(int nErrorCode);

private:
    CHSMdApi* pUserApi;
    uintptr_t goHandle;
    void* dllHandle;
    HSMdSpiWrap* pSpi;
};

#endif

#endif
