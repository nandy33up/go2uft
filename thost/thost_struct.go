package thost

///初始化配置
type CHSInitConfigField struct {
	/// API结构体版本校验
	APICheckVersion HSCheckVersion
	/// 通讯证书
	CommLicense HSLicenseFile
	/// 安全级别
	SafeLevel HSSafeLevel
	/// 通讯密码
	CommPassword HSLicensePassward
	/// SSL版本
	SslVersion HSSslVersion
	/// SSL证书
	CertLicense HSLicenseFile
	/// SSL密码
	CertPassword HSLicensePassward
	/// 国密配置：证书服务商用户名(资金账号)
	SksUser HSSksUser
	/// 国密配置：证书服务商密码(交易密码)
	SksPassword HSSksPassword
	/// 国密配置：证书服务商pin密码
	SksPin HSSksPin
	/// 国密配置：证书服务商根证书校验
	SksCheckCertFlag HSSksCheckCertFlag
}

///响应信息
type CHSRspInfoField struct {
	/// 错误代码
	ErrorID HSErrorID
	/// 错误提示
	ErrorMsg HSErrorMsg
}

///接入认证
type CHSReqAuthenticateField struct {
	/// 账号
	AccountID HSAccountID
	/// 密码
	Password HSPassword
	/// 客户端id
	AppID HSAppID
	/// 认证码
	AuthCode HSAuthCode
}

///接入认证应答
type CHSRspAuthenticateField struct {
	/// 账号
	AccountID HSAccountID
	/// 客户端ID
	AppID HSAppID
	/// 认证码
	AuthCode HSAuthCode
	/// 客户端ID类型
	AppIDType HSAppIDType
}

///用户信息上报请求
type CHSReqUserSystemInfoField struct {
	/// 客户端登陆时间
	AppLoginTime HSStrTime
	/// 客户端公网地址
	AppPublicAdrr HSIPAddress
	/// 客户端系统信息
	AppSysInfo HSAppSysInfo
	/// 客户端系统信息采集完整度
	AppSysInfoIntegrity HSAppSysInfoIntegrity
	/// 客户端系统信息采集异常标识
	AppAbnormalType HSAppAbnormalType
}

///用户信息上报应答
type CHSRspUserSystemInfoField struct {
}

///客户登录
type CHSReqUserLoginField struct {
	/// 账号
	AccountID HSAccountID
	/// 密码
	Password HSPassword
	/// 投资者端应用类别
	UserApplicationType HSUserApplicationType
	/// 投资者端应用信息
	UserApplicationInfo HSUserApplicationInfo
	/// 投资者Mac地址
	MacAddress HSMacAddress
	/// 投资者IP地址
	IPAddress HSIPAddress
	/// 投资者站点信息
	UserStationInfo HSUserStationInfo
}

///客户登录应答
type CHSRspUserLoginField struct {
	/// 营业部号
	BranchID HSNum
	/// 账号
	AccountID HSAccountID
	/// 投资者姓名
	UserName HSUserName
	/// 交易日
	TradingDay HSDate
	/// 上个交易日
	PreTradingDay HSDate
	/// 账单确认标志
	BillConfirmFlag HSBillConfirmFlag
	/// 会话编码
	SessionID HSSessionID
	/// 投资者端应用类别
	UserApplicationType HSUserApplicationType
	/// 投资者端应用信息
	UserApplicationInfo HSUserApplicationInfo
	/// 风险等级
	RiskLevel HSRiskLevel
	/// 投资者上次登陆的Mac地址
	LastMacAddress HSMacAddress
	/// 投资者上次登陆的IP地址
	LastIPAddress HSIPAddress
	/// 上次登录成功时间
	LastLoginTime HSTime
	/// 郑商所当前时间
	CZCETime HSTime
	/// 大商所当前时间
	DCETime HSTime
	/// 上期所当前时间
	SHFETime HSTime
	/// 中金所当前时间
	CFFEXTime HSTime
	/// 能源中心当前时间
	INETime HSTime
	/// 最大报单引用(返回会话上次最大报单引用)
	MaxOrderRef HSRef
	/// 客户编号
	UserID HSUserID
	/// 广期所当前时间
	GFEXTime HSTime
}

///密码更改
type CHSReqUserPasswordUpdateField struct {
	/// 密码类型
	PasswordType HSPasswordType
	/// 密码
	Password HSPassword
	/// 新密码
	NewPassword HSPassword
}

///密码更改应答
type CHSRspUserPasswordUpdateField struct {
}

///报单录入
type CHSReqOrderInsertField struct {
	/// 报单引用
	OrderRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 报单数量
	OrderVolume HSVolume
	/// 报单指令
	OrderCommand HSOrderCommand
	/// 最小成交量
	MinVolume HSVolume
	/// 止损止盈价
	SpringPrice HSPrice
	/// 互换标志
	SwapOrderFlag HSSwapOrderFlag
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
	/// 席位索引
	SeatIndex HSSeatIndex
}

///报单录入应答
type CHSRspOrderInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 经纪公司报单编码
	BrokerOrderID HSBrokerOrderID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 报单数量
	OrderVolume HSVolume
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单指令
	OrderCommand HSOrderCommand
	/// 报单时间
	InsertTime HSTime
	/// 最小成交量
	MinVolume HSVolume
	/// 止损止盈价
	SpringPrice HSPrice
	/// 互换标志
	SwapOrderFlag HSSwapOrderFlag
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
	/// 席位索引
	SeatIndex HSSeatIndex
}

///撤单
type CHSReqOrderActionField struct {
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 报单操作引用
	OrderActionRef HSRef
}

///撤单应答
type CHSRspOrderActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单时间
	InsertTime HSTime
	/// 报单操作引用
	OrderActionRef HSRef
}

///行权录入
type CHSReqExerciseOrderInsertField struct {
	/// 行权引用
	ExerciseRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单数量
	OrderVolume HSVolume
	/// 行权方式
	ExerciseType HSExerciseType
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 期权行权后生成的头寸是否自动平仓
	CloseFlag HSCloseFlag
}

///行权录入应答
type CHSRspExerciseOrderInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 会话编码
	SessionID HSSessionID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
	/// 行权引用
	ExerciseRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单数量
	OrderVolume HSVolume
	/// 行权方式
	ExerciseType HSExerciseType
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单时间
	InsertTime HSTime
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 期权行权后生成的头寸是否自动平仓
	CloseFlag HSCloseFlag
	/// 经纪公司行权报单编码
	ExerciseBrokerOrderID HSBrokerOrderID
}

///行权撤单
type CHSReqExerciseOrderActionField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
	/// 行权引用
	ExerciseRef HSRef
	/// 会话编码
	SessionID HSSessionID
	/// 行权操作引用
	ExecOrderActionRef HSRef
}

///行权撤单应答
type CHSRspExerciseOrderActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
	/// 行权引用
	ExerciseRef HSRef
	/// 会话编码
	SessionID HSSessionID
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 行权操作引用
	ExecOrderActionRef HSRef
}

///锁定录入
type CHSReqLockInsertField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定类型
	LockType HSLockType
	/// 报单数量
	OrderVolume HSVolume
	/// 锁定引用
	LockRef HSRef
}

///锁定录入应答
type CHSRspLockInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 锁定报单编码
	LockOrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定类型
	LockType HSLockType
	/// 报单数量
	OrderVolume HSVolume
	/// 标的物交易账号
	UnderlyingAccountID HSAccountID
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单时间
	InsertTime HSTime
	/// 锁定引用
	LockRef HSRef
}

///询价录入
type CHSReqForQuoteInsertField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///询价录入应答
type CHSRspForQuoteInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单状态
	OrderStatus HSOrderStatus
}

///报价录入
type CHSReqQuoteInsertField struct {
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方套保标志
	BidHedgeType HSHedgeType
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 买方报单引用
	BidOrderRef HSRef
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方套保标志
	AskHedgeType HSHedgeType
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 卖方委托引用
	AskOrderRef HSRef
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
	/// 报价委托引用
	QuoteRef HSRef
	/// 顶单类型
	TopOrderType HSTopOrderType
	/// 报价报单编码
	QuoteSysID HSOrderSysID
}

///报价录入应答
type CHSRspQuoteInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方套保标志
	BidHedgeType HSHedgeType
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 买方委托引用
	BidOrderRef HSRef
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方套保标志
	AskHedgeType HSHedgeType
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 卖方委托引用
	AskOrderRef HSRef
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
	/// 报价委托引用
	QuoteRef HSRef
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
	/// 买方报单编码
	BidOrderSysID HSOrderSysID
	/// 卖方报单编码
	AskOrderSysID HSOrderSysID
	/// 会话编号
	SessionID HSSessionID
	/// 报单时间
	InsertTime HSTime
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 顶单类型
	TopOrderType HSTopOrderType
	/// 报价报单编码
	QuoteSysID HSOrderSysID
}

///报价撤单
type CHSReqQuoteActionField struct {
	/// 报价报单编码
	QuoteSysID HSOrderSysID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 会话编号
	SessionID HSSessionID
	/// 报价委托引用
	QuoteRef HSRef
	/// 报价操作引用
	QuoteActionRef HSRef
	/// 合约代码
	InstrumentID HSInstrumentID
}

///报价撤单应答
type CHSRspQuoteActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 报价报单编码
	QuoteSysID HSOrderSysID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 会话编号
	SessionID HSSessionID
	/// 报价委托引用
	QuoteRef HSRef
	/// 报价操作引用
	QuoteActionRef HSRef
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单时间
	InsertTime HSTime
	/// 合约代码
	InstrumentID HSInstrumentID
}

///申请组合录入
type CHSReqCombActionInsertField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 组合指令方向
	CombDirection HSCombDirection
	/// 报单数量
	OrderVolume HSVolume
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 买卖方向
	Direction HSDirection
	/// 第二腿投机/套保/备兑类型
	SecondHedgeType HSHedgeType
	/// 组合报单引用
	CombOrderRef HSRef
}

///申请组合录入应答
type CHSRspCombActionInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 组合报单编码
	CombOrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 组合指令方向
	CombDirection HSCombDirection
	/// 报单数量
	OrderVolume HSVolume
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单时间
	InsertTime HSTime
	/// 买卖方向
	Direction HSDirection
	/// 第二腿投机/套保/备兑类型
	SecondHedgeType HSHedgeType
	/// 组合报单引用
	CombOrderRef HSRef
}

///最大报单数量获取
type CHSReqQueryMaxOrderVolumeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单指令
	OrderCommand HSOrderCommand
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 组合持仓编码
	CombPositionID HSCombPositionID
}

///最大报单数量获取应答
type CHSRspQueryMaxOrderVolumeField struct {
	/// 账号
	AccountID HSAccountID
	/// 最大报单量
	MaxOrderVolume HSVolume
	/// 单笔最大报单量
	MaxOrderVolumeUnit HSVolume
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单指令
	OrderCommand HSOrderCommand
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 提示信息
	ErrorMsg HSErrorMsg
}

///可锁定数量获取
type CHSReqQryLockVolumeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定类型
	LockType HSLockType
}

///可锁定数量获取应答
type CHSRspQryLockVolumeField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定类型
	LockType HSLockType
	/// 可用持仓数量
	AvailablePositionVolume HSVolume
}

///可行权数量获取
type CHSReqQueryExerciseVolumeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///可行权数量获取应答
type CHSRspQueryExerciseVolumeField struct {
	/// 账号
	AccountID HSAccountID
	/// 最大报单量
	MaxOrderVolumeUnit HSVolume
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///申请组合最大数量获取
type CHSReqQryCombVolumeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 组合指令方向
	CombDirection HSCombDirection
}

///申请组合最大数量获取应答
type CHSRspQryCombVolumeField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 组合指令方向
	CombDirection HSCombDirection
	/// 单笔最大报单量
	MaxOrderVolumeUnit HSVolume
}

///持仓汇总查询请求
type CHSReqQryPositionField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///持仓汇总查询应答
type CHSRspQryPositionField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 上日持仓数量
	YdPositionVolume HSVolume
	/// 持仓数量
	PositionVolume HSVolume
	/// 今日持仓数量
	TodayPositionVolume HSVolume
	/// 可用持仓数量
	AvailablePositionVolume HSVolume
	/// 今日可用持仓数量
	TodayAvailablePositionVolume HSVolume
	/// 持仓保证金
	PositionMargin HSBalance
	/// 权利金
	Premium HSBalance
	/// 手续费
	Commission HSBalance
	/// 开仓冻结数量
	OpenFrozenVolume HSVolume
	/// 平仓冻结数量
	CloseFrozenVolume HSVolume
	/// 组合数量
	CombVolume HSVolume
	/// 行权冻结数量
	ExerciseFrozenVolume HSVolume
	/// 冻结保证金
	FrozenMargin HSBalance
	/// 行权冻结保证金
	ExerciseFrozenMargin HSBalance
	/// 冻结权利金
	FrozenPremium HSBalance
	/// 冻结手续费
	FrozenCommission HSBalance
	/// 开仓量
	OpenVolume HSVolume
	/// 平仓量
	CloseVolume HSVolume
	/// 开仓金额
	OpenBalance HSBalance
	/// 平仓金额
	CloseBalance HSBalance
	/// 开仓成本
	OpenCost HSBalance
	/// 持仓成本
	PositionCost HSBalance
	/// 持仓盈亏
	PositionProfit HSBalance
	/// 平仓盈亏
	CloseProfit HSBalance
	/// 期权市值
	OptionsMarketValue HSBalance
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// TAS持仓数量
	TASPositionVolume HSVolume
}

///资金账户查询
type CHSReqQryTradingAccountField struct {
}

///资金账户查询应答
type CHSRspQryTradingAccountField struct {
	/// 账号
	AccountID HSAccountID
	/// 上日余额
	YdBalance HSBalance
	/// 上日持仓保证金
	YdPositionMargin HSBalance
	/// 上日资金权益
	YdFundEquity HSBalance
	/// 资金权益
	FundEquity HSBalance
	/// 期权市值
	OptionsMarketValue HSBalance
	/// 总权益
	Equity HSBalance
	/// 可用资金
	AvailableBalance HSBalance
	/// 可取资金
	WithdrawBalance HSBalance
	/// 保证金
	Margin HSBalance
	/// 冻结保证金
	FrozenMargin HSBalance
	/// 行权冻结保证金
	ExerciseFrozenMargin HSBalance
	/// 风险度
	RiskDegree HSRiskDegree
	/// 权利金
	Premium HSBalance
	/// 冻结权利金
	FrozenPremium HSBalance
	/// 手续费
	Commission HSBalance
	/// 冻结手续费
	FrozenCommission HSBalance
	/// 平仓盈亏
	CloseProfit HSBalance
	/// 持仓盈亏
	PositionProfit HSBalance
	/// 平仓盯市盈亏
	CloseProfitByDate HSBalance
	/// 持仓盯市盈亏
	PositionProfitByDate HSBalance
	/// 转入金额
	Deposit HSBalance
	/// 转出金额
	Withdraw HSBalance
	/// 货币质押金额
	FundMortgage HSBalance
	/// 仓单质押金额
	WarrantMortgage HSBalance
	/// 冻结资金
	FrozenBalance HSBalance
	/// 解冻资金
	UnFrozenBalance HSBalance
	/// 币种
	CurrencyID HSCurrencyID
	/// 对冲风险度
	HedgeRiskDegree HSRiskDegree
	/// 上海已用限购额度
	ShUsedBuyQuota HSBalance
	/// 深圳已用限购额度
	SzUsedBuyQuota HSBalance
	/// 上海可用限购额度
	ShAvailableBuyQuota HSBalance
	/// 深圳可用限购额度
	SzAvailableBuyQuota HSBalance
	/// 上日仓单质押金额
	YdWarrantMortgage HSBalance
}

///报单查询
type CHSReqQryOrderField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///历史报单查询
type CHSReqQryHistOrderField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///报单信息
type CHSOrderField struct {
	/// 账号
	AccountID HSAccountID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 经纪公司报单编码
	BrokerOrderID HSBrokerOrderID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 报单数量
	OrderVolume HSVolume
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 成交数量
	TradeVolume HSVolume
	/// 撤单数量
	CancelVolume HSVolume
	/// 成交价格
	TradePrice HSPrice
	/// 交易日
	TradingDay HSDate
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 申报时间
	ReportTime HSTime
	/// 报单指令
	OrderCommand HSOrderCommand
	/// 最小成交量
	MinVolume HSVolume
	/// 止损止盈价
	SpringPrice HSPrice
	/// 互换标志
	SwapOrderFlag HSSwapOrderFlag
	/// 强平原因
	ForceCloseReason HSForceCloseReason
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 报单发起方
	OrderSource HSOrderSource
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
	/// 席位索引
	SeatIndex HSSeatIndex
}

///成交查询
type CHSReqQryTradeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///历史成交查询
type CHSReqQryHistTradeField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///成交信息
type CHSTradeField struct {
	/// 账号
	AccountID HSAccountID
	/// 成交编码
	TradeID HSTradeID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 经纪公司报单编码
	BrokerOrderID HSBrokerOrderID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 成交数量
	TradeVolume HSVolume
	/// 成交价格
	TradePrice HSPrice
	/// 交易日
	TradingDay HSDate
	/// 成交时间
	TradeTime HSTime
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 成交手续费
	TradeCommission HSBalance
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///行权查询
type CHSReqQryExerciseField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
}

///历史行权查询
type CHSReqQryHistExerciseField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
}

///行权信息
type CHSExerciseField struct {
	/// 账号
	AccountID HSAccountID
	/// 行权报单编码
	ExerciseOrderSysID HSOrderSysID
	/// 会话编码
	SessionID HSSessionID
	/// 行权引用
	ExerciseRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 行权方式
	ExerciseType HSExerciseType
	/// 行权数量
	OrderVolume HSVolume
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 报单发起方
	OrderSource HSOrderSource
	/// 期权行权后生成的头寸是否自动平仓
	CloseFlag HSCloseFlag
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 行权预冻结手续费
	ExercisePreFrozenCommission HSBalance
	/// 行权预冻结保证金
	ExercisePreFrozenMargin HSBalance
	/// 经纪公司行权报单编码
	ExerciseBrokerOrderID HSBrokerOrderID
	/// 交易编码
	ExchangeAccountID HSExchangeAccountID
}

///锁定查询
type CHSReqQryLockField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定报单编码
	LockOrderSysID HSOrderSysID
}

///历史锁定查询
type CHSReqQryHistLockField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定报单编码
	LockOrderSysID HSOrderSysID
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
}

///锁定信息
type CHSLockField struct {
	/// 账号
	AccountID HSAccountID
	/// 锁定报单编码
	LockOrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 锁定类型
	LockType HSLockType
	/// 锁定数量
	OrderVolume HSVolume
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 交易日
	TradingDay HSDate
	/// 锁定时间
	InsertTime HSTime
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 报单发起方
	OrderSource HSOrderSource
	/// 锁定引用
	LockRef HSRef
	/// 会话编码
	SessionID HSSessionID
}

///申请组合查询
type CHSReqQryCombActionField struct {
	/// 组合报单编码
	CombOrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///历史申请组合查询
type CHSReqQryHistCombActionField struct {
	/// 组合报单编码
	CombOrderSysID HSOrderSysID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
}

///申请组合信息
type CHSCombActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 组合报单编码
	CombOrderSysID HSOrderSysID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 买卖方向
	Direction HSDirection
	/// 组合指令方向
	CombDirection HSCombDirection
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单数量
	OrderVolume HSVolume
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 交易日
	TradingDay HSDate
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 报单发起方
	OrderSource HSOrderSource
	/// 第二腿投机/套保/备兑类型
	SecondHedgeType HSHedgeType
	/// 组合报单引用
	CombOrderRef HSRef
	/// 会话编码
	SessionID HSSessionID
}

///询价订阅，取消订阅请求
type CHSReqForQuoteField struct {
	ExchangeID HSExchangeID
	InstrumentID HSInstrumentID
}

///询价查询
type CHSReqQryForQuoteField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
}

///询价信息
type CHSForQuoteField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
	/// 交易日
	TradingDay HSDate
	/// 最后更新时间
	UpdateTime HSTime
}

///报价查询
type CHSReqQryQuoteField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报价报单编码
	QuoteSysID HSOrderSysID
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
}

///报价信息
type CHSQuoteField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
	/// 报价状态
	OrderStatus HSOrderStatus
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方套保标志
	BidHedgeType HSHedgeType
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方套保标志
	AskHedgeType HSHedgeType
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 报价委托引用
	QuoteRef HSRef
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
	/// 提示信息
	ErrorMsg HSErrorMsg
	/// 报价报单编码
	QuoteSysID HSOrderSysID
	/// 买方报单编码
	BidOrderSysID HSOrderSysID
	/// 卖方报单编码
	AskOrderSysID HSOrderSysID
	/// 买方委托编号
	BidBrokerOrderID HSBrokerOrderID
	/// 卖方委托编号
	AskBrokerOrderID HSBrokerOrderID
	/// 买方委托引用
	BidOrderRef HSRef
	/// 卖方委托引用
	AskOrderRef HSRef
	/// 会话编号
	SessionID HSSessionID
	/// 交易日
	TradingDay HSDate
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 申报时间
	ReportTime HSTime
	/// 顶单类型
	TopOrderType HSTopOrderType
}

///组合持仓明细查询请求
type CHSReqQryPositionCombineDetailField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
}

///组合持仓明细查询应答
type CHSRspQryPositionCombineDetailField struct {
	/// 账号
	AccountID HSAccountID
	/// 组合持仓编码
	CombPositionID HSCombPositionID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 买卖方向
	Direction HSDirection
	/// 上日持仓数量
	YdPositionVolume HSVolume
	/// 可用持仓数量
	AvailablePositionVolume HSVolume
	/// 组合汇总数量
	TotalCombVolume HSVolume
	/// 拆分汇总数量
	TotalSplitVolume HSVolume
	/// 持仓保证金
	PositionMargin HSBalance
	/// 开仓日期
	OpenDate HSDate
	/// 第二腿投机/套保/备兑类型
	SecondHedgeType HSHedgeType
}

///合约信息查询
type CHSReqQryInstrumentField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///合约信息查询应答
type CHSRspQryInstrumentField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 合约名称
	InstrumentName HSInstrumentName
	/// 合约英文名称
	InstrumentEngName HSInstrumentEngName
	/// 合约品种编码
	ProductID HSProductID
	/// 产品类型
	ProductType HSProductType
	/// 市价单最大报单量
	MaxMarketOrderVolume HSVolume
	/// 市价单最小报单量
	MinMarketOrderVolume HSVolume
	/// 限价单最大报单量
	MaxLimitOrderVolume HSVolume
	/// 限价单最小报单量
	MinLimitOrderVolume HSVolume
	/// 合约数量乘数
	VolumeMultiple HSVolume
	/// 最小变动价位
	PriceTick HSPrice
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 行权价
	ExercisePrice HSPrice
	/// 期权类型
	OptionsType HSOptionsType
	/// 交易标志
	TradingFlag HSTradingFlag
	/// 上市日
	MarketDate HSDate
	/// 到期日
	ExpireDate HSDate
	/// 行权开始日期
	BeginExerciseDate HSDate
	/// 行权截至日期
	EndExerciseDate HSDate
	/// 开始交割日
	BeginDeliveryDate HSDate
	/// 结束交割日
	EndDeliveryDate HSDate
	/// 币种
	CurrencyID HSCurrencyID
	/// 组合类型
	CombType HSCombType
}

///备兑缺口查询请求
type CHSReqQryCoveredShortField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
}

///备兑缺口查询应答
type CHSRspQryCoveredShortField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 备兑锁定数量
	CoveredLockVolume HSVolume
	/// 备兑缺口数量
	CoveredShortVolume HSVolume
	/// 备兑预估缺口数量
	CoveredEstimateShortVolume HSVolume
	/// 消息正文
	MsgContent HSMsgContent
}

///行权指派查询请求
type CHSReqQryExerciseAssignField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 持仓类型
	PositionType HSPositionType
}

///行权指派查询应答
type CHSRspQryExerciseAssignField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 持仓类型
	PositionType HSPositionType
	/// 期权类型
	OptionsType HSOptionsType
	/// 行权价
	ExercisePrice HSPrice
	/// 行权数量
	ExerciseVolume HSVolume
	/// 交割数量
	DeliveryVolume HSVolume
	/// 行权冻结资金
	ExerciseFrozenBalance HSBalance
	/// 结算金额
	SettlementBalance HSBalance
}

///银行转账请求
type CHSReqTransferField struct {
	/// 银行代码
	BankID HSBankID
	/// 转账类型
	TransferType HSTransferType
	/// 发生金额
	OccurBalance HSBalance
	/// 资金密码
	FundPassword HSPassword
	/// 银行密码
	BankPassword HSPassword
	/// 币种
	CurrencyID HSCurrencyID
	/// 转账场景
	TransferOccasion HSOccasion
}

///银行转账应答
type CHSRspTransferField struct {
	/// 转账流水
	TransferSerialID HSNum
	/// 银行代码
	BankID HSBankID
	/// 转账类型
	TransferType HSTransferType
	/// 发生金额
	OccurBalance HSBalance
	/// 币种
	CurrencyID HSCurrencyID
	/// 转账场景
	TransferOccasion HSOccasion
}

///银行转账查询请求
type CHSReqQryTransferField struct {
	/// 银行代码
	BankID HSBankID
	/// 转账流水
	TransferSerialID HSNum
}

///银行转账查询应答
type CHSRspQryTransferField struct {
	/// 转账流水
	TransferSerialID HSNum
	/// 银行代码
	BankID HSBankID
	/// 银行名称
	BankName HSBankName
	/// 业务名称
	BusinessName HSBusinessName
	/// 发生金额
	OccurBalance HSBalance
	/// 后资金额
	PostBalance HSBalance
	/// 转账时间
	TransferTime HSTime
	/// 转账状态
	TransferStatus HSTransferStatus
	/// 转账发起方
	TransferSource HSTransferSource
	/// 备注
	Remarks HSRemarks
	/// 币种
	CurrencyID HSCurrencyID
	/// 交易发起方日期
	OrderSourceDate HSDate
	/// 交易日期
	TradingDay HSDate
	/// 转账场景
	TransferOccasion HSOccasion
}

///银行余额查询请求
type CHSReqQueryBankBalanceField struct {
	/// 银行代码
	BankID HSBankID
	/// 资金密码
	FundPassword HSPassword
	/// 银行密码
	BankPassword HSPassword
	/// 币种
	CurrencyID HSCurrencyID
}

///银行余额查询应答
type CHSRspQueryBankBalanceField struct {
	/// 转账流水
	TransferSerialID HSNum
}

///银行账户查询请求
type CHSReqQueryBankAccountField struct {
	/// 银行代码
	BankID HSBankID
	/// 币种
	CurrencyID HSCurrencyID
}

///银行账户查询应答
type CHSRspQueryBankAccountField struct {
	/// 银行代码
	BankID HSBankID
	/// 银行名称
	BankName HSBankName
	/// 银行账号
	BankAccountID HSBankAccountID
	/// 币种
	CurrencyID HSCurrencyID
}

///多中心资金调拨请求
type CHSReqMultiCentreFundTransField struct {
	/// 币种
	CurrencyID HSCurrencyID
	/// 发生金额
	OccurBalance HSBalance
	/// 调拨方向
	TransDirection HSTransDirection
}

///多中心资金调拨应答
type CHSRspMultiCentreFundTransField struct {
	/// 账号
	AccountID HSAccountID
	/// 币种
	CurrencyID HSCurrencyID
	/// 发生金额
	OccurBalance HSBalance
	/// 调拨方向
	TransDirection HSTransDirection
	/// 接入中心资金调拨流水号
	LocalTransferSerialID HSNum
	/// 对端中心资金调拨流水号
	OppositeTransferSerialID HSNum
}

///客户账单查询请求
type CHSReqQueryBillContentField struct {
	/// 开始日期
	BeginDate HSDate
	/// 结束日期
	EndDate HSDate
}

///客户账单查询应答
type CHSRspQueryBillContentField struct {
	/// 账单内容
	BillContent HSBillContent
}

///客户账单确认请求
type CHSReqBillConfirmField struct {
}

///客户账单确认应答
type CHSRspBillConfirmField struct {
	/// 账号
	AccountID HSAccountID
	/// 账单确认状态
	BillConfirmStatus HSBillConfirmStatus
	/// 确认日期
	ConfirmDate HSDate
	/// 确认时间
	ConfirmTime HSTime
}

///保证金查询请求
type CHSReqQryMarginField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
}

///保证金查询应答
type CHSRspQryMarginField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 多头保证金率
	LongMarginRatio HSRate
	/// 多头保证金
	LongMargin HSBalance
	/// 空头保证金率
	ShortMarginRatio HSRate
	/// 空头保证金
	ShortMargin HSBalance
}

///手续费查询请求
type CHSReqQryCommissionField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 产品类型
	ProductType HSProductType
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
}

///手续费查询应答
type CHSRspQryCommissionField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 产品类型
	ProductType HSProductType
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 多头开仓手续费率
	LongOpenCommissionRatio HSRate
	/// 多头开仓手续费
	LongOpenCommission HSBalance
	/// 多头平仓手续费率
	LongCloseCommissionRatio HSRate
	/// 多头平仓手续费
	LongCloseCommission HSBalance
	/// 多头平今手续费率
	LongCloseTodayCommissionRatio HSRate
	/// 多头平今手续费
	LongCloseTodayCommission HSBalance
	/// 空头开仓手续费率
	ShortOpenCommissionRatio HSRate
	/// 空头开仓手续费
	ShortOpenCommission HSBalance
	/// 空头平仓手续费率
	ShortCloseCommissionRatio HSRate
	/// 空头平仓手续费
	ShortCloseCommission HSBalance
	/// 空头平今手续费率
	ShortCloseTodayCommissionRatio HSRate
	/// 空头平今手续费
	ShortCloseTodayCommission HSBalance
	/// 行权手续费率
	ExerciseCommissionRatio HSRate
	/// 行权手续费
	ExerciseCommission HSBalance
	/// 成交面值比例
	TradeValueRatio HSRate
	/// 股票面值
	StockValue HSBalance
}

///汇率查询请求
type CHSReqQryExchangeRateField struct {
	/// 原币种
	FromCurrencyID HSCurrencyID
	/// 目标币种
	ToCurrencyID HSCurrencyID
}

///汇率查询应答
type CHSRspQryExchangeRateField struct {
	/// 原币种
	FromCurrencyID HSCurrencyID
	/// 目标币种
	ToCurrencyID HSCurrencyID
	/// 汇率
	ExchangeRate HSExchangeRate
	/// 原币种数量单位
	FromCurrencyUnit HSVolume
}

///持仓明细查询请求
type CHSReqQryPositionDetailField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///持仓明细查询应答
type CHSRspQryPositionDetailField struct {
	/// 账号
	AccountID HSAccountID
	/// 成交编码
	TradeID HSTradeID
	/// 开仓日期
	OpenDate HSDate
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 成交数量
	TradeVolume HSVolume
	/// 持仓价格
	PositionPrice HSPrice
	/// 今日平仓数量
	TodayCloseVolume HSVolume
	/// 套利持仓号
	ArbitragePositionID HSArbitPositionID
	/// 腿号
	LegID HSLegID
	/// 成交类型
	TradeType HSTradeType
	/// 持仓保证金
	PositionMargin HSBalance
	/// 组合合约代码
	CombInstrumentID HSInstrumentID
	/// 当前持仓数量
	CurrentPositionVolume HSVolume
	/// 结算价
	SettlementPrice HSPrice
	/// 持仓流水号
	PositionSerialID HSNum
}

///经纪公司配置参数查询请求
type CHSReqQrySysConfigField struct {
}

///经纪公司配置参数查询应答
type CHSRspQrySysConfigField struct {
	/// 配置编号
	ConfigNo HSConfigNo
	/// 配置开关状态
	ConfigValue HSConfigValue
}

///行情订阅，取消订阅请求
type CHSReqDepthMarketDataField struct {
	ExchangeID HSExchangeID
	InstrumentID HSInstrumentID
}

///行情查询请求
type CHSReqQryDepthMarketDataField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///行情信息
type CHSDepthMarketDataField struct {
	/// 交易日
	TradingDay HSDate
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 最新价
	LastPrice HSPrice
	/// 昨结算价
	PreSettlementPrice HSPrice
	/// 昨收盘价
	PreClosePrice HSPrice
	/// 开盘价
	OpenPrice HSPrice
	/// 最高价
	HighestPrice HSPrice
	/// 最低价
	LowestPrice HSPrice
	/// 成交数量
	TradeVolume HSVolume
	/// 成交金额
	TradeBalance HSBalance
	/// 总持量
	OpenInterest HSVolume
	/// 收盘价
	ClosePrice HSPrice
	/// 结算价
	SettlementPrice HSPrice
	/// 涨停板价
	UpperLimitPrice HSPrice
	/// 跌停板价
	LowerLimitPrice HSPrice
	/// 昨虚实度
	PreDelta HSDelta
	/// 今虚实度
	CurrDelta HSDelta
	/// 最后更新时间
	UpdateTime HSTime
	/// 申买价一
	BidPrice1 HSPrice
	/// 申买量一
	BidVolume1 HSVolume
	/// 申卖价一
	AskPrice1 HSPrice
	/// 申卖量一
	AskVolume1 HSVolume
	/// 申买价二
	BidPrice2 HSPrice
	/// 申买量二
	BidVolume2 HSVolume
	/// 申卖价二
	AskPrice2 HSPrice
	/// 申卖量二
	AskVolume2 HSVolume
	/// 申买价三
	BidPrice3 HSPrice
	/// 申买量三
	BidVolume3 HSVolume
	/// 申卖价三
	AskPrice3 HSPrice
	/// 申卖量三
	AskVolume3 HSVolume
	/// 申买价四
	BidPrice4 HSPrice
	/// 申买量四
	BidVolume4 HSVolume
	/// 申卖价四
	AskPrice4 HSPrice
	/// 申卖量四
	AskVolume4 HSVolume
	/// 申买价五
	BidPrice5 HSPrice
	/// 申买量五
	BidVolume5 HSVolume
	/// 申卖价五
	AskPrice5 HSPrice
	/// 申卖量五
	AskVolume5 HSVolume
	/// 平均价格
	AveragePrice HSPrice
	/// 昨持仓量
	PreOpenInterest HSVolume
	/// 合约交易状态
	InstrumentTradeStatus HSInstrumentTradeStatus
	/// 合约实时开仓限制
	OpenRestriction HSOpenRestriction
	/// 基金份额参考净值
	IOPV HSPrice
	/// 动态参考价格
	AuctionPrice HSPrice
	/// 行情发送时间
	SendingTime HSTime
}

///撤单信息
type CHSOrderActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 报单编码
	OrderSysID HSOrderSysID
	/// 经纪公司报单编码
	BrokerOrderID HSBrokerOrderID
	/// 会话编码
	SessionID HSSessionID
	/// 报单引用
	OrderRef HSRef
	/// 撤单引用
	OrderActionRef HSRef
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 开平标志
	OffsetFlag HSOffsetFlag
	/// 投机/套保/备兑类型
	HedgeType HSHedgeType
	/// 报单价格
	OrderPrice HSPrice
	/// 报单数量
	OrderVolume HSVolume
	/// 撤单状态
	OrderActionStatus HSOrderStatus
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 交易日
	TradingDay HSDate
	/// 撤单日期
	ActionDate HSDate
	/// 撤单时间
	ActionTime HSTime
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 报单发起方
	OrderSource HSOrderSource
}

///快速交易与集中交易之间资金调拨请求
type CHSReqFundTransField struct {
	/// 调拨方向
	TransDirection HSTransDirection
	/// 币种
	CurrencyID HSCurrencyID
	/// 发生金额
	OccurBalance HSBalance
}

///快速交易与集中交易之间资金调拨应答
type CHSRspFundTransField struct {
	/// 资金调拨流水
	FundTransSerialID HSNum
}

///资金调拨流水查询请求
type CHSReqQryFundTransField struct {
}

///资金调拨流水查询应答
type CHSRspQryFundTransField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易日期
	TradingDay HSDate
	/// 发生日期
	OccurDate HSDate
	/// 发生时间
	OccurTime HSTime
	/// 转账流水
	TransferSerialID HSNum
	/// 币种
	CurrencyID HSCurrencyID
	/// 发生金额
	OccurBalance HSBalance
	/// 后资金额
	PostBalance HSBalance
}

///客户通知查询请求
type CHSReqQryClientNoticeField struct {
}

///客户通知信息
type CHSClientNoticeField struct {
	/// 账号
	AccountID HSAccountID
	/// 消息正文
	MsgBody HSMsgBody
	/// 消息标题
	MsgTitle HSMsgTitle
	/// 消息类型
	MsgType HSMsgType
	/// 消息日期
	MsgDate HSDate
	/// 消息时间
	MsgTime HSTime
}

///期权标的信息查询请求
type CHSReqQryOptUnderlyField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
}

///期权标的信息查询应答
type CHSRspQryOptUnderlyField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 期权对应标的物类型
	UnderlyingType HSStockType
	/// 币种
	CurrencyID HSCurrencyID
	/// 期权对应标的物状态
	UnderlyingStatus HSStatus
	/// 涨停板价
	UpperLimitPrice HSPrice
	/// 跌停板价
	LowerLimitPrice HSPrice
	/// 期权对应标的交易最高数量
	UnderlyingHighAmount HSVolume
	/// 期权对应标的交易最低数量
	UnderlyingLowAmount HSVolume
	/// 合约数量乘数
	VolumeMultiple HSVolume
	/// 买入单位
	BuyUnit HSNum
	/// 卖出单位
	SellUnit HSNum
	/// 最小价差(厘)
	PriceStep HSNum
	/// 上市日
	MarketDate HSDate
	/// 退市日期
	DelistDate HSDate
	/// 期权对应标的市价交易最高数量
	UnderlyingMarketHighAmount HSVolume
	/// 期权对应标的市价交易最低数量
	UnderlyingMarketLowAmount HSVolume
	/// 标的物名称
	UnderlyingName HSInstrumentName
	/// 市价买入单位
	MarketBuyUnit HSNum
	/// 市价卖出单位
	MarketSellUnit HSNum
	/// 期权对应标的物有效日期
	UnderlyingValidDate HSDate
	/// 更新类型
	UpdateType HSNum
}

///证券行情查询请求
type CHSReqQrySecuDepthMarketField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
}

///证券行情查询应答
type CHSRspQrySecuDepthMarketField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 标的物名称
	UnderlyingName HSInstrumentName
	/// 最新价
	LastPrice HSPrice
	/// 开盘价
	OpenPrice HSPrice
	/// 标的物昨收盘价
	UnderlyingPreClosePrice HSPrice
	/// 最高价
	HighestPrice HSPrice
	/// 最低价
	LowestPrice HSPrice
	/// 成交金额
	TradeBalance HSBalance
	/// 成交数量
	TradeVolume HSVolume
	/// 申买价一
	BidPrice1 HSPrice
	/// 申买量一
	BidVolume1 HSVolume
	/// 申卖价一
	AskPrice1 HSPrice
	/// 申卖量一
	AskVolume1 HSVolume
	/// 申买价二
	BidPrice2 HSPrice
	/// 申买量二
	BidVolume2 HSVolume
	/// 申卖价二
	AskPrice2 HSPrice
	/// 申卖量二
	AskVolume2 HSVolume
	/// 申买价三
	BidPrice3 HSPrice
	/// 申买量三
	BidVolume3 HSVolume
	/// 申卖价三
	AskPrice3 HSPrice
	/// 申卖量三
	AskVolume3 HSVolume
	/// 申买价四
	BidPrice4 HSPrice
	/// 申买量四
	BidVolume4 HSVolume
	/// 申卖价四
	AskPrice4 HSPrice
	/// 申卖量四
	AskVolume4 HSVolume
	/// 申买价五
	BidPrice5 HSPrice
	/// 申买量五
	BidVolume5 HSVolume
	/// 申卖价五
	AskPrice5 HSPrice
	/// 申卖量五
	AskVolume5 HSVolume
}

///交易所状态信息
type CHSExchangeStatusField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 交易所状态
	ExchangeStatus HSExchangeStatus
	/// 申报标志
	TradingFlag HSTradingFlag
}

///合约品种状态信息
type CHSProductStatusField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 产品类型
	ProductType HSProductType
	/// 合约品种代码
	ProductID HSProductID
	/// 申报标志
	TradingFlag HSTradingFlag
	/// 合约品种交易状态
	ProductStatus HSExchangeStatus
}

///可取资金查询请求
type CHSReqQryWithdrawFundField struct {
	/// 币种类别
	CurrencyID HSCurrencyID
	/// 系统节点号
	SysNodeID HSSysnodeID
}

///可取资金查询应答
type CHSRspQryWithdrawFundField struct {
	/// 账号
	AccountID HSAccountID
	/// 币种类别
	CurrencyID HSCurrencyID
	/// 系统节点号
	SysNodeID HSSysnodeID
	/// 系统节点名称
	SysNodeName HSSysnodeName
	/// 可取资金
	WithdrawBalance HSBalance
}

///组合合约查询请求
type CHSReqQryCombInstrumentField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///组合合约查询应答
type CHSRspQryCombInstrumentField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买卖方向
	Direction HSDirection
	/// 第二腿买卖方向
	SecondDirection HSDirection
	/// 组合策略类型
	CombStrategyType HSCombStrategyType
	/// 组合保证金系数
	CombMarginCoeff HSVolume
	/// 优先级
	PriorityLevel HSVolume
}

///席位查询请求
type CHSReqQrySeatIDField struct {
	/// 交易类别
	ExchangeID HSExchangeID
}

///席位查询应答
type CHSRspQrySeatIDField struct {
	/// 交易类别
	ExchangeID HSExchangeID
	/// 席位索引
	SeatIndex HSSeatIndex
	/// 席位号
	SeatID HSSeatID
}

///期权对冲设置请求
type CHSReqOptionSelfCloseField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单数量
	OrderVolume HSVolume
	/// 对冲类型
	SelfCloseType HSSelfCloseType
}

///期权对冲设置应答
type CHSRspOptionSelfCloseField struct {
	/// 资金账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单数量
	OrderVolume HSVolume
	/// 对冲类型
	SelfCloseType HSSelfCloseType
}

///期权对冲设置取消请求
type CHSReqOptionSelfCloseActionField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 对冲类型
	SelfCloseType HSSelfCloseType
}

///期权对冲设置取消应答
type CHSRspOptionSelfCloseActionField struct {
	/// 资金账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 对冲类型
	SelfCloseType HSSelfCloseType
}

///查询期权自对冲请求
type CHSReqQryOptionSelfCloseField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///期权自对冲信息
type CHSOptionSelfCloseField struct {
	/// 资金账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单数量
	OrderVolume HSVolume
	/// 对冲类型
	SelfCloseType HSSelfCloseType
	/// 报单状态
	OrderStatus HSOrderStatus
	/// 对冲设置操作类型
	SetSelfCloseType HSSetSelfCloseType
	/// 错误信息
	ErrorMsg HSErrorMsg
	/// 报单编码
	SelfCloseOrderSysID HSOrderSysID
}

///期权对冲设置结果查询请求
type CHSReqQryOptionSelfCloseResultField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
}

///期权对冲设置结果查询应答
type CHSRspQryOptionSelfCloseResultField struct {
	/// 资金账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 对冲类型
	SelfCloseType HSSelfCloseType
	/// 报单数量
	OrderVolume HSVolume
}

///银行转账信息
type CHSTransferField struct {
	/// 账号
	AccountID HSAccountID
	/// 转账流水
	TransferSerialID HSNum
	/// 银行代码
	BankID HSBankID
	/// 银行名称
	BankName HSBankName
	/// 业务名称
	BusinessName HSBusinessName
	/// 发生金额
	OccurBalance HSBalance
	/// 后资金额
	PostBalance HSBalance
	/// 转账时间
	TransferTime HSTime
	/// 转账状态
	TransferStatus HSTransferStatus
	/// 转账发起方
	TransferSource HSTransferSource
	/// 备注
	Remarks HSRemarks
	/// 币种
	CurrencyID HSCurrencyID
	/// 交易发起方日期
	OrderSourceDate HSDate
	/// 交易日期
	TradingDay HSDate
	/// 转账场景
	TransferOccasion HSOccasion
}

///国密证书信息
type CHSSMCertInfoField struct {
	/// 证书标识
	CertID HSCertID
	/// 操作员账号
	OperatorID HSAccountID
	/// 设备标识
	DeviceID HSDeviceID
	/// 证书信息
	CertInfo HSCertInfo
	/// SKSpin密码
	SksPin HSSksPin
	/// 是否当前的设备标识
	IsCurrent HSIsCurrent
}

///股票期权报价录入请求
type CHSReqOptQuoteInsertField struct {
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 报价委托引用
	QuoteRef HSRef
	/// 询价主场单号
	ForQuoteSysID HSOrderSysID
}

///股票期权报价录入应答
type CHSRspOptQuoteInsertField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
	/// 报单时间
	InsertTime HSTime
	/// 买方报单状态
	BidOrderStatus HSOrderStatus
	/// 卖方报单状态
	AskOrderStatus HSOrderStatus
	/// 报价委托引用
	QuoteRef HSRef
	/// 询价主场单号
	ForQuoteSysID HSOrderSysID
	/// 报价状态
	OrderStatus HSOrderStatus
	/// 会话编号
	SessionID HSSessionID
}

///股票期权报价撤单请求
type CHSReqOptQuoteActionField struct {
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 买方报价撤单数量
	BidWithdrawVolume HSVolume
	/// 卖方报价撤单数量
	AskWithdrawVolume HSVolume
	/// 报价撤单委托引用
	QuoteActionRef HSRef
}

///股票期权报价撤单应答
type CHSRspOptQuoteActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易类别
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报单时间
	InsertTime HSTime
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
	/// 买方报价撤单数量
	BidWithdrawVolume HSVolume
	/// 卖方报价撤单数量
	AskWithdrawVolume HSVolume
	/// 买报单状态
	BidOrderStatus HSOrderStatus
	/// 卖报单状态
	AskOrderStatus HSOrderStatus
	/// 报价撤单委托引用
	QuoteActionRef HSRef
	/// 报价状态
	OrderStatus HSOrderStatus
	/// 会话编号
	SessionID HSSessionID
}

///股票期权报价查询请求
type CHSReqQryOptQuoteField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 查询模式
	QuoteQueryType HSType
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
}

///股票期权报价信息
type CHSOptQuoteField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 报价编号
	QuoteBrokerID HSBrokerOrderID
	/// 买方开平方向
	BidOffsetFlag HSOffsetFlag
	/// 买方报价价格
	BidOrderPrice HSPrice
	/// 买方报价数量
	BidOrderVolume HSVolume
	/// 买方成交价格
	BidTradePrice HSPrice
	/// 买方成交数量
	BidTradeVolume HSVolume
	/// 买方撤单数量
	BidCancelVolume HSVolume
	/// 买方委托编号
	BidBrokerOrderID HSBrokerOrderID
	/// 买方报单状态
	BidOrderStatus HSOrderStatus
	/// 卖方开平方向
	AskOffsetFlag HSOffsetFlag
	/// 卖方报价价格
	AskOrderPrice HSPrice
	/// 卖方报价数量
	AskOrderVolume HSVolume
	/// 卖方成交价格
	AskTradePrice HSPrice
	/// 卖方成交数量
	AskTradeVolume HSVolume
	/// 卖方撤单数量
	AskCancelVolume HSVolume
	/// 卖方委托编号
	AskBrokerOrderID HSBrokerOrderID
	/// 卖方报单状态
	AskOrderStatus HSOrderStatus
	/// 提示信息
	ErrorMsg HSErrorMsg
	/// 交易日
	TradingDay HSDate
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 申报时间
	ReportTime HSTime
	/// 会话编号
	SessionID HSSessionID
	/// 期权对应的标的合约代码
	UnderlyingInstrID HSInstrumentID
	/// 报价委托引用
	QuoteRef HSRef
	/// 询价报单编码
	ForQuoteSysID HSOrderSysID
	/// 报价状态
	OrderStatus HSOrderStatus
}

///股票期权报价撤单信息
type CHSOptQuoteActionField struct {
	/// 账号
	AccountID HSAccountID
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 合约代码
	InstrumentID HSInstrumentID
	/// 交易日
	TradingDay HSDate
	/// 报单日期
	InsertDate HSDate
	/// 报单时间
	InsertTime HSTime
	/// 申报时间
	ReportTime HSTime
	/// 会话编号
	SessionID HSSessionID
	/// 报价撤单委托引用
	QuoteActionRef HSRef
	/// 提示信息
	ErrorMsg HSErrorMsg
}

///股票期权组合策略信息查询请求
type CHSReqQryOptCombStrategyField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
}

///股票期权组合策略信息查询应答
type CHSRspQryOptCombStrategyField struct {
	/// 交易所代码
	ExchangeID HSExchangeID
	/// 组合策略编码
	CombStrategyID HSCombStrategyID
	/// 组合策略名称
	CombStrategyName HSCombStrategyName
	/// 到期日是否相同
	EndDateSameFlag HSFlag
	/// 标的证券是否相同
	UnderlySameFlag HSFlag
	/// 合约单位是否相同
	UnitSameFlag HSFlag
	/// 成分合约个数
	ComponentNum HSNum
	/// 第一腿合约种类
	FirstOptionsType HSOptionsType
	/// 第一腿合约持仓类型
	FirstPositionType HSPositionType
	/// 第一腿合约行权价格顺序
	FirstExercisePriceSeq HSNum
	/// 每份第一腿合约数量
	FirstPerInstrumentAmount HSNum
	/// 第二腿合约种类
	SecondOptionsType HSOptionsType
	/// 第二腿合约持仓类型
	SecondPositionType HSPositionType
	/// 第二腿合约行权价格顺序
	SecondExercisePriceSeq HSNum
	/// 每份第二腿合约数量
	SecondPerInstrumentAmount HSNum
	/// 组合到期提前拆分天数
	NearSplitDays HSNum
	/// 是否使用非标合约
	NonStandardInstrumentFlag HSFlag
}
