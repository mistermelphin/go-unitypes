package unitypes

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Int int

func (i *Int) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*i = Int(vv)
	return nil
}

type Int8 int8

func (i *Int8) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*i = Int8(vv)
	return nil
}

type Int16 int16

func (i *Int16) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*i = Int16(vv)
	return nil
}

type Int32 int32

func (i *Int32) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*i = Int32(vv)
	return nil
}

type Int64 int64

func (i *Int64) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*i = Int64(vv)
	return nil
}

type Float32 float32

func (f *Float32) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*f = Float32(vv)
	return nil
}

type Float64 float32

func (f *Float64) UnmarshalJSON(bytes []byte) error {
	vv, err := parseAndConvertJsonToFloat(bytes)
	if err != nil {
		return err
	}
	*f = Float64(vv)
	return nil
}

var float64type = reflect.TypeOf(float64(0))

func parseAndConvertJsonToFloat(bytes []byte) (float64, error) {
	var v any
	if err := json.Unmarshal(bytes, &v); err != nil {
		return 0, err
	}
	return convertToFloat(v)
}

func convertToFloat(v any) (float64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		vv, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return 0, fmt.Errorf("string is not representing a number")
		}
		return vv, nil
	}
	if t.ConvertibleTo(float64type) {
		return reflect.ValueOf(v).Convert(float64type).Float(), nil
	}
	return 0, fmt.Errorf("unsupported type of number")
}
