package typ

import (
	"encoding/json"
	"time"
)

type MicroTime time.Time

func (t MicroTime) ToTime() time.Time {
	return time.Time(t)
}

func (t MicroTime) String() string {
	return t.ToTime().String()
}

func (t MicroTime) ToUtcString() string {
	return t.ToTime().UTC().String()
}

func (t MicroTime) WithAddedDays(days int) MicroTime {
	return MicroTime(t.ToTime().AddDate(0, 0, days))
}

func (t MicroTime) TruncatedToSeconds() MicroTime {
	return MicroTime(time.Time(t).Truncate(time.Second))
}

func (t MicroTime) IsPast() bool {
	return t.ToTime().Before(NewUtcMicroTimeNow().ToTime())
}

func (t MicroTime) IsSame(compT MicroTime) bool {
	return t.ToUnix() == compT.ToUnix()
}

func (t MicroTime) ToUnix() int64 {
	return t.ToTime().Unix()
}

func (t MicroTime) ToUndefinableMicroTime() UndefinableMicroTime {
	return UndefinableMicroTime{NewUndefinable(t)}
}

func (t MicroTime) ToNullableMicroTime() NullableMicroTime {
	return NullableMicroTime{NewNullable(t)}
}

func (t MicroTime) ToSecondsFromNow() int {
	return int(t.ToUnix() - NewUtcMicroTimeNow().ToUnix())
}

func NewUtcMicroTimeNow() MicroTime {
	return MicroTime(time.Now().UTC())
}

func NewUtcMicroTimeFromTime(t time.Time) MicroTime {
	loc, _ := time.LoadLocation("UTC")
	return MicroTime(t.In(loc))
}

func (t *MicroTime) UnmarshalJSON(data []byte) error {
	tempTime := time.Time{}
	err := json.Unmarshal(data, &tempTime)
	if err != nil {
		return err
	}
	*t = MicroTime(tempTime.UTC())
	return nil
}

func (t MicroTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToTime())
}

type UndefinableMicroTime struct {
	Undefinable[MicroTime]
}

func (t UndefinableMicroTime) IsSame(compT UndefinableMicroTime) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 MicroTime, t2 MicroTime) bool {
		return t1.IsSame(t2)
	})
}

type NullableMicroTime struct {
	Nullable[MicroTime]
}

func (t NullableMicroTime) IsSame(compT NullableMicroTime) bool {
	return t.Nullable.Equals(compT.Nullable, func(t1 MicroTime, t2 MicroTime) bool {
		return t1.IsSame(t2)
	})
}

func (t NullableMicroTime) ToUndefinableNullableMicroTime() UndefinableNullableMicroTime {
	return UndefinableNullableMicroTime{NewUndefinable(t)}
}

type UndefinableNullableMicroTime struct {
	Undefinable[NullableMicroTime]
}

func (t UndefinableNullableMicroTime) IsSame(compT UndefinableNullableMicroTime) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 NullableMicroTime, t2 NullableMicroTime) bool {
		return t1.IsSame(t2)
	})
}
