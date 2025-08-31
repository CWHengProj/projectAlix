# Alix , the manager of aliases
## Features

- Add, update, delete, and list shell aliases
- Supports subcommands and aliases for usability
- Detailed listing of aliases with ```ls detailed```

---
## Installation

### Download Pre-built Binary (Recommended)

1. Download the latest release from [GitHub Releases](https://github.com/CWHengProj/projectAlix/releases).
2. Move the binary to `/usr/local/bin`:
   ```sh
   sudo mv <name-of-downloaded-binary> alix
   sudo mv alix /usr/local/bin/
   sudo chmod +x /usr/local/bin/alix
   ```
3. Run `alix` from anywhere:
   ```sh
   alix add myalias "echo hello"
   ```

### Build from Source

1. Install [Go](https://golang.org/doc/install).
2. Clone the repository:
   ```sh
   git clone https://github.com/CWHengProj/projectAlix.git
   cd projectAlix
   ```
3. Build the binary:
   ```sh
   go build -o alix
   ```
4. Move the binary to `/usr/local/bin`:
   ```sh
   sudo mv alix /usr/local/bin/
   ```

## Usage

```sh
alix add myalias "echo hello"
alix ls
alix delete myalias
```
#### Make your changes reflect in your .shrc file
- linux (bash)
```sh
source ~/.bashrc
```
- MacOS (zsh)
```sh
source ~/.zshrc
```
- (Protip) You can turn the source command into an alias as well to optimize your workflow!
---

## How to Contribute

1. **Fork the repo** and clone it locally:
   ```sh
   git clone https://github.com/CWHengProj/projectAlix.git
   cd projectAlix
   ```
2. Create a branch for your changes:
    ```sh 
    git checkout -b feature/my-feature
    ```
3. Make your changes and test them.
4. Push your branch:
    ```sh
    git push origin feature/my-feature
    ```
5. Open a Pull Request on GitHub with a short description of your changes.
### Guidelines
- Keep commits small and descriptive.
- Write clear code and follow Go conventions (gofmt).
- Be respectful in discussions.
- Update your branch from the main branch to avoid conflicts.

### Troubleshooting
- Share your logs
```sh
   ~/.alix/alix.log
```
## License
GNU GPL Version 3.0