package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		e := magazine[i]
		m[e]++
	}
	for i := 0; i < len(ransomNote); i++ {
		e := ransomNote[i]
		if m[e] == 0 {
			return false
		}
		m[e]--
	}
	return true
}
