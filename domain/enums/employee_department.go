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
	DepartmentHRIdx          EmployeeDepartmentIdx = 1
	DepartmentEngineeringIdx EmployeeDepartmentIdx = 2
	DepartmentFinanceIdx     EmployeeDepartmentIdx = 3
	DepartmentOperationsIdx  EmployeeDepartmentIdx = 4
	DepartmentSalesIdx       EmployeeDepartmentIdx = 5

	DepartmentHRKey          EmployeeDepartmentKey = "hr"
	DepartmentEngineeringKey EmployeeDepartmentKey = "engineering"
	DepartmentFinanceKey     EmployeeDepartmentKey = "finance"
	DepartmentOperationsKey  EmployeeDepartmentKey = "operations"
	DepartmentSalesKey       EmployeeDepartmentKey = "sales"
)

var (
	DepartmentHRValue = EmployeeDepartmentValue{
		Idx: int64(DepartmentHRIdx),
		Key: string(DepartmentHRKey),
		LongLabel: map[string]string{
			"id": "Sumber Daya Manusia",
			"en": "Human Resources",
		},
		ShortLabel: map[string]string{
			"id": "HRD",
			"en": "HR",
		},
	}

	DepartmentEngineeringValue = EmployeeDepartmentValue{
		Idx: int64(DepartmentEngineeringIdx),
		Key: string(DepartmentEngineeringKey),
		LongLabel: map[string]string{
			"id": "Teknologi / Engineering",
			"en": "Engineering",
		},
		ShortLabel: map[string]string{
			"id": "ENG",
			"en": "ENG",
		},
	}

	DepartmentFinanceValue = EmployeeDepartmentValue{
		Idx: int64(DepartmentFinanceIdx),
		Key: string(DepartmentFinanceKey),
		LongLabel: map[string]string{
			"id": "Keuangan",
			"en": "Finance",
		},
		ShortLabel: map[string]string{
			"id": "FIN",
			"en": "FIN",
		},
	}

	DepartmentOperationsValue = EmployeeDepartmentValue{
		Idx: int64(DepartmentOperationsIdx),
		Key: string(DepartmentOperationsKey),
		LongLabel: map[string]string{
			"id": "Operasional",
			"en": "Operations",
		},
		ShortLabel: map[string]string{
			"id": "OPS",
			"en": "OPS",
		},
	}

	DepartmentSalesValue = EmployeeDepartmentValue{
		Idx: int64(DepartmentSalesIdx),
		Key: string(DepartmentSalesKey),
		LongLabel: map[string]string{
			"id": "Penjualan / Marketing",
			"en": "Sales",
		},
		ShortLabel: map[string]string{
			"id": "SLS",
			"en": "SLS",
		},
	}
)

var (
	EmployeeDepartmentMapIdx = map[EmployeeDepartmentIdx]EmployeeDepartmentValue{
		DepartmentHRIdx:          DepartmentHRValue,
		DepartmentEngineeringIdx: DepartmentEngineeringValue,
		DepartmentFinanceIdx:     DepartmentFinanceValue,
		DepartmentOperationsIdx:  DepartmentOperationsValue,
		DepartmentSalesIdx:       DepartmentSalesValue,
	}

	EmployeeDepartmentMapReverse = map[EmployeeDepartmentKey]EmployeeDepartmentIdx{
		DepartmentHRKey:          DepartmentHRIdx,
		DepartmentEngineeringKey: DepartmentEngineeringIdx,
		DepartmentFinanceKey:     DepartmentFinanceIdx,
		DepartmentOperationsKey:  DepartmentOperationsIdx,
		DepartmentSalesKey:       DepartmentSalesIdx,
	}
)

func (d *EmployeeDepartmentIdx) Scan(value interface{}) error {
	if value == nil {
		return errors.New("department cannot be null")
	}

	switch v := value.(type) {
	case int64:
		*d = EmployeeDepartmentIdx(v)
		return nil
	case string:
		if idx, ok := EmployeeDepartmentMapReverse[EmployeeDepartmentKey(v)]; ok {
			*d = idx
			return nil
		}
	case []byte:
		if idx, ok := EmployeeDepartmentMapReverse[EmployeeDepartmentKey(string(v))]; ok {
			*d = idx
			return nil
		}
	}

	return errors.New("invalid user department value")
}

func (d EmployeeDepartmentIdx) Value() (driver.Value, error) {
	return int64(d), nil
}
