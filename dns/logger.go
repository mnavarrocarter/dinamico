package dns

import (
	"log"
	"net"
	"net/url"
)

func WithLogging(logger *log.Logger)  {
	if logger == nil {
		logger = log.Default()
	}
	DefaultUpdater = &loggerUpdater{
		updater: DefaultUpdater,
		logger:  logger,
	}
}

type loggerUpdater struct {
	updater RecordUpdater
	logger *log.Logger
}

func (lu *loggerUpdater) Update(record *url.URL, ip net.IP) error {
	lu.logger.Printf("updating hostname %s with ip address %s", record.Host, ip.String())
	err := lu.updater.Update(record, ip)
	if err != nil {
		lu.logger.Printf("error while updating hostname %s with ip %s: %s", record.Host, ip.String(), err.Error())
		return err
	}
	lu.logger.Printf("updated hostname %s with ip address %s successfully", record.Host, ip.String())
	return nil
}

