// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.18

package rand_test

import (
	"bytes"
	"pgregory.net/rand"
	"pgregory.net/rapid"
	"testing"
)

func BenchmarkShuffleSlice(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		rand.ShuffleSlice(r, a)
	}
}

func TestShuffleSlice(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.IntRange(0, small).Draw(t, "n").(int)
		buf1 := make([]byte, n)
		_, _ = r.Read(buf1)
		buf2 := append([]byte(nil), buf1...)
		r.Seed(s)
		r.Shuffle(n, func(i, j int) {
			buf1[i], buf1[j] = buf1[j], buf1[i]
		})
		r.Seed(s)
		rand.ShuffleSlice(r, buf2)
		if !bytes.Equal(buf1, buf2) {
			t.Fatalf("shuffle results differ: %q vs %q", buf1, buf2)
		}
	})
}
