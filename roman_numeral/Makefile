install:
	cd ../
	pipenv install

test:
	cd ../
	pipenv run pytest -f --ff --cov=. --cov-report=term-missing:skip-covered --cov-report=xml:../cov.xml
