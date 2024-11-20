# Program Pertama

Pada kali ini kita akan membuat program pertama di bahasa pemrograman go.

## Membuat program hello world

Program pertama yang akan dibuat adalah program Hello World, sebuah program yang sangat terkenal di kalangan programmer. Program tersebut bertutujan untuk menampilkan teks "Hello world".

```go
package main
import "fmt"

func main(){
    fmt.Println("Hello World")
}
```

## Penjelasan kode diatas

- `package main` merupakan program yang ada di package main
- `import "fmt"` yaitu menyertakan file yang ada di package fmt
- `func main() {}` yaitu sebuah fungsi, kode yang akan di jalankan berada dalam kurung kurawal
- `fmt.Println()` adalah sebuah fungsi yang berada di package fmt, yang berfungsi untuk mencetak teks
