package paxosrpc

type RemotePaxos interface {
	//put methods here
	Prepare(*PrepareArgs, *PrepareReply) error
	Accept(*AcceptArgs, *AcceptReply) error
	Commit(*CommitArgs, *CommitReply) error
	GetLogs(*GetArgs, *GetReply) error
}

type Paxos struct {
	// Embed all methods into the struct.
	RemotePaxos
}

// Wrap wraps s in a type-safe wrapper struct to ensure that only the desired
// StorageServer methods are exported to receive RPCs.
func Wrap(s RemotePaxos) RemotePaxos {
	return &Paxos{s}
}
