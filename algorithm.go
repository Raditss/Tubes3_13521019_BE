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

func levenshteinDistance(s, t string) int {
    n := len(s)
    m := len(t)
    if n == 0 {
        return m
    }
    if m == 0 {
        return n
    }

    // Create a matrix to store the distances between prefixes of s and t
    // The (i,j)th entry of the matrix represents the distance between the first i characters of s and the first j characters of t
    matrix := make([][]int, n+1)
    for i := range matrix {
        matrix[i] = make([]int, m+1)
    }

    // Initialize the first row and column of the matrix
    for i := 0; i <= n; i++ {
        matrix[i][0] = i
    }
    for j := 0; j <= m; j++ {
        matrix[0][j] = j
    }

    // Fill in the rest of the matrix
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            substitutionCost := 1
            if s[i-1] == t[j-1] {
                substitutionCost = 0
            }
            matrix[i][j] = min(matrix[i-1][j]+1, min(matrix[i][j-1]+1, matrix[i-1][j-1]+substitutionCost))
        }
    }

    // The distance between s and t is the bottom right entry of the matrix
    return matrix[n][m]
}

