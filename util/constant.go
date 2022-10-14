package util

const (
	ARG_VERSION   = "version"
	BASE_URL      = "127.0.0.1:7799/service-t-k8s/v1/"
	COMMAND_CHECK = "check"
	COMMAND_INIT  = "init"
	COMMAND_POST  = "post"
	COMMAND_RUN   = "run"
)

var CommandList = map[string][]string{
	"cli":   {"check", "init", "post", "run"},
	"check": {"userdata", "system"},
	"init":  {"system"},
	"post":  {"nexus", "harbor", "image", "rancher", "certs", "kubectl"},
	"run":   {"nexus", "harbor", "image", "rancher", "kubectl"},
}

var ObjectList = map[string]string{
	"userdata": "userdata",
	"system":   "system",
	"nexus":    "file",
	"harbor":   "file",
	"image":    "file",
	"certs":    "file",
	"kubectl":  "file",
	"rancher":  "rancher",
}

type Rsp struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
