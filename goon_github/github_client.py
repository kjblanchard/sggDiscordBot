from github import Github

github_client: Github = None


def initialize_github(token: str):
    global github_client
    github_client = Github(token)
