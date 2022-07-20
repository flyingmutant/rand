// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"flag"
	"fmt"
	"math"
	"math/bits"
	"pgregory.net/rand"
	"testing"
)

var (
	biasFlag = flag.Bool("bias", false, "run bias-detection tests")
)

func uint16nModulo(r *rand.Rand, n uint16) uint16 {
	return uint16(r.Uint32()) % n // biased
}

func uint16nFixedPoint(r *rand.Rand, n uint16, m int) uint16 {
	res, frac := bits.Mul32(uint32(n), r.Uint32()&(1<<(16+m)-1))
	return (uint16(res) << (16 - m)) | (uint16(frac>>16) >> m) // biased with probability 2^-m
}

func uint16nFixedPoint0(r *rand.Rand, n uint16) uint16  { return uint16nFixedPoint(r, n, 0) }
func uint16nFixedPoint2(r *rand.Rand, n uint16) uint16  { return uint16nFixedPoint(r, n, 2) }
func uint16nFixedPoint4(r *rand.Rand, n uint16) uint16  { return uint16nFixedPoint(r, n, 4) }
func uint16nFixedPoint8(r *rand.Rand, n uint16) uint16  { return uint16nFixedPoint(r, n, 8) }
func uint16nFixedPoint10(r *rand.Rand, n uint16) uint16 { return uint16nFixedPoint(r, n, 10) }
func uint16nFixedPoint12(r *rand.Rand, n uint16) uint16 { return uint16nFixedPoint(r, n, 12) }
func uint16nFixedPoint14(r *rand.Rand, n uint16) uint16 { return uint16nFixedPoint(r, n, 14) }
func uint16nFixedPoint16(r *rand.Rand, n uint16) uint16 { return uint16nFixedPoint(r, n, 16) }

func uint16nCanon(r *rand.Rand, n uint16) uint16 {
	res, frac := bits.Mul64(uint64(n), r.Uint64())
	hi, _ := bits.Mul64(uint64(n), r.Uint64())
	_, carry := bits.Add64(frac, hi, 0)
	return uint16(res + carry) // biased with probability 2^-64
}

func uint16nLemire(r *rand.Rand, n uint16) uint16 {
	v := uint16(r.Uint32())
	prod := uint32(v) * uint32(n)
	low := uint16(prod)
	if low < n {
		thresh := -n % n
		for low < thresh {
			v = uint16(r.Uint32())
			prod = uint32(v) * uint32(n)
			low = uint16(prod)
		}
	}
	return uint16(prod >> 16) // unbiased
}

func TestRand_Uint16nBias_Modulo(t *testing.T)       { testRandUint16nBias(t, uint16nModulo) }
func TestRand_Uint16nBias_FixedPoint0(t *testing.T)  { testRandUint16nBias(t, uint16nFixedPoint0) }
func TestRand_Uint16nBias_FixedPoint2(t *testing.T)  { testRandUint16nBias(t, uint16nFixedPoint2) }
func TestRand_Uint16nBias_FixedPoint4(t *testing.T)  { testRandUint16nBias(t, uint16nFixedPoint4) }
func TestRand_Uint16nBias_FixedPoint8(t *testing.T)  { testRandUint16nBias(t, uint16nFixedPoint8) }
func TestRand_Uint16nBias_FixedPoint10(t *testing.T) { testRandUint16nBias(t, uint16nFixedPoint10) }
func TestRand_Uint16nBias_FixedPoint12(t *testing.T) { testRandUint16nBias(t, uint16nFixedPoint12) }
func TestRand_Uint16nBias_FixedPoint14(t *testing.T) { testRandUint16nBias(t, uint16nFixedPoint14) }
func TestRand_Uint16nBias_FixedPoint16(t *testing.T) { testRandUint16nBias(t, uint16nFixedPoint16) }
func TestRand_Uint16nBias_Canon(t *testing.T)        { testRandUint16nBias(t, uint16nCanon) }
func TestRand_Uint16nBias_Lemire(t *testing.T)       { testRandUint16nBias(t, uint16nLemire) }

func testRandUint16nBias(t *testing.T, gen func(*rand.Rand, uint16) uint16) {
	t.Helper()

	if !*biasFlag {
		t.Skip("specify -bias flag to run bias tests")
	}

	const bound = math.MaxUint16 / 4 * 3
	for pow := 10; pow <= 50; pow++ {
		t.Run(fmt.Sprintf("%d/%dbit", bound, pow), func(t *testing.T) {
			r := rand.New()
			data := make([]uint64, bound)
			attempts := 1 << int64(pow)
			for i := 0; i < attempts; i++ {
				ix := gen(r, bound)
				data[ix]++
			}

			var chiSq float64
			expected := float64(attempts) / float64(bound)
			for _, n := range data {
				obs := float64(n)
				chiSq += (obs - expected) * (obs - expected)
			}
			chiSq /= expected

			df := float64(bound - 1)
			t.Logf("χ2 = %.1f, DoF = %v (%v attempts, delta = %.1f%%)", chiSq, df, attempts, (chiSq-df)/df*100)
		})
	}
}
