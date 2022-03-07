package chat

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var (
	Subscribe   = make(chan (chan<- Subscription), 10) // 구독 채널
	Unsubscribe = make(chan (<-chan Event), 10)        // 구독 해지 채널
	Publish     = make(chan Event, 10)                 // 이벤트 발행 채널
)

func Chat() {
	server, err := socketio.NewServer(nil) // socket.io 초기화
	if err != nil {
		fmt.Println("socketio error >> ", err)
		log.Fatal(err)
	}

	go Chatroom() // 채팅방을 처리할 함수 고루틴으로 실행

	// 웹 브라우저에서 socket.io로 접속했을 때 실행할 콜백 설정
	server.On("connection", func(so socketio.Socket) {
		// 웹 브라우저가 접속되면
		s := Subscribe() // 구독 처리
		Join(so.Id())    // 사용자가 채팅방에 들어왔다는 이벤트 발생

		for _, event := range s.Archive { // 지금까지 쌓인 이벤트를 웹 브라우저로 접속한 사용자에게 보냄
			so.Emit("event", event)
		}

		newMessages := make(chan string)

		// 웹 브라우저에서 보내오는 채팅 메세지를 받을 수 있도록 콜백 설정
		so.On("message", func(msg string) {
			newMessages <- msg
		})

		// 웹 브라우저에서 보내오는 채팅 메세지를 받을 수 있도록 콜백 설정
		so.On("message", func(msg string) {
			newMessages <- msg
		})

		// 웹 브라우저의 접속이 끊어졌을 때 콜백 설정
		so.On("disconnnection", func() {
			Leave(so.Id())
			s.Cancel()
		})

		go func() {
			for {
				select {
				case event := s.New: // 채널에 이벤트가 들어오면 이벤트 데이터를 웹 브라우저에 보냄
					so.Emit("event", event)
				case msg := <-newMessages: //  웹 브라우저에서 채팅 메세지를 보내오면 채팅 메세지 이벤트 발생
					Say(so.Id(), msg)
				}
			}
		}()
	}) // end server.On

	http.Handle("/socket.io/", server) // /socket.io/ 경로는 socket.to 인스턴스가 처리하도록 설정

	http.Handle("/", http.FileServer(http.Dir("."))) // 현재 디렉토리를 파일 서버로 설정

	http.ListenAndServe("80", nil) // 80번 포트에서 웹 서버 실행
}

// 이벤트 생성 함수
func NewEvent(evtType, user, msg string) Event {
	return Event{evtType, user, int(time.Now().Unix()), msg}
} // end func NewEvent

// 새로운 사용자가 들어왔을 때 이벤트를 구독할 함수
func Subscribe() Subscription {
	// 채널을 생성하여 구독 채널에 보낸다
	sCh := make(chan Subscription)
	subscribe <- ch

	return <-ch

} // end func Subscribe

// 사용자가 나갔을 때 구독을 취소할 함수
func (s Subscription) Cancel() {
	unsubscribe <- s.New // 구독 해지 채널에 보냄

	for {
		select {
		case _, ok := <-s.New: // 채널에서 값을 모두 꺼냄
			if !ok {
				return
			}
		default:
			return
		}
	}

} // end func Cancel

// 사용자가 들어왔을 때 이벤트 발행
func Join(user string) {
	publish <- NewEvent("join", user, "")
} // end func Join

// 사용자가 채팅 메세지를 보냈을 때 이벤트 발생
func Say(user, message string) {
	publish <- NewEvent("message", user, message)
}

// 사용자가 나갔을 때 이벤트 발행
func Leave(user string) {
	publish <- NewEvent("leave", user, "")
}

// 구독, 구독 해지, 발행된 이벤트를 처리할 함수
func Chatroom() {
	archive := list.New()     // 쌓인 이벤트를 저장할 연결 리스트
	subscrivers := list.New() // 구독자 목록을 저장할 연결 리스트

	for {
		select {
		case c := <-subscribe: // 새로운 사용자가 들어왔을 때
			var events []Event
			for e := archive.Front(); e != nil; e = e.Next() { // 쌓인 이벤트가 있다면
				// events 슬라이스에 이벤트를 저장
				events = append(events, e.Value.(Event))
			}

			subscriber := make(chan Event, 10)    // 이벤트 채널 생성
			subscribers.PushBack(subscriber)      // 이벤트 채널을 구독자 목록에 추가
			c <- Subscription{events, subscriber} // 구독 구조체 인스턴스를 생성하여 채널 c에 보냄

		case event := <-Publish: // 새 이벤트가 발행되었을 때
			// 모든 사용자에게 이벤트 전달
			for e := subscribers.Front(); e != nil; e = e.Next() {
				// 구독자 목록에서 이벤트 채널을 꺼냄
				subscriber := e.Value.(chan Event)

				// 방금 받은 이벤트를 이벤트 채널에 보냄
				subscriber <- event
			}

			if archive.Len() >= 20 { // 저장된 이벤트 개수가 20개가 넘으면 이벤트 삭제
				archive.Remove(archive.Front())
			}
			archive.PushBack(event) // 현재 이벤트를 저장

		case C := <-unsubscribe: // 사용자가 나갔을 때
			for e := subscribers.Front(); e != nil; e = e.Next() {
				subscriber := e.Value.(chan Event) // 구독자 목록에서 이벤트 채널을 꺼낸다.

				if subscriber == c { // 구독자 목록에 들어 있는 이벤트와 채널 c가 같으면 구독자 목록에서 삭제
					subscriber.Remove(e)
					break
				}

			}

		} // end select

	} // end for

} // end func Chatroom()
