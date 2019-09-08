## BackGround
Go はシングルバイナリであることによる、配布のしやすさが利点である言語です。
しかし、時々 Go のソースコードへは含めることのできないファイル (テンプレートファイルなど) を、 
外部に持たせられなければならない場合がります。

そこで、Go では Package (外部ライブラリ) により、 
build した実行隊に静的ファイルを埋め込む機能がいくつか提案されてきました。 

今回はその中でも、最近リリースされた [assets-life](https://github.com/shogo82148/assets-life) という package を使ってみたので紹介します。  

## Target
* 制約なし

## Goal
* `assets-life` を使って Go のバイナリに静的ファイルを埋め込めるようになる

## Environment
* `Go`
 * `1.12`
* `assets-life`
 * `9f8d070f229f24f8dd21ba18b19e3aedfa895d3e`

## assets-life
今回紹介する `assets-life` の他には、 
以下のツールが同じ機能を持つものとしてリリースされています。  

* [jessevdk/go-assets](https://github.com/jessevdk/go-assets)
* [jessevdk/go-bindata](https://github.com/jteeuwen/go-bindata)
* [rakyll/statik](https://github.com/rakyll/statik)

これらのうち、`go-assets` と `go-bindata` は、 
以下の理由から、今開発に採用するのは難しい状態となっています。  

* 長期間メンテナンスされていない

なお、`statik` は、 `Awesome Go` に掲載されており、現在のスタンダードとなっているようです。

今回紹介する、`assets-life` は、2019/7 にリリースされたツールです。 
作者の意図として、これまでの競合ツールより使いやすく、パフォーマンスを向上させる意図があるようです。 

この package には以下の特徴があります。  

* コードの再生成にコマンドの再インストールが不要
* ファイルの検索にバイナリサーチを使用しており、概ねのユースケースで高速 (※)

※ 静的ファイルの数が1000以下の場合

なお、本記事執筆時点のリポジトリを見た所では、以下の点には気をつけた方が良さそうです。 
これからの開発に期待です。  

* エラーハンドリングを一部省略している
    * `filepath.Walk()` の `WalkFunc` のエラー処理をしていない
* Windows で `go generate` によるコード再生成ができない


## Example
使い方を簡単な例を元に説明します。  
コマンドのインストールまでの手順については、github の Usage を参照してください。

まず、以下のような静的ファイルのソースを用意します。  

```
.
├── assets_src
     └── hoge.txt
```

```hoge.txt
hoge
```

次に、`assets-life` コマンドでソースから Go ソースコードを生成します。  

```
$ assets-life ./assets_src ./assets
($ cmd src dst)
```

すると、以下のようにファイルが生成されます。

```
.
├── assets
│   ├── assets-life.go
│   └── filesystem.go
├── assets_src
     └── hoge.txt
```

生成された `filesystem.go` には、読み込んだファイルが以下のように実装されています。 
このファイルを `assets` package として import することができます。  
変数 `Root` は `http.FileSystem` インターフェースを実装しており、
 `Open()` メソッドを利用して、各ファイルにアクセスすることができます。  

```filesystem.go
package assets

(中略)

// Root is the root of the file system.
var Root http.FileSystem = fileSystem{
	file{
		name:    "/",
		content: "",
		mode:    0755 | os.ModeDir,
		next:    0,
		child:   1,
	},
	file{
		name:    "/hoge.txt",
		content: "hoge\r\n",
		mode:    0644,
		next:    -1,
		child:   -1,
	},
}

(以下略)
```

このパッケージを元に、以下のようなユーザプログラムを書くと、
静的ファイルにアクセスすることができます。

```main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/akif999/prototypes/go/assets_life_exp/assets"
)

func main() {
	f, err := assets.Root.Open("/hoge.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))
}
```

```
$ XXX.exe
hoge
```

とても簡単に静的ファイルをバイナリに含めることができました。

また、特徴にある通り、静的ファイルを変更した場合、初回生成に使用したコマンドは不要です。 
以下のように Go 標準ツールの `go generate` によって、再生成することができます。 (※)

```
$ go generate ./assets
```

※ Windows 環境では、生成される `go:generate` への引数となるパスセパレータに問題があり、動作しませんでした

## RoundUp
以下、今回の記事のまとめです。

* Usage がわかりやすく誰でも簡単に使える
* まだ開発中な package である
* 現状は `statik` を使用した方が安定していると考えられる

## Reference
* [github.com/shogo82148/assets-life](https://github.com/shogo82148/assets-life)
* [Goのバイナリに静的ファイルを埋め込むツール assets-life を書いた](https://shogo82148.github.io/blog/2019/07/24/assets-life/)
