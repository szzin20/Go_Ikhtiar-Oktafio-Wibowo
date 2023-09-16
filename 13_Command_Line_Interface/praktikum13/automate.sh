#!/bin/bash

# Periksa jumlah argumen
if [ "$#" -ne 3 ]; then
	echo "Error: You must provide exactly three arguments."
	echo "Usage: $0 <folder_name> <facebook_url> <linkedin_url>"
	exit 1
fi

# Proses argumen
folder_name="$1"
facebook_url="$2"
linkedin_url="$3"

current_date=$(date "+%Y-%m-%d %H:%M:%S")

# Buat folder utama dengan nama yang sesuai
main_folder="$folder_name at $current_date"
mkdir -p "$main_folder"

# Masuk ke folder utama yang sudah dibuat
cd "$main_folder"

# Buat folder about_me dan subfoldernya (personal dan professional)
mkdir -p "about_me/personal"
mkdir -p "about_me/professional"

# Buat folder my_friends
mkdir -p "my_friends"

# Buat folder my_system_info
mkdir -p "my_system_info"

# Membuat file facebook.txt dan linkedin.txt di folder about_me dengan argumen yang ditentukan
echo "https://www.facebook.com/$facebook_url" > "about_me/personal/facebook.txt"
echo "https://www.linkedin.com/$linkedin_url" > "about_me/professional/linkedin.txt"

# Mengisi file list_of_my_friends.txt di folder my_friends denan data sesuai
cat <<EOF > "my_friends/list_of_my_friends.txt"
1) Achmad Miftahul - amulum
2) Achmad Rafiq - achmadrafiq97
3) Adiststi - adiststi
4) Agung - agungajin19
5) Azzahra - al7262
6) Charisma - fadzrichar
7) Farida - ulfarida
8) Garry - garryarielcussoy
9) Hamdi - hamdiranu
10) Hedy Gading - Gading09
11) Ilham - aamfatur
12) Lelianto - Lelianto
13) Mohammad - daffa99
14) Muhammad Fadhil - fabdulkarim
15) Ofbimon - bimon-alta
16) Purnama Syafitri - pipitmnr
17) Setyo - setyoyo07
18) Wildan - wiflash
19) Willy - sumarnowilly94
20) Woka - woka20
EOF

# Mengisi file about_this_laptop.txt dengan data yang sesuai di folder my_system_info
about_this_laptop_info="User: $(whoami)\n$(uname -a)"
echo -e "$about_this_laptop_info" > "my_system_info/about_this_laptop.txt"

# Mengisi file internet_connection.txt di folder my_system_info sebanyak 3 kali
ping -n 3 google.com > "my_system_info/internet_connection.txt"

echo "Skrip otomatis sukses"
