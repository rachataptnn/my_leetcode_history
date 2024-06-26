// https://leetcode.com/problems/2-keys-keyboard/description/

package main

import "fmt"

func main() {
	input := 989 // 66

	fmt.Println(minSteps(input))
}

func minSteps(n int) int {
	if n <= 5 {
		if n == 1 {
			return 0
		}
		return n
	}

	if isPrime(n) {
		return n
	}

	return divideByBigFirst(n)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func divideByBigFirst(input int) int {
	theNums := []int{}
	if input%2 == 0 {
		theNums = []int{77, 35, 19, 13, 11, 7, 5, 3, 2}
	} else {
		theNums = []int{23, 22, 19, 13, 11, 7, 5, 3, 2}
	}

	step := 0
	for {
		for _, theNum := range theNums {
			if input%theNum == 0 {
				step += theNum
				input = input / theNum
			}
		}

		if input == 1 {
			break
		}

		if isPrime(input) {
			step = input + step
			break
		}
	}

	return step
}

// my note <3
// เริ่มจากลองสร้างอะไรที่มัน concrete อย่างการเขียน AAAAA ลงสมุดถถถถถ
// n=3, 4, 10, 11 จะมี AAA... กี่ตัว ละมีการ copy, paste ต่ำสุดกี่ step
// end case type 1 -> แล้วก็เห็น pattern ว่าถ้าเป็น prime number แม่งคืนตัวเองได้เลยนี่หว่า

// ต่อมาก็คิดว่ามันสามารถ breakdown เลขได้ โดยถ้าอยากรู้ minimum เราต้องใช้เลขที่ใหญ่ที่สุดที่เป็นไปได้ break ก่อน
// อันนี้เป็นที่มาของฟังชัน divideByBigFirst()
// end case type 2 -> ก็นั่นแหละ prime number ก็ออก, 1 ก็ออกเช่นกัน
// ทีนี้มันจะมีการคำนวน type 2 ต่างจาก type 1 นิดหน่อยตรงที่ว่า ถ้าเป็นเลขใน set type2 เนี่ย
// มันจะมีการ break มาก่อน กี่ step ก็ว่าไป, จึงต้องเอา step + prime number ที่ break ไม่ได้ละ (line 68 step = input + step)

// theNums นี่แม่งมโนหน่อยๆ
// ตอนแรกยังไม่ได้ทำเป็น loop ยังงั้น ตอนแรกไล่จาก
// if input%5 -> step += theNum & input = input / theNum
// if input%3 -> step += ...
// if input%2 -> ...
// แล้วพอกด submit แม่งได้เคสมาเพิ่มเรื่อยๆ จากตอนแรกมีแค่ 5, 3, 2(ได้มาจาก break เลข 600 ในสมุด)
// ก็ได้พวก 19, 11 อะไรมาด้วย สักพักเลขเริ่มเยอะเลยทำเป็น array ไว้วน

// ทำท่าเดิมไปเรื่อยๆก็ค่อยๆงอกเลข แล้วก็จะมี case 385, 741, 989 นี่แหละที่คิดว่าควร switch theNums โดยมองจาก input เป็นเลขคู่หรือเลขคี่
