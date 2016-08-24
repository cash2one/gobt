package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/btlike/repository"
	"github.com/shiyanhui/dht"
	"github.com/xgfone/gobt/g"
)

type Files []repository.File

func (a Files) Len() int           { return len(a) }
func (a Files) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Files) Less(i, j int) bool { return a[i].Length > a[j].Length }

type torrentSearch struct {
	Name       string
	Length     int64
	Heat       int64
	CreateTime time.Time
}

func storeTorrent(infohash string, metadatainfo []byte) (err error) {
	if len(infohash) != 40 {
		g.Logger.Error("infohash len is not 40", "infohash", infohash)
		return
	}

	defer func() {
		if e := recover(); e != nil {
			g.Logger.Error("Failed to store the torrent", "infohash", infohash, "err", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	g.Logger.Info("Starting to store the torrent", "infohash", infohash)

	var data interface{}
	data, err = dht.Decode(metadatainfo)
	if err != nil {
		g.Logger.Error("Failed to decode the metadata info of the torrent", "infohash", infohash)
		return
	}

	if info, ok := data.(map[string]interface{}); ok {
		var t repository.Torrent
		t.CreateTime = time.Now()
		t.Infohash = infohash

		// get name
		if name, ok := info["name"].(string); ok {
			t.Name = name
			if t.Name == "" {
				g.Logger.Error("store name len is 0", "infohash", infohash)
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

		g.Logger.Info("Start to store the torrent into Mysql.", "infohash", t.Infohash)
		err = g.Repository.CreateTorrent(t)
		if err == nil {
			data := torrentSearch{
				Name:       t.Name,
				Length:     t.Length,
				CreateTime: time.Now(),
			}
			indexType := strings.ToLower(string(t.Infohash[0]))
			g.ElasticClient.Index().Index("torrent").Type(indexType).Id(t.Infohash).BodyJson(data).Refresh(false).Do()
		}
	} else {
		g.Logger.Warn("Invalid data", "infohash", infohash)
	}

	return
}

func checkTorrent(infohash string) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			g.Logger.Error("Failed to check the torrent", "infohash", infohash, "err", err)
			ok = false
		}
	}()

	ok = true
	if t, err := g.Repository.GetTorrentByInfohash(infohash); err != nil || t.Infohash != infohash {
		ok = false
	}
	return
}

func increaseResourceHeat(key string) {
	indexType := strings.ToLower(string(key[0]))
	searchResult, err := g.ElasticClient.Get().Index("torrent").Type(indexType).Id(key).Do()
	if err == nil && searchResult != nil && searchResult.Source != nil {
		var tdata torrentSearch
		err = json.Unmarshal(*searchResult.Source, &tdata)
		if err == nil {
			tdata.Heat++
			_, err = g.ElasticClient.Index().Index("torrent").Type(indexType).Id(key).BodyJson(tdata).Refresh(false).Do()
			if err != nil {
				g.Logger.Warn("Failed to increase the heat of the torrent", "infohash", key, "err", err)
			} else {
				g.Logger.Info("Successfully increase the heat", "infohash", key)
			}
		} else {
			g.Logger.Warn("Failed to increase the heat of the torrent", "infohash", key, "err", err)
		}
	} else {
		g.Logger.Info("Don't increase the heat", "infohash", key)
	}
}
