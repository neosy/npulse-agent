package dto

type RegRequest struct {
	// Client
	IPAddress string `json:"ip_address"`
	HostName  string `json:"host_hame"`
}
