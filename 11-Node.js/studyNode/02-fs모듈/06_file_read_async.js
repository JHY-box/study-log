import fs from "fs"; // FileSystem 모듈 참조

const target = "./output.txt"; // 파일경로

if (fs.existsSync(target)) {
	console.log(target + " 파일을 읽도록 요청했습니다.");
    
	(async () => {
		try {
			const data = await fs.promises.readFile(target,"utf8");
			console.log(data); // 읽어 들인 데이터 출력
		} catch (err) {
			console.log(err);
		}
	})();
} else {
	console.log(target + "파일이 존재하지 않습니다.");
}
