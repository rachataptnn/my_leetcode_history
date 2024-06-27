// https://leetcode.com/problems/add-two-numbers/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := create1stList()
	l2 := create2ndList()

	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := ListNode{}
	add(l1, l2, &resultList, 0)

	return &resultList
}

func add(l1, l2, resultList *ListNode, inCarry int) *ListNode {
	isl1Ended := l1 == nil
	isl2Ended := l2 == nil

	sum := 0
	if !isl1Ended && !isl2Ended {
		sum = l1.Val + l2.Val + inCarry
	} else if !isl1Ended && isl2Ended {
		sum = l1.Val + inCarry
	} else if isl1Ended && !isl2Ended {
		sum = l2.Val + inCarry
	} else {
		sum = inCarry
	}

	lastDigit := sum % 10
	outCarry := sum / 10

	resultList.Val = lastDigit

	isNextL1Ended := false
	isNextL2Ended := false
	if !isl1Ended {
		isNextL1Ended = l1.Next == nil
	} else {
		isNextL1Ended = true
	}
	if !isl2Ended {
		isNextL2Ended = l2.Next == nil
	} else {
		isNextL2Ended = true
	}

	if !isNextL1Ended && !isNextL2Ended {
		resultList.Next = &ListNode{}
		add(l1.Next, l2.Next, resultList.Next, outCarry)
	} else if !isNextL1Ended && isNextL2Ended {
		resultList.Next = &ListNode{}
		add(l1.Next, nil, resultList.Next, outCarry)
	} else if isNextL1Ended && !isNextL2Ended {
		resultList.Next = &ListNode{}
		add(nil, l2.Next, resultList.Next, outCarry)
	} else {
		if outCarry > 0 {
			resultList.Next = &ListNode{}
			resultList.Next.Val = outCarry
		}
		return resultList
	}

	return resultList
}

// helper func

// // Input: l1 = [9,9,9, 9,9,9, 9], l2 = [9,9,9,9]
// // Output: [8,9,9,9,0,0,0,1]
// func create1stList() *ListNode {
// 	list := &ListNode{Val: 9}
// 	list.Next = &ListNode{Val: 9}
// 	list.Next.Next = &ListNode{Val: 9}

// 	list.Next.Next.Next = &ListNode{Val: 9}
// 	list.Next.Next.Next.Next = &ListNode{Val: 9}
// 	list.Next.Next.Next.Next.Next = &ListNode{Val: 9}

// 	list.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

// 	return list
// }

// func create2ndList() *ListNode {
// 	list := &ListNode{Val: 9}
// 	list.Next = &ListNode{Val: 9}
// 	list.Next.Next = &ListNode{Val: 9}

// 	list.Next.Next.Next = &ListNode{Val: 9}

// 	return list
// }

// Input: l1 = [2, 4, 3], l2 = [5, 6, 4]
// Output: [7, 0, 8]
func create1stList() *ListNode {
	list := &ListNode{Val: 2}
	list.Next = &ListNode{Val: 4}
	list.Next.Next = &ListNode{Val: 3}

	return list
}

func create2ndList() *ListNode {
	list := &ListNode{Val: 5}
	list.Next = &ListNode{Val: 6}
	list.Next.Next = &ListNode{Val: 4}

	return list
}

// somenote กันลืม
// ตอนจะทำให้มัน concrete ขึ้นในสมุดจะไปทาง dive ให้เจอ null แล้วเอาเลขมา revers ละแม่ง55555
// นึกไปนึกมา เออผลลัพมันก็ไม่ได้ต้องการ reverse อะไรนี่หว่า มันบวกกันตรงๆทุกหลักเลย
// ก็นั่นแหละ เริ่มเขียนลงสมุดมันก็เห็นเองอะ ต้องมี lastDigit, carry
// ต้องมี fn ifNull แต่คิดแปบนึงก็ได้ละว่า  เออมันบวกกันไปก่อนเลยนี่หว่า แล้วค่อยเช็ค null เป็นเคสๆไป
// เลยได้เป็น add() สั้นๆเนี่ยละ

// งงบูลีนไปพักนึง55555 เช็ค 4 cases เคสละ 2 bool พิมผิดพิมถูก
// แต่คอนเซปได้ก็นั่นแหละจ้า ลู่เข้าหนทางที่ถูกต้องง
