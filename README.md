# Flowstate

Get into your productive flow state. A beautiful terminal todo app that uses simple markdown files.

## Install

```bash
go install github.com/adamaugustinsky/flowstate@latest
```

Make sure `~/go/bin` is in your PATH.

## Usage

```bash
flowstate      # Local todos in current directory
flowstate -g   # Global todos in ~/config/flowstate/
```

Your todos are saved in a `TODO.md` file that's created when you add your first task.

### Keys

**Main view**
- `a` - add todo
- `e` - edit todo  
- `space` - mark done
- `d` - delete
- `enter` - see details
- `q` - quit

**While editing**
- `enter` - next/save
- `esc` - cancel

## Build from source

```bash
git clone https://github.com/adamaugustinsky/flowstate
cd flowstate
go build
```

Requires Go 1.21+

## License

MIT
