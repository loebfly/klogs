package klogs

var (
	null = C("")
)

func Init(ymlPath string) error {
	return config.init(ymlPath)
}

func C(name string) Logger {
	return Logger{category: name}
}

func Debug(format string, args ...interface{}) {
	null.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	null.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	null.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	null.Error(format, args...)
}
