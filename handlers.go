// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// EmailRequest contains the request info for email requests
type EmailRequest struct {
	Email string `json:"email"`
}

func handleEmail(c *gin.Context) {
	var j EmailRequest

	// 1. bind input data and checking email format
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse post data",
			"error":   err.Error(),
		})
		return
	}
	if err := ValidateFormat(j.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email format",
			"error":   err.Error(),
		})
		return
	}

	// do all slow stuff in a goroutine which makes the request even longer
	// I don't think we need a goroutine pool here since we are doing small 
	// project here :)
	go func() {
		// 2. verify email address
		err := ValidateHost(j.Email)
		if err != nil {
			logrus.Errorf("email [%s] cannot be validated, err: %v", j.Email, err)
			return
		}

		// 3. store email address
		err = store.Put(j.Email)
		if err != nil {
			logrus.Errorf("email [%s] cannot be stored, err: %v", j.Email, err)
			return
		}

		// 4. send an email
		err = SendEmail(j.Email)
		if err != nil {
			logrus.Errorf("email [%s] cannot receive messages, err: %v", j.Email, err)
			return
		}

		logrus.Infof("email [%s] is stored and mailed.", j.Email)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "ok", "error": nil})
}

// SendEmail sends an email with given template from configuration file
func SendEmail(email string) error {

	// TODO: send email
	return nil
}
