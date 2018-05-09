package flag

import (
	"fmt"
	"strconv"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type NetworkPort struct {
	StartPort int
	EndPort   int
}

func (np *NetworkPort) UnmarshalFlag(val string) error {
	if strings.HasPrefix(val, "-") {
		return &flags.Error{
			Type:    flags.ErrExpectedArgument,
			Message: fmt.Sprintf("expected argument for flag --port, but got option %s", val),
		}
	}

	ports := strings.Split(val, "-")

	var err error
	np.StartPort, err = strconv.Atoi(ports[0])
	if err != nil {
		return &flags.Error{
			Type:    flags.ErrRequired,
			Message: `PORT must be a positive integer`,
		}
	}

	switch len(ports) {
	case 1:
		np.EndPort = np.StartPort
	case 2:
		np.EndPort, err = strconv.Atoi(ports[1])
		if err != nil {
			return &flags.Error{
				Type:    flags.ErrRequired,
				Message: `PORT must be a positive integer`,
			}
		}
	default:
		return &flags.Error{
			Type:    flags.ErrRequired,
			Message: `PORT syntax must match integer[-integer]`,
		}
	}

	return nil
}
