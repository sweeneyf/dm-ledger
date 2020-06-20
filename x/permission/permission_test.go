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

func (s *permissionSuite) SetupTest() {
	_ = s.Cms.LoadLatestVersion()
}

func (s *permissionSuite) TestHandleMsgCreatePermissionSuccessful() {
	msgCreatePermission := types.NewMsgCreatePermission(s.DS1, s.DC1, "data Location", "data hash")

	res, err := permission.HandleMsgCreatePermission(s.Ctx, s.Keeper, msgCreatePermission)

	resultEvent := res.Events[len(res.Events)-1] // get the latest event
	// now check the event to verify that the permission was created successfully.
	s.Require().Nil(err)
	s.Require().Equal(4, len(resultEvent.Attributes))
	for _, attrib := range resultEvent.Attributes {
		switch string(attrib.GetKey()) {
		case sdk.AttributeKeyModule:
			s.Require().Equal(types.AttributeValueCategory, string(attrib.GetValue()))
		case sdk.AttributeKeyAction:
			s.Require().Equal(types.EventTypeCreatePermission, string(attrib.GetValue()))
		case types.AttributeController:
			s.Require().Equal(msgCreatePermission.Controller.String(), string(attrib.GetValue()))
		case types.AttributeSubject:
			s.Require().Equal(msgCreatePermission.Subject.String(), string(attrib.GetValue()))
		default:
			s.Require().Fail(fmt.Sprintf("we have a non expected attribute type '%v'", string(attrib.GetKey())))
		}
	}
}

func (s *permissionSuite) TestHandleMsgUpdatePermissionAllInsertsSuccessful() {
	// create a test permission
	testPermission := permission.NewPermission(s.DS1, s.DC1, "data_pointer", "data_hash")
	key := testPermission.Subject.String() + testPermission.Controller.String()

	// save the permission to the permission store
	s.Keeper.SetPermission(s.Ctx, key, &testPermission)

	msgUpdatePermission := types.NewMsgUpdatePermission(s.DS1, s.DC1, s.DP1, true, true, true, true)
	_, err := permission.HandleMsgUpdatePermission(s.Ctx, s.Keeper, msgUpdatePermission)

	// try and retrieve it again and assert the changes are as expected
	updatedPermission, err := s.Keeper.GetPermission(s.Ctx, key)

	s.Require().Nil(err)
	s.Require().NotNil(updatedPermission)
	s.Require().Equal(s.DS1, updatedPermission.Subject)
	s.Require().Equal(s.DC1, updatedPermission.Controller)
	s.Require().Equal("data_pointer", updatedPermission.DataPointer)
	s.Require().Equal("data_hash", updatedPermission.DataHash)
	s.Require().Equal(s.DP1, updatedPermission.Policy.Create[1]) // check to see that it was inserted
	s.Require().Equal(s.DP1, updatedPermission.Policy.Read[1])   // check to see that it was inserted
	s.Require().Equal(s.DP1, updatedPermission.Policy.Update[1]) // check to see that it was inserted
	s.Require().Equal(s.DP1, updatedPermission.Policy.Delete[1]) // check to see that it was inserted
}

func (s *permissionSuite) TestHandleMsgUpdatePermissionUpdateDeleteRightSuccessful() {
	// create a test permission
	testPermission := permission.NewPermission(s.DS1, s.DC1, "data_pointer", "data_hash")
	//now add in all rights for user DP1
	testPermission.Policy.Create = append(testPermission.Policy.Create, s.DP1)
	testPermission.Policy.Read = append(testPermission.Policy.Read, s.DP1)
	testPermission.Policy.Update = append(testPermission.Policy.Update, s.DP1)
	testPermission.Policy.Delete = append(testPermission.Policy.Delete, s.DP1)
	key := testPermission.Subject.String() + testPermission.Controller.String()

	// save the permission to the permission store
	s.Keeper.SetPermission(s.Ctx, key, &testPermission)

	// update permission removing the last delete permission for DP1
	msgUpdatePermission := types.NewMsgUpdatePermission(s.DS1, s.DC1, s.DP1, true, true, true, false)
	_, err := permission.HandleMsgUpdatePermission(s.Ctx, s.Keeper, msgUpdatePermission)

	// try and retrieve it again and assert the changes are as expected
	updatedPermission, err := s.Keeper.GetPermission(s.Ctx, key)

	s.Require().Nil(err)
	s.Require().NotNil(updatedPermission)
	s.Require().Equal(s.DS1, updatedPermission.Subject)
	s.Require().Equal(s.DC1, updatedPermission.Controller)
	s.Require().Equal("data_pointer", updatedPermission.DataPointer)
	s.Require().Equal("data_hash", updatedPermission.DataHash)
	s.Require().Equal(1, len(updatedPermission.Policy.Delete)) // check to see delete permission removed and so length 1
}

func (s *permissionSuite) TestHandleMsgUpdatePermissionNotFound() {

	// no test permission created so it should not be found to update it

	msgUpdatePermission := types.NewMsgUpdatePermission(s.DS1, s.DC1, s.DP1, true, true, true, false)
	res, err := permission.HandleMsgUpdatePermission(s.Ctx, s.Keeper, msgUpdatePermission)
	s.Require().Nil(res)
	s.Require().NotNil(err)
	s.Require().Equal("permission does not exist: Cannot locate permission ", err.Error())
}

func (s *permissionSuite) TestHandleMsgDeletePermissionSuccessful() {

	// create a test permission
	testPermission := permission.NewPermission(s.DS1, s.DC1, "data_pointer", "data_hash")
	//now add in all rights for user DP1
	testPermission.Policy.Create = append(testPermission.Policy.Create, s.DP1)
	testPermission.Policy.Read = append(testPermission.Policy.Read, s.DP1)
	testPermission.Policy.Update = append(testPermission.Policy.Update, s.DP1)
	testPermission.Policy.Delete = append(testPermission.Policy.Delete, s.DP1)
	key := testPermission.Subject.String() + testPermission.Controller.String()

	// save the permission to the permission store
	s.Keeper.SetPermission(s.Ctx, key, &testPermission)

	msgDeletePermission := types.NewMsgDeletePermission(s.DS1, s.DC1)
	res, err := permission.HandleMsgDeletePermission(s.Ctx, s.Keeper, msgDeletePermission)

	resultEvent := res.Events[len(res.Events)-1] // get the latest event
	// now check the event to verify that the permission was created successfully.
	s.Require().Nil(err)
	s.Require().NotNil(res)
	s.Require().Equal(4, len(resultEvent.Attributes))
	for _, attrib := range resultEvent.Attributes {
		switch string(attrib.GetKey()) {
		case sdk.AttributeKeyModule:
			s.Require().Equal(types.AttributeValueCategory, string(attrib.GetValue()))
		case sdk.AttributeKeyAction:
			s.Require().Equal(types.EventTypeDeletePermission, string(attrib.GetValue()))
		case types.AttributeSubject:
			s.Require().Equal(msgDeletePermission.Subject.String(), string(attrib.GetValue()))
		case types.AttributeController:
			s.Require().Equal(msgDeletePermission.Controller.String(), string(attrib.GetValue()))
		default:
			s.Require().Fail(fmt.Sprintf("we have a non expected attribute type '%v'", string(attrib.GetKey())))
		}
	}
}

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(permissionSuite))
}
