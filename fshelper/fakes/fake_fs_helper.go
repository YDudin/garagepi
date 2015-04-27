// This file was generated by counterfeiter
package fakes

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/robdimsdale/garagepi/fshelper"
)

type FakeFsHelper struct {
	GetStaticFileSystemStub        func() (http.FileSystem, error)
	getStaticFileSystemMutex       sync.RWMutex
	getStaticFileSystemArgsForCall []struct{}
	getStaticFileSystemReturns struct {
		result1 http.FileSystem
		result2 error
	}
	GetHomepageTemplateStub        func() (*template.Template, error)
	getHomepageTemplateMutex       sync.RWMutex
	getHomepageTemplateArgsForCall []struct{}
	getHomepageTemplateReturns struct {
		result1 *template.Template
		result2 error
	}
}

func (fake *FakeFsHelper) GetStaticFileSystem() (http.FileSystem, error) {
	fake.getStaticFileSystemMutex.Lock()
	fake.getStaticFileSystemArgsForCall = append(fake.getStaticFileSystemArgsForCall, struct{}{})
	fake.getStaticFileSystemMutex.Unlock()
	if fake.GetStaticFileSystemStub != nil {
		return fake.GetStaticFileSystemStub()
	} else {
		return fake.getStaticFileSystemReturns.result1, fake.getStaticFileSystemReturns.result2
	}
}

func (fake *FakeFsHelper) GetStaticFileSystemCallCount() int {
	fake.getStaticFileSystemMutex.RLock()
	defer fake.getStaticFileSystemMutex.RUnlock()
	return len(fake.getStaticFileSystemArgsForCall)
}

func (fake *FakeFsHelper) GetStaticFileSystemReturns(result1 http.FileSystem, result2 error) {
	fake.GetStaticFileSystemStub = nil
	fake.getStaticFileSystemReturns = struct {
		result1 http.FileSystem
		result2 error
	}{result1, result2}
}

func (fake *FakeFsHelper) GetHomepageTemplate() (*template.Template, error) {
	fake.getHomepageTemplateMutex.Lock()
	fake.getHomepageTemplateArgsForCall = append(fake.getHomepageTemplateArgsForCall, struct{}{})
	fake.getHomepageTemplateMutex.Unlock()
	if fake.GetHomepageTemplateStub != nil {
		return fake.GetHomepageTemplateStub()
	} else {
		return fake.getHomepageTemplateReturns.result1, fake.getHomepageTemplateReturns.result2
	}
}

func (fake *FakeFsHelper) GetHomepageTemplateCallCount() int {
	fake.getHomepageTemplateMutex.RLock()
	defer fake.getHomepageTemplateMutex.RUnlock()
	return len(fake.getHomepageTemplateArgsForCall)
}

func (fake *FakeFsHelper) GetHomepageTemplateReturns(result1 *template.Template, result2 error) {
	fake.GetHomepageTemplateStub = nil
	fake.getHomepageTemplateReturns = struct {
		result1 *template.Template
		result2 error
	}{result1, result2}
}

var _ fshelper.FsHelper = new(FakeFsHelper)