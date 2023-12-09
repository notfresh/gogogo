package morestrings

func ReverseRunes(s string) string{
	r := []rune(s)
	for i, j := 0, len(s) -1 ; i < len(r)/2; i, j = i+1, j-1{
		r[i], r[j] = r[j], r[i];
	}	
	return string(r)
}