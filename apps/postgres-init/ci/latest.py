import sys
import os
sys.path.append(os.path.abspath(".github/scripts"))

from apksearch import apksearch

if __name__ == "__main__":
    version = apksearch("postgresql16-client")
    print(f"{version}")
