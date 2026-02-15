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
	PositionAdminIdx  EmployeePositionIdx = 3
	PositionHRDIdx    EmployeePositionIdx = 2
	PositionWorkerIdx EmployeePositionIdx = 1

	PositionAdminKey  EmployeePositionKey = "admin"
	PositionHRDKey    EmployeePositionKey = "hrd"
	PositionWorkerKey EmployeePositionKey = "worker"
)

var (
	PositionAdminValue = EmployeePositionValue{
		Idx: int64(PositionAdminIdx),
		Key: string(PositionAdminKey),
		LongLabel: map[string]string{
			"id": "Admin",
			"en": "Admin",
		},
		ShortLabel: map[string]string{
			"id": "Admin",
			"en": "Admin",
		},
	}

	PositionHRDValue = EmployeePositionValue{
		Idx: int64(PositionHRDIdx),
		Key: string(PositionHRDKey),
		LongLabel: map[string]string{
			"id": "HRD",
			"en": "Human Resources",
		},
		ShortLabel: map[string]string{
			"id": "HRD",
			"en": "HR",
		},
	}

	PositionWorkerValue = EmployeePositionValue{
		Idx: int64(PositionWorkerIdx),
		Key: string(PositionWorkerKey),
		LongLabel: map[string]string{
			"id": "Pekerja",
			"en": "Worker",
		},
		ShortLabel: map[string]string{
			"id": "Pkrj",
			"en": "Wkr",
		},
	}
)

var (
	EmployeePositionMapIdx = map[EmployeePositionIdx]EmployeePositionValue{
		PositionAdminIdx:  PositionAdminValue,
		PositionHRDIdx:    PositionHRDValue,
		PositionWorkerIdx: PositionWorkerValue,
	}

	EmployeePositionMapKey = map[EmployeePositionKey]EmployeePositionValue{
		PositionAdminKey:  PositionAdminValue,
		PositionHRDKey:    PositionHRDValue,
		PositionWorkerKey: PositionWorkerValue,
	}

	EmployeePositionMapReverse = map[EmployeePositionKey]EmployeePositionIdx{
		PositionAdminKey:  PositionAdminIdx,
		PositionHRDKey:    PositionHRDIdx,
		PositionWorkerKey: PositionWorkerIdx,
	}
)

func (p EmployeePositionIdx) String() string {
	if position, ok := EmployeePositionMapIdx[p]; ok {
		return position.Key
	}
	return ""
}

func (p *EmployeePositionIdx) Scan(value interface{}) error {
	if value == nil {
		return errors.New("position cannot be null")
	}

	switch v := value.(type) {
	case int64:
		*p = EmployeePositionIdx(v)
		return nil
	case string:
		if idx, ok := EmployeePositionMapReverse[EmployeePositionKey(v)]; ok {
			*p = idx
			return nil
		}
	case []byte:
		if idx, ok := EmployeePositionMapReverse[EmployeePositionKey(string(v))]; ok {
			*p = idx
			return nil
		}
	}

	return errors.New("invalid employee position value")
}

func (p EmployeePositionIdx) Value() (driver.Value, error) {
	return int64(p), nil
}
