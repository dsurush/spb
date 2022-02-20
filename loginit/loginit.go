package loginit

import (
	"log"

	"github.com/natefinch/lumberjack"
	_ "github.com/natefinch/lumberjack"
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "../logs/foo.log",
		MaxSize:    20, // megabytes
		MaxBackups: 5,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	})
}
