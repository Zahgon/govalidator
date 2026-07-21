package govalidator

import (
	"reflect"
	"regexp"
)

var (
	fieldsRequiredByDefault bool
	nilPtrAllowedByRequired = false
	notNumberRegexp         = regexp.MustCompile("[^0-9]+")
	whiteSpacesAndMinus     = regexp.MustCompile(`[\s-]+`)
	paramsRegexp            = regexp.MustCompile(`\(.*\)$`)
	rxJWT                   = regexp.MustCompile(`^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+$`)
)

const maxURLRuneCount = 2083
const minURLRuneCount = 3
const rfc3339WithoutZone = "2006-01-02T15:04:05"

func SetFieldsRequiredByDefault(value bool) { _ = "STUB: not implemented"; return }

func SetNilPtrAllowedByRequired(value bool) { _ = "STUB: not implemented"; return }

func IsEmail(str string) bool { _ = "STUB: not implemented"; return false }

func IsExistingEmail(email string) bool { _ = "STUB: not implemented"; return false }

func IsURL(str string) bool { _ = "STUB: not implemented"; return false }

func IsRequestURL(rawurl string) bool { _ = "STUB: not implemented"; return false }

func IsRequestURI(rawurl string) bool { _ = "STUB: not implemented"; return false }

func IsAlpha(str string) bool { _ = "STUB: not implemented"; return false }

func IsUTFLetter(str string) bool { _ = "STUB: not implemented"; return false }

func IsAlphanumeric(str string) bool { _ = "STUB: not implemented"; return false }

func IsUTFLetterNumeric(str string) bool { _ = "STUB: not implemented"; return false }

func IsNumeric(str string) bool { _ = "STUB: not implemented"; return false }

func IsUTFNumeric(str string) bool { _ = "STUB: not implemented"; return false }

func IsUTFDigit(str string) bool { _ = "STUB: not implemented"; return false }

func IsHexadecimal(str string) bool { _ = "STUB: not implemented"; return false }

func IsHexcolor(str string) bool { _ = "STUB: not implemented"; return false }

func IsRGBcolor(str string) bool { _ = "STUB: not implemented"; return false }

func IsLowerCase(str string) bool { _ = "STUB: not implemented"; return false }

func IsUpperCase(str string) bool { _ = "STUB: not implemented"; return false }

func HasLowerCase(str string) bool { _ = "STUB: not implemented"; return false }

func HasUpperCase(str string) bool { _ = "STUB: not implemented"; return false }

func IsInt(str string) bool { _ = "STUB: not implemented"; return false }

func IsFloat(str string) bool { _ = "STUB: not implemented"; return false }

func IsDivisibleBy(str, num string) bool { _ = "STUB: not implemented"; return false }

func IsNull(str string) bool { _ = "STUB: not implemented"; return false }

func IsNotNull(str string) bool { _ = "STUB: not implemented"; return false }

func HasWhitespaceOnly(str string) bool { _ = "STUB: not implemented"; return false }

func HasWhitespace(str string) bool { _ = "STUB: not implemented"; return false }

func IsByteLength(str string, min, max int) bool { _ = "STUB: not implemented"; return false }

func IsUUIDv3(str string) bool { _ = "STUB: not implemented"; return false }

func IsUUIDv4(str string) bool { _ = "STUB: not implemented"; return false }

func IsUUIDv5(str string) bool { _ = "STUB: not implemented"; return false }

func IsUUID(str string) bool { _ = "STUB: not implemented"; return false }

var ulidDec = [...]byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x01,
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
	0x0F, 0x10, 0x11, 0xFF, 0x12, 0x13, 0xFF, 0x14, 0x15, 0xFF,
	0x16, 0x17, 0x18, 0x19, 0x1A, 0xFF, 0x1B, 0x1C, 0x1D, 0x1E,
	0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0A, 0x0B, 0x0C,
	0x0D, 0x0E, 0x0F, 0x10, 0x11, 0xFF, 0x12, 0x13, 0xFF, 0x14,
	0x15, 0xFF, 0x16, 0x17, 0x18, 0x19, 0x1A, 0xFF, 0x1B, 0x1C,
	0x1D, 0x1E, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
}

const ulidEncodedSize = 26

func IsULID(str string) bool { _ = "STUB: not implemented"; return false }

func IsCreditCard(str string) bool { _ = "STUB: not implemented"; return false }

func IsISBN10(str string) bool { _ = "STUB: not implemented"; return false }

func IsISBN13(str string) bool { _ = "STUB: not implemented"; return false }

func IsISBN(str string, version int) bool { _ = "STUB: not implemented"; return false }

func IsJSON(str string) bool { _ = "STUB: not implemented"; return false }

func IsMultibyte(str string) bool { _ = "STUB: not implemented"; return false }

func IsASCII(str string) bool { _ = "STUB: not implemented"; return false }

func IsPrintableASCII(str string) bool { _ = "STUB: not implemented"; return false }

func IsFullWidth(str string) bool { _ = "STUB: not implemented"; return false }

func IsHalfWidth(str string) bool { _ = "STUB: not implemented"; return false }

func IsVariableWidth(str string) bool { _ = "STUB: not implemented"; return false }

func IsBase64(str string) bool { _ = "STUB: not implemented"; return false }

func IsJWT(str string) bool { _ = "STUB: not implemented"; return false }

func IsFilePath(str string) (bool, int) { _ = "STUB: not implemented"; return false, 0 }

func IsWinFilePath(str string) bool { _ = "STUB: not implemented"; return false }

func IsUnixFilePath(str string) bool { _ = "STUB: not implemented"; return false }

func IsDataURI(str string) bool { _ = "STUB: not implemented"; return false }

func IsMagnetURI(str string) bool { _ = "STUB: not implemented"; return false }

func IsISO3166Alpha2(str string) bool { _ = "STUB: not implemented"; return false }

func IsISO3166Alpha3(str string) bool { _ = "STUB: not implemented"; return false }

func IsISO693Alpha2(str string) bool { _ = "STUB: not implemented"; return false }

func IsISO693Alpha3b(str string) bool { _ = "STUB: not implemented"; return false }

func IsDNSName(str string) bool { _ = "STUB: not implemented"; return false }

func IsHash(str string, algorithm string) bool { _ = "STUB: not implemented"; return false }

func IsSHA3224(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA3256(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA3384(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA3512(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA512(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA384(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA256(str string) bool { _ = "STUB: not implemented"; return false }

func IsTiger192(str string) bool { _ = "STUB: not implemented"; return false }

func IsTiger160(str string) bool { _ = "STUB: not implemented"; return false }

func IsRipeMD160(str string) bool { _ = "STUB: not implemented"; return false }

func IsSHA1(str string) bool { _ = "STUB: not implemented"; return false }

func IsTiger128(str string) bool { _ = "STUB: not implemented"; return false }

func IsRipeMD128(str string) bool { _ = "STUB: not implemented"; return false }

func IsCRC32(str string) bool { _ = "STUB: not implemented"; return false }

func IsCRC32b(str string) bool { _ = "STUB: not implemented"; return false }

func IsMD5(str string) bool { _ = "STUB: not implemented"; return false }

func IsMD4(str string) bool { _ = "STUB: not implemented"; return false }

func IsDialString(str string) bool { _ = "STUB: not implemented"; return false }

func IsIP(str string) bool { _ = "STUB: not implemented"; return false }

func IsPort(str string) bool { _ = "STUB: not implemented"; return false }

func IsIPv4(str string) bool { _ = "STUB: not implemented"; return false }

func IsIPv6(str string) bool { _ = "STUB: not implemented"; return false }

func IsCIDR(str string) bool { _ = "STUB: not implemented"; return false }

func IsMAC(str string) bool { _ = "STUB: not implemented"; return false }

func IsHost(str string) bool { _ = "STUB: not implemented"; return false }

func IsMongoID(str string) bool { _ = "STUB: not implemented"; return false }

func IsLatitude(str string) bool { _ = "STUB: not implemented"; return false }

func IsLongitude(str string) bool { _ = "STUB: not implemented"; return false }

func IsIMEI(str string) bool { _ = "STUB: not implemented"; return false }

func IsIMSI(str string) bool { _ = "STUB: not implemented"; return false }

func IsRsaPublicKey(str string, keylen int) bool { _ = "STUB: not implemented"; return false }

func IsRegex(str string) bool { _ = "STUB: not implemented"; return false }

func toJSONName(tag string) string { _ = "STUB: not implemented"; return "" }

func prependPathToErrors(err error, path string) error { _ = "STUB: not implemented"; return nil }

func ValidateArray(array []interface{}, iterator ConditionIterator) bool {
	_ = "STUB: not implemented"
	return false
}

func ValidateMap(s map[string]interface{}, m map[string]interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ValidateStruct(s interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ValidateStructAsync(s interface{}) (<-chan bool, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateMapAsync(s map[string]interface{}, m map[string]interface{}) (<-chan bool, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseTagIntoMap(tag string) tagOptionsMap {
	_ = "STUB: not implemented"
	return *new(tagOptionsMap)
}

func isValidTag(s string) bool { _ = "STUB: not implemented"; return false }

func IsSSN(str string) bool { _ = "STUB: not implemented"; return false }

func IsSemver(str string) bool { _ = "STUB: not implemented"; return false }

func IsType(v interface{}, params ...string) bool { _ = "STUB: not implemented"; return false }

func IsTime(str string, format string) bool { _ = "STUB: not implemented"; return false }

func IsUnixTime(str string) bool { _ = "STUB: not implemented"; return false }

func IsRFC3339(str string) bool { _ = "STUB: not implemented"; return false }

func IsRFC3339WithoutZone(str string) bool { _ = "STUB: not implemented"; return false }

func IsISO4217(str string) bool { _ = "STUB: not implemented"; return false }

func ByteLength(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func RuneLength(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func IsRsaPub(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func StringMatches(s string, params ...string) bool { _ = "STUB: not implemented"; return false }

func StringLength(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func MinStringLength(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func MaxStringLength(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func Range(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func IsInRaw(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func IsIn(str string, params ...string) bool { _ = "STUB: not implemented"; return false }

func checkRequired(v reflect.Value, t reflect.StructField, options tagOptionsMap) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func typeCheck(v reflect.Value, t reflect.StructField, o reflect.Value, options tagOptionsMap) (isValid bool, resultErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func stripParams(validatorString string) string { _ = "STUB: not implemented"; return "" }

func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func ErrorByField(e error, field string) string { _ = "STUB: not implemented"; return "" }

func ErrorsByField(e error) map[string]string { _ = "STUB: not implemented"; return nil }

func (e *UnsupportedTypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (sv stringValues) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sv stringValues) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sv stringValues) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (sv stringValues) get(i int) string   { _ = "STUB: not implemented"; return "" }

func IsE164(str string) bool { _ = "STUB: not implemented"; return false }

func IsYYYYMMDD(str string) bool { _ = "STUB: not implemented"; return false }
