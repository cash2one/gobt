package g

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/btlike/repository"
	log "github.com/inconshreveable/log15"
)

var (
	Repository repository.Repository
	Conf       Config
	Logger     log.Logger
)

type Config struct {
	Database string `json:"db"`
	LogFile  string `json:"logfile"`
	LogLevel string `json:"loglevel"`
}

func initConfig(filename string) {
	if f, err := os.Open("config/crawl.conf"); err != nil {
		panic(err)
	} else if data, err := ioutil.ReadAll(f); err != nil {
		panic(err)
	} else if err = json.Unmarshal(data, &Conf); err != nil {
		panic(err)
	}
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

	shandler := log.CallerFileHandler(log.CallerFuncHandler(handler))
	chandler := log.CallerStackHandler("%v", handler)

	handlers := log.MultiHandler(
		log.LvlFilterHandler(log.LvlCrit, chandler),
		log.LvlFilterHandler(lvl, shandler),
	)

	logger = log.New()
	logger.SetHandler(handlers)

	return Logger
}

func Init(config_file string) {
	initConfig(config_file)

	if repo, err := repository.NewMysqlRepository(Conf.Database, 256, 256); err != nil {
		panic(err)
	} else {
		Repository = repo
	}

	if logger, err := NewLogger(Conf.LogLevel, Conf.LogFile); err != nil {
		panic(nil)
	} else {
		Logger = logger
	}
}
