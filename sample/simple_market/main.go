package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nandy33up/go2uft/thost"
	"github.com/nandy33up/go2uft/uft"
)

type MarketSpi struct {
	uft.BaseMdSpi
}

var mdapi thost.CHSMdApi

func NewMarketSpi() *MarketSpi {
	return &MarketSpi{}
}

func (s *MarketSpi) OnFrontConnected() {
	log.Println("Market: Front connected, subscribing to instruments...")

	subscribes := []thost.CHSReqDepthMarketDataField{
		{},
		{},
		{},
		{},
		{},
	}
	copy(subscribes[0].ExchangeID[:], "1")
	copy(subscribes[0].InstrumentID[:], "588000")
	copy(subscribes[1].ExchangeID[:], "1")
	copy(subscribes[1].InstrumentID[:], "588080")
	copy(subscribes[2].ExchangeID[:], "2")
	copy(subscribes[2].InstrumentID[:], "159915")
	copy(subscribes[3].ExchangeID[:], "1")
	copy(subscribes[3].InstrumentID[:], "10011299")
	copy(subscribes[4].ExchangeID[:], "2")
	copy(subscribes[4].InstrumentID[:], "90007293")

	mdapi.ReqDepthMarketDataSubscribe(subscribes, len(subscribes), 1)
	log.Println("Market: Subscription request sent")
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
	log.Println("Go2UFT Market Demo (Static Library Version)")
	log.Println("==========================================")

	mdapi = uft.CreateMdApi()

	spi := NewMarketSpi()
	mdapi.RegisterSpi(spi)

	frontAddr := "tcp://101.71.12.149:8102"
	log.Printf("Connecting to: %s", frontAddr)
	mdapi.RegisterFront(frontAddr)

	log.Printf("API Version: %s", mdapi.GetApiVersion())

	var license thost.HSLicenseFile
	copy(license[:], "/root/go2uft/license.dat")
	initCfg := &thost.CHSInitConfigField{
		APICheckVersion: 240002,
		CommLicense:     license,
	}
	ret := mdapi.Init(initCfg, nil)
	if ret != 0 {
		log.Printf("Market: Init failed with code: %d", ret)
	} else {
		log.Println("Market: Waiting for connection...")
	}

	go func() {
		mdapi.Join()
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Println("Shutting down...")
	}

	mdapi.ReleaseApi()
	log.Println("Done")
}
