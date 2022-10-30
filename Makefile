kata:
	mkdir ${KATA}
	touch ${KATA}/.gitkeep

install:
	pipenv install --dev

test:
	pipenv run bash -c "cd ${KATA} && \
	pytest -f -vvv --cov=. --cov-branch --cov-report=term-missing:skip-covered --cov-report=xml:../cov.xml"
