# istio-test

golang -> docker -> kubernetis -> istio 

목적
----------------
 - pubilc ip가 아닌 private ip로 동작하도록
   (결국은 서버사이드 렌더링 하면 될듯. 서버사이드 렌더링, 클라이언트 렌더링 차이점 확인 필요)

기본 구성
----------------
 - 웹페이지를 보여주는 모듈 하나: potal-server
 - 인증 기능 모듈: auth-server
 - API 기능 모듈: api-server


대략 흐름..
----------------
1. 사용자는 potal-server의 웹페이지에 접근.
2. potal-server는 auth-server로 인증 요청
3. potal-server의 웹단은 potal-server의 백단으로 요청
4. potal-server 백단은 auth-server로 인증 요청
5. auth-server에서 인증 처리후 potal-server로 결과 보내줌
6. auth-server로 부터 결과 값을 받은 potal-server 백단이 받음
7. potal-server 백단은 potal-server 웹단으로 결과 줌.
8. potal-server는 api-server로 기능 요청
9. potal-server의 웹단은 potal-server의 백단으로 요청
10. potal-server 백단은 api-server로 인증 요청
11. api-server에서 인증 처리후 potal-server로 결과 보내줌
12. api-server로 부터 결과 값을 받은 potal-server 백단이 받음
13. potal-server 백단은 potal-server 웹단으로 결과 줌.


부족한 부분
----------------
> 도커이미지가 계속 버전 업이 된 이유 (동일한 버전으로 배포를 못한 이유)
>>쿠베네티스에서 v1의 컨테이너를 만들고 파드를 만들면 갱신이 안됨. 쿠버네티스 상에서 도커 이미지를 지워야 갱신이 될것으로 보임.

> 현재 3개의 서비스 모두가 기본 index경로로 로직 처리함.
>>restapi처럼 사용하려면 경로에 대한 처리를 다르게 하는 로직이 필요

개인 적인 생각들 및 결론
----------------
 - ip가 아닌 yaml파일에 작성된 서비스의 이름으로 통신을 주고 받음.(public ip 아닌, 내부의 지정된 ip를 사용)
 - 소설을 쓰자면 yaml파일의 지정된 서비스 이름으로 DNS(?) 처리 했다고 생각됨.
