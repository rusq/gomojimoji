// Package mojimoji is a port of mojimoji package to Go.
// Original: https://github.com/studio-ousia/mojimoji
package gomojimoji

import (
	"strings"
	"unicode/utf8"
)

type options struct {
	ascii bool
	digit bool
	kana  bool
}

func (oo *options) apply(opt ...Option) *options {
	for _, fn := range opt {
		fn(oo)
	}
	return oo
}

// ASCII enables or disables conversion of ASCII runes (A-Za-z).
func ASCII(enable bool) Option {
	return func(o *options) {
		o.ascii = enable
	}
}

// Digits enables or disables conversion of digit runes (0-9).
func Digits(enable bool) Option {
	return func(o *options) {
		o.digit = enable
	}
}

// Kana enables or disables conversion of Kana runes.
func Kana(enable bool) Option {
	return func(o *options) {
		o.kana = enable
	}
}

// Option is the option function signature.
type Option func(*options)

// ZenToHan converts text to half-width runes.  By default all runes are
// converted, optionally caller can switch off rune-set by passing [Option].
func ZenToHan(text string, opt ...Option) string {
	var options = options{
		ascii: true,
		digit: true,
		kana:  true,
	}
	options.apply(opt...)

	var buf strings.Builder
	for _, c := range text {
		if r, ok := tabASCIIzh[c]; options.ascii && ok {
			buf.WriteRune(r)
		} else if r, ok := tabKANAzh[c]; options.kana && ok {
			buf.WriteRune(r)
		} else if r, ok := tabDIGITzh[c]; options.digit && ok {
			buf.WriteRune(r)
		} else if r, ok := tabKANATENzh[c]; options.kana && ok {
			buf.WriteRune(r)
			buf.WriteRune('ﾞ')
		} else if r, ok := tabKANAMARUzh[c]; options.kana && ok {
			buf.WriteRune(r)
			buf.WriteRune('ﾟ')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

// HanToZen converts text to full-width runes.  By default all runes are
// converted, optionally caller can switch off rune-set by passing [Option].
func HanToZen(text string, opt ...Option) string {
	var options = options{
		ascii: true,
		digit: true,
		kana:  true,
	}
	options.apply(opt...)

	var buf = make([]rune, utf8.RuneCountInString(text))
	var (
		prev rune
		pos  int
	)
	for _, c := range text {
		if r, ok := tabASCIIhz[c]; options.ascii && ok {
			buf[pos] = r
		} else if r, ok := tabDigitHz[c]; options.digit && ok {
			buf[pos] = r
		} else if r, ok := tabKanaTenHz[prev]; options.kana && c == 'ﾞ' && ok {
			pos--
			buf[pos] = r
		} else if r, ok := tabKanaMaruHz[prev]; options.kana && c == 'ﾟ' && ok {
			pos--
			buf[pos] = r
		} else if r, ok := tabKanaHz[c]; options.kana && ok {
			buf[pos] = r
		} else {
			buf[pos] = c
		}
		prev = c
		pos++
	}
	return string(buf[:pos])
}
