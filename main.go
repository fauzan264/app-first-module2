package main

import (
	"fmt"

	go_first_module "github.com/fauzan264/go-first-module/v2"
)

/**
Sebelum Ada Go Modules
- Saat kita membuat aplikasi, biasanya kita akan menggunakan library atau dependency dari project lain.
- Sebelum ada Go Modules, management untuk dependency sangan sulit dilakukan di golang
- Biasanya kita akan meng-copy semua code libary atau dependency lain ke project kita, hal ini membuat project kita menjadi bengkak karena penuh dengna library orang lain.
- Atau biasanya library orang lain akan kita save di GOPATH directory. namun hal ini akan sulit jika ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda.

Go-Modules
- Go-Modules mulai dikenalkan di Golang 1.11 and 1.12
- Dengan Go-Modules, kita bisa membuat project dengan mudah dan dependency management yang sangat mudah
*/

/**
Membuat Module
- Untuk membuat module baru, kita bisa menggunakan perintah: go mod init module-name
- Go-lang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan.

Rilis Module
- Go-Lang terintegrasi baik dengan Git
- Untuk merilis module, kita hanya perlu membuat Tag di Git
*/

// membuat module tidak bisa di private

/**
Upgrade Module
- Untuk melakukan upgrade module, kita hanya cukup membuat tag baru di Git.
*/

/**
Upgrade Dependency
- Untuk upgrade dependency ke versi terbaru, kita bisa menguubah isi go.mod, lalu mengubah tag nya menjadi tag terbaru
- Untuk mendownload versi terbaru, gunakan perintah: go get
- Kalau versinya hilang dan salah memasukan versinya maka buildnya juga pasti error, karena tidak tersedia pada versi tersebut.
*/

/**
Major Upgrade
- Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode program kita, sehingga membuatnya tidak backward compatible.
- Sebaiknya hal ini sebisa mungkin dihindari.
- Namun jika tidak bisa dihindari, strategi terbaik adalah merubah nama module.
- Kalau merubah major changes/upgrade, maka lebih baik meelakukan perubahan nama module seperti menambahkan /new_version
*/

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(go_first_module.First(" Fauzan"), " ", i)
	}
}