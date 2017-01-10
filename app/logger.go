package app

import (
	"log"
	"os"
	"io"
)

const (
	LOGINFO = 0
	LOGERR  = 1
)

type Logger struct {
	infoFile   string
	infoFlags  int
	infoPrefix string
	errFile    string
	errFlags   int
	errPrefix  string
	fileOpen   int
	writer     *os.File
}

type LogMsg struct {
	logType int
	msg     string
	logger  *Logger
}

func NewLogMsg(logType int, msg string, logger *Logger) *LogMsg {
	return &LogMsg{logType, msg, logger,}
}

func (lm *LogMsg)Do() {
	if lm.logType == LOGINFO {
		lm.logger.Info(lm.msg)
	} else if lm.logType == LOGERR{
		lm.logger.Error(lm.msg)
	}
}

func NewLogger(infoFile string, errFile string) *Logger {
	return &Logger{
		infoFile:   infoFile,
		infoPrefix: "[Info]",
		infoFlags:  log.LstdFlags,
		errFile:    errFile,
		errPrefix:  "[Error]",
		errFlags:   log.LstdFlags,
		writer:     nil,
		fileOpen:   0,
	}
}

func (logger *Logger)OpenWriter(logType int) *log.Logger{
	var writer io.Writer
	var prefix string
	var flags  int

	if logType == LOGINFO {
		if logger.infoFile == "" {
			writer = os.Stdout
		} else {
			logger.fileOpen = 1
			writer, _ = os.Create(logger.infoFile)
		}
	} else {
		if logger.errFile == "" {
			writer = os.Stderr
		} else {
			logger.fileOpen = 1
			writer, _ = os.Create(logger.errFile)
		}
	}

	return log.New(writer, prefix, flags)
}

func (logger *Logger)CloseWriter(logType int) {
	if logger.fileOpen == 1 {
		logger.writer.Close()
	}

}

func (logger *Logger)SetInfoFlags(flags int) {
	logger.infoFlags = flags
}

func (logger *Logger)SetErrFlags(flags int) {
	logger.errFlags = flags
}

func (logger *Logger)SetInfoPrefix(prefix string) {
	logger.infoPrefix = prefix
}

func (logger *Logger)SetErrPrefix(prefix string) {
	logger.errPrefix = prefix
}

func (logger *Logger)Info(msg string) {
	info := logger.OpenWriter(LOGINFO)
	defer logger.CloseWriter(LOGINFO)
	info.Println(msg)
}

func (logger *Logger)Error(msg string) {
	err := logger.OpenWriter(LOGERR)
	defer logger.CloseWriter(LOGERR)
	err.Println(msg)
}

