// https://leetcode.com/problems/sort-the-students-by-their-kth-score/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	score := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	k := 2

	fmt.Println(sortTheStudents(score, k))
}

func sortTheStudents(score [][]int, k int) [][]int {
	examKthScore := getExamKthScores(score, k)
	sortedKthScore := sortByKthScores(examKthScore, k)
	sortedScores := sortWholeTable(score, sortedKthScore)

	return sortedScores
}

func getExamKthScores(scores [][]int, k int) []int {
	res := []int{}
	for _, student := range scores {
		res = append(res, student[k])
	}

	return res
}

type studentWithKthScore struct {
	studentIndex int
	score        int
}

func sortByKthScores(examKthScore []int, k int) []studentWithKthScore {
	students := make([]studentWithKthScore, len(examKthScore))

	for i := range examKthScore {
		students[i] = studentWithKthScore{
			studentIndex: i,
			score:        examKthScore[i],
		}
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].score > students[j].score
	})

	return students
}

func sortWholeTable(score [][]int, sortedKthScore []studentWithKthScore) [][]int {
	res := [][]int{}
	for _, kthScore := range sortedKthScore {
		res = append(res, score[kthScore.studentIndex])
	}

	return res
}
