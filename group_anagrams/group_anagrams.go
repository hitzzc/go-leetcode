package group_anagrams

import (
	"sort"
)

type bytes []byte

func (this bytes) Len() int {
	return len(this)
}

func (this bytes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this bytes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	indexMap := map[string]int{}
	for i := range strs {
		key := bytes(strs[i])
		sort.Sort(key)
		keyString := string(key)
		if _, ok := indexMap[keyString]; !ok {
			indexMap[keyString] = len(ret)
			ret = append(ret, []string{strs[i]})
		} else {
			ret[indexMap[keyString]] = append(ret[indexMap[keyString]], strs[i])
		}
	}
	return ret
}
