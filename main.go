package main

import (
	"encoding/hex"
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
		var err error
		for resp := range w.Response() {
			hash := hex.EncodeToString(resp.InfoHash)
			//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
			if err = storeTorrent(hash, resp.MetadataInfo); err != nil {
				logger.Errorf("Failed to store the torrent[%v]: %v", hash, err)
			}
		}
	}()
	go w.Run()

	config.OnAnnouncePeer = func(infohash, ip string, port int) {
		infoHash := []byte(infohash)

		hash := hex.EncodeToString(infoHash)
		logger.Infof("Announce %v on %v:%v", hash, ip, port)
		if !checkTorrent(hash) {
			w.Request(infoHash, ip, port)
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
