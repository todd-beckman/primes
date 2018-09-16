package primes

func GetPrimesUpTo200() []int {
	return []int{
		2, 3, 5, 7, 11, 13, 17, 19,
		23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107,
		109, 113, 127, 131, 137, 139,
		149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199,
	}
}

func GetPrimeMapUpTo200() map[int]bool {
	return map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true,
		17: true, 19: true, 23: true, 29: true, 31: true, 37: true,
		41: true, 43: true, 47: true, 53: true, 59: true, 61: true,
		67: true, 71: true, 73: true, 79: true, 83: true, 89: true,
		97: true, 101: true, 103: true, 107: true, 109: true,
		113: true, 127: true, 131: true, 137: true, 139: true,
		149: true, 151: true, 157: true, 163: true, 167: true,
		173: true, 179: true, 181: true, 191: true, 193: true,
		197: true, 199: true,
	}
}

func GetPrimeFactorsUpTo100(n int) []int {
	knownPrimes := GetPrimeMapUpTo200()

	if knownPrimes[n] {
		return []int{n}
	}

	primeFactors := []int{}

	for i, _ := range knownPrimes {
		if n%i == 0 {
			primeFactors = append(primeFactors, i)
		}

		if i*i > n {
			break
		}
	}

	return primeFactors
}
