package db

import (
	"time"
)

// GameMessage ...
type GameMessage map[string]interface{}

const (
	// ConstIdxInvalid ...
	ConstIdxInvalid = 0xff
)

// Mem is a main memory storage
type Mem struct {
	///////////////////////////////////////////////////////////////////////////
	// persistent

	year     byte
	monthIdx byte

	inChannel  <-chan GameMessage
	outChannel chan<- GameMessage
}

var mem Mem

// GetMem returns static mem
func GetMem() *Mem {
	return &mem
}

// Initialize initialize mem
func (m *Mem) Initialize(inChannel <-chan GameMessage, outChannel chan<- GameMessage) {
	m.year = 0
	m.monthIdx = 0

	m.inChannel = inChannel
	m.outChannel = outChannel
}

// SleepSeconds ...
func (m *Mem) SleepSeconds(seconds int) {
	m.sleepSec(seconds)
}

func (m *Mem) sleepSec(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}

// GetLocalized returns localized string
func (m *Mem) GetLocalized(key string) string {
	return m.getLocalized(key)
}

func (m *Mem) getLocalized(key string) string {
	return getResources().GetMessage(key)
}
