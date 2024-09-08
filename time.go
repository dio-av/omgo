package omgo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const atLayout = "2006-01-02T15:04"
const adLayout = "2006-01-02"
const dateTimeApiLayout = "2006-01-02T15:04"

var nilTime = (time.Time{}).UnixNano()

type ApiTime struct {
	time.Time
}

type ApiDate struct {
	time.Time
}

type ApiDateTime struct {
	time.Time
}

func (ct *ApiTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(atLayout, s)
	return
}

func (ct *ApiTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(atLayout))), nil
}

func (ct *ApiTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

func (ct *ApiDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(adLayout, s)
	return
}

func (ct *ApiDate) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(adLayout))), nil
}

func (ct *ApiDate) IsSet() bool {
	return ct.UnixNano() != nilTime
}

func (dt ApiDateTime) MarshalJSON() ([]byte, error) {
	formatted := dt.Time.Format(dateTimeApiLayout)

	return json.Marshal(formatted)
}

func (dt *ApiDateTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse(dateTimeApiLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	dt.Time = t

	return nil
}
