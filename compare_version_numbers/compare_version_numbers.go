package compare_version_numbers

func compareVersion(version1 string, version2 string) int {
	bytes1, bytes2 := []byte(version1), []byte(version2)
	bToI := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	i, j := -1, -1
	for i < len(bytes1) || j < len(bytes2) {
		var num1, num2 int
		for i++; i < len(bytes1) && bytes1[i] != '.'; i++ {
			num1 = num1*10 + bToI[bytes1[i]]
		}
		for j++; j < len(bytes2) && bytes2[j] != '.'; j++ {
			num2 = num2*10 + bToI[bytes2[j]]
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}
