package test

import (
	"fmt"
	"testing"
)

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("车站中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("车站票已售空")
	}

}

// 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	fmt.Println("代理点开始代理购买")
	proxy.station.sell(name)
	fmt.Println("代理点代理完毕")
}

func TestProxy(t *testing.T) {
	station := &Station{stock: 5}

	station.sell("ryoma")

	stationProxy := &StationProxy{station: station}
	stationProxy.sell("ryoma")
}
