package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := uniquePaths(23, 12)
	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	downAmt := big.NewInt(int64(m - 1))
	rightAmt := big.NewInt(int64(n - 1))
	totalStep := new(big.Int).Add(downAmt, rightAmt)

	ttFactorial := calcFactorial(totalStep)
	dnFactorial := calcFactorial(downAmt)
	rtFactorial := calcFactorial(rightAmt)

	possiblePaths := new(big.Int).Div(ttFactorial, new(big.Int).Mul(dnFactorial, rtFactorial))
	// Check if the result fits in an int
	if possiblePaths.IsInt64() {
		intValue := possiblePaths.Int64()
		if intValue <= int64(^uint(0)>>1) {
			return int(intValue)
		}
	}

	// If it doesn't fit, you might want to handle this case
	fmt.Println("Warning: Result too large for int, returning max int value")
	return int(^uint(0) >> 1) // Returns MaxInt
}

func calcFactorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}

	result := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		result.Mul(result, i)
	}
	return result
}

// note
// ตอนแรกเลย start with CONCRETE lol
// นั่งวาดลูกศรทีละ stepๆ เนี่ยแหละ down, right
// แล้วก็วาด paths แล้วนับ555555
// เลยจะทำ if else เลขน้อยๆ แล้วเลขเยอะจะ return m*n + max(m,n) เพราะ นึกว่า 3*7 + 7 จะได้ all possible paths ;-; --> (3,7) คือ m,n ตัวอย่าง

// ต่อมาไปแตก test case m=5, n=5 ได้ paths=70
// ก็มานะวาดรูป step ลูกศร 5x5 เนี่ยแหละ แล้วมันปิ๊งได้ว่า ถ้าจะนับจากช่อง start ถึง finish
// rightStep ต้อง = 4
// downStep ต้อง = 4

// ทีนี้ได้ sub-problem มาตัวนึงคือ แล้วมันวิธีต่อ D(Down), R(Right) กี่วิธีวะ ถึงจะไปถึง finish point
// เอาไปถาม claude -> claude ให้สูตร (!(m+n)/(!m*!n))
// ถามมันต่อ มันบอก ไปเรียน Math:Combinatoric ซะนะ

// แต่คืนนี้อยาก submit ละเลยหยิบเอาสูตรมาใช้เลย แต่คิดว่าต้องไปอ่านล่ะ น่าจะได้ใช้เยอะแน่ๆโจทย์ต่อๆไป

// ละก็แตกตรง case m=23, n=12

// เลยไปถาม Claude ว่าคำนวนไงวะ factorial แม่ง overflow
// Claude บอกใช้ bigint

// จ๊บบบ

// เจอละ Combination ภาษาไทยใช้คำว่าการจัดหมวดหมู่
// https://www.youtube.com/watch?v=YDJpoFvz1NU
