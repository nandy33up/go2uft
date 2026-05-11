package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"github.com/nandy33up/go2uft/thost"
	"github.com/nandy33up/go2uft/uft"
)

// 交易配置
const (
	FrontAddr   = "tcp://101.71.12.149:8102"
	LicenseFile = "license.dat"

	// 认证信息（根据实际情况修改）
	AccountID = "556611220556650690"
	Password  = "111111"
	AppID     = "111"
	AuthCode  = "222"

	// 登录信息（根据实际情况修改）
	UserApplicationType = "7"
	UserApplicationInfo = "api"
	MacAddress          = "58A02310C3C2"
	IPAddress           = "192.168.1.2:9006"
	UserStationInfo     = "192.168.1.2:9006"
)

type Trader struct {
	uft.BaseTradeSpi
	api thost.CHSTradeApi
	seq int
}

func NewTrader() *Trader {
	return &Trader{api: uft.CreateTradeApi()}
}

// OnFrontConnected 前置机连接成功回调
func (t *Trader) OnFrontConnected() {
	log.Println("========================================")
	log.Println("Trade: Front connected")
	log.Println("Trade: Sending authenticate request...")

	// 发送认证请求
	authField := &thost.CHSReqAuthenticateField{}
	copy(authField.AccountID[:], AccountID)
	copy(authField.Password[:], Password)
	copy(authField.AppID[:], AppID)
	copy(authField.AuthCode[:], AuthCode)

	ret := t.api.ReqAuthenticate(authField, t.nextSeq())
	if ret != 0 {
		log.Printf("Trade: ReqAuthenticate failed with code: %d, %s", ret, t.api.GetApiErrorMsg(ret))
	}
}

// OnFrontDisconnected 前置机断开回调
func (t *Trader) OnFrontDisconnected(nResult int) {
	log.Printf("Trade: Front disconnected, result=%d, %s", nResult, t.api.GetApiErrorMsg(nResult))
}

// OnRspAuthenticate 认证响应回调
func (t *Trader) OnRspAuthenticate(pRspAuthenticateField *thost.CHSRspAuthenticateField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("----------------------------------------")
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		log.Printf("Trade: Authenticate failed, errorID=%d, errorMsg=%s",
			pRspInfo.ErrorID, thost.BytesToGBK(pRspInfo.ErrorMsg[:]))
		os.Exit(1)
	}
}

// OnRspUserLogin 登录响应回调
func (t *Trader) OnRspUserLogin(pRspUserLoginField *thost.CHSRspUserLoginField, pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("----------------------------------------")
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		log.Printf("Trade: Login failed, errorID=%d, errorMsg=%s",
			pRspInfo.ErrorID, thost.BytesToGBK(pRspInfo.ErrorMsg[:]))
		return
	}

	log.Printf("Trade: Login success!")
	log.Printf("  AccountID: %s", strToArrayToString(pRspUserLoginField.AccountID))
	log.Printf("  TradingDay: %d", pRspUserLoginField.TradingDay)
	log.Printf("  SessionID: %d", pRspUserLoginField.SessionID)
	log.Printf("  UserName: %s", strToArrayToString(pRspUserLoginField.UserName))
	log.Printf("  LastLoginTime: %s", strToArrayToString(pRspUserLoginField.LastLoginTime))

	// TODO: 在这里添加业务逻辑，如查询持仓、报单等
	log.Println("========================================")
	log.Println("Trade: Ready for trading operations")
	log.Println("========================================")
}

// OnRtnOrder 订单回报回调
func (t *Trader) OnRtnOrder(pOrder *thost.CHSOrderField) {
	log.Println("----------------------------------------")
	log.Printf("Trade: Order update")
	log.Printf("  InstrumentID: %s", strToArrayToString(pOrder.InstrumentID))
	log.Printf("  OrderRef: %s", strToArrayToString(pOrder.OrderRef))
	log.Printf("  OrderSysID: %s", strToArrayToString(pOrder.OrderSysID))
	log.Printf("  Direction: %c", pOrder.Direction)
	log.Printf("  OffsetFlag: %c", pOrder.OffsetFlag)
	log.Printf("  OrderPrice: %.4f", pOrder.OrderPrice)
	log.Printf("  OrderVolume: %.0f", pOrder.OrderVolume)
	log.Printf("  TradeVolume: %.0f", pOrder.TradeVolume)
	log.Printf("  CancelVolume: %.0f", pOrder.CancelVolume)
	log.Printf("  OrderStatus: %c", pOrder.OrderStatus)
}

// OnRtnTrade 成交回报回调
func (t *Trader) OnRtnTrade(pTrade *thost.CHSTradeField) {
	log.Println("----------------------------------------")
	log.Printf("Trade: Trade execution")
	log.Printf("  TradeID: %s", strToArrayToString(pTrade.TradeID))
	log.Printf("  InstrumentID: %s", strToArrayToString(pTrade.InstrumentID))
	log.Printf("  OrderRef: %s", strToArrayToString(pTrade.OrderRef))
	log.Printf("  Direction: %c", pTrade.Direction)
	log.Printf("  OffsetFlag: %c", pTrade.OffsetFlag)
	log.Printf("  TradePrice: %.4f", pTrade.TradePrice)
	log.Printf("  TradeVolume: %.0f", pTrade.TradeVolume)
	log.Printf("  TradeTime: %s", strToArrayToString(pTrade.TradeTime))
}

// OnRspError 错误响应回调
func (t *Trader) OnRspError(pRspInfo *thost.CHSRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("----------------------------------------")
	if pRspInfo != nil {
		log.Printf("Trade: Error response, errorID=%d, errorMsg=%s, requestID=%d",
			pRspInfo.ErrorID, thost.BytesToGBK(pRspInfo.ErrorMsg[:]), nRequestID)
	}
}

func (t *Trader) nextSeq() int {
	t.seq++
	return t.seq
}

// strToArrayToString 将固定长度数组转换为字符串
func strToArrayToString(v any) string {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Array {
		return ""
	}
	bs := make([]byte, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		bs[i] = byte(rv.Index(i).Uint())
	}
	return string(bytes.Trim(bs, "\x00"))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("========================================")
	log.Println("Go2UFT Trade Demo (Static Library Version)")
	log.Println("========================================")

	trader := NewTrader()
	trader.api.RegisterSpi(trader)

	// 注册前置机
	log.Printf("Connecting to: %s", FrontAddr)
	trader.api.RegisterFront(FrontAddr)

	// 显示API版本
	log.Printf("API Version: %s", trader.api.GetApiVersion())

	// 初始化配置
	initCfg := new(thost.CHSInitConfigField)
	initCfg.APICheckVersion = thost.API_STRUCT_CHECK_VERSION
	copy(initCfg.CommLicense[:], LicenseFile)

	ret := trader.api.Init(initCfg, nil)
	if ret != 0 {
		log.Printf("Trade: Init failed with code: %d, %s", ret, trader.api.GetApiErrorMsg(ret))
		return
	}

	log.Println("Trade: API initialized successfully, waiting for connection...")

	// 启动后台线程
	go func() {
		trader.api.Join()
	}()

	// 等待信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Println("Shutting down...")
	case <-time.After(30 * time.Second):
		log.Println("Timeout, shutting down...")
	}

	trader.api.ReleaseApi()
	log.Println("Done")
}
