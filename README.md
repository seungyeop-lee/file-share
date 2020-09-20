# File Share

## 필요조건

- docker 설치
- shell script 실행 가능 환경

## 기본 설정

- 공유폴더: `share_dir`
- basic auth
  - user: `test`
  - pass: `test`
- port
  - server: `65001`
  - debug: `12345`
- 설정은 `.env`을 통해 변경가능하다.

## 실행방법

### 서버가동

```bash
$ cd script
$ ./start.sh
```

### 접속

- localhost:8080/list
  - 공유 폴더의 list 출력
- localhost:8080/download?path=filepath
  - filepath에 있는 파일을 다운로드한다.

### 서버중지

```
$ cd script
$ ./stop.sh
```

### 