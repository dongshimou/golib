package high_energy_pre_warning

import (
	"github.com/dongshimou/golib/data_structure/segment_tree"
	"github.com/dongshimou/golib/logger"
	"github.com/dongshimou/golib/util"
	"sort"
	"time"
)

type TickData interface {
	GetTimestamp()int64
	GetLastPrice()float64
	GetTradingDay()string
	GetOpenPrice()float64
}

type HighEnergyMsg struct {
	InstrumentID string  `json:"instrument_id"`
	LimitPrice   float64 `json:"limit_price"`
	LimitBase    float64 `json:"limit_base"`
	LimitRate    float64 `json:"limit_rate"`
	LimitMinute  int     `json:"limit_minute"`
	LastPrice    float64 `json:"last_price"`
	UpdateTime   string  `json:"update_time"`
	UpDown       int     `json:"up_down"`
	Timestamp    int64   `json:"timestamp"`
}

type HighEnergyPreWarning struct {
	SegmentTree *segment_tree.SegmentTree
	InFlowStream chan TickData
	OutFlowStream chan HighEnergyMsg
	Minutes []int//需要查询的最近多少分钟
	TickTimes     map[int]int64               //几秒更新一次
	LimitUpRate   float64                     //最高涨幅比例
	LimitDownRate float64                     //最低跌幅比例

	initUpDownBase func(string) (float64, float64) //新的一天如何初始化 baseUp baseDown
	baseUp         float64                   //涨幅基准
	baseDown       float64                   //跌幅基准
	times          []int64                   //时间戳
	tradingDay     string                    //交易日
	instrumentID   string                    //合约id
	pubDown        map[int]int64             //已经发布过跌
	pubUp          map[int]int64             //已经发布过涨
	isForce        bool                      //不足时长是否发布
}

type HighEnergyPreWarningConfig struct {
	InstrumentID string                            //instrumentID 合约id (cu0000)
	Minutes      []struct {
		Minute   int
		ColdTime int
	}                                              //minutes 需要预警的分钟数 (30,30)
	LimitUpRate    float64                         //limitMax 涨幅预警比率 (0.005)
	LimitDownRate  float64                         //limitMin 跌幅预警比率 (0.005)
	InitUpDownBase func(string) (float64, float64) //新的一天如何初始化 baseUp baseDown
	IsForce        bool                            //isForce 不足minutes时,是否发布 (true)
	InStreamSize   int                             //初始化流入大小
	OutStreamSize  int                             //初始化流出大小
}
const (
	Limit_Down = -1
	Limit_Up   = 1
	DefaultStreamSize = 1024
)
var (
	maxFloat = func(l, r interface{}) interface{} {
		if l.(float64) > r.(float64) {
			return l
		} else {
			return r
		}
	}
	minFloat = func(l, r interface{}) interface{} {
		if l.(float64) < r.(float64) {
			return l
		} else {
			return r
		}
	}
)
func New(config *HighEnergyPreWarningConfig) *HighEnergyPreWarning {
	if config.InStreamSize <= 0 {
		config.InStreamSize = DefaultStreamSize
	}
	if config.OutStreamSize <= 0 {
		config.OutStreamSize = DefaultStreamSize
	}

	hepw := &HighEnergyPreWarning{
		instrumentID:  config.InstrumentID,
		LimitUpRate:   config.LimitUpRate,
		LimitDownRate: config.LimitDownRate,
		TickTimes:     map[int]int64{},
		pubDown:       map[int]int64{},
		pubUp:         map[int]int64{},
		InFlowStream:  make(chan TickData, config.InStreamSize),
		OutFlowStream: make(chan HighEnergyMsg, config.OutStreamSize),
		initUpDownBase:config.InitUpDownBase,
		isForce:       config.IsForce,
	}
	for i,_:=range config.Minutes{
		hepw.Minutes=append(hepw.Minutes,config.Minutes[i].Minute)
		hepw.TickTimes[config.Minutes[i].Minute]=int64(config.Minutes[i].ColdTime)
	}
	return hepw
}

func (this *HighEnergyPreWarning) TradingDay() string {
	return this.tradingDay
}
func (this *HighEnergyPreWarning) InstrumentID() string {
	return this.instrumentID
}
func (this *HighEnergyPreWarning) NewDay(tradingDay string, openPrice float64) {

	this.init(tradingDay, openPrice, openPrice)

	if this.initUpDownBase != nil {
		baseUp, baseDown := this.initUpDownBase(this.instrumentID)
		if baseUp!=0&&baseDown!=0{
			this.baseUp,this.baseDown=baseUp,baseDown
		}
	}
}
func (this *HighEnergyPreWarning) init(tradingDay string, baseMax, baseMin float64, args ...interface{}) {
	this.tradingDay = tradingDay
	this.SegmentTree = segment_tree.New([]segment_tree.Function{maxFloat, minFloat}, args)
	this.times = make([]int64, 0, 2048)
	this.baseUp = baseMax
	this.baseDown = baseMin
	logger.Debug(this.instrumentID, "start a new High Energy Pre Warning in", this.tradingDay)
}
func (this *HighEnergyPreWarning) InitTicks(ticks []TickData) {

	for i, _ := range ticks {
		this.process(ticks[i])
	}
}
func (this *HighEnergyPreWarning) setPubUp(minute int,t int64) {
	this.pubUp[minute]=t
}
func (this *HighEnergyPreWarning) setPubDown(minute int,t int64){
	this.pubDown[minute]=t
}
func (this *HighEnergyPreWarning) isPubUp(minute int,t int64)bool {
	if v, exist := this.pubUp[minute]; !exist {
		return false
	} else {
		return t-v < this.TickTimes[minute]
	}
}
func (this *HighEnergyPreWarning) isPubDown(minute int,t int64)bool{
	if v, exist := this.pubDown[minute]; !exist {
		return false
	} else {
		return t-v < this.TickTimes[minute]
	}
}
func (this *HighEnergyPreWarning) lastTime() int64 {
	return this.times[len(this.times)-1]
}
func (this *HighEnergyPreWarning) foundTime(t int64) int {
	tindex := sort.Search(len(this.times), func(i int) bool {
		return this.times[i] == t
	})
	return tindex
}
func (this *HighEnergyPreWarning) empty() bool {
	return len(this.times) == 0
}
func (this *HighEnergyPreWarning) updateLast(v interface{}) {
	this.SegmentTree.Update(this.SegmentTree.Length-1, v)
}

//pos >= len 会执行 append
func (this *HighEnergyPreWarning) update(pos int, v interface{}) {
	this.SegmentTree.Update(pos, v)
}
func (this *HighEnergyPreWarning) append(t int64, v interface{}) {
	this.times = append(this.times, t)
	this.SegmentTree.Append(v)
}
//处理tick数据
func (this *HighEnergyPreWarning) process(tick TickData) {
	if this.tradingDay != tick.GetTradingDay() { //交易日变化,新的一天!
		this.NewDay(tick.GetTradingDay(), tick.GetOpenPrice())
	}
	if !this.empty() {
		if this.lastTime() == tick.GetTimestamp() {
			this.updateLast(tick.GetLastPrice())
		} else if this.lastTime() < tick.GetTimestamp() {
			this.append(tick.GetTimestamp(), tick.GetLastPrice())
		} else {
			tIndex := this.foundTime(tick.GetTimestamp())
			this.update(tIndex, tick.GetLastPrice())
		}
	} else {
		this.append(tick.GetTimestamp(), tick.GetLastPrice())
	}
}
func (this *HighEnergyPreWarning) pubMsg(tick TickData) {
	for _, minute := range this.Minutes {

		//已经发布过了
		if this.isPubUp(minute,tick.GetTimestamp())&&this.isPubDown(minute,tick.GetTimestamp()){
			continue
		}
		//查询距离当前curv.T 最后minute分钟的最值
		maxPrice, minPrice := this.query(tick.GetTimestamp(), minute)
		if maxPrice == 0 && minPrice == 0 {
			continue
		}
		//跌
		if !this.isPubDown(minute,tick.GetTimestamp())&&maxPrice!=0 {
			if (maxPrice-tick.GetLastPrice())/this.baseDown > this.LimitDownRate {
				pubData := HighEnergyMsg{
					InstrumentID: this.instrumentID,
					LimitRate:    this.LimitDownRate,
					LimitPrice:   maxPrice,
					LimitBase:    this.baseDown,
					LimitMinute:  minute,
					UpdateTime:   time.Unix(tick.GetTimestamp(), 0).Format(util.DatetimeFormat),
					LastPrice:    tick.GetLastPrice(),
					Timestamp:    tick.GetTimestamp(),
					UpDown:       Limit_Down,
				}
				this.OutFlowStream <- pubData
				this.setPubDown(minute,tick.GetTimestamp())
			}
		}
		//涨
		if !this.isPubUp(minute,tick.GetTimestamp())&&minPrice!=0 {
			if (minPrice-tick.GetLastPrice())/this.baseUp > this.LimitUpRate {
				pubData := HighEnergyMsg{
					InstrumentID: this.instrumentID,
					LimitRate:    this.LimitUpRate,
					LimitPrice:   minPrice,
					LimitBase:    this.baseUp,
					LimitMinute:  minute,
					UpdateTime:   time.Unix(tick.GetTimestamp(), 0).Format(util.DatetimeFormat),
					LastPrice:    tick.GetLastPrice(),
					Timestamp:    tick.GetTimestamp(),
					UpDown:       Limit_Up,
				}
				this.OutFlowStream <- pubData
				this.setPubUp(minute,tick.GetTimestamp())
			}
		}
		//logger.Debug(this.instrumentID, tick.ActionDay, tick.UpdateTime, maxPrice, minPrice, this.baseDown, this.baseUp)
	}
}
//查询最近多少分钟的最值
func (this *HighEnergyPreWarning) query(t int64, minute int) (max, min float64) {
	preT := time.Unix(t, 0).Add(time.Minute * -time.Duration(minute)).Unix()
	//二分搜索到index
	tindex := sort.Search(len(this.times), func(i int) bool {
		return this.times[i] >= preT
	})
	if tindex > 0 || this.isForce { //找到时间区间或者强制发布
		res := this.SegmentTree.Query(tindex, this.SegmentTree.Length-1)
		if len(res) != 2 {
			return 0, 0
		}
		maxPrice, _ := res[0].(float64)
		minPrice, _ := res[1].(float64)
		return maxPrice, minPrice
	}
	return 0, 0
}
func (this *HighEnergyPreWarning) Run() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {

		case tick := <-this.InFlowStream:
			{
				this.process(tick)
				this.pubMsg(tick)
			}
		case <-ticker.C:
			{
				logger.Debug(this.instrumentID,this.tradingDay,"run HighEnergyPreWarning!")
			}
		}
	}
}
