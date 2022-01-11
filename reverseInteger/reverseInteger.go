package reverseInteger

func reverse(x int) int {
	var result int = 0

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	negative := 1
	if x < 0 {
		negative = -1
	}

	//check for overflow of 32bit signed
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result * negative
}

func Test() {
	println(reverse(123))
	println(reverse(-123))
	println(reverse(120))
	println(reverse(1534236469))

}
