package rb

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Keywords []struct {
		Name string
	}
	Twitter struct {
		ConsumerKey       string `yaml:"consumer_key"`
		ConsumerSecret    string `yaml:"consumer_secret"`
		AccessToken       string `yaml:"access_token"`
		AccessTokenSecret string `yaml:"access_token_secret"`
	}
	StopWords []string `yaml:"stop_words"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config

	dat, err := ioutil.ReadFile("config.yaml")

	err = yaml.Unmarshal([]byte(dat), &config)
	if err != nil {
		return nil, err
	}

	for idx, stopWord := range config.StopWords {
		config.StopWords[idx] = strings.ToLower(stopWord)
	}

	return &config, nil
}
