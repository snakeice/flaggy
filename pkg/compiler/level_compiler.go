package compiler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/snakeice/flaggy/pkg/level"

	"github.com/snakeice/flaggy/pkg/models"
	"github.com/snakeice/flaggy/pkg/utils"
	yaml "gopkg.in/yaml.v2"
)

func CompileLevel(basePath string, config models.LevelsSource) {
	fmt.Println("Compiling levels...")
	levelPaths := getAllFiles(path.Join(basePath, config.Path))

	levels := map[string]*level.LevelData{}

	for _, levelPath := range levelPaths {
		fmt.Printf("Parsing %s\n", levelPath)
		data := loadLevel(levelPath)
		levels[data.Name] = data
	}

	outFile := path.Join(basePath, config.Out)

	fmt.Println("Compressing...")
	err := writeGob(levels, outFile)
	utils.ErrCheck(err)

	var fullOutFile string
	fullOutFile, err = filepath.Abs(outFile)
	if err != nil {
		fullOutFile = outFile

	}
	fmt.Printf("Out file %s\nLevel done!\n", fullOutFile)
}

func loadLevel(levelPath string) *level.LevelData {
	data := &level.LevelData{}
	b, err := ioutil.ReadFile(levelPath)
	utils.ErrCheck(err)

	err = yaml.Unmarshal(b, data)
	utils.ErrCheck(err)

	err = validateLevel(data)
	utils.ErrCheck(err)

	return data
}

func validateLevel(level *level.LevelData) error {
	if level.Name == "" {
		return errors.New("Missing level label")
	}
	if level.Properties.PropInt("width", 0) < 1 ||
		level.Properties.PropInt("height", 0) < 1 {
		return errors.New("Impossible level dimensions")
	}

	return nil
}
