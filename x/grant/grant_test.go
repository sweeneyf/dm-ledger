package grant_test

import (
	"testing"
	"time"

	"github.com/sweeneyf/dm-ledger/util"
	"github.com/sweeneyf/dm-ledger/x/grant"
	"github.com/sweeneyf/dm-ledger/x/grant/keeper"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
	"github.com/sweeneyf/dm-ledger/x/permission"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	perm "github.com/sweeneyf/dm-ledger/x/permission/keeper"
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

type grantSuite struct {
	suite.Suite
	Db         *dbm.MemDB
	Ctx        sdk.Context
	Cms        store.CommitMultiStore
	Cdc        *codec.Codec
	DS1        sdk.AccAddress
	DC1        sdk.AccAddress
	DP1        sdk.AccAddress
	Keeper     keeper.Keeper
	permKeeper perm.Keeper // used for setting up test permissions
}

// This methods is run befor all test in the suite
func (s *grantSuite) SetupSuite() {
	s.Db = dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(s.Db)
	s.Cdc = codec.New()
	s.Ctx = sdk.NewContext(s.Cms, abci.Header{}, false, log.NewNopLogger())
	s.DS1 = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.DC1 = sdk.AccAddress{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.DP1 = sdk.AccAddress{2, 3, 4, 5, 6, 7, 8, 9, 1}

	grant.RegisterCodec(s.Cdc)
	permissionKey := sdk.NewKVStoreKey("permission")
	s.permKeeper = perm.NewKeeper(s.Cdc, permissionKey)
	grantKey := sdk.NewKVStoreKey("grant")
	s.Keeper = grant.NewKeeper(s.Cdc, grantKey, s.permKeeper, nil)

	//load both stores
	s.Cms.MountStoreWithDB(permissionKey, sdk.StoreTypeIAVL, s.Db)
	s.Cms.MountStoreWithDB(grantKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
}

func (s *grantSuite) SetupTest() {
	_ = s.Cms.LoadLatestVersion()
}

func (s *grantSuite) TestHandleMsgRequestAccessSuccessful() {
	// create a test permission
	testPermission := permission.NewPermission(s.DS1, s.DC1, "data_pointer", "data_hash")
	//now add in all rights for user DP1
	testPermission.Policy.Create = append(testPermission.Policy.Create, s.DP1)
	testPermission.Policy.Read = append(testPermission.Policy.Read, s.DP1)
	testPermission.Policy.Update = append(testPermission.Policy.Update, s.DP1)
	key := testPermission.Subject.String() + testPermission.Controller.String()

	// save the permission to the permission store
	s.permKeeper.SetPermission(s.Ctx, key, &testPermission)

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
	s.Require().Equal(0, accGrant.Status)
}

func (s *grantSuite) TestHandleMsgValidateToken() {
	// create a test Grant
	token := util.GenerateUUID()
	//reqTime := ctx.BlockHeader().Time

	testGrant := types.Grant{
		Token:   token,
		Created: time.Now(),
		Status:  0,
		Expires: time.Now().Add(3600 * time.Second),
		Create:  true,
		Read:    true,
		Update:  false,
		Delete:  false,
	}

	s.Keeper.SetGrant(s.Ctx, testGrant.Token, &testGrant)

	msgValidateToken := types.NewMsgValidateToken(s.DS1, token)
	result, err := grant.HandleMsgValidateToken(s.Ctx, s.Keeper, msgValidateToken)

	// get the grant returned from result
	var accGrant grant.Grant
	err = s.Cdc.UnmarshalBinaryLengthPrefixed(result.Data, &accGrant)

	s.Require().Nil(err)
	s.Require().NotNil(accGrant)
	s.Require().Equal(true, testGrant.Created.Equal(accGrant.Created))
	s.Require().Equal(true, testGrant.Expires.Equal(accGrant.Expires))
	s.Require().Equal(testGrant.Create, accGrant.Create)
	s.Require().Equal(testGrant.Read, accGrant.Read)
	s.Require().Equal(testGrant.Update, accGrant.Update)
	s.Require().Equal(testGrant.Delete, accGrant.Delete)
}

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(grantSuite))
}
