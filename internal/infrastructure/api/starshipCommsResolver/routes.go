package starshipcommsresolver

func (server *Server) MapRoutes() error {
	if err := server.TopSecretRoutes(server.engine); err != nil {
		return err
	}
	if err := server.TopSecretSplitRoutes(server.engine); err != nil {
		return err
	}
	return nil

}
