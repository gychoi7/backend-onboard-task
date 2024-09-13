# API 명세

## 1. 사용자 관리
### 1.1 사용자 등록
- Request
    - Method: POST
    - Path: /users/signup
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
        - Body:
            - password: 사용자 비밀번호
            - email: 사용자 이메일

- Response
- Status: 201
- message: 회원가입에 성공했습니다.

### 1.2 로그인
- Request
    - Method: POST
    - Path: /users/signin
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
        - Body:
            - password: 사용자 비밀번호
            - email: 사용자 이메일
- Response
    - Status: 200
    - data:
        - token: 사용자 인증 토큰

## 2. 게시글 관리
### 2.1 특정 사용자의 전체 게시글 조회
- Request
    - Method: GET
    - Path: /posts
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
        - Query:
            - page: 페이지 번호
            - limits: 페이지 당 게시글 수
- Response
    - Status: 200
    - data:
        - posts: 게시글 목록
        - totalCount: 전체 게시글 수
        - totalPages: 전체 페이지 수
        - page: 현재 페이지 번호

### 2.2 게시글 등록
- Request
    - Method: POST
    - Path: /posts/
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
        - Body:
            - title: 게시글 제목
            - content: 게시글 내용
- Response
- Status: 201
- message: 게시글이 성공적으로 생성되었습니다.

### 2.3 게시글 수정
- Request
    - Method: PUT
    - Path: /posts/:id
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
        - Body:
            - title: 게시글 제목
            - content: 게시글 내용
- Response
    - Status: 200
    - message: 게시글이 성공적으로 수정되었습니다.

### 2.4 게시글 삭제
- Request
    - Method: DELETE
    - Path: /posts/:id
    - Header:
        - Content-Type: application/json
        - Authorization: Bearer {token}
        - token: 사용자 인증 토큰
- Response
- Status: 200
- message: 게시글이 성공적으로 삭제되었습니다.
