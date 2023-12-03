/*
This example illustrates the necessity of calling recover() in a deferred
function. Because panics occur after the code which generates them, and
because defers are guaranteed to be run as the stack unrolls, they are the
only sensible place to recover from a panic.
*/
package main
