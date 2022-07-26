// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.18

package rand

import "math"

// Shuffle pseudo-randomizes the order of the elements of s.
func Shuffle[S ~[]E, E any](r *Rand, s S) {
	i := len(s) - 1
	for ; i > math.MaxInt32-1; i-- {
		j := int(r.Uint64n(uint64(i) + 1))
		s[i], s[j] = s[j], s[i]
	}
	for ; i > 0; i-- {
		j := int(r.Uint32n(uint32(i) + 1))
		s[i], s[j] = s[j], s[i]
	}
}
