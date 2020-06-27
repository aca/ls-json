package main

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

type File struct {
	Ext     string    `json:"ext"`
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Bytes   int64     `json:"bytes"`
	Size    string    `json:"size"`
	ModTime time.Time `json:"mod_time"`
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "ls-json [FILE]",
		Short: "List information about the FILE in json format (the current directory by default)",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := new(LSJson)
			cd, err := os.Getwd()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				c.Path = append(c.Path, cd)
			} else {
				c.Path = args[:]
			}

			return c.Run()
		},
	}

	rootCmd.Execute()
}

type LSJson struct {
	Path []string
}

func (c *LSJson) Run() error {
	enc := json.NewEncoder(os.Stdout)

	for _, p := range c.Path {

		f, err := os.Lstat(p)
		if err != nil {
			return err
		}

		// if file
		switch f.IsDir() {
		case true:
			err = filepath.Walk(p, func(root string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				jf, err := toJSON(root, info)
				if err != nil {
					return err
				}

				err = enc.Encode(jf)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			})

			if err != nil {
				log.Fatal(err)
			}
		case false:
			jf, err := toJSON(p, f)
			if err != nil {
				return err
			}

			err = enc.Encode(jf)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func toJSON(root string, info os.FileInfo) (*File, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	f := &File{}
	f.Path = abs
	f.Bytes = info.Size()
	f.ModTime = info.ModTime()
	f.Size = humanize.BigBytes(big.NewInt(info.Size()))
	f.Name = info.Name()
	f.Ext = filepath.Ext(info.Name())

	return f, nil
}
