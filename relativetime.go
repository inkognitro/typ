package typ

import (
	"encoding/json"
	"fmt"
	"time"
)

const relativeTimeFormat = "15:04"

type RelativeTime string

func (t RelativeTime) IsValid() bool {
	t1, err := t.ToTime("UTC")
	if err != nil {
		return false
	}
	return t.ToString() == t1.Format(relativeTimeFormat)
}

func (t RelativeTime) Equals(compT RelativeTime) bool {
	return compT.ToString() == t.ToString()
}

func (t RelativeTime) IsBefore(compT RelativeTime) bool {
	return t.ToTimeOrPanic("UTC").Before(compT.ToTimeOrPanic("UTC"))
}

func (t RelativeTime) ToTime(ianaTimezone string) (time.Time, error) {
	loc, err := time.LoadLocation(ianaTimezone)
	if err != nil {
		return time.Time{}, err
	}
	parsedTime, err := time.ParseInLocation(relativeTimeFormat, string(t), loc)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func (t RelativeTime) ToTimeOrPanic(ianaTimezone string) time.Time {
	parsedTime, err := t.ToTime(ianaTimezone)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func (t RelativeTime) ToString() string {
	return string(t)
}

func (t RelativeTime) String() string {
	return t.ToString()
}

func NewRelativeTime(content string) (RelativeTime, error) {
	t := RelativeTime(content)
	if !t.IsValid() {
		return "", fmt.Errorf(`"%s" has invalid RelativeTime format`, content)
	}
	return t, nil
}

func (t *RelativeTime) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	tempT, err := NewRelativeTime(temp)
	if err != nil {
		return err
	}
	*t = tempT
	return nil
}

func (t RelativeTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToString())
}
