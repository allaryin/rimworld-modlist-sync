package rimworld

import "encoding/xml"

const (
	// while 100% xml, saves are stored with a "rimworld save" file extension
	SavegameExtension = ".rws"
)

type Savegame struct {
	XMLName xml.Name `xml:"savegame"`
	Text    string   `xml:",chardata"`
	Meta    struct {
		Text        string `xml:",chardata"`
		GameVersion string `xml:"gameVersion"`
		ModIds      struct {
			Text string   `xml:",chardata"`
			ModId   []string `xml:"li"`
		} `xml:"modIds"`
		ModNames struct {
			Text string   `xml:",chardata"`
			ModName   []string `xml:"li"`
		} `xml:"modNames"`
	} `xml:"meta"`

	// NB: Actual game save data lives under `game` - which we really don't want to think about
}
