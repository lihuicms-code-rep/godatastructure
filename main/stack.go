package main

import (
	"fmt"
	"github.com/lihuicms-code-rep/godatastructure/stackqueuehash"
)

//栈测试部分
func main() {
	//TestArrayStack()
	TestLinkStack()
}

//测试数组(切片)构造的栈
func TestArrayStack() {
	as := stackqueuehash.NewArrayStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	as.Push(4)
	fmt.Printf("当前栈为:%+v\n", as)
	top := as.Top()
	fmt.Printf("当前栈顶元素:%d\n", top)
	for {
		fmt.Printf("当前栈中元素个数为:%d\n", as.GetNum())
		p := as.Pop()
		if p == -1 {
			break
		}
		fmt.Printf("出栈元素:%d\n", p)
	}
}

//测试链表构造的栈
func TestLinkStack() {
	ls := stackqueuehash.NewLinkStack()
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	ls.Push(4)
	fmt.Printf("当前栈为:%+v\n", ls)
	top := ls.Top()
	fmt.Printf("当前栈顶元素:%d\n", top)
	for {
		fmt.Printf("当前栈中元素个数为:%d\n", ls.GetNum())
		p := ls.Pop()
		if p == -1 {
			break
		}
		fmt.Printf("出栈元素:%d\n", p)
	}

}
