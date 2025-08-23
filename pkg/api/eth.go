package api

import (
	"github.com/pkg/errors"

	"github.com/anonpragmatic/gowaves/pkg/proto"
	"github.com/anonpragmatic/gowaves/pkg/proto/ethabi"
	"github.com/anonpragmatic/gowaves/pkg/ride/serialization"
	"github.com/anonpragmatic/gowaves/pkg/state/stateerr"
)

func (a *App) EthereumDAppMethods(addr proto.WavesAddress) (ethabi.MethodsMap, error) {
	scriptInfo, err := a.state.ScriptInfoByAccount(proto.NewRecipientFromAddress(addr))
	if err != nil {
		if stateerr.IsNotFound(err) {
			return ethabi.MethodsMap{}, errors.Wrap(notFound, "script is not found")
		}
		return ethabi.MethodsMap{}, err
	}
	if len(scriptInfo.Bytes) == 0 {
		return ethabi.MethodsMap{}, errors.Wrap(notFound, "script is empty")
	}
	tree, err := serialization.Parse(scriptInfo.Bytes)
	if err != nil {
		return ethabi.MethodsMap{}, err
	}
	return ethabi.NewMethodsMapFromRideDAppMeta(tree.Meta)
}
