package config

import "time"

// it is better to have etcd with watchers on vars,
// but I think it is enough for this pet

const (
	// APIPortDefault local deploy port
	APIPortDefault = "2022"

	// SectorDroneCronDuration sector-drone cron duration
	SectorDroneCronDuration = 9 * time.Second

	// MinDroneCount ...
	MinDroneCount = 1
	// MaxDroneCount ...
	MaxDroneCount = 101
)
