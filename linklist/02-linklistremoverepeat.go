package linklist

//题目2:从无序链表中移除重复项

//方法1:双重循环依次进行检测,时间复杂度为O(n*n)
func RemoveRepeatListNode(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var outer *ListNode           //待比较结点指针
	var inner, innerPre *ListNode //扫描结点指针,以及其前驱

	for outer = head; outer != nil; outer = outer.Next {
		for inner, innerPre = outer.Next, outer; inner != nil; {
			if outer.Val == inner.Val {
				innerPre.Next = inner.Next
				p := inner //被断开结点
				inner = inner.Next
				p.Next = nil
			} else {
				innerPre = inner
				inner = inner.Next
			}
		}
	}
}

//方法2:借助辅助空间,一般是hash
func RemoveRepeatListNode2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	m := make(map[int]bool)
	m[head.Val] = true

	for p, pre := head.Next, head; p != nil; {
		if _, ok := m[p.Val]; !ok { //如果不存在
			m[p.Val] = true
			pre = p
			p = p.Next
		} else {                    //如果存在则删除
		   tmp := p
		   pre.Next = p.Next
		   p = p.Next
		   tmp.Next = nil
		}
	}
}
