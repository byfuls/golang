package channelMonitor

import (
	"net/http"
	"os"

	"program/cm/channelMonitor/router"
)

func Open() {
	m := router.MakeHandler(os.Getenv("COMMON_SRC") + "/cm/channelMonitor/database/user.db")
	defer m.Close()

	err := http.ListenAndServe(":2219", m)
	if err != nil {
		panic(err)
	}
}
