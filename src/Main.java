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