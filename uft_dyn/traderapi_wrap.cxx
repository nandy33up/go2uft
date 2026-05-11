#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "HSTradeApi.h"
#include "HSStruct.h"
#include "HSDataType.h"

#include "traderapi_wrap.h"

typedef CHSTradeApi* (*NewTradeApiCreator)(const char*);
typedef const char* (*GetTradeApiVersionCreator)();

extern "C" {

void go_td_OnFrontConnected(uintptr_t handle, int _dummy);
void go_td_OnFrontDisconnected(uintptr_t handle, int reason);
void go_td_OnRspAuthenticate(uintptr_t handle, void* pRspAuthenticateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspSubmitUserSystemInfo(uintptr_t handle, void* pRspUserSystemInfoField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspUserLogin(uintptr_t handle, void* pRspUserLoginField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspUserPasswordUpdate(uintptr_t handle, void* pRspUserPasswordUpdateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorOrderInsert(uintptr_t handle, void* pRspOrderInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOrderAction(uintptr_t handle, void* pRspOrderActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorExerciseOrderInsert(uintptr_t handle, void* pRspExerciseOrderInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspExerciseOrderAction(uintptr_t handle, void* pRspExerciseOrderActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorLockInsert(uintptr_t handle, void* pRspLockInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspForQuoteInsert(uintptr_t handle, void* pRspForQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorQuoteInsert(uintptr_t handle, void* pRspQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQuoteAction(uintptr_t handle, void* pRspQuoteActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorCombActionInsert(uintptr_t handle, void* pRspCombActionInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryMaxOrderVolume(uintptr_t handle, void* pRspQueryMaxOrderVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryLockVolume(uintptr_t handle, void* pRspQryLockVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryExerciseVolume(uintptr_t handle, void* pRspQueryExerciseVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombVolume(uintptr_t handle, void* pRspQryCombVolumeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPosition(uintptr_t handle, void* pRspQryPositionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTradingAccount(uintptr_t handle, void* pRspQryTradingAccountField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOrder(uintptr_t handle, void* pRspQryOrderField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTrade(uintptr_t handle, void* pRspQryTradeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExercise(uintptr_t handle, void* pRspQryExerciseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryLock(uintptr_t handle, void* pRspQryLockField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombAction(uintptr_t handle, void* pRspQryCombActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryForQuote(uintptr_t handle, void* pRspQryForQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryQuote(uintptr_t handle, void* pRspQryQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPositionCombineDetail(uintptr_t handle, void* pRspQryPositionCombineDetailField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryInstrument(uintptr_t handle, void* pRspQryInstrumentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCoveredShort(uintptr_t handle, void* pRspQryCoveredShortField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExerciseAssign(uintptr_t handle, void* pRspQryExerciseAssignField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspTransfer(uintptr_t handle, void* pRspTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryTransfer(uintptr_t handle, void* pRspQryTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBankBalance(uintptr_t handle, void* pRspQueryBankBalanceField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBankAccount(uintptr_t handle, void* pRspQueryBankAccountField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspMultiCentreFundTrans(uintptr_t handle, void* pRspMultiCentreFundTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQueryBillContent(uintptr_t handle, void* pRspQueryBillContentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspBillConfirm(uintptr_t handle, void* pRspBillConfirmField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryMargin(uintptr_t handle, void* pRspQryMarginField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCommission(uintptr_t handle, void* pRspQryCommissionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryPositionDetail(uintptr_t handle, void* pRspQryPositionDetailField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryExchangeRate(uintptr_t handle, void* pRspQryExchangeRateField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySysConfig(uintptr_t handle, void* pRspQrySysConfigField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryDepthMarketData(uintptr_t handle, void* pRspQryDepthMarketDataField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspFundTrans(uintptr_t handle, void* pRspFundTransferField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryFundTrans(uintptr_t handle, void* pRspQryFundTransField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryClientNotice(uintptr_t handle, void* pRspQryClientNoticeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptUnderly(uintptr_t handle, void* pRspQryOptUnderlyField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySecuDepthMarket(uintptr_t handle, void* pRspQrySecuDepthMarketField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryHistOrder(uintptr_t handle, void* pRspQryHistOrderField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryHistTrade(uintptr_t handle, void* pRspQryHistTradeField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryWithdrawFund(uintptr_t handle, void* pRspQryWithdrawFundField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryCombInstrument(uintptr_t handle, void* pRspQryCombInstrumentField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQrySeatID(uintptr_t handle, void* pRspQrySeatIDField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptionSelfClose(uintptr_t handle, void* pRspOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptionSelfCloseAction(uintptr_t handle, void* pRspOptionSelfCloseActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptionSelfCloseResult(uintptr_t handle, void* pRspQryOptionSelfCloseResultField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptionSelfClose(uintptr_t handle, void* pRspQryOptionSelfCloseField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspErrorOptQuoteInsert(uintptr_t handle, void* pRspOptQuoteInsertField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspOptQuoteAction(uintptr_t handle, void* pRspOptQuoteActionField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptQuote(uintptr_t handle, void* pRspQryOptQuoteField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRspQryOptCombStrategy(uintptr_t handle, void* pRspQryOptCombStrategyField, void* pRspInfo, int requestID, bool isLast);
void go_td_OnRtnTrade(uintptr_t handle, void* pRtnTradeField);
void go_td_OnRtnOrder(uintptr_t handle, void* pRtnOrderField);
void go_td_OnRtnExercise(uintptr_t handle, void* pRtnExerciseField);
void go_td_OnRtnCombAction(uintptr_t handle, void* pRtnCombActionField);
void go_td_OnRtnLock(uintptr_t handle, void* pRtnLockField);
void go_td_OnErrRtnOrderAction(uintptr_t handle, void* pRtnOrderActionField);
void go_td_OnRtnClientNotice(uintptr_t handle, void* pRtnClientNoticeField);
void go_td_OnRtnForQuote(uintptr_t handle, void* pRtnForQuoteField);
void go_td_OnRtnQuote(uintptr_t handle, void* pRtnQuoteField);
void go_td_OnRtnExchangeStatus(uintptr_t handle, void* pRtnExchangeStatusField);
void go_td_OnRtnProductStatus(uintptr_t handle, void* pRtnProductStatusField);
void go_td_OnRtnOptionSelfClose(uintptr_t handle, void* pRtnOptionSelfCloseField);
void go_td_OnRtnOptQuote(uintptr_t handle, void* pRtnOptQuoteField);
void go_td_OnRtnTransfer(uintptr_t handle, void* pRtnTransferField);
void go_td_OnErrRtnOptQuoteAction(uintptr_t handle, void* pRtnOptQuoteActionField);

}

class HSTradeSpiWrap : public CHSTradeSpi {
public:
    HSTradeSpiWrap(uintptr_t goHandle) : goHandle(goHandle) {}
    virtual ~HSTradeSpiWrap() {}

    virtual void OnFrontConnected() {
        go_td_OnFrontConnected(goHandle, 0);
    }

    virtual void OnFrontDisconnected(int nResult) {
        go_td_OnFrontDisconnected(goHandle, nResult);
    }

    virtual void OnRspAuthenticate(CHSRspAuthenticateField *pRspAuthenticate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspAuthenticate(goHandle, pRspAuthenticate, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspSubmitUserSystemInfo(CHSRspUserSystemInfoField *pRspUserSystemInfo, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspSubmitUserSystemInfo(goHandle, pRspUserSystemInfo, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspUserLogin(CHSRspUserLoginField *pRspUserLogin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspUserLogin(goHandle, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspUserPasswordUpdate(CHSRspUserPasswordUpdateField *pRspUserPasswordUpdate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspUserPasswordUpdate(goHandle, pRspUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorOrderInsert(CHSRspOrderInsertField *pRspOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorOrderInsert(goHandle, pRspOrderInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspOrderAction(CHSRspOrderActionField *pRspOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspOrderAction(goHandle, pRspOrderAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorExerciseOrderInsert(CHSRspExerciseOrderInsertField *pRspExerciseOrderInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorExerciseOrderInsert(goHandle, pRspExerciseOrderInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspExerciseOrderAction(CHSRspExerciseOrderActionField *pRspExerciseOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspExerciseOrderAction(goHandle, pRspExerciseOrderAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorLockInsert(CHSRspLockInsertField *pRspExerciseOrderAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorLockInsert(goHandle, pRspExerciseOrderAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspForQuoteInsert(CHSRspForQuoteInsertField *pRspForQuoteInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspForQuoteInsert(goHandle, pRspForQuoteInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorQuoteInsert(CHSRspQuoteInsertField* pRspQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorQuoteInsert(goHandle, pRspQuoteInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQuoteAction(CHSRspQuoteActionField* pRspQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQuoteAction(goHandle, pRspQuoteAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorCombActionInsert(CHSRspCombActionInsertField *pRspCombActionInsert, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorCombActionInsert(goHandle, pRspCombActionInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQueryMaxOrderVolume(CHSRspQueryMaxOrderVolumeField *pRspQueryMaxOrderVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQueryMaxOrderVolume(goHandle, pRspQueryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryLockVolume(CHSRspQryLockVolumeField *pRspQryLockVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryLockVolume(goHandle, pRspQryLockVolume, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQueryExerciseVolume(CHSRspQueryExerciseVolumeField *pRspQueryExerciseVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQueryExerciseVolume(goHandle, pRspQueryExerciseVolume, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryCombVolume(CHSRspQryCombVolumeField *pRspQryCombVolume, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryCombVolume(goHandle, pRspQryCombVolume, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryPosition(CHSRspQryPositionField *pRspQryPosition, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryPosition(goHandle, pRspQryPosition, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryTradingAccount(CHSRspQryTradingAccountField *pRspQryTradingAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryTradingAccount(goHandle, pRspQryTradingAccount, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOrder(CHSOrderField *pRspQryOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOrder(goHandle, pRspQryOrder, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryTrade(CHSTradeField *pRspQryTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryTrade(goHandle, pRspQryTrade, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryExercise(CHSExerciseField *pRspQryExercise, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryExercise(goHandle, pRspQryExercise, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryLock(CHSLockField *pRspQryLock, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryLock(goHandle, pRspQryLock, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryCombAction(CHSCombActionField *pRspQryCombAction, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryCombAction(goHandle, pRspQryCombAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryForQuote(CHSForQuoteField* pRspQryForQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryForQuote(goHandle, pRspQryForQuote, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryQuote(CHSQuoteField* pRspQryQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryQuote(goHandle, pRspQryQuote, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryPositionCombineDetail(CHSRspQryPositionCombineDetailField *pRspQryPositionCombineDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryPositionCombineDetail(goHandle, pRspQryPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryInstrument(CHSRspQryInstrumentField *pRspQryInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryInstrument(goHandle, pRspQryInstrument, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryCoveredShort(CHSRspQryCoveredShortField *pRspQryCoveredShort, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryCoveredShort(goHandle, pRspQryCoveredShort, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryExerciseAssign(CHSRspQryExerciseAssignField *pRspQryExerciseAssign, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryExerciseAssign(goHandle, pRspQryExerciseAssign, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspTransfer(CHSRspTransferField *pRspTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspTransfer(goHandle, pRspTransfer, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryTransfer(CHSRspQryTransferField *pRspQryTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryTransfer(goHandle, pRspQryTransfer, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQueryBankBalance(CHSRspQueryBankBalanceField *pRspQueryBankBalance, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQueryBankBalance(goHandle, pRspQueryBankBalance, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQueryBankAccount(CHSRspQueryBankAccountField *pRspQueryBankAccount, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQueryBankAccount(goHandle, pRspQueryBankAccount, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspMultiCentreFundTrans(CHSRspMultiCentreFundTransField *pRspMultiCentreFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspMultiCentreFundTrans(goHandle, pRspMultiCentreFundTransfer, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQueryBillContent(CHSRspQueryBillContentField *pRspQueryBillContent, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQueryBillContent(goHandle, pRspQueryBillContent, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspBillConfirm(CHSRspBillConfirmField *pRspBillConfirm, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspBillConfirm(goHandle, pRspBillConfirm, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryMargin(CHSRspQryMarginField *pRspQryMargin, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryMargin(goHandle, pRspQryMargin, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryCommission(CHSRspQryCommissionField *pRspQryCommission, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryCommission(goHandle, pRspQryCommission, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryPositionDetail(CHSRspQryPositionDetailField *pRspQryPositionDetail, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryPositionDetail(goHandle, pRspQryPositionDetail, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryExchangeRate(CHSRspQryExchangeRateField *pRspQryExchangeRate, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryExchangeRate(goHandle, pRspQryExchangeRate, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQrySysConfig(CHSRspQrySysConfigField *pRspQrySysConfig, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQrySysConfig(goHandle, pRspQrySysConfig, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryDepthMarketData(CHSDepthMarketDataField *pRspQryDepthMarketData, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryDepthMarketData(goHandle, pRspQryDepthMarketData, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspFundTrans(CHSRspFundTransField *pRspFundTransfer, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspFundTrans(goHandle, pRspFundTransfer, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryFundTrans(CHSRspQryFundTransField *pRspQryFundTrans, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryFundTrans(goHandle, pRspQryFundTrans, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryClientNotice(CHSClientNoticeField *pRspQryClientNotice, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryClientNotice(goHandle, pRspQryClientNotice, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOptUnderly(CHSRspQryOptUnderlyField *pRspQryOptUnderly, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOptUnderly(goHandle, pRspQryOptUnderly, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQrySecuDepthMarket(CHSRspQrySecuDepthMarketField *pRspQrySecuDepthMarket, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQrySecuDepthMarket(goHandle, pRspQrySecuDepthMarket, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryHistOrder(CHSOrderField *pRspQryHistOrder, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryHistOrder(goHandle, pRspQryHistOrder, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryHistTrade(CHSTradeField *pRspQryHistTrade, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryHistTrade(goHandle, pRspQryHistTrade, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryWithdrawFund(CHSRspQryWithdrawFundField *pRspQryWithdrawFund, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryWithdrawFund(goHandle, pRspQryWithdrawFund, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryCombInstrument(CHSRspQryCombInstrumentField *pRspQryCombInstrument, CHSRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryCombInstrument(goHandle, pRspQryCombInstrument, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQrySeatID(CHSRspQrySeatIDField *pRspQrySeatID, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQrySeatID(goHandle, pRspQrySeatID, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspOptionSelfClose(CHSRspOptionSelfCloseField *pReqOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspOptionSelfClose(goHandle, pReqOptionSelfClose, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspOptionSelfCloseAction(CHSRspOptionSelfCloseActionField *pReqOptionSelfCloseAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspOptionSelfCloseAction(goHandle, pReqOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOptionSelfCloseResult(CHSRspQryOptionSelfCloseResultField *pReqQryOptionSelfCloseResult, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOptionSelfCloseResult(goHandle, pReqQryOptionSelfCloseResult, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOptionSelfClose(CHSOptionSelfCloseField *pRspQryOptionSelfClose, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOptionSelfClose(goHandle, pRspQryOptionSelfClose, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspErrorOptQuoteInsert(CHSRspOptQuoteInsertField *pRspOptQuoteInsert, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspErrorOptQuoteInsert(goHandle, pRspOptQuoteInsert, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspOptQuoteAction(CHSRspOptQuoteActionField *pRspOptQuoteAction, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspOptQuoteAction(goHandle, pRspOptQuoteAction, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOptQuote(CHSOptQuoteField *pRspQryOptQuote, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOptQuote(goHandle, pRspQryOptQuote, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRspQryOptCombStrategy(CHSRspQryOptCombStrategyField *pRspQryOptCombStrategy, CHSRspInfoField* pRspInfo, int nRequestID, bool bIsLast) {
        go_td_OnRspQryOptCombStrategy(goHandle, pRspQryOptCombStrategy, pRspInfo, nRequestID, bIsLast);
    }

    virtual void OnRtnTrade(CHSTradeField *pRtnTrade) {
        go_td_OnRtnTrade(goHandle, pRtnTrade);
    }

    virtual void OnRtnOrder(CHSOrderField *pRtnOrder) {
        go_td_OnRtnOrder(goHandle, pRtnOrder);
    }

    virtual void OnRtnExercise(CHSExerciseField *pRtnExercise) {
        go_td_OnRtnExercise(goHandle, pRtnExercise);
    }

    virtual void OnRtnCombAction(CHSCombActionField *pRtnCombAction) {
        go_td_OnRtnCombAction(goHandle, pRtnCombAction);
    }

    virtual void OnRtnLock(CHSLockField *pRtnLock) {
        go_td_OnRtnLock(goHandle, pRtnLock);
    }

    virtual void OnErrRtnOrderAction(CHSOrderActionField *pRtnOrder) {
        go_td_OnErrRtnOrderAction(goHandle, pRtnOrder);
    }

    virtual void OnRtnClientNotice(CHSClientNoticeField *pRtnClientNotice) {
        go_td_OnRtnClientNotice(goHandle, pRtnClientNotice);
    }

    virtual void OnRtnForQuote(CHSForQuoteField* pRtnForQuote) {
        go_td_OnRtnForQuote(goHandle, pRtnForQuote);
    }

    virtual void OnRtnQuote(CHSQuoteField* pRtnQuote) {
        go_td_OnRtnQuote(goHandle, pRtnQuote);
    }

    virtual void OnRtnExchangeStatus(CHSExchangeStatusField* pRtnExchangeStatus) {
        go_td_OnRtnExchangeStatus(goHandle, pRtnExchangeStatus);
    }

    virtual void OnRtnProductStatus(CHSProductStatusField* pRtnProductStatus) {
        go_td_OnRtnProductStatus(goHandle, pRtnProductStatus);
    }

    virtual void OnRtnOptionSelfClose(CHSOptionSelfCloseField *pRtnOptionSelfClose) {
        go_td_OnRtnOptionSelfClose(goHandle, pRtnOptionSelfClose);
    }

    virtual void OnRtnOptQuote(CHSOptQuoteField* pRtnOptQuote) {
        go_td_OnRtnOptQuote(goHandle, pRtnOptQuote);
    }

    virtual void OnRtnTransfer(CHSTransferField* pRtnTransfer) {
        go_td_OnRtnTransfer(goHandle, pRtnTransfer);
    }

    virtual void OnErrRtnOptQuoteAction(CHSOptQuoteActionField* pRtnOptQuoteAction) {
        go_td_OnErrRtnOptQuoteAction(goHandle, pRtnOptQuoteAction);
    }

private:
    uintptr_t goHandle;
};

class HSTradeApiWrap {
public:
    HSTradeApiWrap(uintptr_t goHandle, const char* dllPath)
        : goHandle(goHandle), pUserApi(nullptr), dllHandle(nullptr), pSpi(nullptr) {
        if (dllPath == nullptr || strlen(dllPath) == 0) {
            fprintf(stderr, "DLL path is empty\n");
            return;
        }

        dllHandle = dlopen(dllPath, RTLD_NOW);
        if (dllHandle == nullptr) {
            fprintf(stderr, "[%s] dlopen error: %s\n", dllPath, dlerror());
            return;
        }

        NewTradeApiCreator creator = (NewTradeApiCreator)dlsym(dllHandle, "NewTradeApi");
        if (creator == nullptr) {
            fprintf(stderr, "[%s] dlsym NewTradeApi error: %s\n", dllPath, dlerror());
            dlclose(dllHandle);
            dllHandle = nullptr;
            return;
        }

        pUserApi = creator(nullptr);
    }

    ~HSTradeApiWrap() {
        ReleaseApi();
    }

    void ReleaseApi() {
        if (pSpi != nullptr) {
            if (pUserApi != nullptr) {
                pUserApi->RegisterSpi(nullptr);
            }
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

    int Init(CHSInitConfigField *pInitCfgField, void* pExtraParam) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->Init(pInitCfgField, pExtraParam);
    }

    int Join() {
        if (pUserApi == nullptr) return -1;
        return pUserApi->Join();
    }

    int RegisterSubModel(SUB_TERT_TYPE eSubType) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->RegisterSubModel(eSubType);
    }

    int RegisterFront(const char *pszFrontAddress) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->RegisterFront(pszFrontAddress);
    }

    int RegisterFensServer(const char *pszFensAddress, const char *pszAccountID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->RegisterFensServer(pszFensAddress, pszAccountID);
    }

    void RegisterSpi(HSTradeSpiWrap* pSpi) {
        if (pUserApi == nullptr) return;
        this->pSpi = pSpi;
        pUserApi->RegisterSpi(pSpi);
    }

    const char* GetApiErrorMsg(int nErrorCode) {
        if (pUserApi == nullptr) return nullptr;
        return pUserApi->GetApiErrorMsg(nErrorCode);
    }

    int GetTradingDate() {
        if (pUserApi == nullptr) return -1;
        return pUserApi->GetTradingDate();
    }

    int BindSessionID(const uint8 nSessionID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->BindSessionID(nSessionID);
    }

    int ReqAuthenticate(CHSReqAuthenticateField *pReqAuthenticate, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqAuthenticate(pReqAuthenticate, nRequestID);
    }

    int ReqSubmitUserSystemInfo(CHSReqUserSystemInfoField *pReqUserSystemInfo, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqSubmitUserSystemInfo(pReqUserSystemInfo, nRequestID);
    }

    int ReqUserLogin(CHSReqUserLoginField *pReqUserLogin, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqUserLogin(pReqUserLogin, nRequestID);
    }

    int ReqUserPasswordUpdate(CHSReqUserPasswordUpdateField *pReqUserPasswordUpdate, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqUserPasswordUpdate(pReqUserPasswordUpdate, nRequestID);
    }

    int ReqOrderInsert(CHSReqOrderInsertField *pReqOrderInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOrderInsert(pReqOrderInsert, nRequestID);
    }

    int ReqOrderAction(CHSReqOrderActionField *pReqOrderAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOrderAction(pReqOrderAction, nRequestID);
    }

    int ReqExerciseOrderInsert(CHSReqExerciseOrderInsertField *pReqExerciseOrderInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqExerciseOrderInsert(pReqExerciseOrderInsert, nRequestID);
    }

    int ReqExerciseOrderAction(CHSReqExerciseOrderActionField *pReqExerciseOrderAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqExerciseOrderAction(pReqExerciseOrderAction, nRequestID);
    }

    int ReqLockInsert(CHSReqLockInsertField *pReqLockInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqLockInsert(pReqLockInsert, nRequestID);
    }

    int ReqForQuoteInsert(CHSReqForQuoteInsertField *pReqForQuoteInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqForQuoteInsert(pReqForQuoteInsert, nRequestID);
    }

    int ReqQuoteInsert(CHSReqQuoteInsertField* pReqQuoteInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQuoteInsert(pReqQuoteInsert, nRequestID);
    }

    int ReqQuoteAction(CHSReqQuoteActionField* pReqQuoteAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQuoteAction(pReqQuoteAction, nRequestID);
    }

    int ReqCombActionInsert(CHSReqCombActionInsertField *pReqCombActionInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqCombActionInsert(pReqCombActionInsert, nRequestID);
    }

    int ReqQueryMaxOrderVolume(CHSReqQueryMaxOrderVolumeField *pReqQueryMaxOrderVolume, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQueryMaxOrderVolume(pReqQueryMaxOrderVolume, nRequestID);
    }

    int ReqQryLockVolume(CHSReqQryLockVolumeField *pReqQryLockVolume, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryLockVolume(pReqQryLockVolume, nRequestID);
    }

    int ReqQueryExerciseVolume(CHSReqQueryExerciseVolumeField *pReqQueryExerciseVolume, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQueryExerciseVolume(pReqQueryExerciseVolume, nRequestID);
    }

    int ReqQryCombVolume(CHSReqQryCombVolumeField *pReqQryCombVolume, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryCombVolume(pReqQryCombVolume, nRequestID);
    }

    int ReqQryPosition(CHSReqQryPositionField *pReqQryPosition, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryPosition(pReqQryPosition, nRequestID);
    }

    int ReqQryTradingAccount(CHSReqQryTradingAccountField *pReqQryTradingAccount, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryTradingAccount(pReqQryTradingAccount, nRequestID);
    }

    int ReqQryOrder(CHSReqQryOrderField *pReqQryOrder, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOrder(pReqQryOrder, nRequestID);
    }

    int ReqQryTrade(CHSReqQryTradeField *pReqQryTrade, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryTrade(pReqQryTrade, nRequestID);
    }

    int ReqQryExercise(CHSReqQryExerciseField *pReqQryExercise, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryExercise(pReqQryExercise, nRequestID);
    }

    int ReqQryLock(CHSReqQryLockField *pReqQryLock, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryLock(pReqQryLock, nRequestID);
    }

    int ReqQryCombAction(CHSReqQryCombActionField *pReqQryCombAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryCombAction(pReqQryCombAction, nRequestID);
    }

    int ReqQryForQuote(CHSReqQryForQuoteField* pReqQryForQuote, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryForQuote(pReqQryForQuote, nRequestID);
    }

    int ReqQryQuote(CHSReqQryQuoteField* pReqQryQuote, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryQuote(pReqQryQuote, nRequestID);
    }

    int ReqQryPositionCombineDetail(CHSReqQryPositionCombineDetailField *pReqQryPositionCombineDetail, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryPositionCombineDetail(pReqQryPositionCombineDetail, nRequestID);
    }

    int ReqQryInstrument(CHSReqQryInstrumentField *pReqQryInstrument, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryInstrument(pReqQryInstrument, nRequestID);
    }

    int ReqQryCoveredShort(CHSReqQryCoveredShortField *pReqQryCoveredShort, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryCoveredShort(pReqQryCoveredShort, nRequestID);
    }

    int ReqQryExerciseAssign(CHSReqQryExerciseAssignField *pReqQryExerciseAssign, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryExerciseAssign(pReqQryExerciseAssign, nRequestID);
    }

    int ReqTransfer(CHSReqTransferField *pReqTransfer, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqTransfer(pReqTransfer, nRequestID);
    }

    int ReqQryTransfer(CHSReqQryTransferField *pReqQryTransfer, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryTransfer(pReqQryTransfer, nRequestID);
    }

    int ReqQueryBankBalance(CHSReqQueryBankBalanceField *pReqQueryBankBalance, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQueryBankBalance(pReqQueryBankBalance, nRequestID);
    }

    int ReqQueryBankAccount(CHSReqQueryBankAccountField *pReqQueryBankAccount, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQueryBankAccount(pReqQueryBankAccount, nRequestID);
    }

    int ReqMultiCentreFundTrans(CHSReqMultiCentreFundTransField *pReqMultiCentreFundTransfer, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqMultiCentreFundTrans(pReqMultiCentreFundTransfer, nRequestID);
    }

    int ReqQueryBillContent(CHSReqQueryBillContentField *pReqQueryBillContent, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQueryBillContent(pReqQueryBillContent, nRequestID);
    }

    int ReqBillConfirm(CHSReqBillConfirmField *pReqBillConfirm, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqBillConfirm(pReqBillConfirm, nRequestID);
    }

    int ReqQryMargin(CHSReqQryMarginField *pReqQryMargin, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryMargin(pReqQryMargin, nRequestID);
    }

    int ReqQryCommission(CHSReqQryCommissionField *pReqQryCommission, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryCommission(pReqQryCommission, nRequestID);
    }

    int ReqQryExchangeRate(CHSReqQryExchangeRateField *pReqQryExchangeRate, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryExchangeRate(pReqQryExchangeRate, nRequestID);
    }

    int ReqQryPositionDetail(CHSReqQryPositionDetailField *pReqQryPositionDetail, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryPositionDetail(pReqQryPositionDetail, nRequestID);
    }

    int ReqQrySysConfig(CHSReqQrySysConfigField *pReqQrySysConfig, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQrySysConfig(pReqQrySysConfig, nRequestID);
    }

    int ReqQryDepthMarketData(CHSReqQryDepthMarketDataField *pReqQryDepthMarketData, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryDepthMarketData(pReqQryDepthMarketData, nRequestID);
    }

    int ReqFundTrans(CHSReqFundTransField *pReqFundTransfer, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqFundTrans(pReqFundTransfer, nRequestID);
    }

    int ReqQryFundTrans(CHSReqQryFundTransField *pReqQryFundTrans, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryFundTrans(pReqQryFundTrans, nRequestID);
    }

    int ReqQryClientNotice(CHSReqQryClientNoticeField *pReqQryClientNotice, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryClientNotice(pReqQryClientNotice, nRequestID);
    }

    int ReqQryOptUnderly(CHSReqQryOptUnderlyField *pReqQryOptUnderly, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOptUnderly(pReqQryOptUnderly, nRequestID);
    }

    int ReqQrySecuDepthMarket(CHSReqQrySecuDepthMarketField *pReqQrySecuDepthMarket, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQrySecuDepthMarket(pReqQrySecuDepthMarket, nRequestID);
    }

    int ReqQryHistOrder(CHSReqQryHistOrderField *pReqQryHistOrder, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryHistOrder(pReqQryHistOrder, nRequestID);
    }

    int ReqQryHistTrade(CHSReqQryHistTradeField *pReqQryHistTrade, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryHistTrade(pReqQryHistTrade, nRequestID);
    }

    int ReqQryWithdrawFund(CHSReqQryWithdrawFundField *pReqQryWithdrawFund, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryWithdrawFund(pReqQryWithdrawFund, nRequestID);
    }

    int ReqQryCombInstrument(CHSReqQryCombInstrumentField *pReqQryCombInstrument, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryCombInstrument(pReqQryCombInstrument, nRequestID);
    }

    int ReqQrySeatID(CHSReqQrySeatIDField *pReqQrySeatID, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQrySeatID(pReqQrySeatID, nRequestID);
    }

    int ReqOptionSelfClose(CHSReqOptionSelfCloseField *pReqOptionSelfClose, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOptionSelfClose(pReqOptionSelfClose, nRequestID);
    }

    int ReqOptionSelfCloseAction(CHSReqOptionSelfCloseActionField *pReqOptionSelfCloseAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOptionSelfCloseAction(pReqOptionSelfCloseAction, nRequestID);
    }

    int ReqQryOptionSelfCloseResult(CHSReqQryOptionSelfCloseResultField *pReqQryOptionSelfCloseResult, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOptionSelfCloseResult(pReqQryOptionSelfCloseResult, nRequestID);
    }

    int ReqQryOptionSelfClose(CHSReqQryOptionSelfCloseField *pReqQryOptionSelfClose, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOptionSelfClose(pReqQryOptionSelfClose, nRequestID);
    }

    int ReqOptQuoteInsert(CHSReqOptQuoteInsertField *pReqOptQuoteInsert, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOptQuoteInsert(pReqOptQuoteInsert, nRequestID);
    }

    int ReqOptQuoteAction(CHSReqOptQuoteActionField *pReqOptQuoteAction, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqOptQuoteAction(pReqOptQuoteAction, nRequestID);
    }

    int ReqQryOptQuote(CHSReqQryOptQuoteField *pReqQryOptQuote, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOptQuote(pReqQryOptQuote, nRequestID);
    }

    int ReqQryOptCombStrategy(CHSReqQryOptCombStrategyField *pReqQryOptCombStrategy, int nRequestID) {
        if (pUserApi == nullptr) return -1;
        return pUserApi->ReqQryOptCombStrategy(pReqQryOptCombStrategy, nRequestID);
    }

private:
    uintptr_t goHandle;
    CHSTradeApi* pUserApi;
    void* dllHandle;
    HSTradeSpiWrap* pSpi;
};

uintptr_t hs_create_td_api(const char* dllPath) {
    return reinterpret_cast<uintptr_t>(new HSTradeApiWrap(0, dllPath));
}

void hs_release_td_api(uintptr_t api) {
    delete reinterpret_cast<HSTradeApiWrap*>(api);
}

const char* hs_td_GetApiVersion(uintptr_t api) {
    return nullptr;
}

int hs_td_Init(uintptr_t api, void* config) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->Init(
        reinterpret_cast<CHSInitConfigField*>(config), nullptr);
}

int hs_td_Join(uintptr_t api) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->Join();
}

int hs_td_RegisterSubModel(uintptr_t api, int eSubType) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->RegisterSubModel((SUB_TERT_TYPE)eSubType);
}

int hs_td_RegisterFront(uintptr_t api, const char* frontAddress) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->RegisterFront(frontAddress);
}

int hs_td_RegisterFensServer(uintptr_t api, const char* fensAddress, const char* accountID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->RegisterFensServer(fensAddress, accountID);
}

void hs_td_RegisterSpi(uintptr_t api, uintptr_t spi) {
    reinterpret_cast<HSTradeApiWrap*>(api)->RegisterSpi(
        reinterpret_cast<HSTradeSpiWrap*>(spi));
}

const char* hs_td_GetApiErrorMsg(int errorCode) {
    return nullptr;
}

int hs_td_GetTradingDate(uintptr_t api) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->GetTradingDate();
}

int hs_td_BindSessionID(uintptr_t api, uint8_t nSessionID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->BindSessionID(nSessionID);
}

int hs_td_ReqAuthenticate(uintptr_t api, void* pReqAuthenticateField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqAuthenticate(
        reinterpret_cast<CHSReqAuthenticateField*>(pReqAuthenticateField), nRequestID);
}

int hs_td_ReqSubmitUserSystemInfo(uintptr_t api, void* pReqUserSystemInfoField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqSubmitUserSystemInfo(
        reinterpret_cast<CHSReqUserSystemInfoField*>(pReqUserSystemInfoField), nRequestID);
}

int hs_td_ReqUserLogin(uintptr_t api, void* pReqUserLoginField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqUserLogin(
        reinterpret_cast<CHSReqUserLoginField*>(pReqUserLoginField), nRequestID);
}

int hs_td_ReqUserPasswordUpdate(uintptr_t api, void* pReqUserPasswordUpdateField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqUserPasswordUpdate(
        reinterpret_cast<CHSReqUserPasswordUpdateField*>(pReqUserPasswordUpdateField), nRequestID);
}

int hs_td_ReqOrderInsert(uintptr_t api, void* pReqOrderInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOrderInsert(
        reinterpret_cast<CHSReqOrderInsertField*>(pReqOrderInsertField), nRequestID);
}

int hs_td_ReqOrderAction(uintptr_t api, void* pReqOrderActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOrderAction(
        reinterpret_cast<CHSReqOrderActionField*>(pReqOrderActionField), nRequestID);
}

int hs_td_ReqExerciseOrderInsert(uintptr_t api, void* pReqExerciseOrderInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqExerciseOrderInsert(
        reinterpret_cast<CHSReqExerciseOrderInsertField*>(pReqExerciseOrderInsertField), nRequestID);
}

int hs_td_ReqExerciseOrderAction(uintptr_t api, void* pReqExerciseOrderActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqExerciseOrderAction(
        reinterpret_cast<CHSReqExerciseOrderActionField*>(pReqExerciseOrderActionField), nRequestID);
}

int hs_td_ReqLockInsert(uintptr_t api, void* pReqLockInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqLockInsert(
        reinterpret_cast<CHSReqLockInsertField*>(pReqLockInsertField), nRequestID);
}

int hs_td_ReqForQuoteInsert(uintptr_t api, void* pReqForQuoteInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqForQuoteInsert(
        reinterpret_cast<CHSReqForQuoteInsertField*>(pReqForQuoteInsertField), nRequestID);
}

int hs_td_ReqQuoteInsert(uintptr_t api, void* pReqQuoteInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQuoteInsert(
        reinterpret_cast<CHSReqQuoteInsertField*>(pReqQuoteInsertField), nRequestID);
}

int hs_td_ReqQuoteAction(uintptr_t api, void* pReqQuoteActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQuoteAction(
        reinterpret_cast<CHSReqQuoteActionField*>(pReqQuoteActionField), nRequestID);
}

int hs_td_ReqCombActionInsert(uintptr_t api, void* pReqCombActionInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqCombActionInsert(
        reinterpret_cast<CHSReqCombActionInsertField*>(pReqCombActionInsertField), nRequestID);
}

int hs_td_ReqQueryMaxOrderVolume(uintptr_t api, void* pReqQueryMaxOrderVolumeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQueryMaxOrderVolume(
        reinterpret_cast<CHSReqQueryMaxOrderVolumeField*>(pReqQueryMaxOrderVolumeField), nRequestID);
}

int hs_td_ReqQryLockVolume(uintptr_t api, void* pReqQryLockVolumeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryLockVolume(
        reinterpret_cast<CHSReqQryLockVolumeField*>(pReqQryLockVolumeField), nRequestID);
}

int hs_td_ReqQueryExerciseVolume(uintptr_t api, void* pReqQueryExerciseVolumeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQueryExerciseVolume(
        reinterpret_cast<CHSReqQueryExerciseVolumeField*>(pReqQueryExerciseVolumeField), nRequestID);
}

int hs_td_ReqQryCombVolume(uintptr_t api, void* pReqQryCombVolumeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryCombVolume(
        reinterpret_cast<CHSReqQryCombVolumeField*>(pReqQryCombVolumeField), nRequestID);
}

int hs_td_ReqQryPosition(uintptr_t api, void* pReqQryPositionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryPosition(
        reinterpret_cast<CHSReqQryPositionField*>(pReqQryPositionField), nRequestID);
}

int hs_td_ReqQryTradingAccount(uintptr_t api, void* pReqQryTradingAccountField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryTradingAccount(
        reinterpret_cast<CHSReqQryTradingAccountField*>(pReqQryTradingAccountField), nRequestID);
}

int hs_td_ReqQryOrder(uintptr_t api, void* pReqQryOrderField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOrder(
        reinterpret_cast<CHSReqQryOrderField*>(pReqQryOrderField), nRequestID);
}

int hs_td_ReqQryTrade(uintptr_t api, void* pReqQryTradeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryTrade(
        reinterpret_cast<CHSReqQryTradeField*>(pReqQryTradeField), nRequestID);
}

int hs_td_ReqQryExercise(uintptr_t api, void* pReqQryExerciseField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryExercise(
        reinterpret_cast<CHSReqQryExerciseField*>(pReqQryExerciseField), nRequestID);
}

int hs_td_ReqQryLock(uintptr_t api, void* pReqQryLockField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryLock(
        reinterpret_cast<CHSReqQryLockField*>(pReqQryLockField), nRequestID);
}

int hs_td_ReqQryCombAction(uintptr_t api, void* pReqQryCombActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryCombAction(
        reinterpret_cast<CHSReqQryCombActionField*>(pReqQryCombActionField), nRequestID);
}

int hs_td_ReqQryForQuote(uintptr_t api, void* pReqQryForQuoteField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryForQuote(
        reinterpret_cast<CHSReqQryForQuoteField*>(pReqQryForQuoteField), nRequestID);
}

int hs_td_ReqQryQuote(uintptr_t api, void* pReqQryQuoteField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryQuote(
        reinterpret_cast<CHSReqQryQuoteField*>(pReqQryQuoteField), nRequestID);
}

int hs_td_ReqQryPositionCombineDetail(uintptr_t api, void* pReqQryPositionCombineDetailField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryPositionCombineDetail(
        reinterpret_cast<CHSReqQryPositionCombineDetailField*>(pReqQryPositionCombineDetailField), nRequestID);
}

int hs_td_ReqQryInstrument(uintptr_t api, void* pReqQryInstrumentField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryInstrument(
        reinterpret_cast<CHSReqQryInstrumentField*>(pReqQryInstrumentField), nRequestID);
}

int hs_td_ReqQryCoveredShort(uintptr_t api, void* pReqQryCoveredShortField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryCoveredShort(
        reinterpret_cast<CHSReqQryCoveredShortField*>(pReqQryCoveredShortField), nRequestID);
}

int hs_td_ReqQryExerciseAssign(uintptr_t api, void* pReqQryExerciseAssignField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryExerciseAssign(
        reinterpret_cast<CHSReqQryExerciseAssignField*>(pReqQryExerciseAssignField), nRequestID);
}

int hs_td_ReqTransfer(uintptr_t api, void* pReqTransferField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqTransfer(
        reinterpret_cast<CHSReqTransferField*>(pReqTransferField), nRequestID);
}

int hs_td_ReqQryTransfer(uintptr_t api, void* pReqQryTransferField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryTransfer(
        reinterpret_cast<CHSReqQryTransferField*>(pReqQryTransferField), nRequestID);
}

int hs_td_ReqQueryBankBalance(uintptr_t api, void* pReqQueryBankBalanceField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQueryBankBalance(
        reinterpret_cast<CHSReqQueryBankBalanceField*>(pReqQueryBankBalanceField), nRequestID);
}

int hs_td_ReqQueryBankAccount(uintptr_t api, void* pReqQueryBankAccountField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQueryBankAccount(
        reinterpret_cast<CHSReqQueryBankAccountField*>(pReqQueryBankAccountField), nRequestID);
}

int hs_td_ReqMultiCentreFundTrans(uintptr_t api, void* pReqMultiCentreFundTransField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqMultiCentreFundTrans(
        reinterpret_cast<CHSReqMultiCentreFundTransField*>(pReqMultiCentreFundTransField), nRequestID);
}

int hs_td_ReqQueryBillContent(uintptr_t api, void* pReqQueryBillContentField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQueryBillContent(
        reinterpret_cast<CHSReqQueryBillContentField*>(pReqQueryBillContentField), nRequestID);
}

int hs_td_ReqBillConfirm(uintptr_t api, void* pReqBillConfirmField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqBillConfirm(
        reinterpret_cast<CHSReqBillConfirmField*>(pReqBillConfirmField), nRequestID);
}

int hs_td_ReqQryMargin(uintptr_t api, void* pReqQryMarginField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryMargin(
        reinterpret_cast<CHSReqQryMarginField*>(pReqQryMarginField), nRequestID);
}

int hs_td_ReqQryCommission(uintptr_t api, void* pReqQryCommissionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryCommission(
        reinterpret_cast<CHSReqQryCommissionField*>(pReqQryCommissionField), nRequestID);
}

int hs_td_ReqQryExchangeRate(uintptr_t api, void* pReqQryExchangeRateField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryExchangeRate(
        reinterpret_cast<CHSReqQryExchangeRateField*>(pReqQryExchangeRateField), nRequestID);
}

int hs_td_ReqQryPositionDetail(uintptr_t api, void* pReqQryPositionDetailField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryPositionDetail(
        reinterpret_cast<CHSReqQryPositionDetailField*>(pReqQryPositionDetailField), nRequestID);
}

int hs_td_ReqQrySysConfig(uintptr_t api, void* pReqQrySysConfigField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQrySysConfig(
        reinterpret_cast<CHSReqQrySysConfigField*>(pReqQrySysConfigField), nRequestID);
}

int hs_td_ReqQryDepthMarketData(uintptr_t api, void* pReqQryDepthMarketDataField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryDepthMarketData(
        reinterpret_cast<CHSReqQryDepthMarketDataField*>(pReqQryDepthMarketDataField), nRequestID);
}

int hs_td_ReqFundTrans(uintptr_t api, void* pReqFundTransferField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqFundTrans(
        reinterpret_cast<CHSReqFundTransField*>(pReqFundTransferField), nRequestID);
}

int hs_td_ReqQryFundTrans(uintptr_t api, void* pReqQryFundTransField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryFundTrans(
        reinterpret_cast<CHSReqQryFundTransField*>(pReqQryFundTransField), nRequestID);
}

int hs_td_ReqQryClientNotice(uintptr_t api, void* pReqQryClientNoticeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryClientNotice(
        reinterpret_cast<CHSReqQryClientNoticeField*>(pReqQryClientNoticeField), nRequestID);
}

int hs_td_ReqQryOptUnderly(uintptr_t api, void* pReqQryOptUnderlyField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOptUnderly(
        reinterpret_cast<CHSReqQryOptUnderlyField*>(pReqQryOptUnderlyField), nRequestID);
}

int hs_td_ReqQrySecuDepthMarket(uintptr_t api, void* pReqQrySecuDepthMarketField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQrySecuDepthMarket(
        reinterpret_cast<CHSReqQrySecuDepthMarketField*>(pReqQrySecuDepthMarketField), nRequestID);
}

int hs_td_ReqQryHistOrder(uintptr_t api, void* pReqQryHistOrderField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryHistOrder(
        reinterpret_cast<CHSReqQryHistOrderField*>(pReqQryHistOrderField), nRequestID);
}

int hs_td_ReqQryHistTrade(uintptr_t api, void* pReqQryHistTradeField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryHistTrade(
        reinterpret_cast<CHSReqQryHistTradeField*>(pReqQryHistTradeField), nRequestID);
}

int hs_td_ReqQryWithdrawFund(uintptr_t api, void* pReqQryWithdrawFundField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryWithdrawFund(
        reinterpret_cast<CHSReqQryWithdrawFundField*>(pReqQryWithdrawFundField), nRequestID);
}

int hs_td_ReqQryCombInstrument(uintptr_t api, void* pReqQryCombInstrumentField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryCombInstrument(
        reinterpret_cast<CHSReqQryCombInstrumentField*>(pReqQryCombInstrumentField), nRequestID);
}

int hs_td_ReqQrySeatID(uintptr_t api, void* pReqQrySeatIDField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQrySeatID(
        reinterpret_cast<CHSReqQrySeatIDField*>(pReqQrySeatIDField), nRequestID);
}

int hs_td_ReqOptionSelfClose(uintptr_t api, void* pReqOptionSelfCloseField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOptionSelfClose(
        reinterpret_cast<CHSReqOptionSelfCloseField*>(pReqOptionSelfCloseField), nRequestID);
}

int hs_td_ReqOptionSelfCloseAction(uintptr_t api, void* pReqOptionSelfCloseActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOptionSelfCloseAction(
        reinterpret_cast<CHSReqOptionSelfCloseActionField*>(pReqOptionSelfCloseActionField), nRequestID);
}

int hs_td_ReqQryOptionSelfCloseResult(uintptr_t api, void* pReqQryOptionSelfCloseResultField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOptionSelfCloseResult(
        reinterpret_cast<CHSReqQryOptionSelfCloseResultField*>(pReqQryOptionSelfCloseResultField), nRequestID);
}

int hs_td_ReqQryOptionSelfClose(uintptr_t api, void* pReqQryOptionSelfCloseField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOptionSelfClose(
        reinterpret_cast<CHSReqQryOptionSelfCloseField*>(pReqQryOptionSelfCloseField), nRequestID);
}

int hs_td_ReqOptQuoteInsert(uintptr_t api, void* pReqOptQuoteInsertField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOptQuoteInsert(
        reinterpret_cast<CHSReqOptQuoteInsertField*>(pReqOptQuoteInsertField), nRequestID);
}

int hs_td_ReqOptQuoteAction(uintptr_t api, void* pReqOptQuoteActionField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqOptQuoteAction(
        reinterpret_cast<CHSReqOptQuoteActionField*>(pReqOptQuoteActionField), nRequestID);
}

int hs_td_ReqQryOptQuote(uintptr_t api, void* pReqQryOptQuoteField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOptQuote(
        reinterpret_cast<CHSReqQryOptQuoteField*>(pReqQryOptQuoteField), nRequestID);
}

int hs_td_ReqQryOptCombStrategy(uintptr_t api, void* pReqQryOptCombStrategyField, int nRequestID) {
    return reinterpret_cast<HSTradeApiWrap*>(api)->ReqQryOptCombStrategy(
        reinterpret_cast<CHSReqQryOptCombStrategyField*>(pReqQryOptCombStrategyField), nRequestID);
}

uintptr_t hs_create_td_spi(uintptr_t goHandle) {
    return reinterpret_cast<uintptr_t>(new HSTradeSpiWrap(goHandle));
}

void hs_release_td_spi(uintptr_t spi) {
    delete reinterpret_cast<HSTradeSpiWrap*>(spi);
}
