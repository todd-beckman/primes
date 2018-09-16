# Primes

A simple command-line prime checker. It uses an academic approach
rather than a highly optimized one. I might try to optimize this
later.

It can confirm most prime numbers within max int64 in 50ms or less
(on my machine). It does not support Big Integers sadly.

## Usage

Simple prime check:

```bash
primes 130619442894397
```

Optionally provide your own k values (see the Math section)

```bash
primes 130619442894397 210
```

## The Math

A number of techniques are used:

### Trivial O(1) Checks

- Any number below 2 is not prime
- Even numbers (other than 2) are not prime. Bitwise logic makes this easy to check in O(1).
- All primes under 200 are hard-coded for O(1) lookup

### `ka + b`

We check if `n` is prime by looking for non-trivial factors.
A factor `f` is non-trivial if:

- `f > 1` as negative factors do not decide if a number is prime, 0 factors no numbers, and 1 factors all numbers.
- `f * f <= n` as any greater factors will be paired with smaller factors.

We construct candidate factors using the form `ka + b` for `b < k`, 
keeping `k` constant and manipulating `a` and `b` to construct 
candidates.

#### Prime Factors of `k`

First check the prime factors of `k`. As it is constant, this step
should be trivial. Simply check if any of these factors divide `n`.
This handles our base case of `a = 1` and `b = 0`.

#### Loop Through Candidates

As we increment `a`, we must check whether `ka + b` for `b < k`
divide `n`. However, we don't need to check every value of `b`.
If `b` is a prime factor of `k`, then we already checked `ka + b`, which is simply a multiple of `b`.

Thus, we can eliminate a substantial number of `b` values. With our
example of `k = 30` (prime factors 2, 3, and 5) we can eliminate

- Multiples of 2: `b = 0, 2, 4, 6, 8, 12, 14, 16, 18, 22, 24, 26, 28`
- Multiples of 3: `b = 3, 9, 15, 21, 27`
- Multiples of 5: `b = 5, 25`

This leaves us only:

`b = 1, 7, 11, 13, 17, 19, 23, 29`

Notice that this is just 1 joined with the set of primes up to
`k` which aren't prime factors of `k`. As we already precomputed the
prime numbers under 200, this list is trivial to construct for
any `k` we can reasonable expect to compute with an int64.

TL;DR: For each `a`, check if `ka + b` divides `n` for either `b = 1` or each prime `b` which doesn't divide `k`.

## Installation

Get the code:

```bash
go get github.com/todd-beckman/primes
```

Globally install (into `$GOPATH/go/bin`):

```bash
go install
```
