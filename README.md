```regex
███████╗██╗███╗   ███╗██████╗ ██╗     ███████╗    ██╗   ██╗██╗   ██╗██╗     ███╗   ██╗███████╗██████╗  █████╗ ██████╗ ██╗     ███████╗    ███████╗ ██████╗ █████╗ ███╗   ██╗███╗   ██╗███████╗██████╗     
██╔════╝██║████╗ ████║██╔══██╗██║     ██╔════╝    ██║   ██║██║   ██║██║     ████╗  ██║██╔════╝██╔══██╗██╔══██╗██╔══██╗██║     ██╔════╝    ██╔════╝██╔════╝██╔══██╗████╗  ██║████╗  ██║██╔════╝██╔══██╗    
███████╗██║██╔████╔██║██████╔╝██║     █████╗      ██║   ██║██║   ██║██║     ██╔██╗ ██║█████╗  ██████╔╝███████║██████╔╝██║     █████╗      ███████╗██║     ███████║██╔██╗ ██║██╔██╗ ██║█████╗  ██████╔╝    
╚════██║██║██║╚██╔╝██║██╔═══╝ ██║     ██╔══╝      ╚██╗ ██╔╝██║   ██║██║     ██║╚██╗██║██╔══╝  ██╔══██╗██╔══██║██╔══██╗██║     ██╔══╝      ╚════██║██║     ██╔══██║██║╚██╗██║██║╚██╗██║██╔══╝  ██╔══██╗    
███████║██║██║ ╚═╝ ██║██║     ███████╗███████╗     ╚████╔╝ ╚██████╔╝███████╗██║ ╚████║███████╗██║  ██║██║  ██║██████╔╝███████╗███████╗    ███████║╚██████╗██║  ██║██║ ╚████║██║ ╚████║███████╗██║  ██║    
╚══════╝╚═╝╚═╝     ╚═╝╚═╝     ╚══════╝╚══════╝      ╚═══╝   ╚═════╝ ╚══════╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝    ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝    
                                                                                                                                                                                                          
```
[![Cybersecurity Projects]
[![Python](https://img.shields.io/badge/Python-3.13+-3776AB?style=flat&logo=python&logoColor=white)](https://www.python.org)
[![PyPI](https://img.shields.io/pypi/v/dnslookup-cli?color=3775A9&logo=pypi&logoColor=white)](https://pypi.org/project/dnslookup-cli/)


## What It Does

- Scans pyproject.toml and requirements.txt for known CVEs via OSV.dev
- Updates all Python dependencies to latest stable versions in one command
- Parallel queries against PyPI with local ETag caching for speed
- Full PEP 440 version parsing with automatic pre-release filtering
- Comment-preserving file updates that keep your formatting intact
- Configurable via .angela.toml or [tool.angela] in pyproject.toml
