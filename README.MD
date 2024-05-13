# FINAL TASK SUBMISSION RAKAMIN PBI 

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
JSON yang akan dikirim
```
{
  "email":"foo@gmail.com",
  "username":"foo",
  "password":"bar123"
}
```

- Login
POST ```/v1/pbiapi/users/login```
JSON yang akan dikirim
```
{
  "email": "foo@gmail.com",
  "password":"bar123"
}
```

- Get User Info
GET ```/v1/pbiapi/users/details/:userid```
Contog Penggunaan
```
http://localhost:8080/v1/pbiapi/users/details/2
```

- Update User Info
PUT  ```/v1/pbiapi/users/update/:userid```
JSON yang akan dikirim :
```
{
  "email":"foo@gmail.com",
  "username":"foo",
  "password":"bar123"
}
```
- Delete User
DELETE  ```/v1/pbiapi/users/delete/:userid```
Contog Penggunaan : 
```
http://localhost:8080/v1/pbiapi/photos/delete/3
```

- Logout / clear cookie
POST  ```/v1/pbiapi/users/logout```
Contog Penggunaan : 
```
http://localhost:8080/v1/pbiapi/users/logout
```

-  Get Photo by id
GET ```/v1/pbiapi/photos/getPhoto/:photoid```
Contoh Penggunaan:
```
http://localhost:8080/v1/pbiapi/photos/getphoto/2
```

- Add Photo
POST ```/v1/pbiapi/photos/post```
JSON Yang akan dikirim :
```
{
  "title" : "Foo",
  "caption" : "Bar",
  "url": "https://placebear.com/500/300"
}
```


-   Update photo
PUT ```/v1/pbiapi/photos/update/:photoid```
JSON Body yang akan dikirim :
```
{
  "title":"foo",
  "caption" : "bar",
  "url":"https://placebear.com/200/300"
}
```

-   Delete Photo
DELETE ```/v1/pbiapi/photos/delete/:photoid```
Contoh Penggunaan :
```
http://localhost:8080/v1/pbiapi/photos/update/2
```