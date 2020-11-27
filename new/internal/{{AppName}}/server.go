package {{AppName}}

import (
	"context"
	"fmt"
	"{{ProjectName}}/internal/{{AppName}}/config"
	"{{ProjectName}}/internal/{{AppName}}/model"
	"{{ProjectName}}/pkg/client/database"
	"{{ProjectName}}/pkg/log"
	"net/http"
)

type Server struct {
	Config *config.Cfg
	Server *http.Server
	err    error
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	s.initCfg()
	s.initLog()
	s.initDB(stopCh)
	s.initHttpServer()
	return s.err
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
		log.Info(fmt.Sprintf("Shutdown server"))
		_ = s.Server.Shutdown(ctx)
	}()
	log.Info(fmt.Sprintf("Start listening on %s", s.Server.Addr))
	err = s.Server.ListenAndServe()
	return nil
}

func (s *Server) initCfg() {
	if s.err != nil {
		return
	}
	s.Config, s.err = config.TryLoadFromDisk()
}

func (s *Server) initDB(stopCh <-chan struct{}) {
	if s.err != nil {
		return
	}
	var c *database.Client
	log.Info(fmt.Sprintf("db init"))
	c, s.err = database.NewDatabaseClient(s.Config.Mysql, stopCh)
	log.Info(fmt.Sprintf("db init over"))
	model.MainDB = c.DB()
	if model.MainDB != nil {
		s.migrate()
	}
}

func (s *Server) initHttpServer() {
	if s.err != nil {
		return
	}
	s.Server = new(http.Server)
	s.Server.Addr = s.Config.Server.Addr
}

func (s *Server) initLog() {
	if s.err != nil {
		return
	}
	s.err = log.NewLog(s.Config.Log)
}

func (s *Server) migrate() {
	model.MainDB.AutoMigrate(
		//new(model.App),
	)
}