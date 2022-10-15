// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"bytes"
	"math"
	"math/bits"
	"pgregory.net/rand"
	"pgregory.net/rapid"
	"testing"
)

func TestRand_Read(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		const N = 32
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		buf := make([]byte, N)
		_, _ = r.Read(buf)
		r.Seed(s)
		buf2 := make([]byte, N)
		for n := 0; n < N; {
			c := rapid.IntRange(0, N-n).Draw(t, "c").(int)
			_, _ = r.Read(buf2[n : n+c])
			n += c
		}
		if !bytes.Equal(buf, buf2) {
			t.Fatalf("got %q instead of %q when reading in chunks", buf2, buf)
		}
	})
}

func TestRand_Float32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		f := r.Float32()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Float64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		f := r.Float64()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Int31n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
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
		r := rand.New(s)
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
		r := rand.New(s)
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
		r := rand.New(s)
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
		r := rand.New(s)
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
		r1 := rand.New(s)
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

func TestRand_Uint32nOpt(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Uint32().Draw(t, "n").(uint32)
		v := rapid.Uint64().Draw(t, "v").(uint64)

		res, frac := bits.Mul32(n, uint32(v>>32))
		hi, _ := bits.Mul32(n, uint32(v))
		_, carry := bits.Add32(frac, hi, 0)
		res += carry

		res2, _ := bits.Mul64(uint64(n), v)

		if uint32(res2) != res {
			t.Fatalf("got %v instead of %v", res2, res)
		}
	})
}
