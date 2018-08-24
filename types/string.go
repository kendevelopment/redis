package types

import (
	"time"
)

const StringType = uint64(0)
const StringTypeName = "string"

type String struct {
	str *string

	expires bool // can exp
	expiry  time.Time
}

func NewString(s *string, expires bool, expiry time.Time) *String {
	return &String{
		str:     s,
		expiry:  expiry,
		expires: expires,
	}
}

func (s *String) Value() interface{} {
	return s.str
}

func (s *String) ValueTypeId() uint64 {
	return StringType
}

func (s *String) ValueType() string {
	return StringTypeName
}

func (s *String) Expiry() time.Time {
	return s.expiry
}

func (s *String) Expires() bool {
	return s.expires
}

func (s *String) OnDelete(key *string) {
}