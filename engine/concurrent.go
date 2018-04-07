package engine

import (
	"crawler/fetcher"

	"crawler/scheduler"
	"crawler/types"
	"log"
)

const count  = 10
func Run(seeds ...types.Request) {
	workerChan := make(chan types.ParseResult, count)
	schedulerChan := make(chan types.Request, count)
	schedulerWorkerChan := make(chan types.Request, count)
	log.Printf("workerChan %d\n", len(schedulerChan))
	log.Printf("schedulerChan %d\n", len(schedulerChan))
	log.Printf("schedulerWorkerChan is "+"%d\n", len(schedulerChan))
	item := 0
	for _, r := range seeds {
		schedulerChan <- r
		log.Println("schedulerChan %d", len(schedulerChan))
		item++
		log.Println("hello")
	}

	log.Println("schedulerChan %d", len(schedulerChan))
	log.Println("seeds 加载完成")
	go scheduler.Run(schedulerChan, schedulerWorkerChan)
	for i := 0; i < count; i++ {
		log.Println(i)
		go createWorker(schedulerWorkerChan, workerChan)
	}
	for {
		result := <-workerChan
		for _, item := range result.Items {
			log.Printf("%v ", item)
		}
		for _, r := range result.Requests {
			schedulerChan <- r
		}
	}

}
func createWorker(in <-chan types.Request, out chan<- types.ParseResult) {
	func() {
		for {
			result, err := work(<-in)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func work(r types.Request) (types.ParseResult, error) {
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, e)
		return types.ParseResult{}, e
	}
	return r.ParserFunc(body), e
}
