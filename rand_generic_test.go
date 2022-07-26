// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.18

package rand_test

import (
	"pgregory.net/rand"
	"testing"
)

func BenchmarkShuffle(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		rand.Shuffle(r, a)
	}
}
