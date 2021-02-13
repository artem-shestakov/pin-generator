# Patching standard socket for gevent
from gevent import monkey; monkey.patch_socket()
from gevent.pywsgi import WSGIServer
from app import create_app
import os

# Get WEBAPP variable and create app object
config_object = os.environ.get("WEBAPP", "dev")
app = create_app(f"config.{config_object.capitalize()}Config")

# Run Gevent server
if __name__ == '__main__':
    http_server = WSGIServer(("0.0.0.0", 5000), application=app)
    http_server.serve_forever()