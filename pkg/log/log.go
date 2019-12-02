package log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/denyu95/life/pkg/convertor"
)

type MyFormatter struct {
	TimestampFormat  string
	DisableTimestamp bool
	DisableFileLine  bool
}

var MapLog = make(map[string]*logrus.Entry)

// @title	Init
// @description	日志初始化动作
// @param	path			string	"需要传文件的绝对路径"
// @param	rotationTime	time	"日志分割时间"
// @param	maxAge			time	"日志保留时间"
func Init(outPath string, rotationTime, maxAge time.Duration) {
	if outPath != "" {
		logrus.AddHook(newLfsHook(outPath, rotationTime, maxAge))
		// 控制台不输出
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Open Src File err", err)
		}
		writer := bufio.NewWriter(src)
		logrus.SetOutput(writer)
	} else {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&MyFormatter{})
	}
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
		strV := convertor.ToString(v)
		output += " " + k + "[" + strV + "]"
	}

	output += " " + entry.Message + "\n"

	return []byte(output), nil
}

func newLfsHook(path string, rotationTime, maxAge time.Duration) logrus.Hook {
	tail := ""
	if rotationTime == time.Minute {
		tail = ".%Y%m%d%H%M"
	} else if rotationTime == time.Hour {
		tail = ".%Y%m%d%H"
	}

	infoWriter, err := rotatelogs.New(
		path + tail,
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(path),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(rotationTime),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(maxAge),
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)

	wfWriter, err := rotatelogs.New(
		path + ".wf" + tail,
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(path+".wf"),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(rotationTime),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(maxAge),
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	logrus.SetReportCaller(true)
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel: infoWriter,
		logrus.WarnLevel: wfWriter,
	}, &MyFormatter{})

	return lfsHook
}
