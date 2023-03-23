package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "04"}
	piscine.BTreeInsertData(root, "01")
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "02")
	piscine.BTreeInsertData(root, "05")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "10")
	piscine.BTreeInsertData(root, "03")
	piscine.BTreeApplyByLevel(root, fmt.Println) // 0401070205120310
}
