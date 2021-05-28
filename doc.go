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
// If the check fails, and the type is a string, array, slice, or complex type,
// the output looks like this:
//
//     --- FAIL: TestAbs (0.00s)
//         test.go:2: test.Eq(t, Abs(-1), 1): int(
//         -: 0
//         +: 1
//         )
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
