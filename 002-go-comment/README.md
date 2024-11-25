![Go Comment](../images/go-comment.png)

# Komentar di Go

`Komentar` adalah teks yang ditulis di dalam kode tetapi diabaikan oleh compiler atau interpreter saat program dijalankan. Komentar digunakan untuk memberikan penjelasan tentang kode, catatan, atau dokuemntasi yang membantu programmer (termasuk diri sendiri di masa depan) untuk memahami logika atau tujuan dari bagian tertentu dari kode.

## Pentingnya komentar

- Membantu kolaborasi tim dalam memahami kode
- Meningkatkan pemeliharaan kode jangka panjang
- Mengurangi kebingungan saat membaca kode kompleks

## Jenis Komentar

### Komentar satu baris

    Program pertama yang akan dibuat adalah program Hello World, sebuah program yang sangat terkenal di kalangan programmer. Program tersebut bertutujan untuk menampilkan teks "Hello world".

    ```go
    package main
    import "fmt"

    func main(){
        fmt.Println("Hello World")
    }
    ```

## Komentar banyak baris

- `package main` merupakan program yang ada di package main.
- `import "fmt"` yaitu menyertakan file yang ada di package fmt.
- `func main() {}` yaitu sebuah fungsi, kode yang akan di jalankan berada dalam kurung kurawal.
- `fmt.Println()` adalah sebuah fungsi yang berada di package fmt, yang berfungsi untuk mencetak teks.
