package main

import (
	"math/big"
	"os"
	"time"

	"github.com/CoderZhi/nubila-token/backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	rpc := os.Getenv("RPC_URL")
	addr := os.Getenv("MANAGER_ADDRESS")
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}
	var token *contract.IERC20
	if addr != "" {
		manager, err := contract.NewVestingManager(common.HexToAddress(addr), client)
		if err != nil {
			panic(err)
		}
		tokenAddr, err := manager.Token(nil)
		if err != nil {
			panic(err)
		}
		token, err = contract.NewIERC20(tokenAddr, client)
		if err != nil {
			panic(err)
		}
	}
	var (
		initSupply    = big.NewFloat(1_000_000_000)
		decimal       = big.NewFloat(1e18)
		totalSupply   *big.Float
		withheld      *big.Float
		totalSupplyTs time.Time
		withheldTs    time.Time
	)
	getTotalSupply := func() (*big.Float, error) {
		if token == nil {
			return initSupply, nil
		}
		if time.Since(totalSupplyTs) < time.Minute && totalSupply != nil {
			return totalSupply, nil
		}
		supply, err := token.TotalSupply(nil)
		if err != nil {
			return nil, err
		}
		totalSupply = new(big.Float).Quo(new(big.Float).SetInt(supply), decimal)
		totalSupplyTs = time.Now()
		return totalSupply, nil
	}
	getWithheld := func() (*big.Float, error) {
		if token == nil {
			return initSupply, nil
		}
		if time.Since(withheldTs) < time.Minute && withheld != nil {
			return withheld, nil
		}
		balance, err := token.BalanceOf(nil, common.HexToAddress(addr))
		if err != nil {
			return nil, err
		}
		withheld = new(big.Float).Quo(new(big.Float).SetInt(balance), decimal)
		withheldTs = time.Now()
		return withheld, nil
	}

	r := gin.Default()

	r.GET("/total", func(c *gin.Context) {
		total, err := getTotalSupply()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.String(200, total.Text('f', 18))
	})
	r.GET("/circulating", func(c *gin.Context) {
		total, err := getTotalSupply()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		withheld, err := getWithheld()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.String(200, new(big.Float).Sub(total, withheld).Text('f', 18))
	})
	r.Run(":8080")
}
