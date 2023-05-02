package main

import "fmt"

func getDayOfWeek(date string) string {
	newDate := extractDate(date)
    day, month, year := parseDate(newDate)
    if day == 0 || month == 0 || year == 0 {
        return "Invalid date format"
    }
    daysOfWeek := [...]string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
    q := day
    m := month
    if month < 3 {
        m += 12
        year--
    }
    k := year % 100
    j := year / 100
    h := (q + ((13*(m+1))/5) + k + (k/4) + (j/4) + (5*j)) % 7
	if h == 0 {
		return daysOfWeek[6]
	}
	return daysOfWeek[h-1]
}

func parseDate(date string) (int, int, int) {
    var day, month, year int
    _, err := fmt.Sscanf(date, "%d/%d/%d", &day, &month, &year)
    if err != nil {
        return 0, 0, 0
    }
    return day, month, year
}

func extractDate(str string) string {
    prefix := "hari apa tanggal "
    if len(str) <= len(prefix) {
        return ""
    }
    dateStr := str[len(prefix):]
    return dateStr
}


