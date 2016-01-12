package main

import (
	"errors"
	"regexp"
	"time"
)

var (
	roleArnRegex *regexp.Regexp = regexp.MustCompile(`^arn:aws:iam::(\d+):role/([^:]+/)?([^:]+?)$`)
)

type RoleArn struct {
	value     string
	path      string
	name      string
	accountId string
}

func NewRoleArn(value string) (RoleArn, error) {
	result := roleArnRegex.FindStringSubmatch(value)

	if result == nil {
		return RoleArn{}, errors.New("invalid role ARN")
	}

	return RoleArn{value, "/" + result[2], result[3], result[1]}, nil
}

func (t RoleArn) RoleName() string {
	return t.name
}

func (t RoleArn) Path() string {
	return t.path
}

func (t RoleArn) AccountId() string {
	return t.accountId
}

func (t RoleArn) String() string {
	return t.value
}

func (t RoleArn) Empty() bool {
	return len(t.value) == 0
}

func (t RoleArn) Equals(other RoleArn) bool {
	return t.value == other.value
}

type RoleCredentials struct {
	AccessKey  string
	SecretKey  string
	Token      string
	Expiration time.Time
}

func (t *RoleCredentials) ExpiredNow() bool {
	return t.ExpiredAt(time.Now())
}

func (t *RoleCredentials) ExpiredAt(at time.Time) bool {
	return at.After(t.Expiration)
}
