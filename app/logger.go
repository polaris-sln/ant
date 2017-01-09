package app

import (
	"log"
	"os"
	"io"
)

const (
	LOGINFO := 0
	LOGERR  := 1
)

type Logger struct {
	infoFile   string
	infoFlags  int
	infoPrefix string
	errFile    string
	errFlags   int
	errPrefix  string
	logWriter  LogWriter
}

type LogWriter struct {
	Writer   io.Writer
}

type LogMsg struct {
	logType int
	msg     string
	logger  Logger
}

func NewLogMsg(logType int, msg string, logger Logger) LogMsg {
	return LogMsg{logType, msg, logger,}
}

func (lm *LogMsg)Do() {
	if lm.logType == LOGINFO {
		logger.Info(lm.msg)
	} else if lm.logType == LOGERR{
		logger.Error(lm.msg)
	}
}

func newLogWriter() *LogWriter {
	return &LogWriter{os.Stdout,}
}

func (lw *LogWriter)Open(fileName string, logType int, prefix string, flags int) log.Logger {
	if fileName == "" {
		if logType == LOGINFO {
			lw.Writer := os.Stderr
		} else {
			lw.Writer := os.Stdout
		}
	} else {
		lw.Writer = os.Create(fileName)
	}

	return log.New(lw.writer, prefix, flags),
}

func (lw *LogWriter)Close() {
	if fileName != "" {
	lw.Writer.Close()
	}
}


func NewLogger(infoFile string, errFile) Logger {
	return &Logger{
		infoFile:   infoFile,
		infoPrefix: "[Info]",
		infoFlags:  log.lstdFlags,
		errFile:    errFile,
		errPrefix:  "[Error]",
		errFlags:   log.lstdFlags,
		logWriter:  NewLogWriter(),
	}
}

func (logger *Logger)SetInfoFlags(flags int) {
	logger.infoFlags = flags
}

func (loger *Logger)SetErrFlags(flags int) {
	logger.errFlags = flags
}

func (logger *Logger)SetInfoPrefix(prefix string) {
	logger.infoPrefix = prefix
}

func (logger *Logger)SetErrPrefix(prefix string) {
	logger.errPrefix = prefix
}

func (logger *Logger)Info(msg string) {
	info := logger.logWriter.Open()
	defer logger.logWriter.Close()
	info.Println(msg)
}

func (logger *Logger)Error(msg string) {
	err := logger.logWriter.Open()
	defer logger.logWriter.Close()
	err.Println(msg)
}

