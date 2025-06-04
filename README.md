# repo2graph
Library that tries to approach a graph simulation of a repository ecosystem

## Build and run

1. Install Go (1.21+ recommended) and clone this repository.
2. Download dependencies using Go modules:
   ```bash
   go mod tidy
   ```
   The available CLI commands are defined in `cmd/commands`.
3. Build the command-line tool (the `-o` flag avoids a name clash with the `cmd`
   directory):
   ```bash
   go build -o repo2graph ./cmd
   ```
4. Run the CLI directly from sources or using the built binary:
   ```bash
   go run ./cmd server      # run without building
   # or
   ./repo2graph server      # using the compiled binary
   ```

## Using the server

1. Start the server using one of the commands above. It listens on `:8080` by default.
2. Open `http://localhost:8080` in your browser.
3. Enter a GitHub token in the text field and click **Reload**. The token must have access to the repositories you want to analyse.
4. The web interface will display a graph rendered with **vis-network**. A list of repository names with checkboxes allows you to toggle each node. The canvas will update to show the graph of repositories and their relationships.
