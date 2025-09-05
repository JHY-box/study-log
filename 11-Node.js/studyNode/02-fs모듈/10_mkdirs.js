/**
 * mkdirs 라이브러리
 * mkdirs는 지정된 경로에 필요한 모든 상위 디렉토리를 재귀적으로 생성해주는 라이브러리
 * 
 * 설치 : yarn add mkdirs
 * 
 * 🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰🟰
 * 
 * rmdir 라이브러리
 * rmdir은 내용의 유무와 상관없이 지정된 디렉토리와 그 하위 항목을 모두 재귀적으로 삭제
 *  내부적으로 rm -rf 명령을 수행하는 것과 유사
 * 
 * 설치 : yarn add rmdir
 */

import fs from "fs"; // FileSystem 모듈 참조

// 지정된 경로를 따라 폴더를 생성하는 라이브러리
// $ yarn add mkdirs
import { mkdirs } from "mkdirs";

// 지정된 경로를 따라 폴더와 그 하위 항목을 모두 삭제하는 라이브러리
// 내부적으로 "rm -rf" 명령을 수행함.
// $yarn add rmdir
import rmdir from "rmdir";

if (!fs.existsSync("./this")) {
	// 현재폴더(./)는 VSCode에서 열려있는 Workspace 폴더 기준
	// 이 라이브러리는 동기식 처리로 되어 있다.
	mkdirs("./this/that/and/the/other");
} else {
	// 이 라이브러리는 비동기식 처리로 되어 있다.
	(async () => {
		try {
			await rmdir("./this");
			console.log("this 폴더를 삭제했습니다.");
		} catch (err) {
			console.log(err);
		}
	})();
}
