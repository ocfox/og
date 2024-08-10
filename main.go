package main

func main() {
	var null *int
	func() {
		_ = *null
	}()
}
