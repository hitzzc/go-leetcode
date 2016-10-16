package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[byte]byte{}
	n := map[byte]byte{}
	for i := range s {
		if _, ok := m[t[i]]; !ok {
			m[t[i]] = s[i]
		}
		if _, ok := n[s[i]]; !ok {
			n[s[i]] = t[i]
		}
		if s[i] != m[t[i]] {
			return false
		}
		if t[i] != n[s[i]] {
			return false
		}
	}
	return true
}
