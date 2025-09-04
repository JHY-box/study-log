/** (1) OS모듈 참조 */
import os from "os";

/** (2) 네트워크 정보 */
const nets = os.networkInterfaces(); // 여기서 nets 는 JSON 객체

/** (3) 네트워크 정보 출력 */
for (const attr in nets) {
	console.group("Network장치 이름: %s", attr);
	const item = nets[attr];
	item.forEach((v, i) => {
		console.log("주소형식: %s", v.family);
		console.log("IP주소: %s", v.address);
		console.log("맥주소: %s", v.mac);
		console.log("넷마스크: %s", v.netmask);
		console.log();
	});

	console.groupEnd();
}
