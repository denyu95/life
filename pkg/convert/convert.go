package convert

import (
	"errors"
)

func ToString(v interface{}) (string, error) {
	if result, ok := v.(string); ok {
		return result, nil
	} else {
		return "", errors.New("convert string err")
	}
}

func ToInt(v interface{}) (int, error) {
	if result, ok := v.(int); ok {
		return result, nil
	} else {
		return 0, errors.New("convert int err ")
	}
}

func ToInt64(v interface{}) (int64, error) {
	if result, ok := v.(int64); ok {
		return result, nil
	} else {
		return 0, errors.New("convert int64 err")
	}
}

func ToInt32(v interface{}) (int32, error) {
	if result, ok := v.(int32); ok {
		return result, nil
	} else {
		return 0, errors.New("convert int32 err")
	}
}

func ToBool(v interface{}) (bool, error) {
	if result, ok := v.(bool); ok {
		return result, nil
	} else {
		return false, errors.New("convert bool err")
	}
}
