package permission_test

import (
	"fmt"
	"testing"

	"github.com/sweeneyf/dm-ledger/x/permission/types"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweeneyf/dm-ledger/x/permission"
	"github.com/sweeneyf/dm-ledger/x/permission/keeper"
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
	s.Keeper = permission.NewKeeper(s.Cdc, permissionKey)
	s.Cms.MountStoreWithDB(permissionKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
}

func (s *permissionSuite) TestHandleMsgCreatepermissionSuccessful() {
	coins, _ := sdk.ParseCoins("10")
	msgCreatepermission := types.NewMsgCreatepermission(s.DS1, s.DC1, s.DP1, "Read", "Location", coins)

	res, err := permission.HandleMsgCreatepermission(s.Ctx, s.Keeper, msgCreatepermission)

	resultEvent := res.Events[0]
	// now check the event to verify that the permission was created successfully.
	s.Require().Nil(err)
	s.Require().Equal(8, len(resultEvent.Attributes))
	for _, attrib := range resultEvent.Attributes {
		switch string(attrib.GetKey()) {
		case sdk.AttributeKeyModule:
			s.Require().Equal(types.AttributeValueCategory, string(attrib.GetValue()))
		case sdk.AttributeKeyAction:
			s.Require().Equal(types.EventTypeCreatepermission, string(attrib.GetValue()))
		case sdk.AttributeKeySender:
			s.Require().Equal(msgCreatepermission.Subject.String(), string(attrib.GetValue()))
		case types.AttributeController:
			s.Require().Equal(msgCreatepermission.Controller.String(), string(attrib.GetValue()))
		case types.AttributeSubject:
			s.Require().Equal(msgCreatepermission.Subject.String(), string(attrib.GetValue()))
		case types.AttributeDataPointer:
			s.Require().Equal(msgCreatepermission.Location, string(attrib.GetValue()))
		case types.AttributeAccessType:
			s.Require().Equal(msgCreatepermission.AccessType, string(attrib.GetValue()))
		case types.AttributeReward:
			s.Require().Equal(msgCreatepermission.Reward.String(), string(attrib.GetValue()))
		default:
			s.Require().Fail(fmt.Sprintf("we have a non expected attribute type '%v'", string(attrib.GetKey())))
		}
	}
}

/*
func (s *permissionSuite) TestHandleMsgDeletepermissionSuccessful() {
	// create a test permission
	key := msgCreatepermission.Subject.String() + msgCreatepermission.Controller.String() + msgCreatepermission.Processor.String()
	testpermissionToInsert := &permission.AccessPermission{
		Subject:    s.DS1,
		Controller: s.DC1,
		Processor:  s.DP1,
		GDPRData: permission.GDPRData{
			Location: "Location",
			EncrKey:  "TODO Change this to encrKey",
			Policy:   permission.Policy{AccessType: "Read"},
		},
	}
	// save the permission to the permission store
	s.Keeper.SetPermission(s.Ctx, key, testpermissionToInsert)

	msgDeletepermission := types.NewMsgDeletepermission(s.DS1, s.DC1, s.DP1, "Location")
	_, err := permission.HandleMsgDeletpermission(s.Ctx, s.Keeper, msgDeletepermission)

	// try and retrieve it again to make sure it is deleted
	deletedpermission, _ = s.Keeper.GetPermission(s.Ctx, key)

	s.Require().Nil(err)
	s.Require().Nil(deletedpermission)
}
*/
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(permissionSuite))
}
