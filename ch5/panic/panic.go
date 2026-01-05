package main

func retWithPanic() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = 10
		}
	}()
	panic("trigger recover")
}
