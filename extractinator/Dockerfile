FROM python:3
ADD extractinator.py /
RUN pip install grpcio
RUN pip install grpcio-reflection
CMD ["python", "./extractinator.py"]