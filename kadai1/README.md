## 画像変換コマンド

### 下記の仕様を満たす
- ディレクトリを指定
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


### 下記を満たすように開発
- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 標準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる

### コマンドラインオプション

 | オプション | 説明 | デフォルト |
 | --- | --- | --- |
 | -pre | 変換前フォーマット | jpeg |
 | -post | 変換後フォーマット | png |


### 対応している画像フォーマット
- png, jpeg(jpg), gif


### 使い方
1. バイナリビルド（実行ファイル作成）
```bash
$ make
```
2. ディレクトリを指定して実行
```bash
$ ./converter [探索ディレクトリ] [変更前拡張子] [変換後拡張子]
```