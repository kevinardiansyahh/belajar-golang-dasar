package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//Append pada slice akan membuat array baru jika slice yg diappend capacitynya sudah penuh
	//jika terbentuk array baru & ada perubahan pada slice tsb, data yg berubah hanya akan
	//terjadi pada slice nya, tidak akan ada perubahan pada array awal
	//Tetapi jika pada saat append capacity slice masih ada, slice masih akan merefer ke array
	//yang sama (data pada array tetap akan berubah jika slice diubah)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)      // [senin selasa rabu kamis jum'at sabtu baru minggu baru]
	fmt.Println(daySlice1) // [sabtu baru minggu baru]

	//append slice capacity penuh
	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2) // [ups minggu baru libur baru]
	fmt.Println(days)      // [senin selasa rabu kamis jum'at sabtu baru minggu baru]

	// make() => membuat langsung slice baru, tanpa array
	sliceName := make([]string, 2, 5)
	sliceName[0] = "Kevin"
	sliceName[1] = "Ardiansyah"
	fmt.Println("sliceNama sebelum append : ", sliceName) // [Kevin Ardiansyah]
	// sliceName[2] = "Hidayat" -> will be error
	sliceName1 := append(sliceName, "Hidayat")
	fmt.Println("sliceNama setelah append : ", sliceName1) // [Kevin Ardiansyah Hidayat]
	fmt.Println(sliceName)

	//COPY SLICE
	copySlice := make([]string, len(sliceName), cap(sliceName))
	copy(copySlice, sliceName)
	fmt.Println(copySlice)
}
