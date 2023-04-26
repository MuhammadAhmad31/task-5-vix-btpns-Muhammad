# task-5-vix-btpns-Muhammad
Tugas Virtual Internship Task akhir VIX Full Stack Developer ini kalian diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Pada kasus ini, kalian diinstruksikan untuk membuat API untuk mengupload dan menghapus gambar. API yang kalian bentuk adalah POST, GET, PUT, dan DELETE.

Ketentuan API Pada bagian User :

Endpoint :
1. POST : /users/register, dan gunakan atribut berikut ini
a. ID (primary key, required)
b. Username (required)
c. Email (unique & required)
d. Password (required & minlength 6)
e. Relasi dengan model Photo (Gunakan constraint cascade)
f. Created At (timestamp)
g. Updated At (timestamp)
2. GET: /users/login
a. Using email & password (required)
3. PUT : /users/:userId (Update User)
4. DELETE : /users/:userId (Delete User)

Photos Endpoint :
1. POST : /photos
b. ID
c. Title
d. Caption
e. PhotoUrl
f. UserID
g. Relasi dengan model User
2. GET : /photos
3. PUT : /photoId
4. DELETE : /:photoId

Requirement :

1. Authorization dapat menggunakan tool Go JWT ○ https://github.com/dgrijalva/jwt-go
2. Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto
