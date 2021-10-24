package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type ApiDataConfig struct {
	configMux sync.RWMutex

	OcFusingThreshold int
	OpenOcFusing      bool

	IpFusingThreshold float64
	OpenIpFusing      bool

	VipFusingThreshold float64
	OpenVipFusing      bool

	PrinceFusingThreshold int
	OpenPrinceFusing      bool
}

type ApiDataManagerOption func(*ApiDataConfig)

func WithIpFusingThreshold(ipFusingThreshold float64, openIpFusing bool)ApiDataManagerOption{
	return func(this *ApiDataConfig) {
		this.IpFusingThreshold = ipFusingThreshold
		this.OpenIpFusing = openIpFusing
	}
}

type OcIspInfoModeId int

type OcIspInfo struct {
	IspId              int
	Ips                []string
	ConstructBandwidth float64
	MinimumBandwidth   float64
	UnitPrice          float64
	ModeId             OcIspInfoModeId
}

func main(){
	fmt.Println(uuid.New().String())
	id := uuid.New()
	fmt.Printf("%T, %v\n", id, id)
	fmt.Printf("%v, %v, %v\n", id.String(), id.Version(), id.Version().String())
}


func test(){
	go func() {
		for{
			fmt.Println("aaa")
			time.Sleep(time.Second)
		}
	}()

	return
}

func test1(){
	test()
}