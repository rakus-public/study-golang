# Study Golang

![](https://img.shields.io/hexpm/l/apa)

## Get Started

[コチラのスライド](https://rakus-public.github.io/study-golang/marp/installation.html) に環境のセットアップおよび `tutorial` の進め方、動作確認などが掲載されていますので一読ください

## Set Up

テストの実行用に以下のタスクランナーを使用する想定です。

- [Task](https://taskfile.dev/)

利用しなくともテストの実行は可能ですが、すぐにインストールできるためよろしければ導入ください。
また、タスクを実行する際はプロジェクトルートにて実行してください。

### Via Go Module

Go をローカルにインストール済みの方は OS を問わず以下のコマンドで導入可能です。

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

### For Mac

```bash
brew install go-task/tap/go-task
```

### For Linux

```bash
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
```

### For Windows

```bash
choco install go-task
```

## Subject

各パッケージとその課題詳細は以下の通りです。

| Tool          | Desc                              |
| :------------ | :-------------------------------- |
| `tutorial`    | [README](./tutorial/README.md)    |
| `curl`        | [README](./curl/README.md)        |
| `cgrep`       | [README](./cgrep/README.md)       |
| `logtransfer` | [README](./logtransfer/README.md) |

## Test

各パッケージごとにテスト実行用のタスクを定義しています。
今回は同一ファイル内で実施日が異なるケースが多々存在するため、関数名を指定して実行するためのタスクを用意しています。
タスクに Args を渡す場合、 `<CMD> -- <Args>` のようにハイフン 2 つを挟むようにしてください。

| Task                 | Args                 | Desc                                                                                     |
| :------------------- | :------------------- | :--------------------------------------------------------------------------------------- |
| `task test_tutorial` | package (chapter) 名 | `tutorial` パッケージ配下の指定したチャプターのテストを実行します                        |
| `task test_curl`     | -                    | `curl` パッケージ配下のテストをすべて実行します                                          |
| `task test_curl_fn`  | 関数・メソッド名     | `curl` パッケージ配下のテストの内、テスト名が `<Args>` 一致するもののみ実行します        |
| `task test_cgrep`    | -                    | `cgrep` パッケージ配下のテストをすべて実行します                                         |
| `task test_cgrep_fn` | 関数・メソッド名     | `cgrep` パッケージ配下のテストの内、テスト名が `<Args>` 一致するもののみ実行します       |
| `task test_lt`       | -                    | `logtransfer` パッケージ配下のテストをすべて実行します                                   |
| `task test_lt_fn`    | 関数・メソッド名     | `logtransfer` パッケージ配下のテストの内、テスト名が `<Args>` 一致するもののみ実行します |

なお、 `cgrep` と `logtransfer` は非同期処理を用いる兼ね合いで、15 秒のタイムアウトを設けています。
正常に実装できていれば 15 秒を超えることは無いので、実装にミスが有ります。

## Build

`tutorial` 以外のパッケージは実装が完了したら実際にバイナリ化して使用することができます。
出力したバイナリに Path が通っていれば、他の CLI と同じように使用することができます。

各ツールのビルド用タスクおよびバイナリ名は以下のとおりです。

| Tool          | Task               | binary  |
| :------------ | :----------------- | :------ |
| `curl`        | `task build_curl`  | `scurl` |
| `cgrep`       | `task build_cgrep` | `cgrep` |
| `logtransfer` | `task build_lt`    | `lt`    |

バイナリは各ツールのルートディレクトリ（go.mod があるところ）に出力されます。

注）出力されるバイナリは実行環境に合わせてコンパイルされます
