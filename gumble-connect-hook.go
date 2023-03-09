package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"net/http"
	"gopkg.in/yaml.v3"
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleutil"
	_ "layeh.com/gumble/opus"
)

var SUPPORTED_METHODS = []string{"GET", "POST"}

type MumbleConfig struct {
	Host string `yaml:"host"`
	Username string `yaml:"username"`
}

type Hook struct {
	Method string `yaml:"method"`
	Url string `yaml:"url"`
}

type Config struct {
	Mumble MumbleConfig
	Hooks []Hook
}

func main() {
	f, err := os.Open("gumble-connect-hook.yml")
	if err != nil {
		log.Fatalln("Could not open gumble-connect-hook.yml. Check if the file is there")
		panic(err)
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	gumbleConfig := gumble.NewConfig()
	gumbleConfig.Username = config.Mumble.Username
	gumbleConfig.Attach(gumbleutil.Listener{
		UserChange: func(e *gumble.UserChangeEvent) {
			if e.Type.Has(gumble.UserChangeConnected) {
				fmt.Printf("User %s connected\n", e.User.Name)
				for _, hook := range config.Hooks {
					invokeHook(hook)
				}
			}
		},
	})
	_, err = gumble.Dial(config.Mumble.Host, gumbleConfig)
	if err != nil {
		panic(err)
	}
	
	for true {
		time.Sleep(time.Second)
	}
}

func invokeHook(hook Hook) {
	if isMethodSupported(hook.Method) {
		client := &http.Client{}

		req, err := http.NewRequest(hook.Method, hook.Url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = client.Do(req)
	} else {
		log.Fatalln("Unknown/Unsupported Method: %s", hook.Method)
	}
	fmt.Printf("invoking %s hook to %s.\n", hook.Method, hook.Url)
}

func isMethodSupported(method string) bool {
	for _, m := range SUPPORTED_METHODS {
		if m == method {
			return true
		}
	}
	return false
}
