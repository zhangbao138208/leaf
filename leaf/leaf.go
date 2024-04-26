package leaf

import (
	"os"
	"os/signal"

	"github.com/golang/glog"
	"github.com/zhangbao138208/leaf/leaf/module"
)

func Run(mods ...module.Module) {

	glog.Errorf("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	glog.Errorf("Leaf closing down (signal: %v)", sig)
	module.Destroy()
}
