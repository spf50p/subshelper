package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf50p/subshelper/conf"
	"github.com/spf50p/subshelper/tpl"
)

var rootCmd = &cobra.Command{
	Use:   "subshelper",
	Short: "Generate structure of directories for subscriptions",
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("conf", "c", os.ExpandEnv("$HOME/.subshelper.yaml"), "config file")
}

func run(cmd *cobra.Command, args []string) {
	confPath, err := cmd.Flags().GetString("conf")
	if err != nil {
		log.Fatalf("Failed to get config file: %v", err)
	}

	// read config
	conf.Load(confPath)

	// create work directory
	if _, err := os.Stat(conf.Conf.WorkDir); os.IsNotExist(err) {
		if err := os.MkdirAll(conf.Conf.WorkDir, 0755); err != nil {
			log.Fatalf("Failed to create work directory: %v", err)
		}
	} else {
		cleanWorkDir()
	}

	// create global headers.caddy file
	globalHeadersCaddyPath := filepath.Join(conf.Conf.WorkDir, ".headers.caddy")
	contentGlobalHeadersCaddy, err := tpl.Execute(tpl.GlobalHeader{Headers: conf.Conf.Subscription.GlobalHeaders}, "globalHeader", tpl.GlobalHeaderTpl)
	if err != nil {
		log.Fatalf("Failed to get global headers content: %v", err)
	}
	if err := os.WriteFile(globalHeadersCaddyPath, []byte(contentGlobalHeadersCaddy), 0644); err != nil {
		log.Fatalf("Failed to write .headers.caddy file: %v", err)
	}

	subsIndex := 0

	// create subs
	for _, sub := range conf.Conf.Subscription.Subs {
		subsIndex++
		// create sub directory
		subDir := filepath.Join(conf.Conf.WorkDir, sub.ID)
		if err := os.MkdirAll(subDir, 0755); err != nil {
			log.Fatalf("Failed to create sub directory: %v", err)
		}

		// .headers.caddy file
		headersCaddyPath := filepath.Join(subDir, ".headers.caddy")
		if len(sub.Headers) > 0 {
			contentSubHeadersCaddy, err := tpl.Execute(tpl.SubHeader{SubID: sub.ID, Headers: sub.Headers}, "subHeader", tpl.SubHeaderTpl)
			if err != nil {
				log.Fatalf("Failed to get sub headers content: %v", err)
			}
			if err := os.WriteFile(headersCaddyPath, []byte(contentSubHeadersCaddy), 0644); err != nil {
				log.Fatalf("Failed to write .headers.caddy file: %v", err)
			}
		} else {
			if _, err := os.Stat(headersCaddyPath); err == nil {
				if err := os.Remove(headersCaddyPath); err != nil {
					log.Fatalf("Failed to remove .headers.caddy file: %v", err)
				}
			}
		}

		// index.txt file
		indexTxtPath := filepath.Join(subDir, "index.txt")
		contentIndexTxt := base64.StdEncoding.EncodeToString([]byte(strings.Join(sub.Links, "\n")))

		if err := os.WriteFile(indexTxtPath, []byte(contentIndexTxt), 0644); err != nil {
			log.Fatalf("Failed to write index.txt file: %v", err)
		}

		// index.html file
		var subLinks []tpl.HtmlIndexSubLink
		indexHtmlPath := filepath.Join(subDir, "index.html")
		for _, link := range sub.Links {
			parts := strings.Split(link, "#")
			var title string
			if len(parts) > 1 {
				title = parts[1]
			} else {
				title = parts[0]
			}
			subLinks = append(subLinks, tpl.HtmlIndexSubLink{Title: strings.TrimSpace(title), Link: link})
		}
		htmlx := tpl.HtmlIndex{
			Title:        conf.Conf.Subscription.Title,
			TitleUrlText: conf.Conf.Subscription.TitleUrlText,
			Url:          filepath.Join(conf.Conf.Subscription.BaseUrl, conf.Conf.Subscription.PathSegment, sub.ID, ""),
			SubLinks:     subLinks,
		}
		contentIndexHTML, err := tpl.Execute(htmlx, "indexHtml", tpl.IndexHTMLTpl)
		if err != nil {
			log.Fatalf("Failed to get index.html content: %v", err)
		}
		if err := os.WriteFile(indexHtmlPath, []byte(contentIndexHTML), 0644); err != nil {
			log.Fatalf("Failed to write index.html file: %v", err)
		}
	}
	fmt.Println("Subs count:", subsIndex)
}

func cleanWorkDir() {
	// get all sub directories
	subDirs, err := os.ReadDir(conf.Conf.WorkDir)
	if err != nil {
		log.Fatalf("Failed to read work directory: %v", err)
	}

	// remove sub directories that are not in conf.Conf.Subs
	for _, subDir := range subDirs {
		if subDir.IsDir() && !slices.ContainsFunc(conf.Conf.Subscription.Subs, func(sub conf.Sub) bool {
			return sub.ID == subDir.Name()
		}) {
			if err := os.RemoveAll(filepath.Join(conf.Conf.WorkDir, subDir.Name())); err != nil {
				log.Fatalf("Failed to remove sub directory: %v", err)
			}
		}
	}
}
