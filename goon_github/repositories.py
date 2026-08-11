import logging

from github.Repository import Repository

from goon_github.github_client import github_client

log = logging.getLogger(__name__)


def get_all_repos() -> list[Repository]:
    try:
        repos = list(github_client.get_user("kjblanchard").get_repos())
        return repos
    except Exception as e:
        log.error(f"Error listing repositories: {e}")
        return []
