# (go) mojimoji

[![Go Reference](https://pkg.go.dev/badge/github.com/rusq/gomojimoji.svg)](https://pkg.go.dev/github.com/rusq/gomojimoji)

This is a port of the excellent [mojimoji][1] library written in Python
to Golang.

It provides two functions:
- **HanToZen** - half-width to full-width character conversion.
- **ZenToHan** - half-width to full-width character conversion.

Each of the functions allow the following options:
- **ASCII** - enable or disable ASCII translation.
- **Digits** - enable or disable Digits translation.
- **Kana** - enable or disable Kana translation.

All options are enabled by default, see examples on their usage.

Logic is implemented as of commit [aca2661][2].

## Examples

### HanToZen

```go
fmt.Println(HanToZen("ﾆｭｰｼﾞｰﾗﾝﾄﾞ"))
fmt.Println(HanToZen("ﾆｭｰｼﾞｰﾗﾝﾄﾞ Auckland 6012", ASCII(true), Digits(false), Kana(false)))

// Output:
// ニュージーランド
// ﾆｭｰｼﾞｰﾗﾝﾄﾞ　Ａｕｃｋｌａｎｄ　6012
```

### ZenToHan

```go
fmt.Println(ZenToHan("ニュージーランド"))
fmt.Println(ZenToHan("ニュージーランド Ａｕｃｋｌａｎｄ ０１２３", Kana(false), Digits(true)))

// Output:
// ﾆｭｰｼﾞｰﾗﾝﾄﾞ
// ニュージーランド Auckland 0123
```

## Benchmark

### Original library etc.
Original mojimoji, zenhan and unicodedata on my system, for comparison:
```python
In [4]: s = u'ＡＢＣＤＥＦＧ０１２３４５' * 10

In [5]: %time for n in range(1000000): mojimoji.zen_to_han(s)
CPU times: user 3.24 s, sys: 1.28 ms, total: 3.24 s
Wall time: 3.24 s

In [6]: %time for n in range(1000000): zenhan.z2h(s)
CPU times: user 26.2 s, sys: 16.3 ms, total: 26.2 s
Wall time: 26.2 s

In [7]: %time for n in range(1000000): unicodedata.normalize('NFKC', s)
CPU times: user 3.12 s, sys: 15.4 ms, total: 3.13 s
Wall time: 3.14 s
```

### This library
ZenToHan and HanToZen use different approaches:

- ZenToHan uses string.Builder, which is simpler to implement.
- HanToZen uses direct slice operations to allow for seeking when needed.

ZenToHan:
```
mojimoji (master)> go test -bench=BenchmarkZenToHanConv
goos: darwin
goarch: amd64
pkg: github.com/rusq/gomojimoji
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkZenToHanConv-16               1        2880823810 ns/op
--- BENCH: BenchmarkZenToHanConv-16
    mojimoji_test.go:98: 2.88079814s
PASS
ok      github.com/rusq/gomojimoji      2.977s
```

HanToZen:
```
mojimoji (master)> go test -bench=BenchmarkHanToZen    
goos: darwin
goarch: amd64
pkg: github.com/rusq/gomojimoji
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkHanToZenConv-16               1        2712209539 ns/op
--- BENCH: BenchmarkHanToZenConv-16
    mojimoji_test.go:107: 2.712166151s
PASS
ok      github.com/rusq/gomojimoji      2.804s
```

[1]: https://github.com/studio-ousia/mojimoji
[2]: https://github.com/studio-ousia/mojimoji/tree/aca26614f4a7a90a845f3a3c384c27d0a925efce
