## TTS

### Environment vars

```
API_KEY={Cognitive Services Key}
REGION={Azure Region}
```

## Usage

```
usage: tts [-i INPUT_FILE] [-s FILE_SEPARATOR] [-l ] [-v VOICE]

Generate batch audio files by speech synthesis service(Provided by Azure Cognitive Services) in your terminal!

options:
  -h, --help            Show this help message and exit
  -i INPUT_FILE         Input file path (default "text.txt")
  -l FILE_SEPARATOR     Locale (default "th-TH")
  -s LOCALE             Columns separator in the input file (default "/")
  -v VOICE              Voice (default "th-TH-AcharaNeural")
```