import { createReadStream, statSync } from "node:fs"
import { createServer } from "node:http"
import { extname, join, normalize, resolve, sep } from "node:path"
import { fileURLToPath } from "node:url"

const projectRoot = resolve(fileURLToPath(new URL("..", import.meta.url)))
const root = resolve(join(projectRoot, "dist"))
const fixtureRoot = resolve(join(projectRoot, "tests", "fixtures"))
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
  const base = normalized.startsWith("/tests/fixtures/") ? fixtureRoot : root
  const relativePath = normalized.startsWith("/tests/fixtures/")
    ? normalized.slice("/tests/fixtures/".length)
    : normalized
  let filePath = resolve(join(base, relativePath))

  if (![root, fixtureRoot].some((allowed) => filePath === allowed || filePath.startsWith(allowed + sep))) {
    response.writeHead(403)
    response.end("Forbidden")
    return
  }

  try {
    let stat = statSync(filePath)
    if (stat.isDirectory()) {
      filePath = join(filePath, "index.html")
      stat = statSync(filePath)
    }
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
