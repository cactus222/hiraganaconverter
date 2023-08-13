package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		input, _ := reader.ReadString('\n')
		output := parse(input)
		fmt.Printf("%v\n", output)
	}
}

func parse(input string) string {
	var result string

	var buffer = []rune{}

	for _, character := range input {
		if isVowel(character) {
			buffer = append(buffer, character)
			output := parseBuffer(buffer)
			result += output
			buffer = []rune{}
		} else if isConsonant(character) {
			// need to check if previous is a solo n
			if len(buffer) == 1 && buffer[0] == 'n' {
				output := parseBuffer(buffer)
				result += output
				buffer = []rune{}
			}
			buffer = append(buffer, character)
		} else {
			// what is this char - dump buffer + dump this char
			output := parseBuffer(buffer)
			result += output
			output = parseBuffer([]rune{character})
			result += output
			buffer = []rune{}
		}
	}

	output := parseBuffer(buffer)
	result += output

	return result
}

var charMapping = map[string]string{
	"a":  "あ",
	"i":  "い",
	"u":  "う",
	"e":  "え",
	"o":  "お",
	"ka": "か", "ga": "が",
	"ki": "き", "gi": "ぎ",
	"ku": "く", "gu": "ぐ",
	"ke": "け", "ge": "げ",
	"ko": "こ", "go": "ご",
	"sa": "さ", "za": "ざ",
	"shi": "し", "ji": "じ",
	"su": "す", "zu": "ず",
	"se": "せ", "ze": "ぜ",
	"so": "そ", "zo": "ぞ",
	"ta": "た", "da": "だ",
	"chi": "ち", "zhi": "ぢ", // or is this ji?... tbh idk if i've seen this char before
	"tsu": "つ", "tzu": "づ",
	"te": "て", "de": "で",
	"to": "と", "do": "ど",
	"na": "な",
	"ni": "に",
	"nu": "ぬ",
	"ne": "ね",
	"no": "の",
	"ha": "は", "ba": "ば", "pa": "ぱ",
	"hi": "ひ", "bi": "び", "pi": "ぴ",
	"fu": "ふ", "bu": "ぶ", "pu": "ぷ",
	"he": "へ", "be": "べ", "pe": "ぺ",
	"ho": "ほ", "bo": "ぼ", "po": "ぽ",
	"ma":  "ま",
	"mi":  "み",
	"mu":  "む",
	"me":  "め",
	"mo":  "も",
	"ya":  "や",
	"yu":  "ゆ",
	"yo":  "よ",
	"ra":  "ら",
	"ri":  "り",
	"ru":  "る",
	"re":  "れ",
	"ro":  "ろ",
	"wa":  "わ",
	"wo":  "を",
	"n":   "ん",
	"kya": "きゃ", "kyu": "きゅ", "kyo": "きょ",
	"sha": "しゃ", "shu": "しゅ", "sho": "しょ",
	"cha": "ちゃ", "chu": "ちゅ", "cho": "ちょ",
	"nya": "にゃ", "nyu": "にゅ", "nyo": "にょ",
	"hya": "ひゃ", "hyu": "ひゅ", "hyo": "ひょ",
	"mya": "みゃ", "myu": "みゅ", "myo": "みょ",
	"rya": "りゃ", "ryu": "りゅ", "ryo": "りょ",
	"gya": "ぎゃ", "gyu": "ぎゅ", "gyo": "ぎょ",
	"ja": "じゃ", "ju": "じゅ", "jo": "じょ",
	"zha": "ぢゃ", "zhu": "ぢゅ", "zho": "ぢょ", // same thing here... ?
	"bya": "びゃ", "byu": "びゅ", "byo": "びょ",
	"pya": "ぴゃ", "pyu": "ぴゅ", "pyo": "ぴょ",
}

func parseBuffer(buffer []rune) string {

	var result string
	// idk when i can use っ, but its just a double const, so lets just insert it if it matches
	if len(buffer) > 2 {
		if buffer[0] == buffer[1] {
			result = result + "っ"
			buffer = buffer[1:]
		}
	}
	if data, exists := charMapping[string(buffer)]; exists {
		return result + data
	}
	return string(buffer)

}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func isConsonant(char rune) bool {
	return (char >= 'a' && char <= 'z') && !isVowel(char)
}
