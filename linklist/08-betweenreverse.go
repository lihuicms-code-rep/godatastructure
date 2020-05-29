package linklist

//题目8:链表相邻元素翻转(非递归写法不是很理解)
/*
 把链表相邻元素翻转，例如给定链表为1->2->3->4->5->6->7，则翻转后的链表变为2->1->4->3->6->5->7
 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */

//解法1:递归,豁然开朗的写法
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next                //第二个结点
	head.Next = SwapPairs(tmp.Next) //第二个结点的后面总之就是已经翻转好的部分
	tmp.Next = head                 //第二个指向第一个

	return tmp
}

//解法2:非递归,需要描述整个过程
func SwapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//这个结点的使用很重要
	h := &ListNode{
		Val:0,
		Next:head,
	}

	pre, cur := h, h.Next

	for {
		if cur == nil {
			break
		}

		if cur.Next == nil {
			break
		}

		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur





	}

	return nil
}

