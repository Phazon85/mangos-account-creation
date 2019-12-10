build-docker:
		docker build -t dgparker/mangos-account-creation:latest -f ./build/package/Dockerfile .

push-docker:
		docker push dgparker/mangos-account-creation:latest