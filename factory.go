package tailtracer

import (
	"context"
	"strconv"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
)

const (
	typeStr = "tailtracer"
	defaultInterval = 1
)

func CreateDefaultConfig() component.ReceiverConfig {

	return &Config{
		ReceiverSettings: config.NewReceiverSettings(component.NewID(typeStr)),
		Interval: strconv.Itoa(defaultInterval),
	}
}

// func CreateTracesReceiver(_ context.Context, params component.ReceiverCreateSettings, bcfg Config,
// 	nextConsumer Component.consumer.Traces) (component.TracesReceiver, error){
// 		return nil, nil
// 	}
func createTracesReceiver(s context.Context, p component.ReceiverCreateSettings, baseCfg component.Config, consumer consumer.Traces) (component.TracesReceiver, error) {
		return nil,nil
	  }

  
func NewFactory()component.ReceiverFactory{
	return component.NewReceiverFactory(
		typeStr,
		CreateDefaultConfig,
		component.WithTracesReceiver(createTracesReceiver))

	)

}