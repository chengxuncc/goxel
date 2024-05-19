# <img src="https://user-images.githubusercontent.com/918212/56085390-af18c780-5e42-11e9-9ae7-7ba453502ddb.png" width="300">

GoXel is a download accelerator written in Golang, which inspired
by [axel](https://github.com/axel-download-accelerator/axel).

## Features

- Unlimited multithreading downloads
- Download multiple files concurrently
- Resume unfinished downloads
- Monitor download progress
- Guess filename from URL path

## Installation

```shell
go install github.com/chengxuncc/goxel/cmd/goxel@latest
```

## Usage

```
$ bin/goxel -h
GoXel is a download accelerator written in Go
Usage: goxel [options] [url1] [url2] [url...]
      --alldebrid-password string         Alldebrid password, can also be passed in the GOXEL_ALLDEBRID_PASSWD environment variable                                                                                 
      --alldebrid-username string         Alldebrid username, can also be passed in the GOXEL_ALLDEBRID_USERNAME environment variable                                                                               
      --buffer-size int                   Buffer size in KB (default 256)
  -f, --file string                       File containing links to download (1 per line)
      --header header-name=header-value   Extra header(s) (default [])
  -h, --help                              This information
      --insecure                          Bypass SSL validation
      --max-conn int                      Max number of connections (default 8)
  -m, --max-conn-file int                 Max number of connections per file (default 4)
      --no-resume                         Don't resume downloads
  -o, --output string                     Output directory
      --overwrite                         Overwrite existing file(s)
  -p, --proxy string                      Proxy string: (http|https|socks5)://0.0.0.0:0000
  -q, --quiet                             No stdout output
  -s, --scroll                            Scroll output instead of in place display
      --version                           Version

Visit https://github.com/m1ck43l/goxel/issues to report bugs.
```

## Benchmark

This benchmark compares Axel and GoXel for multiple downloads using files from https://www.thinkbroadband.com/download.
All links were done using a broadhand connection: 455.0 Mbit/s download, 276.4 Mbit/s upload, lantency 3ms over WiFi.

![Benchmark](https://user-images.githubusercontent.com/918212/56504862-2e308e80-651a-11e9-96de-398bf263b060.png)

## Contributing

Pull requests for new features, bug fixes, and suggestions are welcome!

## License

[Apache 2](https://github.com/m1ck43l/goxel/blob/master/LICENSE)

## Credits

[m1ck43l/goxel](https://github.com/m1ck43l/goxel)