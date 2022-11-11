package response

// swagger:response DashboardMenuResponse
type KplusUponOpenResponse struct {
	Action            *string `json:"action"`
	Title             *string `json:"title"`
	Message           *string `json:"message"`
	Sub_message       *string `json:"sub_message"`
	Image_url         *string `json:"image_url"`
	Show              *string `json:"show"`
	Created_date      *string `json:"created_date"`
	Id                *int    `json:"id"`
	Last_updated_by   *int    `json:"last_updated_by"`
	Last_updated_date *string `json:"last_updated_date"`
	Redirect_link     *string `json:"redirect_link"`
	Created_by        *int    `json:"created_by"`
}

type TryResponse struct {
	Action            *string `json:"action"`
	Title             *string `json:"title"`
	Message           *string `json:"message"`
	Sub_message       *string `json:"sub_message"`
	Image_url         *string `json:"image_url"`
	Show              *string `json:"show"`
	Created_date      *string `json:"created_date"`
	Id                *int    `json:"id"`
	Last_updated_by   *int    `json:"last_updated_by"`
	Last_updated_date *string `json:"last_updated_date"`
	Redirect_link     *string `json:"redirect_link"`
	Created_by        *int    `json:"created_by"`
}
