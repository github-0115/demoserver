package labeltaskdetail

type LabelTaskDetail struct {
	Id         `json:"id"`
	ImagePath  `json:"image_path"`
	CreatedAt  `json:"created_at"`
	FinishedAt `json:"finished_at"`
	TaskId     `json:"task_id"`
	Url        `json:"url"`
	Md5        `json:"md5"`
}
