// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
)

var store storage

func init() {
	store = storage{}
}

// store is a storage that stores all emails
type storage struct {
	mu sync.Mutex
}

// Put puts a given email to the storage
func (s *storage) Put(email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	// TODO: store emails

	fmt.Println("done")
	return nil
}
