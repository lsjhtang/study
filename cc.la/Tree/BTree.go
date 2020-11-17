package Tree

import "fmt"

type BTree struct {
	Value interface{}
	Left *BTree
	Right *BTree
}

func NewBTree(value interface{}) *BTree  {
	return  &BTree{Value:value}
}

func (this *BTree) String()  {
	fmt.Printf("这是二叉树%d\n",this.Value)
	var leftValue interface{}
	if this.Left !=nil {
		leftValue = this.Left.Value
	}
	var rightValue interface{}
	if this.Right !=nil {
		rightValue = this.Right.Value
	}
	fmt.Printf("左节点%v\n右节点%v\n",leftValue, rightValue)
}

func (this *BTree) ConnectLeft(bt interface{}) *BTree {
	if btl,ok := bt.(*BTree);ok {
		this.Left = btl
	}else {
		this.Left = NewBTree(bt)
	}
	return this
}

func (this *BTree) ConnectRight(bt interface{}) *BTree  {
	if btr,ok := bt.(*BTree);ok {
		this.Right = btr
	}else {
		this.Right = NewBTree(bt)
	}
	return this
}

func (this BTrees) String()  {
	for _,bt := range this{
		bt.String()
	}
	fmt.Printf("当前一共有%d个节点\n",len(this))
}

type BTrees []*BTree
func NewBTrees(value ...interface{}) BTrees {
	brees := make(BTrees,len(value))
	for index,v :=range value{
		brees[index] = NewBTree(v)
	}
	return brees
}