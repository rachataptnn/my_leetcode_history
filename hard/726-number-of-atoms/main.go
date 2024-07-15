// https://leetcode.com/problems/number-of-atoms/description/?envType=daily-question&envId=2024-07-14

package main

import (
	"fmt"
	"regexp"
)

//           s  s   e  s                     e  s                  e  s                  e  s  e  e
//           0  1   4  7
// input := "(  (N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"

func main() {
	// input := "H20"
	// input := "K4(ON(SO3)2)2"
	// input := "Mg(OH)2"

	// 13/31
	// input := "((HHe28Be26He)9)34"

	//          s  e   s                     e   s                  e   s                  e   s  e
	input := "((N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"
	// asis ตอนนี้ prepare fn คืน stack, queue เดียว
	// tobe เปลี่ยนเป็น prepare fn ต้องคืน []stack, []queue

	fmt.Println(countOfAtoms(input))
}

func countOfAtoms(formula string) string {
	flatFormula, openBracketIndex := prepareFlatFormula(formula)
	if openBracketIndex != -1 {
		bigFormulas := prepareBigFormulas(formula, openBracketIndex)
		fmt.Println(bigFormulas)
	}
	fmt.Println(flatFormula)
	return ""
}

func peelOuterMultiplier(formula string, openBracketIndex int) (bigFormulas []string) {
	cnt := 0
	var outerMultiplier Queue
	// subText := ""

	i := openBracketIndex
	for {
		if formula[i] == '(' {
			cnt++
		} else if formula[i] == ')' {
			cnt--
		}

		if cnt == 0 {
			if len(formula) > i {
				outerMultiply := formula[i+1:]
				outerMultiplier.Enqueue(outerMultiply)
			}
			formula = formula[1:i]
			fmt.Println("")

		}

		i++
		if i >= len(formula) {
			break
		}
	}

	return bigFormulas
}

func peelOuterMultiply(formula string, openBracketIndex int) (bigFormulas []string) {
	cnt := 0
	var outerMultiplier Queue
	// subText := ""

	i := openBracketIndex
	for {
		if formula[i] == '(' {
			cnt++
		} else if formula[i] == ')' {
			cnt--
		}

		if cnt == 0 {
			if len(formula) > i {
				outerMultiply := formula[i+1:]
				outerMultiplier.Enqueue(outerMultiply)
			}
			formula = formula[1:i]
			fmt.Println("")

		}

		i++
		if i >= len(formula) {
			break
		}
	}

	return bigFormulas
}

func isDigit(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

func prepareFlatFormula(formula string) (flatFormula string, openBracketIndex int) {
	i := 0
	for {
		if formula[i] != '(' {
			flatFormula += string(formula[i])
		} else {
			return flatFormula, i
		}

		i++
		if i == len(formula) {
			return flatFormula, -1
		}
	}
}

type Queue []string

func (q *Queue) Enqueue(v string) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}
