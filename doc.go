// Package test is a package that provides equality functions for testing Go
// code. It is focused on making the most common action in test code simple,
// testing equality expectations.
//
// A simple test function looks like this:
//
//     func TestAbs(t *testing.T) {
//         test.Eq(t, Abs(-1), 1)
//     }
//
// If the check passes, the verbose output looks like this:
//
//     --- PASS: TestAbs (0.00s)
//         test.go:2: test.Eq(t, Abs(-1), 1): got 1
//
// If the check fails, and the type is a bool, int, or float, the output looks
// like this:
//     --- PASS: TestAbs (0.00s)
//         test.go:2: test.Eq(t, Abs(-1), 1): got 0, want 1
//
// If the check fails, and the type is a string, the output looks like this:
//
//     --- FAIL: TestGoGophers (0.00s)
//         test.go:2: test.Eq(t, "Golang\nGophers", "Go\nGophers"):
//         --- got
//         +++ want
//         @@ -1,2 +1,2 @@
//         -Golang
//         +Go
//          Gophers
//
// If the check fails, and the type is an array or slice, the output looks like
// this:
//
//     --- FAIL: TestGoGophers (0.00s)
//         test.go:2: test.Eq(t, []string{"Golang", "Gophers"}, []string{"Go", "Gophers"}):
//         --- got
//         +++ want
//         @@ -1,5 +1,5 @@
//          ([]string) (len=2) {
//         - (string) (len=6) "Golang",
//         + (string) (len=2) "Go",
//           (string) (len=7) "Gophers"
//          }
//
// If the check fails, and the type is a struct value, the output looks like
// this:
//
//     --- FAIL: TestGoGophers (0.00s)
//         test.go:2: test.Eq(t, struct{Name string; Age int}{"A", 44}, struct{Name string; Age int}{"a", 44}):
//         --- got
//         +++ want
//         @@ -1,5 +1,5 @@
//          (struct { Name string; Age int }) {
//         - Name: (string) (len=1) "A",
//         + Name: (string) (len=1) "a",
//           Age: (int) 44
//          }
//
// Got and want
//
// The terms got and want are used to describe what you got as a result of
// running the code, and what you want to have gotten. These terms are commonly
// found in the Go stdlib and it's own testing docs. In some other testing
// libraries they are sometimes referred to actual and expected.
//
// Nesting
//
// Checks can be nested using the bool return value of a prior check.
//
//     func TestAbs(t *testing.T) {
//         if test.Eq(t, Abs(-1), 1) {
//             ...
//         }
//     }
//
// Breaking early
//
// Checks can cause a test to stop at a failure using the bool return value.
//
//     func TestAbs(t *testing.T) {
//         if !test.Eq(t, Abs(-1), 1) {
//             return
//         }
//         ...
//     }
//
// Comparison
//
// Comparison of values is performed using Google's cmp Go module:
// https://github.com/google/go-cmp/cmp
//
package test
