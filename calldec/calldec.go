package calldec

import (
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"mouse/calldec/types"
	"strings"

	"github.com/broothie/qst"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	SPAM_SIGNATURES = []string{
		"join_tg_invmru_haha",
		"CheckOutBoringSecurity",
		"niceFunctionHerePlzClick",
		"watch_tg_invmru",
		"func_2093253501",
		"many_msg_babbage",
		"sign_szabo_bytecode",
		"JunionYoutubeXD_dashhvetozhe",
		"cryethereum_",
		"check_out_my_new_website",
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
func DecodeCalldata(m *types.MouseTx, calldata []byte) (mouseTx *types.MouseTx, err error) {
	if len(calldata) < 4 {
		return nil, errors.New("invalid calldata length (expected at least 4 bytes)")
	}

	m.ByteSignature = hexutil.Encode(calldata[:4])

	// Start from 10 to skip the function selector prefixed with 0x
	c := hexutil.Encode(calldata)
	for i := 10; i < len(c); i += 64 {
		if i+64 > len(c) {
			m.Calldata = append(m.Calldata, c[i:])
		} else {
			m.Calldata = append(m.Calldata, c[i:i+64])
		}
	}

	// Convert the signature to a string (and possible signatures)
	ts, hs, err := SigToText(m.ByteSignature)
	if err != nil {
		return nil, err
	}

	// TODO: Find less hacky way to remove banned signatures
	s, h := filterSpamSignatures(ts, hs)
	for i, sig := range s {
		m.PossibleSignatures = append(m.PossibleSignatures, &types.Signature{
			TextSignature: sig,
			HexSignature:  h[i],
			Arguments: &types.Argument{
				TextArguments: []string{},
				ABIArguments:  []string{},
			},
		})
	}

	if len(m.PossibleSignatures) > 0 {
		err := ParseSignatures(m.PossibleSignatures)
		if err != nil {
			return nil, err
		}
	}

	err = filterSignatures(m)
	if err != nil {
		return nil, err
	}

	for _, sig := range m.PossibleSignatures[0].Arguments.TextArguments {
		if strings.Contains(sig, "[") || strings.Contains(sig, "]") {
			continue
		}

		if strings.Count(sig, "(") > 1 {
			continue
		}
	}

	return m, nil
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
		for _, spam := range SPAM_SIGNATURES {
			if len(sig) >= len(spam) && sig[:len(spam)] == spam {
				textSignatures = append(textSignatures[:i], textSignatures[i+1:]...)
				hexSignatures = append(hexSignatures[:i], hexSignatures[i+1:]...)
			}
		}
	}

	for _, spam := range SPAM_SIGNATURES {
		if len(textSignatures[0]) > len(spam) && textSignatures[0][:len(spam)] == spam {
			textSignatures = textSignatures[1:]
			hexSignatures = hexSignatures[1:]
		}
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
