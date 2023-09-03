**Goroutines** : unit pemrosesan ringan yang dapat dijalankan secara konkuren di dalam program Go.
Membuat goroutine sangat mudah, cukup dengan menambahkan kata kunci go sebelum pemanggilan fungsi yang ingin dieksekusi secara konkuren.
Goroutines sangat efisien dalam penggunaan sumber daya dan dapat digunakan untuk mengeksekusi banyak tugas secara bersamaan.

**Channels** : struktur data yang digunakan untuk berkomunikasi antara goroutines. Channels digunakan untuk mengirim dan menerima data antara goroutines dengan aman, menghindari perlombaan kondisi dan deadlock. Mekanisme channel memungkinkan sinkronisasi yang efisien antara goroutines, memfasilitasi koordinasi tugas konkuren.

**Concurrency Patterns** : Go mendukung berbagai pola konkurensi seperti Worker Pool, Fan-out Fan-in, dan Mutex untuk mengelola akses bersama ke data. Pola konkurensi membantu dalam mengatur dan mengoptimalkan eksekusi goroutines, memungkinkan peningkatan kinerja program dan efisiensi penggunaan sumber daya.