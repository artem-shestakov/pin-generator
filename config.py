class Config(object):
    POSTS_PER_PAGE = 10
    PREFERRED_URL_SCHEME = "https"
    PIN_LEN = 8
    SALT_LEN = 10


class ProdConfig(Config):
    DEBUG = False
    ENV = "production"


class DevConfig(Config):
    DEBUG = True

