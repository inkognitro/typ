package typ

import (
	"encoding/json"
	"fmt"
	"time"
)

const relativeDayTimeFormat = "2006-01-02T15:04"

type RelativeDayTime string

func (t RelativeDayTime) IsValid() bool {
	t1, err := t.ToTime("UTC")
	if err != nil {
		return false
	}
	return t.ToString() == t1.Format(relativeDayTimeFormat)
}

func (t RelativeDayTime) Equals(compT RelativeDayTime) bool {
	return compT.ToString() == t.ToString()
}

func (t RelativeDayTime) IsBefore(compT RelativeDayTime) bool {
	return t.ToTimeOrPanic("UTC").Before(compT.ToTimeOrPanic("UTC"))
}

func (t RelativeDayTime) ToTime(ianaTimezone string) (time.Time, error) {
	loc, err := time.LoadLocation(ianaTimezone)
	if err != nil {
		return time.Time{}, err
	}
	parsedTime, err := time.ParseInLocation(relativeDayTimeFormat, string(t), loc)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func (t RelativeDayTime) ToTimeOrPanic(ianaTimezone string) time.Time {
	parsedTime, err := t.ToTime(ianaTimezone)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func (t RelativeDayTime) ToString() string {
	return string(t)
}

func (t RelativeDayTime) String() string {
	return t.ToString()
}

func NewRelativeDayTime(content string) (RelativeDayTime, error) {
	t := RelativeDayTime(content)
	if !t.IsValid() {
		return "", fmt.Errorf(`"%s" has invalid RelativeDayTime format`, content)
	}
	return t, nil
}

func (t *RelativeDayTime) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	tempT, err := NewRelativeDayTime(temp)
	if err != nil {
		return err
	}
	*t = tempT
	return nil
}

func (t RelativeDayTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToString())
}
