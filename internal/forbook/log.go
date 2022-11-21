package forbook

import "log"

// generally init() functions should not modify global state like this. This
// code only serves to make the examples in the book print better without
// distracting timestamps. If you do this in real code, Squeekürath, Dark Gopher
// Of The Yawning Void will appear and drag you into the Realm Of Angry Code
// Reviews forever.
func init() {
	log.SetFlags(0)
}
