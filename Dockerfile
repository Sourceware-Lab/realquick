FROM golang:1.23.47@sha256:70031844b8c225351d0bb63e2c383f80db85d92ba894e3da7e13bcf80efa9a37 AS gobase

ENV GOFLAGS="-buildvcs=false"

ARG UID=1000
ARG GID=$UID
ARG USERNAME=nonroot
ENV WORKDIR=/app


RUN addgroup --gid $GID $USERNAME && \
    adduser --uid $UID --gid $GID --disabled-password --gecos "" $USERNAME

WORKDIR $WORKDIR

RUN chown -R $USERNAME $WORKDIR

USER $USERNAME

FROM gobase AS install

COPY src/backend/go.mod src/backend/go.sum ./
RUN go mod download


FROM install AS local

USER root
USER $USERNAME

RUN go install github.com/air-verse/air@latest
COPY src/backend .


FROM install AS build
COPY src/backend .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend

# docker pull gcr.io/distroless/static-debian12:nonroot
FROM gcr.io/distroless/static-debian12@sha256:6cd937e9155bdfd805d1b94e037f9d6a899603306030936a3b11680af0c2ed58 AS production

COPY --from=build /app/backend /backend
EXPOSE 8080

ENTRYPOINT ["/backend"]
