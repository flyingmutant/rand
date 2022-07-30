# rand [![PkgGoDev][godev-img]][godev] [![CI][ci-img]][ci]

Fast, high quality alternative to `math/rand` and `golang.org/x/exp/rand`.

Compared to these packages, `pgregory.net/rand`:

- is API-compatible with all `*rand.Rand` methods,
- is significantly faster, while providing state-of-the-art generator quality,
- has simpler generator initialization:
  - `rand.New()` instead of `rand.New(rand.NewSource(time.Now().UnixNano()))`
  - `rand.New(1)` instead of `rand.New(rand.NewSource(1))`
- is deliberately not providing top-level functions like `Float64()` or `Int()`
  and the `Source` interface.

## Benchmarks

All benchmarks were run on `Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz` (frequency-locked),
`linux/amd64`.

Compared to [math/rand](https://pkg.go.dev/math/rand):

```
name                    old time/op    new time/op    delta
Rand_New-8                19.4µs ± 0%     0.1µs ± 4%   -99.54%  (p=0.000 n=8+9)
Rand_ExpFloat64-8         12.1ns ± 0%     8.6ns ± 0%   -28.92%  (p=0.000 n=10+8)
Rand_Float32-8            10.3ns ± 2%     3.4ns ± 1%   -66.60%  (p=0.000 n=10+9)
Rand_Float64-8            7.07ns ± 2%    4.08ns ± 0%   -42.24%  (p=0.000 n=10+9)
Rand_Int-8                6.40ns ± 2%    3.82ns ± 0%   -40.20%  (p=0.000 n=10+9)
Rand_Int31-8              6.27ns ± 2%    2.64ns ± 0%   -57.90%  (p=0.000 n=10+8)
Rand_Int31n-8             16.9ns ± 2%     4.0ns ± 0%   -76.14%  (p=0.000 n=10+8)
Rand_Int31n_Big-8         16.9ns ± 1%     4.0ns ± 0%   -76.07%  (p=0.000 n=10+9)
Rand_Int63-8              6.19ns ± 2%    3.82ns ± 0%   -38.31%  (p=0.000 n=10+9)
Rand_Int63n-8             42.5ns ± 2%     5.6ns ± 0%   -86.86%  (p=0.000 n=10+8)
Rand_Int63n_Big-8         42.5ns ± 1%     9.5ns ± 0%   -77.69%  (p=0.000 n=10+10)
Rand_Intn-8               19.4ns ± 2%     5.6ns ± 0%   -71.29%  (p=0.000 n=10+9)
Rand_Intn_Big-8           45.2ns ± 1%     9.5ns ± 0%   -79.02%  (p=0.000 n=10+9)
Rand_NormFloat64-8        13.3ns ± 0%     8.7ns ± 0%   -34.79%  (p=0.000 n=10+10)
Rand_Perm-8               1.28µs ± 1%    0.39µs ± 0%   -69.35%  (p=0.000 n=9+10)
Rand_Read-8                455ns ± 1%     134ns ± 0%   -70.52%  (p=0.000 n=10+10)
Rand_Seed-8               17.5µs ± 1%     0.0µs ± 0%   -99.74%  (p=0.000 n=9+9)
Rand_Shuffle-8             696ns ± 2%     409ns ± 0%   -41.32%  (p=0.000 n=10+9)
Rand_ShuffleOverhead-8     576ns ± 2%     303ns ± 0%   -47.30%  (p=0.000 n=10+9)
Rand_Uint32-8             6.15ns ± 2%    2.61ns ± 1%   -57.56%  (p=0.000 n=10+9)
Rand_Uint64-8             8.57ns ± 2%    3.87ns ± 0%   -54.88%  (p=0.000 n=10+9)

name                    old alloc/op   new alloc/op   delta
Rand_New-8                5.42kB ± 0%    0.05kB ± 0%   -99.12%  (p=0.000 n=10+10)
Rand_Perm-8                 416B ± 0%      416B ± 0%      ~     (all equal)

name                    old allocs/op  new allocs/op  delta
Rand_New-8                  2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
Rand_Perm-8                 1.00 ± 0%      1.00 ± 0%      ~     (all equal)

name                    old speed      new speed      delta
Rand_Read-8              563MB/s ± 1%  1908MB/s ± 0%  +239.15%  (p=0.000 n=10+10)
Rand_Uint32-8            650MB/s ± 2%  1531MB/s ± 1%  +135.57%  (p=0.000 n=10+9)
Rand_Uint64-8            934MB/s ± 2%  2069MB/s ± 0%  +121.60%  (p=0.000 n=10+9)
```

<details>
<summary>Compared to <a href="https://pkg.go.dev/golang.org/x/exp/rand">golang.org/x/exp/rand</a>:</summary>

```
name                    old time/op    new time/op    delta
Rand_New-8                95.2ns ± 1%    88.9ns ± 4%    -6.68%  (p=0.000 n=10+9)
Rand_ExpFloat64-8         12.3ns ± 0%     8.6ns ± 0%   -30.56%  (p=0.000 n=10+8)
Rand_Float32-8            13.5ns ± 1%     3.4ns ± 1%   -74.38%  (p=0.000 n=9+9)
Rand_Float64-8            11.1ns ± 5%     4.1ns ± 0%   -63.29%  (p=0.000 n=10+9)
Rand_Int-8                7.37ns ± 5%    3.82ns ± 0%   -48.12%  (p=0.000 n=10+9)
Rand_Int31-8              7.10ns ± 4%    2.64ns ± 0%   -62.82%  (p=0.000 n=10+8)
Rand_Int31n-8             26.1ns ± 0%     4.0ns ± 0%   -84.49%  (p=0.000 n=10+8)
Rand_Int31n_Big-8         26.0ns ± 1%     4.0ns ± 0%   -84.50%  (p=0.000 n=9+9)
Rand_Int63-8              6.92ns ± 1%    3.82ns ± 0%   -44.82%  (p=0.000 n=9+9)
Rand_Int63n-8             26.0ns ± 1%     5.6ns ± 0%   -78.55%  (p=0.000 n=9+8)
Rand_Int63n_Big-8         39.8ns ± 0%     9.5ns ± 0%   -76.15%  (p=0.000 n=8+10)
Rand_Intn-8               26.0ns ± 1%     5.6ns ± 0%   -78.55%  (p=0.000 n=9+9)
Rand_Intn_Big-8           39.9ns ± 1%     9.5ns ± 0%   -76.21%  (p=0.000 n=10+9)
Rand_NormFloat64-8        14.0ns ± 0%     8.7ns ± 0%   -38.03%  (p=0.000 n=9+10)
Rand_Perm-8               1.42µs ± 1%    0.39µs ± 0%   -72.23%  (p=0.000 n=9+10)
Rand_Read-8                467ns ± 0%     134ns ± 0%   -71.29%  (p=0.000 n=9+10)
Rand_Seed-8               6.28ns ± 6%   46.18ns ± 0%  +635.79%  (p=0.000 n=10+9)
Rand_Shuffle-8            1.40µs ± 1%    0.41µs ± 0%   -70.75%  (p=0.000 n=9+9)
Rand_ShuffleOverhead-8    1.28µs ± 0%    0.30µs ± 0%   -76.26%  (p=0.000 n=10+9)
Rand_Uint32-8             6.70ns ± 0%    2.61ns ± 1%   -61.02%  (p=0.000 n=10+9)
Rand_Uint64-8             6.71ns ± 0%    3.87ns ± 0%   -42.35%  (p=0.000 n=10+9)
Rand_Uint64n-8            25.0ns ± 1%     5.6ns ± 0%   -77.59%  (p=0.000 n=10+10)
Rand_Uint64n_Big-8        40.0ns ± 1%     9.2ns ± 1%   -76.99%  (p=0.000 n=9+10)
Rand_MarshalBinary-8      36.7ns ± 2%     6.1ns ± 0%   -83.27%  (p=0.000 n=10+10)
Rand_UnmarshalBinary-8    5.31ns ± 0%    6.14ns ± 0%   +15.63%  (p=0.000 n=9+10)

name                    old alloc/op   new alloc/op   delta
Rand_New-8                 48.0B ± 0%     48.0B ± 0%      ~     (all equal)
Rand_Perm-8                 416B ± 0%      416B ± 0%      ~     (all equal)
Rand_MarshalBinary-8       16.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Rand_UnmarshalBinary-8     0.00B          0.00B           ~     (all equal)

name                    old allocs/op  new allocs/op  delta
Rand_New-8                  2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
Rand_Perm-8                 1.00 ± 0%      1.00 ± 0%      ~     (all equal)
Rand_MarshalBinary-8        1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Rand_UnmarshalBinary-8      0.00           0.00           ~     (all equal)

name                    old speed      new speed      delta
Rand_Read-8              548MB/s ± 0%  1908MB/s ± 0%  +248.24%  (p=0.000 n=9+10)
Rand_Uint32-8            597MB/s ± 0%  1531MB/s ± 1%  +156.55%  (p=0.000 n=10+9)
Rand_Uint64-8           1.19GB/s ± 0%  2.07GB/s ± 0%   +73.45%  (p=0.000 n=10+9)
```
</details>

## FAQ

### Why did you write this?

`math/rand` is both slow and [not up to standards](https://gist.github.com/flyingmutant/ad5841f5e594aa8687fe47de34985e6a)
in terms of quality (but can not be changed because of Go 1 compatibility promise).
`golang.org/x/exp/rand` improves the quality, but does not improve the speed,
and it seems that there is no active development happening there.

### How does this thing work?

This package builds on 3 primitives: raw 64-bit generation using `sfc64`, floating-point
generation using floating-point multiplication, and integer generation in range using
32.64 or 64.128 fixed-point multiplication.

### Why is it fast?

The primitives selected are (as far as I am aware) about as fast as you can go
without sacrificing quality. On top of that, it is mainly making sure the compiler
is able to inline code, and a couple of micro-optimizations.

### Why no `Source`?

In Go (but not in C++ or Rust) it is a costly abstraction that provides no real value.
How often do you use a non-default `Source` with `math/rand`?

### Why no top-level functions?

Dislike for global mutable state. Also, without some kind of thread-local state they are
very slow (because global state needs to be mutex-protected). If you like the
convenience of top-level functions, `math/rand` is a fine choice.

### Why `sfc64`?

I like it. It has withstood the test of time, with no known flaws or weaknesses despite
a lot of effort and CPU-hours spent on finding them. Also, it provides guarantees about period
length and distance between generators seeded with different seeds. And it is fast.

### Why not...

#### ...`pcg`?

A bit slow. Otherwise, [`pcg64dxsm`](https://numpy.org/devdocs/reference/random/bit_generators/pcg64dxsm.html)
is probably a fine choice.

#### ...`xoshiro`/`xoroshiro`?

Quite a bit of controversy and people finding weaknesses in variants of this design.
Did you know that `xoshiro256**`, which author describes as an "all-purpose, rock-solid generator"
that "passes all tests we are aware of", fails them in seconds if you multiply the output by 57?

#### ...`splitmix`?

With 64-bit state and 64-bit output, `splitmix64` outputs every 64-bit number exactly once
over its 2^64 period — in other words, the probability of generating the same number is 0.
A [birthday test](https://www.pcg-random.org/posts/birthday-test.html) will quickly find
this problem.

#### ...`wyrand`?

An excellent generator if you are OK with slightly lower quality. Because its output function
(unlike `splitmix`) is not a bijection, some outputs are more likely to appear than others.
You can easily observe this non-uniformity with
a [birthday test](https://gist.github.com/flyingmutant/cb69e96872023f9f580868e746d1128a).

#### ...`romu`?

Very fast, but relatively new and untested. Also, no guarantees about the period length.

## Status

`pgregory.net/rand` is alpha software. API breakage and bugs should be expected before 1.0.

## License

`pgregory.net/rand` is licensed under the [Mozilla Public License Version 2.0](./LICENSE). 

[godev-img]: https://pkg.go.dev/badge/pgregory.net/rand
[godev]: https://pkg.go.dev/pgregory.net/rand
[ci-img]: https://github.com/flyingmutant/rand/workflows/CI/badge.svg
[ci]: https://github.com/flyingmutant/rand/actions
