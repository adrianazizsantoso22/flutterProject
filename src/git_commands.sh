#!/bin/bash

# Memeriksa apakah direktori adalah repositori Git
if [ -d ".git" ]; then
  echo "Direktori ini sudah menjadi repositori Git."
else
  echo "Direktori ini bukan repositori Git. Menginisialisasi repositori..."
  git init
fi

# Menambahkan semua file ke staging area
git add .

# Meminta pesan komit dari pengguna
read -p "Masukkan pesan komit awal: " commit_message

# Membuat commit awal
git commit -m "$commit_message"

# Meminta URL repositori GitHub
read -p "Masukkan URL repositori GitHub: " repo_url

# Menambahkan remote repository
git remote add origin "$repo_url"

# Mengirim commit ke branch master
git push -u origin master || { echo "Gagal mengirim ke remote. Periksa koneksi Anda."; exit 1; }

# Mengatur file .gitignore
echo "build/" >> .gitignore
echo ".idea/" >> .gitignore
echo ".dart_tool/" >> .gitignore
echo "pubspec.lock" >> .gitignore
echo ".DS_Store" >> .gitignore  # Untuk pengguna macOS
echo "*.log" >> .gitignore       # Menambahkan file log
echo "*.tmp" >> .gitignore       # Menambahkan file sementara

# Menambahkan file .gitignore ke staging area
git add .gitignore

# Membuat branch baru untuk pengembangan
git checkout -b development

# Mengirim branch baru ke remote
git push -u origin development || { echo "Gagal mengirim branch 'development' ke remote."; exit 1; }

# Menampilkan pesan sukses
echo "Repositori berhasil diinisialisasi dan branch 'development' sudah dipush ke remote."