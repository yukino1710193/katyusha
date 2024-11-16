gen-junbanmachi-messages:
	protoc -I=. --go_out=. pkg/junbanmachi/messages.proto

gen-fukabunsan-messages:
	protoc -I=. --go_out=. pkg/fukabunsan/messages.proto

gen-outoushuugou-messages:
	protoc -I=. --go_out=. pkg/outoushuugou/messages.proto
