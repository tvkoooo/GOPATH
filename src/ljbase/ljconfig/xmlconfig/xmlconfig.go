package ljconfig

import (
	"encoding/xml"
	"fmt"
	"os"
)

type ConfigXml struct {
	Enabled bool   `xml:"enabled"`
	Path    string `xml:"path"`
}

func (c *ConfigXml) InitConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := xml.NewDecoder(f).Decode(c); err != nil {
		fmt.Println("Error Decode file:", err)
		return
	}

}
