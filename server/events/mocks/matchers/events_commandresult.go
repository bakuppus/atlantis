package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

func AnyEventsCommandResult() events.CommandResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.CommandResult))(nil)).Elem()))
	var nullValue events.CommandResult
	return nullValue
}

func EqEventsCommandResult(value events.CommandResult) events.CommandResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.CommandResult
	return nullValue
}
