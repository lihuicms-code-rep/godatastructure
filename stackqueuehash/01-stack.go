package stackqueuehash

//题目1:实现一个栈的数据结构，使其具有以下方法：压栈、弹栈、取栈顶元素、判断栈是否为空以及获取栈中元素个数
//实现方式:数组实现,链表实现

//方法1:数组(切片)实现
type ArrayStack struct {
	data []int //数据部分
	top  int   //栈顶指针
	num  int   //元素个数
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]int, 0),
		top:  -1,
		num:  0,
	}
}

//入栈
func (as *ArrayStack) Push(item int) {
	as.data = append(as.data, item)
	as.top++
	as.num++
}

//出栈
func (as *ArrayStack) Pop() int {
	if as.IsEmpty() {
		return -1
	}

	item := as.data[as.top]
	as.data = as.data[:as.top]
	as.top--
	as.num--
	return item
}

//栈为空？
func (as *ArrayStack) IsEmpty() bool {
	return as.top == -1
}

//栈元素个数
func (as *ArrayStack) GetNum() int {
	return as.num
}

//获取栈顶元素
func (as *ArrayStack) Top() int {
	if as.IsEmpty() {
		return -1
	}

	return as.data[as.top]
}

//方法2:链表实现
type StackNode struct {
	data int        //数据域
	next *StackNode //下一结点
}

type LinkStack struct {
	head *StackNode  //表头
	num  int         //元素个数
}

func NewLinkStack() *LinkStack {
	ls := &LinkStack{
		head : &StackNode{-1,nil},
		num : 0,
	}

	return ls
}

//入栈,头插法
func (ls *LinkStack) Push(item int) {
	newNode := &StackNode{
		data : item,
		next:nil,
	}

	newNode.next = ls.head.next
	ls.head.next = newNode

	ls.num++
}

//出栈
func (ls *LinkStack) Pop() int {
	if ls.IsEmpty() {
		return -1
	}

	p := ls.head.next     //出栈结点
	ls.head.next = p.next
	p.next = nil
	ls.num--
	return p.data
}


//判断栈空
func (ls *LinkStack) IsEmpty() bool {
	if ls.head.next == nil {
		return true
	}

	return false
}

//栈顶元素
func (ls *LinkStack) Top() int {
	if ls.IsEmpty() {
		return -1
	}
	return ls.head.next.data
}

//元素个数
func (ls *LinkStack) GetNum() int {
	return ls.num
}

