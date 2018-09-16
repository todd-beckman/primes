package primes

const maxK = 100 * 100

/**

Checks if n is prime by checking for factors in the form ka+b
up to the square root of n.

For k>6 this gives a larger overhead; defaults to 6.
*/
func SimpleCheckIfPrime(n, k int) bool {
	// First check the case where a=0
	if n < 2 { // negative, 0, and 1 are not prime
		return false
	} else if n < 4 { // 2 and 3 are prime
		return true
	}

	// Check if n is even every time because this is basically free
	if n == (n>>1)<<1 {
		return false
	}

	// Short circuit: we already know all primes up to 210
	if n < 211 {
		return GetPrimeMapUpTo200()[n]
	}

	// Use k=6 by default because it's fast
	if k > n || k < 7 || k > maxK {
		return simpleCheckIfPrimeWithK6(n)
	}

	return simpleCheckIfPrimeWithSpecifiedK(n, k)
}

// Checks if n is prime by searching all factors 6a+b
// up to the square root of n.
func simpleCheckIfPrimeWithK6(n int) bool {
	if n%3 == 0 { // need to check 2 and 3 and we already checked 2
		return false
	}

	// We only check each case where b=1 and b=5.
	// In this loop, we check i=6a+5 and i=6a+7
	// - i = 6a+5 trivially checks all cases where b=5
	// - i = 6a+7 = 6(a+1) + 1
	//     - Since we are checking for all a,
	//       i=6a+7 checks all cases where b=1
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	// No divisors found
	return true
}

// Checks for primality by searching for its divisors.
// Check all numbers in the form ka + b where b<k
// up to the square root of n. n is prime if no divisors
// are found.
func simpleCheckIfPrimeWithSpecifiedK(n, oldk int) bool {
	knownPrimes := GetPrimesUpTo200()

	numPrimes := 1  // Number of primes to check directly
	k := 1          // Adjusted k to adjust for suboptimal input
	numBChecks := 0 // Number of primes after numPrimes for b vals

	// First find all primes up to sqrt(k)
	// Then find all prime from that to k
	for i := range knownPrimes {
		prime := knownPrimes[i]

		if prime*prime <= oldk {
			k *= prime // calculating the optimal k
			numPrimes++
		} else if prime < k {
			numBChecks++
		} else {
			break
		}
	}

	// Skip 2 since we checked it already
	primeFactors := knownPrimes[1:numPrimes]

	// Now check if any prime factors of k divide n
	// For example, for k=30 (factors 2, 3, 5) we check:
	//
	// b = 0, 2, 4, 6, 8, 12, 14, 16, 18, 22, 24, 26, 28
	// b = 3, 9, 15, 21, 27
	// b = 5, 25
	for i := range primeFactors {
		prime := primeFactors[i]

		if n%prime == 0 {
			return false
		}

		if prime*prime > n {
			break
		}
	}

	// Now we need to check b for prime factors in the form ka+b
	// k is known and we can iterate on a. We already know b but we
	// need to figure out which b's still need checking.
	//
	// We still need to check each b:
	// b = 1, 7, 11, 13, 17, 19, 23, 29
	//
	// Which happens to be 1 and all the primes not dividing into k
	// up to k
	bValues := knownPrimes[numPrimes : numPrimes+numBChecks-1]
	bValues = append(bValues, 1)

	for i := k; i*i <= n; i += k {
		for b := range bValues {
			if n%(i+b) == 0 {
				return false
			}
		}
	}

	return true
}
