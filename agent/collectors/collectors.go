package collectors

import (
	"encoding/json"
	"fmt"
)

type Collector interface {
	GetName() string

	GetEndpoint() string

	Collect(uuid string) (interface{}, error)
}

func PrettyPrintJson(value any) error {
	data, err := json.MarshalIndent(value, "", " ")

	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
