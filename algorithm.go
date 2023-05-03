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

