# ----- 第一階段：建置 (Builder) -----
# 使用官方的 Golang 映像檔作為建置環境
FROM golang:1.24-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 檔案，以便下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製所有專案原始碼
COPY . .

# 進行編譯
# CGO_ENABLED=0  : 產生靜態連結的二進位檔，不依賴 C 函式庫
# -o cloud-butler: 指定輸出的執行檔名稱
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o cloud-butler main.go


# ----- 第二階段：執行 (Final) -----
# 使用一個極度輕量級的 Alpine Linux 作為最終的執行環境
# Alpine 包含了一些基礎工具和 SSL 憑證，是執行 Go 程式的好選擇
FROM alpine:latest

# 設定工作目錄
WORKDIR /app

# 只從第一階段 (builder) 複製編譯好的執行檔到最終映像檔中
# 這是多階段建置的精髓，讓最終映像檔保持極小體積
COPY --from=builder /app/cloud-butler .

# 設定容器啟動時要執行的指令
# 這裡我們讓它預設執行 --help，顯示使用說明
ENTRYPOINT ["./cloud-butler"]
CMD ["--help"]
