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