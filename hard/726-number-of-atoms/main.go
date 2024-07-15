// https://leetcode.com/problems/number-of-atoms/description/?envType=daily-question&envId=2024-07-14

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	// input := "H20"

	// input := "K4(ON(SO3)2)2"
	// input := "Mg(OH)2"

	// 13/31
	// input := "((HHe28Be26He)9)34"

	//          s  e   s                     e   s                  e   s                  e   s  e
	// input := "((N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"
	// asis ตอนนี้ prepare fn คืน stack, queue เดียว
	// tobe เปลี่ยนเป็น prepare fn ต้องคืน []stack, []queue

	// note: day 2
	// หลังจากปวดหัวตอนเช้าเลยไปแง้มดู solution แต่ละคน
	// code ยาวทุกคน ที่เขียนสั้นก็ไม่เข้าใจ
	// แต่มีคนนึง https://leetcode.com/u/reas0ner/
	// เขียนไว้ว่าไม่ใช้ stack
	// วิธีเค้าคือ เปลี่ยนประโยคให้มี สัญลัก + * แล้วใช้ add(), mul()
	// เอาจริงนี่ก็ยังไม่เกทอยู่ดี
	// แต่เหมือนอะไรสักอย่างดลใจให้ลองทำคูณกระจาย
	// แนวคิดคือ ถ้าเจอ '(' แล้วเจอ ')' นั่นก็คือ shorted formula
	// มันคือ formula ที่เล็กที่สุด ที่ควรคูณเข้าไป
	// แล้วก็ลบวงเล็บออก
	// แล้วก็ replace ผลคูณนั้นไป
	// พอเรา recursive วิธีเชี่ยนี่ไปเรื่อยๆ
	// สุดท้าย formula เราจะไม่เหลือ () เลย
	// อันนี้ก็ง่ายละ ไปบวกพอ -> summaryAtomEachElement()
	// สุดท้ายก็ให้ gen ai ไป sort ธาตุมาให้แล้วก็ map ผลรวมของ formula ตามที่ sort จ้บบ
	// ละไอ solution เวนนี่ตอนเขียนเสร็จแล้วเทสผ่านไป 28/31 คือกุนึกถึง pattern matching & term rewritting เฉยเลยถถถถ

	// 28/31
	input := "Mg(H2O)N"

	fmt.Println(countOfAtoms(input))
}

type bracketWithIndex struct {
	bracket string
	index   int
}

func countOfAtoms(formula string) string {
	_, openBracketIndex := getFlatFormula(formula)

	var currBracket, lastBracket bracketWithIndex
	for !isNoBracketsLeft(formula) {
		formula = peelOuterMultipliers(formula, openBracketIndex, currBracket, lastBracket)
		fmt.Println("")
	}

	e := elements{}
	e.init()
	e.summaryAtomEachElement(formula)
	sortedFormula := e.concatFormulaBySortedElement()

	return sortedFormula
}

func getFlatFormula(formula string) (flatFormula string, openBracketIndex int) {
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

func peelOuterMultipliers(formula string, index int, currBracket, lastBracket bracketWithIndex) string {
	if isBracket(formula[index]) {
		currBracket.bracket = string(formula[index])
		currBracket.index = index

		if lastBracket.index < currBracket.index {
			if isShortestBracket(currBracket, lastBracket) {
				flatFormula := string(formula[:lastBracket.index])

				mappedFormula := mapMultiplyToFormula(formula, lastBracket.index, currBracket.index)

				paddingGroupMul := getPaddingGroupMul(formula[currBracket.index:])
				tailFormula := string(formula[currBracket.index+paddingGroupMul:])

				formula = flatFormula + mappedFormula + tailFormula
				return formula
			} else {
				lastBracket.bracket = string(formula[index])
				lastBracket.index = index
			}
		} else if currBracket.bracket != "" {
			lastBracket.bracket = string(formula[index])
			lastBracket.index = index
		}
	}

	return peelOuterMultipliers(formula, index+1, currBracket, lastBracket)
}

func isBracket(token byte) bool {
	tokenStr := string(token)
	return tokenStr == "(" || tokenStr == ")"
}

func isShortestBracket(currBracket, lastBracket bracketWithIndex) bool {
	if lastBracket.bracket == "(" && currBracket.bracket == ")" {
		return true
	}

	return false
}

func isNoBracketsLeft(formula string) bool {
	for _, v := range formula {
		if v == '(' || v == ')' {
			return false
		}
	}

	return true
}

func mapMultiplyToFormula(formula string, startIndex, endIndex int) string {
	groupMul := getGroupMul(formula, endIndex)

	indexAfterStartBracket := startIndex + 1
	subFormula := formula[indexAfterStartBracket:endIndex]

	newSubFormula := ""
	for len(subFormula) != 0 && !isDigit(subFormula) {
		if isSingleCharElement(subFormula) {
			elementMul, mulPadding := getElementMul(subFormula, 1)
			multipliedAtom := elementMul * groupMul
			newSubFormula += subFormula[:1] + strconv.Itoa(multipliedAtom)

			if elementMul == 1 {
				subFormula = subFormula[1+(1-mulPadding):]
			} else {
				subFormula = subFormula[1+(mulPadding):]
			}

		} else if isDoubleCharElement(subFormula) {
			elementMul, mulPadding := getElementMul(subFormula, 2)
			multipliedAtom := elementMul * groupMul
			newSubFormula += subFormula[:2] + strconv.Itoa(multipliedAtom)

			if elementMul == 1 {
				subFormula = subFormula[2+(1-mulPadding):]
			} else {
				subFormula = subFormula[2+(mulPadding):]
			}
		}
	}

	return newSubFormula
}

func isSingleCharElement(subFormula string) bool {
	if len(subFormula) == 1 {
		return true
	}

	is1stUpper := unicode.IsUpper(rune(subFormula[0]))
	is2ndUpper := unicode.IsUpper(rune(subFormula[1]))
	if is1stUpper && is2ndUpper {
		return true
	}

	is2ndDigit := isDigit(string(subFormula[1]))
	if is1stUpper && is2ndDigit {
		return true
	}

	return false
}

func isDoubleCharElement(formula string) bool {
	if len(formula) < 2 {
		return false
	}

	is1stUpper := unicode.IsUpper(rune(formula[0]))
	is2ndLower := unicode.IsLower(rune(formula[1]))
	if is1stUpper && is2ndLower {
		return true
	}

	return false
}

func getElementMul(subFormula string, padding int) (int, int) {
	indexEndMul := len(subFormula)
	for i := padding; i < len(subFormula); i++ {
		if !isDigit(string(subFormula[i])) {
			indexEndMul = i
			break
		}
	}

	mulStr := subFormula[padding:indexEndMul]
	mulNum, _ := strconv.Atoi(mulStr)

	if mulNum != 0 {
		return mulNum, len(mulStr)
	}
	return 1, 1
}

func getGroupMul(formula string, endIndex int) int {
	indexEndMul := len(formula)
	indexAfterEndBracket := endIndex + 1

	for i := indexAfterEndBracket; i < len(formula); i++ {
		if !isDigit(string(formula[i])) {
			indexEndMul = i
			break
		}
	}

	mulStr := formula[indexAfterEndBracket:indexEndMul]
	mulNum, _ := strconv.Atoi(mulStr)

	return mulNum
}

func isDigit(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

func getPaddingGroupMul(tailFormula string) int {
	indexEndMul := len(tailFormula)
	for i := 1; i < len(tailFormula); i++ {
		if !isDigit(string(tailFormula[i])) {
			indexEndMul = i
			break
		}
	}

	return indexEndMul
}

type elements struct {
	elementWithAtomAmt map[string]int
	sortedElements     []string
}

func (e *elements) init() {
	e.elementWithAtomAmt = map[string]int{
		"H": 0, "He": 0, "Li": 0, "Be": 0, "B": 0, "C": 0, "N": 0, "O": 0, "F": 0, "Ne": 0,
		"Na": 0, "Mg": 0, "Al": 0, "Si": 0, "P": 0, "S": 0, "Cl": 0, "Ar": 0, "K": 0, "Ca": 0,
		"Sc": 0, "Ti": 0, "V": 0, "Cr": 0, "Mn": 0, "Fe": 0, "Co": 0, "Ni": 0, "Cu": 0, "Zn": 0,
		"Ga": 0, "Ge": 0, "As": 0, "Se": 0, "Br": 0, "Kr": 0, "Rb": 0, "Sr": 0, "Y": 0, "Zr": 0,
		"Nb": 0, "Mo": 0, "Tc": 0, "Ru": 0, "Rh": 0, "Pd": 0, "Ag": 0, "Cd": 0, "In": 0, "Sn": 0,
		"Sb": 0, "Te": 0, "I": 0, "Xe": 0, "Cs": 0, "Ba": 0, "La": 0, "Ce": 0, "Pr": 0, "Nd": 0,
		"Pm": 0, "Sm": 0, "Eu": 0, "Gd": 0, "Tb": 0, "Dy": 0, "Ho": 0, "Er": 0, "Tm": 0, "Yb": 0,
		"Lu": 0, "Hf": 0, "Ta": 0, "W": 0, "Re": 0, "Os": 0, "Ir": 0, "Pt": 0, "Au": 0, "Hg": 0,
		"Tl": 0, "Pb": 0, "Bi": 0, "Po": 0, "At": 0, "Rn": 0, "Fr": 0, "Ra": 0, "Ac": 0, "Th": 0,
		"Pa": 0, "U": 0, "Np": 0, "Pu": 0, "Am": 0, "Cm": 0, "Bk": 0, "Cf": 0, "Es": 0, "Fm": 0,
		"Md": 0, "No": 0, "Lr": 0, "Rf": 0, "Db": 0, "Sg": 0, "Bh": 0, "Hs": 0, "Mt": 0, "Ds": 0,
		"Rg": 0, "Cn": 0, "Nh": 0, "Fl": 0, "Mc": 0, "Lv": 0, "Ts": 0, "Og": 0,
	}
	e.sortedElements = []string{
		"Ac", "Ag", "Al", "Am", "Ar", "As", "At", "Au", "B", "Ba", "Be", "Bh", "Bi", "Bk",
		"Br", "C", "Ca", "Cd", "Ce", "Cf", "Cl", "Cm", "Cn", "Co", "Cr", "Cs", "Cu", "Db",
		"Ds", "Dy", "Er", "Es", "Eu", "F", "Fe", "Fl", "Fm", "Fr", "Ga", "Gd", "Ge", "H",
		"He", "Hf", "Hg", "Ho", "Hs", "I", "In", "Ir", "K", "Kr", "La", "Li", "Lr", "Lu",
		"Lv", "Mc", "Md", "Mg", "Mn", "Mo", "Mt", "N", "Na", "Nb", "Nd", "Ne", "Nh", "Ni",
		"No", "Np", "O", "Og", "Os", "P", "Pa", "Pb", "Pd", "Pm", "Po", "Pr", "Pt", "Pu",
		"Ra", "Rb", "Re", "Rf", "Rg", "Rh", "Rn", "Ru", "S", "Sb", "Sc", "Se", "Sg", "Si",
		"Sm", "Sn", "Sr", "Ta", "Tb", "Tc", "Te", "Th", "Ti", "Tl", "Tm", "Ts", "U", "V",
		"W", "Xe", "Y", "Yb", "Zn", "Zr",
	}
}

func (e *elements) summaryAtomEachElement(formula string) {
	name := ""
	for i := 0; i < len(formula); i++ {
		for {
			if isSingleCharElement(formula) {
				name = formula[:1]
				elementMul, mulPadding := getElementMul(formula, 1)
				e.elementWithAtomAmt[name] += elementMul

				if elementMul == 1 {
					formula = formula[1+(1-mulPadding):]
				} else {
					formula = formula[1+(mulPadding):]
				}
			}

			if isDoubleCharElement(formula) {
				name = formula[:2]
				elementMul, mulPadding := getElementMul(formula, 2)
				e.elementWithAtomAmt[name] += elementMul

				if elementMul == 1 {
					formula = formula[2+(1-mulPadding):]
				} else {
					formula = formula[2+(mulPadding):]
				}
			}

			if len(formula) == 0 {
				break
			}
		}
	}
}

func (e *elements) concatFormulaBySortedElement() (sortedFormula string) {
	for _, elemName := range e.sortedElements {
		if e.elementWithAtomAmt[elemName] > 0 {
			num := strconv.Itoa(e.elementWithAtomAmt[elemName])
			if num != "1" {
				sortedFormula += elemName + strconv.Itoa(e.elementWithAtomAmt[elemName])
			} else {
				sortedFormula += elemName
			}
		}
	}

	return sortedFormula
}
