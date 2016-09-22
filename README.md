[![Travis Build Status](https://travis-ci.org/craftgear/wincase.svg?branch=master)](https://travis-ci.org/craftgear/wincase)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftgear/wincase)](https://goreportcard.com/report/github.com/craftgear/wincase)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
<!--[![GoDoc](https://godoc.org/github.com/craftgear/wincase?status.svg)](https://godoc.org/github.com/craftgear/wincase)-->

English / [日本語](https://github.com/craftgear/wincase/blob/master/README.ja.md)

# wincase

Windows platforms have forbidden characters on filenames/directory names.

This utility renames the forbidden characters with corresponding 2-byte characters.

For example it renames ``foo*bar.txt`` to ``foo＊bar.txt``.

It searches files recursively from the directory you specify.

# Install

Copy a file from ``bin/your_platform`` directory to your system.

or

``go get github.com/craftgear/wincase``


# Usage

To rename all files under the current directory, run:
``wincase ./``

To display files to rename without renaming(so-called "dry run"), run:
``wincase -dry-run ./``

To run in verbose mode, run:
``wincase -v ./``

To show help, run:
``wincase -h``
