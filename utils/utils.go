package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func GenerateKey(vals ...string) string {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	salt := fmt.Sprintf("%f", r.Float32())
	h := sha1.New()
	h.Write([]byte(salt))
	for _, val := range vals {
		h.Write([]byte(val))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func IsValidEmail(email string) bool {
	return emailPattern.MatchString(email)
}

type PagingInfo interface {
	Offset() int
	Limit() int
}

type Paging struct {
	offset int
	limit  int
}

func NewPaging(offset, limit int) *Paging {
	return &Paging{offset: offset, limit: limit}
}

func (p *Paging) Offset() int {
	return p.offset
}

func (p *Paging) Limit() int {
	return p.limit
}
