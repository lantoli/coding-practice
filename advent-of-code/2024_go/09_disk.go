package main

// 6463499258318 6493634986625 (sample 1928 2858)
func main() {
	solve("09_input.txt", line09a, nil, solve09a, solve09b)
}

var (
	input09 = ""
)

func line09a(line string) {
	input09 = line
}

func solve09a() int {
	nums := make([]int, len(input09))
	for i, ch := range input09 {
		nums[i] = int(ch) - '0'
	}
	checksum, pos := 0, 0
	left, right := 0, len(nums)-1
	for left <= right {
		if left%2 == 0 {
			for range nums[left] {
				checksum += pos * (left / 2)
				pos++
			}
			left++
			continue
		}
		if right%2 == 1 {
			right--
			continue
		}
		for nums[left] > 0 && nums[right] > 0 {
			nums[left]--
			nums[right]--
			checksum += pos * (right / 2)
			pos++
		}
		if nums[left] == 0 {
			left++
		}
		if nums[right] == 0 {
			right--
		}
	}
	return checksum
}

func solve09b() int {
	disk := make([]int, 0, len(input09)*10)
	for i, ch := range input09 {
		id := -1
		if i%2 == 0 {
			id = i / 2
		}
		for range int(ch) - '0' {
			disk = append(disk, id)
		}
	}
	rcount, lcount := 0, 0
	for right := len(disk) - 1; right > 0; right -= rcount {
		rcount = 1
		if disk[right] == -1 {
			continue
		}
		for right-rcount >= 0 && disk[right] == disk[right-rcount] {
			rcount++
		}
	inner:
		for left := 0; left < right; left += lcount {
			lcount = 1
			if disk[left] != -1 {
				continue
			}
			for left+lcount < right && disk[left+lcount] == -1 {
				lcount++
			}
			if rcount <= lcount {
				for count := range rcount {
					disk[left+count] = disk[right-count]
					disk[right-count] = -1
				}
				break inner
			}
		}
	}
	checksum := 0
	for i, id := range disk {
		if id != -1 {
			checksum += i * id
		}
	}
	return checksum
}
