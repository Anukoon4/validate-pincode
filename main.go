package main

import "strconv"

func validatePIN(pin string) bool {
	//1 Check the length of the PIN
	if len(pin) < 6 {
		return false
	}

	//2 Check for more than three repeating digits
	count := 1
	for i := 1; i < len(pin); i++ {
		if pin[i] == pin[i-1] {
			count++
			if count > 3 {
				return false
			}
		} else {
			count = 1
		}
	}

	//3 Check for more than two sequential digits
	for i := 0; i < len(pin)-2; i++ {
		digit1, _ := strconv.Atoi(string(pin[i]))
		digit2, _ := strconv.Atoi(string(pin[i+1]))
		digit3, _ := strconv.Atoi(string(pin[i+2]))

		if digit1+1 == digit2 && digit2+1 == digit3 {
			return false
		}
	}

	//4. Check for more than two pairs of repeating digits
	pairCount := 0
	for i := 0; i < len(pin)-1; i++ {
		if pin[i] == pin[i+1] {
			pairCount++
			if pairCount > 2 {
				return false
			}
			// Skip the next digit since it's part of the current pair
			i++
		}
	}
	// Attempt to convert the PIN to an integer
	_, err := strconv.Atoi(pin)
	if err != nil {
		return false
	}
	return true
}
