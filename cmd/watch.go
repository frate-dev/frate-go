package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"frate-go/config"

	"github.com/spf13/cobra"

	"github.com/fsnotify/fsnotify"
)

var WatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch and rebuild your project",
	Run: func(cmd *cobra.Command, args []string) {
		Watch()
	},
}

func Watch() {
	cfg, err := config.ReadConfig()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	fmt.Println("Watching for changes...")
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					Run()
					continue
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	filepath.Walk(cfg.SourceDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
