# backend-onboard-task

## Backend 채용 전형

게시판을 관리하는 RESTful API를 개발해 주세요. 이때, 다음의 기능을 구현해야 합니다. 데이터베이스의 테이블 설계는 지원자분의 판단에 맡겨져 있습니다. 요구사항을 충족시키는 데 필요하다고 생각되는 구조로 자유롭게 설계해 주세요.

### 과제 소개

- 제약 조건 :
  - 사용 가능 언어 : Go
  - 사용 가능 라이브러리 : 제약 없음 (단 과제 조건의 기능이 이미 구현되어 바로 사용하는 형태 제외)

- [] 과제 1. 사용자 회원가입 엔드포인트
  - [] 이메일과 비밀번호로 회원가입할 수 있는 엔드포인트를 구현해 주세요.
  - [] 이메일과 비밀번호에 대한 유효성 검사를 구현해 주세요.
  - [] 이메일 조건: @ 포함
  - [] 비밀번호 조건: 8자 이상
  - [] 비밀번호는 반드시 암호화하여 저장해 주세요.
  - [] 이메일과 비밀번호의 유효성 검사는 위의 조건만으로 진행해 주세요. 추가적인 유효성 검사 조건은 포함하지 마세요.
- [] 과제 2. 사용자 로그인 엔드포인트
  - [] 사용자가 올바른 이메일과 비밀번호를 제공하면, 사용자 인증을 거친 후에 JWT(JSON Web Token)를 생성하여 사용자에게 반환하도록 해주세요.
  - [] 과제 1과 마찬가지로 회원가입 엔드포인트에 이메일과 비밀번호의 유효성 검사기능을 구현해주세요.
- [] 과제 3. 새로운 게시글을 생성하는 엔드포인트
- [] 과제 4. 게시글 목록을 조회하는 엔드포인트
  - [] (선택) Pagination 기능을 구현해 주세요.
- [] 과제 5. 특정 게시글을 조회하는 엔드포인트
  - [] 게시글의 ID를 받아 해당 게시글을 조회하는 엔드포인트를 구현해 주세요.
- [] 과제 6. 특정 게시글을 수정하는 엔드포인트
  - [] 게시글의 ID와 수정 내용을 받아 해당 게시글을 수정하는 엔드포인트를 구현해 주세요.
  - [] 게시글을 수정할 수 있는 사용자는 게시글 작성자만이어야 합니다.
- [] 과제 7. 특정 게시글을 삭제하는 엔드포인트
  - [] 게시글의 ID를 받아 해당 게시글을 삭제하는 엔드포인트를 구현해 주세요.
  - [] 게시글을 삭제할 수 있는 사용자는 게시글 작성자만이어야 합니다.
- [] 과제 8. 위 과제를 해결하며 나온 산출물에 대한 API 명세 문서 및 실행 가이드를 작성해 주세요.
  - [] API 명세 문서에는 요청 파라미터, 응답 예시가 나와 있어야 합니다.
  - [] 실행 가이드는 해당 가이드만 보고도 누구나 실행이 가능하도록 작성해 주세요.
  - [] API 명세 문서는 Repository Root의 API_SPEC.md 파일에 작성해 주세요.
  - [] 실행 가이드는 Repository Root의 GUIDE.md 파일에 작성해 주세요. 

### 실행 방법

- 본 레포지토리에 포함된 docker-compose 파일은 mysql을 로컬에서 실행 시킵니다.
  - Port: 3306
  - Root Password: 1234
  - User Password: 1234
  - Database Name: onycom

### 제출 방법

- 본인의 레포지토리에 이 Repository를 Fork 합니다.
- 과제의 결과물을 Fork된 Repository에 Commit 합니다.
- 과제 완료 후 해당 Repository 주소를 이메일로 회신 합니다. (ushiii111798@onycom.com)