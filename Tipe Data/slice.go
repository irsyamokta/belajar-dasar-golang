package main

import "fmt"

func main() {
	
    // Membuat slice
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice awal:", nums)

    // Menambah elemen
    nums = append(nums, 6, 7)
    fmt.Println("Setelah append:", nums)

    // Mengambil subslice
    subslice := nums[2:5]
    fmt.Println("Subslice:", subslice)

    // Mengubah elemen
    nums[3] = 10
    fmt.Println("Setelah diubah:", nums)

    // Menyalin slice
    copySlice := make([]int, len(nums))
    copy(copySlice, nums)
    fmt.Println("Hasil copy:", copySlice)

}