# build docker image
`docker build . -t nubila-service`

# start service
Assume $RPC_URL is the RPC of the blockchain where the contracts are deployed to
## without vesting manager contract
Before deploying vesting manager contract, the total supply is the initial supply, 1_000_000_000, and the circulating amount is 0.
`docker run -p 8080:8080 -e RPC_URL="$RPC_URL" nubila-service`

## with vesting manager contract
Once we have the vesting manager contract deployed, the service is started as follows:
`docker run -p 8080:8080 -e MANAGER_ADDRESS="$MANAGER_ADDR" -e RPC_URL="$RPC_URL" nubila-service`

# Query
## total supply
localhost:8080/total

## circulating
localhost:8080/circulating