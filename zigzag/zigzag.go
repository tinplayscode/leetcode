package zigzag

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := ""

	relativePos := 0
	if numRows >= 3 {
		relativePos = numRows - 2
	}

	for i := 0; i < numRows; i++ {
		j := i

		for j < len(s) {
			result = result + string(s[j])

			if i == 0 || i == numRows-1 {
				j = j + numRows + relativePos
			} else {
				pos := j + numRows - i - 1 + relativePos - i + 1

				if pos < len(s) {
					result = result + string(s[pos])
				}
				j = j + numRows + relativePos
			}

		}
	}

	return result
}

func TestConvert() {
	s := "PAYPALISHIRING"
	numRows := 3
	result := convert(s, numRows)
	println(result)

	if result != "PAHNAPLSIIGYIR" {
		println("TestConvert failed")
	}
}
