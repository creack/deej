package util

import (
	"errors"
	"fmt"
)

const (
	MinimumSerialReadSize = 1
	DefaultEditor         = "gedit"
)

func getCurrentWindowProcessNames() ([]string, error) {
	return nil, errors.New("Not implemented")
}

func externalCommand(cmd, arg string) []string {
	return []string{"/usr/bin/env", "bash", "-c", fmt.Sprintf("%s %s", cmd, arg)}
}
