package biz

type UserProfile struct {
	Uuid      string
	Username  string
	Avatar    string
	School    string
	Company   string
	Job       string
	Homepage  string
	Introduce string
}

type UserProfileUpdate struct {
	UserProfile
	Status int32
}

type Credentials struct {
	TmpSecretID  string
	TmpSecretKey string
	SessionToken string
	StartTime    int64
	ExpiredTime  int64
}
