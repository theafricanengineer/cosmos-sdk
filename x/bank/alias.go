// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/bank/types
package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	DefaultCodespace         = types.DefaultCodespace
	CodeSendDisabled         = types.CodeSendDisabled
	CodeInvalidInputsOutputs = types.CodeInvalidInputsOutputs
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	DefaultParamspace        = types.DefaultParamspace
	DefaultSendEnabled       = types.DefaultSendEnabled
)

var (
	// functions aliases
	RegisterCodec          = types.RegisterCodec
	ErrNoInputs            = types.ErrNoInputs
	ErrNoOutputs           = types.ErrNoOutputs
	ErrInputOutputMismatch = types.ErrInputOutputMismatch
	ErrSendDisabled        = types.ErrSendDisabled
	NewMsgSend             = types.NewMsgSend
	NewMsgMultiSend        = types.NewMsgMultiSend
	NewInput               = types.NewInput
	NewOutput              = types.NewOutput
	ValidateInputsOutputs  = types.ValidateInputsOutputs
	ParamKeyTable          = types.ParamKeyTable

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	ParamStoreKeySendEnabled = types.ParamStoreKeySendEnabled
)

type (
	MsgSend      = types.MsgSend
	MsgMultiSend = types.MsgMultiSend
	Input        = types.Input
	Output       = types.Output
)