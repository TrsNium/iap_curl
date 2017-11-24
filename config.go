package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	neturl "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/agext/levenshtein"
)

type Config struct {
	Services []Service `json:"service"`
}

type Service struct {
	URL string `json:"url"`
	Env Env    `json:"env"`
}

type Env struct {
	Credentials string `json:"GOOGLE_APPLICATION_CREDENTIALS"`
	ClientID    string `json:"IAP_CLIENT_ID"`
	Binary      string `json:"IAP_CURL_BIN"`
}

func configDir() (string, error) {
	var dir string

	switch runtime.GOOS {
	default:
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	case "windows":
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data")
		}
	}
	dir = filepath.Join(dir, "iap_curl")

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return dir, fmt.Errorf("cannot create directory: %v", err)
	}

	return dir, nil
}

func (cfg *Config) LoadFile(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		raw, _ := ioutil.ReadFile(file)
		if err := json.Unmarshal(raw, cfg); err != nil {
			return err
		}
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	// Insert sample config map as a default
	if len(cfg.Services) == 0 {
		cfg.Services = []Service{Service{
			URL: "https://iap-protected-app-url",
			Env: Env{
				Credentials: "/path/to/google-credentials.json",
				ClientID:    "foobar.apps.googleusercontent.com",
				Binary:      "curl",
			},
		}}
	}

	return json.NewEncoder(f).Encode(cfg)
}

func (cfg *Config) GetEnv(url string) (env Env, err error) {
	u1, _ := neturl.Parse(url)
	for _, service := range cfg.Services {
		u2, _ := neturl.Parse(service.URL)
		if u1.Host == u2.Host {
			return service.Env, nil
		}
	}
	err = fmt.Errorf("%s: no such host in config file", u1.Host)
	return
}

func (cfg *Config) GetURLs() (list []string) {
	for _, service := range cfg.Services {
		list = append(list, service.URL)
	}
	return
}

func (cfg *Config) Edit() error {
	dir, _ := configDir()
	json := filepath.Join(dir, "config.json")
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	command := fmt.Sprintf("%s %s", editor, json)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func (cfg *Config) SimilarURLs(url string) (urls []string) {
	u1, _ := neturl.Parse(url)
	for _, service := range cfg.Services {
		u2, _ := neturl.Parse(service.URL)
		degree := round(levenshtein.Similarity(u1.Host, u2.Host, nil) * 100)
		if degree > 50 {
			urls = append(urls, service.URL)
		}
	}
	return
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
