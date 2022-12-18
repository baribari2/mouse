package calldec

import (
	"errors"
	"fmt"
	"mouse/calldec/types"
	"strings"
)

// Parses signatures string for name and arguments into a local data structure for easier handling
func ParseSignatures(signatures []*types.Signature) (err error) {
	for _, signature := range signatures {
		if len(signature.HexSignature) != 10 {
			return errors.New("invalid signature length (expected at least 5 bytes)")
		}

		// Find first parenthesis to aid in getting both name and start of arguments
		in := strings.Index(signature.TextSignature, "(")

		if in == -1 {
			return errors.New(fmt.Sprintf("invalid sig: %v", signature.TextSignature))
		}

		signature.Name = signature.TextSignature[:in]

		// Temp arr containing arg string
		args := signature.TextSignature[in+1 : len(signature.TextSignature)-1]

		signature.Arguments.TextArguments = strings.Split(args, ",")

		// If the arguments contain more than one set of parenthesis it's likely a tuple
		tj := strings.Join(signature.Arguments.TextArguments, ",")
		if strings.Contains(tj, "(") || strings.Contains(tj, ")") {
			// log.Printf("complex sig: %v", signature.TextSignature)
			continue
		}

		// Set arguments for ABI string
		for i, arg := range signature.Arguments.TextArguments {
			if len(arg) > 0 {
				if i == len(signature.Arguments.TextArguments)-1 {
					signature.Arguments.ABIArguments = append(signature.Arguments.ABIArguments, fmt.Sprintf(`{"name":"","type":"%s"}`, arg))
				} else {
					signature.Arguments.ABIArguments = append(signature.Arguments.ABIArguments, fmt.Sprintf(`{"name":"","type":"%s"},`, arg))
				}
			}
		}
	}

	return nil
}
