var socket =  io();  // socket.io 객체 생성

// 서버에서 이벤트가 왔을 때 실행할 콜백 함수 설정
socket.on('event', function(data) {
    var msg
    
    switch (data.EvtType) { // 이벤트 타입을 판별하여 메세지 생성
        case 'message':
            msg = data.User + ':' + data.Text;
            break;
        case 'join':
            msg = data.User + "  has joind"
            break;
        case 'leave':
            msg = data.User + " has left"
            break;
    }

    // div 태그를 생성하여 채팅 메세지를 넣어준다.
    col = $('<div>').addClass('col-ms-12').text(msg)
    row = $('<div>').addClass('row').append(col)
    list = $('#messageList').append(row)

    // 채팅 메세지가 15개를 넘어가면 메세지를 삭제한다.
    if (list.children().size() > 15) {
        list.find('div:first').remove();
    }

    // 채팅 메세지를 서버에 보내는 함수
    send = function() {
        msg = $('#message').val()   // 입력 상자에서 메세지를 가져온다.
        // 메세지가 있으면 서버에 메세지를 보낸다.
        if (msg != '') {
            socket.emit('message', msg);
            // 입력한 데이터 삭제
            $('#message').val('');
        }
    }

    // send 버튼으로 메세지를 보낸다.
    $('#send').click(function() {
        alert("클릭")
        send()
    });

    // 엔타키 입력으로 메세지를 보낸다.
    $('#message').keyup(function(e) {
        if (e.keyCode == 13) {  // 13 = 엔타키
            send()
        }
    });
})