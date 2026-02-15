package enums

import "errors"

type EmployeeStatus int64

const (
	Inactive EmployeeStatus = 1
	Active   EmployeeStatus = 2
)

var (
	EmployeeStatusMap = map[EmployeeStatus]string{
		Active:   "Aktif",
		Inactive: "Tidak Aktif",
	}

	EmployeeStatusMapReverse = map[string]EmployeeStatus{
		"Aktif":       Active,
		"Tidak Aktif": Inactive,
	}
)

func (s EmployeeStatus) String() string {
	if status, ok := EmployeeStatusMap[s]; ok {
		return status
	}

	return ""
}

func (s *EmployeeStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*s = EmployeeStatus(v)
	case int32:
		*s = EmployeeStatus(v)
	case string:
		for key, val := range EmployeeStatusMapReverse {
			if key == v {
				*s = val
				return nil
			}
		}

		return errors.New("invalid user status string")
	case []byte:
		for key, val := range EmployeeStatusMapReverse {
			if string(v) == key {
				*s = val
				return nil
			}
		}

		return errors.New("invalid user status []byte")
	}

	return nil
}

func (s EmployeeStatus) Value() (interface{}, error) {
	return int64(s), nil
}
