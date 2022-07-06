package typ

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash struct {
	hash                 []byte
	isUtf8RuneCountValid bool
	utf8RuneCount        int
}

func (t PasswordHash) FindPasswordsUtf8RuneCount() (int, bool) {
	return t.utf8RuneCount, t.isUtf8RuneCountValid
}

func (t PasswordHash) ToString() string {
	return string(t.hash)
}

func (t PasswordHash) ToUndefinablePasswordHash() UndefinablePasswordHash {
	return UndefinablePasswordHash{NewUndefinable(t)}
}

func (t PasswordHash) IsSame(compHash PasswordHash) bool {
	return t.ToString() == compHash.ToString()
}

func (t PasswordHash) String() string {
	return t.ToString()
}

func (t PasswordHash) Validate(password String) bool {
	err := bcrypt.CompareHashAndPassword(t.hash, []byte(password.ToString()))
	return err == nil
}

func NewPasswordHashFromStringType(password String, cost int) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return PasswordHash{}, fmt.Errorf("could not create password hash: %w", err)
	}
	return PasswordHash{hash: hash, utf8RuneCount: password.GetUtf8RuneCount(), isUtf8RuneCountValid: true}, nil
}

func NewPasswordHash(passwordHash string) PasswordHash {
	return PasswordHash{hash: []byte(passwordHash)}
}

func (t *PasswordHash) UnmarshalJSON(data []byte) error {
	tempHash := ""
	if err := json.Unmarshal(data, &tempHash); err != nil {
		return err
	}
	t.hash = []byte(tempHash)
	return nil
}

func (t PasswordHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.hash)
}

type UndefinablePasswordHash struct {
	Undefinable[PasswordHash]
}
