package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}*/

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	if pos == 0 {
		return l
	}
	return ListAt(l.Next, pos-1)
}
