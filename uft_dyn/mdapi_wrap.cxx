#include "mdapi_wrap.h"

HSMdSpiWrap::HSMdSpiWrap(uintptr_t goHandle, const char* dllPath)
    : pUserApi(nullptr), goHandle(goHandle), dllHandle(nullptr) {
    if (dllPath == nullptr || strlen(dllPath) == 0) {
        fprintf(stderr, "DLL path is empty\n");
        return;
    }

    dllHandle = dlopen(dllPath, RTLD_NOW);
    if (dllHandle == nullptr) {
        fprintf(stderr, "[%s] dlopen error: %s\n", dllPath, dlerror());
        return;
    }

    typedef CHSMdApi* (*NewMdApiCreator)(const char*);
    NewMdApiCreator creator = (NewMdApiCreator)dlsym(dllHandle, "NewMdApi");
    if (creator == nullptr) {
        fprintf(stderr, "[%s] dlsym NewMdApi error: %s\n", dllPath, dlerror());
        dlclose(dllHandle);
        dllHandle = nullptr;
        return;
    }

    pUserApi = creator(nullptr);
}

HSMdSpiWrap::~HSMdSpiWrap() {
    if (pUserApi != nullptr) {
        pUserApi->RegisterSpi(nullptr);
        pUserApi->ReleaseApi();
        pUserApi = nullptr;
    }
    if (dllHandle != nullptr) {
        dlclose(dllHandle);
        dllHandle = nullptr;
    }
}

void HSMdSpiWrap::OnFrontConnected() {
    go_md_OnFrontConnected(goHandle, 0);
}

void HSMdSpiWrap::OnFrontDisconnected(int nResult) {
    go_md_OnFrontDisconnected(goHandle, nResult);
}

void HSMdSpiWrap::OnRspDepthMarketDataSubscribe(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_md_OnRspDepthMarketDataSubscribe(goHandle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRspDepthMarketDataCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_md_OnRspDepthMarketDataCancel(goHandle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRtnDepthMarketData(CHSDepthMarketDataField *pDepthMarketData) {
    go_md_OnRtnDepthMarketData(goHandle, pDepthMarketData);
}

void HSMdSpiWrap::OnRspForQuoteSubscribe(CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
    go_md_OnRspForQuoteSubscribe(goHandle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRspForQuoteCancel(CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    go_md_OnRspForQuoteCancel(goHandle, pRspInfo, nRequestID, bIsLast);
}

void HSMdSpiWrap::OnRtnForQuote(CHSForQuoteField *pForQuote) {
    go_md_OnRtnForQuote(goHandle, pForQuote);
}

HSMdApiWrap::HSMdApiWrap(uintptr_t goHandle, const char* dllPath)
    : pUserApi(nullptr), goHandle(goHandle), dllHandle(nullptr), pSpi(nullptr) {
    if (dllPath == nullptr || strlen(dllPath) == 0) {
        fprintf(stderr, "DLL path is empty\n");
        return;
    }

    dllHandle = dlopen(dllPath, RTLD_NOW);
    if (dllHandle == nullptr) {
        fprintf(stderr, "[%s] dlopen error: %s\n", dllPath, dlerror());
        return;
    }

    typedef CHSMdApi* (*NewMdApiCreator)(const char*);
    NewMdApiCreator creator = (NewMdApiCreator)dlsym(dllHandle, "NewMdApi");
    if (creator == nullptr) {
        fprintf(stderr, "[%s] dlsym NewMdApi error: %s\n", dllPath, dlerror());
        dlclose(dllHandle);
        dllHandle = nullptr;
        return;
    }

    pUserApi = creator(nullptr);
}

HSMdApiWrap::~HSMdApiWrap() {
    if (pSpi != nullptr) {
        pUserApi->RegisterSpi(nullptr);
        delete pSpi;
        pSpi = nullptr;
    }
    if (pUserApi != nullptr) {
        pUserApi->ReleaseApi();
        pUserApi = nullptr;
    }
    if (dllHandle != nullptr) {
        dlclose(dllHandle);
        dllHandle = nullptr;
    }
}

const char* HSMdApiWrap::GetApiVersion() {
    return nullptr;
}

void HSMdApiWrap::ReleaseApi() {
    if (pSpi != nullptr) {
        pUserApi->RegisterSpi(nullptr);
        delete pSpi;
        pSpi = nullptr;
    }
    if (pUserApi != nullptr) {
        pUserApi->ReleaseApi();
        pUserApi = nullptr;
    }
}

int HSMdApiWrap::Init(CHSInitConfigField *pInitConfigField, void *pExtraParam) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->Init(pInitConfigField, pExtraParam);
}

int HSMdApiWrap::Join() {
    if (pUserApi == nullptr) return -1;
    return pUserApi->Join();
}

int HSMdApiWrap::RegisterFront(const char *pszFrontAddress) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->RegisterFront(pszFrontAddress);
}

int HSMdApiWrap::RegisterFensServer(const char *pszFensAddress, const char *pszAccountID) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->RegisterFensServer(pszFensAddress, pszAccountID);
}

void HSMdApiWrap::RegisterSpi(HSMdSpiWrap* pSpi) {
    if (pUserApi == nullptr) return;
    this->pSpi = pSpi;
    pUserApi->RegisterSpi(pSpi);
}

int HSMdApiWrap::ReqDepthMarketDataSubscribe(CHSReqDepthMarketDataField pReq[], int nCount, int nRequestID) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->ReqDepthMarketDataSubscribe(pReq, nCount, nRequestID);
}

int HSMdApiWrap::ReqDepthMarketDataCancel(CHSReqDepthMarketDataField pReq[], int nCount, int nRequestID) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->ReqDepthMarketDataCancel(pReq, nCount, nRequestID);
}

int HSMdApiWrap::ReqForQuoteSubscribe(CHSReqForQuoteField pReq[], int nCount, int nRequestID) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->ReqForQuoteSubscribe(pReq, nCount, nRequestID);
}

int HSMdApiWrap::ReqForQuoteCancel(CHSReqForQuoteField pReq[], int nCount, int nRequestID) {
    if (pUserApi == nullptr) return -1;
    return pUserApi->ReqForQuoteCancel(pReq, nCount, nRequestID);
}

const char* HSMdApiWrap::GetApiErrorMsg(int nErrorCode) {
    return pUserApi->GetApiErrorMsg(nErrorCode);
}

uintptr_t hs_create_md_api(const char* dllPath) {
    return reinterpret_cast<uintptr_t>(new HSMdApiWrap(0, dllPath));
}

void hs_release_md_api(uintptr_t api) {
    delete reinterpret_cast<HSMdApiWrap*>(api);
}

const char* hs_md_GetApiVersion(uintptr_t api) {
    return reinterpret_cast<HSMdApiWrap*>(api)->GetApiVersion();
}

int hs_md_Init(uintptr_t api, void* config) {
    return reinterpret_cast<HSMdApiWrap*>(api)->Init(
        reinterpret_cast<CHSInitConfigField*>(config), nullptr);
}

int hs_md_Join(uintptr_t api) {
    return reinterpret_cast<HSMdApiWrap*>(api)->Join();
}

int hs_md_RegisterFront(uintptr_t api, const char* frontAddress) {
    return reinterpret_cast<HSMdApiWrap*>(api)->RegisterFront(frontAddress);
}

int hs_md_RegisterFensServer(uintptr_t api, const char* fensAddress, const char* accountID) {
    return reinterpret_cast<HSMdApiWrap*>(api)->RegisterFensServer(fensAddress, accountID);
}

void hs_md_RegisterSpi(uintptr_t api, uintptr_t spi) {
    reinterpret_cast<HSMdApiWrap*>(api)->RegisterSpi(
        reinterpret_cast<HSMdSpiWrap*>(spi));
}

int hs_md_ReqDepthMarketDataSubscribe(uintptr_t api, void* pReq, int count, int requestID) {
    return reinterpret_cast<HSMdApiWrap*>(api)->ReqDepthMarketDataSubscribe(
        reinterpret_cast<CHSReqDepthMarketDataField*>(pReq), count, requestID);
}

int hs_md_ReqDepthMarketDataCancel(uintptr_t api, void* pReq, int count, int requestID) {
    return reinterpret_cast<HSMdApiWrap*>(api)->ReqDepthMarketDataCancel(
        reinterpret_cast<CHSReqDepthMarketDataField*>(pReq), count, requestID);
}

int hs_md_ReqForQuoteSubscribe(uintptr_t api, void* pReq, int count, int requestID) {
    return reinterpret_cast<HSMdApiWrap*>(api)->ReqForQuoteSubscribe(
        reinterpret_cast<CHSReqForQuoteField*>(pReq), count, requestID);
}

int hs_md_ReqForQuoteCancel(uintptr_t api, void* pReq, int count, int requestID) {
    return reinterpret_cast<HSMdApiWrap*>(api)->ReqForQuoteCancel(
        reinterpret_cast<CHSReqForQuoteField*>(pReq), count, requestID);
}

const char* hs_md_GetApiErrorMsg(int errorCode) {
    return nullptr;
}
