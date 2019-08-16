package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Luzifer/rconfig/v2"
)

var (
	cfg = struct {
		HTML           bool   `flag:"html" default:"true" description:"Print HTML tags with SRI information (If disabled just prints the hashes)"`
		HTMLTag        string `flag:"html-tag" default:"link" description:"Tag to use for HTML mode (supported: link, script)"`
		LogLevel       string `flag:"log-level" default:"info" description:"Log level (debug, info, warn, error, fatal)"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	htmlTemplate = map[string]string{
		"link":   `<link href=%q integrity=%q crossorigin="anonymous">`,
		"script": `<script src=%q integrity=%q></script>`,
	}

	version = "dev"
)

func init() {
	if err := rconfig.ParseAndValidate(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("sri %s\n", version)
		os.Exit(0)
	}

	if l, err := log.ParseLevel(cfg.LogLevel); err != nil {
		log.WithError(err).Fatal("Unable to parse log level")
	} else {
		log.SetLevel(l)
	}
}

func main() {
	for _, url := range rconfig.Args()[1:] {
		logger := log.WithField("url", url)
		logger.Debug("Fetching SRI...")

		hash, err := sriIntegrity(url)
		if err != nil {
			logger.WithError(err).Error("Unable to fetch SRI hash")
			continue
		}

		if cfg.HTML {
			fmt.Printf(htmlTemplate[cfg.HTMLTag], url, hash)
			fmt.Println()
			continue
		}

		fmt.Printf("%s\t%s\n", url, hash)
	}
}

func sriIntegrity(url string) (string, error) {
	resp, err := http.Get(url) //#nosec G107 -- The while intention of this tool is to download files
	if err != nil {
		return "", errors.Wrap(err, "Unable to get URL contents")
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Unable to read body")
	}

	hash := sha512.Sum512(payload)
	return "sha512-" + base64.StdEncoding.EncodeToString(hash[:]), nil
}
