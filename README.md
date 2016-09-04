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
