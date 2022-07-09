package main

var romanMap map[string]int

func romanToInt(s string) int {

	// Make a map of Roman Number to their base 10 number value
	romanMap = make(map[string]int)
	romanMap["I"], romanMap["V"], romanMap["X"], romanMap["L"], romanMap["C"], romanMap["D"], romanMap["M"] = 1, 5, 10, 50, 100, 500, 1000

	romanAmount := 0                      // Set the amount to zero
	previousNum := romanMap[string(s[0])] // previous value to skip first character in loop

	// Iterate over each character
	// Need to know what comes after to know to add or subtract
	for i := 1; i < len(s); i++ {
		if romanMap[string(s[i])] > previousNum {
			romanAmount -= previousNum
		} else {
			romanAmount += previousNum
		}

		previousNum = romanMap[string(s[i])]
	}

	romanAmount += previousNum

	return romanAmount
}
