// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	exprand "golang.org/x/exp/rand"
	"hash/maphash"
	"log"
	"math"
	"math/bits"
	mathrand "math/rand"
	"os"
	"pgregory.net/rand"
)

const (
	chunkSizeBits  = 1 << 16
	chunkSizeBytes = chunkSizeBits / 8
	numChunks      = 1024
	bufSizeBits    = numChunks * chunkSizeBits
	bufSizeBytes   = bufSizeBits / 8
	bufSizeWords   = bufSizeBytes / 8
	maxInt52       = 1<<52 - 1
)

type randGen interface {
	Uint64() uint64
	Float64() float64
	NormFloat64() float64
	ExpFloat64() float64
}

type rand64 struct {
	rng randGen
}

func (r *rand64) raw() uint64 {
	return r.rng.Uint64()
}

func (r *rand64) fromF64() uint64 {
	return floatToUniform(r.rng.Float64(), r.rng.Float64())
}

func (r *rand64) fromNorm() uint64 {
	return floatToUniform(normalCDF(r.rng.NormFloat64()), normalCDF(r.rng.NormFloat64()))
}

func (r *rand64) fromExp() uint64 {
	return floatToUniform(expCDF(r.rng.ExpFloat64()), expCDF(r.rng.ExpFloat64()))
}

func floatToUniform(x float64, y float64) uint64 {
	return uint64(x*maxInt52)<<52 | uint64(y*maxInt52)
}

func normalCDF(x float64) float64 {
	return 0.5 * math.Erfc(-x/math.Sqrt2)
}

func expCDF(x float64) float64 {
	return -math.Expm1(-x)
}

func uint16nModulo(g func() uint64, n uint16) uint16 {
	return uint16(g()) % n // biased
}

func uint16nFixedPoint(g func() uint64, n uint16) uint16 {
	v := uint16(g())
	x := uint32(n) * uint32(v)
	return uint16(x >> 16) // biased
}

func uint16nLongFixedPoint(g func() uint64, n uint16) uint16 {
	res, _ := bits.Mul32(uint32(n), uint32(g()))
	return uint16(res) // biased with probability 2^-16
}

func uint16nLemire(g func() uint64, n uint16) uint16 {
	v := uint16(g())
	prod := uint32(v) * uint32(n)
	low := uint16(prod)
	if low < n {
		thresh := -n % n
		for low < thresh {
			v = uint16(g())
			prod = uint32(v) * uint32(n)
			low = uint16(prod)
		}
	}
	return uint16(prod >> 16) // unbiased
}

func shuffleBits(buf []byte, g func() uint64, b func(func() uint64, uint16) uint16) {
	for i := math.MaxUint16 - 1; i > 0; i-- {
		j := int(b(g, uint16(i+1)))
		bi := getBit(buf, i)
		bj := getBit(buf, j)
		setBit(buf, i, bj)
		setBit(buf, j, bi)
	}
}

func getBit(buf []byte, i int) bool {
	return buf[i/8]&(1<<(i%8)) > 0
}

func setBit(buf []byte, i int, b bool) {
	if b {
		buf[i/8] |= 1 << (i % 8)
	} else {
		buf[i/8] &= ^(1 << (i % 8))
	}
}

func run(gen string, shuffle string) error {
	var (
		randG = rand64{rand.New()}
		stdG  = rand64{mathrand.New(mathrand.NewSource(int64(new(maphash.Hash).Sum64())))}
		xG    = rand64{exprand.New(exprand.NewSource(new(maphash.Hash).Sum64()))}
		g     func() uint64
	)
	switch gen {
	case "rand":
		g = randG.raw
	case "rand-f64":
		g = randG.fromF64
	case "rand-norm":
		g = randG.fromNorm
	case "rand-exp":
		g = randG.fromExp
	case "std":
		g = stdG.raw
	case "std-f64":
		g = stdG.fromF64
	case "std-norm":
		g = stdG.fromNorm
	case "std-exp":
		g = stdG.fromExp
	case "x":
		g = xG.raw
	case "x-f64":
		g = xG.fromF64
	case "x-norm":
		g = xG.fromNorm
	case "x-exp":
		g = xG.fromExp
	default:
		return fmt.Errorf("unknown RNG: %q", gen)
	}

	buf := make([]byte, 8*bufSizeWords)
	switch shuffle {
	case "none":
		return output(buf, g, nil)
	case "mod":
		return output(buf, g, uint16nModulo)
	case "fp":
		return output(buf, g, uint16nFixedPoint)
	case "lfp":
		return output(buf, g, uint16nLongFixedPoint)
	case "lemire":
		return output(buf, g, uint16nLemire)
	default:
		return fmt.Errorf("unknown shuffle method: %q", shuffle)
	}
}

func output(buf []byte, g func() uint64, b func(func() uint64, uint16) uint16) error {
	for {
		if b == nil {
			for i := 0; i < bufSizeWords; i++ {
				binary.LittleEndian.PutUint64(buf[i*8:], g())
			}
		} else {
			for i := 0; i < numChunks; i++ {
				ch := buf[i*chunkSizeBytes : (i+1)*chunkSizeBytes]
				for j := 0; j < len(ch); j++ {
					if j < len(ch)/2 {
						ch[j] = 0xff
					} else {
						ch[j] = 0
					}
				}
				shuffleBits(ch, g, b)
			}
		}

		_, err := os.Stdout.Write(buf)
		if err != nil {
			return err
		}
	}
}

func main() {
	var (
		gen     = flag.String("gen", "sfc", "RNG to use (rand/rand-f64/rand-norm/rand-exp/std/std-f64/std-norm/std-exp/x/x-f64/x-norm/x-exp)")
		shuffle = flag.String("shuffle", "none", "shuffle algorithm to use (none/mod/fp/lfp/lemire)")
	)
	flag.Parse()

	err := run(*gen, *shuffle)
	if err != nil {
		log.Fatal(err.Error())
	}
}
