package grant_test

import (
	"testing"

	"github.com/sweeneyf/dm-ledger/x/grant/types"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweeneyf/dm-ledger/x/grant"
	"github.com/sweeneyf/dm-ledger/x/grant/keeper"
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

type GrantSuite struct {
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
func (s *GrantSuite) SetupSuite() {
	s.Db = dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(s.Db)
	s.Cdc = codec.New()
	s.Ctx = sdk.NewContext(s.Cms, abci.Header{}, false, log.NewNopLogger())
	s.DS1 = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.DC1 = sdk.AccAddress{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.DP1 = sdk.AccAddress{2, 3, 4, 5, 6, 7, 8, 9, 1}

	grant.RegisterCodec(s.Cdc)
	grantKey := sdk.NewKVStoreKey("grant")
	s.Keeper = grant.NewKeeper(s.Cdc, grantKey)
	s.Cms.MountStoreWithDB(grantKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
}

func (s *GrantSuite) TestHandleMsgCreateGrantSuccessful() {
	coins, _ := sdk.ParseCoins("10")
	msgCreateGrant := types.NewMsgCreateGrant(s.DS1, s.DC1, s.DP1, "Read", "Location", coins)
	res, err := grant.HandleMsgCreateGrant(s.Ctx, s.Keeper, msgCreateGrant)

	resultEvent := res.Events[0]
	// now check the event to verify that the Grant was created successfully.
	s.Require().Equal(err, nil)
	s.Require().Equal(8, len(resultEvent.Attributes))
	for _, attrib := range resultEvent.Attributes {
		switch string(attrib.GetKey()) {
		case sdk.AttributeKeyModule:
			s.Require().Equal(types.AttributeValueCategory, string(attrib.GetValue()))
		case sdk.AttributeKeyAction:
			s.Require().Equal(types.EventTypeCreateGrant, string(attrib.GetValue()))
		case sdk.AttributeKeySender:
			s.Require().Equal(msgCreateGrant.Subject.String(), string(attrib.GetValue()))
		case types.AttributeController:
			s.Require().Equal(msgCreateGrant.Controller.String(), string(attrib.GetValue()))
		case types.AttributeSubject:
			s.Require().Equal(msgCreateGrant.Subject.String(), string(attrib.GetValue()))
		case types.AttributeLocation:
			s.Require().Equal(msgCreateGrant.Location, string(attrib.GetValue()))
		case types.AttributeAccessType:
			s.Require().Equal(msgCreateGrant.AccessType, string(attrib.GetValue()))
		case types.AttributeReward:
			s.Require().Equal(msgCreateGrant.Reward.String(), string(attrib.GetValue()))
		default:
			s.Require().Equal("This means we have a non expected attribute type", string(attrib.GetKey()))
		}
	}
}

/*
sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAccessRequest),
sdk.NewAttribute(sdk.AttributeKeySender, msg.Processor.String()),
sdk.NewAttribute(types.AttributeController, msg.Controller.String()),
sdk.NewAttribute(types.AttributeSubject, msg.Subject.String()),
sdk.NewAttribute(types.AttributeLocation, msg.Location),
sdk.NewAttribute(types.AttributeAccessType, msg.AccessType),
sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
*/

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(GrantSuite))
}
