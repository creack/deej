package deej

import (
	"strings"

	"go.uber.org/zap"
)

// Session represents a single addressable audio session
type Session interface {
	GetVolume() float32
	SetVolume(float32) error

	// TODO: Future mute support:
	// GetMute() bool
	// SetMute(m bool) error

	Key() string
	Release()
}

const (
	// Ideally these would share a common ground in baseSession
	// but it will not call the child GetVolume correctly :/.
	sessionCreationLogMessage = "Created audio session instance"

	// Format this with s.humanReadableDesc and whatever the current volume is.
	sessionStringFormat = "<session: %s, vol: %.2f>"
)

type baseSession struct {
	logger *zap.SugaredLogger
	system bool
	master bool

	// Used by Key(), needs to be set by child.
	name string

	// Used by String(), needs to be set by child.
	humanReadableDesc string
}

func (s *baseSession) Key() string {
	if s.system {
		return systemSessionName
	}
	return strings.ToLower(s.name)
}
