# Go: Unexpected Variable Reassignment

This repository demonstrates a common, yet subtle, error in Go: unintended variable reassignment.  The example showcases a simple integer reassignment, but the implications can become more significant with pointers or complex data structures.  The `bug.go` file shows the problematic code, while `bugSolution.go` offers a corrected version and explanation.

The core issue highlights the importance of understanding how variable assignment works in Go, especially within the context of functions and memory management.  Careful consideration of variable scope and potential side effects is crucial for writing robust and predictable Go programs.

## Learning Points

* **Variable Reassignment:**  Understanding how and when variables are reassigned is key to avoiding unexpected behavior.
* **Pointers and Memory:**  Reassignment with pointers can have significant consequences and requires extra care.
* **Data Structures:**  Unexpected behavior can occur when reassigning components of more complex data structures like slices or maps.