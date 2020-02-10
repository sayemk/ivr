package ivr

type NodeInputValidator interface {
	Validate(node *Node) bool
	InvalidMessage(node *Node) string //will return a sound file
}
type Reply struct {
	Res int    //Numeric result of the AGI command.
	Dat string //Additional returned data.
}
type AgiClient interface {
	Answer() (Reply,error)
	ChannelStatus(channel ...string) (Reply, error)
	ControlStreamFile(file, escape string, params ...interface{}) (Reply, error)
	GetOption(filename, escape string, timeout ...int) (Reply, error)
	Hangup(channel ...string) (Reply, error)
	StreamFile(file, escape string, offset ...int) (Reply, error)
	WaitForDigit(timeout int) (Reply, error)
} 


