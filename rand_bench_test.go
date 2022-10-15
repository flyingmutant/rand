// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !benchx && !benchstd && !benchfast

package rand_test

import (
	"math"
	"pgregory.net/rand"
	"testing"
)

var (
	sinkRand *rand.Rand
)

func BenchmarkExpFloat64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s float64
		for pb.Next() {
			s = rand.ExpFloat64()
		}
		sinkFloat64 = s
	})
}

func BenchmarkFloat32(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s float32
		for pb.Next() {
			s = rand.Float32()
		}
		sinkFloat32 = s
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

func BenchmarkInt(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.Int()
		}
		sinkInt = s
	})
}

func BenchmarkInt31(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int32
		for pb.Next() {
			s = rand.Int31()
		}
		sinkInt32 = s
	})
}

func BenchmarkInt31n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int32
		for pb.Next() {
			s = rand.Int31n(small)
		}
		sinkInt32 = s
	})
}

func BenchmarkInt31n_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int32
		for pb.Next() {
			s = rand.Int31n(math.MaxInt32 - small)
		}
		sinkInt32 = s
	})
}

func BenchmarkInt63(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int64
		for pb.Next() {
			s = rand.Int63()
		}
		sinkInt64 = s
	})
}

func BenchmarkInt63n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int64
		for pb.Next() {
			s = rand.Int63n(small)
		}
		sinkInt64 = s
	})
}

func BenchmarkInt63n_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int64
		for pb.Next() {
			s = rand.Int63n(math.MaxInt64 - small)
		}
		sinkInt64 = s
	})
}

func BenchmarkIntn(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.Intn(small)
		}
		sinkInt = s
	})
}

func BenchmarkIntn_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.Intn(math.MaxInt - small)
		}
		sinkInt = s
	})
}

func BenchmarkNormFloat64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s float64
		for pb.Next() {
			s = rand.NormFloat64()
		}
		sinkFloat64 = s
	})
}

func BenchmarkPerm(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Perm(tiny)
		}
	})
}

func BenchmarkRead(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		p := make([]byte, 256)
		b.SetBytes(int64(len(p)))
		for pb.Next() {
			_, _ = rand.Read(p[:])
		}
	})
}

func BenchmarkShuffle(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		a := make([]int, tiny)
		for pb.Next() {
			rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		}
	})
}

func BenchmarkShuffleOverhead(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		a := make([]int, tiny)
		for pb.Next() {
			rand.Shuffle(len(a), func(i, j int) {})
		}
	})
}

func BenchmarkUint32(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint32
		b.SetBytes(4)
		for pb.Next() {
			s = rand.Uint32()
		}
		sinkUint32 = s
	})
}

func BenchmarkUint32n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint32
		for pb.Next() {
			s = rand.Uint32n(small)
		}
		sinkUint32 = s
	})
}

func BenchmarkUint32n_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint32
		for pb.Next() {
			s = rand.Uint32n(math.MaxUint32 - small)
		}
		sinkUint32 = s
	})
}

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

func BenchmarkUint64n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint64
		for pb.Next() {
			s = rand.Uint64n(small)
		}
		sinkUint64 = s
	})
}

func BenchmarkUint64n_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint64
		for pb.Next() {
			s = rand.Uint64n(math.MaxUint64 - small)
		}
		sinkUint64 = s
	})
}

func BenchmarkRand_New(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New(uint64(i))
	}
	sinkRand = s
}

func BenchmarkRand_New0(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New()
	}
	sinkRand = s
}

func BenchmarkRand_New3(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New(uint64(i), uint64(i), uint64(i))
	}
	sinkRand = s
}

func BenchmarkRand_NewInt(b *testing.B) {
	var s int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New().Int()
	}
	sinkInt = s
}

func BenchmarkRand_Get(b *testing.B) {
	var r rand.Rand
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = r.Get()
	}
	sinkRand = s
}

func BenchmarkRand_GetInt(b *testing.B) {
	var r rand.Rand
	var s int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = r.Get().Int()
	}
	sinkInt = s
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.ExpFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Float32(b *testing.B) {
	var s float32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Float32()
	}
	sinkFloat32 = s
}

func BenchmarkRand_Float64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Float64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Int(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int()
	}
	sinkInt = s
}

func BenchmarkRand_Int31(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31()
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31n(small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n_Big(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31n(math.MaxInt32 - small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int63(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63()
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63n(small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n_Big(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63n(math.MaxInt64 - small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Intn(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Intn(small)
	}
	sinkInt = s
}

func BenchmarkRand_Intn_Big(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Intn(math.MaxInt - small)
	}
	sinkInt = s
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.NormFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkRand_Read(b *testing.B) {
	r := rand.New(1)
	p := make([]byte, 256)
	b.SetBytes(int64(len(p)))
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkRand_Seed(b *testing.B) {
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		r.Seed(uint64(i))
	}
}

func BenchmarkRand_Shuffle(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkRand_Uint32(b *testing.B) {
	var s uint32
	r := rand.New(1)
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		s = r.Uint32()
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint32n(b *testing.B) {
	var s uint32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint32n(small)
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint32n_Big(b *testing.B) {
	var s uint32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint32n(math.MaxUint32 - small)
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint64(b *testing.B) {
	var s uint64
	r := rand.New(1)
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s = r.Uint64()
	}
	sinkUint64 = s
}

func BenchmarkRand_Uint64n(b *testing.B) {
	var s uint64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(small)
	}
	sinkUint64 = s
}

func BenchmarkRand_Uint64n_Big(b *testing.B) {
	var s uint64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(math.MaxUint64 - small)
	}
	sinkUint64 = s
}

func BenchmarkRand_MarshalBinary(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		_, _ = r.MarshalBinary()
	}
}

func BenchmarkRand_UnmarshalBinary(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	buf, _ := r.MarshalBinary()
	for i := 0; i < b.N; i++ {
		_ = r.UnmarshalBinary(buf)
	}
}
