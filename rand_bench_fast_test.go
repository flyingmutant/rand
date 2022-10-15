// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build benchfast

package rand_test

import (
	"github.com/valyala/fastrand"
	"testing"
)

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint64
		for pb.Next() {
			s = uint64(fastrand.Uint32())<<32 | uint64(fastrand.Uint32())
		}
		sinkUint64 = s
	})
}

func BenchmarkIntn(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint32
		for pb.Next() {
			s = fastrand.Uint32n(small)
		}
		sinkUint32 = s
	})
}
