rapyd - Run A Python in Docker
=============================

``rapyd`` is a simple CLI to make developing using docker easier. It helps solve the problem of running one off tasks inside existing containers. For example, you entire Django development environment it defined in a docker-compose file and you wish to run management tasks.

Building
--------

``rapyd`` is written in Go, so you will need the Go compile installed. Then you can build against all support architectures with:

.. code-block:: bash

	make build

The resulting executables can be found in the ``build`` directory named by platform and architecture.

You can also build just a single executable for a specific platform using:

.. code-block:: bash

	make linux
	make macos
	make windows


Installation
------------

Once you have the executable for your platform, you will need to either place it in your projects directory or somewhere accessible on the ``$PATH``.


Usage
-----

You should run ``rapyd`` within your project directory as the parent directory name is used to determine the name of the container to run the passed in command within.

.. code-block:: bash

	cd a_project
	rapyd ./manage.py migrate


The above command would execute ``manage.py migrate`` within the docker container named ``aproject_web_1``. The name is derived from what ``docker-compose`` would name a service called ``web``.

The virtual environment will be activated automatically, so you do not need to do it.


Brain Dump Log
--------------

2017-12-20:
    The way in which the container name is determined is pretty crappy, something smarter should be used. Whether it is a config (ini) file, command line switch doesn't really matter. Just needs to be a bit smarter. Perhaps leaving the current approach as a fallback.
