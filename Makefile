build-docker:
		docker build -t phazon85/mangos-account-creation:latest -f ./build/package/Dockerfile .

push-docker:
		docker push phazon85/mangos-account-creation:latest