package configs

import (
	"flag"

	"github.com/kelseyhightower/envconfig"
)

const (
	defaultFilePath  = "text.txt"
	defaultSeparator = "/"
	defaultLocale    = "th-TH"
	defaultVoice     = "th-TH-AcharaNeural"
)

type App struct {
	ApiKey       string `envconfig:"API_KEY" required:"true"`
	Region       string `envconfig:"REGION" required:"true"`
	FilePath     string
	Separator    string
	LocaleString string
	Voice        string
}

// Prepare variables to static configuration
func (c *App) Prepare() (err error) {
	if err := envconfig.Process("", c); err != nil {
		return err
	}

	flag.StringVar(&c.FilePath, "i", defaultFilePath, "input file path")
	flag.StringVar(&c.Separator, "s", defaultSeparator, "columns separator in the input file")
	flag.StringVar(&c.LocaleString, "l", defaultLocale, "locale")
	flag.StringVar(&c.Voice, "v", defaultVoice, "voice")
	flag.Parse()

	return nil
}
