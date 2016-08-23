package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/btlike/repository"
	"github.com/shiyanhui/dht"
	"github.com/xgfone/gobt/g"
	"github.com/xgfone/gobt/logger"
)

type Files []repository.File

func (a Files) Len() int           { return len(a) }
func (a Files) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Files) Less(i, j int) bool { return a[i].Length > a[j].Length }

func storeTorrent(infohash string, data interface{}) (err error) {
	if len(infohash) != 40 {
		logger.Errorf("infohash[%v] len is not 40", infohash)
		return
	}

	defer func() {
		if e := recover(); e != nil {
			logger.Errorf("Failed to store the torrent[%v]: %v", infohash, e)
			err = fmt.Errorf("%v", e)
		}
	}()

	logger.Infof("Starting to store the torrent[%v]", infohash)

	if info, ok := data.(map[string]interface{}); ok {
		var t repository.Torrent
		t.CreateTime = time.Now()
		t.Infohash = infohash

		// get name
		if name, ok := info["name"].(string); ok {
			t.Name = name
			if t.Name == "" {
				logger.Error("store name len is 0")
				return
			}
		}

		// get files
		if v, ok := info["files"]; !ok {
			t.Length = int64(info["length"].(int))
			t.FileCount = 1
			t.Files = append(t.Files, repository.File{Name: t.Name, Length: t.Length})
		} else {
			var tmpFiles Files
			files := v.([]interface{})
			tmpFiles = make(Files, len(files))
			for i, item := range files {
				fl := item.(map[string]interface{})
				flName := fl["path"].([]interface{})
				tmpFiles[i] = repository.File{
					Name:   flName[0].(string),
					Length: int64(fl["length"].(int)),
				}
			}
			sort.Sort(tmpFiles)

			for k, v := range tmpFiles {
				if len(v.Name) > 0 {
					t.Length += v.Length
					t.FileCount++
					if k < 5 {
						t.Files = append(t.Files, repository.File{
							Name:   v.Name,
							Length: v.Length,
						})
					}
				}
			}
		}

		err = g.Repository.CreateTorrent(t)
	}

	return
}

func checkTorrent(infohash string) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("Failed to check the torrent[%v]: %v", infohash, err)
			ok = false
		}
	}()

	ok = true
	if t, err := g.Repository.GetTorrentByInfohash(infohash); err != nil || t.Infohash != infohash {
		ok = false
	}
	return
}

func HandleMetadata(infohash []byte, ip string, port int, mi []byte) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()

	metadata, err := dht.Decode(mi)
	if err != nil {
		return
	}
	info := metadata.(map[string]interface{})

	if _, ok := info["name"]; !ok {
		return
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
		logger.Infof("%s\n\n", data)
	}
}
