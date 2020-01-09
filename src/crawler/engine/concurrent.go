package engine

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (e ConCurrentEngine) Run(seeds ...Request){
	e.Scheduler.Run()
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		createWorker(e.Scheduler.WorkerChan(),out, e.Scheduler)
	}
	//参数seeds的request，要分配任务
	for _, request := range seeds {
		if isDuplicate(request.Url){
			continue
		}
		e.Scheduler.Submit(request)
	}

	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			go func(item interface{}) {
				e.ItemChan <- item
			}(item)

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url){
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,ready  ReadyNotifier) {
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			ready.WorkerReady(in)
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}





var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}