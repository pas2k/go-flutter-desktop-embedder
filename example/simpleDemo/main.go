package main

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"path"
	"runtime"

	gutter "github.com/Drakirus/go-flutter-desktop-embedder"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	var (
		err error
	)

	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFilePath)

	options := []gutter.Option{
		gutter.OptionAssetPath(dir + "/flutter_project/demo/build/flutter_assets"),
		gutter.OptionICUDataPath("/opt/flutter/bin/cache/artifacts/engine/linux-x64/icudtl.dat"),
		gutter.OptionWindowInitializer(setIcon),
		gutter.OptionWindowDimension(800, 600),
		gutter.OptionWindowInitializer(setIcon),
		gutter.OptionPixelRatio(1.2),
		gutter.OptionVMArguments([]string{"--dart-non-checked-mode", "--observatory-port=50300"}),
	}

	if err = gutter.Run(options...); err != nil {
		log.Fatalln(err)
	}

}

func setIcon(window *glfw.Window) error {
	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFilePath)
	imgFile, err := os.Open(dir + "/assets/icon.png")
	if err != nil {
		return err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}
	window.SetIcon([]image.Image{img})
	return nil
}
