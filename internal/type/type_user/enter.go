package type_user

type UserInfoRequest struct {
	ID string `form:"id"`
}
type UserInfoResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID int64 `json:"id,string"`
}

type UserUpdateRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDeleteRequest struct {
	ID string `form:"id"`
}
