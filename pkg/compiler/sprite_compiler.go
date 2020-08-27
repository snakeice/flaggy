package compiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/snakeice/flaggy/pkg/sprite"
	"github.com/snakeice/flaggy/pkg/utils"

	"github.com/snakeice/flaggy/pkg/models"
	yaml "gopkg.in/yaml.v2"
)

func CompileSprites(basePath string, config models.SptiresSource) string {
	fmt.Println("Compiling sprites...")
	spritePaths := getAllFiles(path.Join(basePath, config.Path))

	sprites := map[string]*sprite.SpriteData{}

	for _, spritePath := range spritePaths {
		fmt.Printf("Parsing %s\n", spritePath)
		data := loadSprite(spritePath)

		if _, contains := sprites[data.Name]; contains {
			fmt.Printf("Sprite name '%s' is duplicated!\n", data.Name)
			os.Exit(1)
		}

		sprites[data.Name] = data
	}

	outFile := path.Join(basePath, config.Out)

	fmt.Println("Compressing...")
	err := writeGob(sprites, outFile)
	utils.ErrCheck(err)

	var fullOutFile string
	fullOutFile, err = filepath.Abs(outFile)
	if err != nil {
		fullOutFile = outFile

	}
	fmt.Printf("Out file %s\nSprite done!\n", fullOutFile)
	basePath, _ = filepath.Abs(basePath)
	path, err := filepath.Rel(basePath, fullOutFile)
	utils.ErrCheck(err)

	return path
}

func loadSprite(spritePath string) *sprite.SpriteData {
	data := &sprite.SpriteData{}
	b, err := ioutil.ReadFile(spritePath)
	utils.ErrCheck(err)

	err = yaml.Unmarshal(b, data)
	utils.ErrCheck(err)

	err = validateSprite(spritePath, data)
	utils.ErrCheck(err)

	return data
}

func validateSprite(iname string, data *sprite.SpriteData) error {
	if data.Name == "" {

		return makeErr("%s: Missing name", iname)
	}

	for fri, frame := range data.Frames {
		if len(frame.Data) != data.Height {

			return makeErr("Frame %d bad lines (%d != %d)",
				fri, len(frame.Data), data.Height)
		}
		for lno, line := range frame.Data {
			if len(line) != data.Width {

				return makeErr("%s:%d Frame %d wrong len "+
					"(%d != %d)", iname, lno, fri,
					len(line), data.Width)
			}
			var _, hasDef = data.Glyphs["default"]
			for i := range line {
				if line[i] == ' ' {
					continue
				}
				ss := string(line[i])
				if _, ok := data.Glyphs[ss]; !ok && !hasDef {
					return makeErr("%s:%d Frame %d: "+
						"unknown glyph and no has default",
						iname, i, fri)
				}
			}
		}
	}
	return nil
}
