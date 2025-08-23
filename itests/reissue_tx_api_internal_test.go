//go:build !smoke

package itests

import (
	"fmt"
	"maps"
	"net/http"
	"testing"

	f "github.com/anonpragmatic/gowaves/itests/fixtures"
	"github.com/anonpragmatic/gowaves/itests/testdata"
	utl "github.com/anonpragmatic/gowaves/itests/utilities"
	"github.com/anonpragmatic/gowaves/itests/utilities/issue"
	"github.com/anonpragmatic/gowaves/itests/utilities/reissue"
	"github.com/anonpragmatic/gowaves/pkg/crypto"
	"github.com/stretchr/testify/suite"
)

type ReissueTxAPIPositiveSuite struct {
	f.BaseSuite
}

func (suite *ReissueTxAPIPositiveSuite) Test_ReissueTxAPIPositive() {
	versions := reissue.GetVersions(&suite.BaseSuite)
	for _, v := range versions {
		reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
		itx := issue.BroadcastWithTestData(&suite.BaseSuite, reissuable, v, true)
		tdmatrix := testdata.GetReissuePositiveDataMatrix(&suite.BaseSuite, itx.TxID)
		for name, td := range tdmatrix {
			caseName := utl.GetTestcaseNameWithVersion(name, v)
			suite.Run(caseName, func() {
				tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset :=
					reissue.BroadcastReissueTxAndGetBalances(&suite.BaseSuite, td, v, true)
				errMsg := fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx.TxID.String())
				reissue.PositiveAPIChecks(suite.T(), tx, td, actualDiffBalanceInWaves, actualDiffBalanceInAsset, errMsg)
			})
		}
	}
}

func (suite *ReissueTxAPIPositiveSuite) Test_ReissueTxAPIMaxQuantityPositive() {
	versions := reissue.GetVersions(&suite.BaseSuite)
	for _, v := range versions {
		reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
		itx := issue.BroadcastWithTestData(&suite.BaseSuite, reissuable, v, true)
		tdmatrix := testdata.GetReissueMaxQuantityValue(&suite.BaseSuite, itx.TxID)
		for name, td := range tdmatrix {
			caseName := utl.GetTestcaseNameWithVersion(name, v)
			suite.Run(caseName, func() {
				tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset :=
					reissue.BroadcastReissueTxAndGetBalances(&suite.BaseSuite, td, v, true)
				errMsg := fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx.TxID.String())
				reissue.PositiveAPIChecks(suite.T(), tx, td, actualDiffBalanceInWaves, actualDiffBalanceInAsset, errMsg)
			})
		}
	}
}

func TestReissueTxAPIPositiveSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ReissueTxAPIPositiveSuite))
}

type ReissueTxAPINegativeSuite struct {
	f.BaseNegativeSuite
}

func (suite *ReissueTxAPINegativeSuite) Test_ReissueNotReissuableAPINegative() {
	versions := reissue.GetVersions(&suite.BaseSuite)
	txIds := make(map[string]*crypto.Digest)
	for _, v := range versions {
		reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
		itx := issue.BroadcastWithTestData(&suite.BaseSuite, reissuable, v, true)
		tdmatrix := testdata.GetNotReissuableTestData(&suite.BaseSuite, itx.TxID)
		for name, td := range tdmatrix {
			caseName := utl.GetTestcaseNameWithVersion(name, v)
			suite.Run(caseName, func() {
				//first tx should be successful
				tx1, _, _ := reissue.BroadcastReissueTxAndGetBalances(&suite.BaseSuite, td, v, true)
				errMsg := fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx1.TxID.String())
				utl.StatusCodesCheck(suite.T(), http.StatusOK, http.StatusOK, tx1, errMsg)
				utl.TxInfoCheck(suite.T(), tx1.WtErr.ErrWtGo, tx1.WtErr.ErrWtScala, errMsg)

				//second reissue tx should be failed because of reissuable=false
				tx2, actualDiffBalanceInWaves, actualDiffBalanceInAsset :=
					reissue.BroadcastReissueTxAndGetBalances(
						&suite.BaseSuite, testdata.ReissueDataChangedTimestamp(&td), v, false)
				errMsg = fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx2.TxID.String())
				txIds[name] = &tx2.TxID
				reissue.NegativeAPIChecks(suite.T(), tx2, td, actualDiffBalanceInWaves, actualDiffBalanceInAsset, errMsg)
			})
		}
	}
	actualTxIds := utl.GetTxIdsInBlockchain(&suite.BaseSuite, txIds)
	suite.Lenf(actualTxIds, 0, "IDs: %#v", actualTxIds)
}

func (suite *ReissueTxAPINegativeSuite) Test_ReissueTxAPINFTNegative() {
	versions := reissue.GetVersions(&suite.BaseSuite)
	txIds := make(map[string]*crypto.Digest)
	for _, v := range versions {
		nft := testdata.GetCommonIssueData(&suite.BaseSuite).NFT
		itx := issue.BroadcastWithTestData(&suite.BaseSuite, nft, v, true)
		tdmatrix := testdata.GetReissueNFTData(&suite.BaseSuite, itx.TxID)
		for name, td := range tdmatrix {
			caseName := utl.GetTestcaseNameWithVersion(name, v)
			suite.Run(caseName, func() {
				tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset :=
					reissue.BroadcastReissueTxAndGetBalances(&suite.BaseSuite, td, v, false)
				txIds[name] = &tx.TxID
				errMsg := fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx.TxID.String())
				reissue.NegativeAPIChecks(suite.T(), tx, td, actualDiffBalanceInWaves, actualDiffBalanceInAsset, errMsg)
			})
		}
	}
	actualTxIds := utl.GetTxIdsInBlockchain(&suite.BaseSuite, txIds)
	suite.Lenf(actualTxIds, 0, "IDs: %#v", actualTxIds)
}

func (suite *ReissueTxAPINegativeSuite) Test_ReissueTxAPINegative() {
	versions := reissue.GetVersions(&suite.BaseSuite)
	txIds := make(map[string]*crypto.Digest)
	for _, v := range versions {
		reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
		itx := issue.BroadcastWithTestData(&suite.BaseSuite, reissuable, v, true)
		tdmatrix := testdata.GetReissueNegativeDataMatrix(&suite.BaseSuite, itx.TxID)
		//TODO (ipereiaslavskaia) For v1 of reissue tx negative cases for chainID will be ignored
		if v >= 2 {
			maps.Copy(tdmatrix, testdata.GetReissueChainIDNegativeDataMatrix(&suite.BaseSuite, itx.TxID))
		}
		for name, td := range tdmatrix {
			caseName := utl.GetTestcaseNameWithVersion(name, v)
			suite.Run(caseName, func() {
				tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset :=
					reissue.BroadcastReissueTxAndGetBalances(&suite.BaseSuite, td, v, false)
				txIds[name] = &tx.TxID
				errMsg := fmt.Sprintf("Case: %s; Broadcast Reissue tx: %s", caseName, tx.TxID.String())
				reissue.NegativeAPIChecks(suite.T(), tx, td, actualDiffBalanceInWaves, actualDiffBalanceInAsset, errMsg)
			})
		}
	}
	actualTxIds := utl.GetTxIdsInBlockchain(&suite.BaseSuite, txIds)
	suite.Lenf(actualTxIds, 0, "IDs: %#v", actualTxIds)
}

func TestReissueTxAPINegativeSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ReissueTxAPINegativeSuite))
}
