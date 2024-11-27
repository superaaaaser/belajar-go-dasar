![tutorial go](../images/go-constant.png)

# Golang Constant

- Constant adalah variabel yang memiliki nilai tetap yang tidak dapat diubah
- Untuk membuat constant bisa menggunakan keyword `const`
- Sintaks untuk membuat constant

```go
const NAMACONSTANT tipe = value
```

- Value constant harus di tetapkan pada saat mendeklarasikannya
- Nama constant biasanya ditulis dengan huruf besar, agar mumeudahkan pemneda antara `constant` dan `variabel`
- Constant bisa di deklarasikan didalam maupun di luar fungsi

## Contoh membuat constan di Golang

```go
package main
import "fmt"

const LENGTH = 10
const WIDTH = 5

func main(){
    fmt.Println(LENGTH)
    fmt.Println(WIDTH)
}
```
