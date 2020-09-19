package infrastructure

type QueueInfoDTO struct {
	QueueId     int64  `json:"queueId"`
	Map         string `json:"map"`
	Description string `json:"description"`
}

type Assets struct {
	Queues []QueueInfoDTO
}
