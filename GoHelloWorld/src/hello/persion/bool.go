package main

func main() {
	
}
// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
// itob报告是否为非零。
func itob(i int) bool { return i != 0 }