// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand

import (
	"hash/maphash"
	"math/bits"
)

type sfc64 struct {
	a uint64
	b uint64
	c uint64
	w uint64
}

func (s *sfc64) init(a uint64, b uint64, c uint64) {
	s.a = a
	s.b = b
	s.c = c
	s.w = 1
	for i := 0; i < 12; i++ {
		s.next()
	}
}

//go:noinline
func (s *sfc64) init0() { // noinline makes sure New can be inlined, helping with escape analysis
	s.a = new(maphash.Hash).Sum64()
	s.b = new(maphash.Hash).Sum64()
	s.c = new(maphash.Hash).Sum64()
	s.w = 1
}

//go:noinline
func (s *sfc64) init1(u uint64) { // noinline makes sure NewSeeded can be inlined, helping with escape analysis
	s.a = u
	s.b = u
	s.c = u
	s.w = 1
	for i := 0; i < 12; i++ {
		s.next()
	}
}

func (s *sfc64) next() (out uint64) { // named return value lowers inlining cost a bit
	out = s.a + s.b + s.w
	s.w++
	s.a = s.b ^ (s.b >> 11)
	s.b = s.c + (s.c << 3)
	s.c = bits.RotateLeft64(s.c, 24) + out
	return
}
