package ip

import (
	"log"
	"net"
	"time"
)

func WithLogging(logger *log.Logger)  {
	if logger == nil {
		logger = log.Default()
	}
	DefaultProvider = &loggerProvider{
		DefaultProvider,
		logger,
	}
	GetIpChanges = decorateIpChanges(logger, getIpChangesFromDefaultProvider)
}

type loggerProvider struct {
	prov Provider
	logger *log.Logger
}

func (l *loggerProvider) ObtainIp() (net.IP, error) {
	l.logger.Printf("querying for current ip address")
	ip, err := l.prov.ObtainIp()
	if err != nil {
		l.logger.Printf("error querying for ip address: %s", err.Error())
		return ip, err
	}
	l.logger.Printf("current ip address is %s", ip.String())
	return ip, err
}

func decorateIpChanges(logger *log.Logger, next func(time.Duration) Updates) func(time.Duration) Updates {
	return func(interval time.Duration) Updates {
		log.Printf("Watching for ip changes every %s", interval.String())
		ch := make(chan net.IP)

		go func() {
			innerCh := next(interval)
			for {
				ip := <-innerCh
				logger.Printf("Ip address has changed to %s", ip.String())
				ch <- ip
			}
		}()

		return ch
	}
}
