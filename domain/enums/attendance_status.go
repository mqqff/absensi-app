package enums

import "errors"

type AttendanceStatus int64

const (
	Late   AttendanceStatus = 1
	OnTime AttendanceStatus = 2
)

var (
	AttendanceStatusMap = map[AttendanceStatus]string{
		Late:   "Terlambat",
		OnTime: "Tepat Waktu",
	}

	AttendanceStatusMapReverse = map[string]AttendanceStatus{
		"Terlambat":   Late,
		"Tepat Waktu": OnTime,
	}
)

func (s AttendanceStatus) String() string {
	if status, ok := AttendanceStatusMap[s]; ok {
		return status
	}

	return ""
}

func (s *AttendanceStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*s = AttendanceStatus(v)
	case int32:
		*s = AttendanceStatus(v)
	case string:
		for key, val := range AttendanceStatusMapReverse {
			if key == v {
				*s = val
				return nil
			}
		}

		return errors.New("invalid attendance status string")
	case []byte:
		for key, val := range AttendanceStatusMapReverse {
			if string(v) == key {
				*s = val
				return nil
			}
		}

		return errors.New("invalid attendance status []byte")
	}

	return nil
}

func (s AttendanceStatus) Value() (interface{}, error) {
	return int64(s), nil
}
