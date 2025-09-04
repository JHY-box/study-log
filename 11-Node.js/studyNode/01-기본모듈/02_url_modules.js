/** (1) url 모듈과 URL 클래스 참조하기 */
import url, { URL } from 'url';

/** (2) 주소 문자열을 URL 객체로 만들기 */
const current = new URL(
	'http://blog.hossam.kr:8765/hello/world.html?a=123&b=456#home'
);

console.log('href: ' + current.href);
console.log('protocol: ' + current.protocol);
console.log('host: ' + current.host);
console.log('hostname: ' + current.hostname);
console.log('port: ' + current.port);
console.log('pathname: ' + current.pathname);
console.log('search: ' + current.search);
console.log('hash: ' + current.hash);
console.log();

/** (3) 주소 문자열의 QueryString을 JSON 객체로 변환 */
const params = new URLSearchParams(current.search);
const json = Object.fromEntries(params);
console.log('JSON 객체 : %o', json);
console.log();

/** (4) JSON객체를 주소 문자열로 만들기 */
const info = {
	protocol: 'https:',
	hostname: 'blog.hossam.kr',
	port: '8080',
	pathname: '/hello/world.html',
	search: '?name=nodejs&age=10',
	hash: '#target',
};

const urlString = url.format(info);
console.log('주소 문자열 : %s', urlString);
