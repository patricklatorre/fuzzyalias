Simple HTTP server redirects requests to the nearest matching alias. (see [config](https://github.com/patricklatorre/fuzzyalias/blob/main/config.json.example))

- `https://srv.io/clipnotes` → `https://clipnotes.cc`
- `https://srv.io/clip` → `https://clipnotes.cc`
- `https://srv.io/pnotes` → `https://clipnotes.cc`

Also forwards any string beyond the alias.

- `https://srv.io/gh/fuzzyalias` → `https://github.com/patricklatorre/fuzzyalias`
- `https://srv.io/clip/https://youtu.be/dQw4w9WgXcQ` → `https://clipnotes.cc/https://youtu.be/dQw4w9WgXcQ`

