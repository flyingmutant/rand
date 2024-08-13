// Copyright 2024 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.22 && !unsafe

package rand

import "math/rand/v2"

func rand64() uint64 {
	return rand.Uint64()
}
