package utils

type (
	Conf struct {
		Api      Api      `json:"api"`
		Client   Client   `json:"client"`
		Database Database `json:"database"`
	}
	Api struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Client struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
	}
)
