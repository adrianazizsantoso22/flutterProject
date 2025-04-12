#!/bin/bash

# Memeriksa apakah Flutter terinstal
if ! command -v flutter &> /dev/null
then
    echo "Flutter tidak ditemukan. Pastikan Flutter telah terinstal dan ada di PATH."
    exit 1
fi

# Meminta nama proyek dari pengguna
read -p "Masukkan nama proyek (default: flutter_project): " PROJECT_NAME
PROJECT_NAME=${PROJECT_NAME:-flutter_project}  # Default jika tidak ada input

# Meminta template proyek dari pengguna
echo "Pilih template proyek:"
echo "1) Aplikasi Flutter biasa"
echo "2) Aplikasi Plugin"
read -p "Masukkan pilihan (1/2, default: 1): " TEMPLATE_CHOICE
TEMPLATE_CHOICE=${TEMPLATE_CHOICE:-1}

# Membuat proyek Flutter baru
if [ "$TEMPLATE_CHOICE" -eq 2 ]; then
    flutter create --template=plugin $PROJECT_NAME
else
    flutter create $PROJECT_NAME
fi

# Masuk ke direktori proyek
cd $PROJECT_NAME

# Menginstal dependensi yang ada di pubspec.yaml
flutter pub get

# Membuat folder untuk model, widget, dan layanan
mkdir -p lib/models
mkdir -p lib/widgets
mkdir -p lib/services
mkdir -p lib/screens  # Folder untuk layar

# Membuat file awal untuk model, widget, dan layanan
touch lib/models/example_model.dart
touch lib/widgets/example_widget.dart
touch lib/services/api_service.dart
touch lib/screens/home_screen.dart  # Contoh file layar

# Membuat file .gitignore
echo "build/" >> .gitignore
echo ".idea/" >> .gitignore
echo ".dart_tool/" >> .gitignore
echo "pubspec.lock" >> .gitignore
echo ".DS_Store" >> .gitignore  # Untuk pengguna macOS

# Menginisialisasi repositori Git
git init
git add .
git commit -m "Initial project setup"

# Menambahkan remote repository
read -p "Masukkan URL remote repository (atau tekan Enter untuk melewati): " REPO_URL
if [ -n "$REPO_URL" ]; then
    git remote add origin "$REPO_URL"
    echo "Remote repository ditambahkan."
else
    echo "Anda bisa menambahkan remote repository nanti dengan perintah: git remote add origin <URL>"
fi

# Menampilkan pesan sukses
echo "Proyek Flutter '$PROJECT_NAME' berhasil dibuat dan diinisialisasi!"

# Menampilkan petunjuk untuk langkah selanjutnya
echo "Langkah selanjutnya:"
echo "- Buka proyek di editor pilihan Anda."
echo "- Tambahkan dependensi tambahan di pubspec.yaml jika diperlukan."
echo "- Jalankan proyek dengan 'flutter run' di direktori proyek."