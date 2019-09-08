# assets-life で Go のバイナリに静的ファイルを埋め込む

## BackGround
Go はシングルバイナリであることによる、配布のしやすさが利点である言語です。  
しかし、時々 Go のソースコードへは含めることのできないファイル (テンプレートファイルなど) を、  
外部に持たせられなければならない場合がります。

そこで、Go では Package (外部ライブラリ) により、  
build した実行隊に静的ファイルを埋め込む機能がいくつか提案されてきました。  

今回はその中でも、最近リリースされた [assets-life](https://github.com/shogo82148/assets-life) という package を紹介
します。  

## Target
* 制約なし

## Goal
* assets-life を使って Go のバイナリに静的ファイルを埋め込めるようになる

## assets-life
今回紹介する `assets-life` がリリースされる以前には、  
以下の package が同じ機能を持つものとしてリリースされています。  

* [jessevdk/go-assets](https://github.com/jessevdk/go-assets)
* [jessevdk/go-bindata](https://github.com/jteeuwen/go-bindata)
* [rakyll/statik](https://github.com/rakyll/statik)

これらのうち、`go-assets` と `go-bindata` には、
以下の理由から今採用するのは難しい状態となっています。

* 長期間メンテナンスされていない

`statik` は、 `Awesome Go` にも掲載されており、現在のスタンダードとなっています。

これらの package

 は、2019/7 にリリースされた package です。



## Example

## RoundUp

## Reference
* [github.com/shogo82148/assets-life](https://github.com/shogo82148/assets-life)
* [Goのバイナリに静的ファイルを埋め込むツール assets-life を書いた](https://shogo82148.github.io/blog/2019/07/24/assets-life/)
