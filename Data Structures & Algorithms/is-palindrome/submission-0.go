func isPalindrome(s string) bool {
    mydict := map[rune]rune{
		'a': 'a',
		'b': 'b',
		'c': 'c',
		'd': 'd',
		'e': 'e',
		'f': 'f',
		'g': 'g',
		'h': 'h',
		'i': 'i',
		'j': 'j',
		'k': 'k',
		'l': 'l',
		'm': 'm',
		'n': 'n',
		'o': 'o',
		'p': 'p',
		'q': 'q',
		'r': 'r',
		's': 's',
		't': 't',
		'u': 'u',
		'v': 'v',
		'w': 'w',
		'x': 'x',
		'y': 'y',
		'z': 'z',
		'A': 'a',
		'B': 'b',
		'C': 'c',
		'D': 'd',
		'E': 'e',
		'F': 'f',
		'G': 'g',
		'H': 'h',
		'I': 'i',
		'J': 'j',
		'K': 'k',
		'L': 'l',
		'M': 'm',
		'N': 'n',
		'O': 'o',
		'P': 'p',
		'Q': 'q',
		'R': 'r',
		'S': 's',
		'T': 't',
		'U': 'u',
		'V': 'v',
		'W': 'w',
		'X': 'x',
		'Y': 'y',
		'Z': 'z',
        '0':'0',
        '1':'1',
        '2':'2',
        '3':'3',
        '4':'4',
        '5':'5',
        '6':'6',
        '7':'7',
        '8':'8',
        '9':'9',
	}

	var result string
	for _, char := range s {
		if v, ok := mydict[char]; ok {
			result = result + string(v)
		}
	}

	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-i-1] {
			return false
		}
	}

	return true

}
