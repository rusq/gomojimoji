// Package mojimoji is a port of mojimoji package to Go.
// Original: https://github.com/studio-ousia/mojimoji

package gomojimoji

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestZenToHan(t *testing.T) {
	type args struct {
		text string
		opt  []Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"kana", args{"アイウエオ", nil}, "ｱｲｳｴｵ"},
		{"ten", args{"ガギグゲゴ", nil}, "ｶﾞｷﾞｸﾞｹﾞｺﾞ"},
		{"maru", args{"パピプペポ", nil}, "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ"},
		{"digits", args{"０１２３", nil}, "0123"},
		{"ASCII", args{"ａｂｃＡＢＣ", nil}, "abcABC"},
		{"symbols", args{"＃？！￥", nil}, "#?!¥"},
		{"hiragana", args{"あいうえお", nil}, "あいうえお"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZenToHan(tt.args.text, tt.args.opt...); got != tt.want {
				t.Errorf("ZenToHan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHanToZen(t *testing.T) {
	type args struct {
		text string
		opt  []Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"kana", args{"ｱｲｳｴｵ", nil}, "アイウエオ"},
		{"ten", args{"ｶﾞｷﾞｸﾞｹﾞｺﾞ", nil}, "ガギグゲゴ"},
		{"maru", args{"ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ", nil}, "パピプペポ"},
		{"digits", args{"0123", nil}, "０１２３"},
		{"ascii", args{"abcABC", nil}, "ａｂｃＡＢＣ"},
		{"symbols", args{"#?!¥", nil}, "＃？！￥"},
		{"hiragana", args{"あいうえお", nil}, "あいうえお"},
		{"mixed", args{"ｱﾞｲﾞｳﾞｴﾞｵﾞ", nil}, "ア゛イ゛ヴエ゛オ゛"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HanToZen(tt.args.text, tt.args.opt...); got != tt.want {
				t.Errorf("HanToZen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleZenToHan() {
	fmt.Println(ZenToHan("ニュージーランド"))
	fmt.Println(ZenToHan("ニュージーランド Ａｕｃｋｌａｎｄ ０１２３", Kana(false), Digits(true)))
	//Output:
	// ﾆｭｰｼﾞｰﾗﾝﾄﾞ
	// ニュージーランド Auckland 0123

}

func ExampleHanToZen() {
	fmt.Println(HanToZen("ﾆｭｰｼﾞｰﾗﾝﾄﾞ"))
	fmt.Println(HanToZen("ﾆｭｰｼﾞｰﾗﾝﾄﾞ Auckland 6012", ASCII(true), Digits(false), Kana(false)))
	//Output:
	//ニュージーランド
	//ﾆｭｰｼﾞｰﾗﾝﾄﾞ　Ａｕｃｋｌａｎｄ　6012
}

func BenchmarkZenToHan(b *testing.B) {
	s := strings.Repeat("ＡＢＣＤＥＦＧ０１２３４５", 10)
	for n := 0; n < b.N; n++ {
		ZenToHan(s)
	}
}

func BenchmarkZenToHanConv(b *testing.B) {
	// this benchmark is similar to that for python mojimoji library.
	s := strings.Repeat("ＡＢＣＤＥＦＧ０１２３４５", 10)
	start := time.Now()
	for n := 0; n < 1000000; n++ {
		ZenToHan(s)
	}
	b.Log(time.Since(start))
}

func BenchmarkHanToZen(b *testing.B) {
	s := strings.Repeat("ABCDEFG012345", 10)
	for n := 0; n < b.N; n++ {
		HanToZen(s)
	}
}

func BenchmarkHanToZenConv(b *testing.B) {
	// this benchmark is similar to that for python mojimoji library.
	s := strings.Repeat("ABCDEFG012345", 10)
	start := time.Now()
	for n := 0; n < 1000000; n++ {
		HanToZen(s)
	}
	b.Log(time.Since(start))
}
