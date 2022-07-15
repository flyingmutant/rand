// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkMathRand_New(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.New(rand.NewSource(int64(i)))
	}
}

func BenchmarkMathRand_ExpFloat64(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.ExpFloat64()
	}
}

func BenchmarkMathRand_Float32(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Float32()
	}
}

func BenchmarkMathRand_Float64(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Float64()
	}
}

func BenchmarkMathRand_Int(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int()
	}
}

func BenchmarkMathRand_Int31(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int31()
	}
}

func BenchmarkMathRand_Int31n(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int31n(small)
	}
}

func BenchmarkMathRand_Int63(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}

func BenchmarkMathRand_Int63n(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int63n(small)
	}
}

func BenchmarkMathRand_Intn(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Intn(small)
	}
}

func BenchmarkMathRand_NormFloat64(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.NormFloat64()
	}
}

func BenchmarkMathRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkMathRand_Read(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	p := make([]byte, 256)
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkMathRand_Seed(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Seed(int64(i))
	}
}

func BenchmarkMathRand_Shuffle(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkMathRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkMathRand_Uint32(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Uint32()
	}
}

func BenchmarkMathRand_Uint32n(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int31n(small) // no Uint32n
	}
}

func BenchmarkMathRand_Uint32n_Big(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int31n(math.MaxInt32 - small) // no Uint32n
	}
}

func BenchmarkMathRand_Uint64(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Uint64()
	}
}

func BenchmarkMathRand_Uint64n(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int63n(small) // no Uint64n
	}
}

func BenchmarkMathRand_Uint64n_Big(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Int63n(math.MaxInt64 - small) // no Uint64
	}
}
