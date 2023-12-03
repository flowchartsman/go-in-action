/*
This example illustrates the panic resulting from the assignment to a nil map.

Nil maps are a bit of a tricky situation in Go. They are likely the number one
source of panics amongst new and experienced developers alike; however they are
not that much different from nil slices. Both of them will panic if assigned to
directly, however slices at least have the advantage of using an assignment-based
append workflow, which makes them easier to work with when they are nil.

Unlike slices, however, nil maps are easier to read from, since they will always
return a zero value.
*/
package main
