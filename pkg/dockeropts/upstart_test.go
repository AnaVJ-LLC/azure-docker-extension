package dockeropts

import (
	"testing"
)

func TestUpstartEditor_NoContent(t *testing.T) {
	out, err := UpstartCfgEditor{}.ChangeOpts("", "-d --tlsverify")
	if err != nil {
		t.Fatal(err)
	}
	expected := `DOCKER_OPTS="-d --tlsverify"`
	if out != expected {
		t.Fatal("out:%s\nexpected:%s", out, expected)
	}
}

func TestUpstartEditor_UbuntuDefault(t *testing.T) {
	in := `# Docker Upstart and SysVinit configuration file

# Customize location of Docker binary (especially for development testing).
#DOCKER="/usr/local/bin/dockerd"

# Use DOCKER_OPTS to modify the daemon startup options.
#DOCKER_OPTS="--dns 8.8.8.8 --dns 8.8.4.4"

# If you need Docker to use an HTTP proxy, it can also be specified here.
#export http_proxy="http://127.0.0.1:3128/"

# This is also a handy place to tweak where Docker's temporary files go.
#export TMPDIR="/mnt/bigdrive/docker-tmp"`

	expected := `# Docker Upstart and SysVinit configuration file

# Customize location of Docker binary (especially for development testing).
#DOCKER="/usr/local/bin/dockerd"

# Use DOCKER_OPTS to modify the daemon startup options.
DOCKER_OPTS="-d --tlsverify"

# If you need Docker to use an HTTP proxy, it can also be specified here.
#export http_proxy="http://127.0.0.1:3128/"

# This is also a handy place to tweak where Docker's temporary files go.
#export TMPDIR="/mnt/bigdrive/docker-tmp"`

	out, err := UpstartCfgEditor{}.ChangeOpts(in, "-d --tlsverify")
	if err != nil {
		t.Fatal(err)
	}
	if out != expected {
		t.Fatal("out:%s\nexpected:%s", out, expected)
	}
}
