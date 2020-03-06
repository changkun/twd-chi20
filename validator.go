// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

type smtpError struct {
	Err error
}

func (e smtpError) Error() string {
	return e.Err.Error()
}

func (e smtpError) Code() string {
	return e.Err.Error()[0:3]
}

func newSMTPError(err error) smtpError {
	return smtpError{err}
}

var (
	timeout              = time.Second * 5
	errEmailBadFormat    = errors.New("invalid email format")
	errUnresolveableHost = errors.New("unresolvable email host")
	reEmail              = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}

// dialTimeout returns a new SMTP Client connected to an server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func dialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// ValidateEmail validates if an email is an valid address by checking
// its format and host avaliability.
func ValidateEmail(email string) error {
	err := ValidateFormat(email)
	if err != nil {
		return err
	}
	err = ValidateHost(email)
	if err != nil {
		return err
	}
	return nil
}

// ValidateFormat validates an email contains a correct format
func ValidateFormat(email string) error {
	if !reEmail.MatchString(email) {
		return errEmailBadFormat
	}
	return nil
}

// ValidateHost verifies if an email host does exists
func ValidateHost(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errUnresolveableHost
	}

	client, err := dialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), timeout)
	if err != nil {
		return newSMTPError(err)
	}
	defer client.Close()

	err = client.Hello("changkun.de")
	if err != nil {
		return newSMTPError(err)
	}
	err = client.Mail("hi@changkun.us")
	if err != nil {
		return newSMTPError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return newSMTPError(err)
	}
	return nil
}
