// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand

import (
	"math"
	"math/bits"
)

// Uint64 returns a uniformly distributed pseudo-random 64-bit value as an uint64.
//
// It is safe to call Uint64 concurrently from multiple goroutines, and its performance
// does not degrade when the parallelism increases. However, non-concurrent use of
// multiple instances of [Rand.Uint64] should be generally preferred over the concurrent use
// of Uint64, as [Rand.Uint64] is faster, and it generates higher quality pseudo-random numbers.
func Uint64() uint64 {
	return rand64()
}

// Float64 returns, as a float64, a uniformly distributed pseudo-random number in the half-open interval [0.0, 1.0).
//
// It is safe to call Float64 concurrently from multiple goroutines, and its performance
// does not degrade when the parallelism increases. However, non-concurrent use of
// multiple instances of [Rand.Float64] should be generally preferred over the concurrent use
// of Float64, as [Rand.Float64] is faster, and it generates higher quality pseudo-random numbers.
func Float64() float64 {
	return f64()
}

// Intn returns, as an int, a uniformly distributed non-negative pseudo-random number
// in the half-open interval [0, n). It panics if n <= 0.
//
// It is safe to call Intn concurrently from multiple goroutines, and its performance
// does not degrade when the parallelism increases. However, non-concurrent use of
// multiple instances of [Rand.Intn] should be generally preferred over the concurrent use
// of Intn, as [Rand.Intn] is faster, and it generates higher quality pseudo-random numbers.
func Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if math.MaxInt == math.MaxInt32 {
		return int(u32n(uint32(n)))
	} else {
		return int(u64n(uint64(n)))
	}
}

// same algorithm as Rand.Float64
func f64() float64 {
	return float64(rand64()&int53Mask) * f53Mul
}

// same algorithm as Rand.Uint32n
func u32n(n uint32) uint32 {
	res, _ := bits.Mul64(uint64(n), rand64())
	return uint32(res)
}

// same algorithm as Rand.Uint64n
func u64n(n uint64) uint64 {
	res, frac := bits.Mul64(n, rand64())
	if n <= math.MaxUint32 {
		return res
	}
	hi, _ := bits.Mul64(n, rand64())
	_, carry := bits.Add64(frac, hi, 0)
	return res + carry
}
