package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func mkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

func abort_msg(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func decode(file *os.File, format string) (image.Image, error) {
	if format == "jpeg" {
		return jpeg.Decode(file)
	} else if format == "png" {
		return png.Decode(file)
	}
	return nil, errors.New("invalid format")
}
func encode(file *os.File, img image.Image, format string) error {
	if format == "jpeg" {
		opts := &jpeg.Options{Quality: 100}
		return jpeg.Encode(file, img, opts)
	} else if format == "png" {
		return png.Encode(file, img)
	}
	return errors.New("invalid format")
}

func main() {
	var (
		input  = flag.String("i", "", "input")
		format = flag.String("f", "png", "format")
		naming = flag.String("n", "", "naming rule")
		dir    = flag.String("d", "gosplit", "directory")
		row    = flag.Int("r", 0, "row")
		col    = flag.Int("c", 0, "column")
	)
	flag.Parse()
	// read args
	inputVal := *input
	formatVal := *format
	namingVal := *naming
	dirVal := *dir
	rowVal := *row
	colVal := *col
	// check arg
	if rowVal == 0 {
		abort_msg("please set row: -r 2")
	}
	if colVal == 0 {
		abort_msg("please set col: -c 2")
	}
	if !strings.HasSuffix(dirVal, "/") {
		dirVal += "/"
	}
	// read target files
	if inputVal == "" {
		var buf bytes.Buffer
		for _, arg := range flag.Args() {
			_, err := os.Stat(arg)
			if err == nil {
				buf.WriteString(arg)
				buf.WriteRune(' ')
			}
		}
		inputVal = buf.String()
	}
	image_file, err := os.Open(inputVal)
	defer image_file.Close()
	if err != nil {
		abort_msg(err.Error())
	}
	my_image, err := decode(image_file, formatVal)
	if err != nil {
		abort_msg(err.Error())
	}
	rect := my_image.Bounds()
	cwSize := rect.Size().X / colVal
	chSize := rect.Size().Y / rowVal
	iDir, iFile := filepath.Split(inputVal)
	mkdir(iDir + dirVal)
	for i := 0; i < rowVal; i++ {
		for j := 0; j < colVal; j++ {
			iFile = strings.TrimSuffix(iFile, filepath.Ext(iFile))
			temp := ""
			if namingVal == "" {
				temp = strconv.Itoa(i) + "_" + strconv.Itoa(j)
			} else {
				temp = string([]rune(namingVal)[j+(i*colVal)])
			}
			out_path := iDir + dirVal + iFile + "_" + temp + "." + formatVal
			newfile, err := os.Create(out_path)
			fmt.Println("create: " + out_path)
			defer newfile.Close()
			if err != nil {
				abort_msg(err.Error())
			}
			my_sub_image := my_image.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(image.Rect(j*cwSize, i*chSize, (j*cwSize)+cwSize, (i*chSize)+chSize))
			encode(newfile, my_sub_image, formatVal)
		}
	}

}
