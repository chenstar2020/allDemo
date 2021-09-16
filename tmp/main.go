package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type GetDomainLineDataResp struct { //返回参数
	ProvIspBw []*ProvIspBwData `json:"prov_isp_bw"` //省份运营商数据
	IspBw     []*IspBwData     `json:"isp_bw"`      //运营商数据
	TotalBw   *TotalBwData     `json:"total_bw"`    //总计数据
}

type ProvIspBwData struct {
	IspId  int `json:"isp_id"`
	ProvId int `json:"prov_id"`
}

type IspBwData struct {
	IspId int `json:"isp_id"`
}

type TotalBwData struct {
	Total int `json:"total"`
}
func GetDayZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetLatelyFiveMintueTime() (int64, int64) {
	cur := time.Now()
	nowMintue := cur.Minute()
	if nowMintue < 5 && cur.Hour() == 0 {
		//取0点到当前的时间
		return GetDayZeroTime(cur).Unix(), cur.Unix()
	}
	//超过5分钟 则取前面5分中为开始时间
	fiveMintue := (nowMintue - nowMintue%5) - 5

	begin := time.Date(cur.Year(), cur.Month(), cur.Day(), cur.Hour(), fiveMintue, 0, 0, cur.Location()).Unix()
	end := begin + 300
	return begin, end
}
//交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}


func RandInt(min, max int) int{
	if !(min >= 0 && max >= 0 && max >= min){
		return max
	}

	return rand.Intn(max - min) + min
}

func main(){
	b1 := []byte("aabbcc")
	b2 := []byte("aabbcd")
	fmt.Println(bytes.Compare(b1, b2))


	fmt.Println(RandInt(10, 20))

	s1 := make([]int, 10)

	s2 := s1
	s3 := s1


	s2[3] = 3
    s3[4] = 4

    fmt.Println(s2)


	r := gin.Default()
	err := r.Run("127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
}