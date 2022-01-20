
# Table of Contents



Restored session: Wed Jan 19 12:17:23 EET 2022
➜  cmd git:(#backend-verify-token-api) ✗ make backend/run
make: **\*** No rule to make target \`backend/run'.  Stop.
➜  cmd git:(#backend-verify-token-api) ✗ ls
cmd.go     cmd.go~    server.go  server.go~
➜  cmd git:(#backend-verify-token-api) ✗ cd ..
➜  backend git:(#backend-verify-token-api) ✗ make backend/run
make: **\*** No rule to make target \`backend/run'.  Stop.
➜  backend git:(#backend-verify-token-api) ✗ cd ..
➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem

rest/api/rest<sub>private.go</sub>:41:7: non-name token.resp on left side of :=
rest/api/rest<sub>private.go</sub>:48:21: undefined: resp
make: **\*** [backend/run] Error 2
➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem

ta/token.go:26:80: syntax error: mixed named and unnamed parameters
ta/token.go:27:4: no new variables on left side of :=
ta/token.go:36:11: no new variables on left side of :=
make: **\*** [backend/run] Error 2
➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem

ta/token.go:26:80: syntax error: mixed named and unnamed parameters
make: **\*** [backend/run] Error 2
➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 01:28:16 INFO: All 0 tables opened in 0s
badger 2022/01/20 01:28:16 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 01:28:16 INFO: Set nextTxnTs to 0
badger 2022/01/20 01:28:16 INFO: Deleting empty file: ../obada/trust-anchor/data/000055.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 01:28:16.794951 server.go:54: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:28:16.797843 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:28:24.742873 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:55616 - 201 155B in 786.793µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:34:33.343167 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55717 - 405 0B in 30.825µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:34:47.447146 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55718 - 405 0B in 27.93µs
^C2022/01/20 01:36:49 [WARN] interrupt signal
2022/01/20 01:36:49 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 01:36:49.916911 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 01:36:49.917017 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 01:36:49 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 01:36:49.918050 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 01:36:49 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 01:37:45 INFO: All 0 tables opened in 0s
badger 2022/01/20 01:37:45 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 01:37:45 INFO: Set nextTxnTs to 0
badger 2022/01/20 01:37:45 INFO: Deleting empty file: ../obada/trust-anchor/data/000056.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 01:37:45.835202 server.go:54: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:37:45.839727 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:37:50.717908 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55728 - 405 0B in 64.006µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:39:15.541382 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55730 - 405 0B in 31.432µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:39:23.196452 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55733 - 405 0B in 34.639µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:39:50.601210 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55734 - 405 0B in 43.6µs
^C2022/01/20 01:40:09 [WARN] interrupt signal
2022/01/20 01:40:09 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:09.419521 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:09.419678 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 01:40:09 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:09.420775 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 01:40:09 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 01:40:11 INFO: All 0 tables opened in 0s
badger 2022/01/20 01:40:11 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 01:40:11 INFO: Set nextTxnTs to 0
badger 2022/01/20 01:40:11 INFO: Deleting empty file: ../obada/trust-anchor/data/000057.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:11.113007 server.go:54: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:11.114747 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:40:17.366913 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55737 - 405 0B in 65.265µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:16.751716 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55740 - 405 0B in 30.154µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:20.088105 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55741 - 405 0B in 28.021µs
^C2022/01/20 01:41:43 [WARN] interrupt signal
2022/01/20 01:41:43 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:43.031905 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:43.032192 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 01:41:43 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:43.033226 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 01:41:43 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(#backend-verify-token-api) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anchor/taaut
h.pem
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 01:41:44 INFO: All 0 tables opened in 0s
badger 2022/01/20 01:41:44 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 01:41:44 INFO: Set nextTxnTs to 0
badger 2022/01/20 01:41:44 INFO: Deleting empty file: ../obada/trust-anchor/data/000058.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:44.730082 server.go:54: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:44.733163 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:46.866655 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55744 - 200 22B in 722.158µs
 :: TRUST-ANCHOR  :: 2022/01/20 01:41:53.494845 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:55745 - 200 49B in 150.509µs
^C2022/01/20 11:15:40 [WARN] interrupt signal
2022/01/20 11:15:40 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 11:15:40.328661 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 11:15:40.328845 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 11:15:40 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 11:15:40.330786 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 11:15:40 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(#backend-verify-token-api) ✗ git status
On branch #backend-verify-token-api
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pem
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(#backend-verify-token-api) ✗ git diff Makefile
➜  trust-anchor git:(#backend-verify-token-api) ✗ echo Makefile Makefile.back
Makefile Makefile.back
➜  trust-anchor git:(#backend-verify-token-api) ✗                              
➜  trust-anchor git:(#backend-verify-token-api) ✗ cat Makefile.back
cat: Makefile.back: No such file or directory
➜  trust-anchor git:(#backend-verify-token-api) ✗ cat Makefile > Makefile.back
➜  trust-anchor git:(#backend-verify-token-api) ✗ cat Makefile.back 
backend/compile:
        cd backend && go build -o trust-anchor

backend/run:
        cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.public-key-path=/Users/sr/projects/trust-anchor/taauth.pub.pem &#x2013;auth.private-key-path=/Users/sr/projects/trust-anch
or/taauth.pem

backend/vendor:
        cd backend && go mod tidy && go mod vendor

backend/lint:
        cd backend && golangci-lint &#x2013;config .golangci.yml run &#x2013;print-issued-lines &#x2013;out-format=github-actions ./&#x2026;

backend/docker/build:

backend/docker/publish:

frontend/build: frontend/compile frontend/minify frontend/checksum frontend/clean frontend/embed

frontend/compile:
        cd frontend && elm make src/Main.elm &#x2013;optimize &#x2013;output=./build/app.js

frontend/minify:
        cd frontend && uglifyjs ./build/app.js &#x2013;compress 'pure<sub>funcs</sub>=[F2,F3,F4,F5,F6,F7,F8,F9,A2,A3,A4,A5,A6,A7,A8,A9],pure<sub>getters,keep</sub><sub>fargs</sub>=false,unsafe<sub>comps,unsafe</sub>' | uglifyjs &#x2013;mangle &#x2013;output ./bu
ild/app.min.js

frontend/checksum:
        cd frontend && mv ./build/app.min.js ./build/app.$$(sha512sum ./build/app.min.js | head -c 10).min.js

frontend/clean:
        cd frontend && rm ./build/app.js

APP<sub>FILE</sub> = \((shell ls ./build | grep app)
frontend/embed:
        cd frontend && cp index.html ./build
        sed -i "s/app.min.js/\)(APP<sub>FILE</sub>)/" ./build/index.html

frontend/test:
        cd frontend && elm-test
➜  trust-anchor git:(#backend-verify-token-api) ✗ cat README.md README.md.bak

\## API endpoints
cat: README.md.bak: No such file or directory
➜  trust-anchor git:(#backend-verify-token-api) ✗ cat README.md > README.md.bak
➜  trust-anchor git:(#backend-verify-token-api) ✗ git status
On branch #backend-verify-token-api
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        Makefile.back
        README.md.bak
        taauth.pem
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(#backend-verify-token-api) ✗ git checkout Makefile 
Updated 1 path from the index
➜  trust-anchor git:(#backend-verify-token-api) ✗ git checkout README.md 
Updated 1 path from the index
➜  trust-anchor git:(#backend-verify-token-api) ✗ git pull &#x2013;all
Fetching origin
remote: Enumerating objects: 1164, done.
remote: Counting objects: 100% (1164/1164), done.
remote: Compressing objects: 100% (891/891), done.
remote: Total 1161 (delta 215), reused 1146 (delta 215), pack-reused 0
Receiving objects: 100% (1161/1161), 3.07 MiB | 3.08 MiB/s, done.
Resolving deltas: 100% (215/215), done.
From github.com:obada-foundation/trust-anchor

-   4bbb176&#x2026;79b2461 develop    -> origin/develop  (forced update)

There is no tracking information for the current branch.
Please specify which branch you want to merge with.
See git-pull(1) for details.

git pull <remote> <branch>

If you wish to set tracking information for this branch you can do so with:

git branch &#x2013;set-upstream-to=origin/<branch> #backend-verify-token-api

➜  trust-anchor git:(#backend-verify-token-api) ✗ git checkout develop 
Branch 'develop' set up to track remote branch 'develop' from 'origin'.
Switched to a new branch 'develop'
➜  trust-anchor git:(develop) ✗ git status
On branch develop
Your branch is up to date with 'origin/develop'.

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        Makefile.back
        README.md.bak
        taauth.pem
        taauth.pub.pem

nothing added to commit but untracked files present (use "git add" to track)
➜  trust-anchor git:(develop) ✗ git status
On branch develop
Your branch is up to date with 'origin/develop'.

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        Makefile.back
        README.md.bak
        taauth.pem
        taauth.pub.pem

nothing added to commit but untracked files present (use "git add" to track)
➜  trust-anchor git:(develop) ✗ ls -la
total 64
drwxr-xr-x  17 sr  staff   544 Jan 20 11:18 .
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 ..
drwxr-xr-x  14 sr  staff   448 Jan 20 11:18 .git
-rw-r&#x2013;r&#x2013;   1 sr  staff     6 Jan 19 12:49 .gitignore
-rw-r&#x2013;r&#x2013;   1 sr  staff  1102 Jan 18 22:55 LICENSE
-rw-r&#x2013;r&#x2013;   1 sr  staff  1167 Jan 20 11:17 Makefile
-rw-r&#x2013;r&#x2013;   1 sr  staff  1304 Jan 20 11:17 Makefile.back
-rw-r&#x2013;r&#x2013;   1 sr  staff    15 Jan 20 11:18 README.md
-rw-r&#x2013;r&#x2013;   1 sr  staff    33 Jan 20 11:17 README.md.bak
drwxr-xr-x  10 sr  staff   320 Jan 19 23:29 backend
drwxr-xr-x   2 sr  staff    64 Jan 19 15:15 deployment
drwxr-xr-x   3 sr  staff    96 Jan 18 22:55 docker
drwxr-xr-x   6 sr  staff   192 Jan 19 15:17 docs
drwxr-xr-x  11 sr  staff   352 Jan 19 15:33 frontend
drwx-&#x2013;&#x2014;   3 sr  staff    96 Jan 19 12:48 obada
-rw-r&#x2013;r&#x2013;   1 sr  admin   119 Jan 19 23:42 taauth.pem
-rw-r&#x2013;r&#x2013;   1 sr  staff   113 Jan 19 23:44 taauth.pub.pem
➜  trust-anchor git:(develop) ✗ git checkout -b readme-document
Switched to a new branch 'readme-document'
➜  trust-anchor git:(readme-document) ✗ ls -la
total 64
drwxr-xr-x  17 sr  staff   544 Jan 20 11:18 .
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 ..
drwxr-xr-x  14 sr  staff   448 Jan 20 11:19 .git
-rw-r&#x2013;r&#x2013;   1 sr  staff     6 Jan 19 12:49 .gitignore
-rw-r&#x2013;r&#x2013;   1 sr  staff  1102 Jan 18 22:55 LICENSE
-rw-r&#x2013;r&#x2013;   1 sr  staff  1167 Jan 20 11:17 Makefile
-rw-r&#x2013;r&#x2013;   1 sr  staff  1304 Jan 20 11:17 Makefile.back
-rw-r&#x2013;r&#x2013;   1 sr  staff    15 Jan 20 11:18 README.md
-rw-r&#x2013;r&#x2013;   1 sr  staff    33 Jan 20 11:17 README.md.bak
drwxr-xr-x  10 sr  staff   320 Jan 19 23:29 backend
drwxr-xr-x   2 sr  staff    64 Jan 19 15:15 deployment
drwxr-xr-x   3 sr  staff    96 Jan 18 22:55 docker
drwxr-xr-x   6 sr  staff   192 Jan 19 15:17 docs
drwxr-xr-x  11 sr  staff   352 Jan 19 15:33 frontend
drwx-&#x2013;&#x2014;   3 sr  staff    96 Jan 19 12:48 obada
-rw-r&#x2013;r&#x2013;   1 sr  admin   119 Jan 19 23:42 taauth.pem
-rw-r&#x2013;r&#x2013;   1 sr  staff   113 Jan 19 23:44 taauth.pub.pem
➜  trust-anchor git:(readme-document) ✗ cat README.md.bak > README.md 
➜  trust-anchor git:(readme-document) ✗ rm README.md.bak 
➜  trust-anchor git:(readme-document) ✗ cat README.md 

\## API endpoints
➜  trust-anchor git:(readme-document) ✗ cat Makefile.back > Makefile
➜  trust-anchor git:(readme-document) ✗ rm Makefile.back 
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pem
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(readme-document) ✗ openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta<sub>api.pem</sub>
Algorithm ED25519 not found
usage: genpkey [-algorithm alg] [cipher] [-genparam] [-out file]
    [-outform der | pem] [-paramfile file] [-pass arg]
    [-pkeyopt opt:value] [-text]

 -algorithm name    Public key algorithm to use (must precede -pkeyopt)
 -genparam          Generate a set of parameters instead of a private key
 -out file          Output file to write to (default stdout)
 -outform format    Output format (DER or PEM)
 -paramfile file    File to load public key algorithm parameters from
                    (must precede -pkeyopt)
 -pass arg          Output file password source
 -pkeyopt opt:value Set public key algorithm option to the given value
 -text              Print the private/public key in human readable form
➜  trust-anchor git:(readme-document) ✗ which opessl
opessl not found
➜  trust-anchor git:(readme-document) ✗ which openssl
/usr/bin/openssl
➜  trust-anchor git:(readme-document) ✗ cd /usr/local/opt/openssl
➜  openssl ls -la
total 1600
drwxr-xr-x  14 sr  admin     448 Jan 19 23:40 .
drwxr-xr-x   3 sr  admin      96 Jan 19 23:40 ..
drwxr-xr-x   3 sr  admin      96 Dec 14 18:16 .bottle
drwxr-xr-x   3 sr  admin      96 Dec 14 18:16 .brew
-rw-r&#x2013;r&#x2013;   1 sr  admin     990 Dec 14 18:16 AUTHORS.md
-rw-r&#x2013;r&#x2013;   1 sr  admin  718660 Dec 14 18:16 CHANGES.md
-rw-r&#x2013;r&#x2013;   1 sr  admin    1229 Jan 19 23:40 INSTALL<sub>RECEIPT.json</sub>
-rw-r&#x2013;r&#x2013;   1 sr  admin   10175 Dec 14 18:16 LICENSE.txt
-rw-r&#x2013;r&#x2013;   1 sr  admin   69186 Dec 14 18:16 NEWS.md
-rw-r&#x2013;r&#x2013;   1 sr  admin    6565 Dec 14 18:16 README.md
drwxr-xr-x   4 sr  admin     128 Jan 19 23:42 bin
drwxr-xr-x   3 sr  admin      96 Dec 14 18:16 include
drwxr-xr-x  11 sr  admin     352 Dec 14 18:16 lib
drwxr-xr-x   4 sr  admin     128 Dec 14 18:16 share
➜  openssl cd bin 
➜  bin ls -la
total 1608
drwxr-xr-x   4 sr  admin     128 Jan 19 23:42 .
drwxr-xr-x  14 sr  admin     448 Jan 19 23:40 ..
-r-xr-xr-x   1 sr  admin    6602 Jan 19 23:40 c<sub>rehash</sub>
-r-xr-xr-x   1 sr  admin  814848 Jan 19 23:40 openssl
➜  bin ./openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta<sub>api.pem</sub>
&#x2013;&#x2014;BEGIN PRIVATE KEY&#x2013;&#x2014;
MC4CAQAwBQYDK2VwBCIEIJsG0ZK2hedTkUY0UhIJfIJKSHSaUNLloV1iMhPwCUMV
&#x2013;&#x2014;END PRIVATE KEY&#x2013;&#x2014;
➜  bin cat ~/.ssh/ta<sub>api.pem</sub> 
&#x2013;&#x2014;BEGIN PRIVATE KEY&#x2013;&#x2014;
MC4CAQAwBQYDK2VwBCIEIJsG0ZK2hedTkUY0UhIJfIJKSHSaUNLloV1iMhPwCUMV
&#x2013;&#x2014;END PRIVATE KEY&#x2013;&#x2014;
➜  bin openssl pkey -in ~/.ssh/ta<sub>api</sub>  -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
Error opening key *Users/sr*.ssh/ta<sub>api</sub>
4518893228:error:02FFF002:system library:func(4095):No such file or directory:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2
.8/crypto/bio/bss<sub>file.c</sub>:255:fopen('*Users/sr*.ssh/ta<sub>api</sub>', 'r')
4518893228:error:20FFF002:BIO routines:CRYPTO<sub>internal</sub>:system lib:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2.8/crypto/bi
o/bss<sub>file.c</sub>:257:
unable to load key
➜  bin openssl pkey -in ~/.ssh/ta<sub>api.pem</sub>  -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
unable to load key
4425889452:error:06FFF09C:digital envelope routines:CRYPTO<sub>internal</sub>:unsupported algorithm:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.
3/libressl-2.8/crypto/evp/p<sub>lib.c</sub>:245:
4425889452:error:06FFF076:digital envelope routines:CRYPTO<sub>internal</sub>:unsupported private key algorithm:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/lib
ressl-75.60.3/libressl-2.8/crypto/evp/evp<sub>pkey.c</sub>:85:TYPE=Ed25519
4425889452:error:09FFF00D:PEM routines:CRYPTO<sub>internal</sub>:ASN1 lib:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2.8/crypto/pem/
pem<sub>pkey.c</sub>:143:
➜  bin openssl pkey -in ~/.ssh/ta<sub>api</sub>  -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub> 
Error opening key *Users/sr*.ssh/ta<sub>api</sub>
4464273068:error:02FFF002:system library:func(4095):No such file or directory:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2
.8/crypto/bio/bss<sub>file.c</sub>:255:fopen('*Users/sr*.ssh/ta<sub>api</sub>', 'r')
4464273068:error:20FFF002:BIO routines:CRYPTO<sub>internal</sub>:system lib:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2.8/crypto/bi
o/bss<sub>file.c</sub>:257:
unable to load key
➜  bin openssl pkey -in ~/.ssh/ta<sub>api</sub> -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub> 
Error opening key *Users/sr*.ssh/ta<sub>api</sub>
4584941228:error:02FFF002:system library:func(4095):No such file or directory:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2
.8/crypto/bio/bss<sub>file.c</sub>:255:fopen('*Users/sr*.ssh/ta<sub>api</sub>', 'r')
4584941228:error:20FFF002:BIO routines:CRYPTO<sub>internal</sub>:system lib:/System/Volumes/Data/SWE/macOS/BuildRoots/5b2e67f8af/Library/Caches/com.apple.xbs/Sources/libressl/libressl-75.60.3/libressl-2.8/crypto/bi
o/bss<sub>file.c</sub>:257:
unable to load key
➜  bin ./openssl pkey -in ~/.ssh/ta<sub>api</sub> -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
Could not open file or uri for loading key from *Users/sr*.ssh/ta<sub>api</sub>
00A66A0D01000000:error:16000069:STORE routines:ossl<sub>store</sub><sub>get0</sub><sub>loader</sub><sub>int</sub>:unregistered scheme:crypto/store/store<sub>register.c</sub>:237:scheme=file
00A66A0D01000000:error:80000002:system library:file<sub>open</sub>:No such file or directory:providers/implementations/storemgmt/file<sub>store.c</sub>:269:calling stat(*Users/sr*.ssh/ta<sub>api</sub>)
➜  bin ./openssl pkey -in ~/.ssh/ta<sub>api.pem</sub> -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
&#x2013;&#x2014;BEGIN PUBLIC KEY&#x2013;&#x2014;
MCowBQYDK2VwAyEAqX58enV/v4arHp8e0Rc94LVZAv0Esebeoi+CfycxlwU=
&#x2013;&#x2014;END PUBLIC KEY&#x2013;&#x2014;
➜  bin ./openssl pkey -in ~/.ssh/ta<sub>api.pem</sub> -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
&#x2013;&#x2014;BEGIN PUBLIC KEY&#x2013;&#x2014;
MCowBQYDK2VwAyEAqX58enV/v4arHp8e0Rc94LVZAv0Esebeoi+CfycxlwU=
&#x2013;&#x2014;END PUBLIC KEY&#x2013;&#x2014;
➜  bin ./openssl pkey -in ~/.ssh/ta<sub>api.pem</sub> -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
&#x2013;&#x2014;BEGIN PUBLIC KEY&#x2013;&#x2014;
MCowBQYDK2VwAyEAqX58enV/v4arHp8e0Rc94LVZAv0Esebeoi+CfycxlwU=
&#x2013;&#x2014;END PUBLIC KEY&#x2013;&#x2014;
➜  bin cd ~/projects 
➜  projects ls -la
total 0
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 .
drwxr-x&#x2014;+ 51 sr  staff  1632 Jan 20 11:51 ..
drwxr-xr-x  37 sr  staff  1184 Dec 14 15:22 docker-solr
drwxr-xr-x   9 sr  staff   288 Dec 13 23:36 dot
drwxr-xr-x  16 sr  staff   512 Jan  9 23:17 go-url-fuzz
drwxr-xr-x  16 sr  staff   512 Jan  7 11:05 gr-osmosdr
drwxr-xr-x   4 sr  staff   128 Dec 21 14:38 obada
drwxr-xr-x   4 sr  staff   128 Dec 15 13:38 org
drwxr-xr-x  35 sr  staff  1120 Jan 19 23:40 payments-sample-app
drwxr-xr-x  17 sr  staff   544 Jan  9 23:20 scout
drwxr-xr-x   4 sr  staff   128 Dec 18 19:15 securityrobot
drwxr-xr-x   6 sr  staff   192 Dec 14 15:22 tradeloop
drwxr-xr-x  17 sr  staff   544 Jan 20 11:50 trust-anchor
➜  projects cd trust-anchor 
➜  trust-anchor git:(readme-document) ✗ ls -la
total 56
-rw-r&#x2013;r&#x2013;   1 sr  staff   724 Jan 20 11:50 #README.md#
drwxr-xr-x  17 sr  staff   544 Jan 20 11:50 .
lrwxr-xr-x   1 sr  staff    41 Jan 20 11:49 .#README.md -> sr@SecurityRobots-MacBook-Pro.local.72906
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 ..
drwxr-xr-x  14 sr  staff   448 Jan 20 11:50 .git
-rw-r&#x2013;r&#x2013;   1 sr  staff     6 Jan 19 12:49 .gitignore
-rw-r&#x2013;r&#x2013;   1 sr  staff  1102 Jan 18 22:55 LICENSE
-rw-r&#x2013;r&#x2013;   1 sr  staff  1304 Jan 20 11:20 Makefile
-rw-r&#x2013;r&#x2013;   1 sr  staff   724 Jan 20 11:48 README.md
drwxr-xr-x  10 sr  staff   320 Jan 19 23:29 backend
drwxr-xr-x   2 sr  staff    64 Jan 19 15:15 deployment
drwxr-xr-x   3 sr  staff    96 Jan 18 22:55 docker
drwxr-xr-x   6 sr  staff   192 Jan 19 15:17 docs
drwxr-xr-x  11 sr  staff   352 Jan 19 15:33 frontend
drwx-&#x2013;&#x2014;   3 sr  staff    96 Jan 19 12:48 obada
-rw-r&#x2013;r&#x2013;   1 sr  admin   119 Jan 19 23:42 taauth.pem
-rw-r&#x2013;r&#x2013;   1 sr  staff   113 Jan 19 23:44 taauth.pub.pem
➜  trust-anchor git:(readme-document) ✗ make backend/compile 
cd backend && go build -o trust-anchor
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        backend/trust-anchor
        taauth.pem
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(readme-document) ✗ ./backend/trust-anchor -h
TrustAnchor unknown
(c)OBADA 2022

Usage:
  trust-anchor [OPTIONS] <server>

Application Options:
      &#x2013;db-path= Path to database files (default: /var/obada/trust-anchor/data)

Help Options:
  -h, &#x2013;help     Show this help message

Available commands:
  server

➜  trust-anchor git:(readme-document) ✗ ./backend/trust-anchor server -h
TrustAnchor unknown
(c)OBADA 2022

Usage:
  trust-anchor [OPTIONS] server [server-OPTIONS]

Application Options:
      &#x2013;db-path=                        Path to database files (default: /var/obada/trust-anchor/data)

Help Options:
  -h, &#x2013;help                            Show this help message

[server command options]
          &#x2013;port=                       port (default: 80) [$TA<sub>PORT</sub>]
          &#x2013;address=                    listening address [$TA<sub>ADDRESS</sub>]

ssl:
      &#x2013;ssl.type=[acme|none|static] ssl support (default: none) [$SSL<sub>TYPE</sub>]
      &#x2013;ssl.cert=                   path to cert.pem file [$SSL<sub>CERT</sub>]
      &#x2013;ssl.key=                    path to key.pem file [$SSL<sub>KEY</sub>]

auth:
      &#x2013;auth.public-key-path=
      &#x2013;auth.private-key-path=

➜  trust-anchor git:(readme-document) ✗ rm backend/trust-anchor 
➜  trust-anchor git:(readme-document) ✗ cd ..
➜  projects git status
fatal: not a git repository (or any of the parent directories): .git
➜  projects cd trust-anchor 
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pem
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(readme-document) ✗ rm taauth.pub.pem 
➜  trust-anchor git:(readme-document) ✗ rm taauth.pem 
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(readme-document) ✗ git add Makefile 
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile

Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   README.md

➜  trust-anchor git:(readme-document) ✗ git add README.md 
➜  trust-anchor git:(readme-document) ✗ ls -la
total 32
drwxr-xr-x  13 sr  staff   416 Jan 20 11:57 .
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 ..
drwxr-xr-x  14 sr  staff   448 Jan 20 12:01 .git
-rw-r&#x2013;r&#x2013;   1 sr  staff     6 Jan 19 12:49 .gitignore
-rw-r&#x2013;r&#x2013;   1 sr  staff  1102 Jan 18 22:55 LICENSE
-rw-r&#x2013;r&#x2013;   1 sr  staff  1254 Jan 20 11:56 Makefile
-rw-r&#x2013;r&#x2013;   1 sr  staff   834 Jan 20 11:56 README.md
drwxr-xr-x  10 sr  staff   320 Jan 20 11:57 backend
drwxr-xr-x   2 sr  staff    64 Jan 19 15:15 deployment
drwxr-xr-x   3 sr  staff    96 Jan 18 22:55 docker
drwxr-xr-x   6 sr  staff   192 Jan 20 12:01 docs
drwxr-xr-x  11 sr  staff   352 Jan 19 15:33 frontend
drwx-&#x2013;&#x2014;   3 sr  staff    96 Jan 19 12:48 obada
➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   README.md

Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   docs/Spec.org

➜  trust-anchor git:(readme-document) ✗ git add docs/
➜  trust-anchor git:(readme-document) ✗ git status   
On branch readme-document
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   README.md
        modified:   docs/Spec.org

➜  trust-anchor git:(readme-document) ✗ git commit -m 'Readme and spec update'
[readme-document c9fa76a] Readme and spec update
 3 files changed, 32 insertions(+), 16 deletions(-)
➜  trust-anchor git:(readme-document) git checkout develop
Switched to branch 'develop'
Your branch is up to date with 'origin/develop'.
➜  trust-anchor git:(develop) git checkout -b blockchain-token-design
Switched to a new branch 'blockchain-token-design'
➜  trust-anchor git:(blockchain-token-design) git checkout -b blockchain-jwt-token-claims
Switched to a new branch 'blockchain-jwt-token-claims'
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ ls -la
total 32
drwxr-xr-x  13 sr  staff   416 Jan 20 12:11 .
drwxr-xr-x  13 sr  staff   416 Jan 18 22:55 ..
drwxr-xr-x  14 sr  staff   448 Jan 20 12:16 .git
-rw-r&#x2013;r&#x2013;   1 sr  staff     6 Jan 19 12:49 .gitignore
-rw-r&#x2013;r&#x2013;   1 sr  staff  1102 Jan 18 22:55 LICENSE
-rw-r&#x2013;r&#x2013;   1 sr  staff  1167 Jan 20 12:11 Makefile
-rw-r&#x2013;r&#x2013;   1 sr  staff    15 Jan 20 12:11 README.md
drwxr-xr-x  10 sr  staff   320 Jan 20 11:57 backend
drwxr-xr-x   2 sr  staff    64 Jan 19 15:15 deployment
drwxr-xr-x   3 sr  staff    96 Jan 18 22:55 docker
drwxr-xr-x   6 sr  staff   192 Jan 20 12:11 docs
drwxr-xr-x  11 sr  staff   352 Jan 19 15:33 frontend
drwx-&#x2013;&#x2014;   3 sr  staff    96 Jan 19 12:48 obada
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run 
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data

cmd/server.go:89:22: not enough arguments in call to ta.New
        have (ed25519.PrivateKey, ed25519.PublicKey)
        want (string, interface {}, interface {})
make: **\*** [backend/run] Error 2
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data
TrustAnchor unknown
(c)OBADA 2022

the required flags \`&#x2013;auth.private-key-path' and \`&#x2013;auth.public-key-path' were not specified
 :: TRUST-ANCHOR  :: 2022/01/20 12:31:26.267439 main.go:31: the required flags \`&#x2013;auth.private-key-path' and \`&#x2013;auth.public-key-path' were not specified
exit status 1
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=~/.ssh/ta<sub>token.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

the required flag \`&#x2013;auth.public-key-path' was not specified
 :: TRUST-ANCHOR  :: 2022/01/20 12:34:38.132729 main.go:31: the required flag \`&#x2013;auth.public-key-path' was not specified
exit status 1
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=~/.ssh/ta<sub>token.pem</sub> &#x2013;auth.public-key-path==~/.ssh/ta<sub>token.pem.pub</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:35:15 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:35:15 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:35:15 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:35:15 INFO: Deleting empty file: ../obada/trust-anchor/data/000059.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:35:15.995744 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 12:35:15 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 12:35:15 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open ~/.ssh/ta<sub>token.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16032c0, 0xc000587860})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffabd, 0xc000126008}, {0x7ff7bfeffae8, 0x18})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:63 +0x5d
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0001367a0)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0xab
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0001367a0, {0x168cc5a, 0x7ff7bfeffa8a, 0xc00012ebe0})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1fdedd8, 0xc0001367a0}, {0xc00014b640, 0x0, 0x4})
        /Users/sr/projects/trust-anchor/backend/main.go:54 +0x1a9
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc000242e70, {0xc00012e010, 0x16, 0x4})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc00012ebe0, 0xc000136790)
        /Users/sr/projects/trust-anchor/backend/main.go:61 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:27 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗                      
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run                           
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=~/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path==~/.ssh/ta<sub>api.pem.pub</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:35:56 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:35:56 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:35:56 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:35:56 INFO: Deleting empty file: ../obada/trust-anchor/data/000060.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:35:56.684782 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 12:35:56 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 12:35:56 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open ~/.ssh/ta<sub>api.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16032c0, 0xc0003282d0})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffabd, 0xc000010018}, {0x7ff7bfeffae6, 0x16})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:63 +0x5d
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0001287a0)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0xab
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0001287a0, {0x168cc5a, 0x7ff7bfeffa8a, 0xc000020c80})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1f5ed38, 0xc0001287a0}, {0xc000033680, 0x0, 0x4})
        /Users/sr/projects/trust-anchor/backend/main.go:54 +0x1a9
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc000216e70, {0xc0000200b0, 0x16, 0x4})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc000020c80, 0xc000128790)
        /Users/sr/projects/trust-anchor/backend/main.go:61 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:27 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ cat ~/.ssh/ta<sub>api.p</sub>
cat: *Users/sr*.ssh/ta<sub>api.p</sub>: No such file or directory
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run                           
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=~/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path==~/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:36:29 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:36:29 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:36:29 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:36:29 INFO: Deleting empty file: ../obada/trust-anchor/data/000061.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:36:29.145220 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 12:36:29 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 12:36:29 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open ~/.ssh/ta<sub>api.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16032c0, 0xc0005898c0})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffabe, 0xc0000b6008}, {0x7ff7bfeffae7, 0x16})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:63 +0x5d
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0000c27a0)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0xab
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0000c27a0, {0x168cc5a, 0x7ff7bfeffa8b, 0xc000094c30})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1fa3058, 0xc0000c27a0}, {0xc000091680, 0x0, 0x4})
        /Users/sr/projects/trust-anchor/backend/main.go:54 +0x1a9
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc000244e70, {0xc000094060, 0x16, 0x4})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc000094c30, 0xc0000c2790)
        /Users/sr/projects/trust-anchor/backend/main.go:61 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:27 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ cat ~/.ssh/ta<sub>api.pem</sub>
&#x2013;&#x2014;BEGIN PRIVATE KEY&#x2013;&#x2014;
MC4CAQAwBQYDK2VwBCIEIJsG0ZK2hedTkUY0UhIJfIJKSHSaUNLloV1iMhPwCUMV
&#x2013;&#x2014;END PRIVATE KEY&#x2013;&#x2014;
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run     
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path==~/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:37:39 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:37:39 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:37:39 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:37:39 INFO: Deleting empty file: ../obada/trust-anchor/data/000062.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:37:39.076268 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 12:37:39 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 12:37:39 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open =~/.ssh/ta<sub>api.pub.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16032c0, 0xc000547860})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffabe, 0xc000126008}, {0x7ff7bfeffaef, 0x16})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:68 +0x98
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0001367a0)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0xab
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0001367a0, {0x168cc5a, 0x7ff7bfeffa8b, 0xc00012ebe0})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1fdedd8, 0xc0001367a0}, {0xc00014b640, 0x0, 0x4})
        /Users/sr/projects/trust-anchor/backend/main.go:54 +0x1a9
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc000240e70, {0xc00012e010, 0x16, 0x4})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc00012ebe0, 0xc000136790)
        /Users/sr/projects/trust-anchor/backend/main.go:61 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:27 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path==/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:37:53 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:37:53 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:37:53 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:37:53 INFO: Deleting empty file: ../obada/trust-anchor/data/000063.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:37:53.945640 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 12:37:53 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 12:37:53 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open =/Users/sr/.ssh/ta<sub>api.pub.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16032c0, 0xc0005879b0})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffab6, 0xc000126008}, {0x7ff7bfeffae7, 0x1e})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:68 +0x98
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0001367a0)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0xab
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0001367a0, {0x168cc5a, 0x7ff7bfeffa83, 0xc00012ebe0})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1fdeff8, 0xc0001367a0}, {0xc00014b640, 0x0, 0x4})
        /Users/sr/projects/trust-anchor/backend/main.go:54 +0x1a9
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc000240e70, {0xc00012e010, 0x16, 0x4})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc00012ebe0, 0xc000136790)
        /Users/sr/projects/trust-anchor/backend/main.go:61 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:27 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ cat ~/.ssh/ta<sub>api.pub.pem</sub> 
&#x2013;&#x2014;BEGIN PUBLIC KEY&#x2013;&#x2014;
MCowBQYDK2VwAyEAqX58enV/v4arHp8e0Rc94LVZAv0Esebeoi+CfycxlwU=
&#x2013;&#x2014;END PUBLIC KEY&#x2013;&#x2014;
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ cat ~/.ssh/ta<sub>api.pub.pem</sub>
&#x2013;&#x2014;BEGIN PUBLIC KEY&#x2013;&#x2014;
MCowBQYDK2VwAyEAqX58enV/v4arHp8e0Rc94LVZAv0Esebeoi+CfycxlwU=
&#x2013;&#x2014;END PUBLIC KEY&#x2013;&#x2014;
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run         
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:38:21 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:38:21 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:38:21 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:38:21 INFO: Deleting empty file: ../obada/trust-anchor/data/000064.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:38:21.660401 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 12:38:21.663638 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 12:39:16.666657 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:52938 - 401 22B in 604.092µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:39:19.413967 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:52940 - 401 22B in 191.659µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:39:20.732058 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:52941 - 401 22B in 211.535µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:40:56.662781 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:52946 - 200 49B in 185.224µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:41:40.428279 logger.go:161: "POST <http://localhost/api/v1/token> HTTP/1.1" from [::1]:52948 - 405 0B in 30.395µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:41:45.180286 logger.go:161: "POST <http://localhost/api/v1/token> HTTP/1.1" from [::1]:52949 - 405 0B in 28.224µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:41:52.666436 logger.go:161: "GET <http://localhost/api/v1/token> HTTP/1.1" from [::1]:52950 - 200 204B in 231.369µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:41:54.739810 logger.go:161: "GET <http://localhost/api/v1/token> HTTP/1.1" from [::1]:52951 - 200 204B in 155.881µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:48:51.165984 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53017 - 401 22B in 199.669µs
2022/01/20 12:48:52 [WARN] can't decode request data - invalid character 'd' looking for beginning of object key string - 400 (2) - /api/v1/nft/register - [rest/api/rest<sub>private.go</sub>:30 api.(\*private).regis
terNFT]
 :: TRUST-ANCHOR  :: 2022/01/20 12:48:52.884869 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53018 - 400 124B in 293.3µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:49:19.636634 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53019 - 201 211B in 339.243µs
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:34.713830 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53025 - 201 211B in 366.962µs
^C2022/01/20 12:51:54 [WARN] interrupt signal
2022/01/20 12:51:54 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:54.017779 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:54.017943 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 12:51:54 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:54.019146 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 12:51:54 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 12:51:56 INFO: All 0 tables opened in 0s
badger 2022/01/20 12:51:56 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 12:51:56 INFO: Set nextTxnTs to 0
badger 2022/01/20 12:51:56 INFO: Deleting empty file: ../obada/trust-anchor/data/000065.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:56.828174 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:56.832110 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 12:51:59.398359 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53028 - 201 232B in 857.888µs
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:06.983779 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53396 - 201 232B in 366.127µs
^C2022/01/20 13:33:33 [WARN] interrupt signal
2022/01/20 13:33:33 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:33.498367 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:33.498502 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 13:33:33 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:33.498533 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 13:33:33 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 13:33:35 INFO: All 0 tables opened in 0s
badger 2022/01/20 13:33:35 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 13:33:35 INFO: Set nextTxnTs to 0
badger 2022/01/20 13:33:35 INFO: Deleting empty file: ../obada/trust-anchor/data/000066.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:35.856640 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:35.862539 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:33:39.303632 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53400 - 201 295B in 795.787µs
^C2022/01/20 13:40:27 [WARN] interrupt signal
2022/01/20 13:40:27 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:27.587438 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:27.587544 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 13:40:27 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:27.588694 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 13:40:27 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 13:40:30 INFO: All 0 tables opened in 0s
badger 2022/01/20 13:40:30 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 13:40:30 INFO: Set nextTxnTs to 0
badger 2022/01/20 13:40:30 INFO: Deleting empty file: ../obada/trust-anchor/data/000067.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:30.307610 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:30.312255 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:40:31.632071 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53460 - 201 292B in 982.072µs
^C2022/01/20 13:41:58 [WARN] interrupt signal
2022/01/20 13:41:58 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 13:41:58.296410 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 13:41:58.296601 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 13:41:58 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 13:41:58.297790 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 13:41:58 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 13:42:00 INFO: All 0 tables opened in 0s
badger 2022/01/20 13:42:00 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 13:42:00 INFO: Set nextTxnTs to 0
badger 2022/01/20 13:42:00 INFO: Deleting empty file: ../obada/trust-anchor/data/000068.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 13:42:00.815712 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:42:00.818515 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:42:03.901324 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53466 - 201 300B in 924.012µs
^C2022/01/20 13:54:40 [WARN] interrupt signal
2022/01/20 13:54:40 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 13:54:40.071025 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 13:54:40.071126 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 13:54:40 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 13:54:40.072237 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 13:54:40 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server &#x2013;db-path=../obada/trust-anchor/data &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

the required flag \`&#x2013;url' was not specified
 :: TRUST-ANCHOR  :: 2022/01/20 13:54:42.882830 main.go:32: the required flag \`&#x2013;url' was not specified
exit status 1
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 13:55:52 INFO: All 0 tables opened in 0s
badger 2022/01/20 13:55:52 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 13:55:52 INFO: Set nextTxnTs to 0
badger 2022/01/20 13:55:52 INFO: Deleting empty file: ../obada/trust-anchor/data/000069.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 13:55:52.182456 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:55:52.186142 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 13:56:45.789353 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53540 - 201 333B in 826.488µs
^C2022/01/20 14:00:41 [WARN] interrupt signal
2022/01/20 14:00:41 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:00:41.908491 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:00:41.908641 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:00:41 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:00:41.909923 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:00:41 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>

ta/token.go:28:45: undefined: taVerifyUrl
make: **\*** [backend/run] Error 2
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:01:56 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:01:56 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:01:56 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:01:56 INFO: Deleting empty file: ../obada/trust-anchor/data/000070.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:01:56.090311 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:01:56.093545 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:01:59.973989 logger.go:161: "POST <http://localhost/api/v1/nft/register> HTTP/1.1" from [::1]:53583 - 201 352B in 935.938µs
 :: TRUST-ANCHOR  :: 2022/01/20 14:07:28.702266 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:53604 - 200 22B in 243.167µs
^C2022/01/20 14:08:19 [WARN] interrupt signal
2022/01/20 14:08:19 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:08:19.304456 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:08:19.304576 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:08:19 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:08:19.305762 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:08:19 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ rm  taauth.pub.pem
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status        
On branch blockchain-jwt-token-claims
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run                                                                                       
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=\((HOME)/.ssh/ta_api.pem \
        --auth.public-key-path=\)(HOME)/.ssh/ta<sub>api.pub.pem</sub>
/bin/sh: HOME: command not found
/bin/sh: HOME: command not found
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:09:52 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:09:52 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:09:52 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:09:52 INFO: Deleting empty file: ../obada/trust-anchor/data/000071.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:09:52.417619 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 14:09:52 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 14:09:52 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open /.ssh/ta<sub>api.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16037c0, 0xc00019e270})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffac5, 0x1b6b840}, {0x7ff7bfeffaed, 0x14})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:63 +0x5d
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc0002ac020)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0x95
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc0002ac020, {0x7ff7bfeffa77, 0x10, 0x168d19a})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1f63298, 0xc0002ac020}, {0xc000118c60, 0x0, 0x6})
        /Users/sr/projects/trust-anchor/backend/main.go:56 +0x245
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc00028af50, {0xc000020080, 0x16, 0x6})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc00007ebe0, 0xc0002ac000)
        /Users/sr/projects/trust-anchor/backend/main.go:63 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:28 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=$HOME/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=$HOME/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:10:05 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:10:05 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:10:05 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:10:05 INFO: Deleting empty file: ../obada/trust-anchor/data/000072.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:05.645896 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:05.649196 rest.go:47: [INFO] activate http rest server on :80
^C2022/01/20 14:10:23 [WARN] interrupt signal
2022/01/20 14:10:23 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:23.378895 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:23.379174 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:10:23 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:23.380354 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:10:23 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=$HOME/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path="$HOME"/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:10:25 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:10:25 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:10:25 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:10:25 INFO: Deleting empty file: ../obada/trust-anchor/data/000073.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:25.270095 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:25.275527 rest.go:47: [INFO] activate http rest server on :80
^C2022/01/20 14:10:42 [WARN] interrupt signal
2022/01/20 14:10:42 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:42.585943 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:42.586225 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:10:42 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:42.587429 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:10:42 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=$HOME/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path="OME"/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:10:44 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:10:44 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:10:44 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:10:44 INFO: Deleting empty file: ../obada/trust-anchor/data/000074.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:10:44.400432 server.go:55: [INFO] starting server on port :80
badger 2022/01/20 14:10:44 INFO: Lifetime L0 stalled for: 0s
badger 2022/01/20 14:10:44 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
panic: open OME/.ssh/ta<sub>api.pub.pem</sub>: no such file or directory

goroutine 1 [running]:
github.com/kataras/jwt.glob..func1({0x16037c0, 0xc00010f8f0})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/hmac.go:55 +0x25
github.com/kataras/jwt.MustLoadEdDSA({0x7ff7bfeffab5, 0x1b6b840}, {0x7ff7bfeffae6, 0x17})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/kataras/jwt/eddsa.go:68 +0x98
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).newServerApp(0xc000111d60)
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:82 +0x95
github.com/obada-foundation/trust-anchor/cmd.(\*ServerCommand).Execute(0xc000111d60, {0x7ff7bfeffa67, 0x10, 0x168d19a})
        /Users/sr/projects/trust-anchor/backend/cmd/server.go:67 +0x128
main.run.func1({0x1f62df8, 0xc000111d60}, {0xc000100cc0, 0x0, 0x6})
        /Users/sr/projects/trust-anchor/backend/main.go:56 +0x245
github.com/jessevdk/go-flags.(\*Parser).ParseArgs(0xc00028af50, {0xc000020080, 0x16, 0x6})
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:333 +0x928
github.com/jessevdk/go-flags.(\*Parser).Parse(&#x2026;)
        /Users/sr/projects/trust-anchor/backend/vendor/github.com/jessevdk/go-flags/parser.go:191
main.run(0xc000102be0, 0xc000111d40)
        /Users/sr/projects/trust-anchor/backend/main.go:63 +0xea
main.main()
        /Users/sr/projects/trust-anchor/backend/main.go:28 +0xf2
exit status 2
make: **\*** [backend/run] Error 1
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ make backend/run
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:11:37 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:11:37 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:11:37 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:11:37 INFO: Deleting empty file: ../obada/trust-anchor/data/000075.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:11:37.420886 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:11:37.423691 rest.go:47: [INFO] activate http rest server on :80
^C2022/01/20 14:11:43 [WARN] interrupt signal
2022/01/20 14:11:43 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:11:43.071245 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:11:43.071392 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:11:43 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:11:43.072524 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:11:43 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git add backend/ Makefile 
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status               
On branch blockchain-jwt-token-claims
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git commit -m 'Rename blockchain token fields'
[blockchain-jwt-token-claims 1c2e554] Rename blockchain token fields
 6 files changed, 55 insertions(+), 31 deletions(-)
➜  trust-anchor git:(blockchain-jwt-token-claims) >&#x2026;.                                                                                                                                                     

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git add backend/ Makefile
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git commit -m 'Rename blockchain token fields'
[blockchain-jwt-token-claims 1c2e554] Rename blockchain token fields
 6 files changed, 55 insertions(+), 31 deletions(-)
➜  trust-anchor git:(blockchain-jwt-token-claimsblockchain-jwt-token-claims
zsh: parse error near \`OBADA'
➜  trust-anchor git:(blockchain-jwt-token-claims) git push origin badger 2022/01/20 14:08:19 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ rm  taauth.pub.pem
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status                                                                                                                                        <&#x2026;.
zsh: parse error near \`OBADA'
➜  trust-anchor git:(blockchain-jwt-token-claims) git push origin badger 2022/01/20 14:08:19 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status
On branch blockchain-jwt-token-claims
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   Makefile
        modified:   backend/cmd/cmd.go
        modified:   backend/cmd/server.go
        modified:   backend/main.go
        modified:   backend/rest/api/rest<sub>private.go</sub>
        modified:   backend/ta/token.go

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        taauth.pub.pem

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ rm  taauth.pub.pem
➜  trust-anchor git:(blockchain-jwt-token-claims) ✗ git status                                                                                                                                        <&#x2026;.
➜  trust-anchor git:(blockchain-jwt-token-claims) git status
On branch blockchain-jwt-token-claims
nothing to commit, working tree clean
➜  trust-anchor git:(blockchain-jwt-token-claims) git push origin blockchain-jwt-token-claims 
Enumerating objects: 25, done.
Counting objects: 100% (25/25), done.
Delta compression using up to 8 threads
Compressing objects: 100% (13/13), done.
Writing objects: 100% (13/13), 1.97 KiB | 1008.00 KiB/s, done.
Total 13 (delta 8), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (8/8), completed with 8 local objects.
remote: 
remote: Create a pull request for 'blockchain-jwt-token-claims' on GitHub by visiting:
remote:      <https://github.com/obada-foundation/trust-anchor/pull/new/blockchain-jwt-token-claims>
remote: 
To github.com:obada-foundation/trust-anchor

-   [new branch]      blockchain-jwt-token-claims -> blockchain-jwt-token-claims

➜  trust-anchor git:(blockchain-jwt-token-claims) git checkout readme-document 
Switched to branch 'readme-document'
➜  trust-anchor git:(readme-document) cat README.md 

\## Running OBADA TrustAnchor software

\### Generating EdDSA cryptographical keys

Running software requires two pairs of keys, one is used for private API access and other one is for signing and verifing tokens for OBADA blockchain.

\#### Creating a private API key pair

TODO: add this as a subcommand binary executable

\`\`\`bash
openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta<sub>api.pem</sub>
openssl pkey -in ~/.ssh/ta<sub>api.pem</sub>  -pubout | tee ~/.ssh/ta<sub>api.pub.pem</sub>
\`\`\`

\#### Creating OBADA blockchain key pair

\`\`\`bash
openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta<sub>obada.pem</sub>
openssl pkey -in ~/.ssh/ta<sub>obada.pem</sub>  -pubout | tee ~/.ssh/ta<sub>obada.pub.pem</sub>
\`\`\`

\### Running binary

\`\`\`bash
trust-anchor server &#x2013;auth.public-key-path=~/.ssh/ta<sub>api.pub.pem</sub> &#x2013;auth.private-key-path=~/.ssh/ta<sub>api.pub.pem</sub>
\`\`\`

\## API endpoints
➜  trust-anchor git:(readme-document) git rebase develop
Current branch readme-document is up to date.
➜  trust-anchor git:(readme-document) git checkout develop
Switched to branch 'develop'
Your branch is up to date with 'origin/develop'.
➜  trust-anchor git:(develop) git pull origin develop
remote: Enumerating objects: 25, done.
remote: Counting objects: 100% (25/25), done.
remote: Compressing objects: 100% (5/5), done.
Unpacking objects: 100% (13/13), 2.34 KiB | 1.17 MiB/s, done.
remote: Total 13 (delta 8), reused 12 (delta 8), pack-reused 0
From github.com:obada-foundation/trust-anchor

-   branch            develop    -> FETCH<sub>HEAD</sub>
    79b2461..5cf577a  develop    -> origin/develop

Updating 79b2461..5cf577a
Fast-forward
 Makefile                         |  8 <del><del>++</del></del>&#x2013;
 backend/cmd/cmd.go               |  8 <del><del>+</del></del>&#x2014;
 backend/cmd/server.go            | 18 <del><del><del>++</del></del></del>-----&#x2013;&#x2014;
 backend/main.go                  | 12 <del><del><del>+</del></del></del>&#x2013;&#x2014;
 backend/rest/api/rest<sub>private.go</sub> | 12 <del><del><del><del>++</del></del></del></del>&#x2013;
 backend/ta/token.go              | 28 <del><del><del><del><del><del><del><del><del>+</del></del></del></del></del></del></del></del></del>----&#x2013;&#x2014;
 6 files changed, 55 insertions(+), 31 deletions(-)
➜  trust-anchor git:(develop) git checkout readme-document
Switched to branch 'readme-document'
➜  trust-anchor git:(readme-document) git rebase develop          
Auto-merging Makefile
CONFLICT (content): Merge conflict in Makefile
error: could not apply c9fa76a&#x2026; Readme and spec update
Resolve all conflicts manually, mark them as resolved with
"git add/rm <conflicted<sub>files</sub>>", then run "git rebase &#x2013;continue".
You can instead skip this commit: run "git rebase &#x2013;skip".
To abort and get back to the state before "git rebase", run "git rebase &#x2013;abort".
Could not apply c9fa76a&#x2026; Readme and spec update
➜  trust-anchor git:(5cf577a) ✗ git rebase &#x2013;continue
[detached HEAD 465fc69] Readme and spec update
 2 files changed, 31 insertions(+), 15 deletions(-)
Successfully rebased and updated refs/heads/readme-document.
➜  trust-anchor git:(readme-document) make backend/run                                 
cd backend && go run main.go server \\
        &#x2013;url <http://localhost> \\
        &#x2013;db-path=../obada/trust-anchor/data \\
        &#x2013;auth.private-key-path=/Users/sr/.ssh/ta<sub>api.pem</sub> \\
        &#x2013;auth.public-key-path=/Users/sr/.ssh/ta<sub>api.pub.pem</sub>
TrustAnchor unknown
(c)OBADA 2022

badger 2022/01/20 14:48:50 INFO: All 0 tables opened in 0s
badger 2022/01/20 14:48:50 INFO: Discard stats nextEmptySlot: 0
badger 2022/01/20 14:48:50 INFO: Set nextTxnTs to 0
badger 2022/01/20 14:48:50 INFO: Deleting empty file: ../obada/trust-anchor/data/000076.vlog
 :: TRUST-ANCHOR  :: 2022/01/20 14:48:50.478121 server.go:55: [INFO] starting server on port :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:48:50.482469 rest.go:47: [INFO] activate http rest server on :80
 :: TRUST-ANCHOR  :: 2022/01/20 14:48:53.807659 logger.go:161: "POST <http://localhost/api/v1/verify> HTTP/1.1" from [::1]:53845 - 200 22B in 809.61µs
^C2022/01/20 14:50:48 [WARN] interrupt signal
2022/01/20 14:50:48 [INFO] shutdown initiated
 :: TRUST-ANCHOR  :: 2022/01/20 14:50:48.901839 rest.go:148: [WARN] shutdown rest server
 :: TRUST-ANCHOR  :: 2022/01/20 14:50:48.902035 rest.go:55: [WARN] http server terminated, http: Server closed
badger 2022/01/20 14:50:48 INFO: Lifetime L0 stalled for: 0s
 :: TRUST-ANCHOR  :: 2022/01/20 14:50:48.903280 rest.go:157: [DEBUG] shutdown http server completed
badger 2022/01/20 14:50:48 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
make: **\*** [backend/run] Error 1

➜  trust-anchor git:(readme-document) ✗ git status
On branch readme-document
Changes not staged for commit:
  (use "git add <file>&#x2026;" to update what will be committed)
  (use "git restore <file>&#x2026;" to discard changes in working directory)
        modified:   docs/Spec.org

no changes added to commit (use "git add" and/or "git commit -a")
➜  trust-anchor git:(readme-document) ✗ git add docs/
➜  trust-anchor git:(readme-document) ✗         git status
On branch readme-document
Changes to be committed:
  (use "git restore &#x2013;staged <file>&#x2026;" to unstage)
        modified:   docs/Spec.org

Untracked files:
  (use "git add <file>&#x2026;" to include in what will be committed)
        do.md

➜  trust-anchor git:(readme-document) ✗ rm do
rm: do: No such file or directory
➜  trust-anchor git:(readme-document) ✗ rm do.md 
➜  trust-anchor git:(readme-document) ✗ 

