package logger

import (
	"fmt"
	"os"

	log "github.com/inconshreveable/log15"
)

var Logger log.Logger

func init() {
	Logger, _ = NewLogger("info", "")
}

func NewLogger(level, filepath string) (logger log.Logger, err error) {
	// Logger := log.New(os.Stderr, "app", log.LstdFlags|log.Lshortfile)

	var lvl log.Lvl
	if _level, _err := log.LvlFromString(level); _err != nil {
		err = _err
		return
	} else {
		lvl = _level
	}

	var handler log.Handler
	if filepath == "" {
		handler = log.StreamHandler(os.Stderr, log.LogfmtFormat())
	} else {
		handler, err = log.FileHandler(filepath, log.LogfmtFormat())
		if err != nil {
			return
		}
	}
	handler = log.SyncHandler(handler)

	//shandler := log.CallerFuncHandler(handler)
	shandler := log.CallerFileHandler(handler)
	chandler := log.CallerStackHandler("%v", handler)

	handlers := log.MultiHandler(
		log.LvlFilterHandler(log.LvlCrit, chandler),
		log.LvlFilterHandler(lvl, shandler),
	)

	logger = log.New()
	logger.SetHandler(handlers)

	return
}

func Debug(v ...interface{}) {
	Logger.Debug(fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	Logger.Info(fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	Logger.Warn(fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	Logger.Error(fmt.Sprint(v...))
}

func Crit(v ...interface{}) {
	Logger.Crit(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	Logger.Info(fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	Logger.Error(fmt.Sprintf(format, v...))
}

func Critf(format string, v ...interface{}) {
	Logger.Crit(fmt.Sprintf(format, v...))
}
