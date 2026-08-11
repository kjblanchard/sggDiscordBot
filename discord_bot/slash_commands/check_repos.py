import logging

import discord
from discord import app_commands

from discord_bot.discord_bot import SggBot
from goon_github.repositories import get_all_repos

log = logging.getLogger(__name__)

SYSOP_ROLE_ID = 907314874101665823


def check_if_user_in_role(role_to_check_for: int, roles: list[discord.Role]) -> bool:
    for role in roles:
        if role.id == role_to_check_for:
            return True
    return False


def add_check_repos_slash_command(bot: SggBot, guild: discord.Object):
    @bot.tree.command(name="check-repos", description="Gets a list of 10 of the repos by kjb, starting at the first", guild=guild)
    @app_commands.describe(start_number="The number to start listing repositories from")
    async def check_repos(interaction: discord.Interaction, start_number: int = 0):
        assert isinstance(interaction.user, discord.Member)
        roles = interaction.user.roles
        if not check_if_user_in_role(SYSOP_ROLE_ID, roles):
            await interaction.response.send_message("Only Sysop can perform this action")
            return

        await interaction.response.defer()
        repos = get_all_repos()
        end = start_number + 10

        embed = discord.Embed(title="Repos")
        for i in range(start_number, min(end, len(repos))):
            repo = repos[i]
            embed.add_field(name=repo.name, value=repo.default_branch or "main", inline=False)

        await interaction.followup.send(content="Heres a list of 10 the Repos!", embed=embed)
