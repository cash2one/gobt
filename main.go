package main

import (
	"os"

	"github.com/shiyanhui/dht"
	"github.com/xgfone/gobt/g"
	"github.com/xgfone/gobt/logger"
)

type file struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type bitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Files    []file `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
}

func Start(config *dht.Config, w *dht.Wire) {
	go func() {
		for resp := range w.Response() {
			//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
			storeTorrent(resp.InfoHash, resp.MetadataInfo)
		}
	}()
	go w.Run()

	config.OnAnnouncePeer = func(infohash, ip string, port int) {
		logger.Infof("Announce %v on %v:%v", infohash, ip, port)

		hash := []byte(infohash)

		if !checkTorrent(hash) {
			w.Request(hash, ip, port)
		}
	}

	d := dht.New(config)

	logger.Info("Start Bt ...")
	d.Run()
}

func main() {
	g.Init(os.Args[1])

	config := dht.NewCrawlConfig()
	w := dht.NewWire(65536, 1024, 1024)

	Start(config, w)
}
