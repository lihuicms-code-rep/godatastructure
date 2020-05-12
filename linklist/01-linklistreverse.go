package linklist

import "fmt"

//题目1:实现链表的逆序

//解法1:就地进行逆序
//技巧:使用3指针(pre,cur,next),但是注意最后一个结点next的判断
//时间复杂度O(n),空间复杂度O(1)
func ReverseLinkList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//三指针分别指向当前要改变指向节点的前一个,当前,后一个
	var pre, cur, next *ListNode = nil, head, head.Next
	for {
		if cur == nil { //遍历到尾部,pre就是当前的head
			break
		}
		cur.Next = pre
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}

	return pre
}

//解法2:插入法
//过程比较直观,可以画个图演示过程
//时间复杂度O(n),空间复杂度O(n)
func ReverseLinkList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode = nil
	for {
		if head == nil {
			break
		}
		node := NewListNode(head.Val)
		node.Next = newHead
		newHead = node

		head = head.Next
	}

	return newHead
}

//由此引申的题目:从尾到头输出链表
//解法1:原地逆序,然后顺序输出,这个会导致原始的链表结构改变
//解法2:插入法也会导致结构改变
//解法3:可以借助一个栈,需要额外的空间
//解法4:递归输出
func ReversePrint(head *ListNode) {
	if head == nil {
		return
	}

	ReversePrint(head.Next)
	fmt.Printf("%d->", head.Val)
}