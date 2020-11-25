from http.server import SimpleHTTPRequestHandler
import socketserver

class HttpRequest(SimpleHTTPRequestHandler):
    def get(self):
        if self.path == '/':
            self.path = 'index.html'
        return SimpleHTTPRequestHandler

handler = HttpRequest

port = 3000
server = socketserver.TCPServer(("", port), handler)

server.serve_forever()
