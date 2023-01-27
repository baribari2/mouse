package calldec

import (
	"encoding/json"
	"errors"
	"log"
	"math/big"

	"github.com/baribari2/mouse/common/types"
	"github.com/broothie/qst"
	"github.com/ethereum/go-ethereum/common"
)

var (
	SPAM_SIGNATURES = map[string]bool{
		"join_tg_invmru_haha":          true,
		"CheckOutBoringSecurity":       true,
		"niceFunctionHerePlzClick":     true,
		"watch_tg_invmru":              true,
		"func_2093253501":              true,
		"many_msg_babbage":             true,
		"sign_szabo_bytecode":          true,
		"JunionYoutubeXD_dashhvetozhe": true,
		"cryethereum_":                 true,
		"check_out_my_new_website":     true,
		"please_fix_collisions":        true,
	}
)

func NewMouseTx(hash common.Hash, cost *big.Int, gasLimit uint64, gasPrice *big.Int, nonce uint64, to *common.Address, value *big.Int, txType uint8) *types.MouseTx {
	return &types.MouseTx{
		PossibleSignatures: []*types.Signature{},
		Hash:               hash,
		Cost:               cost,
		GasLimit:           gasLimit,
		GasPrice:           gasPrice,
		Nonce:              nonce,
		To:                 to,
		Value:              value,
		Type:               txType,
	}
}

// Convert ethereum calldata to a local data structure for easier handling
func DecodeCalldata(m *types.MouseTx) (err error) {
	if len(m.RawCalldata) < 7 {
		return errors.New("invalid calldata length (expected at least 8 characters)")
	}

	if m.RawCalldata[:2] == "0x" {
		m.RawCalldata = m.RawCalldata[2:]
	}

	// Convert the signature to a string (and possible signatures)
	ts, hs, err := SigToText(m.RawCalldata[:8])
	if err != nil {
		return err
	}

	s, _ := filterSpamSignatures(ts, hs)
	for _, sig := range s {
		m.PossibleSignatures = append(m.PossibleSignatures, &types.Signature{
			TextSignature: sig,
			Arguments: &types.Argument{
				TextArguments: []string{},
				ABIArguments:  []string{},
			},
		})
	}

	if len(m.PossibleSignatures) > 0 {
		err := ParseSignatures(m.PossibleSignatures)
		if err != nil {
			return err
		}
	}

	err = filterSignatures(m)
	if err != nil {
		return err
	}

	// for _, sig := range m.PossibleSignatures[0].Arguments.TextArguments {
	// 	if strings.Contains(sig, "[") || strings.Contains(sig, "]") {
	// 		continue
	// 	}

	// 	if strings.Count(sig, "(") > 1 {
	// 		continue
	// 	}
	// }

	return nil
}

// Uses 4byte directory to convert a function signature (prefixed with 0x) to a human readable text signature
func SigToText(sig string) (textSignatures []string, hexSignatures []string, err error) {
	// A funciton signature is 4 bytes long, so it should be at least 8 characters long
	if len(sig) < 8 {
		return nil, nil, errors.New("invalid sig length")
	}

	// Query the 4byte directory for the signature
	res, err := qst.Get(
		"https://www.4byte.directory/api/v1/signatures",
		qst.QueryValue("hex_signature", sig),
		qst.Header("Content-Type", "application/json"),
	)

	if err != nil {
		return nil, nil, err
	}

	if res.ContentLength == 0 {
		return nil, nil, errors.New("no content")
	}

	var rs struct {
		Results []struct {
			Id            int    `json:"id"`
			TextSignature string `json:"text_signature"`
			HexSignature  string `json:"hex_signature"`
			ByteSignature string `json:"byte_signature"`
		} `json:"results"`
	}

	err = json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return nil, nil, err
	}

	if len(rs.Results) == 0 {
		return nil, nil, errors.New("no results")
	}

	for _, r := range rs.Results {
		textSignatures = append(textSignatures, r.TextSignature)
		hexSignatures = append(hexSignatures, r.HexSignature)
	}

	return
}

// Uses samczsun's endpoint to convert sig to name
func SigToText2(sig string) (signature string, err error) {
	if len(sig) < 8 {
		return "", errors.New("invalid sig length")
	}

	var rs struct {
		Ok     bool `json:"ok"`
		Result struct {
			Event    struct{} `json:"event"`
			Function struct {
				Sig []struct {
					Name     string `json:"name"`
					Filtered bool   `json:"filtered"`
				}
			} `json:"function"`
		} `json:"result"`
	}

	rss, _ := json.Marshal(rs)
	var am interface{}
	json.Unmarshal(rss, &am)

	cm := am.(map[string]interface{})

	log.Printf("Sig: %+v", cm)

	res, err := qst.Get(
		"https://sig.eth.samczsun.com/api/v1/signatures",
		qst.QueryValue("function", sig),
		qst.Header("Content-Type", "application/json"),
	)

	if err != nil {
		return "", err
	}

	err = json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return "", err
	}

	if !rs.Ok {
		return "", errors.New("not ok")
	}

	return "", nil
}

func filterSpamSignatures(textSignatures, hexSignatures []string) (filteredText, filteredHex []string) {
	for i, sig := range textSignatures {

		if _, exists := SPAM_SIGNATURES[sig]; exists {
			textSignatures = append(textSignatures[:i], textSignatures[i+1:]...)
			hexSignatures = append(hexSignatures[:i], hexSignatures[i+1:]...)
		}
	}

	if _, exists := SPAM_SIGNATURES[textSignatures[0]]; exists {
		textSignatures = textSignatures[1:]
		hexSignatures = hexSignatures[1:]
	}

	for i, sig := range textSignatures {
		filteredText = append(filteredText, sig)
		filteredHex = append(filteredHex, hexSignatures[i])
	}

	return filteredText, filteredHex
}

func filterSignatures(tx *types.MouseTx) error {
	for _, sig := range tx.PossibleSignatures {
		if len(sig.Arguments.TextArguments) > 0 {
			if len(tx.Calldata) < len(sig.Arguments.TextArguments) {
				// Remove the first signature and replace it with the next one
				log.Printf("\x1b[33m%s\x1b[0m%v", "Removing incorrect signature: ", sig.TextSignature)
				if len(tx.PossibleSignatures) > 1 {
					tx.PossibleSignatures = tx.PossibleSignatures[1:]
				}
			}
		}
	}

	return nil
}
