package rimworld

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ModsConfigFilename = "ModsConfig.xml"
)

type ModsConfigData struct {
	XMLName    xml.Name `xml:"ModsConfigData"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version"`
	ActiveMods struct {
		Text string   `xml:",chardata"`
		Mod  []string `xml:"li"`
	} `xml:"activeMods"`
	KnownExpansions struct {
		Text      string `xml:",chardata"`
		Expansion string `xml:"li"`
	} `xml:"knownExpansions"`

	// an open file handle
	fh *os.File `xml:"-"`
}

func LoadModsConfig(configFile string) (mcd *ModsConfigData, err error) {
	if fh, e := os.OpenFile(configFile, os.O_RDWR, 0644); e != nil {
		err = fmt.Errorf("unable to open %q for read/write %w", configFile, e)
	} else if bytes, e := ioutil.ReadAll(fh); e != nil {
		err = fmt.Errorf("unable to read %q %w", configFile, e)
	} else {
		mcd = &ModsConfigData{}
		if e := xml.Unmarshal(bytes, &mcd); e != nil {
			err = fmt.Errorf("unable to parse %q %w", configFile, e)
		}
		// save for later :)
		mcd.fh = fh
	}

	return
}

func (m *ModsConfigData) Save() {
	// TODO: actually save :)
}