package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/go-playground/validator.v9"
)

type ConvertService struct {
	Path      string `validate:"required"`
	BeforeExt string `validate:"oneof=.png .jpeg .jpg .gif"`
	AfterExt  string `validate:"oneof=.png .jpeg .jpg .gif"`
}

func NewConvertService(path string, beforeExt string, afterExt string) *ConvertService {
	cs := &ConvertService{
		Path:      path,
		BeforeExt: "." + beforeExt,
		AfterExt:  "." + afterExt,
	}

	validate := validator.New()
	errors := validate.Struct(cs)
	if errors != nil {
		log.Fatal(errors)
		return nil
	}
	return cs
}

// convert files extension to another extension
func (cs *ConvertService) Convert() error {
	paths := cs.dirwalk(cs.Path)
	for _, path := range paths {
		err := cs.convertFile(path)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// search target files
func (cs *ConvertService) dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, cs.dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}

		if filepath.Ext(file.Name()) == cs.BeforeExt {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

// convert files extesion
func (cs *ConvertService) convertFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	// ファイルオブジェクトを画像オブジェクトに変換

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outPath := cs.replaceExt(path, cs.BeforeExt, cs.AfterExt)
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	switch cs.AfterExt {
	case "png":
		err = png.Encode(out, img)
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "gif":
		err = gif.Encode(out, img, nil)
	default:
		// error
	}

	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	fmt.Printf("%#v\n", "complete")
	return nil
}

// replace extension to target type
func (cs *ConvertService) replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}
