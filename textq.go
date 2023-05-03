package main

import (
	"regexp"
	"math"
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

func findClosestMatch(queryArr []queries, query string) (queries, int) {
    var bestMatch queries
    minDist := math.MaxInt64
    for _, q := range queryArr {
        idx := KMP(q.Question, query)
        if idx != -1 {
            dist := levenshtein.ComputeDistance(q.Question, query)
                minDist = dist
                bestMatch = q
            }
        }
		return bestMatch, minDist
    }




