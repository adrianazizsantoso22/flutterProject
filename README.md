# Flutter Project

oleh: Adrian Aziz Santoso (NRP 5025221229)
untuk memenuhi tugas: Pemrograman Perangkat Bergerak (B)

Proyek ini merupakan aplikasi Flutter yang mmenggunakan Dart SDK dan ObjectBox sebagai database lokal. Aplikasi ini dibuat untuk memenuhi tugas mata kuliah Pemrograman Perangkat Bergerak untuk mengeksplorasi teknologi _mobile database_, dan tak kalah pentingnya adalah untuk

## Fitur utama
- **CRUD Operations**: Menambahkan, mengambil, dan menghapus data pengguna.
- **ObjectBox Integration**: Menggunakan ObjectBox untuk penyimpanan data yang efisien dan cepat.

## Langkah pengerjaan _Flutter project_

![image](https://github.com/user-attachments/assets/82c9d8ec-9a44-49ec-8ed2-4fb5cd26a376)

1. Pertama-tama, pastikan kita telah mengunduh Android Studio. Jika belum, unduhlah Android Studio melalui [tautan berikut](https://developer.android.com/studio?gad_source=1&gbraid=0AAAAAC-IOZnVixm20JUfP0gddBMfqrT83&gclid=Cj0KCQjwnui_BhDlARIsAEo9GutPFzO4kVmx5lW25nYRad8nXDGhMet_m04X3o8KFmk99KhDkrO5ThcaAhQXEALw_wcB&gclsrc=aw.ds)
2. Lalu, pastikan kita juga telah mengunduh Visual Studio Code. Jika belum, unduhlah Visual Studio Code melalui [tautan berikut](https://code.visualstudio.com/Download)
3. Kemudian, pastikan kita juga telah mengunduh Visual Studio Installer. Jika belum, unduhlah Visual Studio Installer melalui [tautan berikut](https://code.visualstudio.com/Download)
4. Pada Visual Studio Installer, unduhlah Visual Studio Community 2022 Preview sembari meng-_install workload_ berikut: a) _ASP.NET and web development_, b) _.NET desktop development_, c) _WinUI application development_, dan d) _Data storage and processing_.
5. Selanjutnya, unduh Dart-SDK dalam bentuk .zip melalui [tautan berikut](https://dart.dev/get-dart/archive)
6. Juga, jangan lupa unduh Flutter dalam bentuk .zip melalui [tautan berikut](https://docs.flutter.dev/get-started/install)
7. Berikutnya, ekstrak Dart-SDK dan Flutter yang sudah diunduh lalu letakkan _folder_ ```dart-sdk``` maupun _folder_ ```flutter``` pada ```C:\Users\<username>\StudioProjects\<nama Flutter Project>```
8. Carilah menu ```Edit the system environment variables``` pada menu ```Start```, lalu klik ```Environment variables``` lalu tambahkan ```C:\Users\<username>\StudioProjects\<nama Flutter Project>\flutter\bin``` dan ```C:\Users\<username>\StudioProjects\<nama Flutter Project>\dart-sdk" pada variabel PATH```
9. Bukalah pengaturan pada Android Studio lalu pada bagian ```Languages & Frameworks```, atur ```Dart SDK Path``` pada ```C:\Users\<username>\StudioProjects\<nama Flutter Project>\flutter```


![image](https://github.com/user-attachments/assets/e21fee64-c505-43a7-8cc2-146f04ace57a)

![image](https://github.com/user-attachments/assets/1aa1d6ee-0295-4c62-8ffc-f24176d8766e)


10. Masih pada bagian  ```Languages & Frameworks```, kini atur ```Flutter Path``` pada ```C:\Users\<username>\StudioProjects\<nama Flutter Project>\flutter```
11. Buatlah sebuah _Flutter project_ di mana Anda membuat kodingan berikut pada _folder_ `src`:
11a. git_commands.sh
```git_commands.sh
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
```
11b. Main.java
```Main.java
import java.util.Scanner;
import java.util.Calendar;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String continueInput;

        do {
            System.out.println("Hello, World!");

            // Meminta nama pengguna
            System.out.print("Masukkan nama Anda: ");
            String name = scanner.nextLine();
            System.out.println("Selamat datang, " + name + "!");

            // Meminta usia pengguna dengan validasi
            int age = -1;
            while (age < 0) {
                System.out.print("Berapa usia Anda? ");
                if (scanner.hasNextInt()) {
                    age = scanner.nextInt();
                    if (age < 0) {
                        System.out.println("Usia tidak bisa negatif. Coba lagi.");
                    }
                } else {
                    System.out.println("Silakan masukkan angka yang valid.");
                    scanner.next(); // Menghapus input yang tidak valid
                }
            }
            System.out.println("Usia Anda adalah " + age + " tahun.");

            // Menghitung tahun lahir
            int currentYear = Calendar.getInstance().get(Calendar.YEAR);
            int birthYear = currentYear - age;
            System.out.println("Anda lahir pada tahun " + birthYear + ".");

            // Memberikan rekomendasi berdasarkan usia
            if (age < 13) {
                System.out.println("Anda masih kecil usianya.");
            } else if (age < 18) {
                System.out.println("Anda masih berusia remaja.");
            } else if (age < 25) {
                System.out.println("Anda sudah berusia dewasa muda.");
            } else if (age < 60) {
                System.out.println("Anda sudah berusia dewasa.");
            } else {
                System.out.println("Anda sudah tua.");
            }

            // Menghitung dan menampilkan jumlah hari hidup lebih akurat
            long daysLived = calculateDaysLived(birthYear, currentYear);
            System.out.println("Anda telah hidup selama sekitar " + daysLived + " hari.");

            // Menanya apakah pengguna ingin memasukkan data lagi
            System.out.print("Apakah Anda ingin memasukkan data lagi? (ya/tidak): ");
            scanner.nextLine(); // Mengonsumsi newline
            continueInput = scanner.nextLine();

        } while (continueInput.equalsIgnoreCase("ya"));

        // Menutup scanner
        scanner.close();
        System.out.println("Terima kasih telah menggunakan program ini!");
    }

    // Metode untuk menghitung jumlah total hari hidup
    private static long calculateDaysLived(int birthYear, int currentYear) {
        long totalDays = 0;
        for (int year = birthYear; year < currentYear; year++) {
            totalDays += isLeapYear(year) ? 366 : 365; // Menambahkan hari sesuai tahun kabisat
        }
        return totalDays;
    }

    // Metode untuk memeriksa apakah tahun adalah tahun kabisat
    private static boolean isLeapYear(int year) {
        return (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0);
    }
}
```
11c. main.dart
```main.dart
import 'package:flutter/material.dart';
import 'package:hive/hive.dart';
import 'package:path_provider/path_provider.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final directory = await getApplicationDocumentsDirectory();
  Hive.init(directory.path);

  // Membuka box Hive untuk menyimpan data
  var box = await Hive.openBox('myBox');

  runApp(MyApp(box: box));
}

class MyApp extends StatelessWidget {
  final Box box;

  MyApp({required this.box});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'My Flutter App',
      home: HomeScreen(box: box),
    );
  }
}

class HomeScreen extends StatefulWidget {
  final Box box;

  HomeScreen({required this.box});

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final TextEditingController _controller = TextEditingController();
  String _greeting = 'Tidak ada pesan';

  @override
  void initState() {
    super.initState();
    _loadGreeting();
  }

  // Memuat pesan dari Hive
  void _loadGreeting() {
    try {
      final greeting = widget.box.get('greeting', defaultValue: 'Tidak ada pesan');
      setState(() {
        _greeting = greeting;
      });
    } catch (e) {
      print('Error loading greeting: $e');
    }
  }

  // Menyimpan pesan ke Hive
  void _saveGreeting() {
    try {
      widget.box.put('greeting', _controller.text);
      _loadGreeting(); // Memuat ulang pesan setelah menyimpan
      ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Pesan disimpan
```
11d. pubspec.yaml
```pubspec.yaml
name: nama_proyek
description: Aplikasi Flutter yang menggunakan Hive dan state management.
version: 1.0.0+1

environment:
  sdk: ">=2.12.0 <3.0.0"
  flutter: ">=2.0.0"

dependencies:
  objectbox: ^2.0.0
  objectbox_flutter_libs: ^2.0.0

  # UI
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.2 # Ikon untuk iOS

  # Database
  hive: ^2.0.0
  hive_flutter: ^2.0.0
  sqflite: ^2.0.0 # Untuk database SQL

  # State Management
  provider: ^6.0.0 # Untuk state management

  # HTTP
  http: ^0.13.0 # Untuk melakukan request HTTP

  # Penyimpanan Sederhana
  shared_preferences: ^2.0.0 # Untuk menyimpan data sederhana

  # Notifikasi
  flutter_local_notifications: ^9.0.0 # Untuk notifikasi lokal

  # Gambar dan Media
  image_picker: ^0.8.0 # Untuk mengambil gambar dari galeri atau kamera
  cached_network_image: ^3.0.0 # Untuk caching gambar dari jaringan

  # Animasi
  flutter_animate: ^2.0.0 # Untuk animasi

  # Internasionalisasi
  flutter_localizations:
    sdk: flutter # Untuk dukungan internasionalisasi

dev_dependencies:
  flutter_test:
    sdk: flutter
  mockito: ^5.0.0 # Untuk mocking dalam pengujian
```
11e. setup_project.sh
```setup_project.sh
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
```
11f. user_model.dart
```user_model.dart
import 'dart:convert'; // Mengimpor pustaka untuk JSON

class UserModel {
  final int id; // ID pengguna
  String name; // Nama pengguna

  UserModel(this.id, this.name) {
    // Validasi saat inisialisasi
    if (id <= 0) {
      throw ArgumentError('ID harus lebih besar dari 0');
    }
    if (name.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong');
    }
  }

  // Metode untuk memperbarui nama pengguna
  void updateName(String newName) {
    if (newName.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong');
    }
    name = newName; // Memperbarui nama
  }

  // Mengonversi objek menjadi map
  Map<String, dynamic> toMap() {
    return {
      'id': id, // Menyimpan ID
      'name': name, // Menyimpan nama
    };
  }

  // Membuat objek UserModel dari map
  factory UserModel.fromMap(Map<String, dynamic> map) {
    if (map['id'] == null || map['name'] == null) {
      throw ArgumentError('Data tidak valid');
    }
    return UserModel(
      map['id'] as int, // Mengambil ID dari map
      map['name'] as String, // Mengambil nama dari map
    );
  }

  // Serialisasi objek ke format JSON
  String toJson() {
    return jsonEncode(toMap()); // Mengonversi ke string JSON
  }

  // Deserialisasi objek dari format JSON
  factory UserModel.fromJson(String json) {
    final Map<String, dynamic> data = jsonDecode(json); // Mengonversi string JSON ke map
    return UserModel.fromMap(data); // Menggunakan fromMap untuk validasi
  }

  @override
  String toString() {
    return 'UserModel{id: $id, name: $name}'; // Mengembalikan representasi string dari objek
  }

  // Metode untuk membuat salinan objek dengan perubahan
  UserModel copyWith({int? id, String? name}) {
    return UserModel(
      id ?? this.id,
      name ?? this.name,
    );
  }
}
```

![image](https://github.com/user-attachments/assets/fbefd834-b2e0-43fd-8f10-c642e788dbac)

12. Bukalah Terminal lalu jalankan perintah ```flutter doctor``` untuk memastikan bahwa semuanya berjalan dengan baik.
13. Lalu, jalankan perintah ```cd C:\Users\<username>\StudioProjects\<nama Flutter Project>\flutter\bin```
14. Berikutnya, jalankan perintah ```flutter pub add hive_flutter:^1.1.0```

![image](https://github.com/user-attachments/assets/fc013b7f-c2d9-4a40-88d8-6f947d2a3f42)

15. Selanjutnya, jalankan perintah ```flutter pub get```
16. Buat _file_ baru di _folder_ `dart-sdk/lib/models` bernama `user_model.dart` dengan isi sebagai berikut:
```dart
import 'dart:convert';
import 'package:objectbox/objectbox.dart';

@Entity()
class UserModel {
  int id = 0; // ID pengguna
  String name; // Nama pengguna

  UserModel({required this.name}) {
    // Validasi saat inisialisasi
    if (id <= 0) {
      throw ArgumentError('ID harus lebih besar dari 0');
    }
    if (name.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong');
    }
  }

  // Metode untuk memperbarui nama pengguna
  void updateName(String newName) {
    if (newName.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong');
    }
    name = newName; // Memperbarui nama
  }

  // Mengonversi objek menjadi Map
  Map<String, dynamic> toMap() {
    return {
      'id': id, // Menyimpan ID
      'name': name, // Menyimpan nama
    };
  }

  // Membuat objek UserModel dari Map
  factory UserModel.fromMap(Map<String, dynamic> map) {
    if (map['id'] == null || map['name'] == null) {
      throw ArgumentError('Data tidak valid');
    }
    return UserModel(
      name: map['name'] as String, // Mengambil nama dari map
    )..id = map['id'] as int; // Mengatur ID
  }

  // Serialisasi objek ke format JSON
  String toJson() {
    return jsonEncode(toMap()); // Mengonversi ke string JSON
  }

  // Deserialisasi objek dari format JSON
  factory UserModel.fromJson(String json) {
    final Map<String, dynamic> data = jsonDecode(json); // Mengonversi string JSON ke map
    return UserModel.fromMap(data); // Menggunakan fromMap untuk validasi
  }

  @override
  String toString() {
    return 'UserModel{id: $id, name: $name}'; // Mengembalikan representasi string dari objek
  }

  // Metode untuk membuat salinan objek dengan perubahan
  UserModel copyWith({int? id, String? name}) {
    return UserModel(
      name: name ?? this.name,
    )..id = id ?? this.id;
  }
}
```
17. Untuk menghasilkan kode ObjectBox:, jalankan perintah `flutter pub run build_runner build` di Terminal.
18. Agar dapat menginisialisasi ObjectBox, ubahlah kode di `main.dart` menjadi:
```dart
import 'package:flutter/material.dart';
import 'package:objectbox/objectbox.dart';
import 'objectbox.g.dart'; // file yang dihasilkan

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final store = await openStore(); // Membuka store ObjectBox
  runApp(MyApp(store: store));
}

class MyApp extends StatelessWidget {
  final Store store;

  MyApp({required this.store});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'My ObjectBox App',
      home: HomeScreen(store: store),
    );
  }
}

class HomeScreen extends StatelessWidget {
  final Store store;

  void addUser(String name) {
    final user = UserModel(name: name);
    store.box<UserModel>().put(user); // Menyimpan data
  }

  List<UserModel> getUsers() {
    return store.box<UserModel>().getAll(); // Mengambil semua data
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('ObjectBox Example')),
      body: Center(child: Text('Hello, ObjectBox!')),
    );
  }
}
```

19. Untuk memastikan semuanya berfungsi dengan baik, jalankan _Flutter project_ kita dengan menjalankan perintah `flutter run`

### Deskripsi _database_
Aplikasi ini menggunakan ObjectBox untuk menyimpan informasi pengguna. Model pengguna didefinisikan dalam _file_ ```user_model.dart```.

### Referensi
- [ObjectBox Documentation](https://objectbox.io/docs/)
- [Flutter Documentation](https://docs.flutter.dev/)
