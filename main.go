package main

import (
    "prom_exporter/server"
	"net/http"
    log "prom_exporter/logger"
    "time"
    "prom_exporter/handle"
)

func main() {
    // 初始化日志
    log.InitLog("./log/exporter.log", "INFO")

    // 初始化mux
    mux := server.New()
    log.Infoln("MUX 路由初始化完成")

    // URL映射
    handle.InitHandle(mux)
    log.Infoln("MUX 路由注册完成")

	srv := &http.Server{
		Handler:      mux.GetRouter(),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Hour,
		ReadTimeout:  15 * time.Hour,
	}

    srv.ListenAndServe()
}
