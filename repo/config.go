package repo

type Config struct {
	AsyncTaskRunner *AsyncTaskConfig
}

type AsyncTaskConfig struct {
	RedisIp   string
	RedisPort int
}

var LocalConfig = &Config{
	AsyncTaskRunner: &AsyncTaskConfig{
		RedisIp:   "127.0.0.1",
		RedisPort: 6379,
	},
}
