/** (1) OS모듈 참조 */
import os from 'os';

/** (2) CPU 정보 */
const cpus = os.cpus();
console.log('CPU 코어 수 : ' + cpus.length);
console.log();

cpus.forEach((v, i) => {
	console.log('[%d번째 CPU] %s', i + 1, v.model);
	console.log('처리속도: %d', v.speed);
	console.log('수행시간 정보: %j', v.times);
	console.log();
});

