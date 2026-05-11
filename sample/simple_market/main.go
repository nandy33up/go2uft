package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nandy33up/go2uft/thost"
	"github.com/nandy33up/go2uft/uft_dyn"
)

type MarketSpi struct {
	uft_dyn.BaseMdSpi
}

func NewMarketSpi() *MarketSpi {
	return &MarketSpi{}
}

func (s *MarketSpi) OnFrontConnected() {
	log.Println("Market: Front connected")

	loginField := &thost.CHSReqUserLoginField{}
	copy(loginField.AccountID[:], "your_account_id")
	copy(loginField.Password[:], "your_password")
	loginField.UserApplicationType = thost.HSUserApplicationType(thost.HS_AT_Investor)

	fmt.Println("Market: Sending login request...")
}

func (s *MarketSpi) OnFrontDisconnected(nResult int) {
	log.Printf("Market: Front disconnected, result=%d", nResult)
}

func (s *MarketSpi) OnRtnDepthMarketData(pDepthMarketData *thost.CHSDepthMarketDataField) {
	instID := string(pDepthMarketData.InstrumentID[:])
	lastPrice := pDepthMarketData.LastPrice
	bid1 := pDepthMarketData.BidPrice1
	bidVol1 := pDepthMarketData.BidVolume1
	ask1 := pDepthMarketData.AskPrice1
	askVol1 := pDepthMarketData.AskVolume1
	log.Printf("Market: %s LastPrice=%.2f Bid1=%.2f(%f) Ask1=%.2f(%f)",
		instID, lastPrice, bid1, bidVol1, ask1, askVol1)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Go2UFT Market Demo (Dynamic Library Version)")
	log.Println("==========================================")

	mdapi := uft_dyn.CreateMdApi(
		uft_dyn.DynamicLibPath("v3.7.4/sdk/linux.x64/libHSMdApi.so"),
	)

	spi := NewMarketSpi()
	mdapi.RegisterSpi(spi)

	frontAddr := "tcp://101.71.12.149:8102"
	log.Printf("Connecting to: %s", frontAddr)
	mdapi.RegisterFront(frontAddr)

	log.Printf("API Version: %s", mdapi.GetApiVersion())

	var license thost.HSLicenseFile
	copy(license[:], "license.dat")
	mdapi.Init(&thost.CHSInitConfigField{CommLicense: license}, nil)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Println("Shutting down...")
	}

	mdapi.ReleaseApi()
	log.Println("Done")
}
