package chat

type (
	// Define the struct of chat event
	Event struct {
		EvtType   string // Event type
		User      string
		Timestamp int
		Text      string // Message text
	}

	// Define the struct of subscription
	Subscription struct {
		Archive []Event      // 지금까지 쌓인 이벤트를 저장할 슬라이스
		New     <-chan Event // 새 이벤트가 생길 때마다 데이터를 받을 수 있도록 이벤트 채널 생성
	}
)
