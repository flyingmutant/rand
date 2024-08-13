// Copyright 2024 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.22 && benchstdpcg

package rand_test

import (
	"math"
	"math/rand/v2"
	"testing"
)

var (
	sinkRand *rand.Rand
)

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint64
		b.SetBytes(8)
		for pb.Next() {
			s = rand.Uint64()
		}
		sinkUint64 = s
	})
}

func BenchmarkFloat64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s float64
		for pb.Next() {
			s = rand.Float64()
		}
		sinkFloat64 = s
	})
}

func BenchmarkIntn(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.IntN(small)
		}
		sinkInt = s
	})
}

func BenchmarkIntn_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.IntN(math.MaxInt - small)
		}
		sinkInt = s
	})
}

func BenchmarkRand_New(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New(rand.NewPCG(uint64(i), uint64(i)))
	}
	sinkRand = s
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.ExpFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Float32(b *testing.B) {
	var s float32
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Float32()
	}
	sinkFloat32 = s
}

func BenchmarkRand_Float64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Float64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Int(b *testing.B) {
	var s int
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int()
	}
	sinkInt = s
}

func BenchmarkRand_Int31(b *testing.B) {
	var s int32
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int32()
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n(b *testing.B) {
	var s int32
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int32N(small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n_Big(b *testing.B) {
	var s int32
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int32N(math.MaxInt32 - small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int63(b *testing.B) {
	var s int64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int64()
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n(b *testing.B) {
	var s int64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int64N(small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n_Big(b *testing.B) {
	var s int64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.Int64N(math.MaxInt64 - small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Intn(b *testing.B) {
	var s int
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.IntN(small)
	}
	sinkInt = s
}

func BenchmarkRand_Intn_Big(b *testing.B) {
	var s int
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.IntN(math.MaxInt - small)
	}
	sinkInt = s
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		s = r.NormFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkRand_Seed(b *testing.B) {
	r := rand.NewPCG(1, 2)
	for i := 0; i < b.N; i++ {
		r.Seed(uint64(i), uint64(i))
	}
}

func BenchmarkRand_Shuffle(b *testing.B) {
	r := rand.New(rand.NewPCG(1, 2))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(rand.NewPCG(1, 2))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkRand_Uint32(b *testing.B) {
	var s uint32
	r := rand.New(rand.NewPCG(1, 2))
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		s = r.Uint32()
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint64(b *testing.B) {
	var s uint64
	r := rand.New(rand.NewPCG(1, 2))
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s = r.Uint64()
	}
	sinkUint64 = s
}
