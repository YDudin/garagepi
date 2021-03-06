// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/robdimsdale/garagepi/os"
)

type FakeOsHelper struct {
	SleepStub        func(d time.Duration)
	sleepMutex       sync.RWMutex
	sleepArgsForCall []struct {
		d time.Duration
	}
}

func (fake *FakeOsHelper) Sleep(d time.Duration) {
	fake.sleepMutex.Lock()
	fake.sleepArgsForCall = append(fake.sleepArgsForCall, struct {
		d time.Duration
	}{d})
	fake.sleepMutex.Unlock()
	if fake.SleepStub != nil {
		fake.SleepStub(d)
	}
}

func (fake *FakeOsHelper) SleepCallCount() int {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return len(fake.sleepArgsForCall)
}

func (fake *FakeOsHelper) SleepArgsForCall(i int) time.Duration {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.sleepArgsForCall[i].d
}

var _ os.OSHelper = new(FakeOsHelper)
