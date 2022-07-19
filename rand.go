// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand

import (
	"encoding/binary"
	"io"
	"math"
	"math/bits"
)

const (
	int24Mask = 1<<24 - 1
	int31Mask = 1<<31 - 1
	int53Mask = 1<<53 - 1
	int63Mask = 1<<63 - 1

	randSizeof = 8*4 + 8 + 1
)

// Rand is a pseudo-random number generator based on the SFC64 algorithm by Chris Doty-Humphrey.
//
// SFC64 has 256 bits of state, average period of ~2^255 and minimum period of at least 2^64.
// Generators returned by New or NewSeeded (with distinct seeds) are guaranteed
// to not run into each other for at least 2^64 iterations.
type Rand struct {
	sfc64
	val uint64
	pos int8
}

// New returns a generator initialized to a non-deterministic state.
func New() *Rand {
	var r Rand
	r.init0()
	return &r
}

// NewSeeded returns a generator seeded with the given value.
func NewSeeded(seed uint64) *Rand {
	var r Rand
	r.init1(seed)
	return &r
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (r *Rand) Seed(seed uint64) {
	r.init1(seed)
	r.pos = 0
	r.val = 0
}

// MarshalBinary returns the binary representation of the current state of the generator.
func (r *Rand) MarshalBinary() ([]byte, error) {
	var data [randSizeof]byte
	binary.LittleEndian.PutUint64(data[0:], r.a)
	binary.LittleEndian.PutUint64(data[8:], r.b)
	binary.LittleEndian.PutUint64(data[16:], r.c)
	binary.LittleEndian.PutUint64(data[24:], r.w)
	binary.LittleEndian.PutUint64(data[32:], r.val)
	data[40] = byte(r.pos)
	return data[:], nil
}

// UnmarshalBinary sets the state of the generator to the state represented in data.
func (r *Rand) UnmarshalBinary(data []byte) error {
	if len(data) < randSizeof {
		return io.ErrUnexpectedEOF
	}
	r.a = binary.LittleEndian.Uint64(data[0:])
	r.b = binary.LittleEndian.Uint64(data[8:])
	r.c = binary.LittleEndian.Uint64(data[16:])
	r.w = binary.LittleEndian.Uint64(data[24:])
	r.val = binary.LittleEndian.Uint64(data[32:])
	r.pos = int8(data[40])
	return nil
}

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
func (r *Rand) Float32() float32 {
	return float32(r.uint32_()&int24Mask) * 0x1.0p-24
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func (r *Rand) Float64() float64 {
	return float64(r.next()&int53Mask) * 0x1.0p-53
}

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	if math.MaxInt == math.MaxInt32 {
		return int(r.uint32_() & int31Mask)
	} else {
		return int(r.next() & int63Mask)
	}
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 {
	return int32(r.uint32_() & int31Mask)
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	return int32(r.Uint32n(uint32(n)))
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 {
	return int64(r.next() & int63Mask)
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	return int64(r.Uint64n(uint64(n)))
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if math.MaxInt == math.MaxInt32 {
		return int(r.Uint32n(uint32(n)))
	} else {
		return int(r.Uint64n(uint64(n)))
	}
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers in the half-open interval [0,n).
func (r *Rand) Perm(n int) []int {
	p := make([]int, n)
	for i := 1; i < len(p); i++ {
		j := r.Uint64n(uint64(i) + 1)
		p[i] = p[j]
		p[j] = i
	}
	return p
}

// Read generates len(p) random bytes and writes them into p. It always returns len(p) and a nil error.
func (r *Rand) Read(p []byte) (n int, err error) {
	val, pos := r.val, r.pos
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val, pos = r.next(), 8
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}
	r.val, r.pos = val, pos
	return
}

// Shuffle pseudo-randomizes the order of elements. n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}
	i := n - 1
	for ; i > math.MaxUint32-1; i-- {
		j := int(r.Uint64n(uint64(i) + 1))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Uint32n(uint32(i) + 1))
		swap(i, j)
	}
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 {
	return uint32(r.uint32_())
}

// uint32_ has a bit lower inlining cost because of uint64 return value
func (r *Rand) uint32_() uint64 {
	// unnatural code to fit into inlining budget of 80
	if r.pos < 4 {
		r.val, r.pos = r.next(), 4
		return r.val >> 32
	} else {
		r.pos = 0
		return r.val
	}
}

// Uint32n returns, as a uint32, a pseudo-random number in [0,n). Uint32n(0) returns 0.
func (r *Rand) Uint32n(n uint32) uint32 {
	// much faster 32-bit version of Uint64n(); result is unbiased with probability 1 - 2^-32.
	// detecting possible bias would require about 2^64 samples, which we consider acceptable
	// since it matches 2^64 guarantees about period length and distance between different seeds
	res, _ := bits.Mul64(uint64(n), r.next())
	return uint32(res)
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64 {
	return r.next()
}

// Uint64n returns, as a uint64, a pseudo-random number in [0,n). Uint64n(0) returns 0.
func (r *Rand) Uint64n(n uint64) uint64 {
	// "An optimal algorithm for bounded random integers" by Stephen Canon, https://github.com/apple/swift/pull/39143
	// making second multiplication unconditional makes the function inlineable, but slows things down for small n
	res, frac := bits.Mul64(n, r.next())
	if frac <= -n {
		return res
	}
	hi, _ := bits.Mul64(n, r.next())
	_, carry := bits.Add64(frac, hi, 0)
	return res + carry
}
