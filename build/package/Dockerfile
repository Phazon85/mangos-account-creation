#############################
### STEP: 1 build binary  ###
#############################
FROM golang:1.12.4 AS builder
WORKDIR /workspace/
COPY . .
RUN GOOS=linux GOARCH=386 GO111MODULE=on go build -o mangos-account-creation cmd/mangos-account-creation/main.go

###########################################
### STEP 2: use small image for deploy  ###
###########################################

FROM scratch
COPY --from=builder /workspace/mangos-account-creation /bin/mangos-account-creation
ENTRYPOINT [ "/bin/mangos-account-creation" ]