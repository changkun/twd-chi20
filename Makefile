# Copyright 2020 Changkun Ou. All rights reserved.
# Use of this source code is governed by a GPLv3
# license that can be found in the LICENSE file.

all:
	go build

s:
	./twd-chi20 -config ./conf.yaml

clean:
	rm -rf twd-chi20