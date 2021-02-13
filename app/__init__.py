from flask import Flask


# create Flask app object and init all modules
def create_app(config_object):
    from .main import create_module as main_create_module
    from app.api.v1 import create_module as api_v1_create_module

    # Init APP
    app = Flask(__name__)
    app.config.from_object(config_object)

    # Init modules
    main_create_module(app)
    app.logger.info("Init Main module")

    app.logger.info('Initializing API v1 module')
    api_v1_create_module(app)
    app.logger.info('API v1 module initialized')

    return app
