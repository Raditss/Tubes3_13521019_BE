package main


func KMP(text, pattern string) int {
    n, m := len(text), len(pattern)
    if m == 0 {
        return 0
    }
    if n < m {
        return -1
    }

    // Compute LPS array
    lps := make([]int, m)
    i, j := 1, 0
    for i < m {
        if pattern[i] == pattern[j] {
            j++
            lps[i] = j
            i++
        } else if j > 0 {
            j = lps[j-1]
        } else {
            lps[i] = 0
            i++
        }
    }

    // Perform string matching
    i, j = 0, 0
    for i < n {
        if pattern[j] == text[i] {
            i++
            j++
            if j == m {
                return i - j
            }
        } else if j > 0 {
            j = lps[j-1]
        } else {
            i++
        }
    }

    return -1
}

func BM(text, pattern string) int {
    n := len(text)
    m := len(pattern)

    if m == 0 {
        return 0
    }

    // build the bad character table
    bc := make(map[rune]int)
    for i := 0; i < m-1; i++ {
        bc[rune(pattern[i])] = m - i - 1
    }

    // build the good suffix table
    suffixes := make([]int, m)
    f := make([]int, m+1)
    var j int
    var k int
    for i := m - 1; i >= 0; i-- {
        for j > 0 && pattern[j-1] != pattern[i] {
            suffixes[j-1] = k
            j = f[j]
        }
        if j > 0 && pattern[j-1] == pattern[i] {
			j--
		} else {
			k = i
		}
        suffixes[i] = k
        f[i] = j
    }
    for j > 0 {
        suffixes[j-1] = k
        j = f[j]
    }

    // search the pattern in the text
    i := m - 1
    j = m - 1
    for i < n && j >= 0 {
        if text[i] == pattern[j] {
            i--
            j--
        } else {
            shift := max(bc[rune(text[i])], suffixes[j])
            i += m - min(j, 1+shift)
            j = m - 1
        }
    }
    if j < 0 {
        return i + 1
    }
    return -1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
