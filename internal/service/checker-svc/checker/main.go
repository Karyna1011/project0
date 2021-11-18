package checker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	bscClient "github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"math/big"

	"gitlab.com/tokend/subgroup/project/internal/config"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"gitlab.com/tokend/subgroup/project/internal/data/postgres"
)

type Service struct {
	bscClient        bscClient.Client
	personQ          data.PersonQ
	debtorQ          data.DebtorQ
	contractAddress  config.ContractConfig
	firstBlockNumber *big.Int
	lastBlockNumber  *big.Int
	log              *logan.Entry
	db               *pgdb.DB
	eth              *ethclient.Client
}

func New(cfg config.Config) *Service {
	log := cfg.Log().WithField("main_service", "checker")

	return &Service{
		personQ:          postgres.NewPersonQ(cfg.DB()),
		debtorQ:          postgres.NewDebtorQ(cfg.DB()),
		contractAddress:  cfg.ContractConfig(),
		firstBlockNumber: cfg.TransferConfig().Block,
		log:              log,
		db:               cfg.DB(),
		eth:              cfg.EthClient(),
	}
}
