# Space Invaders in Go

This is a simple implementation of the classic game Space Invaders, using Go and the [Raylib](https://www.raylib.com/) library.

## Instalation

1. Ensure you have Go and [Raylib Go](https://github.com/gen2brain/raylib-go) installed in your system. You might also need to have a C compiler installed in your system. GCC is a safe bet.

2. Clone the repo:

```bash
git clone https://github.com/dn46/go-space-invaders
```

3. Navigate to the directory:
```bash
cd go-space-invaders
```

4. Install the dependencies:

```bash
go mod tidy
```

5. Run the Makefile:

```bash
make run
```

## Controls

| Action | Key |
| --- | --- |
| Move Left | `A` |
| Move Right | `D` |
| Shoot | `Spacebar` |
