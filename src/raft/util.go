package raft

import "log"

// Debugging
const Debug = 0 //

func Log(rf *Raft, event, action string) {
	var stateStr string
	switch rf.state {
	case Candidate:
		stateStr = "Candidate"
	case Follower:
		stateStr = "Follower"
	case Leader:
		stateStr = "Leader"
	}
	format := "%d server, State: %s, Term: %d, Event: %s, Action: %s"
	DPrintf(format, rf.me, stateStr, rf.term, event, action)
}

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}
