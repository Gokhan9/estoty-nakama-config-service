FROM heroiclabs/nakama-pluginbuilder:3.22.0 AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=1

WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download

COPY nakama/modules ./nakama/modules
RUN go build --trimpath --mod=mod --buildmode=plugin -o /backend/backend.so ./nakama/modules

FROM heroiclabs/nakama:3.22.0

COPY --from=builder /backend/backend.so /nakama/data/modules/backend.so
COPY config/game_config.json /nakama/data/config/game_config.json
