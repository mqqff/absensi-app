package enums

import (
	"database/sql/driver"
	"errors"
)

type EmployeePositionIdx int64
type EmployeePositionKey string
type EmployeePositionValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	PositionUnknownIdx  EmployeePositionIdx = 1
	PositionHRIdx       EmployeePositionIdx = 2
	PositionAdminIdx    EmployeePositionIdx = 3
	PositionEmployeeIdx EmployeePositionIdx = 4

	PositionUnknownKey  EmployeePositionKey = "unknown"
	PositionHRKey       EmployeePositionKey = "hr"
	PositionAdminKey    EmployeePositionKey = "admin"
	PositionEmployeeKey EmployeePositionKey = "employee"
)

var (
	PositionUnknownValue = EmployeePositionValue{
		Idx:        int64(PositionUnknownIdx),
		Key:        string(PositionUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}

	PositionHRValue = EmployeePositionValue{
		Idx:        int64(PositionHRIdx),
		Key:        string(PositionHRKey),
		LongLabel:  map[string]string{"id": "Human Resources", "en": "Human Resources"},
		ShortLabel: map[string]string{"id": "HR", "en": "HR"},
	}

	PositionAdminValue = EmployeePositionValue{
		Idx:        int64(PositionAdminIdx),
		Key:        string(PositionAdminKey),
		LongLabel:  map[string]string{"id": "Administrator", "en": "Administrator"},
		ShortLabel: map[string]string{"id": "Admin", "en": "Admin"},
	}

	PositionEmployeeValue = EmployeePositionValue{
		Idx:        int64(PositionEmployeeIdx),
		Key:        string(PositionEmployeeKey),
		LongLabel:  map[string]string{"id": "Pegawai", "en": "Employee"},
		ShortLabel: map[string]string{"id": "Pegawai", "en": "Emp"},
	}
)

var (
	EmployeePositionMapIdx = map[EmployeePositionIdx]EmployeePositionValue{
		PositionUnknownIdx:  PositionUnknownValue,
		PositionHRIdx:       PositionHRValue,
		PositionAdminIdx:    PositionAdminValue,
		PositionEmployeeIdx: PositionEmployeeValue,
	}

	EmployeePositionMapKey = map[EmployeePositionKey]EmployeePositionValue{
		PositionUnknownKey:  PositionUnknownValue,
		PositionHRKey:       PositionHRValue,
		PositionAdminKey:    PositionAdminValue,
		PositionEmployeeKey: PositionEmployeeValue,
	}

	EmployeePositionMapReverse = map[EmployeePositionKey]EmployeePositionIdx{
		PositionUnknownKey:  PositionUnknownIdx,
		PositionHRKey:       PositionHRIdx,
		PositionAdminKey:    PositionAdminIdx,
		PositionEmployeeKey: PositionEmployeeIdx,
	}
)

func (p EmployeePositionIdx) String() string {
	if position, ok := EmployeePositionMapIdx[p]; ok {
		return position.Key
	}

	return string(PositionUnknownKey)
}

func (p *EmployeePositionIdx) Scan(value interface{}) error {
	if value == nil {
		*p = PositionUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*p = EmployeePositionIdx(v)
		return nil
	case int16:
		*p = EmployeePositionIdx(v)
		return nil
	case int32:
		*p = EmployeePositionIdx(v)
		return nil
	case int64:
		*p = EmployeePositionIdx(v)
		return nil
	case string:
		if idx, ok := EmployeePositionMapReverse[EmployeePositionKey(v)]; ok {
			*p = idx
			return nil
		}
	case []byte:
		strVal := string(v)
		if idx, ok := EmployeePositionMapReverse[EmployeePositionKey(strVal)]; ok {
			*p = idx
			return nil
		}
	}

	return errors.New("invalid user position value")
}

func (p EmployeePositionIdx) Value() (int64, error) {
	return int64(p), nil
}

type NullEmployeePositionIdx struct {
	EmployeePositionIdx EmployeePositionIdx
	Valid               bool
}

func (p *NullEmployeePositionIdx) Scan(value interface{}) error {
	if value == nil {
		p.EmployeePositionIdx, p.Valid = PositionUnknownIdx, false
		return nil
	}

	p.Valid = true
	return p.EmployeePositionIdx.Scan(value)
}

func (p NullEmployeePositionIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.EmployeePositionIdx.Value()
}
