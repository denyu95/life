package convertor

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func ToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case int32:
		return strconv.FormatInt(int64(v.(int32)), 10)
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v.(bool))
	case time.Time:
		return v.(time.Time).Format("2006-01-02 15:04:05")
	case nil:
		return ""
	default:
		r := ""
		buf, err := json.Marshal(v)
		if err != nil {
			r = fmt.Sprint(v)
		} else {
			r = string(buf)
		}
		return r
	}
}

func ToInt(v interface{}) (int, error) {
	strV := ToString(v)
	return strconv.Atoi(strV)
}

func ToInt64(v interface{}) (int64, error) {
	strV := ToString(v)
	return strconv.ParseInt(strV, 10, 64)
}

func ToBool(v interface{}) (bool, error) {
	strV := ToString(v)
	return strconv.ParseBool(strV)
}

func ToFloat32(v interface{}) (float32, error) {
	strV := ToString(v)
	v64, err := strconv.ParseFloat(strV, 32)
	return float32(v64), err
}

func ToFloat64(v interface{}) (float64, error) {
	strV := ToString(v)
	return strconv.ParseFloat(strV, 64)
}
