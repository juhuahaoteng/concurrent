package scheduler

import (
	"crawler/types"
	"log"
)

func Run(in <-chan types.Request, out chan<- types.Request) {
	var requestQ []types.Request
	for {
		log.Printf("length of requestQ is %d", len(requestQ))
		requestQ = append(requestQ, <-in)
		log.Printf("length of requestQ is %d", len(requestQ))
		if len(requestQ) > 0 && len(out) < 10 {
			out <- requestQ[0]
			requestQ = requestQ[1:]
		}
		log.Printf("length of requestQ is %d", len(requestQ))

	}

}
