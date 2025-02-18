# Uninitialized Map Access in Go

This repository demonstrates a common error in Go: accessing a key in an uninitialized map.  Attempting to assign a value to a key in an uninitialized map will *not* initialize the map; instead, it will result in a runtime panic.  This example highlights the importance of initializing maps before use.

## How to Reproduce

1. Clone this repository.
2. Run `go run bug.go`.
3. Observe the runtime panic.

## Solution

The `bugSolution.go` file provides a corrected version, showing the proper way to initialize a map before using it.
