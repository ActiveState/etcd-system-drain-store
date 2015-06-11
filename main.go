package main

import (
	"crypto/sha1"
	"flag"
	"fmt"

	"github.com/cloudfoundry/gunk/workpool"
	"github.com/cloudfoundry/storeadapter"
	"github.com/cloudfoundry/storeadapter/etcdstoreadapter"
)

const (
	EtcdMaxConcurrentRequests = 10
	AppId                     = "system"
)

var (
	deleteKey = flag.Bool("delete", false, "flag to indicate deletion")
	etcdUrl   = flag.String("etcdUrl", "http://127.0.0.1:4001", "etcd url")
	syslogUrl = flag.String("syslogUrl", "", "url for the syslog drain e.g syslog://192.168.6.144:9003")
)

func main() {
	flag.Parse()

	if *syslogUrl == "" {
		panic("syslogUrl can not be empty!")
	}

	workPool := workpool.NewWorkPool(EtcdMaxConcurrentRequests)
	storeAdapter := etcdstoreadapter.NewETCDStoreAdapter([]string{*etcdUrl}, workPool)
	storeAdapter.Connect()

	key := drainKey(AppId, *syslogUrl)

	systemNode := storeadapter.StoreNode{
		Value: []byte(*syslogUrl),
		Key:   key,
	}

	if *deleteKey == true {
		storeAdapter.Delete(key)
	} else {

		storeAdapter.Create(systemNode)
	}

}

func appKey(appId string) string {
	return fmt.Sprintf("/loggregator/services/%s", appId)
}

func drainKey(appId string, drainUrl string) string {
	hash := sha1.Sum([]byte(drainUrl))
	return fmt.Sprintf("%s/%x", appKey(appId), hash)
}
