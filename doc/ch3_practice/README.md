# お付き合いを実践してみる

## どのような入力が与えられても意図通りに動作させる
### 問題設定
さて、最初の実践はシンプルにいってみましょう。問題設定として、与えられたディレクトリパスと同じ階層に、保存用の別ディレクトリを作ることにします。つまり。

```shell:before
(pwd)
  |- workdir/
      |- src/
```
```shell:after
(pwd)
  |- workdir/
      |- src/
      |- {name}/
```

上記のようにディレクトリを操作したいという状況です。ここでは簡単のためにディレクトリの作成はせず、作成するディレクトリのパスを生成するところまでとさせて頂きます。ではいってみましょう。

### 実践
この問題で唯一悩ましいのは入力の与えられ方にバリエーションがあることです。絶対パスは置いておいて、次のようなパターンがありえます。（なお、冗長なので省略しますが、Windowsのパスも与えられる可能性があります。）
```go
# Unix path
"./workdir/src"
"./workdir/src/"
"workdir/src"
"workdir/src/"
```

前の章でも見たように、`Dir`や`Split`を使うと意図しない結果を得る可能性があります。しかし`Clean`さんに登場してもらうことでその悩みから解放されます。
```go
func makeBrotherPath(path, name string) string {
	cleaned := filepath.Clean(path)
	parent := filepath.Dir(cleaned)
	brother := filepath.Join(parent, name)
	return brother
}
```
このように関数を実装することで、次のテストをパスします。
```go
func TestMakeBrotherPath(t *testing.T) {
	tests := []struct {
		src  string
		name string
		want string
	}{
		{"./workdir/src", "dst", "workdir/dst"},
		{"./workdir/src/", "dst", "workdir/dst"},
		{"workdir/src", "dst", "workdir/dst"},
		{"workdir/src/", "dst", "workdir/dst"},
	}
	for _, tt := range tests {
		got := makeBrotherPath(tt.src, tt.name)
		// fit "want" to each OS
		want := filepath.FromSlash(tt.want)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
}
```

期待する出力の文字列に`./`や`trailing slash`を含まない理由は、ここまで本書を読んで頂いた皆さんにはお分かりのことかと思います。ふふふ。

### まとめ
外部からどのような形式でパスが与えられるか予期できない場合、まずは`filepath.Clean`を適用することでパスの形式を一意に定めることができ、思わぬ動作を防ぐことができました。`filepath`パッケージの中には内部的に`Clean`を適用してから処理を始めるメソッドもありますので、各メソッドを深く理解している場合は使い分けるのがもちろん良いですが、最初のうちは渡されたパス文字列に対しまずは`Clean`ぐらいの意識でも良いのかな？と筆者は思います。（個人の意見です）