package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "04"}
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "01")
	//piscine.BTreeInsertData(root, "5")
	fmt.Println(piscine.BTreeLevelCount(root) + 1)
}
