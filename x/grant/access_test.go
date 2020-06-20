package grant_test

import (
	"testing"

	"github.com/sweeneyf/dm-ledger/x/grant"
	"github.com/sweeneyf/dm-ledger/x/grant/keeper"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
	"github.com/sweeneyf/dm-ledger/x/permission"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
)

/*

   BeforeTest(suiteName, testName string) - Runs right before the test starts
   AfterTest(suiteName, testName string) - Runs right after the test finishes
   SetupSuite() - Runs before the tests in the suite
   SetupTest() - Runs before each test in the suite
   TearDownTest() - Runs after each test in the suite
   TearDownSuite() - Runs after all the tests in the suite have been run
*/

type permissionSuite struct {
	suite.Suite
	Db     *dbm.MemDB
	Ctx    sdk.Context
	Cms    store.CommitMultiStore
	Cdc    *codec.Codec
	DS1    sdk.AccAddress
	DC1    sdk.AccAddress
	DP1    sdk.AccAddress
	Keeper keeper.Keeper
}

// This methods is run befor all test in the suite
func (s *permissionSuite) SetupSuite() {
	s.Db = dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(s.Db)
	s.Cdc = codec.New()
	s.Ctx = sdk.NewContext(s.Cms, abci.Header{}, false, log.NewNopLogger())
	s.DS1 = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.DC1 = sdk.AccAddress{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.DP1 = sdk.AccAddress{2, 3, 4, 5, 6, 7, 8, 9, 1}

	permission.RegisterCodec(s.Cdc)
	permissionKey := sdk.NewKVStoreKey("permission")
	permissionKeeper := permission.NewKeeper(s.Cdc, permissionKey)
	grantKey := sdk.NewKVStoreKey("grant")
	s.Keeper = grant.NewKeeper(s.Cdc, grantKey, permissionKeeper, nil)

	s.Cms.MountStoreWithDB(permissionKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
}

func (s *permissionSuite) SetupTest() {
	_ = s.Cms.LoadLatestVersion()
}

func (s *permissionSuite) TestHandleMsgRequestAccessSuccessful() {
	// create a test permission
	testPermission := permission.NewPermission(s.DS1, s.DC1, "data_pointer", "data_hash")
	//now add in all rights for user DP1
	testPermission.Policy.Create = append(testPermission.Policy.Create, s.DP1)
	testPermission.Policy.Read = append(testPermission.Policy.Read, s.DP1)
	testPermission.Policy.Update = append(testPermission.Policy.Update, s.DP1)
	key := testPermission.Subject.String() + testPermission.Controller.String()

	// save the permission to the permission store
	s.Keeper.PermissionKeeper.SetPermission(s.Ctx, key, &testPermission)

	// update permission removing the last delete permission for DP1
	msgRequestAccess := types.NewMsgRequestAccess(s.DS1, s.DC1, s.DP1, nil)
	result, err := grant.HandleMsgRequestAccess(s.Ctx, s.Keeper, msgRequestAccess)

	// get the grant returned from result
	var accGrant grant.Grant
	err = s.Cdc.UnmarshalBinaryLengthPrefixed(result.Data, &accGrant)

	s.Require().Nil(err)
	s.Require().NotNil(accGrant)
	s.Require().Equal(true, accGrant.Create)
	s.Require().Equal(true, accGrant.Read)
	s.Require().Equal(true, accGrant.Update)
	s.Require().Equal(false, accGrant.Delete)
}

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(permissionSuite))
}
