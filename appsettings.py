import os


class AppSettings:
    def __init__(self):
        self.token = ""
        self.app_id = ""
        self.supergoon_games_server_id = ""
        self.github_access_token = ""


application_settings = AppSettings()


def initialize_app_settings():
    application_settings.token = os.getenv("DISCORD_BOT_TOKEN", "")
    application_settings.app_id = os.getenv("DISCORD_APP_ID", "")
    application_settings.supergoon_games_server_id = os.getenv("DISCORD_SUPERGOON_GUILD_ID", "")
    application_settings.github_access_token = os.getenv("GITHUB_ACCESS_TOKEN", "")
