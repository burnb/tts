## TTS

### Environment vars

```
API_KEY={Cognitive Services Key}
REGION={Azure Region}
```

## Usage

```
usage: tts [-h] [-V | -L | -Q | [-t [TEXT] [-p PITCH] [-r RATE] [-S STYLE] [-R ROLE] [-d STYLE_DEGREE] | -s [SSML]]]
              [-f FILE] [-e ENCODING] [-o OUTPUT_PATH] [-l LOCALE] [-v VOICE]
              [--mp3 [-q QUALITY] | --ogg [-q QUALITY] | --webm [-q QUALITY] | --wav [-q QUALITY] | -F FORMAT] 

Try speech synthesis service(Provided by Azure Cognitive Services) in your terminal!

options:
  -h, --help            show this help message and exit
  -V, --version         show program's version number and exit
  -L, --list-voices     list available voices, you can combine this argument with -v and -l
  -Q, --list-qualities-and-formats
                        list available qualities and formats
  -t [TEXT], --text [TEXT]
                        Text to speak. Left blank when reading from file/stdin
  -s [SSML], --ssml [SSML]
                        SSML to speak. Left blank when reading from file/stdin
  -f FILE, --file FILE  Text/SSML file to speak, default to `-`(stdin)
  -e ENCODING, --encoding ENCODING
                        Text/SSML file encoding, default to "utf-8"(Not for stdin!)
  -o OUTPUT_PATH, --output OUTPUT_PATH
                        Output file path, wav format by default
  --mp3                 Use mp3 format for output. (Only works when outputting to a file)
  --ogg                 Use ogg format for output. (Only works when outputting to a file)
  --webm                Use webm format for output. (Only works when outputting to a file)
  --wav                 Use wav format for output
  -F FORMAT, --format FORMAT
                        Set output audio format (experts only)
  -l LOCALE, --locale LOCALE
                        Locale to use, default to en-US
  -v VOICE, --voice VOICE
                        Voice to use
  -q QUALITY, --quality QUALITY
                        Output quality, default to 0
```