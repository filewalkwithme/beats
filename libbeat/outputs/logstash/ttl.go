package logstash

import (
	"time"
)

type ttl struct {
	duration time.Duration
	expired  bool
}

func (t *ttl) init() {
	t.expired = false
	if t.duration > 0 {
		go func(ttl *ttl) {
			<-time.NewTimer(ttl.duration).C
			ttl.expired = true
		}(t)
	}
}
