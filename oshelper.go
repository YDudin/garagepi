package garagepi

import (
	"os/exec"
	"time"
)

type OsHelper interface {
	Exec(executable string, arg ...string) (string, error)
	Sleep(d time.Duration)
}

type OsHelperImpl struct {
	logger Logger
}

func NewOsHelperImpl(
	logger Logger,
) *OsHelperImpl {
	return &OsHelperImpl{
		logger: logger,
	}
}

func (h *OsHelperImpl) Exec(executable string, arg ...string) (string, error) {
	out, err := exec.Command(executable, arg...).CombinedOutput()
	return string(out), err
}

func (h *OsHelperImpl) Sleep(d time.Duration) {
	h.logger.Log("sleeping for " + d.String())
	time.Sleep(d)
}
