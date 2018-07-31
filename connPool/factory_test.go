package connpool

import (
	"fmt"
	"testing"
)

func TestNewFactory(t *testing.T) {
	factoryConfig := &FactoryConfig{connectTimeout: 10, connectRetries: 10, lazyCreate: false, protocol: "tcp", key: "127.0.0.1:8080"}
	factory := NewFactory(factoryConfig)
	_, err := factory.Create()
	if err != nil {
		fmt.Println(err)
	}
}
