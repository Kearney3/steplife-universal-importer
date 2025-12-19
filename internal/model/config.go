package model

type Config struct {
	EnableInsertPointStrategy int     `ini:"enableInsertPointStrategy"`
	InsertPointDistance       int     `ini:"insertPointDistance"`
	PathStartTime             string  `ini:"pathStartTime"`
	PathEndTime               string  `ini:"pathEndTime"`
	PathStartTimestamp        int64
	PathEndTimestamp          int64
	DefaultAltitude           float64 `ini:"defaultAltitude"`
	SpeedMode                 string  `ini:"speedMode"` // "auto" or "manual"
	ManualSpeed               float64 `ini:"manualSpeed"`
	EnableBatchProcessing     int     `ini:"enableBatchProcessing"`
}
