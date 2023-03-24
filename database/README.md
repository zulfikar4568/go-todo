Make Plan
```bash
make migration-plan name=initial_setup POSTGRES_DEFAULT_URL="postgres://zulfikar:123@localhost:5450/todo_service?sslmode=disable" POSTGRES_DEFAULT_TEST_URL="postgres://zulfikar:123@localhost:5450/todo_service_test?sslmode=disable"
```

Make Apply Migration
```bash
make migration-apply POSTGRES_DEFAULT_URL="postgres://zulfikar:123@localhost:5450/todo_service?sslmode=disable" POSTGRES_DEFAULT_TEST_URL="postgres://zulfikar:123@localhost:5450/todo_service_test?sslmode=disable"
```

Clean a Migration
```bash
make migration-clean POSTGRES_DEFAULT_URL="postgres://zulfikar:123@localhost:5450/todo_service?sslmode=disable" POSTGRES_DEFAULT_TEST_URL="postgres://zulfikar:123@localhost:5450/todo_service_test?sslmode=disable"
```