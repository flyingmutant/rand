// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"math"
	"math/bits"
	"pgregory.net/rand"
	"pgregory.net/rapid"
	"sync/atomic"
	"testing"
)

const (
	wyrandAdd = 0xa0761d6478bd642f
	wyrandXor = 0xe7037ed1a0b428db
)

func wyrand64(state *uint64) uint64 {
	s := *state + wyrandAdd
	*state = s
	hi, lo := bits.Mul64(s, s^wyrandXor)
	return hi ^ lo
}

func wyrand64Atomic(state *uint64) uint64 {
	s := atomic.AddUint64(state, wyrandAdd)
	hi, lo := bits.Mul64(s, s^wyrandXor)
	return hi ^ lo
}

func BenchmarkRand64(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s = rand.Uint64()
	}
	sinkUint64 = s
}

func BenchmarkWyRand64(b *testing.B) {
	var s uint64
	var state uint64
	for i := 0; i < b.N; i++ {
		s = wyrand64(&state)
	}
	sinkUint64 = s
}

func BenchmarkWyRand64Atomic(b *testing.B) {
	var s uint64
	var state uint64
	for i := 0; i < b.N; i++ {
		s = wyrand64Atomic(&state)
	}
	sinkUint64 = s
}

func TestFloat32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		f := rand.Float32()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestFloat64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		f := rand.Float64()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestInt31n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Int32Range(1, math.MaxInt32).Draw(t, "n").(int32)
		v := rand.Int31n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestInt63n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Int64Range(1, math.MaxInt64).Draw(t, "n").(int64)
		v := rand.Int63n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestIntn(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.IntRange(1, math.MaxInt).Draw(t, "n").(int)
		v := rand.Intn(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestUint32n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Uint32Range(1, math.MaxUint32).Draw(t, "n").(uint32)
		v := rand.Uint32n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestUint64n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Uint64Range(1, math.MaxUint64).Draw(t, "n").(uint64)
		v := rand.Uint64n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}
