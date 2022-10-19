package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/scroll-tech/go-ethereum/common"

	"scroll-tech/common/utils"

	db_config "scroll-tech/database"
)

// SenderConfig The config for transaction sender
type SenderConfig struct {
	// The RPC endpoint of the ethereum or scroll public node.
	Endpoint string `json:"endpoint"`
	// The time to trigger check pending txs in sender.
	CheckPendingTime uint64 `json:"check_pending_time"`
	// The number of blocks to wait to escalate increase gas price of the transaction.
	EscalateBlocks uint64 `json:"escalate_blocks"`
	// The gap number between a block be confirmed and the latest block.
	Confirmations uint64 `json:"confirmations"`
	// The numerator of gas price escalate multiple.
	EscalateMultipleNum uint64 `json:"escalate_multiple_num"`
	// The denominator of gas price escalate multiple.
	EscalateMultipleDen uint64 `json:"escalate_multiple_den"`
	// The maximum gas price can be used to send transaction.
	MaxGasPrice uint64 `json:"max_gas_price"`
	// the transaction type to use: LegacyTx, AccessListTx, DynamicFeeTx
	TxType string `json:"tx_type"`
}

// L1Config loads l1eth configuration items.
type L1Config struct {
	// Confirmations block height confirmations number.
	Confirmations uint64 `json:"confirmations"`
	// l1 chainID.
	ChainID int64 `json:"chain_id"`
	// l1 eth node url.
	Endpoint string `json:"endpoint"`
	// The start height to sync event from layer 1
	StartHeight uint64 `json:"start_height"`
	// The messenger contract address deployed on layer 1 chain.
	L1MessengerAddress common.Address `json:"l1_messenger_address,omitempty"`
	// The relayer config
	RelayerConfig *RelayerConfig `json:"relayer_config"`
}

// L2Config loads l2geth configuration items.
type L2Config struct {
	// Confirmations block height confirmations number.
	Confirmations uint64 `json:"confirmations"`
	// l2geth chainId.
	ChainID int64 `json:"chain_id"`
	// l2geth node url.
	Endpoint string `json:"endpoint"`
	// The messenger contract address deployed on layer 2 chain.
	L2MessengerAddress common.Address `json:"l2_messenger_address,omitempty"`
	// Proof generation frequency, generating proof every k blocks
	ProofGenerationFreq uint64 `json:"proof_generation_freq"`
	// Skip generating proof when that opcodes appeared
	SkippedOpcodes []string `json:"skipped_opcodes"`
	// The relayer config
	RelayerConfig *RelayerConfig `json:"relayer_config"`
}

// RelayerConfig loads relayer configuration items.
type RelayerConfig struct {
	// RollupContractAddress store the rollup contract address.
	RollupContractAddress common.Address `json:"rollup_contract_address,omitempty"`
	// MessengerContractAddress store the scroll messenger contract address.
	MessengerContractAddress common.Address `json:"messenger_contract_address"`
	// sender config
	SenderConfig *SenderConfig `json:"sender_config"`
	// The private key of the relayer
	PrivateKey string `json:"private_key"`
}

// Config load configuration items.
type Config struct {
	L1Config *L1Config           `json:"l1_config"`
	L2Config *L2Config           `json:"l2_config"`
	DBConfig *db_config.DBConfig `json:"db_config"`
}

// NewConfig returns a new instance of Config.
func NewConfig(file string) (*Config, error) {
	buf, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}

	// cover value by env fields
	cfg.DBConfig.DSN = utils.GetEnvWithDefault("DB_DSN", cfg.DBConfig.DSN)
	cfg.DBConfig.DriverName = utils.GetEnvWithDefault("DB_DRIVER", cfg.DBConfig.DriverName)

	return cfg, nil
}