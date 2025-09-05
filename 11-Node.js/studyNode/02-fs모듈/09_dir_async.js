import fs from "fs"; // FileSystem 모듈 참조

const target = "./docs";

if (!fs.existsSync(target)) {
	(async () => {
		try {
			await fs.promises.mkdir(target);
			await fs.promises.chmod(target, "0777");
			console.log("새로운 docs 폴더를 만들었습니다.");
		} catch (err) {
			console.log(err);
		}
	})();
} else {
	// 파일 삭제 --> 비어있지 않은 폴더는 삭제 못함.
	(async () => {
		try {
			await fs.promises.rmdir(target);
			console.log("docs 폴더를 삭제했습니다.");
		} catch (err) {
			console.log(err);
		}
	})();
}
