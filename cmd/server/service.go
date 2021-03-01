package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/yeboahnanaosei/edith"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct{}

func (s *service) SendText(ctx context.Context, req *edith.Request) (*edith.Response, error) {
	recipient, err := getRecipient(req.Recipient)
	if err != nil {
		log.Println("retrieving recipient name failed: ", err)
		return nil,
			status.Error(codes.InvalidArgument, "retrieving recipient name failed: "+err.Error())
	}

	err = createDestDirIfNotExists(recipient)
	if err != nil {
		log.Println("could not create dest folder:", err)
		return nil, status.Error(codes.Unknown, "an internal error occured and your text could not be delivered")
	}

	mailErr, dbErr := make(chan error), make(chan error)
	go func() { mailErr <- writeMail(recipient, req) }()
	go func() { dbErr <- saveToDB(recipient, req) }()

	select {
	case mErr := <-mailErr:
		if mErr != nil {
			log.Println("could not write mail file:", err)
			return nil, status.Error(codes.Unknown, "an internal error occured. could not deliver your text")
		}
	case dErr := <-dbErr:
		if dErr != nil {
			log.Println("could not save to database:", err)
			return nil, status.Error(codes.Unknown, "an internal error occured. could not record your text")
		}
	case <-time.After(time.Second * 5):
		log.Println("time out sending text item timed out")
		return nil, status.Error(codes.Unknown, "an internal error occured. the system took too long to deliver your text")
	}

	return &edith.Response{
			Msg: "text successfully delivered to " + strings.Title(req.Recipient),
		},
		status.Error(codes.OK, "text successfully delivered to "+strings.Title(req.Recipient))
}

func (s *service) SendFile(ctx context.Context, req *edith.Request) (*edith.Response, error) {
	recipient, err := getRecipient(req.Recipient)
	if err != nil {
		log.Println("could not get recipient name: ", err)
		s := status.New(codes.InvalidArgument, "could not determine recipient: "+err.Error())
		return nil, s.Err()
	}

	err = createDestDirIfNotExists(recipient)
	if err != nil {
		log.Println("could not create dest folder:", err)
		return nil, status.Error(codes.Unknown, "an internal error occured. could not deliver your text")
	}

	mailErr, uploadErr, dbErr := make(chan error), make(chan error), make(chan error)
	go func() { uploadErr <- uploadFile(recipient, req) }()
	go func() { mailErr <- writeMail(recipient, req) }()
	go func() { dbErr <- saveToDB(recipient, req) }()

	if <-mailErr != nil {
		log.Println("could not upload file:", mailErr)
		return nil, status.Error(codes.Unknown, "an internal error occured. could not upload your file")
	}

	if <-uploadErr != nil {
		log.Println("could not upload file:", uploadErr)
		return nil, status.Error(codes.Unknown, "an internal error occured. could not upload your file")
	}

	if <-dbErr != nil {
		log.Println("could not save to database:", dbErr)
		return nil, status.Error(codes.Unknown, "an internal error occured. could not record your file upload request")
	}

	return &edith.Response{Msg: fmt.Sprintf("file %s successfully sent to %s", req.Filename, strings.Title(req.Recipient))},
		status.Error(codes.OK, "Upload successful")
}

func (s *service) GetText(ctx context.Context, req *edith.Request) (*edith.RequestItems, error) {
	_, err := getRecipient(req.Recipient)
	if err != nil {
		log.Println("could not get your name as a recipient: ", err)
		s := status.New(codes.InvalidArgument, "could not determine your name as a recipient: "+err.Error())
		return nil, s.Err()
	}

	_, err = getRecipient(req.Sender)
	if err != nil {
		log.Println("could not get the name of the sender: ", err)
		s := status.New(codes.InvalidArgument, "could not get name of the sender: "+err.Error())
		return nil, s.Err()
	}

	requests, err := getRequestFromDB(req, 5)
	if sql.ErrNoRows == err {
		log.Println("request to retrieve text return no rows", err, req.Sender, req.Recipient)
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no mails found sent from %s to %s", req.Sender, req.Recipient))
	} else if err != nil {
		log.Println("database query failed: ", err)
		return nil, status.New(codes.Unknown, "could not complete you request. an internal error occured").Err()
	}

	r := &edith.RequestItems{Texts: requests}
	return r, nil
}

func (s *service) GetFile(ctx context.Context, req *edith.Request) (*edith.Response, error) {
	return nil, nil
}
