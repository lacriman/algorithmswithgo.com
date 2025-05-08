package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	newNumber := []rune{}
	if dec == 0 {
		return "0"
	}

	const digits = "0123456789ABCDEFG"

	for dec > 0 {
		rem := dec % base
		newNumber = append(newNumber, rune(digits[rem]))
		dec = dec / base
	}
	return Reverse(string(newNumber))
}
