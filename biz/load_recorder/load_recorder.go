package load_recorder

var Loadtotal int64
var LoadMetrics = make(map[int64]int64)
var LoadChan = make(chan int64, 10000)

func Init() {
	// start the loop of load recorder
	go func() {
		for {
			select {
			case t := <-LoadChan:
				LoadMetrics[t]++
				Loadtotal++
			}
		}
	}()
}

func Reset() {
	LoadMetrics = make(map[int64]int64)
	Loadtotal = 0
}
