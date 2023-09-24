package collectors

import (
	"encoding/json"
	"fmt"
	"strconv"
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

func parseUint64(value string) uint64 {
	result, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		panic(err)
	}

	return result
}
