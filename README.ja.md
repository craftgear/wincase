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
