package klogs

const (
	CategoryDefault = "DEFAULT"
	CategoryMysql   = "MYSQL"
	CategoryMongo   = "MONGO"
	CategoryRedis   = "REDIS"
	CategoryHttp    = "HTTP"
)

func Category(name string) Logger {
	return Logger{category: name}
}

func Default() Logger {
	return Category(CategoryDefault)
}

func Mysql() Logger {
	return Category(CategoryMysql)
}

func Mongo() Logger {
	return Category(CategoryMongo)
}

func Redis() Logger {
	return Category(CategoryRedis)
}

func Http() Logger {
	return Category(CategoryHttp)
}

func Debug(format string, args ...interface{}) {
	Category(CategoryDefault).Debug(format, args)
}

func Info(format string, args ...interface{}) {
	Category(CategoryDefault).Info(format, args)
}

func Warn(format string, args ...interface{}) {
	Category(CategoryDefault).Warn(format, args)
}

func Error(format string, args ...interface{}) {
	Category(CategoryDefault).Error(format, args)
}
