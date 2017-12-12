package main

import (
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fogleman/primitive/primitive"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

// Environment variables.
const (
	EnvKeySource      = "DEMO_SOURCE_DIR"
	EnvKeyProcessed   = "DEMO_PROCESSED_DIR"
	EnvKeyPlaceholder = "DEMO_PLACEHOLDER_PICTURE"
)

var (
	src         = "blog-image-demo.files/pictures"
	processed   = "blog-image-demo.files/processed"
	placeholder = "blog-image-demo.files/pictures/vorteil.png"
)

func initialize() {
	s := os.Getenv(EnvKeySource)
	if s != "" {
		src = s
	}

	s = os.Getenv(EnvKeyProcessed)
	if s != "" {
		processed = s
	}

	s = os.Getenv(EnvKeyPlaceholder)
	if s != "" {
		placeholder = s
	}

	fmt.Printf("Source Pictures: %s\n", src)
	fmt.Printf("Processed Pictures: %s\n", processed)

	os.RemoveAll(processed)
	err := os.MkdirAll(processed, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	initialize()

	go processPictures()

	handler := mux.NewRouter()
	handler.HandleFunc("/{path:.*}", handle)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func processPictures() {

	fis, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fis {
		err = process(filepath.Join(src, fi.Name()), filepath.Join(processed, fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("All files processed.\n")
}

func process(inpath, outpath string) error {

	input, err := primitive.LoadImage(inpath)
	if err != nil {
		if err == image.ErrFormat {
			return nil
		}
		return err
	}

	fmt.Printf("Processing file: %s", inpath)

	size := uint(256)

	input = resize.Thumbnail(size, size, input, resize.Bilinear)

	rand.Seed(time.Now().UTC().UnixNano())

	// background color
	bg := primitive.MakeColor(primitive.AverageImageColor(input))

	// run algorithm
	outputSize := 1024
	workers := 1

	model := primitive.NewModel(input, bg, outputSize, workers)
	frame := 0

	count := 100

	Nth := 1

	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	output := tmp.Name()
	defer os.Remove(output)
	tmp.Close()

	mode := 1
	alpha := 128
	repeat := 0

	// fmt.Printf("temp file: %s\n", tmp.Name())

	for i := 0; i < count; i++ {
		frame++

		// fmt.Printf("frame: %v\n", i)

		// find optimal shape and add it to the model
		model.Step(primitive.ShapeType(mode), alpha, repeat)

		// write output image(s)

		ext := strings.ToLower(filepath.Ext(output))
		percent := strings.Contains(output, "%")
		saveFrames := percent && ext != ".gif"
		saveFrames = saveFrames && frame%Nth == 0
		last := i == count-1
		var path string
		if saveFrames || last {
			path = outpath
			if percent {
				path = fmt.Sprintf(outpath, frame)
			}
			err = primitive.SavePNG(path, model.Context.Image())
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("  COMPLETE\n")
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {

	path := mux.Vars(r)["path"]
	path = filepath.Join(processed, path)

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			var fi os.FileInfo
			fi, err = os.Stat(filepath.Join(src, mux.Vars(r)["path"]))
			if err == nil {
				if !fi.IsDir() {
					f, err = os.Open(placeholder)
					if err != nil {
						code := http.StatusInternalServerError
						http.Error(w, http.StatusText(code), code)
						return
					}
					defer f.Close()

					_, err = io.Copy(w, f)
					if err != nil {
						code := http.StatusInternalServerError
						http.Error(w, http.StatusText(code), code)
						return
					}

					return
				}
			}
			code := http.StatusNotFound
			http.Error(w, http.StatusText(code), code)
			return
		}
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

}
