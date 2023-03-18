export GOARCH=386
TARGET ?= ./mfc/mfc.exe
all:
	@go build -ldflags "-s -w" -o $(TARGET) main.go

test: all
	@$(TARGET) uninstall || true
	@$(TARGET) install || true
	@$(TARGET) start || true

clean:
	@rm -f $(TARGET)
