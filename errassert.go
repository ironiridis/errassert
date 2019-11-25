package errassert

import "fmt"

// Must checks that a task didn't fail, and aborts if it does.
func Must(task string, err error) {
	if err == nil {
		return
	}
	fmt.Printf("❌ failed to %s: %s\n", task, err)
	panic(err)
}

// Verbose calls Must, and if it returns, logs the successful task.
func Verbose(task string, err error) {
	Must(task, err)
	fmt.Printf("✅ %s\n", task)
}
