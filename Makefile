TEST_SERVER_PATH 	= ./tests/testServer
TEST_SERVER_TARGET 	= testServer

TEST_CLIENT_PATH 	= ./tests/testClient
TEST_CLIENT_TARGET 	= $(TEST_CLIENT_PATH)/test.py

APP_PATH		= ./app
WEB_PATH		= ./web
UTILS_PATH		= ./utils

.PHONY:install
install: install_web install_utils install_app

.PHONY:install_app
install_app:
	@echo "Insalling app ..."
	@go install $(APP_PATH)

.PHONY:install_web
install_web:
	@echo "Insalling web ..."
	@go install $(WEB_PATH)

.PHONY:install_utils
install_utils:
	@echo "Insalling utils ..."
	@go install $(UTILS_PATH)

.PHONY:test
test: install
	@go install $(TEST_SERVER_PATH) 
	@$(TEST_SERVER_TARGET)&
	@python3 $(TEST_CLIENT_TARGET)
	
