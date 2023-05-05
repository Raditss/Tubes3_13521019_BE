package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

type queries struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func defineQuestionType(question string) string {
	addre := regexp.MustCompile(`^\s*tambahkan\s+pertanyaan\s+(.+)\s+dengan\s+jawaban\s+(.+)$`)
	delre := regexp.MustCompile(`^\s*hapus\s+pertanyaan\s+(.+)$`)
	switch {
	case addre.MatchString(question):
		return "add"
	case delre.MatchString(question):
		return "del"
	default:
		return "question"
	}
}

func addQuestionToDB(input string) string {
	addre := regexp.MustCompile(`^\s*tambahkan\s+pertanyaan\s+(.+)\s+dengan\s+jawaban\s+(.+)$`)
	res := addre.FindStringSubmatch(input)
	var q queries
	var temp queries
	// check if question already exists
	DB.Where("question = ?", res[1]).First(&temp)
	if temp.Question == res[1] {
		temp.Answer = res[2]
		// save to db with where condition
		DB.Model(&q).Where("question = ?", res[1]).Update("answer", res[2])
		return "Pertanyaan sudah diupdate"
	}

	q.Question = res[1]
	q.Answer = res[2]
	DB.Create(&q)
	return "Pertanyaan berhasil ditambahkan"
}

func delQuestion(input string) string {
	delre := regexp.MustCompile(`^\s*hapus\s+pertanyaan\s+(.+)$`)
	res := delre.FindStringSubmatch(input)
	var q queries
	DB.Where("question = ?", res[1]).Delete(&q)
	return "Pertanyaan berhasil dihapus"
}

type queryResult struct {
	Question string
	Answer   string
	Dist     int
}

func getAnswerKMP(input string) string {
	var queryArr []queries
	DB.Find(&queryArr)

	var results []queryResult
	minDist := math.MaxInt32

	for _, query := range queryArr {
		idx := KMP(query.Question, input)
		if idx != -1 {
			if query.Question == input {
				return query.Answer
			}
		}
		dist := levenshteinDistance(query.Question, input)
		if dist < minDist {
			minDist = dist
			results = []queryResult{{Question: query.Question, Answer: query.Answer, Dist: dist}}
		} else if dist == minDist {
			results = append(results, queryResult{Question: query.Question, Answer: query.Answer, Dist: dist})
		}
	}
	threshold := 0.1 * float64(len(input))
	if float64(minDist) > threshold {
		topResults := make([]string, 0, 3)
		sort.Slice(results, func(i, j int) bool {
			return results[i].Dist < results[j].Dist
		})
		for i := 0; i < len(results) && i < 3; i++ {
			topResults = append(topResults, results[i].Question)
		}
		return fmt.Sprintf("Pertanyaan tidak ditemukan. Pertanyaan yang mirip: %s", strings.Join(topResults, ", "))
	}
	return results[0].Answer
}

func getAnswerBM(input string) string {
	var queryArr []queries
	DB.Find(&queryArr)

	var results []queryResult
	minDist := math.MaxInt32

	for _, query := range queryArr {
		idx := BM(query.Question, input)
		if idx != -1 {
			if query.Question == input {
				return query.Answer
			}
		}
		dist := levenshteinDistance(query.Question, input)
		if dist < minDist {
			minDist = dist
			results = []queryResult{{Question: query.Question, Answer: query.Answer, Dist: dist}}
		} else if dist == minDist {
			results = append(results, queryResult{Question: query.Question, Answer: query.Answer, Dist: dist})
		}
	}
	threshold := 0.1 * float64(len(input))
	if float64(minDist) > threshold {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Dist < results[j].Dist
		})
		return fmt.Sprintf("Pertanyaan tidak ditemukan. Pertanyaan yang mirip: %s, %s, %s", results[0].Question, results[1].Question, results[2].Question)
	}
	return results[0].Answer
}
