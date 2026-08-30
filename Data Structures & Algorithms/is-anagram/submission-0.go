func isAnagram(s string, t string) bool {
	alphabets := make([]int, 26)

	for _, v := range s {
		alphabets[v-97]++
	}
	for _, v := range t {
		alphabets[v-97]--
	}
	for _, v := range alphabets {
		if v != 0 {
			return false
		}
	}
	return true

}