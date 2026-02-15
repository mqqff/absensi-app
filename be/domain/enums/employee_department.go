package enums

import (
	"database/sql/driver"
	"errors"
)

type EmployeeDepartmentIdx int64
type EmployeeDepartmentKey string
type EmployeeDepartmentValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	LongLabel  map[string]string `json:"long_label"`
	ShortLabel map[string]string `json:"short_label"`
}

const (
	DepartmentUnknownIdx     EmployeeDepartmentIdx = 0
	DepartmentHRIdx          EmployeeDepartmentIdx = 1
	DepartmentEngineeringIdx EmployeeDepartmentIdx = 2
	DepartmentFinanceIdx     EmployeeDepartmentIdx = 3
	DepartmentOperationsIdx  EmployeeDepartmentIdx = 4
	DepartmentSalesIdx       EmployeeDepartmentIdx = 5

	DepartmentUnknownDepartmentKey EmployeeDepartmentKey = "unknown"
	DepartmentHRKey                EmployeeDepartmentKey = "hr"
	DepartmentEngineeringKey       EmployeeDepartmentKey = "engineering"
	DepartmentFinanceKey           EmployeeDepartmentKey = "finance"
	DepartmentOperationsKey        EmployeeDepartmentKey = "operations"
	DepartmentSalesKey             EmployeeDepartmentKey = "sales"
)

var (
	DepartmentUnknownValue = EmployeeDepartmentValue{
		ShortLabel: map[string]string{"id": "", "en": ""},
		Idx:        int64(DepartmentUnknownIdx),
		Key:        string(DepartmentUnknownDepartmentKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
	}

	DepartmentHRValue = EmployeeDepartmentValue{
		Idx:        int64(DepartmentHRIdx),
		Key:        string(DepartmentHRKey),
		LongLabel:  map[string]string{"id": "Sumber Daya Manusia", "en": "Human Resources"},
		ShortLabel: map[string]string{"id": "HRD", "en": "HR"},
	}

	DepartmentEngineeringValue = EmployeeDepartmentValue{
		Idx:        int64(DepartmentEngineeringIdx),
		Key:        string(DepartmentEngineeringKey),
		LongLabel:  map[string]string{"id": "Engineering / IT", "en": "Engineering"},
		ShortLabel: map[string]string{"id": "ENG", "en": "ENG"},
	}

	DepartmentFinanceValue = EmployeeDepartmentValue{
		Idx:        int64(DepartmentFinanceIdx),
		Key:        string(DepartmentFinanceKey),
		LongLabel:  map[string]string{"id": "Keuangan", "en": "Finance"},
		ShortLabel: map[string]string{"id": "FIN", "en": "FIN"},
	}

	DepartmentOperationsValue = EmployeeDepartmentValue{
		Idx:        int64(DepartmentOperationsIdx),
		Key:        string(DepartmentOperationsKey),
		LongLabel:  map[string]string{"id": "Operasional", "en": "Operations"},
		ShortLabel: map[string]string{"id": "OPS", "en": "OPS"},
	}

	DepartmentSalesValue = EmployeeDepartmentValue{
		Idx:        int64(DepartmentSalesIdx),
		Key:        string(DepartmentSalesKey),
		LongLabel:  map[string]string{"id": "Sales & Marketing", "en": "Sales"},
		ShortLabel: map[string]string{"id": "SLS", "en": "SLS"},
	}
)

var (
	EmployeeDepartmentMapIdx = map[EmployeeDepartmentIdx]EmployeeDepartmentValue{
		DepartmentUnknownIdx:     DepartmentUnknownValue,
		DepartmentHRIdx:          DepartmentHRValue,
		DepartmentEngineeringIdx: DepartmentEngineeringValue,
		DepartmentFinanceIdx:     DepartmentFinanceValue,
		DepartmentOperationsIdx:  DepartmentOperationsValue,
		DepartmentSalesIdx:       DepartmentSalesValue,
	}

	EmployeeDepartmentMapKey = map[EmployeeDepartmentKey]EmployeeDepartmentValue{
		DepartmentUnknownDepartmentKey: DepartmentUnknownValue,
		DepartmentHRKey:                DepartmentHRValue,
		DepartmentEngineeringKey:       DepartmentEngineeringValue,
		DepartmentFinanceKey:           DepartmentFinanceValue,
		DepartmentOperationsKey:        DepartmentOperationsValue,
		DepartmentSalesKey:             DepartmentSalesValue,
	}

	EmployeeDepartmentMapReverse = map[EmployeeDepartmentKey]EmployeeDepartmentIdx{
		DepartmentUnknownDepartmentKey: DepartmentUnknownIdx,
		DepartmentHRKey:                DepartmentHRIdx,
		DepartmentEngineeringKey:       DepartmentEngineeringIdx,
		DepartmentFinanceKey:           DepartmentFinanceIdx,
		DepartmentOperationsKey:        DepartmentOperationsIdx,
		DepartmentSalesKey:             DepartmentSalesIdx,
	}
)

func (d EmployeeDepartmentIdx) String() string {
	if department, ok := EmployeeDepartmentMapIdx[d]; ok {
		return department.Key
	}
	return string(DepartmentUnknownDepartmentKey)
}

func (d *EmployeeDepartmentIdx) Scan(value interface{}) error {
	if value == nil {
		*d = DepartmentUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*d = EmployeeDepartmentIdx(v)
		return nil
	case int16:
		*d = EmployeeDepartmentIdx(v)
		return nil
	case int32:
		*d = EmployeeDepartmentIdx(v)
		return nil
	case int64:
		*d = EmployeeDepartmentIdx(v)
		return nil
	case string:
		if idx, ok := EmployeeDepartmentMapReverse[EmployeeDepartmentKey(v)]; ok {
			*d = idx
			return nil
		}
	case []byte:
		strVal := string(v)
		if idx, ok := EmployeeDepartmentMapReverse[EmployeeDepartmentKey(strVal)]; ok {
			*d = idx
			return nil
		}
	}

	return errors.New("invalid employee department value")
}

func (d EmployeeDepartmentIdx) Value() (int64, error) {
	return int64(d), nil
}

type NullEmployeeDepartmentIdx struct {
	EmployeeDepartmentIdx EmployeeDepartmentIdx
	Valid                 bool
}

func (d *NullEmployeeDepartmentIdx) Scan(value interface{}) error {
	if value == nil {
		d.EmployeeDepartmentIdx, d.Valid = DepartmentUnknownIdx, false
		return nil
	}

	d.Valid = true
	return d.EmployeeDepartmentIdx.Scan(value)
}

func (d NullEmployeeDepartmentIdx) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}

	return d.EmployeeDepartmentIdx.Value()
}
