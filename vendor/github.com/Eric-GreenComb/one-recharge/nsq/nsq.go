package nsq

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

// Producer global nsq Producer
var Producer *nsq.Producer

// Init Init
func Init() error {
	_config := nsq.NewConfig()
	Producer, err := nsq.NewProducer("127.0.0.1:4150", _config)
	if err != nil {
		fmt.Println("nsq err :", Producer.String())
		return err
	}
	return err
}
