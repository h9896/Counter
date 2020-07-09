package database

import (
	"errors"
	"log"
	"sync"
	"time"
)

type (
	// Counter - contain locker, and number map
	Counter struct {
		number   map[string]int
		interval int
		rw       sync.RWMutex
	}
)

// NewCounter - Initialize Counter
func NewCounter(interval int) *Counter {
	return &Counter{number: make(map[string]int), interval: interval}
}

// GetPermission - to know if the IP has the access or not
func (c *Counter) GetPermission(ip string, limit int) (result bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
			log.Fatal(err)
			result = false
		}
	}()
	defer c.rw.Unlock()
	c.rw.Lock()
	if _, ok := c.number[ip]; ok {
		c.number[ip]++
		result = c.number[ip] <= limit
		return result, err
	}
	c.number[ip] = 1
	go c.resetTimer(ip)
	result = true
	return result, err
}

// GetNumber - get the access number of the IP
func (c *Counter) GetNumber(ip string) int {
	defer c.rw.RUnlock()
	c.rw.RLock()
	if val, ok := c.number[ip]; ok {
		return val
	}
	return 0
}

// GetAllNumber - get the access number of all IP in the map
func (c *Counter) GetAllNumber() map[string]int {
	return c.number
}

func (c *Counter) resetTimer(ip string) {
	interval := time.Duration(time.Second * time.Duration(c.interval))
	t := time.NewTicker(interval)
	sapce := 0
	defer t.Stop()
	for {
		if sapce > 3 {
			c.rw.Lock()
			delete(c.number, ip)
			c.rw.Unlock()
			break
		}
		<-t.C
		c.rw.Lock()
		if c.number[ip] == 0 {
			sapce++
		} else {
			c.number[ip] = 0
		}
		c.rw.Unlock()
	}
}
