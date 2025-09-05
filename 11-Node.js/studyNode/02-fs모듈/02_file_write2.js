/** (1) 모듈참조, 필요한 변수 생성 */
import fs from "fs"; // FileSystem 모듈 참조

const target = "./output.txt"; // 파일경로
const content = "Hello World"; // 저장할 내용
const is_exists = fs.existsSync(target); // 파일의 존재 여부 검사


if (!is_exists) {
	/** (2) 파일이 존재하지 않을 경우 새로 저장 */
	// 절대경로 지정, 비동기식 파일 저장
	fs.writeFile(target, content, "utf8", (err) => {   // 콜백함수 사용 (저장 요청 후 바로 다음 코드 실행)
		if (err) {
			return console.log(err); // 에러가 있으면 메시지 출력하고 종료
		}
		console.log(target + "에 데이터 쓰기 완료.");

		// 퍼미션 설정
		fs.chmod(target, "0766", (err) => { 
			if (err) {
				return console.log(err);
			}
			console.log(target + "의 퍼미션 설정 완료");
		});
	});

	console.log(target + "의 파일 저장을 요청했습니다.");
    // 이런 콜백지옥 때문에 async/await 등이 나옴

} else {
	/** (3) 파일이 존재할 경우 파일 삭제 */
	fs.unlink(target, (err) => {
		if (err) {
			return console.log(err);
		}
		console.log(target + "의 파일 삭제 완료");
	});

	console.log(target + "의 파일 삭제를 요청했습니다.");
}
