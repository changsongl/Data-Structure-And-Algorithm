package Binary_Tree

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	bt := New()
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(8)
	bt.Insert(7)
	bt.Insert(0)

	testCase := map[int]bool{3: true, 8: true, 0: true, 10:false, 100: false}

	for num, expect := range testCase{
		_, result := bt.Search(num)
		if result != expect {
			t.Errorf("Binary Tree Search Failed: number %d, expected %t, result %t", num, expect, result)
		}
	}

	bt.InOrderTraversal(bt.root, func(num int){
		fmt.Print(strconv.Itoa(num) + " ")
	})
	fmt.Println()

	bt.PreOrderTraversal(bt.root, func(num int){
		fmt.Print(strconv.Itoa(num) + " ")
	})
	fmt.Println()

	bt.PostOrderTraversal(bt.root, func(num int){
		fmt.Print(strconv.Itoa(num) + " ")
	})
	fmt.Println()
}