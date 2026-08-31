func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)
	for j:=0; j+l <=len(s2); j++ {
		fmt.Println(s2[j:j+l])
		if isAnagram(s1, s2[j:j+l]) {
			return true
		}
	}

	return false
}

func isAnagram(s1 string, s2 string) bool {
	flag := make([]int, 26)

	for i:=0; i<len(s1); i++ {
		flag[s1[i]-'a']+=1
		flag[s2[i]-'a']-=1
	}

	for i:=0; i<26; i++ {
		if flag[i]!=0 {
			return false
		}
	}

	return true
}