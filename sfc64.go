// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rand

import "math/bits"

type sfc64 struct {
	a uint64
	b uint64
	c uint64
	w uint64
}

func (s *sfc64) init(a uint64, b uint64, c uint64, w uint64, n int) {
	s.a = a
	s.b = b
	s.c = c
	s.w = w
	for i := 0; i < n; i++ {
		s.next()
	}
}

func (s *sfc64) next() uint64 {
	out := s.a + s.b + s.w
	s.w++
	s.a = s.b ^ (s.b >> 11)
	s.b = s.c + (s.c << 3)
	s.c = bits.RotateLeft64(s.c, 24) + out
	return out
}
