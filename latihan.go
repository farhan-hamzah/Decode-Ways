package main

import (
	"fmt"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	// dp[i] = jumlah cara decoding untuk s[0:i]
	dp := make([]int, n+1)
	dp[0] = 1 // cara decode string kosong
	dp[1] = 1 // satu cara decode satu karakter pertama (jika bukan '0')

	for i := 2; i <= n; i++ {
		// Cek satu digit terakhir
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// Cek dua digit terakhir
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numDecodings("12"))    // Output: 2 ("AB", "L")
	fmt.Println(numDecodings("226"))   // Output: 3 ("BZ", "VF", "BBF")
	fmt.Println(numDecodings("06"))    // Output: 0 (invalid)
	fmt.Println(numDecodings("11106")) // Output: 2
}
