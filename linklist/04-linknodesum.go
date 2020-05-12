package linklist

import (
	"fmt"
	"math"
)

//题目4:计算两个单链表代表的数之和

/*
 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */

//解法1:先将数据之和求出来,再构造链表
//这种解法通不过leetcode
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var num1, num2 int64
	for head, i := l1, 0; head != nil; head, i = head.Next, i+1 {
		num1 += int64(head.Val) * int64(math.Pow(10, float64(i)))
	}

	for head, i := l2, 0; head != nil; head, i = head.Next, i+1 {
		num2 += int64(head.Val) * int64(math.Pow(10, float64(i)))
	}

	fmt.Printf("num1:%d, num2:%d\n", num1, num2)
	if num1 == 0 && num2 == 0 {
		return l1
	}

	sum := num1 + num2 //有可能会溢出
	var head, tmp *ListNode
	for {
		if sum == 0 {
			break
		}

		val := sum % 10
		node := &ListNode{
			Val:  int(val),
			Next: nil,
		}

		if head == nil {
			head = node
			tmp = head
		} else {
			tmp.Next = node
			tmp = tmp.Next
		}

		sum = sum / 10
	}

	return head
}

//解法2
//链表相加法,记录是否有进位
func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	flag := false           //进位标记
	var head, tmp *ListNode //新链表头,当前尾
	var p1, p2 *ListNode    //l1,l2扫描指针

	for p1, p2 = l1, l2; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		wSum := 0 //当前位两数之和
		if flag {
			wSum = p1.Val + p2.Val + 1
		} else {
			wSum = p1.Val + p2.Val
		}

		if wSum >= 10 {
			flag = true
		} else {
			flag = false
		}

		wSum %= 10
		node := &ListNode{
			Val:  wSum,
			Next: nil,
		}

		if head == nil {
			head = node
			tmp = head
		} else {
			tmp.Next = node
			tmp = tmp.Next
		}
	}

	//p1较短
	if p1 == nil {
		for ; p2 != nil; p2 = p2.Next {
			wSum := 0
			if flag {
				wSum = p2.Val + 1
			} else {
				wSum = p2.Val
			}

			if wSum >= 10 {
				flag = true
			} else {
				flag = false
			}

			wSum %= 10
			node := &ListNode{
				Val:  wSum,
				Next: nil,
			}

			tmp.Next = node
			tmp = tmp.Next
		}

		if flag {               //特别注意最后一个位置的判断
			tmp.Next = &ListNode{
				1,
				nil,
			}
		}
	}

	//p2较短
	if p2 == nil {
		for ; p1 != nil; p1 = p1.Next {
			wSum := 0
			if flag {
				wSum = p1.Val + 1
			} else {
				wSum = p1.Val
			}

			if wSum >= 10 {
				flag = true
			} else {
				flag = false
			}

			wSum %= 10
			node := &ListNode{
				Val:  wSum,
				Next: nil,
			}

			tmp.Next = node
			tmp = tmp.Next
		}

		if flag {
			tmp.Next = &ListNode{
				1,
				nil,
			}
		}
	}

	return head
}
