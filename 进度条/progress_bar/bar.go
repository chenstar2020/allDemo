package progress_bar

import "fmt"

type Bar struct{
	total   int    //总进度
	graph   string //显示符号
}

func NewBar(total int) *Bar{
	return &Bar{
		total: total,
		graph: "█",
	}
}

func NewBarWithGraph(total int, graph string)*Bar{
	return &Bar{
		total: total,
		graph: graph,
	}
}

func (bar *Bar) Display(cur int){
	//计算百分比
	percent := 0
	if cur >= bar.total{
		percent = 100
	}else{
		percent = int(float64(cur) / float64(bar.total) * 100)
	}

	//通过百分比控制字符打印长度
	rate := ""
	for i := 0; i < percent; i += 2{
		rate += bar.graph
	}

	fmt.Printf("\r[%-50s]%d%%   %8d/%d", rate, percent, cur, bar.total)
}
