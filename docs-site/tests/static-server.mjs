import { createReadStream, statSync } from "node:fs"
import { createServer } from "node:http"
import { extname, join, normalize, resolve, sep } from "node:path"
import { fileURLToPath } from "node:url"

const root = resolve(fileURLToPath(new URL("..", import.meta.url)))
const port = Number(process.env.PORT || 4173)

const types = {
  ".css": "text/css; charset=utf-8",
  ".html": "text/html; charset=utf-8",
  ".js": "text/javascript; charset=utf-8",
  ".json": "application/json; charset=utf-8",
}

createServer((request, response) => {
  const url = new URL(request.url || "/", `http://${request.headers.host}`)
  const pathname = decodeURIComponent(url.pathname)
  const normalized = normalize(pathname).replace(/^(\.\.[/\\])+/, "")
  const filePath = resolve(join(root, normalized))

  if (filePath !== root && !filePath.startsWith(root + sep)) {
    response.writeHead(403)
    response.end("Forbidden")
    return
  }

  try {
    const stat = statSync(filePath)
    if (!stat.isFile()) throw new Error("not a file")
    response.writeHead(200, {
      "content-type": types[extname(filePath)] || "application/octet-stream",
    })
    createReadStream(filePath).pipe(response)
  } catch {
    response.writeHead(404)
    response.end("Not found")
  }
}).listen(port, "127.0.0.1", () => {
  console.log(`static test server listening on http://127.0.0.1:${port}`)
})

