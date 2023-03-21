package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	listPointer := l.Head
	var foundAddr *interface{}

	for listPointer != nil {
		if comp(listPointer.Data, ref) {
			foundAddr = &ref
			return foundAddr
		}
		listPointer = listPointer.Next
	}
	return foundAddr
}
