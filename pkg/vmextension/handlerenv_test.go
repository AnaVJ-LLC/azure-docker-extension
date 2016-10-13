package vmextension

import (
	"testing"
)

func Test_ParseHandlerEnv(t *testing.T) {
	json := `[{  "name": "MSOpenTech.Extensions.DockerExtension", "version": 1.0,  "handlerEnvironment": {    "logFolder": "/var/log/azure/MSOpenTech.Extensions.DockerExtension/0.6.0.0",    "configFolder": "/var/lib/waagent/MSOpenTech.Extensions.DockerExtension-0.6.0.0/config",    "statusFolder": "/var/lib/waagent/MSOpenTech.Extensions.DockerExtension-0.6.0.0/status",    "heartbeatFile": "/var/lib/waagent/MSOpenTech.Extensions.DockerExtension-0.6.0.0/heartbeat.log"}}]`
	c, err := ParseHandlerEnv([]byte(json))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Parsed: %#v", c)
}
