package structures

// import ()

type SRequest struct {
	Score			int
	Respond			chan int8
}

type SRequestSlice []SRequest

func (list SRequestSlice) Len() int {
	return len(list)
}

func (list SRequestSlice) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list SRequestSlice) Less(i, j int) bool { 
    return list[i].Score < list[j].Score
}