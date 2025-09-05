// 1. dotenv 모듈을 불러온다.
import dotenv from "dotenv"; // ES Module 방식 (최신)

// 이 함수가 호출되는 순간 .env 파일의 내용이 process.env에 저장된다.
dotenv.config();

// 2. .env 파일에 저장된 값들을 process.env를 통해 사용한다.
console.log("🗄 데이터베이스 설정:");
const dbConfig = {
	host: process.env.DB_HOST,
	port: process.env.DB_PORT,
	user: process.env.DB_USER,
	password: process.env.DB_PASSWORD,
	database: process.env.DB_NAME,
};

console.log(dbConfig);
console.log();

// 3. Boolean 타입 환경변수 처리
// .env 파일의 모든 값은 문자열("true", "false")이므로, Boolean으로 변환해야 한다.
console.log("✅ Boolean 타입 환경변수:");
const isDebugMode = process.env.DEBUG === "true";
console.log(`디버그 모드 활성화: ${isDebugMode}`);
console.log();
