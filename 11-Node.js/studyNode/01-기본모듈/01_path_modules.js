// ctrl + alt + n : 실행
// ctrl + alt + m : 중지
// node 01_path_modules.js : 터미널에서 직접 실행

/** (1) 모듈참조 */
import path from 'path';

/** (2) 디렉토리명과 파일명 합치기 */
const currentPath = path.join('/Users/hello/world', '../photo.jpg');
console.log('파일 패스 : %s', currentPath);
console.log(); // 줄바꿈 (빈 줄 넣기)

/** (3) 경로에서 디렉토리, 파일명, 확장자 구분하기 (분리하기) */
const dirname = path.dirname(currentPath);
const basename = path.basename(currentPath);
const extname = path.extname(currentPath);
console.log('디렉토리 : %s', dirname);
console.log('파일 이름 : %s', basename);
console.log('확장자 : %s', extname);
console.log();

/** (4) 경로정보 파싱 */
const parse = path.parse(currentPath);
console.log('root: ' + parse.root);
console.log('dir: ' + parse.dir);
console.log('name: ' + parse.name);
console.log('ext: ' + parse.ext);
console.log();
// parse 를 사용하면 확장자 제거