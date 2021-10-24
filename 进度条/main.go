package main

import (
	"allDemo/进度条/progress_bar"
	"time"
)

func main(){
	 bar := progress_bar.NewBar(100)
	 for i := 0; i <= 100; i++ {
	 	time.Sleep(time.Millisecond * 100)
	 	bar.Display(i)

	 }
}
