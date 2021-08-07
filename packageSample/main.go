// =======================
// golang 모듈 사용
// =======================
// go init mod [폴더명] 명령어로 .mod 파일 생성 후 사용가능
package main

import (
	greeting "packageSample/myPackage" // myPackage에서 패키지를 불러온다.
)

func main() {
	greeting.Say()

}
