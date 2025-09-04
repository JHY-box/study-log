/** (1) OS모듈 참조 */
import os from 'os';

/** (2) 시스템 기본 정보 */
console.log('홈 디렉토리 : ' + os.homedir());
console.log('시스템 아키텍처 : ' + os.arch());
console.log('os플랫폼 : ' + os.platform());
console.log('시스템 임시 디렉토리 : ' + os.tmpdir());
console.log('시스템의 hostname : %s', os.hostname());
console.log();

/** (3) 사용자 계정정보 */
var userInfo = os.userInfo();
console.log('사용자 계정명 : ' + userInfo.username);
console.log('사용자 홈 디렉토리 : ' + userInfo.homedir);
console.log('사용자 쉘 환경(Linux, Mac) : ' + userInfo.shell);
console.log();

/** (4) 메모리 용량 */
// freemem() --> 시스템에서 현재 사용 가능한 메모리 용량
// totalmem() --> 시스템의 전체 메모리 용량
console.log(
	'시스템의 메모리 : %d(free) / %d(total)',
	os.freemem(),
	os.totalmem()
);

console.log();
