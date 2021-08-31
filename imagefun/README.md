# imagefun

A command line tool that draws an image with random colours.

## Usage

From within the project folder:
```
go build imagefun.go
go run imagefun
```

Use optional command line arguments:
```
go run imagefun -w <width in pixels> -h <height in pixels> -f <filename> -flag true
```
Otherwise default values will be applied.
The option *-flag true* will require additional user input.