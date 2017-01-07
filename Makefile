TEST_SERVER_PATH 	= ./tests/testServer
TEST_SERVER_TARGET 	= testServer

TEST_CLIENT_PATH 	= ./tests/testClient
TEST_CLIENT_TARGET 	= $(TEST_CLIENT_PATH)/test.py


.PHONY:test
test:
	@go install $(TEST_SERVER_PATH) 
	@$(TEST_SERVER_TARGET)&
	@python3 $(TEST_CLIENT_TARGET)
	
