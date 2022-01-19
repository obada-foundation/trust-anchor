backend/compile:
	cd backend && go build -o trust-anchor

backend/run:
	cd backend && go run main.go server --db-path=../obada/trust-anchor/data

backend/vendor:
	cd backend && go mod tidy && go mod vendor

backend/lint:
	cd backend && golangci-lint --config .golangci.yml run --print-issued-lines --out-format=github-actions ./...

backend/docker/build:

backend/docker/publish:

frontend/build: frontend/compile frontend/minify frontend/checksum frontend/clean frontend/embed

frontend/compile:
	cd frontend && elm make src/Main.elm --optimize --output=./build/app.js

frontend/minify:
	cd frontend && uglifyjs ./build/app.js --compress 'pure_funcs=[F2,F3,F4,F5,F6,F7,F8,F9,A2,A3,A4,A5,A6,A7,A8,A9],pure_getters,keep_fargs=false,unsafe_comps,unsafe' | uglifyjs --mangle --output ./build/app.min.js

frontend/checksum:
	cd frontend && mv ./build/app.min.js ./build/app.$$(sha512sum ./build/app.min.js | head -c 10).min.js

frontend/clean:
	cd frontend && rm ./build/app.js

APP_FILE = $(shell ls ./build | grep app)
frontend/embed:
	cd frontend && cp index.html ./build
	sed -i "s/app.min.js/$(APP_FILE)/" ./build/index.html

frontend/test:
	cd frontend && elm-test
