package log

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/lestrrat-go/file-rotatelogs"
	"time"
)

func NewLog(name string) hclog.Logger {
	fileName := fmt.Sprintf("logs/%s.log.%%Y%%m%%d", name)
	rl, err := rotatelogs.New(fileName, rotatelogs.WithMaxAge(120*24*time.Hour), rotatelogs.WithClock(rotatelogs.UTC))
	if err != nil {
		panic(err)
	}
	l := hclog.New(&hclog.LoggerOptions{
		Output: rl,
		Level:  hclog.Info,
		Name:   name,
	})
	return l
}
