package convertor

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func ToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case int32:
		return strconv.FormatInt(v.(int64), 10)
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatFloat(v.(float64), 'f', 5, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', 5, 64)
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
	switch v.(type) {
	case int:
		strV := ToString(v)
		return strconv.Atoi(strV)
	default:
		errText := fmt.Sprintf("%v类型无法使用ToInt转化", reflect.TypeOf(v))
		logrus.Warnf(errText)
		return 0, errors.New(errText)
	}
}

func ToInt64(v interface{}) (int64, error) {
	switch v.(type) {
	case int64:
		strV := ToString(v)
		return strconv.ParseInt(strV, 10, 64)
	default:
		errText := fmt.Sprintf("%v类型无法使用ToInt64转化", reflect.TypeOf(v))
		logrus.Warnf(errText)
		return 0, errors.New(errText)
	}
}

func ToBool(v interface{}) (bool, error) {
	switch v.(type) {
	case int64:
		strV := ToString(v)
		return strconv.ParseBool(strV)
	default:
		errText := fmt.Sprintf("%v类型无法使用ToInt64转化", reflect.TypeOf(v))
		logrus.Warnf(errText)
		return false, errors.New(errText)
	}
}
