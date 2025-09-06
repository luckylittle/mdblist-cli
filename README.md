export MDBLIST_API_KEY="your_actual_api_key_here"

go mod tidy

# Get help
go run . --help

# Get your API limits
go run . get my-limits

# Get your lists
go run . get my-lists
