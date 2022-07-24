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
Rand_New-8                18.5µs ± 0%     0.1µs ± 3%   -99.48%  (p=0.000 n=10+10)
Rand_ExpFloat64-8         12.4ns ± 6%     8.6ns ± 0%   -30.16%  (p=0.000 n=10+10)
Rand_Float32-8            10.7ns ± 5%     3.5ns ± 1%   -67.77%  (p=0.000 n=10+9)
Rand_Float64-8            6.93ns ± 3%    4.09ns ± 1%   -41.00%  (p=0.000 n=9+9)
Rand_Int-8                6.10ns ± 0%    3.82ns ± 0%   -37.38%  (p=0.000 n=8+10)
Rand_Int31-8              6.02ns ± 2%    2.80ns ± 1%   -53.47%  (p=0.000 n=9+10)
Rand_Int31n-8             16.2ns ± 1%     4.1ns ± 2%   -74.94%  (p=0.000 n=8+10)
Rand_Int31n_Big-8         16.1ns ± 1%     4.0ns ± 1%   -74.87%  (p=0.000 n=8+9)
Rand_Int63-8              6.02ns ± 2%    3.82ns ± 0%   -36.50%  (p=0.000 n=10+10)
Rand_Int63n-8             40.8ns ± 1%     5.9ns ± 0%   -85.62%  (p=0.000 n=8+10)
Rand_Int63n_Big-8         39.9ns ± 0%    14.4ns ± 0%   -63.99%  (p=0.000 n=9+9)
Rand_Intn-8               18.4ns ± 1%     5.9ns ± 0%   -68.03%  (p=0.000 n=9+9)
Rand_Intn_Big-8           43.0ns ± 1%    14.3ns ± 0%   -66.88%  (p=0.000 n=8+9)
Rand_NormFloat64-8        13.4ns ± 3%    10.1ns ± 1%   -25.20%  (p=0.000 n=10+10)
Rand_Perm-8               1.30µs ± 7%    0.42µs ± 2%   -67.38%  (p=0.000 n=9+10)
Rand_Read-8                446ns ± 8%     135ns ± 1%   -69.87%  (p=0.000 n=9+10)
Rand_Seed-8               17.7µs ± 3%     0.0µs ± 0%   -99.74%  (p=0.000 n=10+9)
Rand_Shuffle-8             673ns ± 5%     412ns ± 2%   -38.75%  (p=0.000 n=9+9)
Rand_ShuffleOverhead-8     554ns ± 3%     308ns ± 1%   -44.44%  (p=0.000 n=10+10)
Rand_Uint32-8             5.88ns ± 4%    2.82ns ± 1%   -52.02%  (p=0.000 n=9+9)
Rand_Uint64-8             8.68ns ± 2%    4.30ns ±10%   -50.44%  (p=0.000 n=10+10)

name                    old alloc/op   new alloc/op   delta
Rand_New-8                5.42kB ± 0%    0.05kB ± 0%   -99.12%  (p=0.000 n=10+10)
Rand_Perm-8                 416B ± 0%      416B ± 0%      ~     (all equal)

name                    old allocs/op  new allocs/op  delta
Rand_New-8                  2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
Rand_Perm-8                 1.00 ± 0%      1.00 ± 0%      ~     (all equal)

name                    old speed      new speed      delta
Rand_Read-8              568MB/s ±10%  1903MB/s ± 1%  +235.11%  (p=0.000 n=10+10)
Rand_Uint32-8            680MB/s ± 4%  1418MB/s ± 1%  +108.36%  (p=0.000 n=9+9)
Rand_Uint64-8            922MB/s ± 2%  1869MB/s ±10%  +102.73%  (p=0.000 n=10+10)
```

Compared to [golang.org/x/exp/rand](https://pkg.go.dev/golang.org/x/exp/rand):

```
name                    old time/op    new time/op    delta
Rand_New-8                 104ns ± 5%      97ns ± 3%    -6.83%  (p=0.000 n=10+10)
Rand_ExpFloat64-8         13.3ns ± 8%     8.6ns ± 0%   -34.92%  (p=0.000 n=10+10)
Rand_Float32-8            15.1ns ±10%     3.5ns ± 1%   -77.11%  (p=0.000 n=9+9)
Rand_Float64-8            11.4ns ± 1%     4.1ns ± 1%   -64.02%  (p=0.000 n=10+9)
Rand_Int-8                7.37ns ± 2%    3.82ns ± 0%   -48.14%  (p=0.000 n=9+10)
Rand_Int31-8              7.32ns ± 2%    2.80ns ± 1%   -61.73%  (p=0.000 n=10+10)
Rand_Int31n-8             27.0ns ± 1%     4.1ns ± 2%   -85.01%  (p=0.000 n=10+10)
Rand_Int31n_Big-8         27.1ns ± 1%     4.0ns ± 1%   -85.06%  (p=0.000 n=9+9)
Rand_Int63-8              7.29ns ± 4%    3.82ns ± 0%   -47.56%  (p=0.000 n=10+10)
Rand_Int63n-8             27.7ns ± 6%     5.9ns ± 0%   -78.85%  (p=0.000 n=10+10)
Rand_Int63n_Big-8         43.1ns ± 3%    14.4ns ± 0%   -66.67%  (p=0.000 n=10+9)
Rand_Intn-8               26.8ns ± 6%     5.9ns ± 0%   -78.12%  (p=0.000 n=10+9)
Rand_Intn_Big-8           40.1ns ± 0%    14.3ns ± 0%   -64.45%  (p=0.000 n=8+9)
Rand_NormFloat64-8        14.2ns ± 3%    10.1ns ± 1%   -29.18%  (p=0.000 n=10+10)
Rand_Perm-8               1.61µs ±10%    0.42µs ± 2%   -73.66%  (p=0.000 n=9+10)
Rand_Read-8                473ns ± 2%     135ns ± 1%   -71.56%  (p=0.000 n=10+10)
Rand_Seed-8               5.85ns ± 2%   46.27ns ± 0%  +691.44%  (p=0.000 n=8+9)
Rand_Shuffle-8            1.49µs ± 1%    0.41µs ± 2%   -72.33%  (p=0.000 n=10+9)
Rand_ShuffleOverhead-8    1.33µs ± 2%    0.31µs ± 1%   -76.92%  (p=0.000 n=10+10)
Rand_Uint32-8             6.97ns ± 1%    2.82ns ± 1%   -59.54%  (p=0.000 n=9+9)
Rand_Uint64-8             6.99ns ± 1%    4.30ns ±10%   -38.44%  (p=0.000 n=9+10)
Rand_Uint64n-8            27.5ns ± 8%     6.2ns ±10%   -77.42%  (p=0.000 n=10+10)
Rand_Uint64n_Big-8        41.4ns ± 3%     9.8ns ± 1%   -76.36%  (p=0.000 n=10+9)
Rand_MarshalBinary-8      40.6ns ± 9%     6.3ns ± 4%   -84.59%  (p=0.000 n=10+10)
Rand_UnmarshalBinary-8    5.40ns ± 5%    6.15ns ± 0%   +13.88%  (p=0.000 n=10+10)

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
Rand_Read-8              541MB/s ± 2%  1903MB/s ± 1%  +251.62%  (p=0.000 n=10+10)
Rand_Uint32-8            574MB/s ± 1%  1418MB/s ± 1%  +147.16%  (p=0.000 n=9+9)
Rand_Uint64-8           1.15GB/s ± 1%  1.87GB/s ±10%   +63.23%  (p=0.000 n=9+10)
```

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
