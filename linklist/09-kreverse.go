package linklist

//题目9:k个结点为1组进行翻转
/*
K链表翻转是指把每k个相邻的结点看成一组进行翻转，如果剩余结点不足k个，则保持不变。
假设给定链表1->2->3->4->5->6->7和一个数k，如果k的值为2，那么翻转后的链表为2->1->4->3->6->5->7。
如果k的值为3，那么翻转后的链表为：3->2->1->6->5->4->7。
 */

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if k <= 0 {
		return head
	}

	return nil
}
