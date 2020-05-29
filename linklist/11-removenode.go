package linklist

//题目11:如何在只给定单链表中某个结点指针的情况下删除该结点
/*
一般而言，要删除单链表中的一个结点p，首先需要找到结点p的前驱结点pre，然后通过pre.next=p.next来实现对结点p的删除。
对于本题而言，由于无法获取到结点p的前驱结点，因此，不能采用这种传统的方法。那么如何解决这个问题呢？可以分如下两种情况来分析：
（1）如果这个结点是链表的最后一个结点，那么无法删除这个结点。
（2）如果这个结点不是链表的最后一个结点，可以通过把其后继结点的数据复制到当前结点
 */

func RemoveNode(node *ListNode) bool {
	if node == nil || node.Next == nil {      //最后一个结点是无法删除
		return false
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return true
}
