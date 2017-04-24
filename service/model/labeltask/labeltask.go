package labeltask

type LabelTask struct {
	Id            `json:"id"`
	Name          `json:"name"`
	ImageListPath `json:"image_list_path"`
	VerifyTimes   `json:"verify_times"`
	CreatedAt     `json:"created_at"`
	FinishedAt    `json:"finished_at"`
	TypeId        `json:"type_id"`
	State         `json:"state"`
	Done          `json:"done"`
	Total         `json:"total"`
	Exported      `json:"exported"`
	ModelVersion  `json:"model_version"`
	Priority      `json:"priority"`
	Group         `json:"group"`
	ImageFolder   `json:"image_folder"`
}
