// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type config struct {
	Addr string `json:"addr"`
	Mode string `json:"mode"`
}

var conf *config

const help = `
twd-chi202 is a webservice that provides a simple web subscription service.
Usage:
`

func init() {
	c := flag.String("config", "", "path to config file")
	usage := func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(help))
		flag.PrintDefaults()
	}
	flag.Usage = usage
	flag.Parse()
	f := *c
	if len(f) == 0 {
		usage()
		os.Exit(1)
	}

	y, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read configuration file: %v", err)
		os.Exit(1)
	}

	conf = &config{}
	err = yaml.Unmarshal(y, conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse configuration file: %v", err)
		os.Exit(1)
	}

	gin.SetMode(conf.Mode)
}
