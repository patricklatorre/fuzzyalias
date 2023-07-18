Simple HTTP server that redirects to the nearest matching alias. (see [config](https://github.com/patricklatorre/fuzzyalias/blob/main/config.json.example))

- `srv.io/clipnotes` → `clipnotes.cc`
- `srv.io/clip` → `clipnotes.cc`
- `srv.io/pnotes` → `clipnotes.cc`

Also forwards any string beyond the alias.

- `srv.io/gh/fuzzyalias` → `github.com/patricklatorre/fuzzyalias`
- `srv.io/clip/https://youtu.be/dQw4w9WgXcQ` → `clipnotes.cc/https://youtu.be/dQw4w9WgXcQ`

The config is loaded once on server start.
