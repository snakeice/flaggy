package compiler

import (
	"compress/gzip"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/snakeice/flaggy/pkg/utils"
)

func getAllFiles(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		ext := filepath.Ext(file)
		if ext == ".yaml" || ext == ".yml" {
			fullPath, err := filepath.Abs(file)
			utils.ErrCheck(err)
			files = append(files, fullPath)
		}
		return nil
	})
	utils.ErrCheck(err)

	return files
}

func writeGob(data interface{}, outPath string) error {
	outPath, err := filepath.Abs(outPath)
	if err != nil {
		return err
	}

	err = ensureNewFile(outPath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(outPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	gz, err := gzip.NewWriterLevel(file, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer gz.Close()

	enc := gob.NewEncoder(gz)

	return enc.Encode(data)
}

func ReadGob(data interface{}, outPath string) error {
	outPath, err := filepath.Abs(outPath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(outPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	gz, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	defer gz.Close()

	enc := gob.NewDecoder(gz)
	if err != nil {
		return err
	}

	return enc.Decode(data)
}

func ensureNewFile(path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	utils.ErrCheck(err)

	_, err = os.Stat(path)

	//Remove old
	if err == nil {
		err := os.Remove(path)
		if err != nil {
			return err
		}

	}

	file, err := os.Create(path)
	utils.ErrCheck(err)
	defer file.Close()

	return nil
}

func makeErr(format string, a ...interface{}) error {
	return errors.New(
		fmt.Sprintf(format, a...),
	)
}
