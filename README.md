# ONpK-Demo
Demo of the TF+Helm+Docker+Simple voting app. Project sponsor: MaaS


V DockerFiles je možné nájsť multistage buildy pre frontend aj backend aplikácie.
Tieto buildy sa automaticky stavajú cez Github Actions, a sú verejne dostupné na https://hub.docker.com/r/msprg. Buildy zatiaľ nie sú verzované, teoreticky najľahšie by bolo len pridať kúsok hashu pre každý build, ale pre demo snáď stačí.

V adresári TF sú patričné terraform súbory a skripty na IaaC a inštaláciu / rozbehnutie minikubes.


V podadresári workflows je možné nájsť yaml súbory pre GH actions, ktoré sú zodpovedné za build a push docker images, a aplikáciu terraformu on-demand, s tfstate súborom uloženým v hashicorp cloude.
