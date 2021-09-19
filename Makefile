install:
	pipenv install --dev

test:
	pipenv run bash -c "cd ${KATA} && \
	pytest -f --ff --cov=. --cov-branch --cov-report=term-missing:skip-covered --cov-report=xml"
