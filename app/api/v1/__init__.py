def create_module(app):
    from .routes import api_v1_bp

    app.register_blueprint(api_v1_bp)