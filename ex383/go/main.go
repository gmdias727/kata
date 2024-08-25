package main

// func main() {
// 	canConstruct
// }

func canConstruct(ransomNote string, magazine string) bool {
	ransomMap := make(map[rune]int)
	magazineMap := make(map[rune]int)
	for _, char := range magazine {
		magazineMap[char]++
	}
	for _, char := range ransomNote {
		ransomMap[char]++
	}
	for char, count := range ransomMap {
		if magazineMap[char] < count {
			return false
		}
	}
	return true
}
