# FileOrganize

Organize file Command.

## Features

- FileOrganize is CLI tool because file organized.
- You can associate the extension with the save destination directory using the Config file.
- You can use Ableton sub command that organize project music file(mp3, wav).

## Usage

The usage:

```
  FileOrganize -t ./ -o ./output -c ./config.json
```

Available options:

```
  -t, --target  Target Directory.
  -o, --output  Output Directory.
  -c, --config  Specify config json file.
```

Sub Command:

```
  FileOrganize Ableton -t ./ -o ./output
```

Available options:

```
  -t, --target  Target Directory.
  -o, --output  Output Directory.
```

## Author

takuto wakinaka https://github.com/irisu1975

## License

Distributed under the MIT License.
