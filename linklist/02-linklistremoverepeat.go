package linklist
//题目2:从无序链表中移除重复项

func RemoveRepeatListNode(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}

	//p1:待检查结点,p2:扫描结点,p2Pre:p2的前一个结点
	var p1, p2, p2Pre *LNode = head, nil, nil

	for ; p1 != nil; p1 = p1.Next {
		for p2, p2Pre = p1.Next, p1; p2 != nil; {
			if p1.Data == p2.Data {  //找到重复项
				p2Pre.Next = p2.Next
				p2.Next = nil
			}
			p2 = p2Pre.Next
		}
	}
}
