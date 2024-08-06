package log

type Option func(config *LoggerConfig)

// WithFilename set log filename
func WithFilename(filename string) Option {
	return func(cfg *LoggerConfig) {
		cfg.Filename = filename
	}
}

// WithLogDir set log dir
func WithLogDir(dir string) Option {
	return func(cfg *LoggerConfig) {
		cfg.LoggerDir = dir
	}
}
