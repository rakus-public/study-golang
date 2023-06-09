# IC カードと改札

乗車料金と残高とポイントが個別に記録されている IC カードを渡すと、チャージ残高に応じて乗車可否を判定する改札関数を実装せよ。

また乗車可能な場合のみカードからチャージ残高を減らす処理も実装すること

## 前提条件

### 乗車料金

- `charge > 0`

### IC カードの内容

- `Balance >= 0`
- `Point >= 0`

お金の残高を表す `Balance` と、ポイント残高を表す `Point` 2 つのフィールドを持つ構造体として定義しています。

どちらも `int` 型の値を入れることが可能です。

### 乗車可否の判定基準

`Balance` + `Point` が乗車料金以上あれば乗車可とします。

乗車可の場合は戻り値を `true` に、乗車不可の場合は `false` になるように実装してください。

### 残高の減算処理

乗車可能だった場合、IC カードから残高を乗車料金分だけ減算処理を行います。

減算処理は `Point` から優先して差し引かれ、償却しきれない分は `Balance` から差し引かれるようにしてください。

## Tips

### 加減算代入

以下のように、ある変数の現在の値を加減算した上で変数に入れ直したい場合

```go
var n = 10
n = n + 3
// => n == 13
```

代入と計算を一気に行えるショートハンドが以下の通りです。

```go
var n = 10
n += 3
fmt.Println(n) // => 13
```

`for` 構文で出てきた `i++` というのもこれの応用で、加減算する数値が `1` の場合のみ加減算記号を重ねるこのような書き方ができます。

```go
var n = 10

n++
fmt.Println(n) // => 11

n--
fmt.Println(n) // => 10
```

## テスト実行コマンド

```bash
task test_tutorial -- chapter04
```
