# 🎨 ffgen

> @font-face generator for CSS!

## Installing

With a working [Go](https://golang.org) installation, just run this command to install:

```
go get github.com/cjdenio/ffgen/cmd/ffgen
```

You'll need to add your Go bin to your `$PATH`, if you haven't already.

## Running

`ffgen <path to font directory> <path to css file>`

If the CSS file doesn't exist, `ffgen` will create it for you. Otherwise, it'll append to the existing file.

## Building

```
git clone https://github.com/cjdenio/ffgen
cd ffgen
make
```

The built binary will be placed in the `bin/` directory.
