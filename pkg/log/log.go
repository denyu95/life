package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"server.paperang.com/servertools/convertor"
)

var Logger *logrus.Logger

type MyFormatter struct {
	TimestampFormat  string
	DisableTimestamp bool
	DisableFileLine  bool
}

func init() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&MyFormatter{})
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	output := ""

	switch entry.Level {
	case logrus.WarnLevel:
		output += "WARNING: "
		break
	case logrus.InfoLevel:
		output += "NOTICE: "
		break
	}

	// 是否打印时间
	if !m.DisableTimestamp {
		timestampFormat := m.TimestampFormat
		if timestampFormat == "" {
			timestampFormat = "2006-01-02 15:04:05"
		}
		data = make(map[string]interface{})
		output += entry.Time.Format(timestampFormat) + " "
	}

	// 是否打印代码行数
	if !m.DisableFileLine {
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		output += "[" + fileVal + "]"
	}

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	for k, v := range data {
		strV, _ := convertor.ToString(v)
		output += " " + k + "[" + strV + "]"
	}

	output += " " + entry.Message + "\n"

	return []byte(output), nil
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(fields)
}

func WithField(key, value string) *logrus.Entry {
	return Logger.WithField(key, value)
}

func Warn(arg ...interface{}) {
	Logger.Warn(arg...)
}

func Info(arg ...interface{}) {
	Logger.Info(arg...)
}
