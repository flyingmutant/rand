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
)

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
	var g func() uint64
	switch gen {
	case "sfc":
		g = rand.New().Uint64
	case "exp":
		g = exprand.New(exprand.NewSource(new(maphash.Hash).Sum64())).Uint64
	case "std":
		g = mathrand.New(mathrand.NewSource(int64(new(maphash.Hash).Sum64()))).Uint64
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
		gen     = flag.String("gen", "sfc", "RNG to use (sfc/exp/std)")
		shuffle = flag.String("shuffle", "none", "shuffle algorithm to use (none/mod/fp/lfp/lemire)")
	)
	flag.Parse()

	err := run(*gen, *shuffle)
	if err != nil {
		log.Fatal(err.Error())
	}
}
