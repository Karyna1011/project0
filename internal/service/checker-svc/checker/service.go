package checker

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/subgroup/project/internal/config"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"strings"

	"math/big"
	"time"
)

//const myABI="[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_repaidAmount\",\"type\":\"uint256\"}],\"name\":\"BorrowRepaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"DistributionRewardWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_paramKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_paramKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LiquidatorPay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidityWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expectedAllowance\",\"type\":\"uint256\"}],\"name\":\"approveToDelegateBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"}],\"name\":\"borrowFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDistributionRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"}],\"name\":\"claimPoolDistributionRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_borrowerAddr\",\"type\":\"address\"}],\"name\":\"delegateBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"}],\"name\":\"delegateRepayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"}],\"name\":\"disableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disabledCollateralAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"}],\"name\":\"enableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getAvailableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getCurrentBorrowLimitInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentBorrowLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"getLiquidiationInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"borrowAssetKeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"supplyAssetKeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowedAmount\",\"type\":\"uint256\"}],\"internalType\":\"struct IDefiCore.LiquidationInfo[]\",\"name\":\"_resultArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isAdding\",\"type\":\"bool\"}],\"name\":\"getNewBorrowLimitInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getTotalBorrowBalanceInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBorrowBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getTotalSupplyBalanceInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupplyBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserDistributionRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionRewardInUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalanceInUSD\",\"type\":\"uint256\"}],\"internalType\":\"struct IDefiCore.RewardsDistributionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_borrowAssetKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiveAssetKey\",\"type\":\"bytes32\"}],\"name\":\"getUserLiquidationInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"borrowAssetPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receiveAssetPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonusReceiveAssetPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"}],\"internalType\":\"struct IDefiCore.UserLiquidationInfo\",\"name\":\"_liquidationInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"injector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"}],\"name\":\"isCollateralAssetEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_supplyAssetKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_borrowAssetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationAmount\",\"type\":\"uint256\"}],\"name\":\"liquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contract Registry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"}],\"name\":\"updateCompoundRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
//const myABI="[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"
const myABI = "[{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"_userAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"_assetKey\",\"type\": \"bytes32\" },{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_repaidAmount\",\"type\": \"uint256\"}],\"name\": \"BorrowRepaid\",\"type\": \"event\" },{\"anonymous\": false, \"inputs\": [ {\"indexed\": false, \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" },{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_borrowedAmount\",\"type\": \"uint256\"} ], \"name\": \"Borrowed\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"_rewardAmount\", \"type\": \"uint256\" } ], \"name\": \"DistributionRewardWithdrawn\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"_paramKey\", \"type\": \"bytes32\" }, { \"indexed\": false, \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"_amount\", \"type\": \"uint256\" } ], \"name\": \"LiquidateBorrow\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"_paramKey\", \"type\": \"bytes32\" }, { \"indexed\": false, \"internalType\": \"address\", \"name\": \"_liquidatorAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"_amount\", \"type\": \"uint256\" } ], \"name\": \"LiquidatorPay\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"_liquidityAmount\", \"type\": \"uint256\" } ], \"name\": \"LiquidityAdded\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"_liquidityAmount\", \"type\": \"uint256\" } ], \"name\": \"LiquidityWithdrawn\", \"type\": \"event\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"\", \"type\": \"address\" }, {\"internalType\": \"bytes32\",\"name\": \"\",\"type\": \"bytes32\"}],\"name\": \"disabledCollateralAssets\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\",\"constant\": true},{\"inputs\": [],\"name\": \"injector\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"_injector\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_injector\", \"type\": \"address\" } ], \"name\": \"setInjector\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"contract Registry\", \"name\": \"_registry\", \"type\": \"address\" } ], \"name\": \"setDependencies\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" } ], \"name\": \"isCollateralAssetEnabled\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address[]\", \"name\": \"_accounts\", \"type\": \"address[]\" } ], \"name\": \"getLiquidiationInfo\", \"outputs\": [ { \"components\": [ { \"internalType\": \"bytes32[]\", \"name\": \"borrowAssetKeys\", \"type\": \"bytes32[]\" }, { \"internalType\": \"bytes32[]\", \"name\": \"supplyAssetKeys\", \"type\": \"bytes32[]\" }, { \"internalType\": \"uint256\", \"name\": \"totalBorrowedAmount\", \"type\": \"uint256\" } ], \"internalType\": \"struct IDefiCore.LiquidationInfo[]\", \"name\": \"_resultArr\", \"type\": \"tuple[]\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"internalType\": \"bytes32\", \"name\": \"_borrowAssetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"bytes32\", \"name\": \"_receiveAssetKey\", \"type\": \"bytes32\" } ], \"name\": \"getUserLiquidationInfo\", \"outputs\": [ { \"components\": [ { \"internalType\": \"uint256\", \"name\": \"borrowAssetPrice\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"receiveAssetPrice\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"bonusReceiveAssetPrice\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"borrowedAmount\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"supplyAmount\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"maxQuantity\", \"type\": \"uint256\" } ], \"internalType\": \"struct IDefiCore.UserLiquidationInfo\", \"name\": \"_liquidationInfo\", \"type\": \"tuple\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" } ], \"name\": \"getTotalSupplyBalanceInUSD\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"_totalSupplyBalance\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" } ], \"name\": \"getTotalBorrowBalanceInUSD\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"_totalBorrowBalance\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" } ], \"name\": \"getCurrentBorrowLimitInUSD\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"_currentBorrowLimit\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_tokensAmount\", \"type\": \"uint256\" }, { \"internalType\": \"bool\", \"name\": \"_isAdding\", \"type\": \"bool\" } ], \"name\": \"getNewBorrowLimitInUSD\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" } ], \"name\": \"getAvailableLiquidity\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" } ], \"name\": \"getUserDistributionRewards\", \"outputs\": [ { \"components\": [ { \"internalType\": \"address\", \"name\": \"assetAddr\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"distributionReward\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"distributionRewardInUSD\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"userBalance\", \"type\": \"uint256\" }, { \"internalType\": \"uint256\", \"name\": \"userBalanceInUSD\", \"type\": \"uint256\" } ], \"internalType\": \"struct IDefiCore.RewardsDistributionInfo\", \"name\": \"\", \"type\": \"tuple\" } ], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" } ], \"name\": \"updateCompoundRate\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" } ], \"name\": \"enableCollateral\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" } ], \"name\": \"disableCollateral\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_liquidityAmount\", \"type\": \"uint256\" } ], \"name\": \"addLiquidity\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_liquidityAmount\", \"type\": \"uint256\" }, { \"internalType\": \"bool\", \"name\": \"_isMaxWithdraw\", \"type\": \"bool\" } ], \"name\": \"withdrawLiquidity\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_borrowAmount\", \"type\": \"uint256\" }, { \"internalType\": \"address\", \"name\": \"_delegatee\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"_expectedAllowance\", \"type\": \"uint256\" } ], \"name\": \"approveToDelegateBorrow\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_borrowAmount\", \"type\": \"uint256\" } ], \"name\": \"borrow\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_borrowAmount\", \"type\": \"uint256\" }, { \"internalType\": \"address\", \"name\": \"_borrowerAddr\", \"type\": \"address\" } ], \"name\": \"delegateBorrow\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_borrowAmount\", \"type\": \"uint256\" }, { \"internalType\": \"address\", \"name\": \"_recipientAddr\", \"type\": \"address\" } ], \"name\": \"borrowFor\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_repayAmount\", \"type\": \"uint256\" }, { \"internalType\": \"bool\", \"name\": \"_isMaxRepay\", \"type\": \"bool\" } ], \"name\": \"repayBorrow\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"bytes32\", \"name\": \"_assetKey\", \"type\": \"bytes32\" }, { \"internalType\": \"uint256\", \"name\": \"_repayAmount\", \"type\": \"uint256\" }, { \"internalType\": \"address\", \"name\": \"_recipientAddr\", \"type\": \"address\" } ], \"name\": \"delegateRepayBorrow\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"_userAddr\", \"type\": \"address\" }, { \"internalType\": \"bytes32\", \"name\": \"_supplyAssetKey\", \"type\": \"bytes32\"},{\"internalType\": \"bytes32\",\"name\": \"_borrowAssetKey\",\"type\": \"bytes32\"},{\"internalType\": \"uint256\",\"name\": \"_liquidationAmount\",\"type\": \"uint256\"}],\"name\": \"liquidation\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"_assetKey\",\"type\": \"bytes32\"}],\"name\": \"claimPoolDistributionRewards\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"_reward\",\"type\": \"uint256\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"claimDistributionRewards\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"_totalReward\",\"type\": \"uint256\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"

func (s *Service) Run(cfg config.Config, ctx context.Context) {
	s.log.Info("Running checker service")

	contractAbi, err := abi.JSON(strings.NewReader(myABI))
	if err != nil {
		s.log.Error("failed to parse contract ABI")
		return
	}

	var Contract = bind.NewBoundContract(
		common.HexToAddress(s.contractAddress.Address),
		contractAbi,
		s.eth,
		s.eth,
		s.eth,
	)

	d := time.NewTicker(cfg.TransferConfig().Time)
	numRows := 0

	for {
		select {
		case _ = <-d.C:
			currentBlock, err := s.eth.BlockByNumber(ctx, nil)

			if err != nil {
				s.log.Error("failed to fetch current block")
				return
			}

			s.lastBlockNumber = currentBlock.Number()
			s.log.Info(s.lastBlockNumber)

			topicHashString := "e3dcbe559b610f5784b9db6ab7a250c9c7cc6aa560eae7dfc50de489602b5654"
			topicHashBytes, err := hex.DecodeString(topicHashString)
			var topicHash32 [32]byte

			for i := 0; i < 32; i++ {
				topicHash32[i] = topicHashBytes[i]
			}

			logs, err := s.eth.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: s.firstBlockNumber,
				ToBlock:   s.lastBlockNumber,
				Addresses: []common.Address{
					common.HexToAddress(s.contractAddress.Address),
				},
				Topics: [][]common.Hash{{topicHash32}},
			})

			s.firstBlockNumber = s.lastBlockNumber

			if err != nil {
				s.log.Debug("failed to set filter query for filter logs")
			}

			s.log.Info("len logs=", len(logs))

			for i := 0; i < len(logs); i++ {
				s.log.Info("in cycle")
				event := struct {
					UserAddr        common.Address `abi:"_userAddr"`
					AssetKey        [32]byte       `abi:"_assetKey"`
					LiquidityAmount *big.Int       `abi:"_liquidityAmount"`
				}{}
				//userAddress := logs[i].Address

				//s.log.Info(logs[i])
				//s.log.Info(logs[i].Address)
				//s.log.Info(logs[i].Data)
				err := contractAbi.UnpackIntoInterface(&event, "LiquidityAdded", logs[i].Data)
				if err != nil {
					s.log.WithError(err).Error("error unpacking event")
					return
				}

				s.log.Info("my event address=", event.UserAddr.String())

				person, err := s.personQ.FilterByAddress(event.UserAddr.String()).Get()
				if err != nil {
					s.log.WithError(err).Error("failed to get person")
					return
				}

				if person == nil {
					_, err = s.personQ.Insert(data.Person{
						Address: event.UserAddr.String(),
					})
					numRows++
					if err != nil {
						s.log.WithError(err).Error("failed to insert person")
						return
					}
				}

			}
			numRows = 10
			s.log.Info("rows=", numRows)

			flag := false

			for i := 1; flag == false; i++ {
				s.log.Info("in second cycle")
				person, err := s.personQ.FilterById(int64(i)).Get()
				if err != nil {
					s.log.WithError(err).Error("failed to get person")
					return

				}

				if person != nil {

					debtor, err := s.debtorQ.FilterByAddress(person.Address).Get()
					if err != nil {
						s.log.WithError(err).Error("failed to get debtor")
						return
					}

					if Checker(Contract, s.log, common.HexToAddress(person.Address)) == false {
						if debtor == nil {
							_, err = s.debtorQ.Insert(data.Debtor{
								Address: person.Address,
							})
							if err != nil {
								s.log.Error(err, "failed to insert debtor")
								return
							}
						}
					} else {
						if debtor != nil {
							err = s.debtorQ.Delete(person.Address)
							if err != nil {
								s.log.Error(err, "failed to delete debtor")
								return
							}
						}
					}
				} else {
					flag = true
				}
			}
		}
	}
}
