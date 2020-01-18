package dorm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsMeta struct {
	session *session.Session
}

func newAwsMeta(cfg *aws.Config) *awsMeta {
	// Try to load default AWS credentials
	sesh, _ := session.NewSession(cfg)

	return &awsMeta{
		session: sesh,
	}
}
