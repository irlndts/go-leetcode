package leetcode

// https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) string {
	latin := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	result := ""

	for _, d := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		if offset := num / d; offset > 0 {
			result += buildRoman(offset, latin[d])
			num = num - (offset * d)
		}
	}

	return result
}

func buildRoman(length int, l string) string {
	str := ""
	for i := 0; i < length; i++ {
		str += l
	}
	return str
}
