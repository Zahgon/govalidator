package govalidator

import (
	"html"
)

func Contains(str, substring string) bool { _ = "STUB: not implemented"; return false }

func Matches(str, pattern string) bool { _ = "STUB: not implemented"; return false }

func LeftTrim(str, chars string) string { _ = "STUB: not implemented"; return "" }

func RightTrim(str, chars string) string { _ = "STUB: not implemented"; return "" }

func Trim(str, chars string) string { _ = "STUB: not implemented"; return "" }

func WhiteList(str, chars string) string { _ = "STUB: not implemented"; return "" }

func BlackList(str, chars string) string { _ = "STUB: not implemented"; return "" }

func StripLow(str string, keepNewLines bool) string { _ = "STUB: not implemented"; return "" }

func ReplacePattern(str, pattern, replace string) string { _ = "STUB: not implemented"; return "" }

var Escape = html.EscapeString

func addSegment(inrune, segment []rune) []rune { _ = "STUB: not implemented"; return nil }

func UnderscoreToCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func CamelCaseToUnderscore(str string) string { _ = "STUB: not implemented"; return "" }

func Reverse(s string) string { _ = "STUB: not implemented"; return "" }

func GetLines(s string) []string { _ = "STUB: not implemented"; return nil }

func GetLine(s string, index int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func RemoveTags(s string) string { _ = "STUB: not implemented"; return "" }

func SafeFileName(str string) string { _ = "STUB: not implemented"; return "" }

func NormalizeEmail(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Truncate(str string, length int, ending string) string { _ = "STUB: not implemented"; return "" }

func PadLeft(str string, padStr string, padLen int) string { _ = "STUB: not implemented"; return "" }

func PadRight(str string, padStr string, padLen int) string { _ = "STUB: not implemented"; return "" }

func PadBoth(str string, padStr string, padLen int) string { _ = "STUB: not implemented"; return "" }

func buildPadStr(str string, padStr string, padLen int, padLeft bool, padRight bool) string {
	_ = "STUB: not implemented"
	return ""
}

func TruncatingErrorf(str string, args ...interface{}) error { _ = "STUB: not implemented"; return nil }
