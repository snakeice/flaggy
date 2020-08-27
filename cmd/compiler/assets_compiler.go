package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/snakeice/flaggy/pkg/compiler"

	"github.com/snakeice/flaggy/pkg/models"
	"github.com/snakeice/flaggy/pkg/utils"

	yaml "gopkg.in/yaml.v2"
)

func die(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func loadConfig(configPath string) *models.SourceConfig {
	srcConfig := &models.SourceConfig{}

	configPath = path.Join(configPath, "config.yaml")

	if inf, err := os.Open(configPath); err != nil {
		die("%s: open: %v", configPath, err)
	} else {
		b, err := ioutil.ReadAll(inf)
		utils.ErrCheck(err)
		err = yaml.Unmarshal(b, srcConfig)
		utils.ErrCheck(err)
	}

	return srcConfig
}

func main() {
	var configPath string = "."
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	configPath = path.Base(configPath)

	srcConfig := loadConfig(configPath)

	dataConfig := models.DataFiles{}

	dataConfig.Sprites = compiler.CompileSprites(configPath, srcConfig.Sprites)
	dataConfig.Levels = compiler.CompileLevel(configPath, srcConfig.Levels)
	fmt.Printf("\n\n%v\n\n", dataConfig)

}
