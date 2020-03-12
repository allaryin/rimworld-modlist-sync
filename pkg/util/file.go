package util

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

func GetHomeDir() (home string){
	// find our home directory
	home, err := homedir.Dir()
	if err != nil {
		// we failed to look up homedir, this won't bode well - just assume pwd
		if home, err = os.Getwd(); err != nil {
			// we don't have a pwd, the world is falling apart
			home = "."
		}
	}
	return
}

func FileExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return true
	} else {
		return !os.IsNotExist(err)
	}
}

func IsDir(dir string) bool {
	if info, _ := os.Stat(dir); info == nil {
		return false
	} else {
		return info.IsDir()
	}
}
