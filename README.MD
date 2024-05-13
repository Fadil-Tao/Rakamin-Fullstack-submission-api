# FINAL TAS SUBMISSION RAKAMIN PBI 

Membuat REST A.P.I menggunakan Bahasa Pemrograman Go. 

## Functional Requirement 
1. User dapat menambahkan foto profile
2. Sistem dapat mengidentifikasi User ( log in / sign up)
3. Hanya user yang telah login / sign up yang dapat melakukan
delete / tambah foto profil
4. User dapat menghapus gambar yang telah di post
5. User yang berbeda tidak dapat menghapus / mengubah foto yang
telah di buat oleh user lain 

## Technologi yang digunakan 
Dibuat menggunakan bahasa pemrograman Go , web framework gin , GORM dan JWT

## Struktur Project 
- CMD/ API , digunakan untuk menampung router , middleware dan file main go yang akan menajadi gerbang utama untuk menjalankan aplikasi
- Models , digunakan untuk menyimpan file model sebagai blueprint dari objek
- Controllers, sebagai handler dari user dan photos

## Prequisite 
- Go ver 1.22.3
- Postgresql

 ## Cara install
 - Clone 
``` git clone https://github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah.git ```

- Setelah itu buat file .env yang mengandung beberapa variabel dibawah 
``` 
DBHOST=localhost
USER=
PASSWORD=
DB_NAME=
DBPORT=
JWT_SECRET= 
```

## API ENDPOINT
- Register/signup 
POST ```/v1/pbiapi/users/register```
- Login
POST ```/v1/pbiapi/users/login```
- Get User Info
GET ```/v1/pbiapi/users/details/:userid```
- Update User Info
POST  ```/v1/pbiapi/users/update/:userid```
- Delete User
DELETE  ```/v1/pbiapi/users/delete/:userid```
- Logout / clear cookie
POST  ```/v1/pbiapi/users/logout```
-  Get Photo by id
GET ```/v1/pbiapi/photos/getPhoto/:photoid```
- Add Photo
POST ```/v1/pbiapi/photos/post```
-   Update photo
PUT ```/v1/pbiapi/photos/update/:photoid```
-   Delete Photo
DELETE ```/v1/pbiapi/photos/delete/:photoid```