#!/bin/bash
rm -r ~/.dalcli
rm -r ~/.dald

# Initialize configuration files and genesis file
# moniker is the name of your node
dald init dl-node1 --chain-id permission-ledger

# Configure your CLI to eliminate need to declare them as flags
# first line tells cli to interact with the correct node access-ledger
dalcli config chain-id permission-ledger
dalcli config output json
dalcli config indent true
dalcli config trust-node true


# We'll use the "test" keyring backend which save keys unencrypted in the configuration directory of your project (defaults to ~/.dald). You should **never** use the "test" keyring backend in production. 
# For more information about other options for keyring-backend take a look at https://docs.cosmos.network/master/interfaces/keyring.html
dalcli config keyring-backend test 

dalcli keys add ds1  # data subject 1
dalcli keys add dc1 # data controller 1
dalcli keys add dp1 # data processor 1
dalcli keys add ds2  # data subject 2
dalcli keys add dc2 # data controller 2
dalcli keys add dp2 # data processor 2
dalcli keys add sp1 # service provider 1
dalcli keys add sp2 # service provider 2


# Add both accounts, with coins to the genesis file
dald add-genesis-account $(dalcli keys show ds1 -a) 1000xal
dald add-genesis-account $(dalcli keys show dc1 -a) 1000xal,100000000stake
dald add-genesis-account $(dalcli keys show dp1 -a) 1000xal
dald add-genesis-account $(dalcli keys show ds2 -a) 1000xal
dald add-genesis-account $(dalcli keys show dc2 -a) 1000xal,100000000stake
dald add-genesis-account $(dalcli keys show dp2 -a) 1000xal
dald add-genesis-account $(dalcli keys show sp1 -a) 1000xal
dald add-genesis-account $(dalcli keys show sp2 -a) 1000xal


# let the application know that DC1 willbethe only validator
dald gentx --name dc1 --keyring-backend test

# let the application know we are done configuring it
dald collect-gentxs

# validate te genesis file
dald validate-genesis



# ***** sopme basic tests to see if it running correctly
dalcli query account $(dalcli keys show dc1 -a)

#Create a permission
 dalcli tx permission create $(dalcli keys show ds1 -a) $(dalcli keys show dc1 -a) data-location data-hash  --from $(dalcli keys show ds1 -a)

#Delete a permission
dalcli tx permission delete $(dalcli keys show ds1 -a) $(dalcli keys show dc1 -a) --from ds1

#List all permissions
dalcli q permission list

#List specific permissions
dalcli q permission detail "cosmos1uu6lc2370ztj4d29lw6wcrw20093kuvr8gtyd5cosmos1fk0nlu3ysspx46cy8s98l0e3w9tfssnrmf2ntfcosmos1wukwvah8d9cfkvygyk032rasanfh4xdwhf2wp4"


#Create a grant request
 dalcli tx grant request $(dalcli keys show ds1 -a) $(dalcli keys show dc1 -a)  $(dalcli keys show dp1 -a) 12xal  --from $(dalcli keys show dp1 -a)

#List all grants
dalcli q grant list

#query that grant
dalcli q grant detail e0aa9b27290f4c3f97667a27f80a7404

#validate that grant
dalcli tx grant validate $(dalcli keys show sp1 -a) e0aa9b27290f4c3f97667a27f80a7404 --from $(dalcli keys show sp1 -a)



