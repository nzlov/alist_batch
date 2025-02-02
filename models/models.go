package models

type Config struct {
	Url          string `yaml:"url"`
	Auth         *Auth  `yaml:"auth"`
	Token        string `yaml:"token"`
	RefreshToken string `yaml:"refresh_token"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Code    int
	Message string
	Data    *AuthData
}

type AuthData struct {
	Token string
}

type ResData struct {
	Code    int
	Message string
	Data    string
}

type PushData struct {
	MountPath       string `json:"mount_path"`
	Order           int    `json:"order"`
	Remark          string `json:"remark"`
	CacheExpiration int    `json:"cache_expiration"`
	WebProxy        bool   `json:"web_proxy"`
	WebdavPolicy    string `json:"webdav_policy"`
	DownProxyUrl    string `json:"down_proxy_url"`
	ExtractFolder   string `json:"extract_folder"`
	Driver          string `json:"driver"`
	Addition        string `json:"addition"`
}

type Addition struct {
	RefreshToken   string `json:"refresh_token"`
	ShareId        string `json:"share_id"`
	SharePwd       string `json:"share_pwd"`
	RootFolderId   string `json:"root_folder_id"`
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
}

type ShareList struct {
	TvSeries *TvSeries `yaml:"电视剧"`
	Movies   *Movies   `yaml:"电影"`
}

type TvSeries struct {
	XiYouJi      string `yaml:"西游记86版"`
	ShiZongZui   string `yaml:"十宗罪"`
	FaYiQingMing string `yaml:"法医秦明"`
}

type Movies struct {
	ManWei               string `yaml:"漫威合集"`
	XinHaiChengGongQiJun string `yaml:"新海诚&宫崎骏合集"`
	LinZhengYing         string `yaml:"林正英合集"`
}
