package main

import (
	"regexp"

)

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

