package core

import (
	"github.com/cloudquery/cloudquery/pkg/config"
	"os"
)

func ConfigureProxy(pCfg *config.Proxy) {
	if pCfg == nil {
		return
	}

	if pCfg.HttpProxy != nil {
		os.Setenv("HTTP_PROXY", *pCfg.HttpProxy)
	}
	if pCfg.HttpsProxy != nil {
		os.Setenv("HTTPS_PROXY", *pCfg.HttpsProxy)
	}
	if pCfg.RequestMethod != nil {
		os.Setenv("REQUEST_METHOD", *pCfg.HttpsProxy)
	}
}
