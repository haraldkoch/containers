import sys
import os
sys.path.append(os.path.abspath("./scripts"))

from apksearch import apksearch

def get_latest(_: str) -> str | None:
    """Get the latest version of rsync."""
    return apksearch("rsync")

if __name__ == "__main__":
    version = get_latest("stable")
    print(f"{version}")
