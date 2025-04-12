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