// process.env 객체를 통해 모든 환경변수를 JSON 형태로 접근할 수 있다.
console.log("💻 모든 환경변수:", process.env);

// 특정 환경변수 값에 접근
// 운영체제마다 이름이 다르다. (Windows: Path, macOS/Linux: PATH) -전부 대문자 (PATH)
// --> `||` (OR 연산자)를 사용하면 Path가 없을 경우 PATH로 대체하여 조회한다.
const path = process.env.Path || process.env.PATH;

console.log(path);
