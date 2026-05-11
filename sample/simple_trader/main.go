package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nandy33up/go2uft/thost"
	"github.com/nandy33up/go2uft/uft"
)

type TradeSpi struct {
	uft.BaseTradeSpi
}

func NewTradeSpi() *TradeSpi {
	return &TradeSpi{}
}

func (s *TradeSpi) OnFrontConnected() {
	log.Println("Trade: Front connected")

	authField := &thost.CHSReqAuthenticateField{}
	copy(authField.AccountID[:], "your_account_id")
	copy(authField.Password[:], "your_password")
	copy(authField.AppID[:], "your_app_id")
	copy(authField.AuthCode[:], "your_auth_code")

	fmt.Println("Trade: Sending authenticate request...")
}

func (s *TradeSpi) OnFrontDisconnected(nResult int) {
	log.Printf("Trade: Front disconnected, result=%d", nResult)
}

func (s *TradeSpi) OnRspUserLogin(pRspUserLoginField *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		log.Printf("Trade: Login failed, errorID=%d, errorMsg=%s", pRspInfo.ErrorID, string(pRspInfo.ErrorMsg[:]))
		return
	}
	log.Printf("Trade: Login success, account=%s, tradingDay=%08d",
		string(pRspUserLoginField.AccountID[:]),
		pRspUserLoginField.TradingDay)

	fmt.Println("Trade: Querying position...")
}

func (s *TradeSpi) OnRtnOrder(pOrder *thost.CHSOrderField) {
	orderRef := string(pOrder.OrderRef[:])
	orderSysID := string(pOrder.OrderSysID[:])
	orderStatus := pOrder.OrderStatus
	log.Printf("Trade: Order update - Ref=%s SysID=%s Status=%c",
		orderRef, orderSysID, orderStatus)
}

func (s *TradeSpi) OnRtnTrade(pTrade *thost.CHSTradeField) {
	tradeID := string(pTrade.TradeID[:])
	instID := string(pTrade.InstrumentID[:])
	direction := pTrade.Direction
	offsetFlag := pTrade.OffsetFlag
	price := pTrade.TradePrice
	volume := pTrade.TradeVolume
	log.Printf("Trade: Trade - ID=%s %s %c%c Price=%.2f Volume=%d",
		tradeID, instID, direction, offsetFlag, price, volume)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Go2UFT Trade Demo (Static Library Version)")
	log.Println("==========================================")

	tdapi := uft.CreateTradeApi()

	spi := NewTradeSpi()
	tdapi.RegisterSpi(spi)

	frontAddr := "tcp://101.71.12.149:8102"
	log.Printf("Connecting to: %s", frontAddr)
	tdapi.RegisterFront(frontAddr)

	log.Printf("API Version: %s", tdapi.GetApiVersion())

	initCfg := new(thost.CHSInitConfigField)
	initCfg.APICheckVersion = thost.API_STRUCT_CHECK_VERSION
	copy(initCfg.CommLicense[:], "license.dat")
	ret := tdapi.Init(initCfg, nil)
	if ret != 0 {
		log.Printf("Trade: Init failed with code: %d, %s", ret, tdapi.GetApiErrorMsg(ret))
	} else {
		log.Println("Trade: Waiting for connection...")
	}

	go func() {
		tdapi.Join()
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Println("Shutting down...")
	}

	tdapi.ReleaseApi()
	log.Println("Done")
}
