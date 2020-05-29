package linklist

//题目8:检测一个较大的单链表是否有环

/*
单链表有环指的是单链表中某个结点的next域指向的是链表中在它之前的某一个结点，这样在链表的尾部形成一个环形结构。
 */

//解法1:顺序访问,每次访问到的结点指针放入hash中,如果下一个结点在访问时已经出现在hash中,说明有环
//但是题目前提是:较大单链表,这个会导致使用O(n)的空间
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

    m := make(map[*ListNode]bool)
    for p := head; p != nil; p = p.Next {
    	if _, ok := m[p]; ok {
    		return true
		}

    	m[p] = true
	}

    return false
}

//解法2:快慢指针
func HasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for {
		if slow == fast {
			return true
		}

		if fast == nil {
			break
		}

		if fast.Next == nil {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
