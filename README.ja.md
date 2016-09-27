[![Travis Build Status](https://travis-ci.org/craftgear/wincase.svg?branch=master)](https://travis-ci.org/craftgear/wincase)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftgear/wincase)](https://goreportcard.com/report/github.com/craftgear/wincase)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
<!--[![GoDoc](https://godoc.org/github.com/craftgear/wincase?status.svg)](https://godoc.org/github.com/craftgear/wincase)-->

[English](https://github.com/craftgear/wincase/blob/master/README.md)
 / 日本語

# wincase

Windows環境でファイルやディレクトリ名に使えない半角文字を、全角文字に変換するプログラムです。

``hoge*fuga.txt`` を ``hoge＊fuga.txt`` にします。

指定したディレクトリ以下を再帰的に検索して、ファイルとディレクトリの名前を変更します。

# インストール

 ``bin/対応する環境`` からバイナリをコピーするか、``go get github.com/craftgear/wincase`` してください。

# 使い方

カレントディレクトリ以下のファイル/ディレクトリをリネームするには、
``wincase ./``
とします。

実際にリネームを行わず、どうリネームされるかを表示するには
``wincase -dry-run ./``
とします。

実行中にログを出すには、
``wincase -v ./``
とします。

ヘルプを表示するには、
``wincase -h``
とします。

### Author
craftgear (https://twitter.com/craftgear)
