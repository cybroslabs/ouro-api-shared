import os, pkgutil

# Get the absolute path of this package
package_dir = os.path.dirname(__file__)

# Dynamically find all subdirectories inside platform_api
__all__ = [name for _, name, is_pkg in pkgutil.iter_modules([package_dir]) if is_pkg]

# Import all discovered subpackages dynamically
for module in __all__:
    __import__(f"{__name__}.{module}")
