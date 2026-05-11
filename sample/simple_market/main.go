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

type MarketSubscribes struct {
	ExchangeID   string
	InstrumentID string
}

func (s *MarketSpi) OnFrontConnected() {
	log.Println("Market: Front connected, subscribing to instruments...")
	subscribes := []MarketSubscribes{
		{"1", "588000"},
		{"1", "588080"},
		{"2", "159915"},
		{"1", "10011299"},
		{"2", "90007293"},
	}
	pSubscribes := make([]thost.CHSReqDepthMarketDataField, len(subscribes))
	for i, s := range subscribes {
		copy(pSubscribes[i].ExchangeID[:], s.ExchangeID[:])
		copy(pSubscribes[i].InstrumentID[:], s.InstrumentID[:])
	}

	mdapi.ReqDepthMarketDataSubscribe(pSubscribes, len(pSubscribes), 1)
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
	log.Printf("Market: %s LastPrice=%.4f Bid1=%.4f(%.0f), Ask1=%.4f(%.0f)",
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

	initCfg := new(thost.CHSInitConfigField)
	initCfg.APICheckVersion = thost.API_STRUCT_CHECK_VERSION
	copy(initCfg.CommLicense[:], "license.dat")
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
