import logging
import threading

from flask import Flask

from goon_http.handlers.webhook_handler import webhook_bp
from goon_http.handlers.sgg_discord_bot_handler import sgg_discord_bot_bp

log = logging.getLogger(__name__)

app = Flask(__name__)
app.register_blueprint(webhook_bp)
app.register_blueprint(sgg_discord_bot_bp)


def start_server():
    log.info("Starting server..")
    thread = threading.Thread(target=lambda: app.run(host="0.0.0.0", port=80), daemon=True)
    thread.start()
