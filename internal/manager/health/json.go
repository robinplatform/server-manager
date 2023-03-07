package health

import (
	"encoding/json"
)

func (check *HttpHealthCheck) UnmarshalJSON(buf []byte) error {
	type alias HttpHealthCheck
	var a alias
	if err := json.Unmarshal(buf, &a); err != nil {
		return err
	}
	*check = HttpHealthCheck(a)
	return nil
}

func (check *TcpHealthCheck) UnmarshalJSON(buf []byte) error {
	type alias TcpHealthCheck
	var a alias
	if err := json.Unmarshal(buf, &a); err != nil {
		return err
	}
	*check = TcpHealthCheck(a)
	return nil
}

type healthCheck interface {
	json.Unmarshaler
	Check() bool
}

type HealthCheck struct {
	Type string `json:"type"`
	Checker healthCheck `json:"checker"`
}

func (check *HealthCheck) Check() bool {
	return check.Checker.Check()
}

func (check *HealthCheck) UnmarshalJSON(buf []byte) error {
	var checkWithType struct{Type string `json:"type"`}
	if err := json.Unmarshal(buf, &checkWithType); err != nil {
		return err
	}

	checker := HealthCheck{
		Type: checkWithType.Type,
	}

	switch checkWithType.Type {
	case "http":
		checker.Checker = &HttpHealthCheck{}
	case "tcp":
		checker.Checker = &TcpHealthCheck{}
	}

	err := json.Unmarshal(buf, &checker.Checker)
	if err != nil {
		return err
	}

	*check = checker
	return nil
}
