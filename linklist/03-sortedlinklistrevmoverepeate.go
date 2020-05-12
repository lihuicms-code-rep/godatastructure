package linklist

//题目3:从有序链表中删除重复的元素
//题目为题目2的延伸,之前的方法还是可用的

//由于是有序链表,只需要判断前后两个即可
func RemoveRepeateListNodeInSorted(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	for cur := head; cur != nil; cur = cur.Next{
		if cur.Next == nil {
			break
		}

		if cur.Val == cur.Next.Val {
			p := cur.Next
			cur.Next = cur.Next.Next
			p.Next = nil
		}
	}
}

