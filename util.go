package main

import (
	"crypto/md5"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	"net/http"
	"time"
)

func md5hash(orig string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(orig)))
}

func sleepFor(sec int64) {
	for i := 0; i < int(sec); i++ {
		if !IsRunning {
			return
		}
		time.Sleep(time.Second)
	}
}

func contentMiddleware(c martini.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Content-Type") {
	case "application/xml":
		c.MapTo(encoder.XmlEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	default:
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
}
