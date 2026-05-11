package thost

import (
	"bytes"
)

// 基础整数类型
type HSInt8 int8
type HSInt16 int16
type HSInt32 int32
type HSInt64 int64
type HSUint8 uint8
type HSUint16 uint16
type HSUint32 uint32
type HSUint64 uint64

// ////////////////////////////////////////////////////////////////////////
// / HSInstrumentID：合约编码
// ////////////////////////////////////////////////////////////////////////
type HSInstrumentID [81]byte

// ////////////////////////////////////////////////////////////////////////
// / HSOptionsType：期权类型
// ////////////////////////////////////////////////////////////////////////
// / 看涨（认购）
const HS_OT_CallOptions HSOptionsType = 'C'

// / 看跌（认沽）
const HS_OT_PutOptions HSOptionsType = 'P'

type HSOptionsType byte

// ////////////////////////////////////////////////////////////////////////
// / HSExerciseStyle：期权行权方式
// ////////////////////////////////////////////////////////////////////////
// / 美式期权
const HS_American_Options HSExerciseStyle = 'A'

// / 欧式期权
const HS_European_Options HSExerciseStyle = 'E'

type HSExerciseStyle byte

// ////////////////////////////////////////////////////////////////////////
// / HSExchangeID：交易所代码
// ////////////////////////////////////////////////////////////////////////
// / 上海证券交易所
var HS_EI_SSE HSExchangeID = [5]byte{'1', 0, 0, 0, 0}

// / 深圳证券交易所
var HS_EI_SZSE HSExchangeID = [5]byte{'2', 0, 0, 0, 0}

// / 北京证券交易所
var HS_EI_BJSE HSExchangeID = [5]byte{'3', 0, 0, 0, 0}

// / 申万指数市场
var HS_EI_SWI HSExchangeID = [5]byte{'8', 0, 0, 0, 0}

// / 全国股转特转A
var HS_EI_TZASE HSExchangeID = [5]byte{'9', 0, 0, 0, 0}

// / 沪港通
var HS_EI_SHHKSE HSExchangeID = [5]byte{'G', 0, 0, 0, 0}

// / 深港通
var HS_EI_SZHKSE HSExchangeID = [5]byte{'S', 0, 0, 0, 0}

// / 郑州商品交易所
var HS_EI_CZCE HSExchangeID = [5]byte{'F', '1', 0, 0, 0}

// / 大连商品交易所
var HS_EI_DCE HSExchangeID = [5]byte{'F', '2', 0, 0, 0}

// / 上海期货交易所
var HS_EI_SHFE HSExchangeID = [5]byte{'F', '3', 0, 0, 0}

// / 中国金融期货交易所
var HS_EI_CFFEX HSExchangeID = [5]byte{'F', '4', 0, 0, 0}

// / 上海国际能源交易中心股份有限公司
var HS_EI_INE HSExchangeID = [5]byte{'F', '5', 0, 0, 0}

// / 广州期货交易所
var HS_EI_GFEX HSExchangeID = [5]byte{'F', '6', 0, 0, 0}

// / 沪港通
var HS_EI_SSEHK HSExchangeID = [5]byte{'G', 0, 0, 0, 0}

// / 深港通
var HS_EI_SZSEHK HSExchangeID = [5]byte{'S', 0, 0, 0, 0}

type HSExchangeID [5]byte

// ////////////////////////////////////////////////////////////////////////
// / HSInstrumentName：合约名字
// ////////////////////////////////////////////////////////////////////////
type HSInstrumentName [129]byte

// ////////////////////////////////////////////////////////////////////////
// / HSVolume：数量(double)
// ////////////////////////////////////////////////////////////////////////
type HSVolume float64

// ////////////////////////////////////////////////////////////////////////
// / HSIntVolume：数量(int64)
// ////////////////////////////////////////////////////////////////////////
type HSIntVolume int64

// ////////////////////////////////////////////////////////////////////////
// / HSNum64：个数
// ////////////////////////////////////////////////////////////////////////
type HSNum64 int64

// ////////////////////////////////////////////////////////////////////////
// / HSArbitPositionID：套利持仓号
// ////////////////////////////////////////////////////////////////////////
type HSArbitPositionID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSLegID：腿号
// ////////////////////////////////////////////////////////////////////////
type HSLegID float64

// ////////////////////////////////////////////////////////////////////////
// / HSConfigNo：配置编号
// ////////////////////////////////////////////////////////////////////////
type HSConfigNo float64

// ////////////////////////////////////////////////////////////////////////
// / HSConfigValue：配置开关状态
// / 根据不同的配置编号，有不同的字典项
// ////////////////////////////////////////////////////////////////////////
type HSConfigValue [128]byte

// ////////////////////////////////////////////////////////////////////////
// / HSExerciseType：期权行权类型
// ////////////////////////////////////////////////////////////////////////
// / 放弃
const HS_ET_Abandon HSExerciseType = '0'

// / 执行
const HS_ET_Exec HSExerciseType = '1'

type HSExerciseType byte

// ////////////////////////////////////////////////////////////////////////
// / HSPrice：价格
// ////////////////////////////////////////////////////////////////////////
type HSPrice float64

// ////////////////////////////////////////////////////////////////////////
// / HSBalance：金额
// ////////////////////////////////////////////////////////////////////////
type HSBalance float64

// ////////////////////////////////////////////////////////////////////////
// / HSSeqNo：序号
// ////////////////////////////////////////////////////////////////////////
type HSSeqNo int64

// ////////////////////////////////////////////////////////////////////////
// / HSDate：日期 格式YYYYMMDD
// ////////////////////////////////////////////////////////////////////////
type HSDate int32

// ////////////////////////////////////////////////////////////////////////
// / HSMsgContent：消息正文
// ////////////////////////////////////////////////////////////////////////
type HSMsgContent [256]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCombStrategyID：组合策略编码
// ////////////////////////////////////////////////////////////////////////
// / 认购牛市价差
var HS_CSI_CNSJC HSCombStrategyID = [9]byte{'C', 'N', 'S', 'J', 'C', 0, 0, 0, 0}

// / 认沽熊市价差
var HS_CSI_PXSJC HSCombStrategyID = [9]byte{'P', 'X', 'S', 'J', 'C', 0, 0, 0, 0}

// / 认沽牛市价差
var HS_CSI_PNSJC HSCombStrategyID = [9]byte{'P', 'N', 'S', 'J', 'C', 0, 0, 0, 0}

// / 认购熊市价差
var HS_CSI_CXSJC HSCombStrategyID = [9]byte{'C', 'X', 'S', 'J', 'C', 0, 0, 0, 0}

// / 跨式空头
var HS_CSI_KS HSCombStrategyID = [9]byte{'K', 'S', 0, 0, 0, 0, 0, 0, 0}

// / 宽跨式空头
var HS_CSI_KKS HSCombStrategyID = [9]byte{'K', 'K', 'S', 0, 0, 0, 0, 0, 0}

// / 普通转备兑
var HS_CSI_ZBD HSCombStrategyID = [9]byte{'Z', 'B', 'D', 0, 0, 0, 0, 0, 0}

// / 备兑转普通
var HS_CSI_ZXJ HSCombStrategyID = [9]byte{'Z', 'X', 'J', 0, 0, 0, 0, 0, 0}

type HSCombStrategyID [9]byte

// ///////////////////////////////////////////////////////////////////////
// / HSRemarks：备注
// ///////////////////////////////////////////////////////////////////////
type HSRemarks [256]byte

// ///////////////////////////////////////////////////////////////////////
// / HSDirection：买卖方向类型
// ///////////////////////////////////////////////////////////////////////
// / 买
const HS_DC_Buy HSDirection = '1'

// / 卖
const HS_DC_Sell HSDirection = '2'

// / 出借
const HS_DC_Loan HSDirection = 'F'

// / 借入
const HS_DC_Borrow HSDirection = 'G'

type HSDirection byte

// ///////////////////////////////////////////////////////////////////////
// / HSOrderDirection：报单方向类型
// ///////////////////////////////////////////////////////////////////////
// / 买
const HS_OD_Buy HSOrderDirection = 1

// / 卖
const HS_OD_Sell HSOrderDirection = 2

// / 新股申购(包括增发、新债申购)
const HS_OD_Apply HSOrderDirection = 3

// / 配股
const HS_OD_MarthSecu HSOrderDirection = 4

// / 债券转股
const HS_OD_BondConv HSOrderDirection = 5

// / 债券回售
const HS_OD_BondCall HSOrderDirection = 6

// / 质押入库
const HS_OD_PledgeIn HSOrderDirection = 7

// / 质押出库
const HS_OD_PledgeOut HSOrderDirection = 8

// / 正回购
const HS_OD_Repos HSOrderDirection = 9

// / 逆回购
const HS_OD_ReverseRepos HSOrderDirection = 10

// / ETF认购
const HS_OD_EtfSubs HSOrderDirection = 11

// / ETF申购
const HS_OD_EtfPur HSOrderDirection = 12

// / ETF赎回
const HS_OD_EtfRed HSOrderDirection = 13

// / 债券回售撤销
const HS_OD_BondCallCan HSOrderDirection = 14

// / LOF认购
const HS_OD_LofSubs HSOrderDirection = 16

// / LOF申购
const HS_OD_LofPur HSOrderDirection = 17

// / LOF赎回
const HS_OD_LofRed HSOrderDirection = 18

// / LOF转托管
const HS_OD_LofTrams HSOrderDirection = 19

// / 分级基金拆分
const HS_OD_FundSplit HSOrderDirection = 20

// / 分级基金合并
const HS_OD_FundMerger HSOrderDirection = 21

// / 分级基金转换
const HS_OD_FundConv HSOrderDirection = 22

// / 货币基金申购
const HS_OD_CurrFundPur HSOrderDirection = 23

// / 货币基金赎回
const HS_OD_CurrFundRed HSOrderDirection = 24

// / 担保品买入
const HS_OD_CreditBuy HSOrderDirection = 33

// / 担保品卖出
const HS_OD_CreditSell HSOrderDirection = 34

// / 担保品提交
const HS_OD_CollateralIn HSOrderDirection = 35

// / 担保品返回
const HS_OD_CollateralOut HSOrderDirection = 36

// / 融资买入
const HS_OD_MarginBuy HSOrderDirection = 37

// / 融券卖出
const HS_OD_ShortSell HSOrderDirection = 38

// / 卖券还款
const HS_OD_SellRepay HSOrderDirection = 39

// / 买券还券
const HS_OD_BuyRepay HSOrderDirection = 40

// / 现券还券
const HS_OD_HoldRepay HSOrderDirection = 41

// / 港股整手买入
const HS_OD_HkBuyRound HSOrderDirection = 42

// / 港股整手卖出
const HS_OD_HkSellRound HSOrderDirection = 43

// / 港股零股卖出
const HS_OD_HkSellOddLot HSOrderDirection = 44

// / 货币基金申购
const HS_OD_CBPOF_OFC HSOrderDirection = 45

// / 货币基金赎回
const HS_OD_CBPOF_OFR HSOrderDirection = 46

// / REITs基金认购
const HS_OD_ReitsFundSubs HSOrderDirection = 47

// / 场内基金认购
const HS_OD_OFundSubs HSOrderDirection = 48

// / 借券卖出
const HS_OD_BorrowSell HSOrderDirection = 49

type HSOrderDirection int32

// ///////////////////////////////////////////////////////////////////////
// / HSOffsetFlag：开平标志类型
// ///////////////////////////////////////////////////////////////////////
// / 开仓
const HS_OF_Open HSOffsetFlag = '1'

// / 平仓
const HS_OF_Close HSOffsetFlag = '2'

// / 交割
const HS_OF_Delivery HSOffsetFlag = '3'

// / 平今仓
const HS_OF_CloseToday HSOffsetFlag = '4'

// / 组合单边平仓
const HS_OF_CloseSingle HSOffsetFlag = '6'

type HSOffsetFlag byte

// ///////////////////////////////////////////////////////////////////////
// / HSCombDirection：组合指令方向类型
// ///////////////////////////////////////////////////////////////////////
// / 组合
const HS_CD_Comb HSCombDirection = '1'

// / 拆分
const HS_CD_UnComb HSCombDirection = '2'

type HSCombDirection byte

// ////////////////////////////////////////////////////////////////////////
// / HSRef：引用
// ////////////////////////////////////////////////////////////////////////
type HSRef [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTime：时间 格式HHMMSSFFF
// ////////////////////////////////////////////////////////////////////////
type HSTime int32

// ////////////////////////////////////////////////////////////////////////
// / HSDurationTime：总耗时
// ////////////////////////////////////////////////////////////////////////
type HSDurationTime int32

// ///////////////////////////////////////////////////////////////////////
// / HSPositionType：持仓类型类型
// ///////////////////////////////////////////////////////////////////////
// / 权利方
const HS_PT_Right HSPositionType = '0'

// / 义务方
const HS_PT_Voluntary HSPositionType = '1'

// / 备兑方
const HS_PT_Covered HSPositionType = '2'

type HSPositionType byte

// ///////////////////////////////////////////////////////////////////////
// / HSLockType：锁定方向类型
// ///////////////////////////////////////////////////////////////////////
// / 锁定
const HS_LT_Lock HSLockType = '1'

// / 解锁
const HS_LT_Unlock HSLockType = '2'

type HSLockType byte

// ///////////////////////////////////////////////////////////////////////
// / HSOrderStatus：报单状态类型
// ///////////////////////////////////////////////////////////////////////
// / 未报
const HS_OS_NotReported HSOrderStatus = '0'

// / 待报
const HS_OS_ToBeReported HSOrderStatus = '1'

// / 已报
const HS_OS_Reported HSOrderStatus = '2'

// / 已报待撤
const HS_OS_ReportedToBeCancel HSOrderStatus = '3'

// / 部成待撤
const HS_OS_PartsTradedToBeCancel HSOrderStatus = '4'

// / 部撤
const HS_OS_CanceledWithPartsTraded HSOrderStatus = '5'

// / 已撤
const HS_OS_Canceled HSOrderStatus = '6'

// / 部成
const HS_OS_PartsTraded HSOrderStatus = '7'

// / 已成
const HS_OS_Traded HSOrderStatus = '8'

// / 废单
const HS_OS_Abandoned HSOrderStatus = '9'

// / 撤废
const HS_OS_CancelFailed HSOrderStatus = 'D'

// / 已确认待撤
const HS_OS_ConfirmedToBeCancel HSOrderStatus = 'U'

// / 已确认
const HS_OS_Confirmed HSOrderStatus = 'V'

// / 待确认
const HS_OS_ToBeConfirmed HSOrderStatus = 'W'

type HSOrderStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSErrorID：错误编码
// ////////////////////////////////////////////////////////////////////////
type HSErrorID int32

// ////////////////////////////////////////////////////////////////////////
// / HSErrorMsg：错误信息
// ////////////////////////////////////////////////////////////////////////
type HSErrorMsg [256]byte

// ///////////////////////////////////////////////////////////////////////
// / HSOrderSource：报单发起方类型
// ///////////////////////////////////////////////////////////////////////
// / 个人投资者发起
const HS_OSRC_Personal HSOrderSource = '0'

// / 交易所发起
const HS_OSRC_Exchange HSOrderSource = '1'

// / 会员发起
const HS_OSRC_Member HSOrderSource = '2'

// / 机构投资者发起
const HS_OSRC_Organization HSOrderSource = '3'

// / 自营交易发起
const HS_OSRC_Proprietary HSOrderSource = '4'

// / 流动性服务提供商发起
const HS_OSRC_MobileServiceProvider HSOrderSource = '5'

type HSOrderSource byte

// ////////////////////////////////////////////////////////////////////////
// / HSTradeID：成交编码
// ////////////////////////////////////////////////////////////////////////
type HSTradeID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSRiskDegree：风险度
// ////////////////////////////////////////////////////////////////////////
type HSRiskDegree float64

// ////////////////////////////////////////////////////////////////////////
// / HSBankID：银行编码
// ////////////////////////////////////////////////////////////////////////
type HSBankID [8]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransferType：转账类型
// ////////////////////////////////////////////////////////////////////////
// / 银行转经纪公司
const HS_TT_BankToBroker HSTransferType = '1'

// / 经纪公司转银行
const HS_TT_BrokerToBank HSTransferType = '2'

type HSTransferType byte

// ////////////////////////////////////////////////////////////////////////
// / HSBankName：银行名字
// ////////////////////////////////////////////////////////////////////////
type HSBankName [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransferSource：转账发起方
// ////////////////////////////////////////////////////////////////////////
// / 经纪公司
const HS_TSRC_Broker HSTransferSource = '0'

// / 银行
const HS_TSRC_Bank HSTransferSource = '1'

// / 互相
const HS_TSRC_Each HSTransferSource = '2'

// / 第三方
const HS_TSRC_Third HSTransferSource = '3'

type HSTransferSource byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransferStatus：转账状态
// ////////////////////////////////////////////////////////////////////////
// / 未报
const HS_TS_NotReported HSTransferStatus = '0'

// / 已报
const HS_TS_Reported HSTransferStatus = '1'

// / 成功
const HS_TS_Success HSTransferStatus = '2'

// / 作废
const HS_TS_Abandoned HSTransferStatus = '3'

// / 待撤
const HS_TS_ReportedToBeCancel HSTransferStatus = '4'

// / 撤销
const HS_TS_Canceled HSTransferStatus = '5'

// / 待冲正
const HS_TS_PendingReversal HSTransferStatus = '7'

// / 已冲正
const HS_TS_Reversal HSTransferStatus = '8'

// / 待报
const HS_TS_ToBeReported HSTransferStatus = 'A'

// / 正报
const HS_TS_Reporting HSTransferStatus = 'P'

// / 已确认
const HS_TS_Confirmed HSTransferStatus = 'Q'

// / 待确定
const HS_TS_PendingConfirm HSTransferStatus = 'x'

type HSTransferStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSBankAccountID：银行账号编码
// ////////////////////////////////////////////////////////////////////////
type HSBankAccountID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAccountID：账号
// ////////////////////////////////////////////////////////////////////////
type HSAccountID [19]byte

// ////////////////////////////////////////////////////////////////////////
// / HSUserID：客户编号
// ////////////////////////////////////////////////////////////////////////
type HSUserID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBillContent：账单内容
// ////////////////////////////////////////////////////////////////////////
type HSBillContent [2000]byte

// ////////////////////////////////////////////////////////////////////////
// / HSHedgeType：投机(交易)/套保/备兑类型
// ////////////////////////////////////////////////////////////////////////
// / 投机(交易)
const HS_HT_Speculation HSHedgeType = '0'

// / 套保
const HS_HT_Hedge HSHedgeType = '1'

// / 套利
const HS_HT_Arbitrage HSHedgeType = '2'

// / 做市商
const HS_HT_MarketMaker HSHedgeType = '3'

// / 备兑
const HS_HT_Covered HSHedgeType = '4'

type HSHedgeType byte

// ////////////////////////////////////////////////////////////////////////
// / HSRate：费率
// ////////////////////////////////////////////////////////////////////////
type HSRate float64

// ////////////////////////////////////////////////////////////////////////
// / HSSessionID：会话编码
// ////////////////////////////////////////////////////////////////////////
type HSSessionID int32

// ////////////////////////////////////////////////////////////////////////
// / HSUserStationInfo：用户站点信息
// ////////////////////////////////////////////////////////////////////////
type HSUserStationInfo [256]byte

// ////////////////////////////////////////////////////////////////////////
// / HSOrderSysID：报单编号
// ////////////////////////////////////////////////////////////////////////
type HSOrderSysID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCombPositionID：组合持仓编码
// ////////////////////////////////////////////////////////////////////////
type HSCombPositionID [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSPassword：密码
// ////////////////////////////////////////////////////////////////////////
type HSPassword [16]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMacAddress：Mac地址
// ////////////////////////////////////////////////////////////////////////
type HSMacAddress [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSIPAddress：IP地址
// ////////////////////////////////////////////////////////////////////////
type HSIPAddress [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBillConfirmFlag：是否需要确认帐单的标志
// ////////////////////////////////////////////////////////////////////////
// / 不需要确认
const HS_BCF_NO HSBillConfirmFlag = '0'

// / 需要确认
const HS_BCF_YES HSBillConfirmFlag = '1'

type HSBillConfirmFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSBillConfirmStatus：账单确认状态
// ////////////////////////////////////////////////////////////////////////
// / 未确认
const HS_BCFD_NO HSBillConfirmStatus = '0'

// / 已确认
const HS_BCFD_YES HSBillConfirmStatus = '1'

type HSBillConfirmStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSDelta：虚实度
// ////////////////////////////////////////////////////////////////////////
type HSDelta float64

// ////////////////////////////////////////////////////////////////////////
// / HSInstrumentTradeStatus：合约交易状态
// ////////////////////////////////////////////////////////////////////////
// / 启动(开市前)
const HS_IT_Init HSInstrumentTradeStatus = 'S'

// / 集合竞价
const HS_IT_CallAuction HSInstrumentTradeStatus = 'C'

// / 连续交易
const HS_IT_Trinding HSInstrumentTradeStatus = 'T'

// / 休市
const HS_IT_Pause HSInstrumentTradeStatus = 'B'

// / 闭市
const HS_IT_Close HSInstrumentTradeStatus = 'E'

// / 收盘集合竞价
const HS_IT_ClosingCallAuction HSInstrumentTradeStatus = 'U'

// / 波动性中断
const HS_IT_Fusing HSInstrumentTradeStatus = 'V'

// / 临时停牌
const HS_IT_Halt HSInstrumentTradeStatus = 'P'

// / 全天停牌
const HS_IT_HaltAllDay HSInstrumentTradeStatus = '1'

// / 熔断(盘中集合竞价)
const HS_IT_FuseToCallAuction HSInstrumentTradeStatus = 'M'

// / 熔断(暂停交易至闭市)
const HS_IT_FuseToClose HSInstrumentTradeStatus = 'N'

// / 盘后交易
const HS_IT_AfterCloseTrade HSInstrumentTradeStatus = 'A'

type HSInstrumentTradeStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSOpenRestriction：合约实时开仓限制
// ////////////////////////////////////////////////////////////////////////
// / 不限制开仓
var HS_OR_NoLimitOpen HSOpenRestriction = [64]byte{'0'}

// / 限制备兑开仓
var HS_OR_LimitCoveredOpen HSOpenRestriction = [64]byte{'1'}

// / 限制卖出开仓
var HS_OR_LimitSellOpen HSOpenRestriction = [64]byte{'2'}

// / 限制卖出开仓、备兑开仓
var HS_OR_LimitSellAndCovOpen HSOpenRestriction = [64]byte{'3'}

// / 限制买入开仓
var HS_OR_LimitBidOpen HSOpenRestriction = [64]byte{'4'}

// / 限制买入开仓、备兑开仓
var HS_OR_LimitBidAndCovOpen HSOpenRestriction = [64]byte{'5'}

// / 限制买入开仓、卖出开仓
var HS_OR_LimitBidAndSellOpen HSOpenRestriction = [64]byte{'6'}

// / 限制买入开仓、备兑开仓、卖出开仓
var HS_OR_LimitBidAndSellAndCovOpen HSOpenRestriction = [64]byte{'7'}

type HSOpenRestriction [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSProductID：合约品种类别
// ////////////////////////////////////////////////////////////////////////
type HSProductID [5]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMaxMarginSideAlgorithm：大额单边保证金算法类型
// ////////////////////////////////////////////////////////////////////////
// / 不使用大额单边保证金算法
const HS_MMSA_NO HSMaxMarginSideAlgorithm = '0'

// / 使用大额单边保证金算法
const HS_MMSA_YES HSMaxMarginSideAlgorithm = '1'

type HSMaxMarginSideAlgorithm byte

// ////////////////////////////////////////////////////////////////////////
// / HSUserName：客户姓名
// ////////////////////////////////////////////////////////////////////////
type HSUserName [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSForceCloseReason：强平原因
// ////////////////////////////////////////////////////////////////////////
// / 非强平
const HS_FCR_NotForceClose HSForceCloseReason = '0'

// / 资金不足
const HS_FCR_LackDeposit HSForceCloseReason = '1'

// / 客户超仓
const HS_FCR_ClientOverPositionLimit HSForceCloseReason = '2'

// / 会员超仓
const HS_FCR_MemberOverPositionLimit HSForceCloseReason = '3'

// / 持仓非整数倍
const HS_FCR_NotMultiple HSForceCloseReason = '4'

// / 违规
const HS_FCR_Violation HSForceCloseReason = '5'

// / 其它
const HS_FCR_Other HSForceCloseReason = '6'

type HSForceCloseReason byte

// ////////////////////////////////////////////////////////////////////////
// / HSTradingFlag：申报标志
// ////////////////////////////////////////////////////////////////////////
// / 禁止申报
const HS_TF_No HSTradingFlag = '0'

// / 可以申报
const HS_TF_Yes HSTradingFlag = '1'

type HSTradingFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSSwapOrderFlag：互换标志
// ////////////////////////////////////////////////////////////////////////
// / 非互换
const HS_SOF_Normal HSSwapOrderFlag = '0'

// / 互换
const HS_SOF_Swap HSSwapOrderFlag = '1'

type HSSwapOrderFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSPasswordType：密码类型
// ////////////////////////////////////////////////////////////////////////
// / 资金密码
const HS_PWDT_Fund HSPasswordType = '1'

// / 交易密码
const HS_PWDT_Trade HSPasswordType = '2'

type HSPasswordType byte

// ////////////////////////////////////////////////////////////////////////
// / HSOrderCommand：报单指令
// ////////////////////////////////////////////////////////////////////////
// / 限价
const HS_CT_Limit HSOrderCommand = 1

// / 限价即时全部成交否则撤销
const HS_CT_LimitFOK HSOrderCommand = 2

// / 限价任意数量即时成交剩余撤销
const HS_CT_LimitFAK HSOrderCommand = 3

// / 限价止损
const HS_CT_LimitStopLoss HSOrderCommand = 4

// / 限价止盈
const HS_CT_LimitStopProfit HSOrderCommand = 5

// / 市价
const HS_CT_Market HSOrderCommand = 6

// / 市价即时全部成交否则撤销
const HS_CT_MarketFOK HSOrderCommand = 7

// / 市价任意数量即时成交剩余撤销
const HS_CT_MarketFAK HSOrderCommand = 8

// / 市价指定成交数量即时成交剩余撤销
const HS_CT_MarketFAKV HSOrderCommand = 9

// / 市价止损
const HS_CT_MarketStopLoss HSOrderCommand = 10

// / 市价止盈
const HS_CT_MarketStopProfit HSOrderCommand = 11

// / 市价即时成交剩余转限价
const HS_CT_MarketToLimit HSOrderCommand = 12

// / 五档市价即时成交剩余撤销
const HS_CT_Market5FAK HSOrderCommand = 13

// / 五档市价即时成交剩余转限价
const HS_CT_Market5ToLimit HSOrderCommand = 14

// / 最优价即时成交剩余转限价
const HS_CT_Market1ToLimit HSOrderCommand = 15

// / 最优价即时成交剩余撤销
const HS_CT_Market1FAK HSOrderCommand = 16

// / 最优价即时全部成交否则撤销
const HS_CT_Market1FOK HSOrderCommand = 17

// / 本方最优价转限价
const HS_CT_MarketSelfToLimit HSOrderCommand = 18

// / 对手方最优价申报
const HS_CT_CounterPartyBest HSOrderCommand = 19

// / 限价指定数量即时成交剩余撤销
const HS_CT_LimitFAKV HSOrderCommand = 20

// / 盘后固定价格
const HS_CT_LimitPFP HSOrderCommand = 21

// / 港股通当日有效竞价限价盘
const HS_CT_HkAtCrossingLimitGFD HSOrderCommand = 22

// / 港股通即时全部成交否则撤销竞价限价盘
const HS_CT_HkAtCrossingLimitFOK HSOrderCommand = 23

// / 港股通增强限价盘
const HS_CT_HkDayLimit HSOrderCommand = 24

// / 结算价交易TAS
const HS_CT_TAS HSOrderCommand = 25

// / 持仓套保确认
const HS_CT_HoldHedgeConfirm HSOrderCommand = 26

// / 跨期套利确认
const HS_CT_SpreadArbitrageConfirm HSOrderCommand = 27

// / 大宗交易定价委托
const HS_CT_BlockPricing HSOrderCommand = 28

// / 大宗交易确认委托
const HS_CT_BlockConfirm HSOrderCommand = 29

// / 大宗交易互报确认委托
const HS_CT_BlockMutualConfirm HSOrderCommand = 30

// / 大宗交易盘后收盘价委托
const HS_CT_BlockAFC HSOrderCommand = 31

// / 大宗交易盘后加权平均价委托
const HS_CT_BlockAFW HSOrderCommand = 32

// / 大宗交易意向委托
const HS_CT_BlockIntention HSOrderCommand = 33

// / 北证及股转限价
const HS_CT_BJSGZLimit HSOrderCommand = 34

// / 限价小节有效
const HS_CT_LimitGIS HSOrderCommand = 35

// / 限价止损小节有效
const HS_CT_LimitStopLossGIS HSOrderCommand = 36

// / 限价止盈小节有效
const HS_CT_LimitStopProfitGIS HSOrderCommand = 37

// / 市价小节有效
const HS_CT_MarketGIS HSOrderCommand = 38

// / 市价止损小节有效
const HS_CT_MarketStopLossGIS HSOrderCommand = 39

// / 市价止盈小节有效
const HS_CT_MarketStopProfitGIS HSOrderCommand = 40

type HSOrderCommand int32

// ////////////////////////////////////////////////////////////////////////
// / HSNum：数字
// ////////////////////////////////////////////////////////////////////////
type HSNum int32

// ////////////////////////////////////////////////////////////////////////
// / HSBrokerOrderID：经纪公司报单编码
// ////////////////////////////////////////////////////////////////////////
type HSBrokerOrderID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBusinessName：业务名称
// ////////////////////////////////////////////////////////////////////////
type HSBusinessName [65]byte

// ////////////////////////////////////////////////////////////////////////
// / HSUserApplicationType：投资者端应用类别
// ////////////////////////////////////////////////////////////////////////
type HSUserApplicationType byte

// ////////////////////////////////////////////////////////////////////////
// / HSUserApplicationInfo：投资者端应用信息
// ////////////////////////////////////////////////////////////////////////
type HSUserApplicationInfo [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSecurityType:市场业务类别
// ////////////////////////////////////////////////////////////////////////
// / 未知类型
const HS_SET_UnKnown HSSecurityType = '0'

// / 股票
const HS_SET_Stock HSSecurityType = '1'

// / 指数
const HS_SET_Index HSSecurityType = '2'

// / 基金
const HS_SET_Fund HSSecurityType = '3'

// / 债券
const HS_SET_Bond HSSecurityType = '4'

// / 个股期权
const HS_SET_Option HSSecurityType = '5'

// / ETF期权
const HS_SET_ETFOption HSSecurityType = '6'

// / 期货
const HS_SET_FUTU HSSecurityType = '7'

// / 期货期权
const HS_SET_FUTUOption HSSecurityType = '8'

type HSSecurityType byte

// ////////////////////////////////////////////////////////////////////////
// / HSSubSecurityType:市场业务详细类别
// ////////////////////////////////////////////////////////////////////////
// / 未知类型
var HS_SSET_UnKnown HSSubSecurityType = [16]byte{'M'}

// / 指数
var HS_SSET_Index HSSubSecurityType = [16]byte{'M', 'R', 'I'}

// / 主板A股
var HS_SSET_S_A HSSubSecurityType = [16]byte{'E', 'S', 'A', '.', 'M'}

// / 主板B股
var HS_SSET_S_B HSSubSecurityType = [16]byte{'E', 'S', 'A', '.', 'B'}

// / 科创板
var HS_SSET_S_Star HSSubSecurityType = [16]byte{'K', 'S', 'H'}

// / 创业板
var HS_SSET_S_Gem HSSubSecurityType = [16]byte{'G', 'E', 'M'}

// / 优先股
var HS_SSET_S_Prefer HSSubSecurityType = [16]byte{'E', 'R'}

// / 其他股票
var HS_SSET_S_OtherStock HSSubSecurityType = [16]byte{'S'}

// / 做市转让
var HS_SSET_S_ZSZR HSSubSecurityType = [16]byte{'Z', 'S', 'Z', 'R'}

// / 竞价转让
var HS_SSET_S_JJZR HSSubSecurityType = [16]byte{'J', 'J', 'Z', 'R'}

// / 两网及退市
var HS_SSET_S_LWTS HSSubSecurityType = [16]byte{'L', 'W', 'T', 'S'}

// / 连续竞价(北交所)
var HS_SSET_S_LXJJ HSSubSecurityType = [16]byte{'L', 'X', 'J', 'J'}

// / 其他转让
var HS_SSET_S_QTZR HSSubSecurityType = [16]byte{'Q', 'T', 'Z', 'R'}

// / 交易型开放式指数基金(ETF)
var HS_SSET_F_ETF HSSubSecurityType = [16]byte{'E', 'M', '.', 'E', 'T', 'F'}

// / 上市开放基金(LOF)
var HS_SSET_F_LOF HSSubSecurityType = [16]byte{'E', 'M', '.', 'L', 'O', 'F'}

// / 封闭式基金
var HS_SSET_F_CEF HSSubSecurityType = [16]byte{'E', 'M', '.', 'C', 'E', 'F'}

// / 分级基金
var HS_SSET_F_SF HSSubSecurityType = [16]byte{'E', 'M', '.', 'S', 'F'}

// / 开放式基金
var HS_SSET_F_OEF HSSubSecurityType = [16]byte{'E', 'M', '.', 'O', 'E', 'F'}

// / 不动产投资信托基金(REITs)
var HS_SSET_F_REITs HSSubSecurityType = [16]byte{'C', 'B', '.', 'R', 'E', 'T'}

// / 其他基金
var HS_SSET_F_OtherFund HSSubSecurityType = [16]byte{'F'}

// / 国债
var HS_SSET_D_GBF HSSubSecurityType = [16]byte{'D', '.', 'G', 'B', 'F'}

// / 企业债
var HS_SSET_D_CBF HSSubSecurityType = [16]byte{'D', '.', 'C', 'B', 'F'}

// / 公司债
var HS_SSET_D_CPF HSSubSecurityType = [16]byte{'D', '.', 'C', 'P', 'F'}

// / 可转债
var HS_SSET_D_CCF HSSubSecurityType = [16]byte{'D', '.', 'C', 'C', 'F'}

// / 债券回购
var HS_SSET_D_REPO HSSubSecurityType = [16]byte{'D', '.', 'R', 'E', 'P', 'O'}

// / 债券预发行
var HS_SSET_D_WIT HSSubSecurityType = [16]byte{'D', '.', 'W', 'I', 'T'}

// / 其他债券
var HS_SSET_D_OtherBond HSSubSecurityType = [16]byte{'D'}

type HSSubSecurityType [16]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMarketBizType:业务类别
// ////////////////////////////////////////////////////////////////////////
// / 未知类型
const HS_MBT_UnKnown HSMarketBizType = '0'

// / 现货
const HS_MBT_Stock HSMarketBizType = '1'

// / 期权
const HS_MBT_Option HSMarketBizType = '2'

// / 期货
const HS_MBT_Future HSMarketBizType = '3'

type HSMarketBizType byte

// ////////////////////////////////////////////////////////////////////////
// / HSOrdType:委托类型
// ////////////////////////////////////////////////////////////////////////
// / 市价
const HS_ORT_Market HSOrdType = '1'

// / 限价
const HS_ORT_Limit HSOrdType = '2'

// / 本方最优
const HS_ORT_MarketSelf HSOrdType = 'U'

// / 增加委托订单
const HS_ORT_Add HSOrdType = 'A'

// / 删除委托订单
const HS_ORT_Delete HSOrdType = 'D'

// / 产品状态订单
const HS_ORT_Status HSOrdType = 'S'

type HSOrdType byte

// ////////////////////////////////////////////////////////////////////////
// / HSTrdType:成交类型
// ////////////////////////////////////////////////////////////////////////
// / 卖
const HS_TRD_Inside HSTrdType = 'S'

// / 买
const HS_TRD_Outside HSTrdType = 'B'

// / 未知
const HS_TRD_UnKnown HSTrdType = 'N'

// / 撤单
const HS_TRD_Cancel HSTrdType = '4'

// / 普通成交
const HS_TRD_Execute HSTrdType = 'F'

type HSTrdType byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransType:逐笔行情类型
// ////////////////////////////////////////////////////////////////////////
// / 逐笔成交和逐笔委托
const HS_TRANS_All HSTransType = '0'

// / 逐笔成交
const HS_TRANS_Trade HSTransType = '1'

// / 逐笔委托
const HS_TRANS_Entrust HSTransType = '2'

type HSTransType byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransFlag:逐笔行情数据标识
// ////////////////////////////////////////////////////////////////////////
// / 未知标识类型
const HS_TRF_UnKnown HSTransFlag = 0

// / 逐笔成交与委托序号独立编号
const HS_TRF_Alone HSTransFlag = 1

// / 逐笔成交与委托序号统一编号
const HS_TRF_Unified HSTransFlag = 2

type HSTransFlag int32

// ////////////////////////////////////////////////////////////////////////
// / HSInstrumentEngName：合约英文名称
// ////////////////////////////////////////////////////////////////////////
type HSInstrumentEngName [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSProductType：产品类型
// ////////////////////////////////////////////////////////////////////////
// / 期货
const HS_PTYPE_Futures HSProductType = '1'

// / 期货期权
const HS_PTYPE_OptFutu HSProductType = '2'

// / 组合
const HS_PTYPE_Combination HSProductType = '3'

// / 即期
const HS_PTYPE_Spot HSProductType = '4'

// / 期转现
const HS_PTYPE_FutuToSpot HSProductType = '5'

// / 证券
const HS_PTYPE_Securities HSProductType = '6'

// / 股票期权
const HS_PTYPE_OptStock HSProductType = '7'

// / TAS合约
const HS_PTYPE_TAS HSProductType = '8'

type HSProductType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCurrencyID：币种
// ////////////////////////////////////////////////////////////////////////
// / 人民币
const HS_CID_CNY HSCurrencyID = '0'

// / 美元
const HS_CID_USD HSCurrencyID = '1'

// / 港币
const HS_CID_HKD HSCurrencyID = '2'

type HSCurrencyID byte

// ////////////////////////////////////////////////////////////////////////
// / HSCloseFlag：期权行权后生成的头寸是否自动平仓标志
// ////////////////////////////////////////////////////////////////////////
// / 不自动平仓
const HS_CF_NO HSCloseFlag = '0'

// / 自动平仓
const HS_CF_YES HSCloseFlag = '1'

type HSCloseFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSCombType：组合类型
// ////////////////////////////////////////////////////////////////////////
// / 普通
const HS_COMBT_PT HSCombType = '0'

// / 跨期
const HS_COMBT_KQ HSCombType = '1'

// / 跨品种
const HS_COMBT_KPZ HSCombType = '2'

type HSCombType byte

// ////////////////////////////////////////////////////////////////////////
// / HSRiskLevel：风险等级
// ////////////////////////////////////////////////////////////////////////
// / 默认型
var HS_RL_Default HSRiskLevel = [4]byte{'0', 0, 0, 0}

// / 保守型
var HS_RL_Keep HSRiskLevel = [4]byte{'1', 0, 0, 0}

// / 谨慎型
var HS_RL_Cautions HSRiskLevel = [4]byte{'2', 0, 0, 0}

// / 稳健型
var HS_RL_Steady HSRiskLevel = [4]byte{'3', 0, 0, 0}

// / 积极型
var HS_RL_Active HSRiskLevel = [4]byte{'4', 0, 0, 0}

// / 成长型
var HS_RL_Growth HSRiskLevel = [4]byte{'6', 0, 0, 0}

// / 专业投资者
var HS_RL_Profession HSRiskLevel = [4]byte{'9', '9', 0, 0}

// / 自定义风险等级
var HS_RL_Diy HSRiskLevel = [4]byte{'1', '0', '0', 0}

type HSRiskLevel [4]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAppIDType：客户端ID类型
// ////////////////////////////////////////////////////////////////////////
// / 直连的投资者
const HS_AT_Investor HSAppIDType = '1'

// / 所有投资者共享一个操作员连接的中继
const HS_AT_OperatorRelay HSAppIDType = '2'

// / 为每个投资者都创建连接的中继
const HS_AT_InvestorRelay HSAppIDType = '3'

// / 未知
const HS_AT_UnKnown HSAppIDType = '0'

type HSAppIDType byte

// ////////////////////////////////////////////////////////////////////////
// / HSAppID：客户端ID
// ////////////////////////////////////////////////////////////////////////
type HSAppID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAuthCode：认证码
// ////////////////////////////////////////////////////////////////////////
type HSAuthCode [128]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAppSysInfo：客户端系统信息
// ////////////////////////////////////////////////////////////////////////
type HSAppSysInfo [2000]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAppSysInfoIntegrity：客户端系统信息采集完整度
// ////////////////////////////////////////////////////////////////////////
type HSAppSysInfoIntegrity [2000]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAppAbnormalType：客户端客户端系统信息采集异常标识类型
// ////////////////////////////////////////////////////////////////////////
type HSAppAbnormalType byte

// ////////////////////////////////////////////////////////////////////////
// / HSTradeType：成交类型
// ////////////////////////////////////////////////////////////////////////
// / 普通成交、套利组合拆仓
const HS_TT_Common HSTradeType = '0'

// / 期权执行
const HS_TT_OptionsExecution HSTradeType = '1'

// / OTC成交
const HS_TT_OTC HSTradeType = '2'

// / 期转现衍生成交
const HS_TT_EFPDerived HSTradeType = '3'

// / 组合衍生成交
const HS_TT_CombinationDerived HSTradeType = '4'

type HSTradeType byte

// ////////////////////////////////////////////////////////////////////////
// / HSTransDirection：调拨方向
// ////////////////////////////////////////////////////////////////////////
// / 调入接入的中心
const HS_TD_In = '0'

// / 从接入中心调出
const HS_TD_Out = '1'

type HSTransDirection byte

// ////////////////////////////////////////////////////////////////////////
// / HSExchangeRate：汇率
// ////////////////////////////////////////////////////////////////////////
type HSExchangeRate float64

// ////////////////////////////////////////////////////////////////////////
// / HSOccasion：场景
// ////////////////////////////////////////////////////////////////////////
type HSOccasion [32]byte

// ////////////////////////////////////////////////////////////////////////
// / 回报订阅模式
// ////////////////////////////////////////////////////////////////////////
type SUB_TERT_TYPE int32

// / 从本交易日开始重传
const HS_TERT_RESTART SUB_TERT_TYPE = 0

// / 从上次收到的续传
const HS_TERT_RESUME SUB_TERT_TYPE = 1

// / 从登录后最新的开始传
const HS_TERT_QUICK SUB_TERT_TYPE = 2

// / 从指定序号开始续传
const HS_TERT_DESIGNATED SUB_TERT_TYPE = 3

// / 不接收主推消息
const HS_TERT_REFUSE SUB_TERT_TYPE = 4

// ////////////////////////////////////////////////////////////////////////
// / HSStatus:状态
// ////////////////////////////////////////////////////////////////////////
type HSStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSType:类别
// ////////////////////////////////////////////////////////////////////////
type HSType byte

// ////////////////////////////////////////////////////////////////////////
// / HSFlag:标志
// ////////////////////////////////////////////////////////////////////////
type HSFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSRisk:风险
// ////////////////////////////////////////////////////////////////////////
type HSRisk float64

// ////////////////////////////////////////////////////////////////////////
// / HSStockType:证券类别
// ////////////////////////////////////////////////////////////////////////
// / 股票
var HS_ST_Stock HSStockType = [5]byte{'0'}

// / 基金
var HS_ST_Fund HSStockType = [5]byte{'1'}

// / 配股权证
var HS_ST_MarthSecu HSStockType = [5]byte{'3'}

// / 普通申购
var HS_ST_Apply HSStockType = [5]byte{'4'}

// / 记账国债
var HS_ST_RegDebt HSStockType = [5]byte{'9'}

// / 基金申赎
var HS_ST_FundAppRed HSStockType = [5]byte{'A'}

// / LOF基金
var HS_ST_LofFund HSStockType = [5]byte{'L'}

// / ETF基金
var HS_ST_EtfFund HSStockType = [5]byte{'T'}

// / ETF申赎
var HS_ST_EtfAppRed HSStockType = [5]byte{'N'}

// / 创业板
var HS_ST_Gem HSStockType = [5]byte{'c'}

// / 科创板股票
var HS_ST_Star HSStockType = [5]byte{'e'}

// / 科创板存托凭证
var HS_ST_StarCdr HSStockType = [5]byte{'g'}

type HSStockType [5]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSubStockType:证券子类别
// ////////////////////////////////////////////////////////////////////////
// / 注册制创业板
var HS_ST_RegGem HSSubStockType = [5]byte{'p'}

// / 注册制创业板CDR
var HS_ST_RegGemCdr HSSubStockType = [5]byte{'q'}

// / 公开优先股
var HS_SST_PublicPreferStock HSSubStockType = [5]byte{'0', '1'}

// / 非公开优先股
var HS_SST_NonPublicPreferStock HSSubStockType = [5]byte{'0', '2'}

// / 新股市值申购
var HS_SST_MarketValueApply HSSubStockType = [5]byte{'4', '1'}

// / 债券信用申购
var HS_SST_BondCreditApply HSSubStockType = [5]byte{'G', '1'}

type HSSubStockType [5]byte

// ////////////////////////////////////////////////////////////////////////
// / HSStockCode:证券代码
// ////////////////////////////////////////////////////////////////////////
type HSStockCode [7]byte

// ////////////////////////////////////////////////////////////////////////
// / HSStockAccount:证券账户
// ////////////////////////////////////////////////////////////////////////
type HSStockAccount [21]byte

// ////////////////////////////////////////////////////////////////////////
// / HSStockName:证券名称
// ////////////////////////////////////////////////////////////////////////
type HSStockName [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCompactID:合约编号
// ////////////////////////////////////////////////////////////////////////
type HSCompactID [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCashGroupProp:头寸属性
// ////////////////////////////////////////////////////////////////////////
// / 普通头寸
const HS_CGP_Common HSCashGroupProp = '1'

// / 专项头寸
const HS_CGP_Special HSCashGroupProp = '2'

type HSCashGroupProp byte

// ////////////////////////////////////////////////////////////////////////
// / HSTradeStatus：成交状态
// ////////////////////////////////////////////////////////////////////////
// / 成交
const HS_TDS_Traded HSTradeStatus = '0'

// / 废单
const HS_TDS_Abandoned HSTradeStatus = '2'

// / 确认
const HS_TDS_Confirmed HSTradeStatus = '4'

type HSTradeStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSEtfStatus：申购赎回状态
// ////////////////////////////////////////////////////////////////////////
// / 全部禁止
const HS_ES_BothForbid HSEtfStatus = '0'

// / 全部允许
const HS_ES_BothAllow HSEtfStatus = '1'

// / 只能申购
const HS_ES_OnlyPurchase HSEtfStatus = '2'

// / 只能赎回
const HS_ES_OnlyRedeem HSEtfStatus = '3'

type HSEtfStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSEtfType：Etf类别
// ////////////////////////////////////////////////////////////////////////
// / 本市ETF
const HS_ET_SingelExchEtf HSEtfType = '1'

// / 跨境ETF
const HS_ET_CrossBorderEtf HSEtfType = '2'

// / 跨市ETF
const HS_ET_CrossExchEtf HSEtfType = '3'

// / 货币ETF
const HS_ET_CurrencyEtf HSEtfType = '4'

// / 实物债券ETF
const HS_ET_PhysicalBondEtf HSEtfType = '5'

// / 商品ETF
const HS_ET_CommodityEtf HSEtfType = '6'

// / 现金债券ETF
const HS_ET_CashBondEtf HSEtfType = '7'

type HSEtfType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCollateralStatus：担保状态
// ////////////////////////////////////////////////////////////////////////
// / 正常
const HS_COS_Normal HSCollateralStatus = '0'

// / 暂停
const HS_COS_Pause HSCollateralStatus = '1'

// / 作废
const HS_COS_Cancel HSCollateralStatus = '2'

type HSCollateralStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSCreditRateType：利率类别
// ////////////////////////////////////////////////////////////////////////
// / 普通融资利率
const HS_CRT_MarginBuy HSCreditRateType = '1'

// / 普通融券利率
const HS_CRT_ShortSell HSCreditRateType = '2'

// / 普通融资罚息
const HS_CRT_MarginBuyFine HSCreditRateType = '3'

// / 普通融券罚息
const HS_CRT_ShortSellFine HSCreditRateType = '4'

// / 专项融资利率
const HS_CRT_SpeMarginBuy HSCreditRateType = '5'

// / 专项融券利率
const HS_CRT_SpeShortSell HSCreditRateType = '6'

// / 专项融资罚息
const HS_CRT_SpeMarginBuyFine HSCreditRateType = '7'

// / 专项融券罚息
const HS_CRT_SpeShortSellFine HSCreditRateType = '8'

type HSCreditRateType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCompactType：合约类别
// ////////////////////////////////////////////////////////////////////////
// / 融资合约
const HS_COT_MarginBuy HSCompactType = '0'

// / 融券合约
const HS_COT_ShortSell HSCompactType = '1'

// / 其他负债
const HS_COT_OtherDebit HSCompactType = '2'

type HSCompactType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCompactSource：合约来源
// ////////////////////////////////////////////////////////////////////////
// / 普通头寸
const HS_COS_Public HSCompactSource = '0'

// / 专项头寸
const HS_COS_Special HSCompactSource = '1'

type HSCompactSource byte

// ////////////////////////////////////////////////////////////////////////
// / HSOrderPartition：报单分区
// ////////////////////////////////////////////////////////////////////////
// / 普通竞价
const HS_OP_Secu HSOrderPartition = 1

// / 普通盘后
const HS_OP_SecuAfof HSOrderPartition = 2

// / 普通综合（证券创业板盘后固定价格业务的报单分区是该值）
const HS_OP_SecuCbp HSOrderPartition = 3

// / 港股通
const HS_OP_SecuHk HSOrderPartition = 4

// / 信用竞价
const HS_OP_Crdt HSOrderPartition = 5

// / 信用盘后
const HS_OP_CrdtAfof HSOrderPartition = 6

// / 信用综合（融资融券创业板盘后固定价格业务的报单分区是该值）
const HS_OP_CrdtCbp HSOrderPartition = 7

// / 普通大宗
const HS_OP_SecuBt HSOrderPartition = 8

type HSOrderPartition int32

// ////////////////////////////////////////////////////////////////////////
// / HSAssetProp：资产属性
// ////////////////////////////////////////////////////////////////////////
type HSAssetProp byte

// ////////////////////////////////////////////////////////////////////////
// / HSSessionNo：会话编号
// ////////////////////////////////////////////////////////////////////////
type HSSessionNo int32

// ////////////////////////////////////////////////////////////////////////
// / HSSysnodeID：系统节点编号
// ////////////////////////////////////////////////////////////////////////
type HSSysnodeID int32

// ////////////////////////////////////////////////////////////////////////
// / HSBatchNo：批次号
// ////////////////////////////////////////////////////////////////////////
type HSBatchNo int32

// ////////////////////////////////////////////////////////////////////////
// / HSOrderID：交易所申报编号
// ////////////////////////////////////////////////////////////////////////
type HSOrderID [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCompactIDStr：合约编号串
// ////////////////////////////////////////////////////////////////////////
type HSCompactIDStr [2001]byte

// ////////////////////////////////////////////////////////////////////////
// / HSDelistFlag：退市标志
// ////////////////////////////////////////////////////////////////////////
// / 正常
const HS_DF_Normal HSDelistFlag = '0'

// / 退市
const HS_DF_Delist HSDelistFlag = '1'

type HSDelistFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSChannelType：通道类型
// ////////////////////////////////////////////////////////////////////////
// / 无关
const HS_CNT_Normal HSChannelType = '0'

// / 场内申赎
const HS_CNT_Of HSChannelType = '1'

// / 盘后申赎
const HS_CNT_Afof HSChannelType = '2'

type HSChannelType byte

// ////////////////////////////////////////////////////////////////////////
// / HSRealActionType：交易变动类型
// ////////////////////////////////////////////////////////////////////////
// / 存管调整
const HS_RA_TodayAdjustment HSRealActionType = '0'

// / 交易冻结
const HS_RA_Frozen HSRealActionType = '1'

// / 回报解冻
const HS_RA_Unfrozen HSRealActionType = '2'

// / 上账
const HS_RA_Deposit HSRealActionType = 'D'

// / 下账
const HS_RA_Withdrawal HSRealActionType = 'W'

// / 锁定
const HS_RA_Lock HSRealActionType = 'L'

// / 解锁
const HS_RA_Unlock HSRealActionType = 'U'

type HSRealActionType byte

// ////////////////////////////////////////////////////////////////////////
// / HSBusinessFlag：业务标志
// ////////////////////////////////////////////////////////////////////////
// / 客户资金调拨
const HS_BF_FundTrans HSBusinessFlag = 2034

// / 客户股份调拨
const HS_BF_StockTrans HSBusinessFlag = 3005

// / 客户ETF持仓调拨
const HS_BF_EtfStockTrans HSBusinessFlag = 3006

// / 融资买入
const HS_BF_FinBuy HSBusinessFlag = 31000

// / 卖券还款
const HS_BF_SellRepay HSBusinessFlag = 31011

// / 自动融资负债归还
const HS_BF_AutoFinRepay HSBusinessFlag = 31014

// / 自动融资费用支付
const HS_BF_PayOff_CashRepay HSBusinessFlag = 31016

// / 融资平仓
const HS_BF_PayOff_FinRepay HSBusinessFlag = 31017

// / 融券卖出
const HS_BF_SloSell HSBusinessFlag = 31100

// / 买券还券
const HS_BF_BuyRepay HSBusinessFlag = 31101

// / 平仓买券还券
const HS_BF_PayOff_BuyRepay HSBusinessFlag = 31102

// / 直接还券
const HS_BF_HoldRepay HSBusinessFlag = 31103

// / 平仓现券还券
const HS_BF_PayOff_HoldRepay HSBusinessFlag = 31104

// / 融券现金了结
const HS_BF_SloCashRepay HSBusinessFlag = 31107

// / 融资融券合约展期状态变更
const HS_BF_CompactApply HSBusinessFlag = 31118

type HSBusinessFlag int32

// ///////////////////////////////////////////////////////////////////////
// / HSOrderAssStatus：报单辅助状态
// ///////////////////////////////////////////////////////////////////////
// / 报单已录入
const HS_OCR_OrderInserted HSOrderAssStatus = '0'

// / 报单已申报
const HS_OCR_OrderReported HSOrderAssStatus = '1'

// / 报单已被交易所确认
const HS_OCR_OrderConfirmed HSOrderAssStatus = '2'

// / 报单被交易所撮合成交
const HS_OCR_OrderTraded HSOrderAssStatus = '3'

// / 报单被交易所废单
const HS_OCR_OrderAbandoned HSOrderAssStatus = '4'

// / 撤单已录入
const HS_OCR_CancelInserted HSOrderAssStatus = '5'

// / 撤单已申报
const HS_OCR_CancelReported HSOrderAssStatus = '6'

// / 撤单已被交易所确认
const HS_OCR_CancelConfirmed HSOrderAssStatus = '7'

// / 撤单被交易所废单
const HS_OCR_CancelFailed HSOrderAssStatus = '8'

type HSOrderAssStatus byte

// ///////////////////////////////////////////////////////////////////////
// / HSStockProperty：股份性质
// ///////////////////////////////////////////////////////////////////////
// / 无限售流通股
var HS_SP_Unlimited HSStockProperty = [3]byte{'0', '0', 0}

// / IPO后限售股
var HS_SP_Limited HSStockProperty = [3]byte{'0', '1', 0}

type HSStockProperty [3]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMsgBody：消息正文
// ////////////////////////////////////////////////////////////////////////
type HSMsgBody [4000]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMsgTitle：消息标题
// ////////////////////////////////////////////////////////////////////////
type HSMsgTitle [256]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMsgType：消息类型
// ////////////////////////////////////////////////////////////////////////
type HSMsgType byte

// ////////////////////////////////////////////////////////////////////////
// / HSStrTime：时间（字符串格式）
// ////////////////////////////////////////////////////////////////////////
type HSStrTime [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTopOrderType：顶单类型
// ////////////////////////////////////////////////////////////////////////
// / 不执行顶单业务
const HS_TOT_UnTopOrder HSTopOrderType = '0'

// / 顶单最近一笔委托
const HS_TOT_LastOrder HSTopOrderType = '1'

// / 顶单指定的委托
const HS_TOT_SpecifiedOrder HSTopOrderType = '2'

type HSTopOrderType byte

// ////////////////////////////////////////////////////////////////////////
// / HSExchangeStatus：交易所状态
// ////////////////////////////////////////////////////////////////////////
// / 连接断开
const HS_EXS_DisConnect HSExchangeStatus = '0'

// / 开盘前
const HS_EXS_BeforeOpening HSExchangeStatus = '1'

// / 集合竞价申报
const HS_EXS_AuctionOrder HSExchangeStatus = '2'

// / 集合竞价平衡
const HS_EXS_AuctionBalance HSExchangeStatus = '3'

// / 集合竞价撮合
const HS_EXS_AuctionMatch HSExchangeStatus = '4'

// / 连续交易
const HS_EXS_Trade HSExchangeStatus = '5'

// / 暂停交易
const HS_EXS_TradePause HSExchangeStatus = '6'

// / 交易闭市
const HS_EXS_TradeClose HSExchangeStatus = '7'

// / 未知状态
const HS_EXS_Unknown HSExchangeStatus = 'X'

type HSExchangeStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSClientOrderID：客户端报单编号
// ////////////////////////////////////////////////////////////////////////
type HSClientOrderID int64

// ////////////////////////////////////////////////////////////////////////
// / HSCompactStatus:合约状态
// ////////////////////////////////////////////////////////////////////////
// / 开仓未归还
const HS_CS_OpenNoret HSCompactStatus = '0'

// / 部分归还
const HS_CS_PartRet HSCompactStatus = '1'

// / 合约已过期
const HS_CS_Overload HSCompactStatus = '2'

// / 客户自行归还
const HS_CS_CustEnd HSCompactStatus = '3'

// / 手工了结
const HS_CS_ManEnd HSCompactStatus = '4'

// / 未形成负债
const HS_CS_NoDebit HSCompactStatus = '5'

type HSCompactStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSExchangeAccountID：交易编码
// ////////////////////////////////////////////////////////////////////////
type HSExchangeAccountID [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCombStrategyType：组合策略类型
// ////////////////////////////////////////////////////////////////////////
// / 基本定单
const HS_CST_NoReport HSCombStrategyType = '0'

// / 套利定单
const HS_CST_Arbitrage HSCombStrategyType = '1'

// / 组合定单
const HS_CST_Comb HSCombStrategyType = '2'

// / 批量定单
const HS_CST_Batch HSCombStrategyType = '3'

// / 期权执行申请定单
const HS_CST_OptExec HSCombStrategyType = '4'

// / 双边报价定单
const HS_CST_Quote HSCombStrategyType = '5'

// / 互换订单
const HS_CST_Swap HSCombStrategyType = '6'

// / 跨期套利
const HS_CST_ArbiSP HSCombStrategyType = '7'

// / 跨品种套利
const HS_CST_ArbiSPC HSCombStrategyType = '8'

// / 期权垂直买权套利
const HS_CST_FOptBUL HSCombStrategyType = '9'

// / 期权垂直卖权套利
const HS_CST_FOptBER HSCombStrategyType = 'A'

// / 期权水平买权套利
const HS_CST_FOptBLT HSCombStrategyType = 'B'

// / 期权水平卖权套利
const HS_CST_FOptBRT HSCombStrategyType = 'C'

// / 期权跨式套利
const HS_CST_FOptSTD HSCombStrategyType = 'D'

// / 期权宽跨式套利
const HS_CST_FOptSTG HSCombStrategyType = 'E'

// / 期权转换CNV
const HS_CST_CNV HSCombStrategyType = 'F'

// / 期权三期跨期套利SPZ
const HS_CST_SPZ HSCombStrategyType = 'G'

// / 期权期货保护PRT
const HS_CST_PRT HSCombStrategyType = 'H'

// / 同合约对锁
const HS_CST_Lock HSCombStrategyType = 'I'

// / 期权日历价差
const HS_CST_FOptCAS HSCombStrategyType = 'J'

// / 买入期货期权
const HS_CST_FOptBFO HSCombStrategyType = 'K'

// / 卖出期货期权
const HS_CST_FOptSFO HSCombStrategyType = 'L'

type HSCombStrategyType byte

// ////////////////////////////////////////////////////////////////////////
// / HSSysnodeName：系统节点名称
// ////////////////////////////////////////////////////////////////////////
type HSSysnodeName [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSQueryType：查询类型
// ////////////////////////////////////////////////////////////////////////
// / 查询全部委托
const HS_QT_All HSQueryType = '0'

// / 查询可撤委托
const HS_QT_Cancel HSQueryType = '1'

type HSQueryType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCheckVersion：API结构体版本校验
// ////////////////////////////////////////////////////////////////////////
type HSCheckVersion int32

// ////////////////////////////////////////////////////////////////////////
// / HSLicenseFile：证书文件
// ////////////////////////////////////////////////////////////////////////
type HSLicenseFile [255]byte

// ////////////////////////////////////////////////////////////////////////
// / HSLicensePassward：证书密码
// ////////////////////////////////////////////////////////////////////////
type HSLicensePassward [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSafeLevel：安全级别
// ////////////////////////////////////////////////////////////////////////
// / 无加密
var HS_SL_None HSSafeLevel = [32]byte{'n', 'o', 'n', 'e', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / 通讯证书加密
var HS_SL_Pwd HSSafeLevel = [32]byte{'p', 'w', 'd', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / SSL加密
var HS_SL_SSL HSSafeLevel = [32]byte{'s', 's', 'l', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HSSafeLevel [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSHkszBusiness：港股非交易业务类别
// ////////////////////////////////////////////////////////////////////////
// / 配股认购申报
var HS_HKSZBS_PGRG HSHkszBusiness = [5]byte{'P', 'G', 'R', 'G', 0}

// / 红利选择权申报
var HS_HKSZBS_QPSB HSHkszBusiness = [5]byte{'Q', 'P', 'S', 'B', 0}

// / 收购保管申报
var HS_HKSZBS_SGBG HSHkszBusiness = [5]byte{'S', 'G', 'B', 'G', 0}

// / 投票申报
var HS_HKSZBS_TPSB HSHkszBusiness = [5]byte{'T', 'P', 'S', 'B', 0}

type HSHkszBusiness [5]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCorbehaviorCode：公司行为代码
// ////////////////////////////////////////////////////////////////////////
type HSCorbehaviorCode [17]byte

// ////////////////////////////////////////////////////////////////////////
// / HSPlacardId：公共编号
// ////////////////////////////////////////////////////////////////////////
type HSPlacardId [21]byte

// ////////////////////////////////////////////////////////////////////////
// / HSMotionId：议案编号
// ////////////////////////////////////////////////////////////////////////
type HSMotionId [21]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAllotNo：订单申请编号
// ////////////////////////////////////////////////////////////////////////
type HSAllotNo [25]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCsdcBusiSerialNo：中登业务流水号
// ////////////////////////////////////////////////////////////////////////
type HSCsdcBusiSerialNo [17]byte

// ////////////////////////////////////////////////////////////////////////
// / HSInfoKind：通知信息类型
// ////////////////////////////////////////////////////////////////////////
// / H01-港股通权益登记通知信息
var HS_IK_H01 HSInfoKind = [4]byte{'H', '0', '1', 0}

// / H05-港股通结算汇兑比率通知信息
var HS_IK_H05 HSInfoKind = [4]byte{'H', '0', '5', 0}

// / H06-港股通投票公告通知信息
var HS_IK_H06 HSInfoKind = [4]byte{'H', '0', '6', 0}

// / H07-港股通现金收购通知信息
var HS_IK_H07 HSInfoKind = [4]byte{'H', '0', '7', 0}

// / H08-港股通股份收购通知信息
var HS_IK_H08 HSInfoKind = [4]byte{'H', '0', '8', 0}

// / H09-港股通现金和股份收购通知信息
var HS_IK_H09 HSInfoKind = [4]byte{'H', '0', '9', 0}

// / H10-股份分拆合并通知信息
var HS_IK_H10 HSInfoKind = [4]byte{'H', '1', '0', 0}

// / H12-港股通投票议案通知信息
var HS_IK_H12 HSInfoKind = [4]byte{'H', '1', '2', 0}

// / H13-港股通公开配售通知信息
var HS_IK_H13 HSInfoKind = [4]byte{'H', '1', '3', 0}

// / H14-港股通供股通知信息
var HS_IK_H14 HSInfoKind = [4]byte{'H', '1', '4', 0}

type HSInfoKind [4]byte

// ////////////////////////////////////////////////////////////////////////
// / HSReturnCode：返回代码
// ////////////////////////////////////////////////////////////////////////
type HSReturnCode [9]byte

// ////////////////////////////////////////////////////////////////////////
// / HSReturnInfo：返回信息
// ////////////////////////////////////////////////////////////////////////
type HSReturnInfo [256]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBusinessType：业务类别
// ////////////////////////////////////////////////////////////////////////
// / '1'-证券配股, 对应港股非交易申报业务:配股认购-'PGRG'
const HS_BT_ZQPG HSBusinessType = '1'

// / '6'-股息入账, 对应港股非交易申报业务:股息入账-'QPSB'
const HS_BT_GXRZ HSBusinessType = '6'

// / 'U'-收购入账, 对应港股非交易申报业务:收购保管-'SGBG'
const HS_BT_SGRZ HSBusinessType = 'U'

// / 'n'-权益数据, 对应港股非交易申报业务:投票申报-'TPSB'
const HS_BT_QYSJ HSBusinessType = 'n'

type HSBusinessType byte

// ////////////////////////////////////////////////////////////////////////
// / HSNoticeNo:提示编号
// ////////////////////////////////////////////////////////////////////////
// / 5-该REITs最近20个交易日累计涨跌幅超过20%或当日涨跌幅超过5%，请注意投资风险
const HS_NOTICE_NO_RISK_OVER_LIMIT HSNoticeNo = 5

type HSNoticeNo int32

// ////////////////////////////////////////////////////////////////////////
// / HSSslVersion：SSL版本
// ////////////////////////////////////////////////////////////////////////
// / SSLV2
var HS_SV_SSLV2 HSSslVersion = [32]byte{'S', 'S', 'L', 'V', '2', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / SSLV3
var HS_SV_SSLV3 HSSslVersion = [32]byte{'S', 'S', 'L', 'V', '3', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / SSLV23
var HS_SV_SSLV23 HSSslVersion = [32]byte{'S', 'S', 'L', 'V', '2', '3', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / TLSV1
var HS_SV_TLSV1 HSSslVersion = [32]byte{'T', 'L', 'S', 'V', '1', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / 格尔国密
var HS_SV_GEER HSSslVersion = [32]byte{'g', 'e', 'r', 'g', 'm', 's', 's', 'l', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// / 信安世纪国密
var HS_SV_XASJ HSSslVersion = [32]byte{'x', 'a', 's', 'j', 'g', 'm', 's', 's', 'l', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HSSslVersion [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSHktTradeLimit：港股通订单交易限制类型
// ////////////////////////////////////////////////////////////////////////
// / 未知
const HS_HKT_UNKNOWN HSHktTradeLimit = '0'

// / 港股通订单限制交易
const HS_HKT_ORDER_LIMIT HSHktTradeLimit = '1'

// / 港股通订单不限制交易
const HS_HKT_ORDER_UNLIMIT HSHktTradeLimit = '2'

type HSHktTradeLimit byte

// ////////////////////////////////////////////////////////////////////////
// / HSMemberID：交易商代码
// ////////////////////////////////////////////////////////////////////////
type HSMemberID [7]byte

// ////////////////////////////////////////////////////////////////////////
// / HSQuoteID：报价消息编号
// ////////////////////////////////////////////////////////////////////////
type HSQuoteID [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSInvestorType：债券交易主体类型
// ////////////////////////////////////////////////////////////////////////
// / 自营
var HS_INT_ZY HSInvestorType = [3]byte{'0', '1', 0}

// / 资管
var HS_INT_ZG HSInvestorType = [3]byte{'0', '2', 0}

// / 机构经纪
var HS_INT_JGJJ HSInvestorType = [3]byte{'0', '3', 0}

// / 个人经纪
var HS_INT_GRJJ HSInvestorType = [3]byte{'0', '4', 0}

type HSInvestorType [3]byte

// ////////////////////////////////////////////////////////////////////////
// / HSInvestorID：债券交易主体代码
// ////////////////////////////////////////////////////////////////////////
type HSInvestorID [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTraderCode：交易员代码
// ////////////////////////////////////////////////////////////////////////
type HSTraderCode [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSettlPeriod：结算周期
// ////////////////////////////////////////////////////////////////////////
type HSSettlPeriod uint16

// ////////////////////////////////////////////////////////////////////////
// / HSSettlType：结算方式
// ////////////////////////////////////////////////////////////////////////
// / 多边净额
const HS_SETTLTYPE_DBJE HSSettlType = 103

// / 逐笔全额
const HS_SETTLTYPE_ZBQE HSSettlType = 104

type HSSettlType uint16

// ////////////////////////////////////////////////////////////////////////
// / HSBondTradeType：债券交易方式
// ////////////////////////////////////////////////////////////////////////
// / 匹配成交
const HS_BTT_PPCJ HSBondTradeType = 1

// / 协商成交
const HS_BTT_XSCJ HSBondTradeType = 2

// / 点击成交
const HS_BTT_DJCJ HSBondTradeType = 3

// / 询价成交
const HS_BTT_XJCJ HSBondTradeType = 4

// / 竞买成交
const HS_BTT_JMCJ HSBondTradeType = 5

// / 意向申报
const HS_BTT_YXSB HSBondTradeType = 6

// / 匹配成交大额
const HS_BTT_PPDE HSBondTradeType = 7

// / 质押式匹配成交
const HS_BTT_ZYPPCJ HSBondTradeType = 8

// / 未知类型
const HS_BTT_UNKNOWN HSBondTradeType = 0

type HSBondTradeType int8

// ////////////////////////////////////////////////////////////////////////
// / HSBondBidExecInstType：债券竞买成交方式
// ////////////////////////////////////////////////////////////////////////
// / 单一主体中标
const HS_BEIT_DYZT HSBondBidExecInstType = 1

// / 多主体单一价格
const HS_BEIT_DZTDYJ HSBondBidExecInstType = 2

// / 多主体多重价格
const HS_BEIT_DZTDCJ HSBondBidExecInstType = 3

// / 未知类型
const HS_BEIT_UNKNOWN HSBondBidExecInstType = 0

type HSBondBidExecInstType int16

// ////////////////////////////////////////////////////////////////////////
// / HSSecondaryOrderID：竞买场次编号
// ////////////////////////////////////////////////////////////////////////
type HSSecondaryOrderID [17]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBondMemo：债券交易备注
// ////////////////////////////////////////////////////////////////////////
type HSBondMemo [161]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBondBidTransType：债券竞买业务类别
// ////////////////////////////////////////////////////////////////////////
// / 竞买预约申报
const HS_BBTT_YYSB HSBondBidTransType = 1

// / 竞买发起申报
const HS_BBTT_FQSB HSBondBidTransType = 2

// / 竞买应价申报
const HS_BBTT_YJSB HSBondBidTransType = 3

// / 未知类型
const HS_BBTT_UNKNOWN HSBondBidTransType = 0

type HSBondBidTransType int16

// ////////////////////////////////////////////////////////////////////////
// / HSRebuildTransType：重建返回逐笔类型
// ////////////////////////////////////////////////////////////////////////
// / 逐笔成交
const HS_RTT_Trade HSRebuildTransType = '1'

// / 逐笔委托
const HS_RTT_Entrust HSRebuildTransType = '2'

// / 债券逐笔成交
const HS_RTT_BOND_Trade HSRebuildTransType = '3'

// / 债券逐笔委托
const HS_RTT_BOND_Entrust HSRebuildTransType = '4'

type HSRebuildTransType byte

// ////////////////////////////////////////////////////////////////////////
// / HSSeatID：席位号
// ////////////////////////////////////////////////////////////////////////
type HSSeatID [17]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSeatIndex：席位索引
// ////////////////////////////////////////////////////////////////////////
type HSSeatIndex uint32

// ////////////////////////////////////////////////////////////////////////
// / HSSelfCloseType：对冲类型
// ////////////////////////////////////////////////////////////////////////
type HSSelfCloseType byte

// ////////////////////////////////////////////////////////////////////////
// / HSSetSelfCloseType：对冲设置操作类型
// ////////////////////////////////////////////////////////////////////////
type HSSetSelfCloseType byte

// ////////////////////////////////////////////////////////////////////////
// / HSCertID：证书标识
// ////////////////////////////////////////////////////////////////////////
type HSCertID [50]byte

// ////////////////////////////////////////////////////////////////////////
// / HSDeviceID：设备标识
// ////////////////////////////////////////////////////////////////////////
type HSDeviceID [100]byte

// ////////////////////////////////////////////////////////////////////////
// / HSCertInfo：证书信息
// ////////////////////////////////////////////////////////////////////////
type HSCertInfo [1024]byte

// ////////////////////////////////////////////////////////////////////////
// / HSIsCurrent：是否暂不返回
// ////////////////////////////////////////////////////////////////////////
type HSIsCurrent int

// ////////////////////////////////////////////////////////////////////////
// / HSSksPin：证书服务商pin密码
// ////////////////////////////////////////////////////////////////////////
type HSSksPin [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSksServer：证书服务商地址
// ////////////////////////////////////////////////////////////////////////
type HSSksServer [64]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSksUser：证书服务商用户名
// ////////////////////////////////////////////////////////////////////////
type HSSksUser [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSksPassword：证书服务商密码
// ////////////////////////////////////////////////////////////////////////
type HSSksPassword [32]byte

// ////////////////////////////////////////////////////////////////////////
// / HSSksCheckCertFlag：证书服务商根证书校验
// ////////////////////////////////////////////////////////////////////////
// / 不进行根证书校验
const HS_SCC_UnCheck HSSksCheckCertFlag = '0'

// / 进行根证书校验
const HS_SCC_Check HSSksCheckCertFlag = '1'

type HSSksCheckCertFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSErrorLog:错误信息
// ////////////////////////////////////////////////////////////////////////
type HSErrorLog [6000]byte

// ////////////////////////////////////////////////////////////////////////
// / HSRelationName:联系人姓名
// ////////////////////////////////////////////////////////////////////////
type HSRelationName [121]byte

// ////////////////////////////////////////////////////////////////////////
// / HSRelationTel:联系人电话
// ////////////////////////////////////////////////////////////////////////
type HSRelationTel [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSPropSeatNo:对方席位号
// ////////////////////////////////////////////////////////////////////////
type HSPropSeatNo [9]byte

// ////////////////////////////////////////////////////////////////////////
// / HSConferId:约定号
// ////////////////////////////////////////////////////////////////////////
type HSConferId [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSNationality:国籍地区
// ////////////////////////////////////////////////////////////////////////
type HSNationality [4]byte

// ////////////////////////////////////////////////////////////////////////
// / HSAddress:地址
// ////////////////////////////////////////////////////////////////////////
type HSAddress [121]byte

// ////////////////////////////////////////////////////////////////////////
// / HSName:姓名
// ////////////////////////////////////////////////////////////////////////
type HSName [65]byte

// ////////////////////////////////////////////////////////////////////////
// / HSOrganName:机构名称
// ////////////////////////////////////////////////////////////////////////
type HSOrganName [61]byte

// ////////////////////////////////////////////////////////////////////////
// / HSGroupName:组别名称
// ////////////////////////////////////////////////////////////////////////
type HSGroupName [13]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTelCode:联系电话
// ////////////////////////////////////////////////////////////////////////
type HSTelCode [21]byte

// ////////////////////////////////////////////////////////////////////////
// / HSTriggerFlag:触警标志
// ////////////////////////////////////////////////////////////////////////
// / 未触警
const HS_TF_NoRisk HSTriggerFlag = '0'

// / 触警
const HS_TF_Alert HSTriggerFlag = '1'

type HSTriggerFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSRiskContent:触警信息
// ////////////////////////////////////////////////////////////////////////
type HSRiskContent [513]byte

// ////////////////////////////////////////////////////////////////////////
// / HSRiskName:风控名称
// ////////////////////////////////////////////////////////////////////////
type HSArcsName [121]byte

// ////////////////////////////////////////////////////////////////////////
// / HsClientName:客户名称
// ////////////////////////////////////////////////////////////////////////
type HSClientName [257]byte

// ////////////////////////////////////////////////////////////////////////
// / HSBTradeconferId:约定号
// ////////////////////////////////////////////////////////////////////////
type HSBTradeconferId [11]byte

// ////////////////////////////////////////////////////////////////////////
// / HSEtfComponentCode:成分股代码
// ////////////////////////////////////////////////////////////////////////
type HSEtfComponentCode [21]byte

// ////////////////////////////////////////////////////////////////////////
// / HSLockStockQuoteID：锁券行情编号(收益互换借券业务)
// ////////////////////////////////////////////////////////////////////////
type HSLockStockQuoteID [33]byte

// ////////////////////////////////////////////////////////////////////////
// / HSLockCompactStatus：锁券合约状态(收益互换借券业务)
// ////////////////////////////////////////////////////////////////////////
// / 合约未生效
const HS_LCS_InEffective HSLockCompactStatus = '0'

// / 合约已生效
const HS_LCS_Effective HSLockCompactStatus = '1'

// / 合约已失效
const HS_LCS_Expired HSLockCompactStatus = '2'

type HSLockCompactStatus byte

// ////////////////////////////////////////////////////////////////////////
// / HSEarlyRecallMode：提前召回模式(收益互换借券业务)
// ////////////////////////////////////////////////////////////////////////
// / 不召回
const HS_ERM_NoRecall HSEarlyRecallMode = '0'

// / T+0
const HS_ERM_Today HSEarlyRecallMode = '1'

// / T+1
const HS_ERM_Morrow HSEarlyRecallMode = '2'

type HSEarlyRecallMode byte

// ////////////////////////////////////////////////////////////////////////
// / HSTickStatusFlag：上海逐笔产品状态订单状态标识
// ////////////////////////////////////////////////////////////////////////
// / 产品未上市
const HS_TSF_ADD HSTickStatusFlag = '1'

// / 启动
const HS_TSF_START HSTickStatusFlag = '2'

// / 开市集合竞价
const HS_TSF_OCALL HSTickStatusFlag = '3'

// / 连续自动撮合
const HS_TSF_TRADE HSTickStatusFlag = '4'

// / 停牌
const HS_TSF_SUSP HSTickStatusFlag = '5'

// / 收盘集合竞价
const HS_TSF_CCALL HSTickStatusFlag = '6'

// / 闭市
const HS_TSF_CLOSE HSTickStatusFlag = '7'

// / 交易结束
const HS_TSF_ENDTR HSTickStatusFlag = '8'

type HSTickStatusFlag byte

// ////////////////////////////////////////////////////////////////////////
// / HSCombStrategyName：组合策略名称
// ////////////////////////////////////////////////////////////////////////
type HSCombStrategyName [128]byte

// String() 方法

func BytesToString(bs []byte) string {
	return string(bytes.TrimRight(bs, "\x00"))
}

func (s HSInstrumentID) String() string {
	return BytesToString(s[:])
}

func (s HSExchangeID) String() string {
	return BytesToString(s[:])
}

func (s HSInstrumentName) String() string {
	return BytesToString(s[:])
}

func (s HSArbitPositionID) String() string {
	return BytesToString(s[:])
}

func (s HSConfigValue) String() string {
	return BytesToString(s[:])
}

func (s HSMsgContent) String() string {
	return BytesToString(s[:])
}

func (s HSCombStrategyID) String() string {
	return BytesToString(s[:])
}

func (s HSRemarks) String() string {
	return BytesToString(s[:])
}

func (s HSRef) String() string {
	return BytesToString(s[:])
}

func (s HSErrorMsg) String() string {
	return BytesToString(s[:])
}

func (s HSTradeID) String() string {
	return BytesToString(s[:])
}

func (s HSBankID) String() string {
	return BytesToString(s[:])
}

func (s HSBankName) String() string {
	return BytesToString(s[:])
}

func (s HSBankAccountID) String() string {
	return BytesToString(s[:])
}

func (s HSAccountID) String() string {
	return BytesToString(s[:])
}

func (s HSUserID) String() string {
	return BytesToString(s[:])
}

func (s HSBillContent) String() string {
	return BytesToString(s[:])
}

func (s HSUserStationInfo) String() string {
	return BytesToString(s[:])
}

func (s HSOrderSysID) String() string {
	return BytesToString(s[:])
}

func (s HSCombPositionID) String() string {
	return BytesToString(s[:])
}

func (s HSPassword) String() string {
	return BytesToString(s[:])
}

func (s HSMacAddress) String() string {
	return BytesToString(s[:])
}

func (s HSIPAddress) String() string {
	return BytesToString(s[:])
}

func (s HSProductID) String() string {
	return BytesToString(s[:])
}

func (s HSUserName) String() string {
	return BytesToString(s[:])
}

func (s HSBrokerOrderID) String() string {
	return BytesToString(s[:])
}

func (s HSBusinessName) String() string {
	return BytesToString(s[:])
}

func (s HSInstrumentEngName) String() string {
	return BytesToString(s[:])
}

func (s HSRiskLevel) String() string {
	return BytesToString(s[:])
}

func (s HSAppID) String() string {
	return BytesToString(s[:])
}

func (s HSAuthCode) String() string {
	return BytesToString(s[:])
}

func (s HSAppSysInfo) String() string {
	return BytesToString(s[:])
}

func (s HSAppSysInfoIntegrity) String() string {
	return BytesToString(s[:])
}

func (s HSOccasion) String() string {
	return BytesToString(s[:])
}

func (s HSStockType) String() string {
	return BytesToString(s[:])
}

func (s HSSubStockType) String() string {
	return BytesToString(s[:])
}

func (s HSStockCode) String() string {
	return BytesToString(s[:])
}

func (s HSStockAccount) String() string {
	return BytesToString(s[:])
}

func (s HSStockName) String() string {
	return BytesToString(s[:])
}

func (s HSCompactID) String() string {
	return BytesToString(s[:])
}

func (s HSSubSecurityType) String() string {
	return BytesToString(s[:])
}

func (s HSCompactIDStr) String() string {
	return BytesToString(s[:])
}

func (s HSStockProperty) String() string {
	return BytesToString(s[:])
}

func (s HSMsgBody) String() string {
	return BytesToString(s[:])
}

func (s HSMsgTitle) String() string {
	return BytesToString(s[:])
}

func (s HSStrTime) String() string {
	return BytesToString(s[:])
}

func (s HSSysnodeName) String() string {
	return BytesToString(s[:])
}

func (s HSLicenseFile) String() string {
	return BytesToString(s[:])
}

func (s HSLicensePassward) String() string {
	return BytesToString(s[:])
}

func (s HSSafeLevel) String() string {
	return BytesToString(s[:])
}

func (s HSHkszBusiness) String() string {
	return BytesToString(s[:])
}

func (s HSCorbehaviorCode) String() string {
	return BytesToString(s[:])
}

func (s HSPlacardId) String() string {
	return BytesToString(s[:])
}

func (s HSMotionId) String() string {
	return BytesToString(s[:])
}

func (s HSAllotNo) String() string {
	return BytesToString(s[:])
}

func (s HSCsdcBusiSerialNo) String() string {
	return BytesToString(s[:])
}

func (s HSInfoKind) String() string {
	return BytesToString(s[:])
}

func (s HSReturnCode) String() string {
	return BytesToString(s[:])
}

func (s HSReturnInfo) String() string {
	return BytesToString(s[:])
}

func (s HSSslVersion) String() string {
	return BytesToString(s[:])
}

func (s HSOrderID) String() string {
	return BytesToString(s[:])
}

func (s HSExchangeAccountID) String() string {
	return BytesToString(s[:])
}

func (s HSUserApplicationInfo) String() string {
	return BytesToString(s[:])
}

func (s HSMemberID) String() string {
	return BytesToString(s[:])
}

func (s HSQuoteID) String() string {
	return BytesToString(s[:])
}

func (s HSInvestorType) String() string {
	return BytesToString(s[:])
}

func (s HSInvestorID) String() string {
	return BytesToString(s[:])
}

func (s HSTraderCode) String() string {
	return BytesToString(s[:])
}

func (s HSSecondaryOrderID) String() string {
	return BytesToString(s[:])
}

func (s HSBondMemo) String() string {
	return BytesToString(s[:])
}

func (s HSSeatID) String() string {
	return BytesToString(s[:])
}

func (s HSCertID) String() string {
	return BytesToString(s[:])
}

func (s HSDeviceID) String() string {
	return BytesToString(s[:])
}

func (s HSCertInfo) String() string {
	return BytesToString(s[:])
}

func (s HSSksPin) String() string {
	return BytesToString(s[:])
}

func (s HSSksServer) String() string {
	return BytesToString(s[:])
}

func (s HSSksUser) String() string {
	return BytesToString(s[:])
}

func (s HSSksPassword) String() string {
	return BytesToString(s[:])
}

func (s HSErrorLog) String() string {
	return BytesToString(s[:])
}

func (s HSRelationName) String() string {
	return BytesToString(s[:])
}

func (s HSRelationTel) String() string {
	return BytesToString(s[:])
}

func (s HSPropSeatNo) String() string {
	return BytesToString(s[:])
}

func (s HSConferId) String() string {
	return BytesToString(s[:])
}

func (s HSNationality) String() string {
	return BytesToString(s[:])
}

func (s HSAddress) String() string {
	return BytesToString(s[:])
}

func (s HSName) String() string {
	return BytesToString(s[:])
}

func (s HSOrganName) String() string {
	return BytesToString(s[:])
}

func (s HSGroupName) String() string {
	return BytesToString(s[:])
}

func (s HSTelCode) String() string {
	return BytesToString(s[:])
}

func (s HSRiskContent) String() string {
	return BytesToString(s[:])
}

func (s HSArcsName) String() string {
	return BytesToString(s[:])
}

func (s HSClientName) String() string {
	return BytesToString(s[:])
}

func (s HSBTradeconferId) String() string {
	return BytesToString(s[:])
}

func (s HSEtfComponentCode) String() string {
	return BytesToString(s[:])
}

func (s HSLockStockQuoteID) String() string {
	return BytesToString(s[:])
}

func (s HSCombStrategyName) String() string {
	return BytesToString(s[:])
}
