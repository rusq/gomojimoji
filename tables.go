package gomojimoji

// Character tables are taken from the original mojimoji python library.
var (
	ASCII_ZENKAKU_CHARS    = []rune{'ａ', 'ｂ', 'ｃ', 'ｄ', 'ｅ', 'ｆ', 'ｇ', 'ｈ', 'ｉ', 'ｊ', 'ｋ', 'ｌ', 'ｍ', 'ｎ', 'ｏ', 'ｐ', 'ｑ', 'ｒ', 'ｓ', 'ｔ', 'ｕ', 'ｖ', 'ｗ', 'ｘ', 'ｙ', 'ｚ', 'Ａ', 'Ｂ', 'Ｃ', 'Ｄ', 'Ｅ', 'Ｆ', 'Ｇ', 'Ｈ', 'Ｉ', 'Ｊ', 'Ｋ', 'Ｌ', 'Ｍ', 'Ｎ', 'Ｏ', 'Ｐ', 'Ｑ', 'Ｒ', 'Ｓ', 'Ｔ', 'Ｕ', 'Ｖ', 'Ｗ', 'Ｘ', 'Ｙ', 'Ｚ', '！', '”', '＃', '＄', '％', '＆', '’', '（', '）', '＊', '＋', '，', '－', '．', '／', '：', '；', '＜', '＝', '＞', '？', '＠', '［', '￥', '］', '＾', '＿', '‘', '｛', '｜', '｝', '～', '\u3000'}
	ASCII_HANKAKU_CHARS    = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '¥', ']', '^', '_', '`', '{', '|', '}', '~', ' '}
	tabASCIIzh, tabASCIIhz = indexRunes(ASCII_ZENKAKU_CHARS, ASCII_HANKAKU_CHARS)

	KANA_ZENKAKU_CHARS   = []rune{'ア', 'イ', 'ウ', 'エ', 'オ', 'カ', 'キ', 'ク', 'ケ', 'コ', 'サ', 'シ', 'ス', 'セ', 'ソ', 'タ', 'チ', 'ツ', 'テ', 'ト', 'ナ', 'ニ', 'ヌ', 'ネ', 'ノ', 'ハ', 'ヒ', 'フ', 'ヘ', 'ホ', 'マ', 'ミ', 'ム', 'メ', 'モ', 'ヤ', 'ユ', 'ヨ', 'ラ', 'リ', 'ル', 'レ', 'ロ', 'ワ', 'ヲ', 'ン', 'ァ', 'ィ', 'ゥ', 'ェ', 'ォ', 'ッ', 'ャ', 'ュ', 'ョ', '。', '、', '・', '゛', '゜', '「', '」', 'ー'}
	KANA_HANKAKU_CHARS   = []rune{'ｱ', 'ｲ', 'ｳ', 'ｴ', 'ｵ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾅ', 'ﾆ', 'ﾇ', 'ﾈ', 'ﾉ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾏ', 'ﾐ', 'ﾑ', 'ﾒ', 'ﾓ', 'ﾔ', 'ﾕ', 'ﾖ', 'ﾗ', 'ﾘ', 'ﾙ', 'ﾚ', 'ﾛ', 'ﾜ', 'ｦ', 'ﾝ', 'ｧ', 'ｨ', 'ｩ', 'ｪ', 'ｫ', 'ｯ', 'ｬ', 'ｭ', 'ｮ', '｡', '､', '･', 'ﾞ', 'ﾟ', '｢', '｣', 'ｰ'}
	tabKANAzh, tabKanaHz = indexRunes(KANA_ZENKAKU_CHARS, KANA_HANKAKU_CHARS)

	DIGIT_ZENKAKU_CHARS    = []rune{'０', '１', '２', '３', '４', '５', '６', '７', '８', '９'}
	DIGIT_HANKAKU_CHARS    = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	tabDIGITzh, tabDigitHz = indexRunes(DIGIT_ZENKAKU_CHARS, DIGIT_HANKAKU_CHARS)

	KANA_TEN_MAP = map[rune]rune{
		'ガ': 'ｶ', 'ギ': 'ｷ', 'グ': 'ｸ', 'ゲ': 'ｹ', 'ゴ': 'ｺ',
		'ザ': 'ｻ', 'ジ': 'ｼ', 'ズ': 'ｽ', 'ゼ': 'ｾ', 'ゾ': 'ｿ',
		'ダ': 'ﾀ', 'ヂ': 'ﾁ', 'ヅ': 'ﾂ', 'デ': 'ﾃ', 'ド': 'ﾄ',
		'バ': 'ﾊ', 'ビ': 'ﾋ', 'ブ': 'ﾌ', 'ベ': 'ﾍ', 'ボ': 'ﾎ',
		'ヴ': 'ｳ',
	}
	tabKANATENzh, tabKanaTenHz = inverse(KANA_TEN_MAP)

	KANA_MARU_MAP = map[rune]rune{
		'パ': 'ﾊ', 'ピ': 'ﾋ', 'プ': 'ﾌ', 'ペ': 'ﾍ', 'ポ': 'ﾎ',
	}
	tabKANAMARUzh, tabKanaMaruHz = inverse(KANA_MARU_MAP)
)

// indexRunes maps runes from rune-set a to rune-set b and vice versa.  a and
// b should be of equal size, otherwise the function panics.
func indexRunes(a []rune, b []rune) (ab map[rune]rune, ba map[rune]rune) {
	if len(a) != len(b) {
		panic("slices of unequal sizes")
	}
	ab, ba = make(map[rune]rune, len(a)), make(map[rune]rune, len(b))
	for i := range a {
		ab[a[i]] = b[i]
		ba[b[i]] = a[i]
	}
	return
}

// inverse inverts the mapping in abIn, and returns both inverted and
// non-inverted maps.
func inverse(abIn map[rune]rune) (ab map[rune]rune, ba map[rune]rune) {
	ab = abIn
	ba = make(map[rune]rune, len(ab))
	for k, v := range ab {
		ba[v] = k
	}
	return
}
