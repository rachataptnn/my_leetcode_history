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
	input := "Mg(OH)2"

	fmt.Println(countOfAtoms(input))
}

func countOfAtoms(formula string) string {
	e := elements{}
	e.init()

	flatFormula, openBracketIndex := prepareFlatFormula(formula)
	if openBracketIndex == -1 {
		e.bro(flatFormula)

		fuck := ""
		for _, elemName := range e.elementSymbolsSorted {
			if e.elements[elemName] > 0 {
				num := strconv.Itoa(e.elements[elemName])
				if num != "1" {
					fuck += elemName + strconv.Itoa(e.elements[elemName])
				} else {
					fuck += elemName
				}
			}
		}

		return fuck
	}

	subFormulas, groupMultipliers := prepareSubFormulasAndGroupMul(formula, openBracketIndex)

	beforeSort := ""
	round := len(subFormulas)
	for i := 0; i < round; i++ {
		subFormula := beforeSort + subFormulas.Pop()
		groupMultiply := groupMultipliers.Dequeue()

		some := getElementArr(subFormula)
		groupMultiplyInt, _ := strconv.Atoi(groupMultiply)
		beforeSort = mulSubFormula(some, groupMultiplyInt)
	}

	beforeSort = flatFormula + beforeSort
	e.bro(beforeSort)

	fuck := ""
	for _, elemName := range e.elementSymbolsSorted {
		if e.elements[elemName] > 0 {
			num := strconv.Itoa(e.elements[elemName])
			if num != "1" {
				fuck += elemName + strconv.Itoa(e.elements[elemName])
			} else {
				fuck += elemName
			}
		}
	}

	return fuck
}

// elements functions
type elements struct {
	elements             map[string]int
	elementSymbolsSorted []string
}

type element struct {
	name   string
	number int
}

func (e *elements) init() {
	e.elements = map[string]int{
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
	e.elementSymbolsSorted = []string{
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

func getElementArr(subFormula string) (elements []element) {
	name := ""
	num := 0
	numLength := 0
	for {
		name = isSingleCharElement(subFormula)
		if name != "" {
			num, numLength = getElemMul(subFormula, 1)

			elements = append(elements, element{
				name:   name,
				number: num,
			})

			if num == 1 {
				subFormula = subFormula[1+(1-numLength):]
			} else {
				subFormula = subFormula[1+(numLength):]
			}

		} else {
			name = isDoubleCharElement(subFormula)
			num, numLength = getElemMul(subFormula, 2)

			elements = append(elements, element{
				name:   name,
				number: num,
			})

			if num == 1 {
				subFormula = subFormula[2+(1-numLength):]
			} else {
				subFormula = subFormula[2+(numLength):]
			}
		}

		if len(subFormula) == 0 {
			return elements
		}
	}
}

func isSingleCharElement(subFormula string) string {
	if len(subFormula) == 1 {
		return subFormula
	}

	is1stUpper := unicode.IsUpper(rune(subFormula[0]))
	is2ndUpper := unicode.IsUpper(rune(subFormula[1]))
	if is1stUpper && is2ndUpper {
		return subFormula[:1]
	}

	is2ndDigit := isDigit(string(subFormula[1]))
	if is1stUpper && is2ndDigit {
		return subFormula[:1]
	}

	return ""
}

func isDoubleCharElement(subFormula string) string {
	if len(subFormula) < 1 {
		return ""
	}

	is1stUpper := unicode.IsUpper(rune(subFormula[0]))
	is2ndLower := unicode.IsLower(rune(subFormula[1]))
	if is1stUpper && is2ndLower {
		return subFormula[:2]
	}

	return ""
}

func getElemMul(subFormula string, elemNameAmt int) (int, int) {
	digits := ""
	foundDigit := false

	if elemNameAmt == 1 {
		if len(subFormula) == 1 {
			return 1, 1
		} else {
			is2ndUpper := unicode.IsUpper(rune(subFormula[1]))
			if is2ndUpper {
				return 1, 1
			}
		}
	}
	if elemNameAmt == 2 {
		if len(subFormula) == 2 {
			return 1, 1
		} else {
			is3ndUpper := unicode.IsUpper(rune(subFormula[2]))
			if is3ndUpper {
				return 1, 1
			}
		}
	}

	i := 0
	for i < len(subFormula) {

		if len(subFormula) > 1 {
			is2ndUpper := unicode.IsUpper(rune(subFormula[1]))
			if is2ndUpper {
				return 1, 1
			}
		}

		if !isDigit(string(subFormula[i])) && foundDigit {
			result, _ := strconv.Atoi(digits)
			return result, len(digits)
		}

		if isDigit(string(subFormula[i])) {
			digits += string(subFormula[i])
			foundDigit = true
		}

		i++
	}

	if len(digits) > 0 {
		result, _ := strconv.Atoi(digits)
		return result, len(digits)
	}

	return 1, 1
}

func mulSubFormula(elements []element, groupMultiply int) string {
	result := ""
	for _, element := range elements {
		result += element.name + strconv.Itoa(element.number*groupMultiply)
	}

	return result
}

func (e *elements) bro(beforeSort string) {
	for i := 0; i < len(beforeSort); i++ {
		name := ""
		num := 0
		numLength := 0
		for {
			name = isSingleCharElement(string(beforeSort))
			if name != "" {
				num, numLength = getElemMul(string(beforeSort), 1)
				e.elements[name] += num

				if num == 1 {
					beforeSort = beforeSort[1+(1-numLength):]
				} else {
					beforeSort = beforeSort[1+(numLength):]
				}

			} else {
				name = isDoubleCharElement(string(beforeSort))
				num, numLength = getElemMul(string(beforeSort), 2)
				e.elements[name] += num

				if num == 1 {
					beforeSort = beforeSort[2+(1-numLength):]
				} else {
					beforeSort = beforeSort[2+(numLength):]
				}

			}

			if len(beforeSort) == 0 {
				break
			}
		}
	}
}

// stack functions
type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	member := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return member
}

type Queue []string

func (q *Queue) Enqueue(v string) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() string {
	element := (*q)[0]
	*q = (*q)[1:]
	return element
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

func prepareSubFormulasAndGroupMul(formula string, openBracketIndex int) (subFormulas stack, groupMultipliers Queue) {
	i := openBracketIndex
	subText := ""
	for {
		if formula[i] != '(' && formula[i] != ')' {
			subText += string(formula[i])
		} else {
			if isDigit(subText) {
				groupMultipliers.Enqueue(subText)
			} else if subText != "" {
				subFormulas.Push(subText)
			}
			subText = ""
		}

		if i+1 == len(formula) && isDigit(subText) {
			groupMultipliers.Enqueue(subText)
		}

		i++
		if i == len(formula) {
			return subFormulas, groupMultipliers
		}
	}
}

func isDigit(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}
