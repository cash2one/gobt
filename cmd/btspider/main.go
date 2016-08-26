package main

import (
	"encoding/hex"
	"os"
	"strings"

	"github.com/shiyanhui/dht"
	"github.com/xgfone/go-utils/log"
	"github.com/xgfone/gobt/g"
	"github.com/xgfone/gobt/store"
)

func Start(config *dht.Config, w *dht.Wire) {
	go func() {
		for resp := range w.Response() {
			go func(resp dht.Response) {
				var err error
				hash := hex.EncodeToString(resp.InfoHash)
				//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
				if err = store.StoreTorrent(hash, resp.MetadataInfo); err != nil {
					log.Errorj("Failed to store the torrent", "infohash", hash, "err", err)
				} else {
					log.Infoj("Successfully store the torrent", "infohash", hash)
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
				log.Warnj("infohash is invalid", "infohash", hash)
				return
			} else {
				hash = strings.ToLower(hash)
			}
			log.Infoj("OnAnnouncePeer", "infohash", hash, "ip", ip, "port", port)

			if !store.CheckTorrent(hash) {
				w.Request(infoHash, ip, port)
			} else {
				store.IncreaseResourceHeat(hash)
			}
		}(infohash, ip, port)
	}

	d := dht.New(config)

	log.Infoj("Start Bt Service")
	d.Run()
}

func main() {
	g.Init(os.Args[1])

	config := dht.NewCrawlConfig()
	w := dht.NewWire(65536, 1024, 1024)

	Start(config, w)
}
