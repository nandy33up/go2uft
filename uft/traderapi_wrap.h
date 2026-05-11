#ifndef _HS_TRADE_API_WRAP_H_
#define _HS_TRADE_API_WRAP_H_

#include "HSTradeApi.h"

class HSTradeSpiWrap : public CHSTradeSpi {
public:
    HSTradeSpiWrap(uintptr_t handle) : m_handle(handle) {}
    
    virtual void OnFrontConnected();
    virtual void OnFrontDisconnected(int nResult);
    virtual void OnRspAuthenticate(CHSRspAuthenticateField *pRspAuthenticate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspSubmitUserSystemInfo(CHSRspUserSystemInfoField *pRspUserSystemInfo, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspUserLogin(CHSRspUserLoginField *pRspUserLogin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspUserPasswordUpdate(CHSRspUserPasswordUpdateField *pRspUserPasswordUpdate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorOrderInsert(CHSRspOrderInsertField *pRspOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspOrderAction(CHSRspOrderActionField *pRspOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorExerciseOrderInsert(CHSRspExerciseOrderInsertField *pRspExerciseOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspExerciseOrderAction(CHSRspExerciseOrderActionField *pRspExerciseOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorLockInsert(CHSRspLockInsertField *pRspLockInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspForQuoteInsert(CHSRspForQuoteInsertField *pRspForQuoteInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorQuoteInsert(CHSRspQuoteInsertField* pRspQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQuoteAction(CHSRspQuoteActionField* pRspQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorCombActionInsert(CHSRspCombActionInsertField *pRspCombActionInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQueryMaxOrderVolume(CHSRspQueryMaxOrderVolumeField *pRspQueryMaxOrderVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryLockVolume(CHSRspQryLockVolumeField *pRspQryLockVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQueryExerciseVolume(CHSRspQueryExerciseVolumeField *pRspQueryExerciseVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryCombVolume(CHSRspQryCombVolumeField *pRspQryCombVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryPosition(CHSRspQryPositionField *pRspQryPosition, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryTradingAccount(CHSRspQryTradingAccountField *pRspQryTradingAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOrder(CHSOrderField *pRspQryOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryTrade(CHSTradeField *pRspQryTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryExercise(CHSExerciseField *pRspQryExercise, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryLock(CHSLockField *pRspQryLock, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryCombAction(CHSCombActionField *pRspQryCombAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryForQuote(CHSForQuoteField* pRspQryForQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryQuote(CHSQuoteField* pRspQryQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryPositionCombineDetail(CHSRspQryPositionCombineDetailField *pRspQryPositionCombineDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryInstrument(CHSRspQryInstrumentField *pRspQryInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryCoveredShort(CHSRspQryCoveredShortField *pRspQryCoveredShort, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryExerciseAssign(CHSRspQryExerciseAssignField *pRspQryExerciseAssign, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspTransfer(CHSRspTransferField *pRspTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryTransfer(CHSRspQryTransferField *pRspQryTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQueryBankBalance(CHSRspQueryBankBalanceField *pRspQueryBankBalance, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQueryBankAccount(CHSRspQueryBankAccountField *pRspQueryBankAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspMultiCentreFundTrans(CHSRspMultiCentreFundTransField *pRspMultiCentreFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQueryBillContent(CHSRspQueryBillContentField *pRspQueryBillContent, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspBillConfirm(CHSRspBillConfirmField *pRspBillConfirm, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryMargin(CHSRspQryMarginField *pRspQryMargin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryCommission(CHSRspQryCommissionField *pRspQryCommission, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryPositionDetail(CHSRspQryPositionDetailField *pRspQryPositionDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryExchangeRate(CHSRspQryExchangeRateField *pRspQryExchangeRate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQrySysConfig(CHSRspQrySysConfigField *pRspQrySysConfig, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryDepthMarketData(CHSDepthMarketDataField *pRspQryDepthMarketData, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspFundTrans(CHSRspFundTransField *pRspFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryFundTrans(CHSRspQryFundTransField *pRspQryFundTrans, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryClientNotice(CHSClientNoticeField *pRspQryClientNotice, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOptUnderly(CHSRspQryOptUnderlyField *pRspQryOptUnderly, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQrySecuDepthMarket(CHSRspQrySecuDepthMarketField *pRspQrySecuDepthMarket, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryHistOrder(CHSOrderField *pRspQryHistOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryHistTrade(CHSTradeField *pRspQryHistTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryWithdrawFund(CHSRspQryWithdrawFundField *pRspQryWithdrawFund, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryCombInstrument(CHSRspQryCombInstrumentField *pRspQryCombInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQrySeatID(CHSRspQrySeatIDField *pRspQrySeatID, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspOptionSelfClose(CHSRspOptionSelfCloseField *pReqOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspOptionSelfCloseAction(CHSRspOptionSelfCloseActionField *pReqOptionSelfCloseAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOptionSelfCloseResult(CHSRspQryOptionSelfCloseResultField *pReqQryOptionSelfCloseResult, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOptionSelfClose(CHSOptionSelfCloseField *pRspQryOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspErrorOptQuoteInsert(CHSRspOptQuoteInsertField *pRspOptQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspOptQuoteAction(CHSRspOptQuoteActionField *pRspOptQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOptQuote(CHSOptQuoteField *pRspQryOptQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRspQryOptCombStrategy(CHSRspQryOptCombStrategyField *pRspQryOptCombStrategy, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
    virtual void OnRtnTrade(CHSTradeField *pRtnTrade);
    virtual void OnRtnOrder(CHSOrderField *pRtnOrder);
    virtual void OnRtnExercise(CHSExerciseField *pRtnExercise);
    virtual void OnRtnCombAction(CHSCombActionField *pRtnCombAction);
    virtual void OnRtnLock(CHSLockField *pRtnLock);
    virtual void OnErrRtnOrderAction(CHSOrderActionField *pRtnOrderAction);
    virtual void OnRtnClientNotice(CHSClientNoticeField *pRtnClientNotice);
    virtual void OnRtnForQuote(CHSForQuoteField* pRtnForQuote);
    virtual void OnRtnQuote(CHSQuoteField* pRtnQuote);
    virtual void OnRtnExchangeStatus(CHSExchangeStatusField* pRtnExchangeStatus);
    virtual void OnRtnProductStatus(CHSProductStatusField* pRtnProductStatus);
    virtual void OnRtnOptionSelfClose(CHSOptionSelfCloseField *pRtnOptionSelfClose);
    virtual void OnRtnOptQuote(CHSOptQuoteField* pRtnOptQuote);
    virtual void OnRtnTransfer(CHSTransferField* pRtnTransfer);
    virtual void OnErrRtnOptQuoteAction(CHSOptQuoteActionField* pRtnOptQuoteAction);

private:
    uintptr_t m_handle;
};

class HSTradeApiWrap {
public:
    HSTradeApiWrap(CHSTradeApi* api) : m_api(api) {}
    ~HSTradeApiWrap() { if (m_spi) delete m_spi; }
    
    CHSTradeApi* GetApi() { return m_api; }
    void SetSpi(HSTradeSpiWrap* spi) { m_spi = spi; }
    
private:
    CHSTradeApi* m_api;
    HSTradeSpiWrap* m_spi;
};

extern "C" {
    uintptr_t hs_create_td_api_static(const char* flowPath);
    void hs_release_td_api_static(uintptr_t api);
    const char* hs_td_GetApiVersion_static(uintptr_t api);
    int hs_td_Init_static(uintptr_t api, void* config);
    int hs_td_Join_static(uintptr_t api);
    int hs_td_RegisterSubModel_static(uintptr_t api, int eSubType);
    int hs_td_RegisterFront_static(uintptr_t api, const char* frontAddress);
    int hs_td_RegisterFensServer_static(uintptr_t api, const char* fensAddress, const char* accountID);
    void hs_td_RegisterSpi_static(uintptr_t api, uintptr_t spi);
    const char* hs_td_GetApiErrorMsg_static(uintptr_t api, int errorCode);
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
    int hs_td_ReqQryOptionSelfClose_static(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID);
    int hs_td_ReqOptQuoteInsert_static(uintptr_t api, void* pReqOptQuoteInsertField, int nRequestID);
    int hs_td_ReqOptQuoteAction_static(uintptr_t api, void* pReqOptQuoteActionField, int nRequestID);
    int hs_td_ReqQryOptQuote_static(uintptr_t api, void* pReqQryOptQuoteField, int nRequestID);
    int hs_td_ReqQryOptCombStrategy_static(uintptr_t api, void* pReqQryOptCombStrategyField, int nRequestID);
}

#endif