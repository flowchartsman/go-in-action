/*
This example illustrates how a panic unrolls the stack. If you run
it, you should see the stack trace at the site of the panic, with
the entries flowing backwards through the call stack until they
reach the call to a() at line 12 of main.main().
*/
package main
