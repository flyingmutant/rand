// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"golang.org/x/exp/rand"
	"math"
	"testing"
)

func BenchmarkExpRand_New(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.New(rand.NewSource(uint64(i)))
	}
}

func BenchmarkExpRand_ExpFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.ExpFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkExpRand_Float32(b *testing.B) {
	var s float32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Float32()
	}
	sinkFloat32 = s
}

func BenchmarkExpRand_Float64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Float64()
	}
	sinkFloat64 = s
}

func BenchmarkExpRand_Int(b *testing.B) {
	var s int
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int()
	}
	sinkInt = s
}

func BenchmarkExpRand_Int31(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31()
	}
	sinkInt32 = s
}

func BenchmarkExpRand_Int31n(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31n(small)
	}
	sinkInt32 = s
}

func BenchmarkExpRand_Int63(b *testing.B) {
	var s int64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int63()
	}
	sinkInt64 = s
}

func BenchmarkExpRand_Int63n(b *testing.B) {
	var s int64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int63n(small)
	}
	sinkInt64 = s
}

func BenchmarkExpRand_Intn(b *testing.B) {
	var s int
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Intn(small)
	}
	sinkInt = s
}

func BenchmarkExpRand_NormFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.NormFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkExpRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkExpRand_Read(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	p := make([]byte, 256)
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkExpRand_Seed(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Seed(uint64(i))
	}
}

func BenchmarkExpRand_Shuffle(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkExpRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkExpRand_Uint32(b *testing.B) {
	var s uint32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Uint32()
	}
	sinkUint32 = s
}

func BenchmarkExpRand_Uint32n(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31n(small) // no Uint32n
	}
	sinkInt32 = s
}

func BenchmarkExpRand_Uint32n_Big(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31n(math.MaxInt32 - small) // no Uint32n
	}
	sinkInt32 = s
}

func BenchmarkExpRand_Uint64(b *testing.B) {
	var s uint64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Uint64()
	}
	sinkUint64 = s
}

func BenchmarkExpRand_Uint64n(b *testing.B) {
	var s uint64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(small)
	}
	sinkUint64 = s
}

func BenchmarkExpRand_Uint64n_Big(b *testing.B) {
	var s uint64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(math.MaxUint64 - small)
	}
	sinkUint64 = s
}
