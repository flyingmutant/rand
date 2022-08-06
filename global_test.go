// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand_test

import (
	"math/bits"
	"pgregory.net/rand"
	"sync/atomic"
	"testing"
)

var (
	wyrandState uint64
)

func wyrand64() uint64 {
	s := atomic.AddUint64(&wyrandState, 0xa0761d6478bd642f)
	hi, lo := bits.Mul64(s, s^0xe7037ed1a0b428db)
	return hi ^ lo
}

func BenchmarkRand64(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s = rand.Rand64()
	}
	sinkUint64 = s
}

func BenchmarkWyRand64(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s = wyrand64()
	}
	sinkUint64 = s
}
