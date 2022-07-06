package typ

import (
	"encoding/json"
	"testing"
)

type anyOfCarAudi struct {
	IsS bool `json:"isS"`
}

type anyOfCarBmw struct {
	IsM bool `json:"isM"`
}

type CarAnyOfRuler struct{}

func (CarAnyOfRuler) Discriminator() string {
	return "type"
}

func (CarAnyOfRuler) SupportedValuesByDiscriminatorValue() map[string]any {
	return map[string]any{
		"audi": &anyOfCarAudi{},
		"bmw":  &anyOfCarBmw{},
	}
}

type anyOfTestCar struct {
	AnyOf[CarAnyOfRuler]
}

func TestParseAnyOfCorrectly(t *testing.T) {
	strToParse := `{"foo":"bar","type":"bmw","payload":{}}`
	bytesToParse := []byte(strToParse)
	entry := anyOfTestCar{}
	err := json.Unmarshal(bytesToParse, &entry)
	if err != nil {
		t.Errorf("could not execute json.Unmarshal: %v", err)
		return
	}
	actualBytes, err := entry.MarshalJSON()
	actualStr := string(actualBytes)
	if err != nil {
		panic(err)
	}
	expectedStr := `{"isM":false,"type":"bmw"}`
	if actualStr != expectedStr {
		t.Errorf(`did receive %s instead of %s`, actualStr, expectedStr)
	}
}

func TestCannotParseInvalidAnyOfFormat(t *testing.T) {
	strToParse := "\"foo\""
	bytesToParse := []byte(strToParse)
	entry := anyOfTestCar{}
	err := json.Unmarshal(bytesToParse, &entry)
	if err == nil {
		t.Errorf("parsed foo as valid json")
	}
}
