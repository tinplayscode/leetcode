package atoi

func myAtoi(s string) (result int) {
	result = 0
	mark := 1
	marked := false
	readNumber := false
	whitespaced := false

loop:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '-':
			if readNumber {
				break loop
			}
			if marked || readNumber {
				return 0
			}

			if mark == -1 {
				mark = 1
			} else {
				mark = -1
			}
			marked = true

			break
		case s[i] == '+':
			if readNumber {
				break loop
			}
			if marked || readNumber {
				return 0
			}
			marked = true
			break

		case s[i] == ' ':
			if readNumber {
				break loop
			}
			if marked {
				return 0
			}
			whitespaced = true

			break

		case s[i] >= '0' && s[i] <= '9':
			if readNumber && whitespaced {
				return 0
			}

			if !readNumber && i+1 < len(s) && !(s[i+1] >= '0' && s[i+1] <= '9') {
				readNumber = true
			}
			result = result*10 + int(s[i]-'0')

			if mark == 1 && result > 2147483647 {
				return 2147483647
			}
			if mark == -1 && result > 2147483648 {
				return -2147483648
			}
			break
		default:
			break loop
		}
	}

	result = result * mark

	return
}

func Test() {
	//println(myAtoi("    -88827   5655  U"))
	//println(myAtoi("-5-"))
	//println(myAtoi("-13+-"))
	println(myAtoi("123   456"))
}
