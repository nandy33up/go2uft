package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nandy33up/go2uft/thost"
	"github.com/nandy33up/go2uft/uft"
)

const (
	FrontAddr   = "tcp://101.71.12.149:8102"
	LicenseFile = "license.dat"
)

type Quoter struct {
	uft.BaseMdSpi
	api thost.CHSMdApi
	seq int
}

func NewQuoter() *Quoter {
	return &Quoter{api: uft.CreateMdApi()}
}

type MarketSubscribes struct {
	ExchangeID   string
	InstrumentID string
}

func (q *Quoter) OnFrontConnected() {
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

	q.api.ReqDepthMarketDataSubscribe(pSubscribes, len(pSubscribes), q.nextSeq())
	log.Println("Market: Subscription request sent")
}

func (q *Quoter) OnFrontDisconnected(nResult int) {
	log.Printf("Market: Front disconnected, result=%d", nResult)
}

func (q *Quoter) OnRtnDepthMarketData(pDepthMarketData *thost.CHSDepthMarketDataField) {
	instID := string(pDepthMarketData.InstrumentID[:])
	lastPrice := pDepthMarketData.LastPrice
	bid1 := pDepthMarketData.BidPrice1
	bidVol1 := pDepthMarketData.BidVolume1
	ask1 := pDepthMarketData.AskPrice1
	askVol1 := pDepthMarketData.AskVolume1
	log.Printf("Market: %s LastPrice=%.4f Bid1=%.4f(%.0f), Ask1=%.4f(%.0f)",
		instID, lastPrice, bid1, bidVol1, ask1, askVol1)
}

func (q *Quoter) nextSeq() int {
	q.seq++
	return q.seq
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Go2UFT Market Demo (Static Library Version)")
	log.Println("==========================================")

	quoter := NewQuoter()
	quoter.api.RegisterSpi(quoter)

	log.Printf("Connecting to: %s", FrontAddr)
	quoter.api.RegisterFront(FrontAddr)

	log.Printf("API Version: %s", quoter.api.GetApiVersion())

	initCfg := new(thost.CHSInitConfigField)
	initCfg.APICheckVersion = thost.API_STRUCT_CHECK_VERSION
	copy(initCfg.CommLicense[:], LicenseFile)
	ret := quoter.api.Init(initCfg, nil)
	if ret != 0 {
		log.Printf("Market: Init failed with code: %d, %s", ret, quoter.api.GetApiErrorMsg(ret))
	} else {
		log.Println("Market: Waiting for connection...")
	}

	go func() {
		quoter.api.Join()
	}()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	<-exitSignal
	log.Println("Shutting down...")
	quoter.api.ReleaseApi()
	log.Println("Done")
}
