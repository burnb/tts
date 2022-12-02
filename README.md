## TTS

### Environment vars

```
API_KEY={Cognitive Services Key}
REGION={Azure Region}
```

## Usage

```
Usage: tts [OPTIONS]
Generate batch audio files with speech synthesis service(Provided by Azure Cognitive Services)

Options:
  -h, --help            Show this help message and exit
  -i INPUT_FILE         Input file path (default "text.txt")
  -s FILE_SEPARATOR     Columns separator in the input file (default "/")
  -l LOCALE             Locale (default "th-TH")
  -v VOICE              Voice (default "th-TH-AcharaNeural")
```