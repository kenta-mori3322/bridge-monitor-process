package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	maddr "github.com/multiformats/go-multiaddr"
	"github.com/spf13/viper"

	"github.com/humansdotai/bridge-monitor-process/common"
)

type Configuration struct {
	Signer      SignerConfiguration  `json:"signer" mapstructure:"signer"`
	Humanschain ClientConfiguration  `json:"Humanschain" mapstructure:"Humanschain"`
	Metrics     MetricsConfiguration `json:"metrics" mapstructure:"metrics"`
	Chains      []ChainConfiguration `json:"chains" mapstructure:"chains"`
	TSS         TSSConfiguration     `json:"tss" mapstructure:"tss"`
	BackOff     BackOff              `json:"back_off" mapstructure:"back_off"`
}

// SignerConfiguration all the configures need by signer
type SignerConfiguration struct {
	SignerDbPath  string                    `json:"signer_db_path" mapstructure:"signer_db_path"`
	BlockScanner  BlockScannerConfiguration `json:"block_scanner" mapstructure:"block_scanner"`
	RetryInterval time.Duration             `json:"retry_interval" mapstructure:"retry_interval"`
}

// BackOff configuration
type BackOff struct {
	InitialInterval     time.Duration `json:"initial_interval" mapstructure:"initial_interval"`
	RandomizationFactor float64       `json:"randomization_factor" mapstructure:"randomization_factor"`
	Multiplier          float64       `json:"multiplier" mapstructure:"multiplier"`
	MaxInterval         time.Duration `json:"max_interval" mapstructure:"max_interval"`
	MaxElapsedTime      time.Duration `json:"max_elapsed_time" mapstructure:"max_elapsed_time"`
}

// ChainConfiguration configuration
type ChainConfiguration struct {
	ChainID             common.Chain              `json:"chain_id" mapstructure:"chain_id"`
	ChainHost           string                    `json:"chain_host" mapstructure:"chain_host"`
	ChainNetwork        string                    `json:"chain_network" mapstructure:"chain_network"`
	UserName            string                    `json:"username" mapstructure:"username"`
	Password            string                    `json:"password" mapstructure:"password"`
	RPCHost             string                    `json:"rpc_host" mapstructure:"rpc_host"`
	HTTPostMode         bool                      `json:"http_post_mode" mapstructure:"http_post_mode"` // Bitcoin core only supports HTTP POST mode
	DisableTLS          bool                      `json:"disable_tls" mapstructure:"disable_tls"`       // Bitcoin core does not provide TLS by default
	BlockScanner        BlockScannerConfiguration `json:"block_scanner" mapstructure:"block_scanner"`
	BackOff             BackOff
	OptToRetire         bool `json:"opt_to_retire" mapstructure:"opt_to_retire"` // don't emit support for this chain during keygen process
	ParallelMempoolScan int  `json:"parallel_mempool_scan" mapstructure:"parallel_mempool_scan"`
	Disabled            bool `json:"disabled" mapstructure:"disabled"`
}

// TSSConfiguration
type TSSConfiguration struct {
	BootstrapPeers []string `json:"bootstrap_peers" mapstructure:"bootstrap_peers"`
	Rendezvous     string   `json:"rendezvous" mapstructure:"rendezvous"`
	P2PPort        int      `json:"p2p_port" mapstructure:"p2p_port"`
	InfoAddress    string   `json:"info_address" mapstructure:"info_address"`
	ExternalIP     string   `json:"external_ip" mapstructure:"external_ip"`
}

// BlockScannerConfiguration settings for BlockScanner
type BlockScannerConfiguration struct {
	RPCHost                    string        `json:"rpc_host" mapstructure:"rpc_host"`
	StartBlockHeight           int64         `json:"start_block_height" mapstructure:"start_block_height"`
	BlockScanProcessors        int           `json:"block_scan_processors" mapstructure:"block_scan_processors"`
	HTTPRequestTimeout         time.Duration `json:"http_request_timeout" mapstructure:"http_request_timeout"`
	HTTPRequestReadTimeout     time.Duration `json:"http_request_read_timeout" mapstructure:"http_request_read_timeout"`
	HTTPRequestWriteTimeout    time.Duration `json:"http_request_write_timeout" mapstructure:"http_request_write_timeout"`
	MaxHTTPRequestRetry        int           `json:"max_http_request_retry" mapstructure:"max_http_request_retry"`
	BlockHeightDiscoverBackoff time.Duration `json:"block_height_discover_back_off" mapstructure:"block_height_discover_back_off"`
	BlockRetryInterval         time.Duration `json:"block_retry_interval" mapstructure:"block_retry_interval"`
	EnforceBlockHeight         bool          `json:"enforce_block_height" mapstructure:"enforce_block_height"`
	DBPath                     string        `json:"db_path" mapstructure:"db_path"`
	ChainID                    common.Chain  `json:"chain_id" mapstructure:"chain_id"`
	SuggestedFeeVersion        int           `json:"suggested_fee_version" mapstructure:"suggested_fee_version"`
}

// ClientConfiguration
type ClientConfiguration struct {
	ChainID         common.Chain `json:"chain_id" mapstructure:"chain_id" `
	ChainHost       string       `json:"chain_host" mapstructure:"chain_host"`
	ChainRPC        string       `json:"chain_rpc" mapstructure:"chain_rpc"`
	ChainHomeFolder string       `json:"chain_home_folder" mapstructure:"chain_home_folder"`
	SignerName      string       `json:"signer_name" mapstructure:"signer_name"`
	SignerPasswd    string
	BackOff         BackOff
}

type MetricsConfiguration struct {
	Enabled      bool           `json:"enabled" mapstructure:"enabled"`
	PprofEnabled bool           `json:"pprof_enabled" mapstructure:"pprof_enabled"`
	ListenPort   int            `json:"listen_port" mapstructure:"listen_port"`
	ReadTimeout  time.Duration  `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration  `json:"write_timeout" mapstructure:"write_timeout"`
	Chains       []common.Chain `json:"chains" mapstructure:"chains"`
}

type CredentialConfiguration struct {
	// Ethereum RPC Node Provider URL from Alchemy
	URL_Ethereum_RPC_Node_Provider string

	// Ethereum RPC Node Provider WSS URL from Alchemy, rinkeby
	URL_Ethereum_RPC_Node_Provider_WSS string

	// Ethereum Rinkeby USDK Contract Address
	Ethereum_USDK_Token_Address string

	// Ethereum Pool Account Address
	Ethereum_Pool_Address string

	// Ethereum Pool Account Private Key
	Ethereum_Owner_Account_Private_Key string

	// Human Pool Contract Address
	Humanchain_Pool_Address string

	// Human Pool Owner Address
	Humanchain_Pool_Owner_Address string

	// Human Pool Owner Signer Key Name
	Humanchain_Pool_Owner_Signer_KeyName string
}

// NewConfig create a new instance of configuration
func NewCredentialConfig() (*CredentialConfiguration, error) {
	return &CredentialConfiguration{
		URL_Ethereum_RPC_Node_Provider:       "",
		URL_Ethereum_RPC_Node_Provider_WSS:   "",
		Ethereum_USDK_Token_Address:          "",
		Ethereum_Pool_Address:                "",
		Ethereum_Owner_Account_Private_Key:   "",
		Humanchain_Pool_Address:              "",
		Humanchain_Pool_Owner_Address:        "",
		Humanchain_Pool_Owner_Signer_KeyName: "",
	}, nil
}

func (o *CredentialConfiguration) LoadConfig() error {
	// load .env file from given path
	// we keep it empty it will load .env from current directory

	rootPath := os.Getenv("HOME")
	err := godotenv.Load(rootPath + "/.humans/.env")

	if err != nil {
		return err
	}

	o.URL_Ethereum_RPC_Node_Provider = os.Getenv("URL_Ethereum_RPC_Node_Provider")
	o.URL_Ethereum_RPC_Node_Provider_WSS = os.Getenv("URL_Ethereum_RPC_Node_Provider_WSS")
	o.Ethereum_USDK_Token_Address = os.Getenv("Ethereum_USDK_Token_Address")
	o.Ethereum_Pool_Address = os.Getenv("Ethereum_Pool_Address")
	o.Ethereum_Owner_Account_Private_Key = os.Getenv("Ethereum_Owner_Account_Private_Key")
	o.Humanchain_Pool_Address = os.Getenv("Humanchain_Pool_Address")
	o.Humanchain_Pool_Owner_Address = os.Getenv("Humanchain_Pool_Owner_Address")
	o.Humanchain_Pool_Owner_Signer_KeyName = os.Getenv("Humanchain_Pool_Owner_Signer_KeyName")

	return nil
}

// LoadProcessorConfig read the processor configuration from the given file
func LoadProcessorConfig(file string) (*Configuration, error) {
	applyDefaultConfig()
	var cfg Configuration
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(file))
	viper.SetConfigName(strings.TrimRight(path.Base(file), ".json"))
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fail to read from config file: %w", err)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fail to unmarshal: %w", err)
	}

	for i, chain := range cfg.Chains {
		if err := chain.ChainID.Validate(); err != nil {
			return nil, err
		}
		if err := chain.BlockScanner.ChainID.Validate(); err != nil {
			return nil, err
		}
		cfg.Chains[i].BackOff = cfg.BackOff
	}

	return &cfg, nil
}

// GetBootstrapPeers return the internal bootstrap peers in a slice of maddr.Multiaddr
func (c TSSConfiguration) GetBootstrapPeers() ([]maddr.Multiaddr, error) {
	var addrs []maddr.Multiaddr
	for _, item := range c.BootstrapPeers {
		if len(item) > 0 {
			addr, err := maddr.NewMultiaddr(item)
			if err != nil {
				return nil, fmt.Errorf("fail to parse multi addr(%s): %w", item, err)
			}
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}

func applyDefaultConfig() {
	viper.SetDefault("metrics.listen_port", "9000")
	viper.SetDefault("metrics.read_timeout", "30s")
	viper.SetDefault("metrics.write_timeout", "30s")
	viper.SetDefault("metrics.chains", common.Chains{common.ETHChain})
	viper.SetDefault("humanchain.chain_id", "humanchain")
	viper.SetDefault("humanchain.chain_host", "localhost:1317")
	viper.SetDefault("back_off.initial_interval", 500*time.Millisecond)
	viper.SetDefault("back_off.randomization_factor", 0.5)
	viper.SetDefault("back_off.multiplier", 1.5)
	viper.SetDefault("back_off.max_interval", 3*time.Minute)
	viper.SetDefault("back_off.max_elapsed_time", 168*time.Hour) // 7 days. Due to node sync time's being so random
	applyDefaultSignerConfig()
}

func applyBlockScannerDefault(path string) {
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.start_block_height", path), "0")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.block_scan_processors", path), "2")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.http_request_timeout", path), "30s")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.http_request_read_timeout", path), "30s")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.http_request_write_timeout", path), "30s")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.max_http_request_retry", path), "10")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.block_height_discover_back_off", path), "1s")
	viper.SetDefault(fmt.Sprintf("%s.block_scanner.block_retry_interval", path), "1s")
}

func applyDefaultSignerConfig() {
	viper.SetDefault("signer.signer_db_path", "signer_db")
	applyBlockScannerDefault("signer")
	viper.SetDefault("signer.retry_interval", "2s")
	viper.SetDefault("signer.block_scanner.chain_id", "ThorChain")
}
