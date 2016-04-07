package main

type State struct {
	followers map[string]map[string]bool
}

func (s State) follow(from, to string) {
	m := s.followers[to]
	if m == nil {
		m = make(map[string]bool)
		s.followers[to] = m
	}
	m[from] = true
}

func (s State) unfollow(from, to string) {
	if m := s.followers[to]; m != nil {
		m[from] = false
	}
}

func (s State) isFollower(followee, follower string) bool {

	if m := s.followers[followee]; m != nil {
		return m[follower]
	}
	return false
}

func NewState() State {
	return State{make(map[string]map[string]bool)}
}

/*
Follow: Only the To User Id should be notified
Unfollow: No clients should be notified
Broadcast: All connected user clients should be notified
Private Message: Only the To User Id should be notified
Status Update: All current followers of the From User ID should be notified
*/
func (s State) Process(msg Message) Predicate {

	switch msg.Type {
	case "F":
		s.follow(msg.From, msg.To)
	case "U":
		s.unfollow(msg.From, msg.To)
	}

	return s.IsRecipient(msg)
}

type Predicate func(string) bool

func (s State) IsRecipient(msg Message) Predicate {

	return func(peerId string) bool {

		switch msg.Type {
		case "F", "P":
			return peerId == msg.To
		case "S":
			return s.isFollower(msg.From, peerId)
		case "B":
			return true

		default:
			return false
		}
	}
}
