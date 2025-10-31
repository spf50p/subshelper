package conf

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

// default config
var Conf = Config{
	WorkDir: "/var/lib/subshelper",
	Subscription: Subscription{
		Title:        "SubsHelper",
		TitleUrlText: "Subscription URL",
		BaseUrl:      "https://s.subshelper.42",
		GlobalHeaders: map[string]string{
			"profile-title":           "SubsHelper",
			"profile-update-interval": "10",
		},
	},
}

type Config struct {
	WorkDir      string       `yaml:"work_dir"`
	Subscription Subscription `yaml:"subscription"`
}

type Subscription struct {
	Title         string            `yaml:"title"`
	TitleUrlText  string            `yaml:"title_url_text"`
	BaseUrl       string            `yaml:"base_url"`
	GlobalHeaders map[string]string `yaml:"global_headers"`
	Subs          []Sub             `yaml:"subs"`
}

type Sub struct {
	ID      string            `yaml:"id"`
	Links   []string          `yaml:"links"`
	Headers map[string]string `yaml:"headers"`
}

func Load(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = yaml.Unmarshal(b, &Conf)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
}
