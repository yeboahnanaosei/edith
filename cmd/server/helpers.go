package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yeboahnanaosei/edith"
)

type config struct {
	Recipients map[string]string `json:"recipients,omitempty"`
}

// getRecipient gets the known name of the recipient of the request
func getRecipient(name string) (string, error) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", fmt.Errorf("could not read config file: %s", err)
	}

	c := config{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return "", fmt.Errorf("could not load recipients from config file: %s", err)
	}

	recipient, ok := c.Recipients[name]
	if !ok {
		recipients := []string{}
		for r := range c.Recipients {
			recipients = append(recipients, r)
		}
		return "", fmt.Errorf("unknown recipient name '%s', expected one of %v", name, recipients)
	}

	return recipient, nil
}

// createDestDirIfNotExists creates the destination folder for the recipient
// if it does not exist
func createDestDirIfNotExists(recipient string) error {
	recipient = strings.Title(recipient)
	destDir := filepath.Join(workingDir, recipient)
	_, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("could not stat recipients folder path %s: %v", destDir, err)
	}

	return nil
}

// saveToDB saves a copy of the mail to a database.
// This is to make it easy for retrieval
func saveToDB(recipient string, req *edith.Request) error {
	db, err := sql.Open("sqlite3", filepath.Join(serverRoot, "edithd.db"))
	if err != nil {
		return err
	}
	defer db.Close()

	accra, err := time.LoadLocation("Africa/Accra")
	if err != nil {
		return fmt.Errorf("error: could not load timezone: %v", err)
	}

	var payload string
	switch strings.ToLower(req.Type) {
	case "text":
		payload = string(req.Body)
	case "file":
		payload = path.Base(req.Filename)
	default:
		return fmt.Errorf("saveToDB: unknown request item type")
	}

	timestamp := time.Now().In(accra).Unix()
	_, err = db.Exec(
		"INSERT INTO mails (sender, recipient, payload, item_sent, date_sent) VALUES (?, ?, ?, ?, ?)",
		strings.ToLower(req.Sender),
		strings.ToLower(recipient),
		payload,
		req.Type,
		timestamp,
	)
	if err != nil {
		return fmt.Errorf("database query failed: %v", err)
	}

	return nil
}

// writeMail writes the text to the recipient's edith_mail file
func writeMail(recipient string, req *edith.Request) error {
	accra, err := time.LoadLocation("Africa/Accra")
	if err != nil {
		return fmt.Errorf("error: could not load timezone: %v", err)
	}

	timestamp := time.Now().In(accra).Format("02 Jan 2006, 03:04PM")
	mailFile, err := os.OpenFile(
		filepath.Join(workingDir, recipient, "edith_mail.txt"),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		os.ModePerm,
	)
	if err != nil {
		return fmt.Errorf("could open mail file for %s: %v", recipient, err)
	}
	defer mailFile.Close()

	t := `-------------------------------------
Date: {{.Date}}
From: {{.From}}
Item: {{.Type}}

{{.Content}}

`
	tmpl, err := template.New("").Parse(t)
	if err != nil {
		return fmt.Errorf("could not parse template files: %v", err)
	}

	var content string
	switch strings.ToLower(req.Type) {
	case "text":
		content = string(req.Body)
	case "file":
		content = path.Base(string(req.Filename))
	}

	err = tmpl.Execute(
		mailFile,
		map[string]string{
			"Date":    timestamp,
			"From":    strings.Title(req.Sender),
			"Content": content,
			"Type":    req.Type,
		},
	)
	if err != nil {
		return fmt.Errorf("could not execute template: %v", err)
	}

	return nil
}

// uploadFile uploads a file from a sender to a recipient
func uploadFile(recipient string, req *edith.Request) error {
	basename := filepath.Base(strings.TrimPrefix(req.Filename, "."))
	if basename == "" || len(basename) == 0 {
		return fmt.Errorf("supplied file name appears to be empty %s. basename: %s", req.Filename, basename)
	}

	uploadedFile, err := os.OpenFile(
		filepath.Join(workingDir, recipient, basename),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		os.ModePerm,
	)
	if err != nil {
		return fmt.Errorf("could open mail file for %s: %v", recipient, err)
	}
	defer uploadedFile.Close()

	_, err = io.Copy(uploadedFile, bytes.NewReader(req.Body))
	if err != nil {
		return fmt.Errorf("could not upload file %s: %v", req.Filename, err)
	}

	return nil
}

// getRequestFromDB saves a copy of the mail to a database.
// This is to make it easy for retrieval
func getRequestFromDB(req *edith.Request, limit int) ([]*edith.Request, error) {
	db, err := sql.Open("sqlite3", filepath.Join(serverRoot, "edithd.db"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		fmt.Sprintf("SELECT sender, recipient, item_sent, payload FROM mails WHERE sender = ? AND recipient = ? AND item_sent = ? ORDER BY date_sent DESC LIMIT %d", limit),
		strings.ToLower(req.Sender), strings.ToLower(req.Recipient), strings.ToLower(req.Type),
	)

	
	if err != nil {
		return nil, err
	}

	requests := []*edith.Request{}
	for rows.Next() {
		r := &edith.Request{}
		rows.Scan(&r.Sender, &r.Recipient, &r.Type, &r.Body)
		requests = append(requests, r)
	}
	return requests, nil
}
