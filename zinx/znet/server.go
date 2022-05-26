package znet

// 实现IServer接口  （依次实现了IServer接口中的所有方法即可）
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的IP版本
	IPVersion string
	// 服务器IP
	IP string
	// 服务器端口号
	Port int
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}
