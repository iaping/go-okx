package common

import (
	"fmt"
	"time"
)

type Signature struct {
	Key, Timestamp, Method, Path, Body string
	IsUnix                             bool
}

// The Base64-encoded signature (see Signing Messages subsection for details).
func (s *Signature) Build() string {
	if s.Timestamp == "" {
		s.setNowTimestamp()
	}
	data := s.Timestamp + s.Method + s.Path + s.Body
	return HmacSHA256([]byte(data), []byte(s.Key))
}

// The timestamp of your request.e.g : 2020-12-08T09:08:57.715Z
func (s *Signature) setNowTimestamp() {
	if s.IsUnix {
		s.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
		return
	}
	s.Timestamp = time.Now().UTC().Format(time.RFC3339)
}
