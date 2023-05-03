package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/agnivade/levenshtein"
)

type queries struct {
	Question string	`json:"question"`
	Answer   string `json:"answer"`
}

func defineQuestionType(question string) string {
	addre:=regexp.MustCompile(`^\s*tambahkan\s+pertanyaan\s+(.+)\s+dengan\s+jawaban\s+(.+)$`)
	delre:=regexp.MustCompile(`^\s*hapus\s+pertanyaan\s+(.+)$`)
	switch {
	case addre.MatchString(question):
		return "add"
	case delre.MatchString(question):
		return "del"
	default:
		return "question"
	}
}

func addQuestionToDB(input string) string{
	addre:=regexp.MustCompile(`^\s*tambahkan\s+pertanyaan\s+(.+)\s+dengan\s+jawaban\s+(.+)$`)
	res:=addre.FindStringSubmatch(input)
	var q queries
	q.Question=res[1]
	q.Answer=res[2]
	DB.Create(&q)
	return "Pertanyaan berhasil ditambahkan"
}

func delQuestion(input string) string{
	delre:=regexp.MustCompile(`^\s*hapus\s+pertanyaan\s+(.+)$`)
	res:=delre.FindStringSubmatch(input)
	var q queries
	DB.Where("question = ?", res[1]).Delete(&q)
	return "Pertanyaan berhasil dihapus"
}


func getAnswerKMP(input string) string {
	var queryArr []queries
	DB.Find(&queryArr)
	bestMatch := ""
	minDist := math.MaxInt32

	for _, query := range queryArr {
		idx := KMP(query.Question, input)
		if idx != -1 {
			if query.Question == input {
				return query.Answer
			}
		}
		dist := levenshtein.ComputeDistance(query.Question, input)
		if dist < minDist {
			bestMatch = query.Answer
			minDist = dist
		}
	}
	threshold := 0.1 * float64(len(input))
	fmt.Println(threshold)
	if float64(minDist) > threshold {
		return "Pertanyaan tidak ditemukan"
	}
	return bestMatch
}


func getAnswerBM(input string) string {
	var queryArr []queries
	DB.Find(&queryArr)
	bestMatch := ""
	minDist := math.MaxInt32

	for _, query := range queryArr {
		idx := BM(query.Question, input)
		if idx != -1 {
			if query.Question == input {
				return query.Answer
			}
		}
		dist := levenshtein.ComputeDistance(query.Question, input)
		if dist < minDist {
			bestMatch = query.Answer
			minDist = dist
		}
	}
	threshold := 0.1 * float64(len(input))
	fmt.Println(threshold)
	if float64(minDist) > threshold {
		return "Pertanyaan tidak ditemukan"
	}
	return bestMatch
}






