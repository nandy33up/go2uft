package thost

import "unsafe"

type CHSMdApi interface {
	GetApiVersion() string
	ReleaseApi()
	Init(pInitCfgField *CHSInitConfigField, pExtraParam unsafe.Pointer) int
	Join() int
	RegisterFront(pszFrontAddress string) int
	RegisterFensServer(pszFensAddress string, pszAccountID string) int
	RegisterSpi(pSpi CHSMdSpi)
	ReqDepthMarketDataSubscribe(pReqDepthMarketDataSubscribe []CHSReqDepthMarketDataField, nCount int, nRequestID int) int
	ReqDepthMarketDataCancel(pReqDepthMarketDataCancel []CHSReqDepthMarketDataField, nCount int, nRequestID int) int
	ReqForQuoteSubscribe(pReqForQuoteSubscribe []CHSReqForQuoteField, nCount int, nRequestID int) int
	ReqForQuoteCancel(pReqForQuoteCancel []CHSReqForQuoteField, nCount int, nRequestID int) int
	GetApiErrorMsg(nErrorCode int) string
}

type CHSMdSpi interface {
	OnFrontConnected()
	OnFrontDisconnected(nResult int)
	OnRspDepthMarketDataSubscribe(pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspDepthMarketDataCancel(pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRtnDepthMarketData(pDepthMarketData *CHSDepthMarketDataField)
	OnRspForQuoteSubscribe(pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRspForQuoteCancel(pRspInfo *CHSRspInfoField, nRequestID int, bIsLast bool)
	OnRtnForQuote(pForQuote *CHSForQuoteField)
}
