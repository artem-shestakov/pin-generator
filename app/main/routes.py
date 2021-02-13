from flask import Blueprint, render_template, abort
from jinja2 import TemplateNotFound
from flask import current_app, request

# Create blueprint
main_blueprint = Blueprint(
    "main",
    __name__,
    template_folder="../templates/main",
)


# Route for index page
@main_blueprint.route("/")
def index():
    try:
        current_app.logger.info(f"{request.method} Request '/' url from {request.remote_addr}")
        return render_template("index.html", title="Main page")
    except TemplateNotFound:
        abort(404)
        current_app.logger.error(f"TemplateNotFound exception")
