/*
This example illustrates the use of a recover statement to capture a panic.
Because deferred statements are executed after the function has completed
evaluation, they can be used to handle a panic statement using the built-in
recover function. This function will return the "value" of the panic as an
`any` value (or nil, if there was no panic). Once recover has been called,
the panic is effectively stopped, and the function will return zero values
for any of its return types.
*/
package main
