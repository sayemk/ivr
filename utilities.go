package ivr

type DefaultValidator struct {
	message string
}

func (valid *DefaultValidator) Validate(node *Node) bool  {
	return true
}

func (valid *DefaultValidator) InvalidMessage(node *Node) string  {
	return valid.message
}
