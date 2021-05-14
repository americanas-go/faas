package cloudevents

import "strings"

type HandlerWrapperOptions struct {
	IDsToDiscard []string
}

func DefaultHandlerWrapperOptions() (*HandlerWrapperOptions, error) {
	o := new(HandlerWrapperOptions)

	rawIDs := HandleDiscardEventsIDValue()
	ids := strings.Split(rawIDs, ",")
	if nil != ids && len(ids) >= 1 && ids[0] == "" {
		return o, nil
	}

	o.IDsToDiscard = ids

	return o, nil
}
