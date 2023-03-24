# hello-golang

[A Tour of Go](https://go-tour-jp.appspot.com/) に沿って Go を学習する際に、より深く理解できる課題集です。下記一覧にある通り、 A Tour of Go の特定の範囲に対応する課題が当パッケージ配下に Chapter という形で用意されています。

各 Chapter は課題となる `function` が虫食い状態で用意されており、どのように実装すべきかを同じディレクトリ内の README に記載しています。また当該 `function` に対するユニットテストも実装されており、このテストをパスすると課題クリアとなります。

当リポジトリ上では同様のテストを Push 時に実行する CI も用意しておりますので、Push 権限のある方は実際に Push してテスト結果を共有しましょう！

## 課題一覧

| Chapter    | A Tour of Go                                                                                                                                                                           | Subject                                 |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------- |
| Chapter 01 | [Packages](https://go-tour-jp.appspot.com/basics/1)<span style="padding: auto 2rem">~</span>[Short variable declarations](https://go-tour-jp.appspot.com/basics/10)                    | [README](./subject/chapter01/README.md) |
| Chapter 02 | [Basic types](https://go-tour-jp.appspot.com/basics/11)<span style="padding: auto 2rem">~</span>[Numeric Constants](https://go-tour-jp.appspot.com/basics/16)                          | [README](./subject/chapter02/README.md) |
| Chapter 03 | [For](https://go-tour-jp.appspot.com/flowcontrol/1)<span style="padding: auto 2rem">~</span>[Stacking defers](https://go-tour-jp.appspot.com/flowcontrol/13)                           | [README](./subject/chapter03/README.md) |
| Chapter 04 | [Pointers](https://go-tour-jp.appspot.com/moretypes/1)<span style="padding: auto 2rem">~</span>[Struct Literals](https://go-tour-jp.appspot.com/moretypes/5)                           | [README](./subject/chapter04/README.md) |
| Chapter 05 | [Arrays](https://go-tour-jp.appspot.com/moretypes/6)<span style="padding: auto 2rem">~</span>[Range continued](https://go-tour-jp.appspot.com/moretypes/17)                            | [README](./subject/chapter05/README.md) |
| Chapter 06 | [Maps](https://go-tour-jp.appspot.com/moretypes/19)<span style="padding: auto 2rem">~</span>[Methods continued](https://go-tour-jp.appspot.com/methods/3)                              | [README](./subject/chapter06/README.md) |
| Chapter 07 | [Pointer receivers](https://go-tour-jp.appspot.com/methods/4)<span style="padding: auto 2rem">~</span>[Choosing a value or pointer receiver](https://go-tour-jp.appspot.com/methods/8) | [README](./subject/chapter07/README.md) |
| Chapter 08 | [Interfaces](https://go-tour-jp.appspot.com/methods/9)<span style="padding: auto 2rem">~</span>[The empty interface](https://go-tour-jp.appspot.com/methods/14)                        | [README](./subject/chapter08/README.md) |
| Chapter 09 | [Type assertions](https://go-tour-jp.appspot.com/methods/15)<span style="padding: auto 2rem">~</span>[Images](https://go-tour-jp.appspot.com/methods/24)                               | [README](./subject/chapter09/README.md) |

※ A Tour of Go の内、Exercise は対象外です。
