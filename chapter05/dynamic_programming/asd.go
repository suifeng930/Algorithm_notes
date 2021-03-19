package main

import "log"

func main() {

	nums := []int{0}
	number := missingNumber(nums)
	log.Println(number)
	parkingSystem := Constructor(1, 1, 0)
	log.Println(parkingSystem.AddCar(1))
	log.Println(parkingSystem.AddCar(2))
	log.Println(parkingSystem.AddCar(3))
	log.Println(parkingSystem.AddCar(1))

}

func missingNumber(nums []int) int {

	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := right + left // 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

type ParkingSystem struct {
	bigCar   int
	midCar   int
	smallCar int
}

func Constructor(big int, medium int, small int) ParkingSystem {

	return ParkingSystem{
		bigCar:   big,
		midCar:   medium,
		smallCar: small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {

	switch carType {
	case 1:
		if this.bigCar == 0 {
			return false
		}
		this.bigCar = this.bigCar - 1
		return true
	case 2:
		if this.midCar == 0 {
			return false
		}
		this.midCar = this.midCar - 1
		return true
	case 3:
		if this.smallCar == 0 {
			return false
		}
		this.smallCar = this.smallCar - 1
		return true
	default:
		return false
	}
}
