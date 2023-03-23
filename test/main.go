package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "2")
	piscine.BTreeInsertData(root, "9")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	root.Data = "9"
	fmt.Println(piscine.BTreeIsBinary(root))
}
