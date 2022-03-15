package discussion

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidID    = errors.New("invalid discussion ID")
	ErrInvalidType  = fmt.Errorf("discussion type is invalid, supported types are: %s", strings.Join(SupportedTypes, ","))
	ErrInvalidState = fmt.Errorf("discussion state is invalid, supported states are: %s", strings.Join(SupportedStates, ","))
)

type NotFoundError struct {
	CommentID    string
	DiscussionID string
}

func (e NotFoundError) Error() string {
	fields := []string{"could not find"}
	if e.DiscussionID != "" {
		fields = append(fields, fmt.Sprintf(" discussion with id \"%s\"", e.DiscussionID))
	}
	if e.CommentID != "" {
		fields = append(fields, fmt.Sprintf(" and comment with id \"%s\"", e.CommentID))
	}
	return strings.Join(fields, " ")
}

type InvalidError struct {
	CommentID    string
	DiscussionID string
}

func (e InvalidError) Error() string {
	fields := []string{"invalid"}
	if e.DiscussionID != "" {
		fields = append(fields, fmt.Sprintf("discussion id \"%s\"", e.DiscussionID))
	}
	if e.CommentID != "" {
		fields = append(fields, fmt.Sprintf("comment id \"%s\"", e.CommentID))
	}
	return strings.Join(fields, " ")
}
