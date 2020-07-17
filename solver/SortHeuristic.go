package solver

import (
	// l "container/list"
	"fmt"
	s "n-puzzle/structures"
	"sort"
)

func manageRequests(requests s.SRequestSlice, requestChan chan s.SRequest) {
	for {
		select {
		case request := <- requestChan:
			requests = append(requests, request)
		default:
			break
		}
	}
	sort.Sort(requests)
	for i := 0; i < 10; i++ {
		requests[i].Respond <- 1
	}
	requests = requests[10:]
	if len(requests) > 1000 {
		length := len(requests)
		for i := 1001; i < length; i++ {
			requests[i].Respond <- -1
		}
		requests = requests[:1000]
	}
}

func stopGoRoutines(requests s.SRequestSlice, requestChan chan s.SRequest) {
	for {
		select {
		case request := <- requestChan:
			requests = append(requests, request)
		default:
			break
		}
	}
	length := len(requests)
	for i := 0; i < length; i++ {
		requests[i].Respond <- -1
	}
}

func finalResult(result s.SResult, requests s.SRequestSlice) {
	fmt.Println(requests)
	fmt.Println(result)
}

func SortHeuristic(ctx *s.SContext) {
	// var result s.SResult
	resultChan := make(chan s.SResult, 10)
	requestChan := make(chan s.SRequest, 400)
	var requests s.SRequestSlice
	firstRoutine(ctx, resultChan, requestChan)
	for {
		select {
		case result := <-resultChan:
			stopGoRoutines(requests, requestChan)
			finalResult(result, requests)
			return
		default:
			manageRequests(requests, requestChan)
		}
	}
	fmt.Println("end")
}