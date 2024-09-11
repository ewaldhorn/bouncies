# Bouncies

A simple game written in Go and Ebitengine.

## Technologies

I used a variety of technologies to make this all happen:

- Go: <https://go.dev/>
- Ebitengine: <https://ebitengine.org/>
- Task: <https://taskfile.dev/>

### Tasks

For convenience, I have a Task file. This helps save me from having to remember commands. Instead, Task does that for me!

Run `task` to see a list of available tasks. Some are:

| Task            | Action taken                                                                    |
| --------------- | ------------------------------------------------------------------------------- |
| buildproduction | Builds a production WASM binary, minus debug information                        |
| buildwasm       | Builds the WASM project                                                         |
| clean           | Removes the './bin/' folder                                                     |
| default         | Lists available tasks                                                           |
| run             | Runs the files in the ./bin/ folder as-is, no build steps envoked. On port 9000 |
| runwasm         | Runs the project in WASM on port 9000                                           |
| setupexecjs     | Copies the wasm_exec.js and html files to the './bin/' folder                   |
| test            | Runs the desktop project                                                        |
