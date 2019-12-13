# 如何新增icon

## 執行方法

在https://github.com/fyne-io/fyne/tree/master/cmd/fyne下有個工具
go build之後，有個執行檔
可以參考官方的[範例](https://github.com/fyne-io/examples/tree/develop/img/icon)

```bash
#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package icon -name life life.svg > $FILE
$BIN/fyne bundle -package icon -append -name lifeBitmap life.png >> $FILE

$BIN/fyne bundle -package icon -append -name bugBitmap bug.png >> $FILE
$BIN/fyne bundle -package icon -append -name calculatorBitmap calculator.png >> $FILE
$BIN/fyne bundle -package icon -append -name clockBitmap clock.png >> $FILE
$BIN/fyne bundle -package icon -append -name fractalBitmap fractal.png >> $FILE
$BIN/fyne bundle -package icon -append -name solitaireBitmap solitaire.png >> $FILE
$BIN/fyne bundle -package icon -append -name textEditorBitmap texteditor.png >> $FILE
$BIN/fyne bundle -package icon -append -name xkcdBitmap xkcd.png >> $FILE
```

或是可以直接執行

```bash
$./fyne bundle -package icon -name logo JLOGO.png > bundled.go
```

## 理解

實際上產生出來的檔也是讀了照片的二進位資料之後，寫進bundled.go
bundle.go裡面有fyne.StaticResource的struct，定義為

```go
type StaticResource struct {
    StaticName    string
    StaticContent []byte
}
```

StaticName是要在資源內的唯一名稱
StaticContent就是二進位資料了
