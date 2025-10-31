// Author :
// https://github.com/MaulanaR/mr-auto-typer

package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/*
var assets embed.FS

func main() {
	app := NewApp()

	icon, err := assets.ReadFile("frontend/mrlabs.ico")
	if err != nil {
		log.Fatalf("Error reading embedded file: %v", err)
	}

	if err := wails.Run(&options.App{
		Title:       "Mr Auto Typer",
		Width:       1060,
		Height:      740,
		OnStartup:   app.startup,
		AssetServer: &assetserver.Options{Assets: assets},
		Bind:        []interface{}{app},
		Frameless:   false,
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "Mr Auto Typer",
				Message: "Its free, pls support for this project. https://github.com/MaulanaR/mr-auto-typer",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "Mr Auto Typer",
		},
	}); err != nil {
		log.Fatal(err)
	}
}
