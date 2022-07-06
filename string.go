package typ

import (
	"crypto/rand"
	"encoding/json"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type String string

func (t String) ToString() string {
	return string(t)
}

func (t String) ToUndefinableString() UndefinableString {
	return UndefinableString{NewUndefinable(t)}
}

func (t String) ToNullableString() NullableString {
	return NullableString{NewNullable(t)}
}

func (t String) ToUndefinableNullableString() UndefinableNullableString {
	return t.ToNullableString().ToUndefinableNullableString()
}

func (t String) ToPasswordHash(cost int) (PasswordHash, error) {
	return NewPasswordHashFromStringType(t, cost)
}

func (t String) ToLowerCaseString() string {
	return strings.ToLower(string(t))
}

func (t String) GetUtf8RuneCount() int {
	return utf8.RuneCountInString(string(t))
}

func (t String) ToStrings() Strings {
	return []String{t}
}

func (t String) String() string {
	return t.ToString()
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (t String) HasEmailAddressFormat() bool {
	s := t.ToString()
	if len(s) < 3 && len(s) > 254 {
		return false
	}
	return emailRegex.MatchString(s)
}

func (t String) HasUrlFormat() bool {
	_, err := url.ParseRequestURI(string(t))
	return err == nil
}

func (t String) EqualsCaseInSensitive(compT String) bool {
	return strings.EqualFold(t.ToString(), compT.ToString())
}

func (t String) EqualsCaseSensitive(compT String) bool {
	return t.ToString() == compT.ToString()
}

func NewString(content string) String {
	return String(content)
}

const alphaNumericChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewSecureRandomAlphaNumericString(length int) String {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	for i, b := range bytes {
		bytes[i] = alphaNumericChars[b%byte(len(alphaNumericChars))]
	}
	return String(bytes)
}

func (t *String) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*t = String(temp)
	return nil
}

func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

type NullableString struct {
	Nullable[String]
}

func (t NullableString) Value() (String, bool) {
	return t.Nullable.Value()
}

func (t NullableString) EqualsCaseSensitive(compT NullableString) bool {
	return t.Nullable.Equals(compT.Nullable, func(t1 String, t2 String) bool {
		return t1.EqualsCaseSensitive(t2)
	})
}

func (t NullableString) ToUndefinableNullableString() UndefinableNullableString {
	return UndefinableNullableString{NewUndefinable(t)}
}

type UndefinableString struct {
	Undefinable[String]
}

func (t UndefinableString) ToStringType() (String, bool) {
	return t.Value()
}

func (t UndefinableString) ToString() (string, bool) {
	strType, ok := t.ToStringType()
	return strType.ToString(), ok
}

type UndefinableNullableString struct {
	Undefinable[NullableString]
}

func (t UndefinableNullableString) EqualsCaseSensitive(compT UndefinableNullableString) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 NullableString, t2 NullableString) bool {
		return t1.EqualsCaseSensitive(t2)
	})
}
