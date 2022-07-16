// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"bytes"
	"math"
	"pgregory.net/rand"
	"pgregory.net/rapid"
	"testing"
)

const (
	tiny  = 100
	small = 1000
)

func BenchmarkRand_New(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.New()
	}
}

func BenchmarkRand_NewSeeded(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.NewSeeded(1)
	}
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.ExpFloat64()
	}
}

func BenchmarkRand_Float32(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Float32()
	}
}

func BenchmarkRand_Float64(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Float64()
	}
}

func BenchmarkRand_Int(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Int()
	}
}

func BenchmarkRand_Int31(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Int31()
	}
}

func BenchmarkRand_Int31n(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Int31n(small)
	}
}

func BenchmarkRand_Int63(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}

func BenchmarkRand_Int63n(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Int63n(small)
	}
}

func BenchmarkRand_Intn(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Intn(small)
	}
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.NormFloat64()
	}
}

func BenchmarkRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkRand_Read(b *testing.B) {
	r := rand.NewSeeded(1)
	p := make([]byte, 256)
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkRand_Seed(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Seed(uint64(i))
	}
}

func BenchmarkRand_Shuffle(b *testing.B) {
	r := rand.NewSeeded(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkRand_ShuffleOverhead(b *testing.B) {
	r := rand.NewSeeded(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkRand_Uint32(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint32()
	}
}

func BenchmarkRand_Uint32n(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint32n(small)
	}
}

func BenchmarkRand_Uint32n_Big(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint32n(math.MaxUint32 - small)
	}
}

func BenchmarkRand_Uint64(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint64()
	}
}

func BenchmarkRand_Uint64n(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint64n(small)
	}
}

func BenchmarkRand_Uint64n_Big(b *testing.B) {
	r := rand.NewSeeded(1)
	for i := 0; i < b.N; i++ {
		r.Uint64n(math.MaxUint64 - small)
	}
}

func TestRand_Float32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		f := r.Float32()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Float64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		f := r.Float64()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Int31n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		n := rapid.Int32Range(1, math.MaxInt32).Draw(t, "n").(int32)
		v := r.Int31n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Int63n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		n := rapid.Int64Range(1, math.MaxInt64).Draw(t, "n").(int64)
		v := r.Int63n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Intn(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		n := rapid.IntRange(1, math.MaxInt).Draw(t, "n").(int)
		v := r.Intn(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Uint32n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		n := rapid.Uint32Range(1, math.MaxUint32).Draw(t, "n").(uint32)
		v := r.Uint32n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Uint64n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.NewSeeded(s)
		n := rapid.Uint64Range(1, math.MaxUint64).Draw(t, "n").(uint64)
		v := r.Uint64n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_MarshalBinary_Roundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r1 := rand.NewSeeded(s)
		data1, err := r1.MarshalBinary()
		if err != nil {
			t.Fatalf("got unexpected marshal error: %v", err)
		}
		var r2 rand.Rand
		err = r2.UnmarshalBinary(data1)
		if err != nil {
			t.Fatalf("got unexpected unmarshal error: %v", err)
		}
		data2, err := r2.MarshalBinary()
		if err != nil {
			t.Fatalf("got unexpected marshal error: %v", err)
		}
		if !bytes.Equal(data1, data2) {
			t.Fatalf("data %q / %q after marshal/unmarshal", data1, data2)
		}
	})
}
