package ljconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigJson struct {
	Enabled bool
	Path    string
}

func (c *ConfigJson) InitConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
