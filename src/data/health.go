package data

const (
	// ServiceRunning represents service is Running
	ServiceRunning = "Running"
	// ServiceDegraded represents service health is Degraded
	ServiceDegraded = "Degraded"
	// ServiceStopped represents service is Stopped
	ServiceStopped = "Stopped"
	// ConnectionActive represents connection is active
	ConnectionActive = "Active"
	// ConnectionDisconnected represents connection is disconnected
	ConnectionDisconnected = "Disconnected"
)

// Health represents health response
type Health struct {
	TimeStampUTC     string `json:"timeStampUTC,omitempty"`
	ServiceName      string `json:"serviceName,omitempty"`
	ServiceProvider  string `json:"serviceProvider,omitempty"`
	ServiceVersion   string `json:"serviceVersion,omitempty"`
	ServiceStatus    string `json:"serviceStatus,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Hostname         string `json:"hostname,omitempty"`
	OS               string `json:"OS,omitempty"`
}
