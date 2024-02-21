package main

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxVolume := 0

	for left < right {
		volume := (right - left) * min(height[left], height[right])
		maxVolume = max(maxVolume, volume)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolume
}
