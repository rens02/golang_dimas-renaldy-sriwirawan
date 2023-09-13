# Set variabel folder utama dan buat folder utama
main_folder="$1_$at$(date +'%a %b %d %T %Z %Y')"
mkdir "$main_folder"

# Masuk ke folder utama
cd "$main_folder"

# Buat folder about me
mkdir about_me
# Masuk folder about_me
cd about_me

# Buat folder personal dan professional
mkdir personal
mkdir professional

# Keluar dari folder about me
cd ..

# Buat folder my_friends dan my_system_info
mkdir my_friends
mkdir my_system_info

# Masukkan url ke file masing masing
echo "https://www.facebook.com/$2" > "./about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "./about_me/professional/linkedin.txt"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "./my_friends/list_of_my_friends.txt"

# Memasukkan username dan info device ke dalam file masing masing
echo "My username: $USER" > "./my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "./my_system_info/about_this_laptop.txt"

# Memasukkan hasil dari ping google ke file
ping -c 3 google.com > "./my_system_info/internet_connection.txt"

tree

echo "Script telah selesai dijalankan."

