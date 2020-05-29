package linklist

//题目5:对链表进行重新排序
/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
①在原来链表的基础上进行排序，即不能申请新的结点；②只能修改结点的next域，不能修改数据域

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 */

//获取链表的中间结点
func getMidNodeInLinkList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head //快慢指针
	for {
		if fast == nil { //偶数个结点情况
			break
		}

		if fast.Next == nil { //奇数个结点情况
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//反转链表
func revereLinkList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, cur, next *ListNode = nil, head, head.Next
	for {
		if cur == nil {
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

func ReorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	mid := getMidNodeInLinkList(head)
	if mid == nil {
		return
	}

	revMid := revereLinkList(mid)      //将后一部分进行反转
	cur1, cur2 := head, revMid         //cur1:第一部分表头,cur2:第二部分表头
	for {
		if cur1.Next == nil {
			break
		}

		if cur2.Next == nil {
			break
		}
		tmp := cur1.Next
		cur1.Next = cur2
		cur1 = tmp

		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
}