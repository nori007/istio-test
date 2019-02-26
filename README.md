# istio-test

golang -> docker -> kubernetis -> istio 

#목적
 - pubilc ip가 아닌 private ip로 동작하도록
   (결국은 서버사이드 렌더링 하면 될듯. 서버사이드 렌더링, 클라이언트 렌더링 차이점 확인 필요)

#기본 구성
 - 웹페이지를 보여주는 모듈 하나: potal-server
 - 인증 기능 모듈: auth-server
 - API 기능 모듈: api-server


#대략 흐름..
 1. 사용자는 potal-server의 웹페이지에 접근.
 2. potal-server는 auth-server로 인증 요청
    2-1. potal-server의 웹단은 potal-server의 백단으로 요청
    2-2. potal-server 백단은 auth-server로 인증 요청
 3. auth-server에서 인증 처리후 potal-server로 결과 보내줌
    3-1. auth-server로 부터 결과 값을 받은 potal-server 백단이 받음
    3-2. potal-server 백단은 potal-server 웹단으로 결과 줌.
 4. potal-server는 api-server로 기능 요청
    4-1. potal-server의 웹단은 potal-server의 백단으로 요청
    4-2. potal-server 백단은 api-server로 인증 요청
 5. api-server에서 인증 처리후 potal-server로 결과 보내줌
    5-1. api-server로 부터 결과 값을 받은 potal-server 백단이 받음
    5-2. potal-server 백단은 potal-server 웹단으로 결과 줌.

