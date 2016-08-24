package main

import (
	"encoding/hex"
	"os"
	"strings"

	"github.com/shiyanhui/dht"
	"github.com/xgfone/gobt/g"
)

func Start(config *dht.Config, w *dht.Wire) {
	go func() {
		for resp := range w.Response() {
			go func(resp dht.Response) {
				var err error
				hash := hex.EncodeToString(resp.InfoHash)
				//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
				if err = storeTorrent(hash, resp.MetadataInfo); err != nil {
					g.Logger.Error("Failed to store the torrent", "infohash", hash, "err", err)
				} else {
					g.Logger.Info("Successfully store the torrent", "infohash", hash)
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
				g.Logger.Warn("infohash is invalid", "infohash", hash)
				return
			} else {
				hash = strings.ToLower(hash)
			}
			g.Logger.Info("OnAnnouncePeer", "infohash", hash, "ip", ip, "port", port)

			if !checkTorrent(hash) {
				w.Request(infoHash, ip, port)
			} else {
				increaseResourceHeat(hash)
			}
		}(infohash, ip, port)
	}

	d := dht.New(config)

	g.Logger.Info("Start Bt Service")
	d.Run()
}

func main() {
	g.Init(os.Args[1])

	config := dht.NewCrawlConfig()
	w := dht.NewWire(65536, 1024, 1024)

	Start(config, w)
}
