package main

import (
	"github.com/dongshimou/golib/high_energy_pre_warning"
	"github.com/dongshimou/golib/logger"
	"time"
)
type Tick struct {
	InstrumentID string
	Timestamp int64
	LastPrice float64
	TradingDay string
	OpenPrice float64
}
func(this*Tick)GetTimestamp()int64{
	return this.Timestamp
}
func(this*Tick)GetLastPrice()float64{
	return this.LastPrice
}
func(this*Tick)GetTradingDay()string{
	return this.TradingDay
}
func(this*Tick)GetOpenPrice()float64{
	return this.OpenPrice
}

func Test(){
	InstrumentIDs := map[string]struct{
		UpRate   float64
		DownRate float64
	}{
		"cu0000":    {0.007,0.007},
		"al0000":    {0.007,0.007},
		"pb0000":    {0.007,0.007},
		"zn0000":    {0.007,0.007},
		"ni0000":    {0.007,0.007},
		"sn0000":    {0.007,0.007},
	}
	curChanMap := map[string]chan high_energy_pre_warning.TickData{}
	pubMsg := func(stream chan high_energy_pre_warning.HighEnergyMsg) { //发布高能预警
		for {
			func() {
				for {
					data := <-stream
					logger.Debug(data.InstrumentID, data.UpdateTime, "PubHighEnergy ==> last:", data.LastPrice, "limit:", data.LimitPrice, "rate:", data.LimitRate)
					//do
				}
			}()
		}
	}

	for id, v := range InstrumentIDs {

		//每个合约创建预警
		hepwSystem := high_energy_pre_warning.New(
			&high_energy_pre_warning.HighEnergyPreWarningConfig{
				InstrumentID:   id,
				Minutes: []struct {
					Minute   int
					ColdTime int
				}{{30,30}}        , //30分钟预警
				LimitUpRate:    v.UpRate,
				LimitDownRate:  v.DownRate,
				InitUpDownBase: nil,
				IsForce:        true,
			})
		curChanMap[hepwSystem.InstrumentID()] = hepwSystem.InFlowStream

		//ticks:=qcache.GetAllTick(conn,id)
		//hepwSystem.InitTicks(ticks)
		go pubMsg(hepwSystem.OutFlowStream)
		go hepwSystem.Run()
	}

	start:=time.Now().Unix()
	tick:=Tick{}
	tick.Timestamp=start
	tick.OpenPrice=40000
	tick.LastPrice=40000
	tick.TradingDay="2019-05-25"
	tick.InstrumentID="cu0000"

	curChanMap["cu0000"]<- &tick

	tick.Timestamp=start+1
	tick.OpenPrice=40000
	tick.LastPrice=50000
	tick.TradingDay="2019-05-25"
	tick.InstrumentID="cu0000"

	tick.GetLastPrice()
	curChanMap["cu0000"]<-&tick
}

func main(){
	Test()
	for{
		time.Sleep(time.Second)
	}
}