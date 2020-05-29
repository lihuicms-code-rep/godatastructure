package linklist
//题目10:合并两个有序链表

/*
已知两个链表head1和head2 各自有序（例如升序排列），请把它们合并成一个链表，要求合并后的链表依然有序
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode
	var p1, p2 *ListNode
	var cur *ListNode       //待合并链表的尾部
	for p1, p2 = l1, l2; p1 != nil && p2 != nil; {
		if p1.Val <= p2.Val {
			if head == nil {
				head = p1
				cur = head
			} else {
				cur.Next = p1
				cur = p1
			}
			p1 = p1.Next
		} else {
			if head == nil {
				head = p2
				cur = head
			} else {
				cur.Next = p2
				cur = p2
			}
			p2 = p2.Next
		}
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return head

}