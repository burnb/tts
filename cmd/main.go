package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
	"github.com/joho/godotenv"

	"github.com/burnb/tts/internal/configs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if _, ok := err.(*os.PathError); !ok {
			log.Fatal(err)
		}
	}

	cfg := &configs.App{}
	if err := cfg.Prepare(); err != nil {
		log.Fatal(err)
	}

	stream, err := audio.CreatePullAudioOutputStream()
	if err != nil {
		log.Panic(fmt.Sprintf("create pull audio output stream error: %v", err))
	}
	defer stream.Close()

	speechConfig, err := speech.NewSpeechConfigFromSubscription(cfg.ApiKey, cfg.Region)
	if err != nil {
		log.Panic(err)
	}
	err = speechConfig.SetSpeechSynthesisOutputFormat(common.Audio24Khz48KBitRateMonoMp3)
	if err != nil {
		log.Panic(err)
	}
	defer speechConfig.Close()

	if cfg.Voice != "" {
		err := speechConfig.SetSpeechSynthesisVoiceName(cfg.Voice)
		if err != nil {
			log.Panic(err)
		}
	}

	speechSynthesizer, err := speech.NewSpeechSynthesizerFromConfig(speechConfig, nil)
	if err != nil {
		log.Panic(err)
	}
	defer speechSynthesizer.Close()

	file, err := os.Open(cfg.FilePath)
	if err != nil {
		log.Panic(fmt.Sprintf("unable to open input file %s: %v", cfg.FilePath, err))
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var i uint16
	for scanner.Scan() {
		i++
		row := scanner.Text()
		parts := strings.Split(row, cfg.Separator)

		fileName := fmt.Sprintf("%d-%s.mp3", i, parts[0])

		definition := parts[0]
		if len(parts) > 1 {
			definition = parts[1]
		}

		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			log.Printf("Synthesize: %s to %s\n", definition, fileName)
			task := speechSynthesizer.SpeakTextAsync(definition)

			var outcome speech.SpeechSynthesisOutcome
			select {
			case outcome = <-task:
			case <-time.After(10 * time.Second):
				log.Panic("Synthesize timed out")
			}
			defer outcome.Close()

			if outcome.Error != nil {
				log.Panic(outcome.Error.Error())
			}

			if outcome.Result.Reason == common.SynthesizingAudioCompleted {
				if err = os.WriteFile(fileName, outcome.Result.AudioData, 0644); err != nil {
					log.Panic(fmt.Sprintf("unable to write file, received %v", err))
				}
			} else {
				cancellation, err := speech.NewCancellationDetailsFromSpeechSynthesisResult(outcome.Result)
				if err != nil {
					log.Panic(fmt.Sprintf("unable to get cancellation details %v", err))
				}
				if cancellation.Reason == common.Error {
					log.Panic(fmt.Printf("ErrorCode:%d\n Error:%s\n", cancellation.ErrorCode, cancellation.ErrorDetails))
				}
			}
		}
	}

	log.Print("Done.")
}
