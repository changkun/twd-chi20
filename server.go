// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// serveFS is a middleware that allows static files serves
// in the root router like "addr:port/""
type serveFS interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func loacalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(l.root, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if !l.indexes && stats.IsDir() {
			return false
		}
		return true
	}
	return false
}

// staticServe returns a middleware handler that serves static files
// in the given directory.
func staticServe(urlPrefix string, fs serveFS) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func main() {
	// To whom it may concern:
	// I don't think we need gracefully shutdown the server at this
	// moment. If you in using this code in more production level,
	// you will need modify the following server life cycle.

	r := gin.Default()

	// register routers
	r.Use(staticServe("/", loacalFile("./public", true)))
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	v1.POST("/email", handleEmail)

	logrus.Infof("twd-chi20 is running at: http://%s", conf.Addr)
	r.Run(conf.Addr)
}
