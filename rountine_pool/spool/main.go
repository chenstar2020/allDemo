package main

import(
	"allDemo/rountine_pool/spool/internal"
	"fmt"
	"math/rand"
	"time"
)

//任务列表
var nameSlices = []string{"001", "002", "003", "004", "005", "006", "007", "008", "009"}


type workersp struct{
	ID string
}

func (m *workersp) Task() error{
	timen := rand.Intn(3)
	time.Sleep(time.Second * time.Duration(timen))
	fmt.Println("job:" + m.ID + "over")
	return nil
}

func (m *workersp) GetTaskID() interface{}{
	return m.ID
}

func main(){
	spool := internal.NewWPool(3, cap(nameSlices), 2, true)


	for _, id := range nameSlices{
		np := workersp{ID: id}
		spool.Commit(&np)
	}

	spool.Release()
	time.Sleep(time.Second * 1)
}
