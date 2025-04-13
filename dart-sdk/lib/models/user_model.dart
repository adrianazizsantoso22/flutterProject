import 'dart:convert';
import 'package:objectbox/objectbox.dart';

@Entity()
class UserModel {
  @Id()
  int id = 0; // ID pengguna
  String name; // Nama pengguna

  UserModel({required this.name}) {
    // Validasi saat inisialisasi
    if (id <= 0) {
      throw ArgumentError('ID harus lebih besar dari 0'); [[1]](https://poe.com/citation?message_id=375478100477&citation=1)
    }
    if (name.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong'); [[1]](https://poe.com/citation?message_id=375478100477&citation=1)
    }
  }

  // Metode untuk memperbarui nama pengguna
  void updateName(String newName) {
    if (newName.isEmpty) {
      throw ArgumentError('Nama tidak boleh kosong'); [[1]](https://poe.com/citation?message_id=375478100477&citation=1)
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
      throw ArgumentError('Data tidak valid'); [[1]](https://poe.com/citation?message_id=375478100477&citation=1)
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

  void main() {
    final user = UserModel(name: 'John Doe');
    print(user); // Output: UserModel{id: 0, name: John Doe} [[2]](https://poe.com/citation?message_id=375478100477&citation=2)

    user.updateName('Jane Doe');
    print(user); // Output: UserModel{id: 0, name: Jane Doe} [[2]](https://poe.com/citation?message_id=375478100477&citation=2)
  }
}
