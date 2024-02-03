# macos 에서 실행
  - homebrew 설치
    - https://brew.sh/ko/
  - docker 설치
    - 터미널에서 `brew install --cask docker` 실행
    ```shell
    brew install --cask docker
    ```
  - docker 실행
    - launchpad 에서 docker app 실행
  - docker 실행 화면에서 로그인 혹은 로그인없이 시작
  - 터미널 실행 - 프로젝트 경로로 이동
  - image build 및 docker container 실행(`docker-compose up -d` 명령어 실행)
    ```shell
    docker-compose up -d
    ```
  - browser 실행 후 접속
    - http://localhost:8080


# windows 에서 실행
  - chocolatey 설치
    - https://chocolatey.org/install
  - docker 설치
    - powershell 관리자권한 실행 후 `choco install docker-desktop` 명령어 실행
    ```shell
    choco install docker-desktop
    ```
  - docker 실행(아래 두 가지 중 한 가지 방법으로 실행)
    - `시작` 메뉴에서 docker 검색 후 실행
    - `C:\Program Files\Docker\Docker\Docker Desktop.exe` 실행
  - docker 실행 화면에서 로그인 혹은 로그인없이 시작
  - 터미널 실행 - 프로젝트 경로로 이동
  - image build 및 docker container 실행(`docker-compose up -d` 명령어 실행)
    ```shell
    docker-compose up -d
    ```
  - browser 실행 후 접속
    - http://localhost:8080


# test case
- 유저
  - 회원가입
    - 올바른 휴대폰번호 유효성 검사
    - 이미 가입된 휴대폰번호
    - 비밀번호-비밀번호확인 불일치
  - 로그인
    - 가입되지 않은 유저
    - 비밀번호 불일치
- 상품 
  - 비로그인시 접근
  - 로그인 도중 쿠키삭제 후 접근
  - 상품 리스트 표시
  - 상품 추가
  - 상품 수정
  - 상품 삭제
  - 상품 상세보기
  - 상품 검색
    - 상품 일반 키워드 검색
    - 상품 초성 키워드 검색

# 기타
- 웹페이지 동작을 위해 우선 html 처리하여 요구사항에 맞는 json response 는 웹페이지 하단 별도 추가
- 테스트 데이터 불필요시 init.sql 하단에 있는 테스트데이터 쿼리 주석 후 docker image build
- `ph` 실행 소요시간 보다 `mysql` 이미지 실행에 시간이 더 길어 container 배포 직후엔 `connection refused` 오류가 발생하나 5초에 한 번씩 DB 연결 재시도하므로 몇 번 연결 실패 후 자동으로 연결됨 