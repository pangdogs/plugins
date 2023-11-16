package zap_log

import (
	"go.uber.org/zap"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/option"
	"kit.golaxy.org/golaxy/util/types"
	"kit.golaxy.org/plugins/log"
	"reflect"
)

func newLogger(settings ...option.Setting[LoggerOptions]) log.Logger {
	return &_Logger{
		options: option.Make(Option{}.Default(), settings...),
	}
}

type _Logger struct {
	options       LoggerOptions
	sugaredLogger *zap.SugaredLogger
}

// InitSP 初始化服务插件
func (l *_Logger) InitSP(ctx service.Context) {
	options := []zap.Option{zap.AddCallerSkip(l.options.CallerSkip)}
	if l.options.ServiceInfo {
		options = append(options, zap.Fields(zap.String("service", ctx.String())))
	}
	l.sugaredLogger = l.options.ZapLogger.WithOptions(options...).Sugar()

	log.Infof(ctx, "init service plugin %q with %q", plugin.Name, types.AnyFullName(*l))
}

// ShutSP 关闭服务插件
func (l *_Logger) ShutSP(ctx service.Context) {
	log.Infof(ctx, "shut service plugin %q", plugin.Name)
}

// InitRP 初始化运行时插件
func (l *_Logger) InitRP(ctx runtime.Context) {
	options := []zap.Option{zap.AddCallerSkip(l.options.CallerSkip)}
	if l.options.ServiceInfo {
		options = append(options, zap.Fields(zap.String("service", ctx.String())))
	}
	if l.options.RuntimeInfo {
		options = append(options, zap.Fields(zap.String("runtime", ctx.String())))
	}
	l.sugaredLogger = l.options.ZapLogger.WithOptions(options...).Sugar()

	log.Infof(ctx, "init runtime plugin %q with %q", plugin.Name, reflect.TypeOf(_Logger{}))
}

// ShutRP 关闭运行时插件
func (l *_Logger) ShutRP(ctx runtime.Context) {
	log.Infof(ctx, "shut runtime plugin %q", plugin.Name)
}

// Log writes a log entry, spaces are added between operands when neither is a string and a newline is appended
func (l *_Logger) Log(level log.Level, v ...interface{}) {
	sugaredLogger := l.sugaredLogger

	switch level {
	case log.TraceLevel:
		sugaredLogger.Debug(v...)
	case log.DebugLevel:
		sugaredLogger.Debug(v...)
	case log.InfoLevel:
		sugaredLogger.Info(v...)
	case log.WarnLevel:
		sugaredLogger.Warn(v...)
	case log.ErrorLevel:
		sugaredLogger.Error(v...)
	case log.DPanicLevel:
		sugaredLogger.DPanic(v...)
	case log.PanicLevel:
		sugaredLogger.Panic(v...)
	case log.FatalLevel:
		sugaredLogger.Fatal(v...)
	}
}

// Logln writes a log entry, spaces are always added between operands and a newline is appended
func (l *_Logger) Logln(level log.Level, v ...interface{}) {
	sugaredLogger := l.sugaredLogger

	switch level {
	case log.TraceLevel:
		sugaredLogger.Debugln(v...)
	case log.DebugLevel:
		sugaredLogger.Debugln(v...)
	case log.InfoLevel:
		sugaredLogger.Infoln(v...)
	case log.WarnLevel:
		sugaredLogger.Warnln(v...)
	case log.ErrorLevel:
		sugaredLogger.Errorln(v...)
	case log.DPanicLevel:
		sugaredLogger.DPanicln(v...)
	case log.PanicLevel:
		sugaredLogger.Panicln(v...)
	case log.FatalLevel:
		sugaredLogger.Fatalln(v...)
	}
}

// Logf writes a formatted log entry
func (l *_Logger) Logf(level log.Level, format string, v ...interface{}) {
	sugaredLogger := l.sugaredLogger

	switch level {
	case log.TraceLevel:
		sugaredLogger.Debugf(format, v...)
	case log.DebugLevel:
		sugaredLogger.Debugf(format, v...)
	case log.InfoLevel:
		sugaredLogger.Infof(format, v...)
	case log.WarnLevel:
		sugaredLogger.Warnf(format, v...)
	case log.ErrorLevel:
		sugaredLogger.Errorf(format, v...)
	case log.DPanicLevel:
		sugaredLogger.DPanicf(format, v...)
	case log.PanicLevel:
		sugaredLogger.Panicf(format, v...)
	case log.FatalLevel:
		sugaredLogger.Fatalf(format, v...)
	}
}
