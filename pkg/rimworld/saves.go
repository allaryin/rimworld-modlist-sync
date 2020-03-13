package rimworld

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

const (
	SaveGameExtension = ".rws"
)

type Saves struct {
	files []string
}

func ScanSaves(saveDir string) (s *Saves, err error) {
	if saves, e := ioutil.ReadDir(saveDir); e != nil {
		err = fmt.Errorf("unable to scan saves dir %w", e)
	} else {
		s = &Saves{
			files: make([]string, 0),
		}
		for _, file := range saves {
			if file.IsDir() {
				continue
			}
			fname := file.Name()
			if ext := filepath.Ext(fname); ext == SaveGameExtension {
				fname = strings.TrimSuffix(fname, ext)
				s.files = append(s.files, fname)
			}
		}
		sort.Strings(s.files)
	}

	return
}

func (s *Saves) ListAll() []string {
	return s.files
}
