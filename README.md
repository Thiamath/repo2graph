# repo2graph
Library that tries to approach a graph simulation of a repository ecosystem

## Build and run

1. Ensure the Go toolchain is installed.
2. Build the command line tool:
   ```bash
   go build -o repo2graph ./cmd
   ```
3. Alternatively, run it directly from the sources:
   ```bash
   go run ./cmd server
   ```
   or execute the built binary:
   ```bash
   ./repo2graph server
   ```

## Using the server

1. Start the server using one of the commands above. It listens on `:8080` by default.
2. Open `http://localhost:8080` in your browser.
3. Enter a GitHub token in the text field and click **Reload**. The token must have access to the repositories you want to analyse.
4. The web interface will display a graph rendered with **vis-network**. A list of repository names with checkboxes allows you to toggle each node. The canvas will update to show the graph of repositories and their relationships.
