package solver

import (
	// l "container/list"
	"fmt"
	s "n-puzzle/structures"
	"sort"
	// "time"
)

func manageRequests(requests s.SRequestSlice, requestChan chan s.SRequest) {
	for {
		select {
		case request := <- requestChan:
			requests = append(requests, request)
		default:
			if len(requests) != 0 {
				goto ForEnd
			}
		}
	}
	ForEnd:
	sort.Sort(requests)
	length := len(requests)
	// if length > 10000 {
	// 	for i := 0; i < length; i++ {
	// 		fmt.Print(requests[i].Score, " ")
	// 	}
	// 	time.Sleep(2 * time.Second)
	// 	panic("exit")
	// }
	if length > 1 {
		length = 1
	}
	for i := 0; i < length; i++ {
		requests[i].Respond <- 1
	}
	requests = requests[length:]
	// length = len(requests)
	// if length > 1000 {
	// 	for i := 1001; i < length; i++ {
	// 		requests[i].Respond <- -1
	// 	}
	// 	requests = requests[:1000]
	// }
}

func stopGoRoutines(requests s.SRequestSlice, requestChan chan s.SRequest) {
	for {
		select {
		case request := <- requestChan:
			requests = append(requests, request)
		default:
			goto ForEnd
		}
	}
	ForEnd:
	length := len(requests)
	for i := 0; i < length; i++ {
		requests[i].Respond <- -1
	}
}

func finalResult(result s.SResult, requests s.SRequestSlice) {
	fmt.Println(requests)
	fmt.Println(result.Res)
}

func SortHeuristic(ctx *s.SContext) {
	ctx.ResultChan = make(chan s.SResult, 10)
	ctx.RequestChan = make(chan s.SRequest, 400)
	requests := s.SRequestSlice{}
	firstRoutine(ctx)
	// fmt.Println("routines launched")
	for {
		select {
		case result := <-ctx.ResultChan:
			stopGoRoutines(requests, ctx.RequestChan)
			finalResult(result, requests)
			return
		default:
			// fmt.Println("before manage requests")
			manageRequests(requests, ctx.RequestChan)
		}
	}
	// fmt.Println("end")
}