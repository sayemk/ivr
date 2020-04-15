package ivr

import (
	"time"
)

const (
	DTMF_ANY = "0123456789*#"
	DTMF_STAR = "*"
	DTMF_HASH = "#"
	DTMF_NUMERIC = "1234567890"
	DTMF_NONNUMERIC = "*#"
	DTMF_0 = 0
	DTMF_1 = 1
	DTMF_2 = 2
	DTMF_3 = 3
	DTMF_4 = 4
	DTMF_5 = 5
	DTMF_6 = 6
	DTMF_7 = 7
	DTMF_8 = 8
	DTMF_9 = 9
	DTMF_NONE = ""
	STATE_NOT_RUN = 1
	STATE_CANCEL = 2
	STATE_COMPLETE = 3
	STATE_TIMEOUT = 4
	STATE_MAX_INPUTS_REACHED = 5
	INPUT_END = 1
	INPUT_CANCEL = 2
	INPUT_NORMAL = 3
	TIME_INFINITE = -1
)
/*
* Node type Declaration
*/
type Node struct {
	client *AgiClient
	input string
	state int
	endOfInputDigit string
	cancelDigit string
	minInput int
	maxInput int
	timeBetweenDigits time.Duration
	totalTimeForInput time.Duration
	promptMessage string
	prePromptMessage string
	prePromptMessagesInterruptable bool
	acceptPrePromptInputAsInput bool
	name string
	inputValidator NodeInputValidator
	maxAttemptsForInput int
	playOnMaxValidInputAttempts string
	promptsCanBeInterrupted bool
	validInterruptDigits string
	executeOnValidInput func(node *Node)
	executeOnInputFailed func(node *Node)
	cancelWithInputRetriesInput bool
	inputAttemptsUsed int
	executeBeforeRun func(node *Node)
	executeAfterRun func(node *Node)
	executeAfterFailedValidation func(node *Node)
	playOnNoInputInLastAttempt bool
	playOnNoInputMessage string
	playOnNoInputInLastAttemptMessage string
}
/*
* New Node creation
* @return Node ptr 
*/

func NewNode(client *AgiClient) *Node {
	return &Node{
		client: client,
		input:DTMF_NONE,
		state:STATE_NOT_RUN,
		endOfInputDigit: "X",
		cancelDigit: "X",
		minInput:1,
		maxInput:1,
		timeBetweenDigits:300*time.Millisecond,
		totalTimeForInput:30*time.Second,
		prePromptMessagesInterruptable: true,
		acceptPrePromptInputAsInput:true,
		name:DTMF_NONE,
		inputValidator: new(DefaultValidator),
		maxAttemptsForInput:DTMF_1,
		promptsCanBeInterrupted:true,
		validInterruptDigits:DTMF_ANY,
		inputAttemptsUsed:0,
		playOnNoInputInLastAttempt:true,



	}
}



func (node *Node) SetPrePromptMessagesInterruptable( status bool)  {
	node.prePromptMessagesInterruptable = status
}

func (node *Node) SetAcceptPrePromptInputAsInput(status bool)  {
	node.acceptPrePromptInputAsInput = status
}

func (node *Node) PlayOnNoInput(filename string)  {
	node.playOnNoInputMessage = filename
}

func (node *Node) PlayNoInputMessageOnLastAttempt(filename string)  {
	node.playOnNoInputInLastAttemptMessage = filename
}

func (node *Node) PlayOnMaxInputAttemptsReached(filename string)  {
	node.playOnMaxValidInputAttempts = filename
}


func (node *Node) MaxAttemptsForInput(attempts int)  {
	node.maxAttemptsForInput = attempts
}

func (node *Node) ClearPromptMessages()  {
	node.playOnNoInputInLastAttemptMessage = ""
	node.playOnMaxValidInputAttempts = ""
	node.playOnNoInputMessage = ""

}

func (node *Node) AddPrePromptMessage(filename string)  {
	node.prePromptMessage = filename
}





