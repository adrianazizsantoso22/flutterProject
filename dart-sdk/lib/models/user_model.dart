import 'dart:convert';
import 'package:objectbox/objectbox.dart';

@Entity()
class UserModel {
  int id; // ID pengguna
  String name; // Nama pengguna

  UserModel({required this.id, required this.name}) {
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
      'id': id,
      'name': name,
    };
  }

  // Membuat objek UserModel dari Map
  factory UserModel.fromMap(Map<String, dynamic> map) {
    if (map['id'] == null || map['name'] == null) {
      throw ArgumentError('Data tidak valid');
    }
    return UserModel(
      id: map['id'] as int, // Mengambil ID dari map
      name: map['name'] as String,
    );
  }

  // Serialisasi objek ke format JSON
  String toJson() {
    return jsonEncode(toMap());
  }

  // Deserialisasi objek dari format JSON
  factory UserModel.fromJson(String json) {
    final Map<String, dynamic> data = jsonDecode(json);
    return UserModel.fromMap(data);
  }

  @override
  String toString() {
    return 'UserModel{id: $id, name: $name}';
  }

  // Metode untuk membuat salinan objek dengan perubahan
  UserModel copyWith({int? id, String? name}) {
    return UserModel(
      id: id ?? this.id,
      name: name ?? this.name,
    );
  }
}

void main() {
  // Contoh penggunaan UserModel
  try {
    // Membuat objek UserModel dengan ID yang valid
    UserModel user = UserModel(id: 1, name: "John Doe");
    print(user.toString()); // Menampilkan detail pengguna

    // Memperbarui nama pengguna
    user.updateName("Jane Doe");
    print("Nama baru: ${user.name}");

    // Serialisasi ke JSON
    String json = user.toJson();
    print("JSON: $json");

    // Deserialisasi dari JSON
    UserModel newUser = UserModel.fromJson(json);
    print("Pengguna baru: ${newUser.toString()}");
  } catch (e) {
    print("Terjadi kesalahan: $e");
  }
}
