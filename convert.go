package main

import(
    "os"
    "fmt"
    "bufio"
    "image"
    _ "image/jpeg"
    "image/png"
)

func main() {
    // 入力画像パス.
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Input source image file path  >>")
    if !scanner.Scan() {
        fmt.Println("Please input source image file path.")
        return;
    }
    srcPath := scanner.Text()

   // ファイルオープン
    file, err := os.Open(srcPath)
    assert(err, "Invalid image file path " + srcPath)
    defer file.Close()

    // ファイルオブジェクトを画像オブジェクトに変換
    img, _, err :=  image.Decode(file)
    assert(err, "Failed to convert file to image.")

    // 出力画像パス.
    fmt.Print("Input output image file path  >>")
    if !scanner.Scan() {
        fmt.Println("Please input output image file path.")
        return;
    }
    dstPath := scanner.Text()

    // 出力ファイルを生成
    out, err := os.Create(dstPath)
    assert(err, "Failed to create destination path.")
    defer out.Close()

    // 画像ファイル出力
//    jpeg.Encode(out, img, nil)
    png.Encode(out, img)

}

// errorオブジェクトをチェックし、nilの場合例外を送出
func assert(err error, msg string) {
    if err != nil {
        panic(err.Error() + ":" + msg)
    }
}