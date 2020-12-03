package global

import (
	"bytes"
	"errors"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)
var Logger *logrus.Logger


type textFormatter struct {}

func (f *textFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b = &bytes.Buffer{}

	// 按照日志规范输出
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.000"))
	b.WriteString(" ")
	b.WriteString("[" + strings.ToUpper(entry.Level.String()) + "]")
	b.WriteString(" ")
	b.WriteString(entry.Message)
	b.WriteString(" ")

	b.WriteRune('\n')
	return b.Bytes(), nil
}


func InitLog(l *log) error {
	path := l.Dir
	if path == "" {
		return errors.New("The log path is empty")
	}

	Logger = logrus.New()
	Logger.SetLevel(logrus.InfoLevel)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	Logger.Formatter = new(textFormatter)
	filename := "%Y-%m-%d.log"
	logFileName := fmt.Sprintf("%s/%s", path, filename)
	rl, err := rotatelogs.New(logFileName)
	if nil != err {
		return err
	}
	Logger.SetOutput(rl)
	return nil
}
