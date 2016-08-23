package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/shiyanhui/dht"
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

func HandleMetadata(infohash []byte, ip string, port int, mi []byte) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	metadata, err := dht.Decode(mi)
	if err != nil {
		continue
	}
	info := metadata.(map[string]interface{})

	if _, ok := info["name"]; !ok {
		continue
	}

	bt := bitTorrent{
		InfoHash: hex.EncodeToString(infohash),
		Name:     info["name"].(string),
	}

	if v, ok := info["files"]; ok {
		files := v.([]interface{})
		bt.Files = make([]file, len(files))

		for i, item := range files {
			f := item.(map[string]interface{})
			bt.Files[i] = file{
				Path:   f["path"].([]interface{}),
				Length: f["length"].(int),
			}
		}
	} else if _, ok := info["length"]; ok {
		bt.Length = info["length"].(int)
	}

	data, err := json.Marshal(bt)
	if err == nil {
		fmt.Printf("%s\n\n", data)
	}
}

func Start(config *dht.Config, w *dht.Wire) {
	go func() {
		for resp := range w.Response() {
			//HandleMetadata(resp.InfoHash, resp.IP, resp.Port, resp.MetadataInfo)
			storeTorrent(resp.InfoHash, resp.MetadataInfo)
		}
	}()
	go w.Run()

	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		hash := []byte(infohash)

		if !checkTorrent(hash) {
			w.Request(hash, ip, port)
		}
	}

	d := dht.New(config)
	d.Run()
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	config := dht.NewCrawlConfig()
	w := dht.NewWire(65536, 1024, 1024)

	Start(config, w)
}
