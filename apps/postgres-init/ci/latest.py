import sys
import os
sys.path.append(os.path.abspath(".github/scripts"))

from apksearch import apksearch

def get_latest(_: str) -> str | None:
    """Get the latest version of the PostgreSQL client."""
    return apksearch("postgresql16-client")

if __name__ == "__main__":
    version = get_latest("stable")
    print(f"{version}")
