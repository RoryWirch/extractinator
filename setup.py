from setuptools import setup
setup(
  name = 'extractinator',
  packages = ['extractinator'],
  version = '0.1',
  license='MIT',
  description = 'Simple tool for extracting bundle data from containers using gRPC',
  author = 'Rory Wirch',
  url = 'https://github.com/RoryWirch/extractinator',
  install_requires=[
          #TODO: 'here',
      ],
  classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
)