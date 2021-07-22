package codef

import (
	"go-gin-codef-api/src/config"

	ecg "github.com/codef-io/easycodefgo"
)

func GetCodefInstance() *ecg.Codef {
	config := config.InitConfig()

	codef := &ecg.Codef{
		PublicKey: config.CodefPublicKey,
	}

	codef.SetClientInfoForDemo(config.CodefClientId, config.CodefClientSecret)

	return codef
}

func GetServiceType(serviceType int) ecg.ServiceType {
	switch serviceType {
	case int(ecg.TypeProduct):
		return ecg.TypeProduct
	case int(ecg.TypeDemo):
		return ecg.TypeDemo
	default:
		return ecg.TypeSandbox
	}
}
