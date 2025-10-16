package models

import "fmt"

type Host struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type HostManager struct {
	hosts  []Host
	nextID int
}

func NewHostManager() *HostManager {
	return &HostManager{
		hosts:  make([]Host, 0),
		nextID: 1,
	}
}

func (hm *HostManager) AddHost(hostname, username, password string, port int) Host {
	host := Host{
		ID:       hm.nextID,
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
		Name:     fmt.Sprintf("%s@%s", username, hostname),
	}

	hm.hosts = append(hm.hosts, host)
	hm.nextID++
	return host
}

func (hm *HostManager) GetHosts() []Host {
	return hm.hosts
}

func (hm *HostManager) GetHostByID(id int) *Host {
	for i, host := range hm.hosts {
		if host.ID == id {
			return &hm.hosts[i]
		}
	}
	return nil
}

func (hm *HostManager) RemoveHost(id int) bool {
	for i, host := range hm.hosts {
		if host.ID == id {
			hm.hosts = append(hm.hosts[:i], hm.hosts[i+1:]...)
			return true
		}
	}
	return false
}
