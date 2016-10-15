package repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
	m := map[string]int{}
	rets := []string{}
	for i := 0; i <= len(s)-10; i++ {
		sub := string(s[i : i+10])
		m[sub]++
		if m[sub] == 2 {
			rets = append(rets, sub)
		}
	}
	return rets
}
