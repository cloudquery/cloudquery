package client

import (
	"fmt"
	"regexp"
	"sync"
)

// this is used to hook ParseResourceGroup and to have easier codegen
var debug = false

const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`

var resourceIDPattern = regexp.MustCompile(resourceIDPatternText)

func ParseResourceGroup(resourceID string) (string, error) {
	if debug {
		return "debug", nil
	}
	match := resourceIDPattern.FindStringSubmatch(resourceID)
	if len(match) == 0 {
		return "", fmt.Errorf("parsing failed for %s. Invalid resource Id format", resourceID)
	}
	return match[2], nil
}

type syncData struct {
	data any
	err  error
	once *sync.Once
}

func loadOrStore(m *sync.Map, key string, f func() (any, error)) (any, error) {
	temp, _ := m.LoadOrStore(key, syncData{
		data: nil,
		err:  nil,
		once: &sync.Once{},
	})
	d := temp.(syncData)
	if d.data == nil && d.err == nil {
		d.once.Do(func() {
			dt, err := f()
			if err != nil {
				dt = nil
				d.once = &sync.Once{}
			}
			m.Store(key, syncData{
				data: dt,
				err:  err,
				once: d.once,
			})
		})
		temp, _ = m.Load(key)
		d = temp.(syncData)
	}

	return d.data, d.err
}
