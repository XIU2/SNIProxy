package main

type configModel struct {
	ForwardRules  []string `yaml:"rules,omitempty"`
	ListenAddr    string   `yaml:"listen_addr,omitempty"`
	EnableSocks   bool     `yaml:"enable_socks5,omitempty"`
	SocksAddr     string   `yaml:"socks_addr,omitempty"`
	AllowAllHosts bool     `yaml:"allow_all_hosts,omitempty"`
}
