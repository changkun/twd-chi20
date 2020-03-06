// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

type testdata struct {
	email string
	valid bool
}

func TestValidateEmail(t *testing.T) {
	emails := []testdata{
		testdata{email: "hi@changkun.us", valid: true},
	}

	// TODO: do fuzz tests

	for _, tt := range emails {
		got := ValidateEmail(tt.email)
		if !reflect.DeepEqual(got, tt.valid) {
			t.Fatalf("email [%s] cannot be validated email, want %v, got %v.", tt.email, tt.valid, got)
		}
	}
}
