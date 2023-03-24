---
marp: true
size: 16:9
paginate: true
headingDivider: 1
style: |
  div.title {
    text-align: center;
    font-size: 60px;
    color: #66CCFF;
    font-weight: bold;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  div.author {
    text-align: right;
    font-size: 40px;
    color: #ffffff;
    font-weight: bold;
    position: absolute;
    bottom: 20%;
    right: 3%;
  }
  section h1 {
    font-size: 50px;
    color: #66CCFF;
  }
  section h2 {
    color: #ffffff;
  }
  section li {
    color: #ffffff;
  }
  section p {
    color: #ffffff;
  }
---

<div class="title">
  Go勉強会ガイダンス
</div>

<div class="author">
  インフラ開発部 SRE課（今本/上村)
</div>

![bg brightness:0.3](./img/gopher.png)

# Go 言語の特徴

- Go 言語は Google が開発したプログラミング言語
- 文法がシンプル、高速、並行処理が得意などの特徴を持つ

![bg brightness:0.3](./img/gopher.png)

# 当リポジトリが想定するゴール

- アプリケーション開発に詳しくない人が参加することを想定した内容です。
  - ただし「変数」「関数」等、ごく初歩的なプログラミング用語はある程度知っている前提です（must ではない）
- 勉強会の開催期間は半年で、前後半に分けてテーマを変更します。
- 前半：基礎編
  - 演習問題を各自で解きながら Go の基本的な文法を習得します。
  - 詳しい進め方については次のスライドから説明します。
- 後半：実践編
  - 週替わりで実装担当者を決めて、Go の CLI ツールを 1 つ完成させることで実装力を高めます。
  - 完成させるツールの詳細については検討中のため、このガイダンス内での説明はありません。

![bg brightness:0.3](./img/gopher.png)

# 進め方について

当リポジトリは 4 つの章で構成されています

1. Tutorial
2. 機能限定版 `curl` アプリ実装
3. ファイルの内容一致検索アプリ実装
4. ログ転送アプリ実装

![bg brightness:0.3](./img/gopher.png)

# Tutorial の進め方

- チャプターごとに対応する範囲の A Tour of Go を完了させる
  - 範囲は `tutorial` 配下の README に掲載
- 演習問題を各自で実装する
  - そのチャプターまでで出てきた知識のみでクリア可能になっている
  - 各自でブランチを切って作業する
  - 未実装になっているメソッドを実装する形をとる
- チャプター内のテストがすべて通ればクリア
- （Push 権限があれば）GitHub に Push して、CI をパスする

![bg brightness:0.3](./img/gopher.png)

# その他ルール

- Push する際のブランチ名は`講座開始月/自身の名前`としてください。
  - 識別できれば良いので、名前部分は GitHub のアカウント名とかで大丈夫です
- 解答を載せたブランチ（`feature/subject-examples`）がありますが、事前の閲覧は控えてください。
  - 演習問題を自力で解いてもらいたいため

![bg brightness:0.3](./img/gopher.png)

# 環境構築

- 開発スタイルによって環境構築作業が異なります。
  - ローカルの実行環境に Go をインストールして開発する場合
  - VSCode の DevContainers 拡張機能を使って Docker コンテナ上で開発する場合

![bg brightness:0.3](./img/gopher.png)

# DevContainers とは

- Docker コンテナを起動させて、コンテナ上にあるソースコードを直接編集する VSCode の機能
- Docker コンテナ起動時にローカル環境のソースコードをマウントさせることで、ローカル環境のファイルとの同期も取られる

![bg brightness:0.3](./img/gopher.png)

# DevContainers の利点

- 開発環境を Dockerfile や docker-compose.yml に明記できるので、開発に必要なツールがインストールされた状態で開発を始められる
  - 例えば今回の場合、Go の実行環境やタスクランナー(go-task)、開発用ツールなどが事前にインストール済み
- devcontainer.json の値を書き換えれば、vscode の設定や拡張機能も共通化できるようになる
  - 開発環境を統一できる
- 開発環境がコンテナ内で完結するので、ローカル環境を汚さずに開発ができる

![bg brightness:0.3](./img/gopher.png)

# 事前準備について

- docker コマンドを実行する環境を準備済みの前提です

![bg brightness:0.3](./img/gopher.png)

# ローカルに Go をインストールして開発する場合

## 必要な作業

1. ローカル環境への Go のインストールする - [手順](https://go.dev/doc/install)
   a. Linux(ログインシェル bash)の例
   `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz`
   `vi ~/.bashrc`で"export PATH=$PATH:/usr/local/go/bin"を追記
   `source ~/.bashrc`
   `go version`
2. ローカル環境へのタスクランナー(go-task)インストールする - [手順](https://taskfile.dev/installation/#go-modules)
   `go install github.com/go-task/task/v3/cmd/task@latest`

![bg brightness:0.3](./img/gopher.png)

# ローカルに Go をインストールして開発する場合

## 必要な作業

3. [リポジトリ](https://github.com/rakus-public/study-golang)を clone
   最終ページに注意事項あり
4. clone したプロジェクトのルートディレクトリでローカル環境でテストを実行してみる
   `task test_tutorial -- chapter00`

![bg brightness:0.3](./img/gopher.png)

# VSCode の DevContainers 拡張機能を使って Docker コンテナ上で開発する場合

## 必要な作業

1. [リポジトリ](https://github.com/rakus-public/study-golang)を clone
   最終ページに注意事項あり
1. VSCode をインストールし、clone したリポジトリを開く
1. Remote Development 拡張を VSCode にインストール
1. VSCode の左下の `><` アイコンをクリックし `Reopen in Container` を選択
1. コンテナ上でソースコードが閲覧・編集できることを確認
1. VSCode のターミナル上からテストを実行してみる
   `task test_tutorial -- chapter00`

※ DevContainer に Go の実行環境が内包されているので、Go のインストールは不要

![bg brightness:0.3](./img/gopher.png)

# リモートリポジトリの clone に関する注意点

グローバルの Git ユーザーがこのリポジトリのユーザーと異なる場合、プッシュできなくなっていたりコミットログが荒れたりしてしまいます。
以下の手順で clone して勉強会用リポジトリの設定を追加してください。

1. Personal Access Token を`Settings`>`Developer Settings`>`Personal access tokens`>`Generate new token`で生成
1. `git clone https://[アカウント名]:[Personal Access Token]@github.com/kurupeku/hello-golang`で clone
1. `cd hello-golang`
1. `git config --local user.name [アカウント名]`
1. `git config --local user.email [メールアドレス]`
   `Settings`>`Emails`で`Keep my email addresses private`を有効化し、記載されたメアドを入力

![bg brightness:0.3](./img/gopher.png)

# CI の実行（要 Push 権限）

1. リモートリポジトリを clone
   次ページに注意事項あり
1. 自分用のブランチを作成
1. VSCode 上で DevContainer を開く
1. サンプル問題を実装する
1. ローカル環境でテストがパスすることを確認
1. リモートリポジトリに push
1. GitHub 上で CI が通っていることを確認

![bg brightness:0.3](./img/gopher.png)

# はじめよう

1. `chapter00` の関数をコメントどおりに編集し、テストをパスさせてください
2. 無事にパスしたら A Tour of Go を進めましょう

![bg brightness:0.3](./img/gopher.png)
