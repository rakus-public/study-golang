# カーシェアサービスの利用料金計算

カーシェアサービスの利用時間と利用した車のグレードに基づいて料金を算出せよ。

## 前提条件

### 利用時間

`minutes > 0`

引数 `minutes` として利用した時間が `int` 型で渡されてきます。

単位は分です。

### グレード

以下の 3 種類のグレードを計算できるように実装してください

- `Basic`
- `Middle`
- `Premium`

### 料金表

各グレードにおいて、 `15分ごとの料金` と `6時間以内最大料金` が設定されています。

各料金は税込みとし、税率計算は不要です。

|                      |    Basic     |    Middle    |   Premium    |
| :------------------: | :----------: | :----------: | :----------: |
|   15 分ごとの料金    | 220 円/15 分 | 330 円/15 分 | 440 円/15 分 |
| 最大料金(6 時間まで) |   4,290 円   |   6,490 円   |   8,690 円   |

#### 15 分ごとの料金

15 分ごとに加算される料金です。

次のような要件が存在します。

- 15 分を超えると次の 15 分分の料金が加算される
  - `e.g.) 40分 => 15分ごとの料金 * 3`
- 6 時間未満かつ最大料金を超える場合は最大料金を適応する
  - `e.g.) Basic 330 分 => ✕4840 ◯4290`
- 6 時間を超える度に計算がリセットされる
  - `e.g.) 380分 => 最大料金 * 1 + 15分ごとの料金 * 2`

#### 最大料金

6 時間ごとに加算される料金です。

利用時間が 6 時間に到達するごとに時間計算がリセットされ、その地点から新たな最大料金適応時間が開始されます。

`e.g.) 735分 => 最大料金 * 2 + 15分`

### 実装方法

各グレードを `type Car interface` を満たす構造体になるようにメソッドを実装し、 `func Calc(car Car, minutes int) int` 内で単一のロジックで各グレードの料金を正しく計算できるように実装してください。

## Tips

各グレードを表す構造体にフィールドは不要です。

## テスト実行コマンド

```bash
task test -- chapter08
```