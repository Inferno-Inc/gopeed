package libgopeed

// #cgo LDFLAGS: -static-libstdc++
import "C"
import (
	"encoding/json"
	"github.com/monkeyWie/gopeed/pkg/rest"
	"github.com/monkeyWie/gopeed/pkg/rest/model"
)

func Start(cfg string) (int, error) {
	var config model.StartConfig
	if err := json.Unmarshal([]byte(cfg), &config); err != nil {
		return 0, err
	}
	return rest.Start(&config)
}

func Stop() {
	rest.Stop()
}
