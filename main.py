import asyncio
import logging
import signal

from appsettings import initialize_app_settings, application_settings
from discord_bot.discord_bot import initialize_discord, open_discord_connection, close_discord
from discord_bot.slash_commands_setup import add_all_slash_commands
from goon_github.github_client import initialize_github
from goon_http.goon_http import start_server

logging.basicConfig(level=logging.INFO)
log = logging.getLogger(__name__)


async def main():
    initialize_app_settings()
    bot = initialize_discord(
        application_settings.token,
        application_settings.app_id,
        application_settings.supergoon_games_server_id,
    )
    initialize_github(application_settings.github_access_token)
    add_all_slash_commands(bot)
    loop = asyncio.get_event_loop()
    stop = loop.create_future()
    loop.add_signal_handler(signal.SIGINT, stop.set_result, None)
    await open_discord_connection(bot)
    start_server()
    log.info("Press Ctrl+C to exit")
    await stop
    await close_discord(bot)


if __name__ == "__main__":
    asyncio.run(main())
