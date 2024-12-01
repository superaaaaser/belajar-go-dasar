![Go Variabel](../images/go-variabel.png)

# Variabel di Golang

- variabel dalah tempat untuk menyimpan value data
- untuk membuat variable di golang ada dua cara yaitu dengan keyword var atau dengan tanda :=
- Syntax dengan keyword var: var namaVariabel tipe = value
- Syntax dengan tanda := : namaVariabel := value

## Perbedaan Antara var dan :=

Ada sedikit perbedaan antara var dan :=
| var | := |
|:--: |:--:|
| Data baris 1 | Data baris 1 |
| Data baris 2 | Data baris 2 |
| Data baris 3 | Data baris 3 |

## Deklarasi variabel tanpa nilai awal

```go
package main
import "fmt"

func main(){
    var a string
    var i int
    var u bool

    fmt.Println(a)
    fmt.Println(i)
    fmt.Println(u)
}
```

## Deklarasi Multiple Variabel

```go
package main
import "fmt"

func main(){
    var d, e, f int = 2, 3, 4

    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
}
```
