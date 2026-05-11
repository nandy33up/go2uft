package uft_dyn

import (
	"github.com/nandy33up/go2uft/thost"
)

var _ thost.CHSMdSpi = &BaseMdSpi{}

type BaseMdSpi struct {
	OnFrontConnectedCallback              func()
	OnFrontDisconnectedCallback           func(int)
	OnHeartBeatWarningCallback            func(int)
	OnRspDepthMarketDataSubscribeCallback func(*thost.CHSRspInfoField, int, bool)
	OnRspDepthMarketDataCancelCallback    func(*thost.CHSRspInfoField, int, bool)
	OnRspForQuoteSubscribeCallback        func(*thost.CHSRspInfoField, int, bool)
	OnRspForQuoteCancelCallback           func(*thost.CHSRspInfoField, int, bool)
	OnRtnDepthMarketDataCallback          func(*thost.CHSDepthMarketDataField)
	OnRtnForQuoteCallback                 func(*thost.CHSForQuoteField)
	OnRspErrorCallback                    func(*thost.CHSRspInfoField, int, bool)
	OnRspUserLoginCallback                func(*thost.CHSRspUserLoginField, *thost.CHSRspInfoField, int, bool)
	OnRspErrorMsgCallback                 func(*thost.CHSRspInfoField, int, bool)
	OnRspUserSystemInfoCallback           func(*thost.CHSRspUserSystemInfoField, *thost.CHSRspInfoField, int, bool)
}

func (s *BaseMdSpi) OnFrontConnected() {
	if s.OnFrontConnectedCallback != nil {
		s.OnFrontConnectedCallback()
	}
}

func (s *BaseMdSpi) OnFrontDisconnected(nReason int) {
	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(nReason)
	}
}

func (s *BaseMdSpi) OnHeartBeatWarning(nTimeLapse int) {
	if s.OnHeartBeatWarningCallback != nil {
		s.OnHeartBeatWarningCallback(nTimeLapse)
	}
}

func (s *BaseMdSpi) OnRspDepthMarketDataSubscribe(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspDepthMarketDataSubscribeCallback != nil {
		s.OnRspDepthMarketDataSubscribeCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspDepthMarketDataCancel(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspDepthMarketDataCancelCallback != nil {
		s.OnRspDepthMarketDataCancelCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspForQuoteSubscribe(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteSubscribeCallback != nil {
		s.OnRspForQuoteSubscribeCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspForQuoteCancel(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteCancelCallback != nil {
		s.OnRspForQuoteCancelCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRtnDepthMarketData(pDepthMarketData *thost.CHSDepthMarketDataField) {
	if s.OnRtnDepthMarketDataCallback != nil {
		s.OnRtnDepthMarketDataCallback(pDepthMarketData)
	}
}

func (s *BaseMdSpi) OnRtnForQuote(pForQuote *thost.CHSForQuoteField) {
	if s.OnRtnForQuoteCallback != nil {
		s.OnRtnForQuoteCallback(pForQuote)
	}
}

func (s *BaseMdSpi) OnRspError(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspUserLogin(pRspUserLogin *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserLoginCallback != nil {
		s.OnRspUserLoginCallback(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspUserLogout(pUserLogout *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
}

func (s *BaseMdSpi) OnRspErrorMsg(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorMsgCallback != nil {
		s.OnRspErrorMsgCallback(pRspInfo, nRequestID, bIsLast)
	}
}

func (s *BaseMdSpi) OnRspUserSystemInfo(pRspUserSystemInfo *thost.CHSRspUserSystemInfoField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserSystemInfoCallback != nil {
		s.OnRspUserSystemInfoCallback(pRspUserSystemInfo, pRspInfo, nRequestID, bIsLast)
	}
}
