package logger

import "log"

func Error(v ...any) {
	log.Println(append([]any{"[Error]"}, v...)...)
}

func Warn(v ...any) {
	log.Println(append([]any{"[Warn]"}, v...)...)
}

func Info(v ...any) {
	log.Println(append([]any{"[Info]"}, v...)...)
}

func Debug(v ...any) {
	log.Println(append([]any{"[Debug]"}, v...)...)
}
