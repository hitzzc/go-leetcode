package remove_element

func removeElement(nums []int, val int) int {
    l := 0
    for _, v := range nums {
        if v != val {
            nums[l] = v
            l++
        }
    }

    return l
}
