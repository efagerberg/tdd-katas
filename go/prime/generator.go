package generator

func primeNumberGenerator() func() int {
	var previousPrimes []int
	currentNumber := 2

	f := func() int {
		checkPrime(previousPrimes, &currentNumber)
		previousPrimes = append(previousPrimes, currentNumber)
		currentNumber++
		return currentNumber - 1
	}

	return f
}

func checkPrime(previousPrimes []int, currentNumber *int) {
	for _, previousPrime := range previousPrimes {
		if *currentNumber%previousPrime == 0 {
			*currentNumber++
			checkPrime(previousPrimes, currentNumber)
		}
	}
}
