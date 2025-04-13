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
      name: map['name'] as String,
    )..id = map['id'] as int; // Mengatur ID
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
      name: name ?? this.name,
    )..id = id ?? this.id;
  }
}

void main() {
  // Contoh penggunaan UserModel
  try {
    // Membuat objek UserModel
    UserModel user = UserModel(name: "John Doe");
    user.id = 1; // Mengatur ID
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
