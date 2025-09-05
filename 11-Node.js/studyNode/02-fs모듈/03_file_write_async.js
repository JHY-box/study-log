/** (1) 모듈참조, 필요한 변수 생성 */
import fs from "fs"; // FileSystem 모듈 참조

const target = "./output.txt"; // 파일경로
const content = "Hello World"; // 저장할 내용
const is_exists = fs.existsSync(target); // 파일의 존재 여부 검사


if (!is_exists) {
	/** (2) 파일이 존재하지 않을 경우 새로 저장 */
	// 절대경로 지정, 비동기식 파일 저장
	(async () => {
		console.log(target + "의 파일 저장을 요청했습니다.");
		// async~await는 비동기식 처리를 동기식으로 작동하도록 제어함
		try {
            await fs.promises.writeFile(target,content,"utf8");
			console.log(target + "에 데이터 쓰기 완료.");

			await fs.promises.chmod(target, "0766");
			console.log(target + "의 퍼미션 설정 완료");
		} catch (err) {
			console.log(err);
			return;
		}
	})();
} else {
	/** (3) 파일이 존재할 경우 파일 삭제 */
	console.log(target + "의 파일 삭제를 요청했습니다.");
    
	(async () => {
		try {
			await fs.promises.unlink(target);
			console.log(target + "의 파일 삭제 완료");
		} catch (err) {
			console.log(err);
			return;
		}
	})();
}
