package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
	EnvKeyPolygons    = "DEMO_POLYGONS"
	EnvKeyPixels      = "DEMO_PIXELS"
	EnvKeyResolution  = "DEMO_RESOLUTION"
)

var (
	src         = "blog-image-demo.files/pictures"
	processed   = "blog-image-demo.files/processed"
	placeholder = "blog-image-demo.files/pictures/vorteil.png"
	polygons    = 50
	resolution  = 64
	pixels      = 512
)

var list []tuple

type tuple struct {
	Name  string `json:"name"`
	Ready bool   `json:"ready"`
}

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

	s = os.Getenv(EnvKeyPolygons)
	if s != "" {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if x <= 1000 && x >= 10 {
			polygons = x
		}
	}

	s = os.Getenv(EnvKeyResolution)
	if s != "" {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if x <= 1024 && x >= 64 {
			resolution = x
		}
	}

	s = os.Getenv(EnvKeyPixels)
	if s != "" {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if x <= 4096 && x >= 64 {
			pixels = x
		}
	}

	fmt.Printf("Source Pictures: %s\n", src)
	fmt.Printf("Processed Pictures: %s\n", processed)
	fmt.Printf("Polygons Per Picture: %v\n", polygons)
	fmt.Printf("Processing Resolution: %v\n", resolution)
	fmt.Printf("Thumbnail Pixel Width: %v\n", pixels)

	os.RemoveAll(processed)
	err := os.MkdirAll(processed, 0777)
	if err != nil {
		log.Fatal(err)
	}

	list = make([]tuple, 0)
	fis, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fis {
		_, err = primitive.LoadImage(filepath.Join(src, fi.Name()))
		if err != nil {
			continue
		}
		list = append(list, tuple{Name: fi.Name()})
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

	fmt.Printf("Listening on port: 8080\n")

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

	size := uint(resolution)

	input = resize.Thumbnail(size, size, input, resize.Bilinear)

	rand.Seed(time.Now().UTC().UnixNano())

	// background color
	bg := primitive.MakeColor(primitive.AverageImageColor(input))

	// run algorithm
	outputSize := pixels
	workers := 1

	model := primitive.NewModel(input, bg, outputSize, workers)
	frame := 0

	count := polygons

	mode := 1
	alpha := 128
	repeat := 0

	for i := 0; i < count; i++ {
		frame++

		// find optimal shape and add it to the model
		model.Step(primitive.ShapeType(mode), alpha, repeat)

		last := i == count-1
		var path string
		if last {
			path = outpath
			err = primitive.SavePNG(path, model.Context.Image())
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("  COMPLETE\n")

	for i := range list {
		if list[i].Name == filepath.Base(inpath) {
			list[i].Ready = true
			break
		}
	}
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	path := mux.Vars(r)["path"]

	if path == "" {
		data, err := Asset("web/index.html")
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, err.Error(), code)
			return
		}
		w.Write(data)
		return
	}

	if path == "api/pictures" {
		data, err := json.Marshal(list)
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			return
		}
		w.Write(data)
		return
	}

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
