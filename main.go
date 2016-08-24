package main

import (
	"encoding/hex"
	"os"
	"strings"

	"github.com/shiyanhui/dht"
	"github.com/xgfone/gobt/g"
	"github.com/xgfone/gobt/logger"
)

func Start(config *dht.Config, w *dht.Wire) {
	go func() {
		for resp := range w.Response() {
			go func(resp dht.Response) {
				var err error
				hash := hex.EncodeToString(resp.InfoHash)
				//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
				if err = storeTorrent(hash, resp.MetadataInfo); err != nil {
					logger.Errorf("Failed to store the torrent[%v]: %v", hash, err)
				} else {
					logger.Infof("Successfully store the torrent[%v]", hash)
				}
			}(resp)
		}
	}()
	go w.Run()

	config.OnAnnouncePeer = func(infohash, ip string, port int) {
		go func(infohash, ip string, port int) {
			infoHash := []byte(infohash)

			hash := hex.EncodeToString(infoHash)
			if len(hash) != 40 {
				logger.Warnf("infohash is invalid: %v", hash)
				return
			} else {
				hash = strings.ToLower(hash)
			}
			logger.Infof("Announce %v on %v:%v", hash, ip, port)

			if !checkTorrent(hash) {
				w.Request(infoHash, ip, port)
			} else {
				increaseResourceHeat(hash)
			}
		}(infohash, ip, port)
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
