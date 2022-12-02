package configs

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

const (
	Name = "tts"

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
	flagSet      *flag.FlagSet
}

// Prepare variables to static configuration
func (c *App) Prepare() (err error) {
	c.flagSet = flag.NewFlagSet(Name, flag.ExitOnError)
	c.flagSet.Usage = c.printUsage

	if err := envconfig.Process("", c); err != nil {
		return err
	}

	help := c.flagSet.Bool("h", false, "Show this help message and exit")
	c.flagSet.StringVar(&c.FilePath, "i", defaultFilePath, "input file path")
	c.flagSet.StringVar(&c.Separator, "s", defaultSeparator, "columns separator in the input file")
	c.flagSet.StringVar(&c.LocaleString, "l", defaultLocale, "locale")
	c.flagSet.StringVar(&c.Voice, "v", defaultVoice, "voice")

	if err = c.flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if *help {
		c.flagSet.Usage()
		os.Exit(1)
	}

	return nil
}

func (c *App) printUsage() {
	fmt.Println("\nUsage: tts [OPTIONS]\nGenerate batch audio files with speech synthesis service(Provided by Azure Cognitive Services)\n\nOptions:")
	c.flagSet.PrintDefaults()
}
