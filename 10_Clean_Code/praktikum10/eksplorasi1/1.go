package main

type user struct {
	id       int         // Kekurangan 1: Nama field 'username' dan 'password' seharusnya berupa string, bukan int
	username int         // Alasan: username biasanya berupa string, bukan angka
	password int         // Alasan: password biasanya berupa string, bukan angka
}

type userservice struct {
	t []user           // Kekurangan 2: Struktur data 't' seharusnya memiliki nama yang lebih deskriptif
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}        // Kekurangan 3: Seharusnya mengembalikan nilai error jika user tidak ditemukan
}
// 4. Ubah tipe data username dan password dari int menjadi string jika mereka akan berisi teks atau karakter.

// 5. Beri nama yang lebih deskriptif pada field t dalam userservice, misalnya users.

// 6. Pada fungsi getuserbyid, sebaiknya kembalikan nilai error jika user dengan id yang dicari tidak ditemukan. 
	//Ini memungkinkan pemanggil fungsi untuk menangani kasus di mana user tidak ditemukan.