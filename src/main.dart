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
