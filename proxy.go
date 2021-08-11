package main

import (
	"fmt"
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
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}

}

// 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	fmt.Print("代理模式代理购票开始")
	proxy.station.sell(name)
	fmt.Print("代理模式代理购票结束")
}


func main()  {
	station := &Station{stock: 10}
	sp := &StationProxy{station: station}
	sp.sell("王密斯")
}
