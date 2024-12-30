# **RomanRepository Module**

## Deskripsi

Modul `roman-repository` adalah sebuah modul di Go yang menyediakan repository generik untuk menangani operasi CRUD (Create, Read, Update, Delete) pada berbagai entitas model menggunakan GORM. Modul ini mendukung berbagai jenis database, seperti MySQL, PostgreSQL, SQLite, dan lain-lain.

## Fitur

- Menyediakan interface generik untuk operasi database dasar.
- Mendukung berbagai jenis model yang dapat digantikan dengan tipe data apapun.
- Memudahkan interaksi dengan database tanpa menulis ulang kode untuk setiap model.
- Mendukung operasi CRUD standar: `FindAll`, `FindByID`, `Save`, dan `DeleteByID`.

---

## **Instalasi**

Untuk menggunakan modul ini, Anda perlu menginstalnya di proyek Go Anda.

### **Langkah 1: Menambahkan Modul ke Proyek Anda**

Tambahkan modul `roman-repository` ke proyek Go Anda dengan menjalankan perintah berikut:

```bash
go get github.com/romanmufid16/roman-repository
```

### **Langkah 2: Instal GORM dan Driver Database**

Modul ini bergantung pada GORM sebagai ORM dan membutuhkan driver database yang sesuai. Misalnya, jika Anda menggunakan MySQL, jalankan perintah berikut untuk menginstal GORM dan driver MySQL:

```bash
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

Jika Anda menggunakan database lain, cukup ganti driver sesuai dengan database yang digunakan, misalnya `postgres`, `sqlite`, atau `sqlserver`.

---

## **Penggunaan**

Berikut adalah contoh cara menggunakan `roman-repository` dalam proyek Anda.

### **Langkah 1: Menyiapkan Model**

Buat model yang ingin Anda gunakan dengan repository. Berikut adalah contoh model `User`:

```go
package main

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string
}
```

### **Langkah 2: Menyiapkan Koneksi Database**

Koneksi ke database menggunakan GORM. Berikut adalah contoh untuk mengonfigurasi koneksi ke MySQL:

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
    dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to database:", err)
        return nil
    }
    return db
}
```

### **Langkah 3: Menggunakan `RomanRepository`**

Setelah koneksi database siap, Anda dapat membuat dan menggunakan repository generik.

```go
package main

import (
    "fmt"
    "github.com/username/roman-repository"
    "gorm.io/gorm"
)

func main() {
    // Koneksi ke database
    db := ConnectDB()
    if db == nil {
        return
    }

    // Auto-migrasi untuk model User
    db.AutoMigrate(&User{})

    // Membuat instance RomanRepository untuk model User
    userRepo := NewRomanRepository[User](db)

    // Menyimpan data user baru
    user := User{Name: "John", Email: "john@example.com"}
    err := userRepo.Save(&user)
    if err != nil {
        fmt.Println("Failed to save user:", err)
        return
    }
    fmt.Println("User saved:", user)

    // Mengambil semua user dari database
    users, err := userRepo.FindAll()
    if err != nil {
        fmt.Println("Failed to fetch users:", err)
        return
    }
    fmt.Println("Users:", users)
}
```

---

## **Metode Repository**

### **1. `FindAll`**
Mengambil semua entitas dari database.

```go
FindAll() ([]T, error)
```

### **2. `FindByID`**
Mencari entitas berdasarkan ID.

```go
FindByID(id uint) (*T, error)
```

### **3. `Save`**
Menyimpan entitas ke dalam database (Create atau Update).

```go
Save(entity *T) error
```

### **4. `DeleteByID`**
Menghapus entitas berdasarkan ID.

```go
DeleteByID(id uint) error
```

---