// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

#include "vendor/nanobench.h"
#include <cstdio>
#include <cstdint>

struct sfc64 {
    uint64_t a;
    uint64_t b;
    uint64_t c;
    uint64_t w;
};

uint64_t next(sfc64& s) {
    uint64_t out = s.a + s.b + s.w;
    s.w++;
    s.a = s.b ^ (s.b>>11);
    s.b = s.c + (s.c<<3);
    s.c = ((s.c << 24) | (s.c >> (64-24))) + out;
    return out;
}

uint32_t bound_mod(uint32_t n, uint64_t v) {
    return uint32_t(v) % n;
}

uint32_t bound_fp_32x32(uint32_t n, uint64_t v) {
    uint64_t r = uint64_t(n) * uint64_t(uint32_t(v));
    return r >> 32;
}

uint32_t bound_fp_32x64(uint32_t n, uint64_t v) {
    __uint128_t r = __uint128_t(n) * __uint128_t(v);
    return uint32_t(r >> 64);
}

uint32_t nearlydivisionless(uint32_t n, sfc64& s) {
    uint32_t x = uint32_t(next(s));
    uint64_t m = uint64_t(x) * uint64_t(n);
    uint32_t l = uint32_t(m);
    if (l < n) {
        uint32_t t = -n % n;
        while (l < t) {
            x = uint32_t(next(s));
            m = uint64_t(x) * uint64_t(n);
            l = uint32_t(m);
        }
    }
    return m >> 32;
}

int main() {
    ankerl::nanobench::Rng rng;
    uint32_t bound = rng();

    {
        uint64_t val = rng();

        ankerl::nanobench::Bench b;
        b.title("modulo reduction").unit("uint32_t").epochs(239).relative(true);

        b.run("modulo (biased)", [&]() {
            uint32_t x = bound_mod(bound, val);
            b.doNotOptimizeAway(x);
        });
        b.run("32x32 fixed point (biased)", [&]() {
            uint32_t x = bound_fp_32x32(bound, val);
            b.doNotOptimizeAway(x);
        });
        b.run("32x64 fixed point (unbiased*)", [&]() {
            uint32_t x = bound_fp_32x64(bound, val);
            b.doNotOptimizeAway(x);
        });
    }
    {
        sfc64 s{rng(), rng(), rng(), 1};

        ankerl::nanobench::Bench b;
        b.title("random number generation in range (sfc64)").unit("uint32_t").epochs(239).relative(true);

        b.run("raw uint32_t", [&]() {
            uint32_t x = next(s);
            b.doNotOptimizeAway(x);
        });
        b.run("modulo (biased)", [&]() {
            uint32_t x = bound_mod(bound, next(s));
            b.doNotOptimizeAway(x);
        });
        b.run("32x32 fixed point (biased)", [&]() {
            uint32_t x = bound_fp_32x32(bound, next(s));
            b.doNotOptimizeAway(x);
        });
        b.run("32x64 fixed point (unbiased*)", [&]() {
            uint32_t x = bound_fp_32x64(bound, next(s));
            b.doNotOptimizeAway(x);
        });
        b.run("Lemire's \"Nearly Divisionless\" (unbiased)", [&]() {
            uint32_t x = nearlydivisionless(bound, s);
            b.doNotOptimizeAway(x);
        });
    }
}
