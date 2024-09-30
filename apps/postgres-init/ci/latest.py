import sys
import os
sys.path.append(os.path.abspath(".github/scripts"))

from apksearch import get_latest_version

if __name__ == "__main__":
    version = get_latest_version("postgresql16-client")
    print(f"{version}")
