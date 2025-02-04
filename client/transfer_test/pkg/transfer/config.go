package transfer

import "time"

type duration time.Duration

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	*(*time.Duration)(d), err = time.ParseDuration(string(text))
	return err
}

type basePluginConfig struct {
	Enabled bool `toml:"enabled"`
	Threads int  `toml:"threads"`
}

type flashbackConfig struct {
	basePluginConfig
	MinSeconds int `toml:"min_seconds" default:"10"`
	MaxSeconds int `toml:"max_seconds" default:"20"`
}

type readLongConfig struct {
	basePluginConfig
	Interval duration `toml:"interval"`
	Total    duration `toml:"total"`
}

type largeTransferPluginConfig struct {
	basePluginConfig
	Count int `toml:"count"`
}

type backgroundTaskConfig struct {
	Enabled  bool     `toml:"enabled"`
	Interval duration `toml:"interval"`
}

type replicaFlashbackConfig struct {
	basePluginConfig
	MinSeconds      int    `toml:"min_seconds" default:"10"`
	MaxSeconds      int    `toml:"max_seconds" default:"20"`
	ReplicaDsn      string `toml:"replica_dsn" default:""`
	ReplicaReadHint string `toml:"replica_read_hint" default:""`
	SessionVar      string `toml:"session_var" default:""`
}

type replicaReadConfig struct {
	basePluginConfig
	ReplicaDsn      string `toml:"replica_dsn" default:""`
	ReplicaReadHint string `toml:"replica_read_hint" default:""`
	SessionVar      string `toml:"session_var" default:""`
}

type cdcReadConfig struct {
	basePluginConfig
	DownstreamDsn string `toml:"downstream_dsn" default:""`
}

type Config struct {
	RowCount                 int    `toml:"row_count"`
	InitialBalance           int    `toml:"initial_balance"`
	Verbose                  bool   `toml:"verbose" default:"true"`
	EnableCts                bool   `toml:"enable_cts" default:"false"`
	EnableAsyncCommit        bool   `toml:"enable_async_commit" default:"false"`
	EnableSsot               bool   `toml:"enable_ssot" default:"false"`
	ForXDB                   bool   `toml:"for_mysql" default:"false"`
	Dsn                      string `toml:"dsn" default:""`
	IgnoreReadError          bool   `toml:"ignore_read_error" default:"false"`
	ReplicaStrongConsistency bool   `toml:"replica_strong_consistency" default:"false"`
	CreateTableSuffix        string `toml:"create_table_suffix" default:""`
	DbType                   string `toml:"db_type" default:"mysql"`

	CheckBalance        basePluginConfig `toml:"check_balance"`
	ReadSnapshot        basePluginConfig `toml:"read_snapshot"`
	ReadTooOldSnapshot  basePluginConfig `toml:"read_too_old_snapshot"`
	ReadLong            readLongConfig   `toml:"read_long"`
	CheckSecIdx         basePluginConfig `toml:"check_secondary_index"`
	ReadCurrentSnapshot basePluginConfig `toml:"read_current_snapshot"`

	TransferBasic    basePluginConfig          `toml:"transfer_basic"`
	TransferOnePhase basePluginConfig          `toml:"transfer_one_phase"`
	TransferTwoXA    basePluginConfig          `toml:"transfer_two_xa"`
	TransferSimple   basePluginConfig          `toml:"transfer_simple"`
	TransferLarge    largeTransferPluginConfig `toml:"transfer_large"`

	RunEvict   backgroundTaskConfig `toml:"run_evict"`
	RunPurge   backgroundTaskConfig `toml:"run_purge"`
	ShowStatus backgroundTaskConfig `toml:"show_status"`
	Heartbeat  backgroundTaskConfig `toml:"heartbeat"`

	ReplicaRead  replicaReadConfig `toml:"replica_read"`
	CheckVersion replicaReadConfig `toml:"check_version"`

	CheckSessionHint   basePluginConfig       `toml:"session_hint"`
	CheckFlashback     flashbackConfig        `toml:"flashback_query"`
	ReplicaSessionHint replicaReadConfig      `toml:"replica_session_hint"`
	ReplicaFlashback   replicaFlashbackConfig `toml:"replica_flashback_query"`
	CheckCdc           cdcReadConfig          `toml:"check_cdc"`

	Paritions int `toml:"partitions"`
}

func (c *Config) Normalize() {
	if c.RowCount == 0 {
		c.RowCount = 100
	}
	if c.InitialBalance == 0 {
		c.InitialBalance = 1000
	}
	if c.ReadLong.Enabled {
		if c.ReadLong.Interval == 0 {
			c.ReadLong.Interval = duration(5 * time.Second)
		}
		if c.ReadLong.Total == 0 {
			c.ReadLong.Total = duration(600 * time.Second)
		}
	}
}
