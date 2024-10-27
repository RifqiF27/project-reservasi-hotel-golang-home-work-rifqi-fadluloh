# Aplikasi Reservasi Hotel

Aplikasi Reservasi Hotel adalah aplikasi berbasis CLI (Command Line Interface) yang memungkinkan pengguna untuk mendaftar, login, dan melakukan reservasi kamar hotel. Aplikasi ini juga menyediakan fitur bagi admin untuk mengelola data kamar dan reservasi.

## Fitur

### Untuk Pengguna (Customer)

- **Registrasi**: Pengguna baru dapat mendaftar untuk membuat akun.
- **Login**: Pengguna dapat masuk ke akun mereka.
- **Reservasi Kamar**: Pengguna dapat melakukan reservasi kamar dengan memilih tanggal dan kamar yang tersedia.
- **Lihat Reservasi Saya**: Pengguna dapat melihat reservasi yang telah mereka buat.
- **Batalkan Reservasi**: Pengguna dapat membatalkan reservasi yang belum dikonfirmasi.
- **Melihat Kamar**: Pengguna dapat melihat semua kamar yang tersedia.

### Untuk Admin

- **Kelola Reservasi**: Admin dapat melihat semua reservasi, mengonfirmasi atau membatalkan reservasi.
- **Kelola Kamar**: Admin dapat menambahkan, mengedit, dan menghapus data kamar.
- **Lihat Semua Pelanggan**: Admin dapat melihat daftar semua pengguna terdaftar.

## Penggunaan

Setelah Anda menjalankan aplikasi, Anda akan disambut dengan menu utama. Pilih opsi berikut untuk melakukan tindakan:

1. **Registrasi**: Masukkan data diri Anda untuk membuat akun baru.
2. **Login**: Masukkan username dan password untuk masuk ke akun Anda.
3. **Reservasi Kamar**: ID kamar, input tanggal, input pembayaran dan lakukan reservasi.
4. **Lihat Reservasi Saya**: Melihat daftar reservasi yang telah Anda buat.
5. **Batalkan Reservasi**: Masukkan ID reservasi untuk membatalkan reservasi yang belum dikonfirmasi.

### Contoh Penggunaan

Berikut adalah contoh interaksi pengguna dengan aplikasi:

1. **Registrasi Pengguna**

```
Silakan masukkan username: John
Silakan masukkan password: *******
Silakan masukkan phone number: 021123123
Akun berhasil dibuat!
```

2. **Login Pengguna**

```
   Silakan masukkan username: John
   Silakan masukkan password: *******
   login berhasil
```

3. **Reservasi Kamar**

```
masukkan ID Kamar yang tersedia: 1
Masukkan tanggal check-in (YYYY-MM-DD): 2024-11-01
Masukan pembayaran: 200
Reservasi berhasil dibuat!
```

### Note

registrasi akun hanya sebagai customer.

```
Login untuk admin :
username: admin
password: admin123
```

```
Login untuk customer :
username: customer
password: customer123
```
