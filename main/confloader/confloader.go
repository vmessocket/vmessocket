package confloader

import (
	"io"
	"os"
)

var (
	EffectiveConfigFileLoader configFileLoader
	EffectiveExtConfigLoader  extconfigLoader
)

type (
	configFileLoader func(string) (io.Reader, error)
	extconfigLoader  func([]string, io.Reader) (io.Reader, error)
)

func LoadConfig(file string) (io.Reader, error) {
	if EffectiveConfigFileLoader == nil {
		newError("external config module not loaded, reading from stdin").AtInfo().WriteToLog()
		return os.Stdin, nil
	}
	return EffectiveConfigFileLoader(file)
}

func LoadExtConfig(files []string, reader io.Reader) (io.Reader, error) {
	if EffectiveExtConfigLoader == nil {
		return nil, newError("external config module not loaded").AtError()
	}
	return EffectiveExtConfigLoader(files, reader)
}
