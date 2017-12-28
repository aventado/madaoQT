package exchange

import (
	"log"
	"strings"

	"github.com/kataras/golog"
)

/*
	初始化日志句柄
*/
var Logger *golog.Logger

func init() {
	logger := golog.New()
	Logger = logger
	Logger.SetLevel("debug")
	Logger.SetTimeFormat("2006-01-02 06:04:05")
}

/* 交易类型：买，卖，开多，开空，平多，平空 */
type ExchangeType int8
type EventType int8
type TradeType int8
type OrderStatusType int8

const (
	ExchangeTypeFuture ExchangeType = iota
	ExchangeTypeSpot
)

const (
	EventConnected EventType = iota
	EventError
)

const (
	TradeTypeOpenLong TradeType = iota
	TradeTypeOpenShort
	TradeTypeCloseLong
	TradeTypeCloseShort
	TradeTypeBuy
	TradeTypeSell
	TradeTypeCancel
	TradeTypeUnknown
)

var TradeTypeString = map[TradeType]string{
	TradeTypeOpenLong:   "OpenLong",
	TradeTypeOpenShort:  "OpenShort",
	TradeTypeCloseLong:  "CloseLong",
	TradeTypeCloseShort: "CloseShort",
	TradeTypeBuy:        "Buy",
	TradeTypeSell:       "Sell",
	TradeTypeCancel:     "cancel",
	TradeTypeUnknown:    "Unknown_TradeType",
}

const (
	OrderStatusOpen OrderStatusType = iota
	OrderStatusPartDone
	OrderStatusDone
	OrderStatusCanceling
	OrderStatusCanceled
	OrderStatusUnknown
)

var OrderStatusString = map[OrderStatusType]string{
	OrderStatusOpen:      "Open",
	OrderStatusPartDone:  "PartDone",
	OrderStatusDone:      "Done",
	OrderStatusCanceling: "Canceling",
	OrderStatusCanceled:  "Canceled",
	OrderStatusUnknown:   "Unknown_OrderStatus",
}

type InitConfig struct {
	Api    string
	Secret string
	Ticker ITicker
	Custom map[string]interface{}
}

type ITicker interface {
	Ticker(exchange string, value TickerValue)
}

type TickerListItem struct {
	Tag    string // 用户调用者匹配
	Name   string // 用户交易所匹配
	Time   string
	Period string // 合约周期
	Value  interface{}

	ticket    int64
	oldticket int64
}

type DepthListItem struct {
	Coin       string
	Name       string
	Time       string
	Depth      string
	AskAverage float64
	AskQty     float64
	BidAverage float64
	BidQty     float64
	AskByOrder float64
	BidByOrder float64
}

type DepthValue struct {
	Time       string
	AskAverage float64
	AskQty     float64
	BidAverage float64
	BidQty     float64
	AskByOrder float64 // 下单深度均价
	AskPrice   float64 // 下单价格
	BidByOrder float64
	BidPrice   float64

	LimitTradePrice  float64
	LimitTradeAmount float64
}

type TickerValue struct {
	High   float64
	Low    float64
	Volume float64
	Last   float64
	Time   string
	Period string // 合约周期
}

type TradeConfig struct {
	Batch string
	Coin  string

	/* buy or sell */
	Type   TradeType
	Price  float64
	Amount float64
	Limit  float64
}

type OrderInfo struct {
	Pair       string
	OrderID    string
	Price      float64
	Amount     float64
	DealAmount float64
	AvgPrice   float64
	Type       TradeType
	Status     OrderStatusType
}

type TradeResult struct {
	Error   error
	OrderID string
}

/* 获取深度价格 */
type DepthPrice struct {
	price float64
	qty   float64
}

type IExchange interface {
	GetExchangeName() string
	Init(config InitConfig)
	Start()
	Close()

	// AddTicker(coinA string, coinB string, config interface{}, tag string)
	GetTickerValue(tag string) *TickerValue
	WatchEvent() chan EventType
	GetDepthValue(coin string, price float64, limit float64, orderQuantity float64, tradeType TradeType) *DepthValue
	GetBalance() map[string]interface{}
	Trade(configs TradeConfig) *TradeResult

	CancelOrder(order OrderInfo) *TradeResult
	GetOrderInfo(filter OrderInfo) []OrderInfo
}

func RevertTradeType(tradeType TradeType) TradeType {
	switch tradeType {
	case TradeTypeOpenLong:
		return TradeTypeCloseLong
	case TradeTypeOpenShort:
		return TradeTypeCloseShort
	// case TradeTypeCloseLong:
	// case TradeTypeCloseShort:
	case TradeTypeBuy:
		return TradeTypeSell
	case TradeTypeSell:
		return TradeTypeBuy
	}

	return TradeTypeUnknown
}

func RevertDepthArray(array []DepthPrice) []DepthPrice {
	var tmp DepthPrice
	var length int

	if len(array)%2 != 0 {
		length = len(array) / 2
	} else {
		length = len(array)/2 - 1
	}
	for i := 0; i <= length; i++ {
		tmp = array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = tmp

	}
	return array
}

/*
	实际意义不大
*/
func GetDepthAveragePrice(items []DepthPrice) (float64, float64) {

	if items == nil || len(items) == 0 {
		return -1, -1
	}

	var total float64
	var quantity float64

	for _, item := range items {
		total += item.price * item.qty
		quantity += item.qty
	}

	return total / quantity, quantity
}

/*
	返回：（下单均价，下单价格）
*/
func GetDepthPriceByOrder(items []DepthPrice, orderQty float64) (float64, float64) {
	if items == nil || len(items) == 0 {
		return -1, -1
	}

	// log.Printf("Depth:%v", items)
	var total float64
	var amount float64
	for _, item := range items {
		total += item.qty
		amount += (item.qty * item.price)
	}

	if orderQty > total {
		log.Printf("深度不够：%v", total)
		return amount / total, -2

	}

	var depth int
	balance := orderQty

	for i, item := range items {
		if balance-item.qty <= 0 {
			depth = i
			break
		} else {
			balance -= item.qty
		}
	}

	total = 0
	for i := 0; i < depth; i++ {
		total += items[i].price * items[i].qty
	}

	total += (items[depth].price * balance)

	return total / orderQty, items[depth].price
}

/*
	返回：（下单价格，下单数量）
*/
func GetDepthPriceByPrice(items []DepthPrice, price float64, limit float64, quantity float64) (float64, float64) {
	if items == nil || len(items) == 0 || limit <= 0 {
		return -1, -1
	}

	limitPriceHigh := price * (1 + limit)
	limitPriceLow := price * (1 - limit)
	Logger.Debugf("有效价格范围：%v-%v", limitPriceLow, limitPriceHigh)

	var tradePrice float64
	var tradeQuantity float64
	for _, item := range items {
		if item.price >= limitPriceLow && item.price <= limitPriceHigh {

			tradePrice = item.price
			tradeQuantity += item.qty

			if tradeQuantity > quantity {
				tradeQuantity = quantity
				break
			}

		} else {
			Logger.Debugf("超出价格范围")
			break
		}
	}

	Logger.Debugf("限价价格：%v 限价数量：%v", tradePrice, tradeQuantity)
	return tradePrice, tradeQuantity
}

func GetRatio(value1 float64, value2 float64) float64 {

	var big, small float64

	if value1 >= value2 {
		big = value1
		small = value2
	} else {
		big = value2
		small = value1
	}

	return (big - small) * 100 / small ///????
}

// Example: ETH/USDT
func ParsePair(pair string) []string {
	return strings.Split(pair, "/")
}

type Exchanges struct {
	exchanges map[string]IExchange
}

// func (e *Exchanges) Init() {
// 	/* add exchange list */
// 	okexfuture := new(OKExAPI)
// 	okexfuture.Init(InitConfig{
// 		Api:    constOKEXApiKey,
// 		Secret: constOEXSecretKey,
// 		Custom: map[string]interface{}{"tradeType": ExchangeTypeFuture},
// 	})

// 	e.exchanges[okexfuture.GetExchangeName()] = okexfuture

// 	okexspot := new(OKExAPI)
// 	okexspot.Init(InitConfig{
// 		Api:    constOKEXApiKey,
// 		Secret: constOEXSecretKey,
// 		Custom: map[string]interface{}{"tradeType": ExchangeTypeSpot},
// 	})

// 	e.exchanges[okexfuture.GetExchangeName()] = okexspot
// }

// func (e *Exchanges) Start() {

// }
