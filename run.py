from app import create_app
import os

config_object = os.environ.get("WEBAPP_ENV", "dev")
app = create_app(f"config.{config_object.capitalize()}Config")

if __name__ == '__main__':
    app.logger.info("Starting application")
    app.run(host="0.0.0.0")
