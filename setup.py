from setuptools import setup
from setuptools import find_packages

about = {}
with open('__version__.py') as f:
    exec(f.read(), about)

REQUIREMENTS = [
    'click==7.1.2',
]

setup(name='aaec',
      python_requires='>=3.8.5',
      version=about['__version__'],
      packages=find_packages(),
      entry_points={'console_scripts': ['aaec=aaec.main:cli']},
      author_email='greyschwinger@gmail.com',
      install_requires=REQUIREMENTS,
      zip_safe=False)
